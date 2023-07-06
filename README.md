# Накопительная система лояльности «Гофермарт»

Система представляет собой HTTP API со следующими требованиями к бизнес-логике:
регистрация, аутентификация и авторизация пользователей;
приём номеров заказов от зарегистрированных пользователей;
учёт и ведение списка переданных номеров заказов зарегистрированного пользователя;
учёт и ведение накопительного счёта зарегистрированного пользователя;
проверка принятых номеров заказов через систему расчёта баллов лояльности;
начисление за каждый подходящий номер заказа положенного вознаграждения на счёт лояльности пользователя.

#API
*POST /api/user/register — регистрация пользователя (JSON {"login": "<login>","password": "<password>"});

*POST /api/user/login — аутентификация пользователя (JSON {"login": "<login>","password": "<password>"});
*POST /api/user/orders — загрузка пользователем номера заказа для расчёта body (12345678903);
*GET /api/user/orders — получение списка загруженных пользователем номеров заказов, статусов их обработки и информации о начислениях;
*GET /api/user/balance — получение текущего баланса счёта баллов лояльности пользователя;
*POST /api/user/balance/withdraw — запрос на списание баллов с накопительного счёта в счёт оплаты нового заказа;
*GET /api/user/withdrawals — получение информации о выводе средств с накопительного счёта пользователем.

## Аргументы командной строки

* Флаг -a отвечает за адрес запуска HTTP-сервера (по умолчанию: localhost:8080)
* Флаг -d отвечает за строку подключения к БД (postgreSQL: "host=localhost user=postgres password=123 dbname=postgres sslmode=disable")
* Флаг -r адрес системы расчёта начислений (по умолчанию localhost:8081)

## Переменные окружения

Имеется возможность конфигурировать сервис с помощью переменных окружения:

* RUN_ADDRESS - Адрес запуска HTTP-сервера (по умолчанию: localhost:8080)
* DATABASE_URI - адрес подключения к базе данных (postgreSQL: "host=localhost user=postgres password=123 dbname=postgres sslmode=disable")
* ACCRUAL_SYSTEM_ADDRESS — адрес системы расчёта начислений (по умолчанию localhost:8081)
 