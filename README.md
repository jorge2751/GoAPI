# Go Hello World API

A simple Go API that returns a hello world message at the `/hello_world` endpoint. This project demonstrates how to create a basic REST API in Go and deploy it to Render.

## Features

- Simple HTTP server using Go's standard library
- JSON response format
- Configurable port via environment variables
- Deployable to Render

## Prerequisites

- Go 1.20 or later
- Git

## Getting Started

### Installation

1. Clone the repository:
   ```
   git clone https://github.com/[your-username]/go-hello-api.git
   cd go-hello-api
   ```

2. Run the application:
   ```
   go run cmd/api/main.go
   ```

3. The server will start on port 8080 by default. You can change this by setting the `PORT` environment variable.

## API Endpoints

### GET /hello_world

Returns a hello world message in JSON format.

**Response Example:**

```json
{
  "message": "Hello World from Go API!"
}
```

## Deployment

This API can be deployed to Render by connecting your GitHub repository and using the following settings:

- **Build Command:** `go build -o app ./cmd/api`
- **Start Command:** `./app`
- **Environment Variables:** Set `PORT` if needed

## Development

This project follows the standard Go project layout:

- `cmd/api`: Application entry point
- `internal`: Private application code
- `pkg`: Public libraries that can be used by external applications

## License

MIT 