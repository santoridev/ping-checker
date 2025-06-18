


#  Ping Checker API

A lightweight and concurrent URL availability checker built with Go and Gin. This service accepts a list of URLs via a POST request and returns the HTTP status of each URL, including response time.

##  Features

-  Concurrent URL checking using goroutines
-  Configurable timeout (default: 2 seconds)
-  Clean JSON API
-  Uses Go's `context` for request timeouts

---

##  Technologies

- [Go](https://golang.org/)
- [Gin Web Framework](https://github.com/gin-gonic/gin)
- `context`, `http`, `time`, and `sync` from Go stdlib

---

##  Project Structure

```

.
├── main.go              # Entry point for the server
├── checker/
│   └── checker.go       # Core logic for checking URLs
└── handlers/
└── handlers.go      # HTTP handler for the /check endpoint

````

---

##  API Usage

### POST `/check`

Check the availability of multiple URLs.

####  Request Body:

```json
{
  "urls": [
    "https://google.com",
    "https://example.com",
    "https://nonexistent.tld"
  ]
}
````

####  Response:

```json
[
  {
    "url": "https://google.com",
    "status": "200 OK (120ms)"
  },
  {
    "url": "https://example.com",
    "status": "200 OK (85ms)"
  },
  {
    "url": "https://nonexistent.tld",
    "status": "error"
  }
]
```

---

##  Running the App

1. **Clone the repository**:

```bash
git clone https://github.com/santoridev/ping-checker.git
cd ping-checker
```

2. **Run the server**:

```bash
go run .
```

The server will start at `http://localhost:8080`.

---

##  Example Curl Request

```bash
curl -X POST http://localhost:8080/check \
  -H "Content-Type: application/json" \
  -d '{
    "urls": ["https://google.com", "https://example.com"]
  }'
```

---

## Configuration

You can easily modify the default timeout in the `checker.NewCheckURLs()` function. Currently set to `2 * time.Second`.

---

## License

MIT License — feel free to use this project as you wish.

---

## Credits

Built by SANTORI using Go and the Gin framework.


