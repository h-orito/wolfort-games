package db

import (
	"time"
	model "wolfort-games/domain/model"
)

type Player struct {
	ID         uint32
	PlayerName string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (p Player) ToModel() *model.Player {
	return &model.Player{
		ID:   p.ID,
		Name: p.PlayerName,
	}
}
