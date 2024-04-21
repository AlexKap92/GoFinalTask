# Итоговое Домашнее задание 

- Реализован REST API для управления пользователями и их транзакциями.
- База данных создана на ресурсе [elephantsql.com](https://elephantsql.com)
- Подключение к базе данных реализовано ч.з. конфигурационный файл [configs.yaml](./configure/configs.yaml)
- Включено логирование(пакет logrus)
- Подготовлены unit тесты

## Запуск приложения

В Терминале выполнить 

go run .

## Пользователи
### Создание пользователей

curl -v -X GET http://localhost:8089/users

curl -v -X POST http://localhost:8089/users -H "Content-Type: application/json" -d '{"name":"John1","email":"john1@gmail.com", "password":"Qwerty1"}'

curl -v -X POST http://localhost:8089/users -H "Content-Type: application/json" -d '{"name":"John2","email":"john2@gmail.com", "password":"Qwerty1"}'

curl -v -X POST http://localhost:8089/users -H "Content-Type: application/json" -d '{"name":"John3","email":"john3@gmail.com", "password":"Qwerty1"}'

### Удаление пользователей

curl -v -X DELETE http://localhost:8089/users/1 -H "Content-Type: application/json" -d '{}'

curl -v -X DELETE http://localhost:8089/users/2 -H "Content-Type: application/json" -d '{}'

curl -v -X GET http://localhost:8089/users

## Транзакции
### Добавление Транзакции

curl -v -X POST http://localhost:8089/transactions --header "Content-Type: application/json" --data '{"user_id":1, "amount":2000.20,"currency":"RUB","type":"income","category":"Зарплата","description":"Описание 1"}' 

curl -v -X POST http://localhost:8089/transactions --header "Content-Type: application/json" --data '{"user_id":2, "amount":1000.10,"currency":"USD","type":"transfer","category":"Зарплата", "description":"Описание 2"}' 

curl -v -X POST http://localhost:8089/transactions --header "Content-Type: application/json" --data '{"user_id":3, "amount":5000.00,"currency":"RUB","type":"перевод","category":"Зарплата", "description":"Описание 3"}' 

### Обновление Транзакции
curl -v -X GET http://localhost:8089/transactions

curl -v -X PUT http://localhost:8089/transactions/3 --header "Content-Type: application/json" --data '{"user_id":3, "amount":6000.00,"currency":"RUB","type":"перевод","category":"Зарплата","date":"2024-02-03","description":"Описание 3_1"}' 

### Удаление Транзакции

curl -v -X DELETE http://localhost:8089/transactions/1 --header "Content-Type: application/json"

curl -v -X GET http://localhost:8089/transactions --header "Content-Type: application/json"

curl -v -X DELETE http://localhost:8089/transactions/1000 --header "Content-Type: application/json"


### Просмотр всех транзакций

curl -v http://localhost:8089/transactions -H "Content-Type: application/json"

### Просмотр транзакции

curl -v http://localhost:8089/transactions/1 -H "Content-Type: application/json"

### Просмотр транзакции с указанием валюты для конвертации суммы

curl -v http://localhost:8089/transaction/1/USD -H "Content-Type: application/json"

## Логирование

Подключено логирование(пакет logrus), Включен режим debugging, вывод логов в консоль. 

> настройка в init package main.

## Unit тесты

Из корневой директории проекта запустить тесты:
go test -v ./...
