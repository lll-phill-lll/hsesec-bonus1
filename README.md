**How to run**
1. Run postgres: `docker-compose -f docker-compose.yml up -d`
2. Run one of binaries from `/bin` or install go and execute `go build cmd/hsesec/main.go`, then run binary
3. Go `localhost:9090`

**Note**
To stop docker run `docker-compose -f docker-compose.yml down -v`