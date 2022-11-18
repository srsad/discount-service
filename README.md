# Описание

1. `git clone https://github.com/srsad/discount-service.git`
2. `docker-compose buid`, последующии запуски `docker-compose up`
    1. `docker-compose up` запустит сервис postrges (+pgadmin) и `gofiber` - бэк.
    2. для работы с fiber, надо выпонить `docker-compose run --service-ports gofiber bash` (мб подумать об отдельном запуске)
    3. для работы на фронте, выполнить `docker-compose run --service-ports frontend bash`

**TODO**
  - `+` добавить образы для go(fiber)
  - `+` nodejs(vue3)
      - собрать базу для vue, проверить порты
  - добавиьт описание к разворачиванию окружения (docker-compose buid || docker-compose up)
  - разобраться с pgadmin
  - 
  - ...

**Fiber**
  - добавить тестовый эндпоинт
  - настройки линтера
  - 
  - ...

**Vue**
  - подключить TS
  - настройки линтера
  - выборать библиотеку компонентов
  - проработать архитектуру проекта (учесть composition api + pinia (нужна/нет))
  - создать дэмо страницу/страницы
  - 
  - ...
