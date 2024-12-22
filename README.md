# _**Итоговая задача модуля 1 курса Яндекс Лицей.**_ 👀
### **Веб-сервис для вычисления арифметических выражений**
___
## _Описание_
### Проект реализует веб-сервис,который вычисляет арифметические выражения,переданные пользователем через HTTP-запрос.
___
## Структура Проекта
+ ### cmd/ - точка входа
  + ### main.go
+ ### internal/ - логика приложения
  + ### application_test.go - тест приложения
  + ### application.go
+ ### pkg/ - вспомогательный пакет.
  + ### calculator.go - консольный калькулятор.
## _Запуск сервера_
## 1. Скопируйте проект с GitHub:
```
git clone https://github.com/your-username/calc_service.git
```
## 2.Перейдите в папку проекта и запустите сервер:
```
go run ./cmd/main.go
```
## 3.Сервис будет доступен по адресу: http://localhost:8080/api/v1/calculate.

> [!IMPORTANT]
> ### Эндпоинты
> | Эндпоинт | Методы | Описание |
> | --- | --- | --- |
> | /api/v1/calculate | POST | Получает POST-запрос c телом запроса в формате JSON, содержащим выражение. Отдает результат или ошибку в формате JSON |
## Пример запроса с использованием PowerShell
### Пример успешного запроса
```
Invoke-RestMethod -Uri "http://localhost:8080/api/v1/calculate" `
-Method POST `
-Headers @{"Content-Type"="application/json"} `
-Body '{"expression": "2+2*2"}'

{
  "result": "6.000000"
}
```
### Пример ошибки 500
```
Invoke-RestMethod -Uri "http://localhost:8080/api/v1/calculate" `
-Method POST `
-Headers @{"Content-Type"="application/json"} `
-Body '{"expression": "1+$2"}'

{
  "error": "Некорректное выражение"
}
```
# Примеры запросов с POSTMAN
## Успешный запрос:
```
curl --location 'localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2+2*2"
}'
```
## Ошибка 422(необрабатываемый объект):
```
curl --location 'localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2/0"
}'
```
##  Ошибка 500(внутренняя ошибка сервера).
### Реализована с помощью символа $.
```
curl --location 'localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "$"
}'
```
# Пример использования с cURL:
cURL команда                                   | Ответ                                     | HTTP код
|------------------------------------------------|-------------------------------------------| ----------------------------- |
| ```curl -XPOST -d '{ "expression" : "2 + 2*2"}' 'http://localhost:8080/api/v1/calculate'```  | ```{"result":6} ``` | 200 |
| ```curl -XPOST -d '{ "expression" : "1-"}' 'http://localhost:8080/api/v1/calculate'``` | ```{"error":"invalid expression"}```|422|
| ```curl -XGET -d '{ "expression" : "2 -2"}```' 'http://localhost:8080/api/v1/calculate' | ```{"error":"Invalid request method"} ```|405|
| ```curl -XPOST -d '{ "expression" : "1-1"'``` 'http://localhost:8080/api/v1/calculate' | ```{"error":"invalid request body"}```|400|


