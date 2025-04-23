# MyAnimeGameList Backend

## Requirements

- golang 1.24.1
- magl-docker

### 1. install sql-migrate

```sh
go install github.com/rubenv/sql-migrate/...@latest
```

### 2. install wire

```sh
go install github.com/google/wire/cmd/wire@latest
```

### 3. install swaggo

```sh
go install github.com/swaggo/swag/cmd/swag@latest
```

## Infrastructure

### 1. cd to magl-docker

```sh
cd ~/magl/magl-docker
```

### 2. run databases

```sh
make dev
```

### 4. cd to malg-backend

```sh
cd ~/magl/magl-backend
```

### 5. migration

```sh
sql-migrate up
```

## Run backend

### 1. install air (recommand)

```
go install github.com/air-verse/air@latest
```

### 2. run app

```sh
air
# if you dont install air
go run cmd/main.go
```
