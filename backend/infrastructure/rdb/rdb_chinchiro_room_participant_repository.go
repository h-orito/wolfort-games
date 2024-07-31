package db

import (
	"context"
	"errors"
	"fmt"
	model "wolfort-games/domain/model"
	"wolfort-games/util/array"

	"gorm.io/gorm"
)

type ChinchiroRoomParticipantRepository struct {
	db *DB
}

func NewChinchiroRoomParticipantRepository(db *DB) *ChinchiroRoomParticipantRepository {
	return &ChinchiroRoomParticipantRepository{
		db: db,
	}
}

func (repo *ChinchiroRoomParticipantRepository) FindRoomParticipants(query model.ChinchiroRoomParticipantsQuery) (participants model.ChinchiroRoomParticipants, err error) {
	return findChinchiroRoomParticipants(repo.db.Connection, query)
}

func (repo *ChinchiroRoomParticipantRepository) FindRoomParticipant(query model.ChinchiroRoomParticipantQuery) (participant *model.ChinchiroRoomParticipant, err error) {
	return findChinchiroRoomParticipant(repo.db.Connection, query)
}

func (repo *ChinchiroRoomParticipantRepository) RegisterRoomParticipant(ctx context.Context, roomID uint32, participant model.ChinchiroRoomParticipant) (saved *model.ChinchiroRoomParticipant, err error) {
	tx, ok := GetTx(ctx)
	if !ok {
		return nil, fmt.Errorf("failed to get tx from context")
	}
	p, err := registerRdbChinchiroRoomParticipant(tx, roomID, participant)
	if err != nil {
		return nil, err
	}
	return findChinchiroRoomParticipant(tx, model.ChinchiroRoomParticipantQuery{
		ID: &p.ID,
	})
}

func (repo *ChinchiroRoomParticipantRepository) UpdateRoomParticipant(ctx context.Context, participant model.ChinchiroRoomParticipant) (err error) {
	tx, ok := GetTx(ctx)
	if !ok {
		return fmt.Errorf("failed to get tx from context")
	}
	return updateRdbChinchiroRoomParticipant(tx, participant)
}

func (repo *ChinchiroRoomParticipantRepository) DeleteRoomParticipant(ctx context.Context, participantID uint32) (err error) {
	tx, ok := GetTx(ctx)
	if !ok {
		return fmt.Errorf("failed to get tx from context")
	}
	return deleteRdbChinchiroRoomParticipant(tx, participantID)

}

// -----------------------

func findChinchiroRoomParticipants(db *gorm.DB, query model.ChinchiroRoomParticipantsQuery) (participants model.ChinchiroRoomParticipants, err error) {
	rdbParticipants, err := findRdbChinchiroRoomParticipants(db, query)
	if err != nil {
		return model.ChinchiroRoomParticipants{}, err
	}
	if rdbParticipants == nil {
		return model.ChinchiroRoomParticipants{}, nil
	}
	return model.ChinchiroRoomParticipants{
		Count: len(rdbParticipants),
		List: array.Map(rdbParticipants, func(r ChinchiroRoomParticipant) model.ChinchiroRoomParticipant {
			return *r.ToModel()
		}),
	}, nil
}

func findChinchiroRoomParticipant(db *gorm.DB, query model.ChinchiroRoomParticipantQuery) (participant *model.ChinchiroRoomParticipant, err error) {
	rdb, err := findRdbChinchiroRoomParticipant(db, query)
	if err != nil {
		return nil, err
	}
	if rdb == nil {
		return nil, nil
	}
	return rdb.ToModel(), nil
}

// -----------------------

func findRdbChinchiroRoomParticipants(db *gorm.DB, query model.ChinchiroRoomParticipantsQuery) (_ []ChinchiroRoomParticipant, err error) {
	var rdbRoomParticipants []ChinchiroRoomParticipant
	result := db.Model(&ChinchiroRoomParticipant{})
	if query.Paging != nil {
		result = result.Scopes(Paginate(query.Paging))
	} else {
		result = result.Scopes(Paginate(&model.PagingQuery{
			PageSize:   10000,
			PageNumber: 1,
		}))
	}
	if query.IDs != nil {
		result = result.Where("id in (?)", *query.IDs)
	}
	if query.Name != nil {
		result = result.Where("participant_name like ?", fmt.Sprintf("%%%s%%", *query.Name))
	}
	if query.RoomIDs != nil {
		result = result.Where("room_id in (?)", *query.RoomIDs)
	}
	if query.IsExcludeGone != nil && *query.IsExcludeGone {
		result = result.Where("is_gone = ?", false)
	}
	result = result.Find(&rdbRoomParticipants)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if result.Error != nil {
		return nil, fmt.Errorf("failed to find: %s", result.Error)
	}

	return rdbRoomParticipants, nil
}

func findRdbChinchiroRoomParticipant(db *gorm.DB, query model.ChinchiroRoomParticipantQuery) (_ *ChinchiroRoomParticipant, err error) {
	var rdbRoomParticipant ChinchiroRoomParticipant
	result := db.Model(&ChinchiroRoomParticipant{}).Where("is_gone = ?", false)
	if query.ID != nil {
		result = result.Where("id = ?", *query.ID)
	}
	if query.RoomID != nil {
		result = result.Where("room_id = ?", *query.RoomID)
	}
	if query.PlayerID != nil {
		result = result.Where("player_id = ?", *query.PlayerID)
	}
	result = result.First(&rdbRoomParticipant)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if result.Error != nil {
		return nil, fmt.Errorf("failed to find: %s", result.Error)
	}
	return &rdbRoomParticipant, nil
}

func registerRdbChinchiroRoomParticipant(db *gorm.DB, roomID uint32, participant model.ChinchiroRoomParticipant) (*ChinchiroRoomParticipant, error) {
	p := ChinchiroRoomParticipant{
		RoomID:          roomID,
		PlayerID:        participant.PlayerID,
		ParticipantName: participant.Name,
		IsGone:          false,
	}
	if result := db.Create(&p); result.Error != nil {
		return nil, fmt.Errorf("failed to save: %s", result.Error)
	}
	return &p, nil
}

func updateRdbChinchiroRoomParticipant(db *gorm.DB, participant model.ChinchiroRoomParticipant) error {
	if err := db.Model(&ChinchiroRoomParticipant{}).Where("id = ?", participant.ID).Updates(ChinchiroRoomParticipant{
		ParticipantName: participant.Name,
		IsGone:          participant.IsGone,
	}).Error; err != nil {
		return err
	}
	return nil
}

func deleteRdbChinchiroRoomParticipant(db *gorm.DB, participantID uint32) error {
	if err := db.Model(&ChinchiroRoomParticipant{}).Where("id = ?", participantID).Update("is_gone", true).Error; err != nil {
		return err
	}
	return nil
}
