# proxyYar

## Реализовано
Проксирование http +
Проксирование https +
Повторная отправка проксированных запросов +
Command injection -

## Инструкция
Добавить сертификат rootCert.cert

git clone git@github.com:AzisHop/proxyYar.git
sudo docker build . -t yaro
sudo docker run -p 8080:8080 -p 8081:8081 --name yaro -t yaro

## Функциональность repeater
localhost:8081/requests - история всех запросов
localhost:8081/request/{id} - повторяет запрос по id запроса


