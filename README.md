# CHATTIMUS SERVER

## Folder struct

### build

In this folder there are configuration files for build the project

### cmd

Here are the diferents entrypoints of the applications

### config

Parametizable variables

### pkg

Reusable package for the aplications
The application logic must be in this folder separated in packages

## RUN PROJECT

With docker

```bash
$ docker-compose up -d
```

Without docker

1. For web server

```bash
$ go run ./cmd/server
```

2. For socket server

```bash
$ go run ./cmd/socket
```
