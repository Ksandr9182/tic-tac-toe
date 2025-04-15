package infrastructure

import (
	"errors"
	"sync"

	"tic_tac_toe/domain"
)

// InMemoryGameRepository stores games in memory (сохраняет игры в памяти)
type InMemoryGameRepository struct {
	games map[string]*domain.Game
	mu    sync.RWMutex
}

// NewInMemoryGameRepository creates a new in-memory repository (создает новый репозиторий в памяти)
func NewInMemoryGameRepository() *InMemoryGameRepository {
	return &InMemoryGameRepository{
		games: make(map[string]*domain.Game),
	}
}

// Save stores the game in memory (сохраняет игру в памяти)
func (r *InMemoryGameRepository) Save(game *domain.Game) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.games[game.ID()] = game
	return nil
}

// FindByID retrieves a game by ID (извлекает игру по идентификатору)
func (r *InMemoryGameRepository) FindByID(id string) (*domain.Game, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	game, exists := r.games[id]
	if !exists {
		return nil, errors.New("game not found")
	}
	return game, nil
}
