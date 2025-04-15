package application

import (
	"tic_tac_toe/domain"
)

// GameRepository defines the interface for game storage (определяет интерфейс для хранения игр)
type GameRepository interface {
	Save(game *domain.Game) error
	FindByID(id string) (*domain.Game, error)
}

// GameService manages game operations (управляет игровыми операциями)
type GameService struct {
	repo GameRepository
}

// NewGameService creates a new game service (создает новый игровой сервис)
func NewGameService(repo GameRepository) *GameService {
	return &GameService{repo: repo}
}

// CreateGame creates a new game with the given ID (создает новую игру с указанным идентификатором)
func (s *GameService) CreateGame(id string) (*domain.Game, error) {
	game := domain.NewGame(id)
	if err := s.repo.Save(game); err != nil {
		return nil, err
	}
	return game, nil
}

// MakeMove makes a move in the game with the given ID (делает ход в игре с указанным идентификатором)
func (s *GameService) MakeMove(gameID string, index int) (*domain.Game, error) {
	game, err := s.repo.FindByID(gameID)
	if err != nil {
		return nil, err
	}

	if err := game.MakeMove(index); err != nil {
		return nil, err
	}

	if err := s.repo.Save(game); err != nil {
		return nil, err
	}
	return game, nil
}

// GetGame retrieves the game by ID (извлекает игру по идентификатору)
func (s *GameService) GetGame(id string) (*domain.Game, error) {
	return s.repo.FindByID(id)
}

// ResetGame resets the game with the given ID (сбрасывает игру с указанным ID)
func (s *GameService) ResetGame(id string) (*domain.Game, error) {
	game, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	game.Reset()
	if err := s.repo.Save(game); err != nil {
		return nil, err
	}
	return game, nil
}
