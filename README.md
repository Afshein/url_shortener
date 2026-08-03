# URL Shortener (exercise to learn GO)

#### Redis 
`docker run -p 6379:6379 redis:latest`

#### Server
`export $(grep -v '^#' .env | xargs) && go run ./cmd/web -PORT=$PORT`


#### Example POST Request
`curl -d "url=https://www.google.com" -X POST http://localhost:4000/`
`

#### Example GET Request

`curl http://localhost:4000/shorten/12424224562945606824`