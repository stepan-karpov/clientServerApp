# 1C-department-selection
This repository destined to implement and submit 1C-department selection competition

Общий пайплайн работы с приложением:

1) Запустить приложение-сервер. Оно запустится на 8080 порте, инициализирует базу данных и начнет слушать клиентов на предмет подписки. При этом при добавлении каждого нового клиента будет перерисовываться таблица с участниками, зарегистрированными на эксперимент

Общие заметки:

Статистики по экспериментам легко масштабируются
Можно добавить отписку
Почему клиенты идентифицуруются по IP
Поддержка concurrency