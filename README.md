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
.
├── cmd/
│   └── http.go         # CLI command to start HTTP server
├── internal/
│   ├── config/         # Viper setup
│   ├── logger/         # Logrus setup
│   ├── server/         # Fiber server bootstrap
│   └── ...
├── domain/             # Core business logic (hexagonal core)
├── ports/              # Interfaces for external adapters
├── adapters/           # Infrastructure implementation (DB, HTTP, etc.)
├── go.mod
├── main.go
└── .env                # Local dev environment variables
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
