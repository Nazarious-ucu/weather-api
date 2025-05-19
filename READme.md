# 🌦️ Weather Subscription API

**Weather Subscription API** — це REST-сервіс для отримання поточних погодних умов у заданому місті та підписки користувачів на оновлення погоди.
Проєкт реалізований на Go з використанням Gin, Swagger, SQLite та Docker.

## Можливості


* Отримання поточної погоди за містом
* Підписка на email-оновлення погоди (щогодинно або щодневно)
* Підтвердження підписки через токен
* Відписка за токеном
* Swagger-документація

## Технології

* Go 1.23+
* Gin (HTTP фреймворк)
* SQLite (вбудована база даних)
* Goose (міграції)
* Swagger + swag CLI
* Docker + Docker Compose

## 📁 Структура проєкту

```
.
├── cmd/WeatherSubscriptionAPI         
├── internal/
│   ├── handlers                       
│   ├── services                       
│   └── repository                     
├── migrations/                        
├── web/                               
├── docs/                              
├── api/swagger.yaml                   
├── Dockerfile                         
├── docker-compose.yml                 
└── .env                               
```

## Налаштування та запуск

1. Створи файл `.env` у корені:

```
WEATHER_API_KEY=your_api_key_here
```

2. Запусти застосунок:

```bash
docker-compose up --build
```

3. Перевірка:

* Swagger: [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)
* Форма підписки: [http://localhost:8080/](http://localhost:8080/)

## Приклади запитів

### Отримання погоди

```
GET /weather?city=Lviv
```

### Створення підписки

```
POST /subscribe
Content-Type: application/x-www-form-urlencoded

email=test@example.com&city=Lviv&frequency=daily
```

### Підтвердження

```
GET /confirm/{token}
```

### Відписка

```
GET /unsubscribe/{token}
```
MIT License © 2025 [Nazar Parnosov](https://github.com/nazarparnosov)
