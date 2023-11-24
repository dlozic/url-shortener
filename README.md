# Shortenex

A simple URL shortener written in Go.

## Docker

### API
```bash
docker build -t shortenex .
docker run -p 8080:8080 shortenex
```

### API calls
```
POST /api/urls/shorten
GET /api/urls
GET /api/urls/{id}
GET /api/users
```

## Local development

### API
```bash
make build
make run
```