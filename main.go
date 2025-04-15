package main

import (
	"fmt"
	"log"
	"net/http"

	"tic_tac_toe/application"
	"tic_tac_toe/infrastructure"
	"tic_tac_toe/transport"
)

func main() {
	// Creates a new in-memory game repository (Создает новый репозиторий игры в памяти)
	repo := infrastructure.NewInMemoryGameRepository()
	// Creates a new game service with the repository (Создает новый игровой сервис с репозиторием)
	service := application.NewGameService(repo)
	// Creates a new HTTP handler with the game service (Создает новый HTTP-обработчик с игровым сервисом)
	handler := transport.NewGameHandler(service)

	// Creates a new HTTP serve multiplexer (Создает новый HTTP-мультиплексор)
	mux := http.NewServeMux()
	// Registers HTTP routes with the multiplexer (Регистрирует HTTP-маршруты с помощью мультиплексора)
	handler.RegisterRoutes(mux)

	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", mux))
}
