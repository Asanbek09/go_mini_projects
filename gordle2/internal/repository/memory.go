package repository

import (
	"fmt"
	"gordle2/internal/session"
	"log"
	"sync"
)

type GameRepository struct {
	mutex sync.Mutex
	storage map[session.GameID]session.Game
}

func New() *GameRepository {
	return &GameRepository{
		storage: make(map[session.GameID]session.Game),
	}
}

func (gr *GameRepository) Add(game session.Game) error {
	log.Print("Adding a game...")

	gr.mutex.Lock()
	defer gr.mutex.Unlock()

	_, ok := gr.storage[game.ID]
	if ok {
		return fmt.Errorf("%w (%s)", ErrConflictingID, game.ID)
	}

	gr.storage[game.ID] = game

	return nil
}

func (gr *GameRepository) Find(id session.GameID) (session.Game, error) {
	log.Printf("Looking for game %s...", id)

	game, found := gr.storage[id]
	if !found {
		return session.Game{}, fmt.Errorf("can't find game %s: %w", id, ErrNotFound)
	}

	return game, nil
}

func (gr *GameRepository) Update(game session.Game) error {
	_, found := gr.storage[game.ID]
	if !found {
		return fmt.Errorf("can't find game %s: %w", game.ID, ErrNotFound)
	}

	gr.storage[game.ID] = game
	return nil
}