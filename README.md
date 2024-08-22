# go_calculator_backend_api


Start the API server
```shell
go run main.go logging_middleware.go api_handlers.go
```

Make requests to the API server:
```shell
curl localhost:8080/status
```

```shell
curl -X POST localhost:8080/multiply -d '{"a": 8, "b": 2, "c":4}'
```