# Classroom main service

[![Main CI](https://github.com/Binaretech/classroom-main/actions/workflows/go.yml/badge.svg)](https://github.com/Binaretech/classroom-main/actions/workflows/go.yml)

## Arquitecture

The full classroom is made up with three services routed by traefik, two databases and a cache store

This repository contains the auth service made with go

The following image describe the full arquitecture and the tecnologies used
![arquitecture](https://github.com/Binaretech/classroom/blob/main/img/classroom-diagram.png?raw=true)

Full source is available on https://github.com/Binaretech/classroom

## Description

This service is designed to handle with users and classes process

## Running

This service depends on the `Main service` only for the login interface, also this project uses `Redis` for JWT verification

Meeting the requirements, just run:

```bash
go run ./cmd/service
```

Dockefile and docker-compose files are availables for development to mount containers with the running service, `Postgres` and `Redis` with its canonical port exposed to the host for debug purposes. A `MinIO` container will be mount to act as a S3 storage. In addition every change to the code will be live reloaded. Just run:

```bash
docker-compose up -d
```

