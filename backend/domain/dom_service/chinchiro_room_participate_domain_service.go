package dom_service

import (
	"wolfort-games/domain/model"
)

type ChinchiroRoomParticipateDomainService interface {
	AssertParticipate(room model.ChinchiroRoom, player model.Player, authorities []model.PlayerAuthority, password *string) error
}

type chinchiroRoomParticipateDomainService struct {
	roomMasterDomainService ChinchiroRoomMasterDomainService
}

func NewChinchiroRoomParticipateDomainService(
	roomMasterDomainService ChinchiroRoomMasterDomainService,
) ChinchiroRoomParticipateDomainService {
	return &chinchiroRoomParticipateDomainService{
		roomMasterDomainService: roomMasterDomainService,
	}
}

func (s *chinchiroRoomParticipateDomainService) AssertParticipate(
	room model.ChinchiroRoom,
	player model.Player,
	authorities []model.PlayerAuthority,
	password *string,
) error {
	if room.Status == model.ChinchiroRoomStatusFinished {
		// 終了後は参加不可
		return model.NewErrBusiness("終了後は参加不可です")
	}

	if room.Settings.Password.HasPassword() {
		if password == nil || *room.Settings.Password.Password != *password {
			return model.NewErrBusiness("パスワードが違います")
		}
	}

	return nil
}
