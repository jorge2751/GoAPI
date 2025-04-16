# Go API Project Notes

This document provides a detailed breakdown of our Go API project for someone learning Go. It covers the project structure, steps to set up the environment, and commands to test and run the API.

## Project Overview

This is a simple REST API built with Go's standard library that:
- Serves a JSON response at the `/hello_world` endpoint
- Serves random quotes at the `/quotes/random` endpoint
- Uses proper HTTP status codes and content types
- Handles errors appropriately
- Can be configured via environment variables
- Has been set up for deployment on Render
- Features request logging middleware

## Project Structure

```
GoAPI/
├── cmd/
│   └── api/
│       └── main.go         # Application entry point
├── internal/               # Private application code
│   └── api/
│       ├── data/           # Data models and services
│       │   └── quotes.go   # Quote data service
│       ├── middleware/     # HTTP middleware
│       │   └── logging.go  # Logging middleware
│       └── routes/         # HTTP route handlers
│           ├── routes.go   # Route registration
│           └── quotes.go   # Quote-related handlers
├── test/                   # Test files
│   ├── quotes_test.go      # Tests for quote routes
│   └── routes_test.go      # Tests for hello world route
├── pkg/                    # Public libraries
├── .gitignore              # Git ignore file
├── go.mod                  # Go module definition
├── README.md               # Project documentation
├── render.yaml             # Render deployment configuration
└── TASK.md                 # Task list for the project
```

This structure follows the standard Go project layout, which is a common convention in the Go community.

## Environment Setup

Here are the steps used to set up the Go environment:

1. **Install Go**:
   ```
   brew install go
   ```

2. **Verify Installation**:
   ```
   go version
   ```

3. **Initialize Go Module**:
   ```
   go mod init github.com/jorge2751/GoAPI
   ```

4. **Create Project Structure**:
   ```
   mkdir -p cmd/api pkg internal test
   ```

## API Implementation Breakdown

### Main Function (Entry Point)

```go
func main() {
    // Get port from environment variable or use default
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    // Define HTTP server
    mux := http.NewServeMux()

    // Register routes with middleware
    routes.RegisterRoutes(mux, middleware.LoggingMiddleware)

    // Start server
    addr := fmt.Sprintf(":%s", port)
    fmt.Printf("Server starting on port %s...\n", port)
    log.Fatal(http.ListenAndServe(addr, mux))
}
```

This main function:
1. Checks for a PORT environment variable, defaulting to "8080" if not set
2. Creates a new HTTP server multiplexer (router) with `http.NewServeMux()`
3. Registers the routes with the logging middleware
4. Starts the HTTP server with the configured port

### Handler Functions

#### Hello World Handler

```go
func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
    // Set content type
    w.Header().Set("Content-Type", "application/json")
    
    // Create response
    response := Response{
        Message: "Hello World from Go API!",
    }
    
    // Encode and send response
    err := json.NewEncoder(w).Encode(response)
    if err != nil {
        http.Error(w, "Failed to encode response", http.StatusInternalServerError)
        return
    }
}
```

#### Random Quote Handler

```go
func RandomQuoteHandler(w http.ResponseWriter, r *http.Request) {
    // Set content type
    w.Header().Set("Content-Type", "application/json")

    // Create a new quote service
    quoteService := data.NewQuoteService()
    
    // Get a random quote
    randomQuote := quoteService.GetRandomQuote()

    // Create response
    response := QuoteResponse{
        Status: "success",
        Data:   randomQuote,
    }

    // Encode and send response
    err := json.NewEncoder(w).Encode(response)
    if err != nil {
        http.Error(w, "Failed to encode response", http.StatusInternalServerError)
        return
    }
}
```

These handlers:
1. Set the response content type to "application/json"
2. Create the appropriate response structures
3. Encode the structs as JSON and write them to the response
4. Handle any encoding errors

### Response Structures

