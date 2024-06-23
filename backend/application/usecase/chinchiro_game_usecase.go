package usecase

import (
	"context"
	"fmt"
	"wolfort-games/application/app_service"
	"wolfort-games/domain/dom_service"
	"wolfort-games/domain/model"
	"wolfort-games/util/array"
)

type ChinchiroGameUsecase interface {
	// game
	FindChinchiroGames(query model.ChinchiroGamesQuery) (model.ChinchiroGames, error)
	FindChinchiroGame(ID uint32) (*model.ChinchiroGame, error)
	RegisterChinchiroGame(ctx context.Context, user model.User, roomID uint32, game model.ChinchiroGame) (*model.ChinchiroGame, error)
	UpdateChinchiroGameStatus(ctx context.Context, user model.User, gameID uint32, status model.ChinchiroGameStatus) error
	// game participant
	FindChinchiroGameParticipants(query model.ChinchiroGameParticipantsQuery) (model.ChinchiroGameParticipants, error)
	FindMyChinchiroGameParticipant(gameID uint32, user model.User) (*model.ChinchiroGameParticipant, error)
	// game turn
	FindChinchiroGameTurns(query model.ChinchiroGameTurnsQuery) (model.ChinchiroGameTurns, error)
	FindChinchiroGameTurn(ID uint32) (*model.ChinchiroGameTurn, error)
	UpdateChinchiroGameTurnStatus(ctx context.Context, user model.User, turnID uint32, status model.ChinchiroGameTurnStatus) error
	// game turn roll
	FindChinchiroGameTurnRolls(query model.ChinchiroGameTurnRollsQuery) (model.ChinchiroGameTurnRolls, error)
	FindChinchiroGameTurnResults(query model.ChinchiroGameTurnResultsQuery) (model.ChinchiroGameTurnResults, error)
	BetChinchiroGameTurn(ctx context.Context, user model.User, turnID uint32, betAmount int) error
	RollChinchiroGameTurn(ctx context.Context, user model.User, turnID uint32) error
}

type chinchiroGameUsecase struct {
	chinchiroRoomService             app_service.ChinchiroRoomService
	chinchiroGameService             app_service.ChinchiroGameService
	playerService                    app_service.PlayerService
	chinchiroRoomMasterDomainService dom_service.ChinchiroRoomMasterDomainService
	transaction                      Transaction
}

func (u *chinchiroGameUsecase) FindChinchiroGames(query model.ChinchiroGamesQuery) (model.ChinchiroGames, error) {
	return u.chinchiroGameService.FindChinchiroGames(query)
}

func (u *chinchiroGameUsecase) FindChinchiroGame(ID uint32) (*model.ChinchiroGame, error) {
	return u.chinchiroGameService.FindChinchiroGame(ID)
}

func (u *chinchiroGameUsecase) RegisterChinchiroGame(ctx context.Context, user model.User, roomID uint32, game model.ChinchiroGame) (*model.ChinchiroGame, error) {
	r, err := u.transaction.DoInTx(ctx, func(ctx context.Context) (interface{}, error) {
		err := u.assertRegisterGame(user, roomID)
		if err != nil {
			return nil, err
		}
		// game
		return u.chinchiroGameService.RegisterChinchiroGame(ctx, roomID, game)
	})
	if err != nil {
		return nil, err
	}
	return r.(*model.ChinchiroGame), nil
}

func (u *chinchiroGameUsecase) UpdateChinchiroGameStatus(ctx context.Context, user model.User, gameID uint32, status model.ChinchiroGameStatus) error {
	_, err := u.transaction.DoInTx(ctx, func(ctx context.Context) (interface{}, error) {
		err := u.assertUpdateChinchiroGameStatus(user, gameID)
		if err != nil {
			return nil, err
		}
		return nil, u.chinchiroGameService.UpdateChinchiroGameStatus(ctx, gameID, status)
	})
	return err
}

func (u *chinchiroGameUsecase) FindChinchiroGameParticipants(query model.ChinchiroGameParticipantsQuery) (model.ChinchiroGameParticipants, error) {
	return u.chinchiroGameService.FindChinchiroGameParticipants(query)
}

func (u *chinchiroGameUsecase) FindMyChinchiroGameParticipant(gameID uint32, user model.User) (*model.ChinchiroGameParticipant, error) {
	return u.findMyChinchiroGameParticipant(gameID, user)
}

func (u *chinchiroGameUsecase) FindChinchiroGameTurns(query model.ChinchiroGameTurnsQuery) (model.ChinchiroGameTurns, error) {
	return u.chinchiroGameService.FindChinchiroGameTurns(query)
}

