package model

import (
	"context"
	"wolfort-games/util/array"
)

type ChinchiroRoom struct {
	ID             uint32
	Name           string
	Status         ChinchiroRoomStatus
	RoomMasterIDs  []uint32
	ParticipantIDs []uint32
	GameIDs        []uint32
	Settings       ChinchiroRoomSettings
}

type ChinchiroRoomStatus int

const (
	ChinchiroRoomStatusOpened ChinchiroRoomStatus = iota
	ChinchiroRoomStatusFinished
)

func (rs ChinchiroRoomStatus) String() string {
	switch rs {
	case ChinchiroRoomStatusOpened:
		return "Opened"
	case ChinchiroRoomStatusFinished:
		return "Finished"
	default:
		return ""
	}
}

func ChinchiroRoomStatusValues() []ChinchiroRoomStatus {
	return []ChinchiroRoomStatus{
		ChinchiroRoomStatusOpened,
		ChinchiroRoomStatusFinished,
	}
}

func ChinchiroRoomStatusValueOf(s string) *ChinchiroRoomStatus {
	return array.Find(ChinchiroRoomStatusValues(), func(rs ChinchiroRoomStatus) bool {
		return rs.String() == s
	})
}

type ChinchiroRoomsQuery struct {
	IDs      *[]uint32
	Name     *string
	Statuses *[]ChinchiroRoomStatus
	Paging   *PagingQuery
}

type ChinchiroRoomMaster struct {
	ID       uint32
	PlayerID uint32
}

type ChinchiroRoomMastersQuery struct {
	IDs    *[]uint32
	RoomID *uint32
}

type ChinchiroRoomParticipants struct {
	Count int
	List  []ChinchiroRoomParticipant
}

type ChinchiroRoomParticipantsQuery struct {
	IDs           *[]uint32
	Name          *string
	RoomIDs       *[]uint32
	IsExcludeGone *bool
	Paging        *PagingQuery
}

type ChinchiroRoomParticipant struct {
	ID       uint32
	Name     string
	PlayerID uint32
	RoomID   uint32
	IsGone   bool
}

type ChinchiroRoomSettings struct {
	Password ChinchiroRoomPasswordSettings
}

type ChinchiroRoomPasswordSettings struct {
	Password *string
}

func (s ChinchiroRoomPasswordSettings) HasPassword() bool {
	return s.Password != nil && *s.Password != ""
}

type ChinchiroRoomParticipantQuery struct {
	ID       *uint32
	RoomID   *uint32
	PlayerID *uint32
}

// -----------------------

type ChinchiroRoomRepository interface {
	// room
	FindRooms(query ChinchiroRoomsQuery) (rooms []ChinchiroRoom, err error)
	FindRoom(ID uint32) (room *ChinchiroRoom, err error)
	RegisterRoom(ctx context.Context, room ChinchiroRoom) (saved *ChinchiroRoom, err error)
	UpdateRoom(ctx context.Context, room ChinchiroRoom) (err error)
	UpdateRoomSettings(ctx context.Context, roomID uint32, settings ChinchiroRoomSettings) (err error)
	// room master
	FindRoomMasters(query ChinchiroRoomMastersQuery) (masters []ChinchiroRoomMaster, err error)
	RegisterRoomMaster(ctx context.Context, roomID uint32, master ChinchiroRoomMaster) (saved *ChinchiroRoomMaster, err error)
	DeleteRoomMaster(ctx context.Context, roomMasterID uint32) (err error)
}

type ChinchiroRoomParticipantRepository interface {
	// room participant
	FindRoomParticipants(query ChinchiroRoomParticipantsQuery) (participants ChinchiroRoomParticipants, err error)
	FindRoomParticipant(query ChinchiroRoomParticipantQuery) (participant *ChinchiroRoomParticipant, err error)
	RegisterRoomParticipant(ctx context.Context, roomID uint32, participant ChinchiroRoomParticipant) (saved *ChinchiroRoomParticipant, err error)
	UpdateRoomParticipant(ctx context.Context, participant ChinchiroRoomParticipant) (err error)
	DeleteRoomParticipant(ctx context.Context, participantID uint32) (err error)
}
