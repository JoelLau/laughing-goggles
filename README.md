# laughing-goggles

highly inspired by previous projects [user-analytics-service](https://github.com/JoelLau/user-analytics-service/)
and [personal-accounting](https://github.com/JoelLau/personal-accounting)

## Pre-requisites

1. [docker](https://www.docker.com/products/docker-desktop/)
1. [docker compose](https://docs.docker.com/compose/)
1. [Go](https://go.dev/) >= 1.26.3
1. [optional] [goose](https://github.com/pressly/goose)
1. [optional] [sqlc](https://sqlc.dev/)

## Running this project

These are the recommended instructions for running this project
on a unix-like machine (e.g. macOS, Linux).

1. configure project by setting environment variables
    1. `cp .env.example .env`: create an _.env_ file using the provided example
    1. _modify the contents of .env_, PLEASE CHANGE THE DEFAULT PASSWORD
    1. `export $(grep -v '^#' .env | xargs)`: set the contents of the .env file
        into your local environment
1. set up database
    1. `docker compose up -d` run docker compose in background
        - spins up postgresql database instance
        - runs database migrations to create necessary tables, etc
1. install go dependencies
    - `go mod tidy`
1. start the transfers service:
    - `go run cmd/server/main.go` starts up the REST API
1. hit the endpoints you need
    1. [swagger docs are available at http://localhost:8080/swagger/](http://localhost:8080/swagger/)
    1. click the "Try it out" button on the corresponding endpoint
    1. modify the parameters
    1. click execute to see results

## Commands

| Commands | Remarks |
| --- | --- |
| `docker compose up -d` | starts external services |
| `docker compose down -v` | stops external services |
| `go mod tidy` | install go dependencies |
| `cp .env.example .env` | create .env from examples |
| `export $(grep -v '^#' .env | xargs)` | load environment variables |
| `go run cmd/server/main.go` | starts transfers service REST API server |
| `go test ./...` | run tests (requires docker) |
| `go generate ./...` | code generation - sql helpers, openapi interfaces |
