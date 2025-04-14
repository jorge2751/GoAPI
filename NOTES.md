# Go API Project Notes

This document provides a detailed breakdown of our Go API project for someone learning Go. It covers the project structure, steps to set up the environment, and commands to test and run the API.

## Project Overview

This is a simple REST API built with Go's standard library that:
- Serves a JSON response at the `/hello_world` endpoint
- Uses proper HTTP status codes and content types
- Handles errors appropriately
- Can be configured via environment variables
- Has been set up for deployment on Render

## Project Structure

```
go-hello-api/
├── cmd/
│   └── api/
│       ├── main.go         # Application entry point
│       └── main_test.go    # Tests for the API
├── internal/               # Private application code
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
   go mod init github.com/jorgeleon/go-hello-api
   ```

4. **Create Project Structure**:
   ```
   mkdir -p cmd/api pkg internal
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

    // Register routes
    mux.HandleFunc("/hello_world", helloWorldHandler)

    // Start server
    addr := fmt.Sprintf(":%s", port)
    fmt.Printf("Server starting on port %s...\n", port)
    log.Fatal(http.ListenAndServe(addr, mux))
}
```

This main function:
1. Checks for a PORT environment variable, defaulting to "8080" if not set
2. Creates a new HTTP server multiplexer (router) with `http.NewServeMux()`
3. Registers the `/hello_world` route with its handler function
4. Starts the HTTP server with the configured port

### Handler Function

```go
func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
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

This handler function:
1. Sets the response content type to "application/json"
2. Creates a Response struct with a message
3. Encodes the struct as JSON and writes it to the response
4. Handles any encoding errors

### Response Structure

```go
// Response represents the API response structure
type Response struct {
    Message string `json:"message"`
}
```

This struct:
1. Defines the structure of our JSON response
2. Uses a struct tag (`json:"message"`) to specify the JSON field name
3. Will be automatically marshaled to JSON when using `json.NewEncoder().Encode()`

## Testing

### Unit Tests

The project includes a unit test for the `helloWorldHandler` function in `main_test.go`. This test:
1. Creates a fake HTTP request
2. Records the response using `httptest.NewRecorder()`
3. Verifies the response status code, content type, and body

To run the tests:
```
cd cmd/api
go test
```

For verbose output:
```
go test -v
```

### Manual Testing

You can also test the API manually using curl:

```
# Start the server
go run cmd/api/main.go

# In another terminal
curl http://localhost:8080/hello_world
```

Expected response:
```json
{"message":"Hello World from Go API!"}
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

## Deployment

The project is configured for deployment on Render using the `render.yaml` file:

```yaml
services:
  - type: web
    name: go-hello-api
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
1. Adding middleware for logging
2. Implementing graceful shutdown
3. Adding more endpoints with different functionality
4. Using a router package like gorilla/mux or a framework like Gin
5. Adding a database connection
