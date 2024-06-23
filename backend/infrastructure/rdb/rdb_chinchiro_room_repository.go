package db

import (
	"context"
	"errors"
	"fmt"
	model "wolfort-games/domain/model"
	"wolfort-games/util/array"

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
	return findChinchiroRooms(repo.db.Connection, query)
}

func (repo *ChinchiroRoomRepository) FindRoom(ID uint32) (room *model.ChinchiroRoom, err error) {
	return findChinchiroRoom(repo.db.Connection, ID)
}

func (repo *ChinchiroRoomRepository) RegisterRoom(ctx context.Context, room model.ChinchiroRoom) (saved *model.ChinchiroRoom, err error) {
	tx, ok := GetTx(ctx)
	if !ok {
		return nil, fmt.Errorf("failed to get tx from context")
	}
	// room
	r, err := registerRdbChinchiroRoom(tx, room)
	if err != nil {
		return nil, err
	}
	// room settings
	if err := registerRdbChinchiroRoomSettings(tx, r.ID, room.Settings); err != nil {
		return nil, err
	}
	return findChinchiroRoom(tx, r.ID)
}

func (repo *ChinchiroRoomRepository) UpdateRoom(ctx context.Context, room model.ChinchiroRoom) (err error) {
	tx, ok := GetTx(ctx)
	if !ok {
		return fmt.Errorf("failed to get tx from context")
	}
	return updateRdbChinchiroRoom(tx, room)
}

func (repo *ChinchiroRoomRepository) UpdateRoomSettings(ctx context.Context, roomID uint32, settings model.ChinchiroRoomSettings) (err error) {
	tx, ok := GetTx(ctx)
	if !ok {
		return fmt.Errorf("failed to get tx from context")
	}
	if err := updateRdbChinchiroRoomSettings(tx, roomID, settings); err != nil {
		return err
	}
	return nil

}

func (repo *ChinchiroRoomRepository) FindRoomMasters(query model.ChinchiroRoomMastersQuery) (masters []model.ChinchiroRoomMaster, err error) {
	return findChinchiroRoomMasters(repo.db.Connection, query)
}

func (repo *ChinchiroRoomRepository) RegisterRoomMaster(ctx context.Context, roomID uint32, master model.ChinchiroRoomMaster) (saved *model.ChinchiroRoomMaster, err error) {
	tx, ok := GetTx(ctx)
	if !ok {
		return nil, fmt.Errorf("failed to get tx from context")
	}
	rm, err := registerRdbChinchiroRoomMaster(tx, roomID, master)
	if err != nil {
		return nil, err
	}
	return rm.ToModel(), nil
}

func (repo *ChinchiroRoomRepository) DeleteRoomMaster(ctx context.Context, roomMasterID uint32) (err error) {
	tx, ok := GetTx(ctx)
	if !ok {
		return fmt.Errorf("failed to get tx from context")
	}
	return deleteRdbChinchiroRoomMaster(tx, roomMasterID)
}

// -----------------------

func findChinchiroRooms(db *gorm.DB, query model.ChinchiroRoomsQuery) (rooms []model.ChinchiroRoom, err error) {
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
	if err != nil {
		return nil, err
	}
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

func findChinchiroRoom(db *gorm.DB, ID uint32) (room *model.ChinchiroRoom, err error) {
	rdbRoom, err := findRdbChinchiroRoom(db, ID)
	if err != nil {
		return nil, err
	}
	if rdbRoom == nil {
		return nil, nil
	}
	roomMasters, err := findRdbChinchiroRoomMasterPlayers(db, model.ChinchiroRoomMastersQuery{
		RoomID: &ID,
	})
	if err != nil {
		return nil, err
	}
	participants, err := findRdbChinchiroRoomParticipants(db, model.ChinchiroRoomParticipantsQuery{
		RoomIDs: &[]uint32{ID},
	})
	if err != nil {
		return nil, err
	}
	games, err := findRdbChinchiroGames(db, model.ChinchiroGamesQuery{
		RoomID: &ID,
	})
	if err != nil {
		return nil, err
	}
	settings, err := findRdbChinchiroRoomSettings(db, chinchiroRoomSettingsQuery{RoomID: &ID})
	if err != nil {
		return nil, err
	}
	roomSettings := ToChinchiroRoomSettingsModel(settings)
	return rdbRoom.ToModel(roomMasters, participants, games, *roomSettings), nil
}

func findChinchiroRoomMasters(db *gorm.DB, query model.ChinchiroRoomMastersQuery) (masters []model.ChinchiroRoomMaster, err error) {
	rdbMasters, err := findRdbChinchiroRoomMasterPlayers(db, query)
	if err != nil {
		return nil, err
	}
	if rdbMasters == nil {
		return nil, nil
	}
	return array.Map(rdbMasters, func(r ChinchiroRoomMasterPlayer) model.ChinchiroRoomMaster {
		return *r.ToModel()
	}), nil
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
		return nil, fmt.Errorf("failed to find: %s", result.Error)
	}

	return rdbRooms, nil
}