```go
// Response represents the API response structure for hello world endpoint
type Response struct {
    Message string `json:"message"`
}

// Quote represents a quote with its text and author
type Quote struct {
    Text   string `json:"text"`
    Author string `json:"author"`
}

// QuoteResponse is the response structure for quote endpoints
type QuoteResponse struct {
    Status string `json:"status"`
    Data   Quote  `json:"data"`
}
```

These structs:
1. Define the structure of our JSON responses
2. Use struct tags (`json:"field"`) to specify the JSON field names
3. Will be automatically marshaled to JSON when using `json.NewEncoder().Encode()`

## Testing

### Unit Tests

The project includes unit tests for all handlers in the `test` directory. These tests:
1. Create fake HTTP requests
2. Record the responses using `httptest.NewRecorder()`
3. Verify the response status codes, content types, and bodies

To run the tests:
```
go test ./test
```

For verbose output:
```
go test -v ./test
```

### Manual Testing

You can also test the API manually using curl:

```
# Start the server
go run cmd/api/main.go

# In another terminal
curl http://localhost:8080/hello_world
curl http://localhost:8080/quotes/random
```

Expected responses:
```json
{"message":"Hello World from Go API!"}
```

```json
{
  "status": "success",
  "data": {
    "text": "Life is what happens when you're busy making other plans.",
    "author": "John Lennon"
  }
}
```

## Accessing the Deployed API

The API is deployed and accessible at https://goapi-idtt.onrender.com. You can interact with it using the following methods:

### Using curl

To access the endpoints with curl:

```
curl https://goapi-idtt.onrender.com/hello_world
curl https://goapi-idtt.onrender.com/quotes/random
```

You can also use curl with additional options to see more details:

```
# Format the JSON output nicely (if you have jq installed)
curl https://goapi-idtt.onrender.com/quotes/random | jq

# Make a verbose request to see the full HTTP transaction
curl -v https://goapi-idtt.onrender.com/hello_world
```

## Go Concepts Used in This Project

### 1. HTTP Server
Go's standard library provides a powerful HTTP server through the `net/http` package. No external frameworks are required for basic web services.

### 2. Structs and JSON
Go uses structs to define data structures. The `encoding/json` package handles JSON serialization and deserialization.

### 3. Error Handling
Go uses explicit error checking rather than exceptions. Functions often return an error as their last return value.

### 4. Environment Variables
The `os.Getenv()` function retrieves environment variables, which is a common way to configure applications.

### 5. Testing
Go has built-in testing support via the `testing` package and `go test` command.

### 6. Middleware Pattern
The middleware pattern is implemented to add request logging functionality. This demonstrates:
- Function wrapping with closures
- HTTP handler chaining
- Custom ResponseWriter implementation
- Time measurement for performance tracking

Our logging middleware captures:
- HTTP method and URL path
- Response status code
- Request processing duration

```go
// Example of our logging middleware implementation
func LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        startTime := time.Now()
        
        crw := &customResponseWriter{
            ResponseWriter: w,
            statusCode:     http.StatusOK,
        }
        
        log.Printf("Request: %s %s", r.Method, r.URL.Path)
        
        next(crw, r)
        
        duration := time.Since(startTime)
        
        log.Printf("Response: %s %s - Status: %d - Duration: %v", 
            r.Method, r.URL.Path, crw.statusCode, duration)
    }
}
```

### 7. Package Organization
Go encourages organizing code into packages:
- `internal`: Private code that can't be imported by other modules
- `cmd`: Entry points for applications
- `pkg`: Library code that can be imported by other modules
- `test`: Dedicated test files

## Deployment

The project is configured for deployment on Render using the `render.yaml` file:

```yaml
services:
  - type: web
    name: GoAPI
    env: go
    buildCommand: go build -o app ./cmd/api
    startCommand: ./app
    envVars:
      - key: PORT
        value: 10000
```

This configuration tells Render:
1. This is a web service using Go
2. How to build the application
3. How to start the application
4. What environment variables to set

## Next Steps

To improve this API, consider:
1. Implementing graceful shutdown
2. Adding more endpoints with different functionality
3. Using a router package like gorilla/mux or a framework like Gin
4. Adding a database connection
5. Adding authentication
