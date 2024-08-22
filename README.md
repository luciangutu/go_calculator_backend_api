# Go API Project

This Go project provides a simple API with endpoints for basic arithmetic operations and a status check. It demonstrates the use of middleware for logging HTTP requests and responses.

## Features

- **/status**: Returns the service status.
- **/add**: Adds multiple numbers.
- **/subtract**: Subtracts multiple numbers.
- **/multiply**: Multiplies multiple numbers.
- **/divide**: Divides numbers, ensuring no division by zero.

## Project Structure

- `main.go`: The entry point of the application that sets up routes and starts the server.
- `logging_middleware.go`: Contains middleware for logging request and response details.
- `api_handlers.go`: Contains handler functions for the API endpoints.

## Installation

1. **Clone the Repository**:

```shell
   git clone https://github.com/luciangutu/go_calculator_backend_api.git
   cd <repository-directory>
```

2. **Install Dependencies**:

Go modules are used for dependency management. Ensure you have Go installed, then run:

```shell
go mod tidy
```

## Running the Application

1. **Run the Application**

   Compile and run the application with:

```shell
   go run main.go logging_middleware.go api_handlers.go
```

2. **Access the API**

Use the following endpoints to interact with the API:

- **Status**: `GET http://localhost:8080/status`
- **Add**: `POST http://localhost:8080/add`
  - Request body example: `{"a": 1, "b": 2, "c": 8}`
- **Subtract**: `POST http://localhost:8080/subtract`
  - Request body example: `{"number1": 5, "number2": 3}`
- **Multiply**: `POST http://localhost:8080/multiply`
  - Request body example: `{"number1": 2, "number2": 3}`
- **Divide**: `POST http://localhost:8080/divide`
  - Request body example: `{"number1": 6, "number2": 3}`


## Testing

You can test the API using tools like `curl` or Postman. For example, using `curl`:

```shell
curl -X POST http://localhost:8080/add -H "Content-Type: application/json" -d '{"a": 5, "b": 10}'
```

## Middleware

The `LoggingMiddleware` logs details of each request and its duration. It wraps around each handler function to provide logging functionality.


## Running Tests

To ensure the functionality of the application, you can run the tests using the Go testing tool.

  Execute the following command to run all tests:

```shell
   go test -v
```

## Contributing

Feel free to contribute by submitting issues or pull requests.