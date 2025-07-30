🧩 Go Server Template

A modular Go server application using:
•	Cobra CLI – for structured command-line interface
•	Fiber – fast, Express-like web framework
•	Viper – config management via .env or environment variables
•	Logrus – structured logging (JSON-formatted, production-ready)
•	Hexagonal architecture – clean separation between domain, transport, and infrastructure

⸻

📦 Project Structure
```
go-server/
├── cmd/
│   └── http.go                # Cobra command to start HTTP server
├── config/
│   └── config.go              # Viper config setup (InitConfig)
├── internal/
│   ├── adapter/
│   │   ├── http/
│   │   │   ├── handler/
│   │   │   │   └── user_handler.go    # Example handler
│   │   │   ├── router.go              # Fiber route setup
│   │   │   └── server.go              # Fiber app instance creation
│   │   └── repository/
│   │       └── user_repo.go          # DB implementation (outbound)
│   ├── domain/
│   │   ├── model/
│   │   │   └── user.go               # Entity definition
│   │   └── port/
│   │       ├── repository.go        # Outbound ports (interfaces)
│   │       └── service.go           # Inbound ports (use case interfaces)
│   ├── usecase/
│   │   └── user_service.go          # Business logic / Application Service
│   └── logger/
│       └── logger.go                # Logrus logger setup
├── .env                            # Only used with `--dev`
├── go.mod
├── go.sum
└── main.go                         # Entrypoint (calls cobra.Execute)
```

🚀 Commands

Run in development mode (local machine)

Uses .env file:
```shell
go run main.go http --dev
```

Run in production mode (Docker/Kubernetes)

Uses system environment variables:
```shell
go run main.go http
```

⚙️ Configuration

.env example (for local dev)
```shell
PORT=8080
LOG_LEVEL=debug
DEV_MODE=true
```
For staging/production, set environment variables via Kubernetes ConfigMap or Docker ENV.


📖 Logging

- Uses **Logrus** with `JSONFormatter` for log aggregation (e.g., Elasticsearch, Loki).
- Automatically logs:
    - `@timestamp`
    - `level`
    - `message`
    - `caller` (file:line)
    - custom fields like `request_id`, `action`, etc.

Sample log:
```json
{
  "@timestamp": "2025-07-30T08:45:21+07:00",
  "level": "info",
  "message": "User logged in",
  "caller": "internal/auth/handler.go:42",
  "request_id": 1234,
  "action": "login"
}
```

🧪 Testing

You can add tests using standard Go test tools:
```shell
go test ./...
```

🛠 Requirements

•	Go 1.22+
•	Cobra CLI (go install github.com/spf13/cobra-cli@latest)
•	Fiber, Viper, Logrus (go get handles them)
