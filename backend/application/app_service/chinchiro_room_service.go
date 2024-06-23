package app_service

import (
	"context"
	"wolfort-games/domain/model"
)

type ChinchiroRoomService interface {
	// room
	FindChinchiroRooms(query model.ChinchiroRoomsQuery) ([]model.ChinchiroRoom, error)
	FindChinchiroRoom(ID uint32) (*model.ChinchiroRoom, error)
	RegisterChinchiroRoom(ctx context.Context, room model.ChinchiroRoom) (*model.ChinchiroRoom, error)
	UpdateChinchiroRoomStatus(ctx context.Context, roomID uint32, status model.ChinchiroRoomStatus) error
	UpdateChinchiroRoomSettings(ctx context.Context, roomID uint32, settings model.ChinchiroRoomSettings) error
	// room master
	FindChinchiroRoomMasters(query model.ChinchiroRoomMastersQuery) ([]model.ChinchiroRoomMaster, error)
	RegisterChinchiroRoomMaster(ctx context.Context, roomID uint32, playerID uint32) (*model.ChinchiroRoomMaster, error)
	DeleteChinchiroRoomMaster(ctx context.Context, masterID uint32) error
	// room participant
	FindChinchiroRoomParticipants(query model.ChinchiroRoomParticipantsQuery) (model.ChinchiroRoomParticipants, error)
	FindChinchiroRoomParticipant(query model.ChinchiroRoomParticipantQuery) (*model.ChinchiroRoomParticipant, error)
	RegisterChinchiroRoomParticipant(ctx context.Context, roomID uint32, participant model.ChinchiroRoomParticipant) (*model.ChinchiroRoomParticipant, error)
	UpdateChinchiroRoomParticipant(ctx context.Context, participantID uint32, participant model.ChinchiroRoomParticipant) error
	DeleteChinchiroRoomParticipant(ctx context.Context, participantID uint32) error
}

type chinchiroRoomService struct {
	chinchiroRoomRepository            model.ChinchiroRoomRepository
	chinchiroRoomParticipantRepository model.ChinchiroRoomParticipantRepository
}

func (s *chinchiroRoomService) FindChinchiroRooms(query model.ChinchiroRoomsQuery) ([]model.ChinchiroRoom, error) {
	return s.chinchiroRoomRepository.FindRooms(query)
}

func (s *chinchiroRoomService) FindChinchiroRoom(ID uint32) (*model.ChinchiroRoom, error) {
	return s.chinchiroRoomRepository.FindRoom(ID)
}

func (s *chinchiroRoomService) RegisterChinchiroRoom(ctx context.Context, room model.ChinchiroRoom) (*model.ChinchiroRoom, error) {
	return s.chinchiroRoomRepository.RegisterRoom(ctx, room)
}

func (s *chinchiroRoomService) UpdateChinchiroRoomStatus(ctx context.Context, roomID uint32, status model.ChinchiroRoomStatus) error {
	return s.chinchiroRoomRepository.UpdateRoom(ctx, model.ChinchiroRoom{
		ID:     roomID,
		Status: status,
	})
}

func (s *chinchiroRoomService) UpdateChinchiroRoomSettings(ctx context.Context, roomID uint32, settings model.ChinchiroRoomSettings) error {
	return s.chinchiroRoomRepository.UpdateRoomSettings(ctx, roomID, settings)
}

func (s *chinchiroRoomService) FindChinchiroRoomMasters(query model.ChinchiroRoomMastersQuery) ([]model.ChinchiroRoomMaster, error) {
	return s.chinchiroRoomRepository.FindRoomMasters(query)
}

func (s *chinchiroRoomService) RegisterChinchiroRoomMaster(
	ctx context.Context,
	roomID uint32,
	playerID uint32,
) (*model.ChinchiroRoomMaster, error) {
	return s.chinchiroRoomRepository.RegisterRoomMaster(ctx, roomID, model.ChinchiroRoomMaster{
		PlayerID: playerID,
	})
}

func (s *chinchiroRoomService) DeleteChinchiroRoomMaster(ctx context.Context, masterID uint32) error {
	return s.chinchiroRoomRepository.DeleteRoomMaster(ctx, masterID)
}

func (s *chinchiroRoomService) FindChinchiroRoomParticipants(query model.ChinchiroRoomParticipantsQuery) (model.ChinchiroRoomParticipants, error) {
	return s.chinchiroRoomParticipantRepository.FindRoomParticipants(query)
}

func (s *chinchiroRoomService) FindChinchiroRoomParticipant(query model.ChinchiroRoomParticipantQuery) (*model.ChinchiroRoomParticipant, error) {
	return s.chinchiroRoomParticipantRepository.FindRoomParticipant(query)
}

func (s *chinchiroRoomService) RegisterChinchiroRoomParticipant(ctx context.Context, roomID uint32, participant model.ChinchiroRoomParticipant) (*model.ChinchiroRoomParticipant, error) {
	return s.chinchiroRoomParticipantRepository.RegisterRoomParticipant(ctx, roomID, participant)
}

func (s *chinchiroRoomService) UpdateChinchiroRoomParticipant(ctx context.Context, participantID uint32, participant model.ChinchiroRoomParticipant) error {
	return s.chinchiroRoomParticipantRepository.UpdateRoomParticipant(ctx, participant)
}

func (s *chinchiroRoomService) DeleteChinchiroRoomParticipant(ctx context.Context, participantID uint32) error {
	participant, err := s.chinchiroRoomParticipantRepository.FindRoomParticipant(model.ChinchiroRoomParticipantQuery{
		ID: &participantID,
	})
	if err != nil {
		return err
	}
	// RoomMasterから削除
	rms, err := s.chinchiroRoomRepository.FindRoomMasters(model.ChinchiroRoomMastersQuery{
		RoomID: &participant.RoomID,
	})
	if err != nil {
		return err
	}
	for _, rm := range rms {
		if rm.PlayerID == participant.PlayerID {
			err = s.chinchiroRoomRepository.DeleteRoomMaster(ctx, rm.ID)
			if err != nil {
				return err
			}
		}
	}
	err = s.chinchiroRoomParticipantRepository.DeleteRoomParticipant(ctx, participantID)
	if err != nil {
		return err
	}
	// RoomMasterがいなくなったら部屋を削除
	if len(rms) == 1 && rms[0].PlayerID == participant.PlayerID {
		return s.chinchiroRoomRepository.UpdateRoom(ctx, model.ChinchiroRoom{
			ID:     participant.RoomID,
			Status: model.ChinchiroRoomStatusFinished,
		})
	}

	// 全員退出したら部屋を削除
	isExcludeGone := true
	participants, err := s.chinchiroRoomParticipantRepository.FindRoomParticipants(model.ChinchiroRoomParticipantsQuery{
		RoomIDs:       &[]uint32{participant.RoomID},
		IsExcludeGone: &isExcludeGone,
	})
	if err != nil {
		return err
	}
	if len(participants.List) == 0 {
		return s.chinchiroRoomRepository.UpdateRoom(ctx, model.ChinchiroRoom{
			ID:     participant.RoomID,
			Status: model.ChinchiroRoomStatusFinished,
		})
	}
	return nil
}

func NewChinchiroRoomService(
	chinchiroRoomRepository model.ChinchiroRoomRepository,
	chinchiroRoomParticipantRepository model.ChinchiroRoomParticipantRepository,
) ChinchiroRoomService {
	return &chinchiroRoomService{
		chinchiroRoomRepository:            chinchiroRoomRepository,
		chinchiroRoomParticipantRepository: chinchiroRoomParticipantRepository,
	}
}
