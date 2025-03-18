# BuildProject

## Начало работы
Для начала работы следует создать файл `.env` в корне проекта. Примерно выглядит так:
```python
POSTGRES_DB=db_name
POSTGRES_USER=username
POSTGRES_PASSWORD=password
JWT_KEY=key
```

## Сборка проекта
```shell
$ docker compose build
```

## Запуск проекта
```shell
$ docker compose up
```

## Остановка проекта
```shell
$ docker compose stop
```
