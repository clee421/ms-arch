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

## Kubernetes

### Link

https://kubernetes.io/docs/tasks/tools/install-kubectl/#install-kubectl-on-macos

### Installation

```shell
$ brew install kubernetes-cli
$ go get -d k8s.io/kubernetes
```

## Minikube

### Link

https://kubernetes.io/docs/tasks/tools/install-minikube/#before-you-begin

### Installation

```shell
$ brew cask install minikube
```

## VirtualBox

### Link

https://www.virtualbox.org/wiki/Downloads

## Skaffold

### Link

https://skaffold.dev/docs/getting-started/#installing-skaffold

### Installation

```shell
# macOS

$ brew install skaffold
```

# Environment

## Kubernetes

```shell
$ kubectl create -f pod.yaml
```

# Notes

## Kubernetes

https://github.com/GoogleCloudPlatform/postgresql-docker/blob/master/9/README.md#using-kubernetes

### PostgresQL

Create Pod
```shell
$ kubectl create -f pod.yaml
```

Expose port
```shell
$ kubectl expose pod auth-psql --name auth-psql-5432 \
  --type LoadBalancer --port 5432 --protocol TCP
```

Run a PostgreSQL client directly within the container.
```shell
$ kubectl exec -it auth-psql -- psql --username ms_auth_psql
```

Run in a new terminal to expose external ip
```shell
$ minikube tunnel
```

# Migrations

## Auth: PostgresQL

### CLI

```shell
$ migrate -database postgres://ms_auth_psql:password@localhost:5432/auth_db?sslmode=disable -path ./migrations up
```

### Connecting Manuall

```shell
$ psql -h localhost -p 5432 -d auth_db -U ms_auth_psql
```

# TODOs

1. Configuration Files
2. Service: Logging
3. Service: Error Handling
4. Instructions to Setup