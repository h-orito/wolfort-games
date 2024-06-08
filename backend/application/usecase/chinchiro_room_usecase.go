package usecase

import (
	"chat-role-play/application/app_service"
	"chat-role-play/domain/model"
	"context"
)

type ChinchiroRoomUsecase interface {
	// room
	FindChinchiroRooms(query model.ChinchiroRoomsQuery) ([]model.ChinchiroRoom, error)
	FindChinchiroRoom(ID uint32) (*model.ChinchiroRoom, error)
	RegisterChinchiroRoom(ctx context.Context, user model.User, room model.ChinchiroRoom) (*model.ChinchiroRoom, error)
	UpdateChinchiroRoomStatus(ctx context.Context, user model.User, roomID uint32, status model.ChinchiroRoomStatus) error
	UpdateChinchiroRoomSettings(ctx context.Context, user model.User, roomID uint32, settings model.ChinchiroRoomSettings) error
	// room master
	FindChinchiroRoomMasters(query model.ChinchiroRoomMastersQuery) ([]model.ChinchiroRoomMaster, error)
	RegisterChinchiroRoomMaster(ctx context.Context, user model.User, roomID uint32, playerID uint32) (model.ChinchiroRoomMaster, error)
	DeleteChinchiroRoomMaster(ctx context.Context, user model.User, roomID uint32, masterID uint32) error
	// room participant
	FindChinchiroRoomParticipants(query model.ChinchiroRoomParticipantsQuery) (model.ChinchiroRoomParticipants, error)
	FindChinchiroRoomParticipant(ID uint32) (*model.ChinchiroRoomParticipant, error)
	FindMyChinchiroRoomParticipant(user model.User, roomID uint32) (*model.ChinchiroRoomParticipant, error)
	ParticipateChinchiroRoom(ctx context.Context, user model.User, roomID uint32, participant model.ChinchiroRoomParticipant, password *string) (*model.ChinchiroRoomParticipant, error)
	UpdateChinchiroRoomParticipant(ctx context.Context, user model.User, participantID uint32, participant model.ChinchiroRoomParticipant) error
	LeaveChinchiroRoom(ctx context.Context, user model.User, roomID uint32) error
	DeleteChinchiroRoomParticipant(ctx context.Context, user model.User, participantID uint32) error
}

type chinchiroRoomUsecase struct {
	chinchiroRoomService app_service.ChinchiroRoomService
}

// DeleteChinchiroRoomParticipant implements ChinchiroRoomUsecase.
func (*chinchiroRoomUsecase) DeleteChinchiroRoomParticipant(ctx context.Context, user model.User, participantID uint32) error {
	panic("unimplemented")
}

// UpdateChinchiroRoomParticipant implements ChinchiroRoomUsecase.
func (*chinchiroRoomUsecase) UpdateChinchiroRoomParticipant(ctx context.Context, user model.User, participantID uint32, participant model.ChinchiroRoomParticipant) error {
	panic("unimplemented")
}

// UpdateChinchiroRoomSettings implements ChinchiroRoomUsecase.
func (*chinchiroRoomUsecase) UpdateChinchiroRoomSettings(ctx context.Context, user model.User, roomID uint32, settings model.ChinchiroRoomSettings) error {
	panic("unimplemented")
}

// UpdateChinchiroRoomStatus implements ChinchiroRoomUsecase.
func (*chinchiroRoomUsecase) UpdateChinchiroRoomStatus(ctx context.Context, user model.User, roomID uint32, status model.ChinchiroRoomStatus) error {
	panic("unimplemented")
}

// DeleteChinchiroRoomMaster implements ChinchiroRoomUsecase.
func (*chinchiroRoomUsecase) DeleteChinchiroRoomMaster(ctx context.Context, user model.User, roomID uint32, masterID uint32) error {
	panic("unimplemented")
}

// RegisterChinchiroRoomMaster implements ChinchiroRoomUsecase.
func (*chinchiroRoomUsecase) RegisterChinchiroRoomMaster(ctx context.Context, user model.User, roomID uint32, playerID uint32) (model.ChinchiroRoomMaster, error) {
	panic("unimplemented")
}

// FindChinchiroRoomMasters implements ChinchiroRoomUsecase.
func (*chinchiroRoomUsecase) FindChinchiroRoomMasters(query model.ChinchiroRoomMastersQuery) ([]model.ChinchiroRoomMaster, error) {
	panic("unimplemented")
}

// FindChinchiroRoom implements ChinchiroRoomUsecase.
func (*chinchiroRoomUsecase) FindChinchiroRoom(ID uint32) (*model.ChinchiroRoom, error) {
	panic("unimplemented")
}

// FindChinchiroRoomParticipant implements ChinchiroRoomUsecase.
func (*chinchiroRoomUsecase) FindChinchiroRoomParticipant(ID uint32) (*model.ChinchiroRoomParticipant, error) {
	panic("unimplemented")
}

// FindChinchiroRoomParticipants implements ChinchiroRoomUsecase.
func (*chinchiroRoomUsecase) FindChinchiroRoomParticipants(query model.ChinchiroRoomParticipantsQuery) (model.ChinchiroRoomParticipants, error) {
	panic("unimplemented")
}

// FindChinchiroRooms implements ChinchiroRoomUsecase.
func (*chinchiroRoomUsecase) FindChinchiroRooms(query model.ChinchiroRoomsQuery) ([]model.ChinchiroRoom, error) {
	panic("unimplemented")
}

// FindMyChinchiroRoomParticipant implements ChinchiroRoomUsecase.
func (*chinchiroRoomUsecase) FindMyChinchiroRoomParticipant(user model.User, roomID uint32) (*model.ChinchiroRoomParticipant, error) {
	panic("unimplemented")
}

// LeaveChinchiroRoom implements ChinchiroRoomUsecase.
func (*chinchiroRoomUsecase) LeaveChinchiroRoom(ctx context.Context, user model.User, roomID uint32) error {
	panic("unimplemented")
}

// ParticipateChinchiroRoom implements ChinchiroRoomUsecase.
func (*chinchiroRoomUsecase) ParticipateChinchiroRoom(ctx context.Context, user model.User, roomID uint32, participant model.ChinchiroRoomParticipant, password *string) (*model.ChinchiroRoomParticipant, error) {
	panic("unimplemented")
}

// RegisterChinchiroRoom implements ChinchiroRoomUsecase.
func (*chinchiroRoomUsecase) RegisterChinchiroRoom(ctx context.Context, user model.User, room model.ChinchiroRoom) (*model.ChinchiroRoom, error) {
	panic("unimplemented")
}

func NewChinchiroRoomUsecase(chinchiroRoomService app_service.ChinchiroRoomService) ChinchiroRoomUsecase {
	return &chinchiroRoomUsecase{
		chinchiroRoomService: chinchiroRoomService,
	}
}
