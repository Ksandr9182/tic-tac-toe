package domain

import "errors"

// Player represents a player in the game ("X" or "O") (представляет игрока в игре («X» или «O»)
type Player string

const (
	PlayerX Player = "X"
	PlayerO Player = "O"
)

// Game represents the Tic-Tac-Toe game state and logic (представляет собой состояние и логику игры «Крестики-нолики»)
type Game struct {
	id            string
	board         [9]string
	currentPlayer Player
	winner        Player
	winningCells  []int
}

// NewGame creates a new game with the given ID (создает новую игру с указанным идентификатором)
func NewGame(id string) *Game {
	return &Game{
		id:            id,
		currentPlayer: PlayerX,
	}
}

// ID returns the game ID (возвращает идентификатор игры)
func (g *Game) ID() string {
	return g.id
}

// Board returns the current board state (возвращает текущее состояние доски)
func (g *Game) Board() [9]string {
	return g.board
}

// CurrentPlayer returns the current player (возвращает текущего игрока)
func (g *Game) CurrentPlayer() Player {
	return g.currentPlayer
}

// Winner returns the winner, if any (возвращает победителя, если таковой имеется)
func (g *Game) Winner() Player {
	return g.winner
}

// WinningCells returns the winning cells, if any (возвращает выигрышные ячейки, если таковые имеются)
func (g *Game) WinningCells() []int {
	return g.winningCells
}

// MakeMove attempts to make a move at the given index (пытается сделать ход по указанному индексу)
func (g *Game) MakeMove(index int) error {
	if index < 0 || index >= 9 {
		return errors.New("index out of bounds")
	}
	if g.board[index] != "" {
		return errors.New("cell already occupied")
	}
	if g.winner != "" {
		return errors.New("game already finished")
	}

	g.board[index] = string(g.currentPlayer)
	g.checkWinner()
	g.switchPlayer()
	return nil
}

// Reset resets the game state (сбрасывает состояние игры)
func (g *Game) Reset() {
	g.board = [9]string{}
	g.currentPlayer = PlayerX
	g.winner = ""
	g.winningCells = nil
}

// switchPlayer changes player (меняет игрока)
func (g *Game) switchPlayer() {
	if g.currentPlayer == PlayerX {
		g.currentPlayer = PlayerO
	} else {
		g.currentPlayer = PlayerX
	}
}

// checkWinner checks winning combinations (проверяет выигрышные комбинации)
func (g *Game) checkWinner() {
	winningCombinations := [][]int{
		{0, 1, 2}, {3, 4, 5}, {6, 7, 8},
		{0, 3, 6}, {1, 4, 7}, {2, 5, 8},
		{0, 4, 8}, {2, 4, 6},
	}

	for _, combo := range winningCombinations {
		if g.board[combo[0]] != "" &&
			g.board[combo[0]] == g.board[combo[1]] &&
			g.board[combo[0]] == g.board[combo[2]] {
			g.winner = Player(g.board[combo[0]])
			g.winningCells = combo
			return
		}
	}
}
