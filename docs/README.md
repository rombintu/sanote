# SANOTE - Backend

## Зависимости
* Golang >= 1.16
* MongoDB
* Docker
* Docker-compose
## Быстрый запуск
```bash
$ cd sanote/backend
$ make run
```

## Сборка
```bash
$ cd sanote/backend
$ make build
```

## Docker-compose (Рекомендуется)
```bash
$ cd sanote/backend
$ export MONGO_USER=mongo
$ export MONGO_PASS=mongo
$ make docker-compose
```