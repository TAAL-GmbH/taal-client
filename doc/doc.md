# Developer documentation

## Building

Build the application by running the build script `build.sh`

### Prequisites

For the building the appliaction the following tools need to be preinstalled
- `go` >= v1.17
- `npm`

## Integration tests

Currently two types of databases are supported: `sqlite` and `postgres`

In order to run the integration tests with either database use the environment variable `DB`

### SqLite

For integration tests with `sqlite` run

```
DB=SQLITE go test -v repository/repository_integration_test.go
```

### Postgres

For integration tests with `postgres` run

```
DB=POSTGRES go test -v repository/repository_integration_test.go
```

[Docker](https://www.docker.com/products/docker-desktop/) needs to be preinstalled for this integration test to run
