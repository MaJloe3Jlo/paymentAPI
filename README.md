`EN` # Payment REST-API by m3 v.0.0.1

## Task ##

Need to make an HTTP REST API only on default packages of Go.
The app need includes 3 methods: Index, Block и Charge.
The app need to recieve incoming POST requests by 7000 port.
The app need recieve JSON requests.
No DB support - sving in memory.

## Start ##
```
go run main.go
```
POST-methods /block/, /charge/. СontentType: application/json. Control requests in folder ./jsons.
For testing you can use curl or postman.

## Additional ##

Some tests attached in folder "test"


`RU` # Платежное REST-API от m3 вер.0.0.1

## Задача ##

Необходимо реализовать HTTP REST API сервис на базе исключительно стандартных библиотек Go.
Сервис должен реализовать три метода: Index, Block и Charge.
Сервис должен принимать входящие POST запросы на 7000 порту.
Сервис должен вести общение в формате JSON.
Соединение с БД реализовывать не нужно - все операции проводить "мнимо".

## Запуск ##

go run main.go

POST-методы /block/, /charge/. СontentType: application/json. Контрольные запросы находятся в папке ./jsons.
Для тестирования можно использовать curl или postman.

## Дополнения ##

Представлено несколько тестов в папке test
