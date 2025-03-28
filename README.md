# Rest sso

### Стек технологий

Backend: Golang
Frontend: Html, Css
Database: PostgreSQL, pgAdmin

### Запуск веб-сайта


1. добавьте необходимые зависимости:
```bash
make download
```

2. Добавьте конфигурации для базы данных в файле config.yaml по маршруту configs
Добавьте файл .env и введите пароль для базы данных
```
DB_PASSWORD=
```
или напишите в терминале bash:
```
export DB_PASSWORD=
```

3. Запустите веб-сайт:
```bash
make
```

ИЛИ

2. Запустите веб-сайт через докер:
```bash
make build$$dokcer
```
~~~~
сайт доступен по хосту:
```
http://localhost:8080
```
