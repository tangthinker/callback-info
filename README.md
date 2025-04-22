# Callback Info

A simple web application for testing and collecting callback requests, with a clean web interface to display the collected data.

## Features

- Receive and store callback requests
- Record request IP, parameters, and timestamp
- Return 200 OK response
- Web interface to display collected data
- Support pagination and sorting
- Store up to 10,000 latest records
- JSON parameter formatting and highlighting
- Configurable database location and server port

## Installation

1. Clone the repository:
```bash
git clone https://github.com/yourusername/callback-info.git
cd callback-info
```

2. Install dependencies:
```bash
go mod tidy
```

## Usage

### Basic Usage
```bash
go run main.go
```
This will start the server with default settings:
- Database file: `./callback.db`
- Server port: `8080`

### Custom Configuration
```bash
# Custom database location
go run main.go -db /path/to/your/callback.db

# Custom server port
go run main.go -port 9090

# Both custom database location and port
go run main.go -db /path/to/your/callback.db -port 9090
```

### Testing Callback
Send a POST request to test the callback:
```bash
curl -X POST -H "Content-Type: application/json" -d '{"test":"data"}' http://localhost:8080/api/callback
```

## Web Interface

Access the web interface at `http://localhost:8080` to:
- View all callback records
- See formatted JSON parameters
- Navigate through pages
- Sort by timestamp

## Data Storage

- Uses SQLite database
- Stores up to 10,000 latest records
- Automatically removes oldest records when limit is reached
- Records include:
  - Request timestamp
  - IP address
  - Request parameters

## License

MIT License