# Вездекод Финал 2022. Go

## 10
### Компиляция:
```
go build ./10/main.go
```

### Запуск:
```
./main <path>
```

Где path - путь до файла с задачами


## 20
### Компиляция:
```
go build ./20/main.go
```

### Запуск:
```
./main <path>
```

Где path - путь до файла с задачами

## 30
### Компиляция:
```
go build ./30/main.go
```

### Запуск:
```
./main <path>
```

Где path - путь до файла с задачами

## 40
### Компиляция:
```
go build ./40/main.go
```

### Запуск:
```
./main <addr>
```

Где addr - адрес на котором нужно запустить сервер

### Методы

#### POST /add
Параметры:

`type` - тип sync/async

`timeDuration` - время в формате time.Duration

Пример ответа:

```json
{
  "ok": true,
  "data": {
    "id": 1,
    "duration": 15000000000
  }
}
```

#### GET /schedule
Пример ответа:

```json
{
  "ok": true,
  "data": [
    {
      "id": 1,
      "duration": 15000000000
    },
    {
      "id": 2,
      "duration": 24000000000
    },
    {
      "id": 5,
      "duration": 50000000000
    }
  ]
}
```

#### GET /time
Пример ответа:

```json
{
  "ok": true,
  "data": 124000000000
}
```