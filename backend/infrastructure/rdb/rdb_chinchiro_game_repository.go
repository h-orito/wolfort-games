package db

import (
	model "chat-role-play/domain/model"
	"chat-role-play/util/array"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type ChinchiroGameRepository struct {
	db *DB
}

func NewChinchiroGameRepository(db *DB) *ChinchiroGameRepository {
	return &ChinchiroGameRepository{
		db: db,
	}
}

func findRdbChinchiroGames(db *gorm.DB, query model.ChinchiroGamesQuery) (games []ChinchiroGame, err error) {
	var rdbs []ChinchiroGame
	result := db.Model(&ChinchiroGame{})
	if query.Paging != nil {
		result = result.Scopes(Paginate(query.Paging))
	} else {
		result = result.Scopes(Paginate(&model.PagingQuery{
			PageSize:   10000,
			PageNumber: 1,
			Desc:       false,
		}))
	}
	if query.IDs != nil {
		result = result.Where("id in (?)", *query.IDs)
	}
	if query.RoomID != nil {
		result = result.Where("room_id = ?", *query.RoomID)
	}
	if query.Name != nil {
		result = result.Where("game_name like ?", fmt.Sprintf("%%%s%%", *query.Name))
	}
	if query.Statuses != nil {
		statuses := array.Map(*query.Statuses, func(s model.ChinchiroGameStatus) string {
			return s.String()
		})
		result = result.Where("game_status_code in (?)", statuses)
	}
	result = result.Find(&rdbs)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if result.Error != nil {
		return nil, fmt.Errorf("failed to find: %s \n", result.Error)
	}

	return rdbs, nil
}
