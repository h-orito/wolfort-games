package app_service

import (
	"chat-role-play/domain/model"
	"context"
)

type ChinchiroRoomService interface {
	// room
	FindChinchiroRooms(query model.ChinchiroRoomsQuery) ([]model.ChinchiroRoom, error)
	FindChinchiroRoom(ID uint32) (*model.ChinchiroRoom, error)
	RegisterChinchiroRoom(ctx context.Context, room model.ChinchiroRoom) (*model.ChinchiroRoom, error)
	UpdateChinchiroRoomStatus(ctx context.Context, roomID uint32, status model.ChinchiroRoomStatus) error
	// room master
	FindChinchiroRoomMasters(query model.ChinchiroRoomMastersQuery) ([]model.ChinchiroRoomMaster, error)
	RegisterChinchiroRoomMaster(ctx context.Context, roomID uint32, playerID uint32) error
	DeleteChinchiroRoomMaster(ctx context.Context, masterID uint32) error
	// room participant
	FindChinchiroRoomParticipants(query model.ChinchiroRoomParticipantsQuery) (model.ChinchiroRoomParticipants, error)
	FindChinchiroRoomParticipant(ID uint32) (*model.ChinchiroRoomParticipant, error)
	FindMyChinchiroRoomParticipant(user model.User, roomID uint32) (*model.ChinchiroRoomParticipant, error)
	RegisterChinchiroRoomParticipant(ctx context.Context, roomID uint32, participant model.ChinchiroRoomParticipant) (*model.ChinchiroRoomParticipant, error)
	UpdateChinchiroRoomParticipant(ctx context.Context, participantID uint32, participant model.ChinchiroRoomParticipant) (*model.ChinchiroRoomParticipant, error)
	DeleteChinchiroRoomParticipant(ctx context.Context, participantID uint32) error
}

type chinchiroRoomService struct {
	chinchiroRoomRepository model.ChinchiroRoomRepository
}

// FindChinchiroRoomMasters implements ChinchiroRoomService.
func (*chinchiroRoomService) FindChinchiroRoomMasters(query model.ChinchiroRoomMastersQuery) ([]model.ChinchiroRoomMaster, error) {
	panic("unimplemented")
}

// FindMyChinchiroRoomParticipant implements ChinchiroRoomService.
func (*chinchiroRoomService) FindMyChinchiroRoomParticipant(user model.User, roomID uint32) (*model.ChinchiroRoomParticipant, error) {
	panic("unimplemented")
}

// DeleteChinchiroRoomMaster implements ChinchiroRoomService.
func (*chinchiroRoomService) DeleteChinchiroRoomMaster(ctx context.Context, masterID uint32) error {
	panic("unimplemented")
}

// DeleteChinchiroRoomParticipant implements ChinchiroRoomService.
func (*chinchiroRoomService) DeleteChinchiroRoomParticipant(ctx context.Context, participantID uint32) error {
	panic("unimplemented")
}

// FindChinchiroRoom implements ChinchiroRoomService.
func (*chinchiroRoomService) FindChinchiroRoom(ID uint32) (*model.ChinchiroRoom, error) {
	panic("unimplemented")
}

// FindChinchiroRoomParticipant implements ChinchiroRoomService.
func (*chinchiroRoomService) FindChinchiroRoomParticipant(ID uint32) (*model.ChinchiroRoomParticipant, error) {
	panic("unimplemented")
}

// FindChinchiroRoomParticipants implements ChinchiroRoomService.
func (*chinchiroRoomService) FindChinchiroRoomParticipants(query model.ChinchiroRoomParticipantsQuery) (model.ChinchiroRoomParticipants, error) {
	panic("unimplemented")
}

// FindChinchiroRooms implements ChinchiroRoomService.
func (*chinchiroRoomService) FindChinchiroRooms(query model.ChinchiroRoomsQuery) ([]model.ChinchiroRoom, error) {
	panic("unimplemented")
}

// RegisterChinchiroRoom implements ChinchiroRoomService.
func (*chinchiroRoomService) RegisterChinchiroRoom(ctx context.Context, room model.ChinchiroRoom) (*model.ChinchiroRoom, error) {
	panic("unimplemented")
}

// RegisterChinchiroRoomMaster implements ChinchiroRoomService.
func (*chinchiroRoomService) RegisterChinchiroRoomMaster(ctx context.Context, roomID uint32, masterID uint32) error {
	panic("unimplemented")
}

// RegisterChinchiroRoomParticipant implements ChinchiroRoomService.
func (*chinchiroRoomService) RegisterChinchiroRoomParticipant(ctx context.Context, roomID uint32, participant model.ChinchiroRoomParticipant) (*model.ChinchiroRoomParticipant, error) {
	panic("unimplemented")
}

// UpdateChinchiroRoomParticipant implements ChinchiroRoomService.
func (*chinchiroRoomService) UpdateChinchiroRoomParticipant(ctx context.Context, participantID uint32, participant model.ChinchiroRoomParticipant) (*model.ChinchiroRoomParticipant, error) {
	panic("unimplemented")
}

// UpdateChinchiroRoomStatus implements ChinchiroRoomService.
func (*chinchiroRoomService) UpdateChinchiroRoomStatus(ctx context.Context, roomID uint32, status model.ChinchiroRoomStatus) error {
	panic("unimplemented")
}

func NewChinchiroRoomService(
	chinchiroRoomRepository model.ChinchiroRoomRepository,
) ChinchiroRoomService {
	return &chinchiroRoomService{
		chinchiroRoomRepository: chinchiroRoomRepository,
	}
}
