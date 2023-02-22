# RATE_LIMITER

## Инструкция

1. Склонировать репозиторий
```
git clone https://github.com/vladjong/rate_limiter
```

2. Добавить `.env` файл в проект

3. Запустить проекта через docker compose
```
make docker
```
4. Завершить проект
```
make clean
```
5. Запустить тесты
```
make test
```

## Пример запросов к API

1. PING-PONG (тестовый хендлер)
```
curl --location --request GET 'http://0.0.0.0:8080' \
--header 'X-Forwarded-For: 123.45.67.1'
```

2. Refresh
```
curl --location --request GET 'http://0.0.0.0:8080/refresh' \
--header 'X-Forwarded-For: 123.45.67.89'
```


