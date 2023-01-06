# Описание

#### Начало работы
1. `git clone https://github.com/srsad/discount-service.git`
2. создать файл `.env` в корне проекта и скопировать в него данные из `example.env`
3. `docker-compose buid`, *последующии запуски `docker-compose up -d`**
    - `docker-compose up -d` запустит сервис postrges (+pgadmin) и `gofiber` - бэк.
    - Для того чтоб попасть в контейнер fiber `docker-compose run --service-ports gofiber bash`
4. поднять миграции в контейнере gofiber - `make migrate-up`

---

#### Работа на бэке
Все новые команды/запуск сервисов создавать/изменять только в Makefile

##### Миграции
[Работа с миграциями описанна тут](MIGRATIONS.md)

---
#### Настройка pgadmin
pgadmin достепен по `http://localhost:5050/browser/`, для его настройки понадобиться:
1. установить master key
2. в главном меню `Object -> Register -> Server`
3. вкладка `General -> Name` (на твое усмотрение)
4. вкладка `Connection`:
    - `Host name/address` = `postgres`
    - `Maintenance database` = `DS`
    - `Username` = `root`
    - `Password` = `root`
    - Можно сохранить пароль

---
#### Работа на фронте
Необходима nodejs версии `16.13.0` или выше. Рекуомендуется установить [nvm](https://github.com/coreybutler/nvm-windows) для безболезненной смены ноды.
  После установки: 
  1. перейти в каталок `frontend`
  2. `npm i`
  3. `npm run dev`
##### Зависимости
- [axios](https://www.npmjs.com/package/axios) - http клиент
- [pinia](https://www.npmjs.com/package/pinia) - стейт менеджер
- [vue-router](https://www.npmjs.com/package/vue-router) - роутинг
- [vue-toastification](https://www.npmjs.com/package/vue-toastification) - уведомления
- [vue-final-modal](https://www.npmjs.com/package/vue-final-modal) - модальные окна

---
**TODO**
  - ВАЖНО! Почитать за различные концепции архитектуры приложений на go + почитать за модульность. 
    - ЧТИВО:
    - [Effective Go](https://go.dev/doc/effective_go)
    - [Go-clean-template: шаблон чистой архитектуры для сервисов на Go](https://evrone.ru/Go-clean-template)
  - 
  - ...

