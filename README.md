# Microservices Playground

# Requirements

## Go

### Link

https://golang.org/doc/install

## DEP

### Link

https://golang.github.io/dep/docs/installation.html

### Installation

```shell
# macOS

$ brew install dep
$ brew upgrade dep
```

User `dep` to install vendor files

## Docker

### Link

https://docs.docker.com/docker-for-mac/install/

## Docker Compose

### Link

https://docs.docker.com/compose/install/

## Yarn

### Link

https://yarnpkg.com/lang/en/docs/install/#mac-stable

### Installation

```shell
$ brew install yarn
```

## golang-migrate

### Link

https://github.com/golang-migrate/migrate/tree/master/cmd/migrate

### Installation

```shell
$ brew install golang-migrate
```

# Environment

## Docker Compose

```shell
$ docker-compose up -d
```

# Migrations

## Auth: PostgresQL

### CLI

```shell
$ migrate -database postgres://ms_auth_psql:password@localhost:5432/auth_db?sslmode=disable -path ./migrations up
$ migrate -database postgres://ms_cc_psql:password@localhost:5433/code_challenge_db?sslmode=disable -path ./migrations up
```

### Seeding

```shell
$ go run services/utilities/seed/admin.go
```

### Connecting Manually

```shell
$ psql -h localhost -p 5432 -d auth_db -U ms_auth_psql
$ psql -h localhost -p 5433 -d code_challenge_db -U ms_cc_psql
```

# Starting Application

```shell
// New tab
$ cd services/auth/
$ go run auth_server/main.go auth_server/server.go auth_server/config.go

// New tab
$ cd services/code-challenge/server/
$ go run server.go config.go

// New tab
$ cd services/code-challenge/service
$ go run main.go service.go config.go

// New tab
$ cd uis/msp-overview/
$ yarn start
```

# TODOs

1. Configuration Files (DONE) / NOTE: Requires refactoring
    * https://medium.com/@felipedutratine/manage-config-in-golang-to-get-variables-from-file-and-env-variables-33d876887152
2. Service: Logging (DONE) / NOTE: code challenge architecture is all wrong! MUST REFACTOR
    * logrus - https://github.com/sirupsen/logrus
3. Service: Error Handling
    * https://github.com/juju/errors
4. Testing
    * Ginkgo - https://onsi.github.io/ginkgo/
5. Instructions to Setup
6. Productivity?
    * Cobra - https://github.com/spf13/cobra