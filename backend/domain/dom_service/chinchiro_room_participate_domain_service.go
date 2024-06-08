package dom_service

import (
	"chat-role-play/domain/model"
)

type ParticipateDomainService interface {
	AssertParticipate(room model.ChinchiroRoom, player model.Player, authorities []model.PlayerAuthority, password *string) error
}

type participateDomainService struct {
	roomMasterDomainService ChinchiroRoomMasterDomainService
}

func NewParticipateDomainService(
	roomMasterDomainService ChinchiroRoomMasterDomainService,
) ParticipateDomainService {
	return &participateDomainService{
		roomMasterDomainService: roomMasterDomainService,
	}
}

func (s *participateDomainService) AssertParticipate(
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
