package repository

import "gordle2/internal/session"

type GameRepository struct {
	storage map[session.GameID]session.Game
}