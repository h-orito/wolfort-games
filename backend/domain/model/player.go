package model

import (
	"context"
)

type Player struct {
	ID   uint32
	Name string
}

type PlayersQuery struct {
	IDs    *[]uint32
	Name   *string
	Paging *PagingQuery
}

type PlayerRepository interface {
	FindPlayers(query PlayersQuery) ([]Player, error)
	Find(ID uint32) (player *Player, err error)
	FindByName(name string) (player *Player, err error)
	FindByUserName(userName string) (player *Player, err error)
	Save(ctx context.Context, player *Player) (saved *Player, err error)
}
