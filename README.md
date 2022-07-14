# L0

* `docker-compose up -d` для запуска PostgreSQL и Nats-streaming сервера
* `go run ./cmd/service/service.go` для запуска сервиса
* `go run ./cmd/publisher/publisher.go` для отправки в Nats тестовых файлов (валидный и невалидный)

## Для проверки отображения:
  * `http://127.0.0.1:8081/view` простой интерфейс для поиска order по uid
  * `http://127.0.0.1:8081/api/orders/{uid}` выводит json обьект по uid
