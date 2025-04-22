# Callback Info

A simple and efficient web application for testing and collecting callback requests, with a clean web interface to display the collected data.

## Features

- Receive and store callback requests
- Record request IP, parameters, and timestamp
- Return 200 OK response
- Web interface to display collected data
- Support pagination and sorting
- Store up to 10,000 latest records
- JSON parameter formatting and highlighting
- Configurable root directory and server port
- Automatic cleanup of old records
- Real-time data display

## Installation

1. Clone the repository:
```bash
git clone https://github.com/tangthinker/callback-info.git
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
- Root directory: `.` (current directory)
- Server port: `8080`
- Database file: `callback.db` (in root directory)

### Custom Configuration
```bash
# Custom root directory
go run main.go -root /path/to/your/root

# Custom server port
go run main.go -port 9090

# Both custom root directory and port
go run main.go -root /path/to/your/root -port 9090
```

### Testing Callback
Send a POST request to test the callback:
```bash
# Basic test
curl -X POST -H "Content-Type: application/json" -d '{"test":"data"}' http://localhost:8080/api/callback

# Test with complex JSON
curl -X POST -H "Content-Type: application/json" -d '{
  "name": "test",
  "value": 123,
  "items": ["item1", "item2"],
  "nested": {
    "key": "value"
  }
}' http://localhost:8080/api/callback
```

## Web Interface

Access the web interface at `http://localhost:8080` to:
- View all callback records
- See formatted JSON parameters with syntax highlighting
- Navigate through pages
- Sort by timestamp (newest first)
- View request details including:
  - Timestamp
  - IP address
  - Request parameters

## Data Storage

- Uses SQLite database for efficient storage
- Database file is stored in the root directory
- Stores up to 10,000 latest records
- Automatically removes oldest records when limit is reached
- Records include:
  - Request timestamp
  - IP address
  - Request parameters
- Data is automatically cleaned up to maintain performance

## Development

### Building
```bash
go build -o callback-info
```

### Running Tests
```bash
go test ./...
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

MIT License

Copyright (c) 2024 Callback Info

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.