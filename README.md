# Система бронирования отелей
## Описание проекта
Система предназначена для бронирования номеров в отелях и управления информацией об отелях. Реализована в архитектуре микросервисов с использованием современных подходов и инструментов.

## Основные возможности:
1. **Клиент**: бронирует отели, получает уведомления о статусе бронирования.
2. **Отельер**: управляет информацией об отелях и номерах.
3. **Уведомления**: отправка уведомлений через интеграцию с системой доставки.
4. **Оплата**: интеграция с системой оплаты для имитации обработки транзакций.
5. **Очереди и события**: использование Kafka для обработки событий бронирования.
## Архитектура системы
### Основные компоненты:

1. **Booking Svc (Сервис бронирования)**
Управляет процессом бронирования.
Хранит данные о бронированиях в базе Booking Data.
Отправляет события о новых бронированиях в Kafka.

2. **Hotel Svc (Сервис отелей)**
Управляет информацией об отелях, включая номера и цены.
Хранит данные в базе Hotels Data.
Предоставляет данные для обработки бронирования.

3. **Notification Svc (Сервис уведомлений)**
Обрабатывает события из Kafka.
Отправляет уведомления клиентам и отельерам через внешний сервис доставки (Delivery System).

4. **Delivery System (Система доставки)**
Отвечает за доставку уведомлений (email, SMS и т.д.).
Внешний сервис, с которым производится интеграция.

5. **Payment System (Система оплаты)**
Обрабатывает платежи через интеграцию с Booking Svc.
Возвращает статус оплаты через webhook.

6. **Очередь (Kafka)**
Обеспечивает асинхронную обработку событий.

7. **Базы данных (PostgreSQL)**
Booking Data: хранит данные о бронированиях.
Hotels Data: хранит данные об отелях и номерах.

<img width="3458" alt="components-diagram" src="https://github.com/user-attachments/assets/7b2e95a9-b4a1-44af-82ff-4fe8e4672e40" />
