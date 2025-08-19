# URL Shortener — Fiber + GORM + Clean Architecture

## Step to run
1. `crete database`
2. `cp .env.example .env` (opsional edit)
3. `setup .env, adjust to your db connection`
4. `go mod tidy`
5. `go run ./cmd/server`


## Example Request

### Create
curl -X POST http://localhost:8080/api/v1/urls \  -H 'Content-Type: application/json' \  -d '{"original_url":"https://golang.org","custom_alias":"go","ttl_hours":24}'

### Redirect
curl -i http://localhost:8080/go

### Get Stats
curl http://localhost:8080/api/v1/urls/go

### Delete
curl -X DELETE http://localhost:8080/api/v1/urls/go -i
