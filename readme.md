
# Task Manager API

### Описание
Это простое приложение для управления задачами с использованием Go и фреймворка Gin. Приложение предоставляет возможности для создания, обновления, получения и удаления задач. Задачи хранятся в оперативной памяти (in-memory), что делает его подходящим для обучения и демонстрации работы с HTTP-запросами и базовыми CRUD операциями.

### Функционал:
- Добавление задачи.
- Получение списка всех задач.
- Получение задачи по её ID (через query-параметр).
- Обновление статуса задачи (completed).
- Удаление задачи по её ID.

### Требования:
- Go 1.19+
- Модуль `github.com/gin-gonic/gin`

### Установка и запуск:

1. Склонируйте репозиторий или скопируйте код в папку проекта:

   ```bash
   git clone https://github.com/x3m4llk0/GO-Task-Manager.git
   cd go-task-manager
   ```

2. Установите зависимости:

   ```bash
   go mod tidy
   ```

3. Запустите сервер:

   ```bash
   go run main.go
   ```

4. Сервер будет запущен по адресу: `http://localhost:8080`.

### API Эндпоинты:

#### 1. Создание задачи

- **URL:** `/task`
- **Метод:** `POST`
- **Тело запроса:** (JSON)

  ```json
  {
    "title": "New Task",
    "description": "Description of the new task"
  }
  ```

- **Ответ:** (201 Created)

  ```json
  {
    "id": 1,
    "title": "New Task",
    "description": "Description of the new task",
    "completed": false,
    "created_at": "2024-01-01T12:00:00Z",
    "updated_at": "2024-01-01T12:00:00Z"
  }
  ```

#### 2. Получение списка всех задач

- **URL:** `/tasks`
- **Метод:** `GET`
- **Ответ:** (200 OK)

  ```json
  [
    {
      "id": 1,
      "title": "New Task",
      "description": "Description of the new task",
      "completed": false,
      "created_at": "2024-01-01T12:00:00Z",
      "updated_at": "2024-01-01T12:00:00Z"
    }
  ]
  ```

#### 3. Получение задачи по ID

- **URL:** `/task?id=1`
- **Метод:** `GET`
- **Ответ:** (200 OK)

  ```json
  {
    "id": 1,
    "title": "New Task",
    "description": "Description of the new task",
    "completed": false,
    "created_at": "2024-01-01T12:00:00Z",
    "updated_at": "2024-01-01T12:00:00Z"
  }
  ```

- **Ошибка:** (400 Bad Request)

  ```json
  {
    "error": "ID query parameter is missing"
  }
  ```

#### 4. Обновление статуса задачи

- **URL:** `/task/:id`
- **Метод:** `PATCH`
- **Тело запроса:** (JSON)

  ```json
  {
    "completed": true
  }
  ```

- **Ответ:** (200 OK)

  ```json
  {
    "id": 1,
    "title": "New Task",
    "description": "Description of the new task",
    "completed": true,
    "created_at": "2024-01-01T12:00:00Z",
    "updated_at": "2024-01-01T14:00:00Z"
  }
  ```

#### 5. Удаление задачи

- **URL:** `/task/:id`
- **Метод:** `DELETE`
- **Ответ:** (200 OK)

  ```json
  {
    "message": "Task removed"
  }
  ```

### Примечания:
- Все задачи хранятся в памяти, поэтому они будут удалены после перезапуска сервера.
- Этот проект можно использовать как учебный пример для работы с HTTP API на Go с использованием Gin.
