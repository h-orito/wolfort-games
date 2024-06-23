package usecase

import (
	"context"
	"fmt"
	"wolfort-games/application/app_service"
	"wolfort-games/domain/dom_service"
	"wolfort-games/domain/model"
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
	chinchiroRoomService                  app_service.ChinchiroRoomService
	playerService                         app_service.PlayerService
	chinchiroRoomMasterDomainService      dom_service.ChinchiroRoomMasterDomainService
	chinchiroRoomParticipateDomainService dom_service.ChinchiroRoomParticipateDomainService
	transaction                           Transaction
}

func (u *chinchiroRoomUsecase) FindChinchiroRooms(query model.ChinchiroRoomsQuery) ([]model.ChinchiroRoom, error) {
	return u.chinchiroRoomService.FindChinchiroRooms(query)
}

func (u *chinchiroRoomUsecase) FindChinchiroRoom(ID uint32) (*model.ChinchiroRoom, error) {
	return u.chinchiroRoomService.FindChinchiroRoom(ID)
}

func (u *chinchiroRoomUsecase) RegisterChinchiroRoom(ctx context.Context, user model.User, room model.ChinchiroRoom) (*model.ChinchiroRoom, error) {
	r, err := u.transaction.DoInTx(ctx, func(ctx context.Context) (interface{}, error) {
		player, err := u.playerService.FindByUserName(user.UserName)
		if err != nil {
			return nil, err
		}
		err = u.assertRegisterRoom(user, room)
		if err != nil {
			return nil, err
		}
		// room
		rm, err := u.chinchiroRoomService.RegisterChinchiroRoom(ctx, room)
		if err != nil {
			return nil, err
		}
		// room master
		_, err = u.chinchiroRoomService.RegisterChinchiroRoomMaster(ctx, rm.ID, player.ID)
		if err != nil {
			return nil, err
		}
		return rm, nil
	})
	if err != nil {
		return nil, err
	}
	return r.(*model.ChinchiroRoom), nil
}

func (u *chinchiroRoomUsecase) assertRegisterRoom(
	user model.User,
	room model.ChinchiroRoom,
) error {
	if room.Name == "" {
		return fmt.Errorf("room name is required")
	}
	return nil
}

func (u *chinchiroRoomUsecase) UpdateChinchiroRoomStatus(ctx context.Context, user model.User, roomID uint32, status model.ChinchiroRoomStatus) error {
	_, err := u.transaction.DoInTx(ctx, func(ctx context.Context) (interface{}, error) {
		room, err := u.chinchiroRoomService.FindChinchiroRoom(roomID)
		if err != nil {
			return nil, err
		}
		player, err := u.playerService.FindByUserName(user.UserName)
		if err != nil {
			return nil, err
		}
		authorities, err := u.playerService.FindAuthorities(player.ID)
		if err != nil {
			return nil, err
		}
		err = u.chinchiroRoomMasterDomainService.AssertModifyRoom(*room, *player, authorities)
		if err != nil {
			return nil, err
		}
		return nil, u.chinchiroRoomService.UpdateChinchiroRoomStatus(ctx, roomID, status)
	})
	return err
}

func (u *chinchiroRoomUsecase) UpdateChinchiroRoomSettings(ctx context.Context, user model.User, roomID uint32, settings model.ChinchiroRoomSettings) error {
	_, err := u.transaction.DoInTx(ctx, func(ctx context.Context) (interface{}, error) {
		room, err := u.chinchiroRoomService.FindChinchiroRoom(roomID)
		if err != nil {
			return nil, err
		}
		player, err := u.playerService.FindByUserName(user.UserName)
		if err != nil {
			return nil, err
		}
		authorities, err := u.playerService.FindAuthorities(player.ID)
		if err != nil {
			return nil, err
		}
		err = u.chinchiroRoomMasterDomainService.AssertModifyRoom(*room, *player, authorities)
		if err != nil {
			return nil, err
		}
		return nil, u.chinchiroRoomService.UpdateChinchiroRoomSettings(ctx, roomID, settings)
	})
	return err
}