func findRdbChinchiroRoom(db *gorm.DB, ID uint32) (*ChinchiroRoom, error) {
	var rdb ChinchiroRoom
	result := db.Model(&ChinchiroRoom{}).First(&rdb, ID)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if result.Error != nil {
		return nil, fmt.Errorf("failed to find: %s", result.Error)
	}
	return &rdb, nil
}

func findRdbChinchiroRoomMasterPlayers(db *gorm.DB, query model.ChinchiroRoomMastersQuery) (_ []ChinchiroRoomMasterPlayer, err error) {
	var rdbs []ChinchiroRoomMasterPlayer
	result := db.Model(&ChinchiroRoomMasterPlayer{})

	if query.IDs != nil {
		result = result.Where("id in (?)", *query.IDs)
	}
	if query.RoomID != nil {
		result = result.Where("room_id = ?", *query.RoomID).Find(&rdbs)
	}
	result = result.Find(&rdbs)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if result.Error != nil {
		return nil, fmt.Errorf("failed to find: %s", result.Error)
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
		return nil, fmt.Errorf("failed to find: %s", result.Error)
	}
	return rdbs, nil
}

type chinchiroRoomSettingsQuery struct {
	RoomIDs *[]uint32
	RoomID  *uint32
}

func registerRdbChinchiroRoom(db *gorm.DB, room model.ChinchiroRoom) (*ChinchiroRoom, error) {
	r := ChinchiroRoom{
		RoomName:       room.Name,
		RoomStatusCode: room.Status.String(),
	}
	if result := db.Create(&r); result.Error != nil {
		return nil, fmt.Errorf("failed to save: %s", result.Error)
	}
	return &r, nil
}

func updateRdbChinchiroRoom(db *gorm.DB, room model.ChinchiroRoom) error {
	if err := db.Model(&ChinchiroRoom{}).Where("id = ?", room.ID).Updates(ChinchiroRoom{
		RoomName:       room.Name,
		RoomStatusCode: room.Status.String(),
	}).Error; err != nil {
		return err
	}
	return nil
}

func registerRdbChinchiroRoomMaster(db *gorm.DB, roomID uint32, master model.ChinchiroRoomMaster) (*ChinchiroRoomMasterPlayer, error) {
	r := ChinchiroRoomMasterPlayer{
		RoomID:   roomID,
		PlayerID: master.PlayerID,
	}
	if result := db.Create(&r); result.Error != nil {
		return nil, fmt.Errorf("failed to save: %s", result.Error)
	}
	return &r, nil
}

func deleteRdbChinchiroRoomMaster(db *gorm.DB, roomMasterID uint32) error {
	if err := db.Where("id = ?", roomMasterID).Delete(&ChinchiroRoomMasterPlayer{}).Error; err != nil {
		return err
	}
	return nil
}

func registerRdbChinchiroRoomSettings(db *gorm.DB, roomID uint32, settings model.ChinchiroRoomSettings) error {
	if err := registerRdbChinchiroRoomSetting(db, roomID, ChinchiroRoomSettingKeyPassword, orEmpty(settings.Password.Password)); err != nil {
		return err
	}
	return nil
}

func updateRdbChinchiroRoomSettings(db *gorm.DB, roomID uint32, settings model.ChinchiroRoomSettings) error {
	if err := updateRdbChinchiroRoomSetting(db, roomID, ChinchiroRoomSettingKeyPassword, orEmpty(settings.Password.Password)); err != nil {
		return err
	}
	return nil
}

func registerRdbChinchiroRoomSetting(db *gorm.DB, roomID uint32, key ChinchiroRoomSettingKey, value string) (err error) {
	if err := db.Create(&ChinchiroRoomSetting{
		RoomID:           roomID,
		RoomSettingKey:   key.String(),
		RoomSettingValue: value,
	}).Error; err != nil {
		return err
	}
	return nil
}

func updateRdbChinchiroRoomSetting(db *gorm.DB, roomID uint32, key ChinchiroRoomSettingKey, value string) (err error) {
	if err := db.Model(&ChinchiroRoomSetting{}).Where("room_id = ? and room_setting_key = ?", roomID, key.String()).Updates(&ChinchiroRoomSetting{
		RoomSettingValue: value,
	}).Error; err != nil {
		return err
	}
	return nil
}

// -----------------------

func orEmpty(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
