# Scalable-E-Commerce-Platform-models

Этот проект содержит модели данных для [Scalable E-Commerce Platform](https://github.com/DurkaVerder/Scalable-E-Commerce-Platform). <br>
Модели описывают основные сущности, используемые в системе, такие как пользователи, продукты, заказы и уведомления.

## Структура проекта

- **User**: Представляет пользователя системы с такими полями, как ID, имя пользователя, email и пароль.
- **Product**: Описывает продукт с полями ID, название, цена, описание и количество.
- **Order**: Представляет заказ, включая ID, ID пользователя, сумму, статус, дату создания и список продуктов в заказе.
- **OrderProduct**: Описывает продукт в заказе с названием, ценой и количеством.
- **Notification**: Используется для отправки уведомлений с полями email, тема и тело сообщения.