func (u *chinchiroGameUsecase) FindChinchiroGameTurn(ID uint32) (*model.ChinchiroGameTurn, error) {
	return u.chinchiroGameService.FindChinchiroGameTurn(ID)
}

func (u *chinchiroGameUsecase) UpdateChinchiroGameTurnStatus(ctx context.Context, user model.User, turnID uint32, status model.ChinchiroGameTurnStatus) error {
	_, err := u.transaction.DoInTx(ctx, func(ctx context.Context) (interface{}, error) {
		err := u.assertUpdateChinchiroGameTurnStatus(user, turnID)
		if err != nil {
			return nil, err
		}
		return nil, u.chinchiroGameService.UpdateChinchiroGameTurn(ctx, turnID, status)
	})
	return err
}

func (u *chinchiroGameUsecase) FindChinchiroGameTurnRolls(query model.ChinchiroGameTurnRollsQuery) (model.ChinchiroGameTurnRolls, error) {
	return u.chinchiroGameService.FindChinchiroGameTurnRolls(query)
}

func (u *chinchiroGameUsecase) FindChinchiroGameTurnResults(query model.ChinchiroGameTurnResultsQuery) (model.ChinchiroGameTurnResults, error) {
	return u.chinchiroGameService.FindChinchiroGameTurnResults(query)
}

func (u *chinchiroGameUsecase) BetChinchiroGameTurn(ctx context.Context, user model.User, turnID uint32, betAmount int) error {
	_, err := u.transaction.DoInTx(ctx, func(ctx context.Context) (interface{}, error) {
		game, err := u.findGameByTurnID(turnID)
		if err != nil {
			return nil, err
		}
		if game == nil {
			return nil, fmt.Errorf("game not found")
		}
		gameParticipant, err := u.findMyChinchiroGameParticipant(game.ID, user)
		if err != nil {
			return nil, err
		}
		if gameParticipant == nil {
			return nil, fmt.Errorf("game participant not found")
		}
		return nil, u.chinchiroGameService.RegisterChinchiroGameTurnBet(ctx, turnID, gameParticipant.ID, betAmount)
	})
	return err
}

func (u *chinchiroGameUsecase) RollChinchiroGameTurn(ctx context.Context, user model.User, turnID uint32) error {
	_, err := u.transaction.DoInTx(ctx, func(ctx context.Context) (interface{}, error) {
		game, err := u.findGameByTurnID(turnID)
		if err != nil {
			return nil, err
		}
		if game == nil {
			return nil, fmt.Errorf("game not found")
		}
		gameParticipant, err := u.findMyChinchiroGameParticipant(game.ID, user)
		if err != nil {
			return nil, err
		}
		if gameParticipant == nil {
			return nil, fmt.Errorf("game participant not found")
		}
		return nil, u.chinchiroGameService.RegisterChinchiroGameTurnRoll(
			ctx,
			turnID,
			gameParticipant.ID,
			model.RollChinchiroDice(),
		)
	})
	return err
}

// ----------------------------

func (u *chinchiroGameUsecase) assertRegisterGame(
	user model.User,
	roomID uint32,
) error {
	player, authorities, err := u.findPlayerAndAuthorities(user)
	if err != nil {
		return err
	}
	if player == nil || authorities == nil {
		return fmt.Errorf("player not found")
	}
	room, err := u.chinchiroRoomService.FindChinchiroRoom(roomID)
	if err != nil {
		return err
	}
	err = u.chinchiroRoomMasterDomainService.AssertModifyRoom(*room, *player, authorities)
	if err != nil {
		return err
	}
	games, err := u.chinchiroGameService.FindChinchiroGames(model.ChinchiroGamesQuery{
		RoomID: &roomID,
	})
	if err != nil {
		return err
	}
	if array.Any(games.List, func(g model.ChinchiroGame) bool {
		return g.Status == model.ChinchiroGameStatusProgress
	}) {
		return fmt.Errorf("room already has opened game")
	}
	isExcludeGone := false
	participants, err := u.chinchiroRoomService.FindChinchiroRoomParticipants(model.ChinchiroRoomParticipantsQuery{
		RoomIDs:       &[]uint32{roomID},
		IsExcludeGone: &isExcludeGone,
	})
	if err != nil {
		return err
	}
	if len(participants.List) < 2 {
		return fmt.Errorf("room has not enough participants")
	}
	return nil
}

