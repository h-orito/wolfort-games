package db

import (
	"context"
	"errors"
	"fmt"
	model "wolfort-games/domain/model"
	"wolfort-games/util/array"

	"gorm.io/gorm"
)

type PlayerRepository struct {
	db *DB
}

func NewPlayerRepository(db *DB) model.PlayerRepository {
	return &PlayerRepository{
		db: db,
	}
}

func (repo *PlayerRepository) FindPlayers(query model.PlayersQuery) ([]model.Player, error) {
	return repo.findPlayers(repo.db.Connection, query)
}

func (repo *PlayerRepository) Find(ID uint32) (_ *model.Player, err error) {
	return repo.findPlayer(repo.db.Connection, ID)
}

func (repo *PlayerRepository) FindByName(name string) (_ *model.Player, err error) {
	rdbPlayer, err := repo.findRdbPlayerByName(repo.db.Connection, name)
	if err != nil {
		return nil, err
	}
	if rdbPlayer == nil {
		return nil, nil
	}
	return rdbPlayer.ToModel(), nil
}

func (repo *PlayerRepository) FindByUserName(username string) (_ *model.Player, err error) {
	return repo.findByUserName(repo.db.Connection, username)
}

func (repo *PlayerRepository) Save(ctx context.Context, p *model.Player) (_ *model.Player, err error) {
	tx, ok := GetTx(ctx)
	if !ok {
		return nil, fmt.Errorf("failed to get tx from context")
	}

	rdbPlayer, err := repo.findRdbPlayer(tx, p.ID)
	if err != nil {
		return nil, err
	}
	if rdbPlayer == nil {
		return nil, nil
	}
	rdbPlayer.PlayerName = p.Name
	result := tx.Save(&rdbPlayer)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to save: %s", result.Error)
	}
	return repo.findPlayer(tx, rdbPlayer.ID)
}

// ----------------------------------------------------------------------------

func (repo *PlayerRepository) findPlayers(db *gorm.DB, query model.PlayersQuery) ([]model.Player, error) {
	rdbPlayers, err := repo.findRdbPlayers(db, query)
	if err != nil {
		return nil, err
	}
	return array.Map(rdbPlayers, func(p Player) model.Player {
		return *p.ToModel()
	}), nil
}

func (repo *PlayerRepository) findPlayer(db *gorm.DB, ID uint32) (_ *model.Player, err error) {
	rdbPlayer, err := repo.findRdbPlayer(db, ID)
	if err != nil {
		return nil, err
	}
	if rdbPlayer == nil {
		return nil, nil
	}

	return rdbPlayer.ToModel(), nil
}

func (repo *PlayerRepository) findRdbPlayers(db *gorm.DB, query model.PlayersQuery) ([]Player, error) {
	var rdbPlayers []Player
	result := db.Model(&Player{})
	if query.IDs != nil {
		result = result.Where("id IN (?)", *query.IDs)
	}
	if query.Name != nil {
		result = result.Where("player_name LIKE ?", fmt.Sprintf("%%%s%%", *query.Name))
	}
	if query.Paging != nil {
		result = result.Scopes(Paginate(query.Paging))
	} else {
		result = result.Scopes(Paginate(&model.PagingQuery{
			PageSize:   10000,
			PageNumber: 1,
			Desc:       false,
		}))
	}
	result = result.Find(&rdbPlayers)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if result.Error != nil {
		return nil, fmt.Errorf("failed to find: %s", result.Error)
	}
	return rdbPlayers, nil
}

func (repo *PlayerRepository) findRdbPlayer(db *gorm.DB, ID uint32) (_ *Player, err error) {
	var rdbPlayer Player
	result := db.Model(&Player{}).First(&rdbPlayer, ID)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if result.Error != nil {
		return nil, fmt.Errorf("failed to find: %s", result.Error)
	}
	return &rdbPlayer, nil
}

func (repo *PlayerRepository) findRdbPlayerByName(db *gorm.DB, name string) (_ *Player, err error) {
	var rdbPlayer Player
	result := db.Model(&Player{}).Where("player_name = ?", name).First(&rdbPlayer)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if result.Error != nil {
		return nil, fmt.Errorf("failed to find: %s", result.Error)
	}
	return &rdbPlayer, nil
}

func (repo *PlayerRepository) findByUserName(db *gorm.DB, username string) (_ *model.Player, err error) {
	var rdbPlayerAccount PlayerAccount
	result := repo.db.Connection.
		Model(&PlayerAccount{}).
		Where("user_name = ?", username).
		First(&rdbPlayerAccount)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if result.Error != nil {
		return nil, fmt.Errorf("failed to find: %s", result.Error)
	}
	return repo.findPlayer(db, rdbPlayerAccount.PlayerID)
}
