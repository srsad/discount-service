# Миграции

#### Создание миграции
```bash
make migrate-create название_миграции
``` 
Например, `make migrate-create users`.
В папке `./migrations` появится два файла:
1. <версия>_users.up.sql
2. <версия>_users.down.sql

#### Применение изменений
Для выполнения вызвать:
```bash
make migrate-up
```
Можно передать версию:
```bash
make migrate-up v=2
```

#### Откат
Для выполнения вызвать:
```bash
make migrate-down
```
Дополнительные аргументы передаются через `a`:
```bash
make migrate-down a=--all
```
Для указания версии:
```bash
make migrate-down a=2
```
