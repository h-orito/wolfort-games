package dom_service

import (
	"chat-role-play/domain/model"
	"chat-role-play/util/array"
	"fmt"
)

type ChinchiroRoomMasterDomainService interface {
	IsRoomMaster(room model.ChinchiroRoom, player model.Player, authorities []model.PlayerAuthority) bool
	AssertModifyRoomMaster(room model.ChinchiroRoom, player model.Player, authorities []model.PlayerAuthority) error
}

type chinchiroRoomMasterDomainService struct {
}

func NewChinchiroRoomMasterDomainService() ChinchiroRoomMasterDomainService {
	return &chinchiroRoomMasterDomainService{}
}

func (ds *chinchiroRoomMasterDomainService) IsRoomMaster(
	room model.ChinchiroRoom,
	player model.Player,
	authorities []model.PlayerAuthority,
) bool {
	// AdminはRM扱い
	if array.Any(authorities, func(a model.PlayerAuthority) bool {
		return a.IsAdmin()
	}) {
		return true
	}
	return array.Any(room.RoomMasterIDs, func(id uint32) bool {
		return id == player.ID
	})
}

func (ds *chinchiroRoomMasterDomainService) AssertModifyRoomMaster(
	room model.ChinchiroRoom,
	player model.Player,
	authorities []model.PlayerAuthority,
) error {
	if !ds.IsRoomMaster(room, player, authorities) {
		return fmt.Errorf("player is not room master")
	}
	if room.Status != model.ChinchiroRoomStatusOpened {
		return fmt.Errorf("room is not opened")
	}
	return nil
}
