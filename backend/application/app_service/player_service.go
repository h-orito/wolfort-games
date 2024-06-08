package app_service

import (
	"chat-role-play/domain/model"
	"context"
)

type PlayerService interface {
	FindPlayers(query model.PlayersQuery) ([]model.Player, error)
	Find(ID uint32) (player *model.Player, err error)
	FindByName(name string) (player *model.Player, err error)
	FindByUserName(name string) (player *model.Player, err error)
	Save(ctx context.Context, player model.Player) (saved *model.Player, err error)
	FindAuthorities(ID uint32) (authorities []model.PlayerAuthority, err error)
}

type playerService struct {
	playerRepository model.PlayerRepository
	userRepository   model.UserRepository
}

func NewPlayerService(
	playerRepository model.PlayerRepository,
	userRepository model.UserRepository,
) PlayerService {
	return &playerService{
		playerRepository: playerRepository,
		userRepository:   userRepository,
	}
}

func (s *playerService) FindPlayers(query model.PlayersQuery) ([]model.Player, error) {
	return s.playerRepository.FindPlayers(query)
}

func (s *playerService) Find(ID uint32) (player *model.Player, err error) {
	return s.playerRepository.Find(ID)
}

func (s *playerService) FindByName(name string) (player *model.Player, err error) {
	return s.playerRepository.FindByName(name)
}

func (s *playerService) FindByUserName(name string) (player *model.Player, err error) {
	return s.playerRepository.FindByUserName(name)
}

func (s *playerService) Save(ctx context.Context, player model.Player) (saved *model.Player, err error) {
	return s.playerRepository.Save(ctx, &player)
}

// FindAuthorities implements PlayerService.
func (s *playerService) FindAuthorities(ID uint32) (authorities []model.PlayerAuthority, err error) {
	return s.userRepository.FindPlayerAuthorities(ID)
}
