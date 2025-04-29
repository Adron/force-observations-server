# Force Observations Server

A Go-based server for managing computer vision activities across multiple cameras. This service provides API endpoints for camera management, video processing, and observation data collection.

## Features

- Multi-camera support
- RESTful API endpoints
- Structured logging
- Real-time camera status monitoring

## Prerequisites

- Go 1.21 or later
- Git

## Getting Started

1. Clone the repository:
```bash
git clone [repository-url]
cd force-observations-server
```

2. Install dependencies:
```bash
go mod download
```

3. Run the server:
```bash
go run main.go
```

The server will start on port 8080 by default.

## API Endpoints

- `GET /health` - Service health check
- `GET /cameras` - List all connected cameras
- `GET /cameras/{id}` - Get specific camera details
- `POST /observations` - Create a new observation

## Development

### Project Structure

```
.
├── internal/
│   ├── api/
│   ├── camera/
│   └── logging/
├── pkg/
├── main.go
├── go.mod
└── README.md
```

### Adding New Features

1. Create new packages in the `internal` directory
2. Add corresponding tests in the same package
3. Update the API endpoints in `internal/api`

## License

[License information to be added] 