# Caching Proxy Server

A simple caching proxy server written in Go that caches responses from an origin server. When a request is made, it first checks if the response is available in the cache. If found, it returns the cached response; otherwise, it forwards the request to the origin server and caches the response.

## Features

- Forward HTTP GET requests to an origin server
- In-memory caching of responses
- Cache hit/miss logging
- Command-line interface for configuration
- Cache clearing functionality
- Custom port configuration

## Project Structure

```
.
├── main.go           # Entry point and CLI handling
├── proxy/
│   └── proxy.go      # Proxy server implementation
└── cache/
    └── cache.go      # Cache implementation
```

## Prerequisites

- Go 1.16 or higher

## Installation

1. Clone the repository:
```bash
git clone https://github.com/RohithBN/caching-proxy.git
cd caching-proxy
```

2. Build the project:
```bash
go build
```

## Usage

### Starting the Server

Start the proxy server by specifying the port and origin server:

```bash
./caching-proxy --port <port_number> --origin <origin_url>
```

Example:
```bash
./caching-proxy --port 8080 --origin http://dummyjson.com
```

### Clearing the Cache

To clear the cache:
```bash
./caching-proxy --clear-cache
```

### Command Line Flags

- `--port`: Port number on which the proxy server will listen (default: 8080)
- `--origin`: URL of the origin server to which requests will be forwarded
- `--clear-cache`: Clear the in-memory cache

## Testing

1. Start the server:
```bash
./caching-proxy --port 8080 --origin http://dummyjson.com
```

2. Make a request (First request - Cache Miss):
```bash
curl http://localhost:8080/products
```

3. Make the same request again (Cache Hit):
```bash
curl http://localhost:8080/products
```

4. Clear the cache:
```bash
./caching-proxy --clear-cache
```

## Limitations

- Only supports GET requests
- In-memory cache (data is lost when server restarts)
- No cache expiration mechanism
- No concurrent request handling protection
- No HTTPS support
- No response headers forwarding

## Future Improvements

- Add support for other HTTP methods
- Implement cache expiration/TTL
- Add persistent storage
- Add concurrent request handling
- Add HTTPS support
- Forward response headers
- Add cache size limits
- Add metrics and monitoring
- Add tests

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

---

For questions and suggestions, please open an issue in the repository.

https://roadmap.sh/projects/caching-server