func (u *chinchiroGameUsecase) assertUpdateChinchiroGameStatus(
	user model.User,
	gameID uint32,
) error {
	player, authorities, err := u.findPlayerAndAuthorities(user)
	if err != nil {
		return err
	}
	if player == nil || authorities == nil {
		return fmt.Errorf("player not found")
	}

	game, err := u.chinchiroGameService.FindChinchiroGame(gameID)
	if err != nil {
		return err
	}
	if game == nil {
		return fmt.Errorf("game not found")
	}
	if game.Status == model.ChinchiroGameStatusFinished {
		return fmt.Errorf("game already finished")
	}
	room, err := u.chinchiroRoomService.FindChinchiroRoom(game.ID)
	if err != nil {
		return err
	}
	err = u.chinchiroRoomMasterDomainService.AssertModifyRoom(*room, *player, authorities)
	if err != nil {
		return err
	}
	return nil
}

func (u *chinchiroGameUsecase) assertUpdateChinchiroGameTurnStatus(
	user model.User,
	turnID uint32,
) error {
	player, authorities, err := u.findPlayerAndAuthorities(user)
	if err != nil {
		return err
	}
	if player == nil || authorities == nil {
		return fmt.Errorf("player not found")
	}
	room, err := u.findRoomByTurnID(turnID)
	if err != nil {
		return err
	}
	if room == nil {
		return fmt.Errorf("room not found")
	}
	err = u.chinchiroRoomMasterDomainService.AssertModifyRoom(*room, *player, authorities)
	if err != nil {
		return err
	}
	return nil
}

func (u *chinchiroGameUsecase) findMyChinchiroGameParticipant(gameID uint32, user model.User) (*model.ChinchiroGameParticipant, error) {
	player, err := u.playerService.FindByUserName(user.UserName)
	if err != nil {
		return nil, err
	}
	if player == nil {
		return nil, fmt.Errorf("player not found")
	}
	game, err := u.chinchiroGameService.FindChinchiroGame(gameID)
	if err != nil {
		return nil, err
	}
	if game == nil {
		return nil, fmt.Errorf("game not found")
	}
	roomParticipant, err := u.chinchiroRoomService.FindChinchiroRoomParticipant(model.ChinchiroRoomParticipantQuery{
		PlayerID: &player.ID,
		RoomID:   &game.RoomID,
	})
	if err != nil {
		return nil, err
	}
	if roomParticipant == nil {
		return nil, fmt.Errorf("room participant not found")
	}
	return u.chinchiroGameService.FindChinchiroGameParticipant(model.ChinchiroGameParticipantQuery{
		GameID:            &gameID,
		RoomParticipantID: &roomParticipant.ID,
	})
}

func (u *chinchiroGameUsecase) findGameByTurnID(turnID uint32) (*model.ChinchiroGame, error) {
	turn, err := u.chinchiroGameService.FindChinchiroGameTurn(turnID)
	if err != nil {
		return nil, err
	}
	if turn == nil {
		return nil, fmt.Errorf("turn not found")
	}
	return u.chinchiroGameService.FindChinchiroGame(turn.GameID)
}

func (u *chinchiroGameUsecase) findRoomByTurnID(turnID uint32) (*model.ChinchiroRoom, error) {
	game, err := u.findGameByTurnID(turnID)
	if err != nil {
		return nil, err
	}
	if game == nil {
		return nil, fmt.Errorf("game not found")
	}
	return u.chinchiroRoomService.FindChinchiroRoom(game.RoomID)
}

func (u *chinchiroGameUsecase) findPlayerAndAuthorities(
	user model.User,
) (*model.Player, []model.PlayerAuthority, error) {
	player, err := u.playerService.FindByUserName(user.UserName)
	if err != nil {
		return nil, nil, err
	}
	authorities, err := u.playerService.FindAuthorities(player.ID)
	if err != nil {
		return nil, nil, err
	}
	return player, authorities, nil
}

// ----------------------------

func NewChinchiroGameUsecase(
	chinchiroRoomService app_service.ChinchiroRoomService,
	chinchiroGameService app_service.ChinchiroGameService,
	playerService app_service.PlayerService,
	chinchiroRoomMasterDomainService dom_service.ChinchiroRoomMasterDomainService,
	tx Transaction,
) ChinchiroGameUsecase {
	return &chinchiroGameUsecase{
		chinchiroRoomService:             chinchiroRoomService,
		chinchiroGameService:             chinchiroGameService,
		playerService:                    playerService,
		chinchiroRoomMasterDomainService: chinchiroRoomMasterDomainService,
		transaction:                      tx,
	}
}
