# API Documentation

## Users

### GET /users
#### Description
Получить список пользователей с поддержкой пагинации и фильтрации.

#### Query Parameters
- `name` (optional): Фильтровать по имени.
- `age` (optional): Фильтровать по возрасту.
- `page` (optional): Номер страницы (по умолчанию 1).
- `limit` (optional): Количество пользователей на странице (по умолчанию 10).

#### Response
```json
{
    "total": 3,
    "users": [
        {"id": 1, "name": "Alice", "age": 30},
        {"id": 2, "name": "Bob", "age": 25},
        {"id": 3, "name": "Charlie", "age": 35}
    ]
}
