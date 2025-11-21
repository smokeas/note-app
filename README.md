Шаг 1. Подготовка проекта
# Создаём структуру папок
mkdir -p note-app/{cmd/app,internal/{api,service,storage/{postgres,redis},events},config,docker}

# Переходим в проект
cd note-app

# Инициализируем Go-модуль
go mod init note-app

Шаг 2. Установка зависимостей
go get github.com/spf13/viper!

Эта команда:

Скачает библиотеку Viper
Добавит запись в go.mod
Сгенерирует go.sum с контрольными суммами
Шаг 3. Создание файлов
Создай все файлы из структуры, скопировав код, который мы разбирали.

Шаг 4. Запуск зависимостей
# Запускаем Postgres и Redis
docker-compose up -d
Флаг -d означает "detached mode" — сервисы будут работать в фоне.

Шаг 5. Запуск приложения
# Собираем и запускаем
go run cmd/app/main.go

В консоли должно появиться:
2023/11/21 14:30:00 Starting server on port 8080

Шаг 6. Проверка работы
Открой в браузере:

http://localhost:8080/register
http://localhost:8080/login
Должны отобразиться соответствующие надписи.




cd D:\note-app
docker compose -f "docker\docker-compose.yml" up -d

go run cmd/app/main.go