func (u *chinchiroRoomUsecase) FindChinchiroRoomMasters(query model.ChinchiroRoomMastersQuery) ([]model.ChinchiroRoomMaster, error) {
	return u.chinchiroRoomService.FindChinchiroRoomMasters(query)
}

func (u *chinchiroRoomUsecase) RegisterChinchiroRoomMaster(ctx context.Context, user model.User, roomID uint32, playerID uint32) (model.ChinchiroRoomMaster, error) {
	rm, err := u.transaction.DoInTx(ctx, func(ctx context.Context) (interface{}, error) {
		room, err := u.chinchiroRoomService.FindChinchiroRoom(roomID)
		if err != nil {
			return nil, err
		}
		player, err := u.playerService.FindByUserName(user.UserName)
		if err != nil {
			return nil, err
		}
		authorities, err := u.playerService.FindAuthorities(player.ID)
		if err != nil {
			return nil, err
		}
		err = u.chinchiroRoomMasterDomainService.AssertModifyRoom(*room, *player, authorities)
		if err != nil {
			return nil, err
		}
		return u.chinchiroRoomService.RegisterChinchiroRoomMaster(ctx, roomID, playerID)
	})
	return *rm.(*model.ChinchiroRoomMaster), err
}

func (u *chinchiroRoomUsecase) DeleteChinchiroRoomMaster(ctx context.Context, user model.User, roomID uint32, masterID uint32) error {
	_, err := u.transaction.DoInTx(ctx, func(ctx context.Context) (interface{}, error) {
		room, err := u.chinchiroRoomService.FindChinchiroRoom(roomID)
		if err != nil {
			return nil, err
		}
		player, err := u.playerService.FindByUserName(user.UserName)
		if err != nil {
			return nil, err
		}
		authorities, err := u.playerService.FindAuthorities(player.ID)
		if err != nil {
			return nil, err
		}
		err = u.chinchiroRoomMasterDomainService.AssertModifyRoom(*room, *player, authorities)
		if err != nil {
			return nil, err
		}
		return nil, u.chinchiroRoomService.DeleteChinchiroRoomMaster(ctx, masterID)
	})
	return err
}

func (u *chinchiroRoomUsecase) FindChinchiroRoomParticipants(query model.ChinchiroRoomParticipantsQuery) (model.ChinchiroRoomParticipants, error) {
	return u.chinchiroRoomService.FindChinchiroRoomParticipants(query)
}

func (u *chinchiroRoomUsecase) FindChinchiroRoomParticipant(ID uint32) (*model.ChinchiroRoomParticipant, error) {
	return u.chinchiroRoomService.FindChinchiroRoomParticipant(model.ChinchiroRoomParticipantQuery{
		ID: &ID,
	})
}

func (u *chinchiroRoomUsecase) FindMyChinchiroRoomParticipant(user model.User, roomID uint32) (*model.ChinchiroRoomParticipant, error) {
	player, err := u.playerService.FindByUserName(user.UserName)
	if err != nil {
		return nil, err
	}
	if player == nil {
		return nil, nil
	}
	return u.chinchiroRoomService.FindChinchiroRoomParticipant(model.ChinchiroRoomParticipantQuery{
		PlayerID: &player.ID,
		RoomID:   &roomID,
	})
}

func (u *chinchiroRoomUsecase) ParticipateChinchiroRoom(
	ctx context.Context,
	user model.User,
	roomID uint32,
	participant model.ChinchiroRoomParticipant,
	password *string,
) (*model.ChinchiroRoomParticipant, error) {
	p, err := u.transaction.DoInTx(ctx, func(ctx context.Context) (interface{}, error) {
		room, err := u.chinchiroRoomService.FindChinchiroRoom(roomID)
		if err != nil {
			return nil, err
		}
		player, err := u.playerService.FindByUserName(user.UserName)
		if err != nil {
			return nil, err
		}
		authorities, err := u.playerService.FindAuthorities(player.ID)
		if err != nil {
			return nil, err
		}
		err = u.chinchiroRoomParticipateDomainService.AssertParticipate(*room, *player, authorities, password)
		if err != nil {
			return nil, err
		}
		return u.chinchiroRoomService.RegisterChinchiroRoomParticipant(ctx, roomID, model.ChinchiroRoomParticipant{
			Name:     participant.Name,
			PlayerID: player.ID,
			RoomID:   roomID,
			IsGone:   false,
		})
	})
	return p.(*model.ChinchiroRoomParticipant), err
}

