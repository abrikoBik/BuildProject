# BuildProject

## Начало работы
```shell
$ make init
```
Далее в новый файл `.env` добавить переменные `POSTGRES_DB`, `POSTGRES_USER`, `POSTGRES_PASSWORD` и `JWT_KEY`

## Сборка проекта
```shell
$ make
```

### Сборка бэкенда
```shell
$ make backend
```

### Сборка фронтенда
```
$ make frontend
```

## Запуск бэкенда и базы данных
```shell
$ make start
```

## Остановка бэкенда и базы данных
```shell
$ make stop
```
