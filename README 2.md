## Реализовано
Проксирование http +
Проксирование https +
Отправка сохраненного запроса +
Command injection -

## Инструкция
Добавить сертификат в rootCert.cert


sudo docker build -t alex https://github.com/AzisHop/proxyYar.git
sudo docker run -p 8080:8080 -p 8081:8081 --name yaro -t yaro

## Функциональность repeater
localhost:8081/requests - выдает информацию о последних 10 запросах

localhost:8081/requests/{id} - повторяет запрос по id запроса


