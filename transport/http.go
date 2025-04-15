package transport

import (
	"encoding/json"
	"net/http"
	"strings"

	"tic_tac_toe/application"
	"tic_tac_toe/domain"
)

// GameHandler handles HTTP requests for the game (обрабатывает HTTP-запросы для игры)
type GameHandler struct {
	service *application.GameService
}

// NewGameHandler creates a new game handler (создает новый обработчик игры)
func NewGameHandler(service *application.GameService) *GameHandler {
	return &GameHandler{service: service}
}

// RegisterRoutes registers the HTTP routes (регистрирует HTTP-маршруты)
func (h *GameHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", h.handleRoot)
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	mux.HandleFunc("/game/", h.handleGame)
}

// handleRoot serves the main HTML page (обслуживает главную HTML-страницу)
func (h *GameHandler) handleRoot(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/index.html")
}

// handleGame routes game-related HTTP requests based on method (маршрутизирует HTTP-запросы, связанные с игрой, на основе метода)
func (h *GameHandler) handleGame(w http.ResponseWriter, r *http.Request) {
	gameID := strings.TrimPrefix(r.URL.Path, "/game/")
	if gameID == "" {
		gameID = "default"
	}

	switch r.Method {
	case http.MethodGet:
		h.getGame(w, r, gameID)
	case http.MethodPost:
		h.makeMove(w, r, gameID)
	case http.MethodDelete:
		h.resetGame(w, r, gameID)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

// getGame retrieves the game state by ID (извлекает состояние игры по идентификатору)
func (h *GameHandler) getGame(w http.ResponseWriter, r *http.Request, gameID string) {
	game, err := h.service.GetGame(gameID)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			game, err = h.service.CreateGame(gameID)
			if err != nil {
				http.Error(w, "failed to create game", http.StatusInternalServerError)
				return
			}
		} else {
			http.Error(w, "failed to get game", http.StatusInternalServerError)
			return
		}
	}
	h.writeGame(w, game)
}

// makeMove processes a move request for the game (обрабатывает запрос хода для игры)
func (h *GameHandler) makeMove(w http.ResponseWriter, r *http.Request, gameID string) {
	var req struct {
		Index int `json:"index"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	game, err := h.service.MakeMove(gameID, req.Index)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	h.writeGame(w, game)
}

// resetGame resets the game state by ID (сбрасывает состояние игры по ID)
func (h *GameHandler) resetGame(w http.ResponseWriter, r *http.Request, gameID string) {
	game, err := h.service.ResetGame(gameID)
	if err != nil {
		http.Error(w, "failed to reset game", http.StatusInternalServerError)
		return
	}
	h.writeGame(w, game)
}

// writeGame writes the game state as a JSON response (записывает состояние игры как ответ JSON)
func (h *GameHandler) writeGame(w http.ResponseWriter, game *domain.Game) {
	resp := struct {
		Board        [9]string `json:"board"`
		Turn         string    `json:"turn"`
		Winner       string    `json:"winner"`
		WinningCells []int     `json:"winning_cells"`
	}{
		Board:        game.Board(),
		Turn:         string(game.CurrentPlayer()),
		Winner:       string(game.Winner()),
		WinningCells: game.WinningCells(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
