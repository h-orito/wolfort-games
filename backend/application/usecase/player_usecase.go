package usecase

import (
	"wolfort-games/application/app_service"
	"wolfort-games/domain/model"
)

type PlayerUsecase interface {
	FindPlayers(query model.PlayersQuery) ([]model.Player, error)
	Find(ID uint32) (player *model.Player, err error)
	FindByName(name string) (player *model.Player, err error)
	FindByUserName(username string) (player *model.Player, err error)
	FindAuthorities(ID uint32) (authorities []model.PlayerAuthority, err error)
}

type playerUsecase struct {
	playerService app_service.PlayerService
	transaction   Transaction
}

func NewPlayerUsecase(playerService app_service.PlayerService,
	tx Transaction) PlayerUsecase {
	return &playerUsecase{
		playerService: playerService,
		transaction:   tx,
	}
}

func (s *playerUsecase) FindPlayers(query model.PlayersQuery) ([]model.Player, error) {
	return s.playerService.FindPlayers(query)
}

func (s *playerUsecase) Find(ID uint32) (player *model.Player, err error) {
	return s.playerService.Find(ID)
}

func (s *playerUsecase) FindByName(name string) (player *model.Player, err error) {
	return s.playerService.FindByName(name)
}

func (s *playerUsecase) FindByUserName(username string) (player *model.Player, err error) {
	return s.playerService.FindByUserName(username)
}

// FindAuthorities implements PlayerUsecase.
func (s *playerUsecase) FindAuthorities(ID uint32) (authorities []model.PlayerAuthority, err error) {
	return s.playerService.FindAuthorities(ID)
}
