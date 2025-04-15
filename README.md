
English
-----------------------------------------------------------------------------------------------------------------------------------------------------------------
# Tic-Tac-Toe

A web-based Tic-Tac-Toe game built with Go and JavaScript, designed using Domain-Driven Design (DDD) principles.

## Overview

This project is a single-player Tic-Tac-Toe game where users can alternate between X and O moves on a 3x3 board. The backend is implemented in Go, providing a REST API to manage game state, while the frontend uses JavaScript to render the game board and handle user interactions. The project showcases a clean architecture with DDD, separating concerns into domain, application, infrastructure, and transport layers.

### Features
- Interactive 3x3 game board with X and O moves.
- Real-time winner detection with highlighted winning cells.
- Game reset functionality via a "Restart" button.
- Responsive UI with images for X and O.
- REST API for game operations.
- In-memory storage for game state.

## Requirements
- [Go](https://golang.org/dl/) 1.21 or higher
- A modern web browser (e.g., Chrome, Firefox, Edge)
- No external dependencies required

## Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/tic-tac-toe.git
   cd tic-tac-toe


Verify Go is installed:
-- go version

################################### Running the Game ###################################
1. Start the server:
-- go run main.go

You should see:
-- Server running on http://localhost:8080
2. Open your browser and navigate to:
http://localhost:8080

################################### How to Play ###################################

Make a move: Click an empty cell to place X or O (players alternate automatically).
View status: The status bar shows whose turn it is ("Ходит: X") or the winner ("Победитель: X").
Reset game: Click the "Restart" button to start a new game.
Note: Clicks are disabled after a winner is determined to prevent further moves.


################################### Project Structure ###################################
The project is organized according to Domain-Driven Design (DDD):

domain/: Core logic for Tic-Tac-Toe.
Defines the Game entity with methods like MakeMove, Reset, and checkWinner.
Implements rules for valid moves, winner detection, and player switching.
application/: Application layer coordinating use cases.
GameService manages operations like creating games, making moves, and resetting.
Uses a GameRepository interface for data access.
infrastructure/: Infrastructure layer.
InMemoryGameRepository implements GameRepository for in-memory storage.
transport/: HTTP layer.
GameHandler processes REST API requests and serves static files (HTML, JS, images).
static/: Frontend assets.
script.js: JavaScript for rendering the board and handling user interactions.
cross.png, zero.png: Images for X and O.
templates/: HTML template (index.html) for the game UI.

################################### API Endpoints ###################################
The backend exposes a REST API for game operations at http://localhost:8080/game/.

Method → GET	Endpoint → /game/	Description → Get current game state		Request Body → None		Response Body → { "board": [9]string, "turn": "X"/"O", "winner": ""/"X"/"O", "winning_cells": []int }
Method → POST	Endpoint → /game/	Description → Make a move at specified index	Request Body → { "index": 0-8 }	Response Body → Same as GET
Method → DELETE	Endpoint → /game/	Description → Reset the game to initial state	Request Body → None		Response Body → Same as GET

Example Response
json
{
  "board": ["", "X", "", "", "O", "", "", "", ""],
  "turn": "X",
  "winner": "",
  "winning_cells": []
}

Error Responses
400 Bad Request: Invalid move (e.g., occupied cell) or malformed JSON.
500 Internal Server Error: Server-side issues (e.g., repository failure).

################################### Development ###################################
To extend or modify the project:

Backend:
Edit Go files in domain/, application/, infrastructure/, or transport/.
Run go mod tidy to sync dependencies (if any are added).

Format code:
-- go fmt ./...

Frontend:
Update static/script.js for UI logic.
Modify templates/index.html for layout changes.
Add new assets (e.g., CSS) to static/.

Testing:
Currently, no unit tests are implemented. To add tests:
-- go test ./...

Build:
Compile the project:
-- go build ./...

###################################### Running with Docker ######################################
You can run the game in a Docker container for easy deployment.

## Prerequisites
-- [Docker](https://www.docker.com/get-started) is installed on your computer

## Steps
1. Build the Docker image:
docker build -t tic-tac-toe .

2. Run the container:
docker run -p 8080:8080 --rm tic-tac-toe

3. Open your browser and navigate to:
http://localhost:8080

Notes
The -p 8080:8080 flag maps port 8080 in the container to port 8080 on your host
The --rm flag removes the container after it is stopped
Ensure that port 8080 on your computer is free


Russian
-----------------------------------------------------------------------------------------------------------------------------------------------------------------
# Крестики-нолики

Веб-игра крестики-нолики, созданная с использованием Go и JavaScript, разработанная с использованием принципов Domain-Driven Design (DDD).

## Обзор

Этот проект представляет собой однопользовательскую игру крестики-нолики, в которой пользователи могут чередовать ходы X и O на доске 3x3. Бэкэнд реализован на Go, предоставляя REST API для управления состоянием игры, в то время как фронтэнд использует JavaScript для визуализации игрового поля и обработки взаимодействий. Проект демонстрирует чистую архитектуру с DDD, разделяя задачи на уровни домена, приложения, инфраструктуры и транспорта.

### Особенности
- Интерактивное игровое поле 3x3 с ходами X и O
- Определение победителя в реальном времени с выделенными выигрышными ячейками
- Функция сброса игры с помощью кнопки «Перезапустить»
- Адаптивный пользовательский интерфейс с изображениями для X и O
- REST API для игровых операций
- Хранилище в памяти для состояния игры

## Требования
- [Go](https://golang.org/dl/) 1.21 или выше
- Современный веб-браузер (например, Chrome, Firefox, Edge)
- Не требуется никаких внешних зависимостей

## Установка
1. Клонируйте репозиторий:
```bash
git clone https://github.com/your-username/tic-tac-toe.git
cd tic-tac-toe

Проверьте, установлен ли Go:
-- go version

#################################### Запуск игры #####################################
1. Запустите сервер:
-- go run main.go

Вы должны увидеть:
-- Server running on http://localhost:8080
2. Откройте браузер и перейдите по адресу:
http://localhost:8080

###################################### Как играть #####################################

Сделать ход: Щелкните пустую ячейку, чтобы поставить X или O (игроки меняются автоматически)
Просмотр статуса: В строке статуса отображается чей ход («Ходит: X») или победитель («Победитель: X»)
Сбросить игру: Нажмите кнопку «Перезапустить», чтобы начать новую игру
Примечание: После определения победителя щелчки отключаются, чтобы предотвратить дальнейшие ходы

######################################## Структура проекта ########################################
Проект организован в соответствии с Domain-Driven Design (DDD):

domain/: Основная логика для Tic-Tac-Toe.
Определяет сущность Game с такими методами, как MakeMove, Reset и checkWinner
Реализует правила для допустимых ходов, определения победителя и переключения игроков
application/: Уровень приложения, координирующий варианты использования
GameService управляет такими операциями, как создание игр, выполнение ходов и сброс
Использует интерфейс GameRepository для доступа к данным
infrastructure/: Уровень инфраструктуры
InMemoryGameRepository реализует GameRepository для хранения в памяти
transport/: Уровень HTTP
GameHandler обрабатывает запросы REST API и обслуживает статические файлы (HTML, JS, изображения)
static/: Ресурсы интерфейса
script.js: JavaScript для рендеринга доски и обработки взаимодействий с пользователем
cross.png, zero.png: Изображения для X и O
templates/: HTML-шаблон (index.html) для игрового пользовательского интерфейса

##################################### Конечные точки API ####################################
Бэкэнд предоставляет REST API для игровых операций по адресу http://localhost:8080/game/.

Method → GET	Endpoint → /game/ Получить текущее состояние игры	Request Body → None		Response Body → { "board": [9]string, "turn": "X"/"O", "winner": ""/"X"/"O", "winning_cells": []int }
Method → POST	Endpoint → /game/ Сделать ход по указанному индексу	Request Body → { "index": 0-8 }	Response Body → То же, что и GET
Method → DELETE	Endpoint → /game/ Сбросить игру в начальное состояние	Request Body → None		Response Body → То же, что и GET

Пример ответа
json
{
"board": ["", "X", "", "", "O", "", "", "", ""],
"turn": "X",
"winner": "",
"winning_cells": []
}

Ответы об ошибках
400 Неправильный запрос: Недопустимый ход (например, занятая ячейка) или некорректный JSON
500 Внутренняя ошибка сервера: Проблемы на стороне сервера (например, сбой репозитория)

###################################### Разработка ######################################
Чтобы расширить или изменить проект:

Бэкэнд:
Измените файлы Go в domain/, application/, infrastructure/ или transport/.
Запустите go mod tidy, чтобы синхронизировать зависимости (если таковые добавлены)

Код форматирования:
-- go fmt ./...

Фронтенд:
Обновите static/script.js для логики пользовательского интерфейса
Измените templates/index.html для изменения макета
Добавьте новые ресурсы (например, CSS) в static/

Тестирование:
В настоящее время модульные тесты не реализованы. Чтобы добавить тесты:
-- go test ./...

Сборка:
Скомпилируйте проект:
-- go build ./...


###################################### Запуск с Docker ###################################### 
Вы можете запустить игру в контейнере Docker для простого развертывания

## Предварительные условия
-- [Docker](https://www.docker.com/get-started) установлен на вашем компьютере

## Шаги
1. Соберите образ Docker:
docker build -t tic-tac-toe .

2. Запустите контейнер:
docker run -p 8080:8080 --rm tic-tac-toe

3. Откройте браузер и перейдите по адресу:
http://localhost:8080

Примечания
Флаг -p 8080:8080 сопоставляет порт 8080 в контейнере с портом 8080 на вашем хосте
Флаг --rm удаляет контейнер после его остановки
Убедитесь, что порт 8080 на вашем компьютере свободен

