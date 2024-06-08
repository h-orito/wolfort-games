package db

import (
	model "chat-role-play/domain/model"
	"chat-role-play/util/array"
	"strconv"
	"time"
)

type ChinchiroRoom struct {
	ID             uint32
	RoomName       string
	RoomStatusCode string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func (r ChinchiroRoom) ToModel(
	roomMasterPlayers []ChinchiroRoomMasterPlayer,
	participants []ChinchiroRoomParticipant,
	games []ChinchiroGame,
	settings model.ChinchiroRoomSettings,
) *model.ChinchiroRoom {
	return &model.ChinchiroRoom{
		ID:             r.ID,
		Name:           r.RoomName,
		Status:         *model.ChinchiroRoomStatusValueOf(r.RoomStatusCode),
		RoomMasterIDs:  array.Map(roomMasterPlayers, func(m ChinchiroRoomMasterPlayer) uint32 { return m.ID }),
		ParticipantIDs: array.Map(participants, func(p ChinchiroRoomParticipant) uint32 { return p.ID }),
		GameIDs:        array.Map(games, func(g ChinchiroGame) uint32 { return g.ID }),
		Settings:       settings,
	}
}

func (r ChinchiroRoom) ToSimpleModel(
	ptsCount int,
	settings model.ChinchiroRoomSettings,
) *model.ChinchiroRoom {
	return &model.ChinchiroRoom{
		ID:             r.ID,
		Name:           r.RoomName,
		Status:         *model.ChinchiroRoomStatusValueOf(r.RoomStatusCode),
		RoomMasterIDs:  []uint32{},
		ParticipantIDs: make([]uint32, ptsCount),
		GameIDs:        []uint32{},
		Settings:       settings,
	}
}

type ChinchiroRoomMasterPlayer struct {
	ID        uint32
	RoomID    uint32
	PlayerID  uint32
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (g ChinchiroRoomMasterPlayer) ToModel() *model.ChinchiroRoomMaster {
	return &model.ChinchiroRoomMaster{
		ID:       g.ID,
		PlayerID: g.PlayerID,
	}
}

type ChinchiroRoomParticipant struct {
	ID              uint32
	RoomID          uint32
	PlayerID        uint32
	ParticipantName string
	IsGone          bool
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func (p ChinchiroRoomParticipant) ToModel() *model.ChinchiroRoomParticipant {
	return &model.ChinchiroRoomParticipant{
		ID:       p.ID,
		Name:     p.ParticipantName,
		PlayerID: p.PlayerID,
		IsGone:   p.IsGone,
	}
}

type ChinchiroRoomSetting struct {
	ID               uint32
	RoomID           uint32
	RoomSettingKey   string
	RoomSettingValue string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type ChinchiroRoomSettingKey int

const (
	ChinchiroRoomSettingKeyPassword ChinchiroRoomSettingKey = iota
)

func (k ChinchiroRoomSettingKey) String() string {
	switch k {
	case ChinchiroRoomSettingKeyPassword:
		return "Password"
	default:
		return ""
	}
}

func ChinchiroRoomSettingKeyValues() []ChinchiroRoomSettingKey {
	return []ChinchiroRoomSettingKey{
		ChinchiroRoomSettingKeyPassword,
	}
}

func ChinchiroRoomSettingKeyValueOf(s string) *ChinchiroRoomSettingKey {
	return array.Find(ChinchiroRoomSettingKeyValues(), func(k ChinchiroRoomSettingKey) bool {
		return k.String() == s
	})
}

func ToChinchiroRoomSettingsModel(
	settings []ChinchiroRoomSetting,
) *model.ChinchiroRoomSettings {
	password := roomSettingsToString(settings, ChinchiroRoomSettingKeyPassword, nil)
	return &model.ChinchiroRoomSettings{
		Password: model.ChinchiroRoomPasswordSettings{
			Password: password,
		},
	}
}

func roomSettingsToBool(settings []ChinchiroRoomSetting, Key ChinchiroRoomSettingKey, _default bool) bool {
	for _, s := range settings {
		if s.RoomSettingKey == Key.String() {
			return s.RoomSettingValue == "true"
		}
	}
	return _default
}

func roomSettingsToUint32(settings []ChinchiroRoomSetting, Key ChinchiroRoomSettingKey, _default uint32) uint32 {
	for _, s := range settings {
		if s.RoomSettingKey == Key.String() {
			i, err := strconv.Atoi(s.RoomSettingValue)
			if err != nil {
				return _default
			}
			return uint32(i)
		}
	}
	return _default
}

func roomSettingsToString(settings []ChinchiroRoomSetting, key ChinchiroRoomSettingKey, _default *string) *string {
	for _, s := range settings {
		if s.RoomSettingKey == key.String() {
			return &s.RoomSettingValue
		}
	}
	return _default
}

func roomSettingsToTime(settings []ChinchiroRoomSetting, Key ChinchiroRoomSettingKey, _default time.Time) time.Time {
	for _, s := range settings {
		if s.RoomSettingKey == Key.String() {
			t, err := time.Parse(time.RFC3339, s.RoomSettingValue)
			if err != nil {
				return _default
			}
			return t
		}
	}
	return _default
}