func (u *chinchiroRoomUsecase) UpdateChinchiroRoomParticipant(
	ctx context.Context,
	user model.User,
	participantID uint32,
	participant model.ChinchiroRoomParticipant,
) error {
	_, err := u.transaction.DoInTx(ctx, func(ctx context.Context) (interface{}, error) {
		p, err := u.chinchiroRoomService.FindChinchiroRoomParticipant(model.ChinchiroRoomParticipantQuery{
			ID: &participantID,
		})
		if err != nil {
			return nil, err
		}
		if p == nil {
			return nil, fmt.Errorf("participant not found")
		}
		player, err := u.playerService.FindByUserName(user.UserName)
		if err != nil {
			return nil, err
		}
		if p.PlayerID != player.ID {
			return nil, fmt.Errorf("you are not the participant")
		}
		return nil, u.chinchiroRoomService.UpdateChinchiroRoomParticipant(ctx, participantID, participant)
	})
	return err
}

func (u *chinchiroRoomUsecase) LeaveChinchiroRoom(ctx context.Context, user model.User, roomID uint32) error {
	_, err := u.transaction.DoInTx(ctx, func(ctx context.Context) (interface{}, error) {
		player, err := u.playerService.FindByUserName(user.UserName)
		if err != nil {
			return nil, err
		}
		p, err := u.chinchiroRoomService.FindChinchiroRoomParticipant(model.ChinchiroRoomParticipantQuery{
			RoomID:   &roomID,
			PlayerID: &player.ID,
		})
		if err != nil {
			return nil, err
		}
		if p == nil {
			return nil, fmt.Errorf("participant not found")
		}
		return nil, u.chinchiroRoomService.DeleteChinchiroRoomParticipant(ctx, p.ID)
	})
	return err
}

func (u *chinchiroRoomUsecase) DeleteChinchiroRoomParticipant(ctx context.Context, user model.User, participantID uint32) error {
	_, err := u.transaction.DoInTx(ctx, func(ctx context.Context) (interface{}, error) {
		player, err := u.playerService.FindByUserName(user.UserName)
		if err != nil {
			return nil, err
		}
		authorities, err := u.playerService.FindAuthorities(player.ID)
		if err != nil {
			return nil, err
		}
		p, err := u.chinchiroRoomService.FindChinchiroRoomParticipant(model.ChinchiroRoomParticipantQuery{
			ID: &participantID,
		})
		if err != nil {
			return nil, err
		}
		room, err := u.chinchiroRoomService.FindChinchiroRoom(p.RoomID)
		if err != nil {
			return nil, err
		}
		err = u.chinchiroRoomMasterDomainService.AssertModifyRoom(*room, *player, authorities)
		if err != nil {
			return nil, err
		}
		return nil, u.chinchiroRoomService.DeleteChinchiroRoomParticipant(ctx, participantID)
	})
	return err
}

func NewChinchiroRoomUsecase(
	chinchiroRoomService app_service.ChinchiroRoomService,
	playerService app_service.PlayerService,
	chinchiroRoomMasterDomainService dom_service.ChinchiroRoomMasterDomainService,
	chinchiroRoomParticipateDomainService dom_service.ChinchiroRoomParticipateDomainService,
	tx Transaction,
) ChinchiroRoomUsecase {
	return &chinchiroRoomUsecase{
		chinchiroRoomService:                  chinchiroRoomService,
		playerService:                         playerService,
		chinchiroRoomMasterDomainService:      chinchiroRoomMasterDomainService,
		chinchiroRoomParticipateDomainService: chinchiroRoomParticipateDomainService,
		transaction:                           tx,
	}
}
