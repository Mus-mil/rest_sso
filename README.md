# Мини-чат

(проект еще не доделан)

### Стек технологий

Backend: Golang
Frontend: Html, Css
Database: PostgreSQL, pgAdmin

### Запуск веб-сайта


добавьте необходимые зависимости:
```bash
make download
```

Добавьте конфигурации для базы данных в файле config.yaml по маршруту configs
Добавьте файл .env и введите пароль для базы данных
```
DB_PASSWORD=
```
или напишите в терминале bash:
```
export DB_PASSWORD=
```

Запустите веб-сайт:
```bash
make
```
сайт доступен по хосту:
```
http://localhost:8080
```
