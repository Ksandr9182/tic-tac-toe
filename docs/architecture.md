English
-----------------------------------------------------------------------------------------------------------------------------------------------------------------
# Architecture

This project follows Domain-Driven Design (DDD) principles with a clear separation of concerns:

## Domain Layer
- Located in `domain/`.
- Contains the basic logic of the game "Tic Tac Toe"
- Key entity: `Game` struct with methods like `MakeMove`, `Reset`, `checkWinner` and switchPlayer.
- Defines rules for moves, winner detection, and player switching.

## Application Layer
- Located in `application/`.
- Coordinates use cases (create game, make move, reset).
- `GameService` uses the `GameRepository` interface (port) to interact with storage.
- Ensures domain logic is applied correctly and persists changes.

## Infrastructure Layer
- Located in `infrastructure/` and `transport/`.
- `infrastructure/inmemory_repository.go`: Implements `GameRepository` for in-memory storage.
- `transport/http.go`: HTTP adapter handling REST API requests and responses.

## Interaction Flow
1. HTTP request → `GameHandler` (adapter) → `GameService` (application) → `Game` (domain).
2. Domain updates state and returns results.
3. Application saves state via repository.
4. Handler formats response as JSON.

This structure allows easy replacement of storage (e.g., database) or transport (e.g., WebSocket) without changing the domain.


Russian
-----------------------------------------------------------------------------------------------------------------------------------------------------------------
# Архитектура

Этот проект следует принципам предметно-ориентированного проектирования (DDD) с чётким разделением ответственности:

## Доменный уровень
- Расположен в domain/.
- Содержит основную логику игры "Крестики-нолики".
- Ключевая сущность: структура Game с методами MakeMove, Reset, checkWinner и switchPlayer
- Определяет правила ходов, проверку победителя и переключение игроков.

## Уровень приложения
- Расположен в application/.
- Координирует сценарии использования (создание игры, выполнение хода, сброс и загрузка игры).
- GameService использует интерфейс GameRepository (порт) для взаимодействия с хранилищем.
- Обеспечивает правильное применение доменной логики и сохранение изменений.

## Инфраструктурный уровень
- Расположен в infrastructure/ и transport/.
- infrastructure/inmemory_repository.go: реализация GameRepository для хранения в памяти.
- transport/http.go: HTTP-адаптер, обрабатывающий REST API-запросы и ответы.


## Поток взаимодействия
1. HTTP-запрос → GameHandler (адаптер) → GameService (приложение) → Game (домен).
2. Домен обновляет состояние и возвращает результат.
3. Приложение сохраняет состояние через репозиторий.
4. Обработчик форматирует ответ в JSON.

Такая структура позволяет легко заменить хранилище (например, на базу данных) или транспорт (например, на WebSocket) без изменения доменной логики.