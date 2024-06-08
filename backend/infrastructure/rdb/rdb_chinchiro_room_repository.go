package db

import (
	model "chat-role-play/domain/model"
	"chat-role-play/util/array"
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type ChinchiroRoomRepository struct {
	db *DB
}

func NewChinchiroRoomRepository(db *DB) *ChinchiroRoomRepository {
	return &ChinchiroRoomRepository{
		db: db,
	}
}

func (repo *ChinchiroRoomRepository) FindRooms(query model.ChinchiroRoomsQuery) (rooms []model.ChinchiroRoom, err error) {
	return findRooms(repo.db.Connection, query)
}

func (repo *ChinchiroRoomRepository) FindRoom(ID uint32) (room *model.ChinchiroRoom, err error) {
	return findRoom(repo.db.Connection, ID)
}

func (repo *ChinchiroRoomRepository) RegisterRoom(ctx context.Context, room model.ChinchiroRoom) (saved *model.ChinchiroRoom, err error) {
	tx, ok := GetTx(ctx)
	if !ok {
		return nil, fmt.Errorf("failed to get tx from context")
	}
	// room
	r, err := registerRdbRoom(tx, room)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	// room settings
	if err := registerRoomSettings(tx, r.ID, room.Settings); err != nil {
		return nil, err
	}
	return findRoom(tx, r.ID)
}

func (repo *ChinchiroRoomRepository) UpdateRoom(ctx context.Context, room model.ChinchiroRoom) (err error) {
	tx, ok := GetTx(ctx)
	if !ok {
		return fmt.Errorf("failed to get tx from context")
	}
	return updateRoom(tx, room)
}

func (repo *ChinchiroRoomRepository) FindRoomMasters(query model.ChinchiroRoomMastersQuery) (masters []model.ChinchiroRoomMaster, err error) {
	return findRoomMasters(repo.db.Connection, query)
}

func (repo *ChinchiroRoomRepository) RegisterRoomMaster(ctx context.Context, roomID uint32, master model.ChinchiroRoomMaster) (saved *model.ChinchiroRoomMaster, err error) {
	tx, ok := GetTx(ctx)
	if !ok {
		return nil, fmt.Errorf("failed to get tx from context")
	}
	return registerRoomMaster(tx, roomID, master)
}

func (repo *ChinchiroRoomRepository) DeleteRoomMaster(ctx context.Context, roomMasterID uint32) (err error) {
	tx, ok := GetTx(ctx)
	if !ok {
		return fmt.Errorf("failed to get tx from context")
	}
	return deleteRoomMaster(tx, roomMasterID)
}

// -----------------------

func findRooms(db *gorm.DB, query model.ChinchiroRoomsQuery) (rooms []model.ChinchiroRoom, err error) {
	rdbRooms, err := findRdbChinchiroRooms(db, query)
	if err != nil {
		return nil, err
	}
	if rdbRooms == nil {
		return nil, nil
	}
	ids := array.Map(rdbRooms, func(r ChinchiroRoom) uint32 {
		return r.ID
	})
	if len(ids) == 0 {
		return nil, nil
	}
	pts, err := findRdbChinchiroRoomParticipants(db, model.ChinchiroRoomParticipantsQuery{
		RoomIDs: &ids,
	})
	settings, err := findRdbChinchiroRoomSettings(db, chinchiroRoomSettingsQuery{RoomIDs: &ids})
	if err != nil {
		return nil, err
	}

	return array.Map(rdbRooms, func(r ChinchiroRoom) model.ChinchiroRoom {
		ptsCount := array.Count(pts, func(p ChinchiroRoomParticipant) bool {
			return p.RoomID == r.ID
		})
		rdbSettings := array.Filter(settings, func(s ChinchiroRoomSetting) bool {
			return s.RoomID == r.ID
		})
		roomSettings := ToChinchiroRoomSettingsModel(rdbSettings)
		return *r.ToSimpleModel(ptsCount, *roomSettings)
	}), nil
}

func findRoom(db *gorm.DB, ID uint32) (room *model.ChinchiroRoom, err error) {
	rdbRoom, err := findRdbChinchiroRoom(db, ID)
	if err != nil {
		return nil, err
	}
	if rdbRoom == nil {
		return nil, nil
	}
	roomMasters, err := findRdbChinchiroRoomMasterPlayers(db, ID)
	if err != nil {
		return nil, err
	}
	participants, err := findRdbChinchiroRoomParticipants(db, model.ChinchiroRoomParticipantsQuery{
		RoomIDs: &[]uint32{ID},
	})
	games, err := findRdbChinchiroGames(db, model.ChinchiroGamesQuery{
		RoomID: &ID,
	})
	settings, err := findRdbChinchiroRoomSettings(db, chinchiroRoomSettingsQuery{RoomID: &ID})
	if err != nil {
		return nil, err
	}
	roomSettings := ToChinchiroRoomSettingsModel(settings)
	return rdbRoom.ToModel(roomMasters, participants, games, *roomSettings), nil
}

// -----------------------

func findRdbChinchiroRooms(db *gorm.DB, query model.ChinchiroRoomsQuery) ([]ChinchiroRoom, error) {
	var rdbRooms []ChinchiroRoom
	result := db.Model(&ChinchiroRoom{})
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
	if query.Name != nil {
		result = result.Where("room_name like ?", fmt.Sprintf("%%%s%%", *query.Name))
	}
	if query.Statuses != nil {
		statuses := array.Map(*query.Statuses, func(s model.ChinchiroRoomStatus) string {
			return s.String()
		})
		result = result.Where("room_status_code in (?)", statuses)
	}
	result = result.Find(&rdbRooms)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if result.Error != nil {
		return nil, fmt.Errorf("failed to find: %s \n", result.Error)
	}

	return rdbRooms, nil
}

func findRdbChinchiroRoom(db *gorm.DB, ID uint32) (*ChinchiroRoom, error) {
	var rdb ChinchiroRoom
	result := db.Model(&Game{}).First(&rdb, ID)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if result.Error != nil {
		return nil, fmt.Errorf("failed to find: %s \n", result.Error)
	}
	return &rdb, nil
}

func findRdbChinchiroRoomMasterPlayers(db *gorm.DB, roomID uint32) (_ []ChinchiroRoomMasterPlayer, err error) {
	var rdbs []ChinchiroRoomMasterPlayer
	result := db.Model(&ChinchiroRoomMasterPlayer{}).Where("room_id = ?", roomID).Find(&rdbs)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if result.Error != nil {
		return nil, fmt.Errorf("failed to find: %s \n", result.Error)
	}
	return rdbs, nil

}

func findRdbChinchiroRoomSettings(db *gorm.DB, query chinchiroRoomSettingsQuery) (_ []ChinchiroRoomSetting, err error) {
	var rdbs []ChinchiroRoomSetting
	result := db.Model(&ChinchiroRoomSetting{})
	if query.RoomIDs != nil {
		result = result.Where("room_id in (?)", *query.RoomIDs)
	}
	if query.RoomID != nil {
		result = result.Where("room_id = ?", *query.RoomID)
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

type chinchiroRoomSettingsQuery struct {
	RoomIDs *[]uint32
	RoomID  *uint32
}

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
		return nil, fmt.Errorf("failed to find: %s \n", result.Error)
	}

	return rdbRoomParticipants, nil
}

func registerRdbRoom(db *gorm.DB, room model.ChinchiroRoom) (*ChinchiroRoom, error) {
	r := ChinchiroRoom{
		RoomName:       room.Name,
		RoomStatusCode: room.Status.String(),
	}
	if result := db.Create(&r); result.Error != nil {
		return nil, fmt.Errorf("failed to save: %s \n", result.Error)
	}
	return &r, nil
}
