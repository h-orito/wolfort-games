package usecase

import (
	"context"
	"wolfort-games/application/app_service"
	"wolfort-games/domain/model"
)

type PlayerUsecase interface {
	FindPlayers(query model.PlayersQuery) ([]model.Player, error)
	Find(ID uint32) (player *model.Player, err error)
	FindByName(name string) (player *model.Player, err error)
	FindByUserName(username string) (player *model.Player, err error)
	FindAuthorities(ID uint32) (authorities []model.PlayerAuthority, err error)
	UpdatePlayerProfile(ctx context.Context, user model.User, player model.Player) error
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

func (s *playerUsecase) FindAuthorities(ID uint32) (authorities []model.PlayerAuthority, err error) {
	return s.playerService.FindAuthorities(ID)
}

func (u *playerUsecase) UpdatePlayerProfile(ctx context.Context, user model.User, player model.Player) error {
	_, err := u.transaction.DoInTx(ctx, func(ctx context.Context) (interface{}, error) {
		p, err := u.playerService.FindByUserName(user.UserName)
		if err != nil {
			return nil, err
		}
		player.ID = p.ID
		u.playerService.Save(ctx, player)
		return nil, nil
	})
	return err
}
