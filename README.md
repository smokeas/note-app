Шаг 1. Подготовка проекта
# Создаём структуру папок
mkdir -p note-app/{cmd/app,internal/{api,service,storage/{postgres,redis},events},config,docker}

# Переходим в проект
cd note-app

# Инициализируем Go-модуль
go mod init note-app

Шаг 2. Установка зависимостей
go get github.com/spf13/viper

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




Докер 
docker compose -f .\docker\docker-compose.yml up -d --build

или 

cd .\docker
docker compose up -d --build
# или
docker compose up -d


go mod init github.com/you/note-app        
go get github.com/golang-jwt/jwt/v5
go get golang.org/x/crypto/bcrypt


Как запускать (локально)

Запусти Postgres и применить миграцию (можно через psql):

# если у тебя docker-compose поднял postgres:
psql "postgres://postgres:secret@localhost:5432/noteapp?sslmode=disable" -c "\dt"
# чтобы применить миграции — используй psql < migrations/0001_init.up.sql
psql "postgres://postgres:secret@localhost:5432/noteapp?sslmode=disable" -f migrations/0001_init.up.sql


Экспортируй переменные (пример в PowerShell):

psql "postgres://postgres:secret@localhost:5432/noteapp?sslmode=disable" -f migrations/0001_init.up.sql


Запусти сервис:

go run cmd/app/main.go





Тестирование через curl

Регистрация:

curl -i -X POST http://localhost:8080/register \
  -H "Content-Type: application/json" \
  -d '{"email":"test@example.com","password":"12345"}'


Ожидаемый ответ: 201 Created + JSON {id, email}.

Логин:

curl -i -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{"email":"test@example.com","password":"12345"}'


Ожидаемый ответ: 200 OK + JSON с полем token.

Доступ к защищённому ресурсу (пример):

Допустим, у тебя есть endpoint /me, который возвращает инфо о пользователе и защищён AuthMiddleware. Тогда:

curl -i http://localhost:8080/me -H "Authorization: Bearer <token>"


cd D:\note-app
docker compose -f .\docker\docker-compose.yml up -d

docker ps
# или посмотреть логи
docker compose -f .\docker\docker-compose.yml logs postgres --tail=100

# зайдём внутрь контейнера и выполним команду psql
docker exec -it $(docker ps --filter "name=postgres" -q) psql -U postgres -d noteapp -c "\dt"

go run cmd/app/main.go
