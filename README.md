# BuildProject

## Начало работы
> [!IMPORTANT]
> Для начала работы следует создать файл `.env` в корне проекта. Примерно выглядит так:

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

## Запуск миграций
> [!NOTE]
> Перед запуском миграций необходимо запустить контейнеры `backend` и `db`.

```shell
$ docker compose exec backend ./mp1_backend -migrate
```

## Запуск проекта
> [!IMPORTANT]
> Необходимо запустить миграцию после каждого изменения моделей в `backend/models/`, либо перед первым запуском.

```shell
$ docker compose up
```


## Остановка проекта
```shell
$ docker compose stop
```
