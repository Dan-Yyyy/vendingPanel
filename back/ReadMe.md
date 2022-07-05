### Скачать образ postgres
`docker pull postgres`

### Запустить образ
`docker run --name=vending-panel -e POSTGRES_PASSWORD="postgres" -p 5432:5432 -d --rm postgres`

### Установить утилиту миграций, если не установлена
`Set-ExecutionPolicy RemoteSigned -Scope CurrentUser`

`irm get.scoop.sh | iex`

`scoop install migrate`


### Запустить миграции
`migrate -path ./migrations -database 'postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable' up`

### Запустить бекенд
`go run ./back/cmd/main.go`

### Команда для генерации документации
`swag init -g ./back/cmd/main.go`

### Документация находится по адресу
`http://localhost:8000/swagger/index.html`