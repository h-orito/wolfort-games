package graphql

import (
	"context"
	"fmt"
	"wolfort-games/adaptor/auth"
	"wolfort-games/domain/model"
	"wolfort-games/middleware/graph/gqlmodel"
	"wolfort-games/util/array"

	"github.com/graph-gophers/dataloader"
)

func (r *chinchiroGameResolver) participants(ctx context.Context, obj *gqlmodel.ChinchiroGame) ([]*gqlmodel.ChinchiroGameParticipant, error) {
	if obj.ParticipantIDs == nil || len(obj.ParticipantIDs) == 0 {
		return nil, nil
	}
	thunk := r.loaders.ChinchiroGameParticipantLoader.LoadMany(ctx, dataloader.NewKeysFromStrings(obj.ParticipantIDs))
	c, errs := thunk()
	if errs != nil || len(errs) > 0 {
		return nil, errs[0]
	}
	return array.Map(c, func(c interface{}) *gqlmodel.ChinchiroGameParticipant {
		p := c.(*model.ChinchiroGameParticipant)
		return MapToChinchiroGameParticipant(p)
	}), nil
}

func (r *chinchiroGameResolver) turns(ctx context.Context, obj *gqlmodel.ChinchiroGame) ([]*gqlmodel.ChinchiroGameTurn, error) {
	if obj.TurnIDs == nil || len(obj.TurnIDs) == 0 {
		return nil, nil
	}
	thunk := r.loaders.ChinchiroGameTurnLoader.LoadMany(ctx, dataloader.NewKeysFromStrings(obj.TurnIDs))
	c, errs := thunk()
	if errs != nil || len(errs) > 0 {
		return nil, errs[0]
	}
	return array.Map(c, func(c interface{}) *gqlmodel.ChinchiroGameTurn {
		t := c.(*model.ChinchiroGameTurn)
		return MapToChinchiroGameTurn(t)
	}), nil
}

func (r *chinchiroGameParticipantResolver) roomParticipant(ctx context.Context, obj *gqlmodel.ChinchiroGameParticipant) (*gqlmodel.ChinchiroRoomParticipant, error) {
	thunk := r.loaders.ChinchiroRoomParticipantLoader.Load(ctx, dataloader.StringKey(obj.RoomParticipantID))
	p, err := thunk()
	if err != nil {
		return nil, err
	}
	roomParticipant := p.(*model.ChinchiroRoomParticipant)
	return MapToRoomParticipant(roomParticipant), nil
}

func (r *chinchiroGameTurnResolver) dealer(ctx context.Context, obj *gqlmodel.ChinchiroGameTurn) (*gqlmodel.ChinchiroGameParticipant, error) {
	thunk := r.loaders.ChinchiroGameParticipantLoader.Load(ctx, dataloader.StringKey(obj.DealerID))
	p, err := thunk()
	if err != nil {
		return nil, err
	}
	gameParticipant := p.(*model.ChinchiroGameParticipant)
	return MapToChinchiroGameParticipant(gameParticipant), nil
}

func (r *chinchiroGameTurnResolver) nextRoller(ctx context.Context, obj *gqlmodel.ChinchiroGameTurn) (*gqlmodel.ChinchiroGameParticipant, error) {
	thunk := r.loaders.ChinchiroGameParticipantLoader.Load(ctx, dataloader.StringKey(*obj.NextRollerID))
	p, err := thunk()
	if err != nil {
		return nil, err
	}
	gameParticipant := p.(*model.ChinchiroGameParticipant)
	return MapToChinchiroGameParticipant(gameParticipant), nil
}

func (r *chinchiroGameTurnResolver) rolls(ctx context.Context, obj *gqlmodel.ChinchiroGameTurn) ([]*gqlmodel.ChinchiroGameTurnParticipantRoll, error) {
	if obj.RollIDs == nil || len(obj.RollIDs) == 0 {
		return nil, nil
	}
	thunk := r.loaders.ChinchiroGameTurnRollLoader.LoadMany(ctx, dataloader.NewKeysFromStrings(obj.RollIDs))
	c, errs := thunk()
	if errs != nil || len(errs) > 0 {
		return nil, errs[0]
	}
	return array.Map(c, func(c interface{}) *gqlmodel.ChinchiroGameTurnParticipantRoll {
		r := c.(*model.ChinchiroGameTurnRoll)
		return MapToChinchiroGameParticipantRoll(r)
	}), nil
}

func (r *chinchiroGameTurnResolver) results(ctx context.Context, obj *gqlmodel.ChinchiroGameTurn) ([]*gqlmodel.ChinchiroGameTurnParticipantResult, error) {
	if obj.ResultIDs == nil || len(obj.ResultIDs) == 0 {
		return nil, nil
	}
	thunk := r.loaders.ChinchiroGameTurnResultLoader.LoadMany(ctx, dataloader.NewKeysFromStrings(obj.ResultIDs))
	c, errs := thunk()
	if errs != nil || len(errs) > 0 {
		return nil, errs[0]
	}
	return array.Map(c, func(c interface{}) *gqlmodel.ChinchiroGameTurnParticipantResult {
		r := c.(*model.ChinchiroGameTurnResult)
		return MapToChinchiroGameParticipantResult(r)
	}), nil
}

func (r *chinchiroGameTurnParticipantResultResolver) turn(ctx context.Context, obj *gqlmodel.ChinchiroGameTurnParticipantResult) (*gqlmodel.ChinchiroGameTurn, error) {
	thunk := r.loaders.ChinchiroGameTurnLoader.Load(ctx, dataloader.StringKey(obj.GameTurnID))
	p, err := thunk()
	if err != nil {
		return nil, err
	}
	turn := p.(*model.ChinchiroGameTurn)
	return MapToChinchiroGameTurn(turn), nil
}

func (r *chinchiroGameTurnParticipantResultResolver) participant(ctx context.Context, obj *gqlmodel.ChinchiroGameTurnParticipantResult) (*gqlmodel.ChinchiroGameParticipant, error) {
	thunk := r.loaders.ChinchiroGameParticipantLoader.Load(ctx, dataloader.StringKey(obj.ParticipantID))
	p, err := thunk()
	if err != nil {
		return nil, err
	}
	turn := p.(*model.ChinchiroGameParticipant)
	return MapToChinchiroGameParticipant(turn), nil
}

func (r *chinchiroGameTurnParticipantRollResolver) turn(ctx context.Context, obj *gqlmodel.ChinchiroGameTurnParticipantRoll) (*gqlmodel.ChinchiroGameTurn, error) {
	thunk := r.loaders.ChinchiroGameTurnLoader.Load(ctx, dataloader.StringKey(obj.TurnID))
	p, err := thunk()
	if err != nil {
		return nil, err
	}
	turn := p.(*model.ChinchiroGameTurn)
	return MapToChinchiroGameTurn(turn), nil
}

func (r *chinchiroGameTurnParticipantRollResolver) participant(ctx context.Context, obj *gqlmodel.ChinchiroGameTurnParticipantRoll) (*gqlmodel.ChinchiroGameParticipant, error) {
	thunk := r.loaders.ChinchiroGameParticipantLoader.Load(ctx, dataloader.StringKey(obj.ParticipantID))
	p, err := thunk()
	if err != nil {
		return nil, err
	}
	turn := p.(*model.ChinchiroGameParticipant)
	return MapToChinchiroGameParticipant(turn), nil
}

func (r *mutationResolver) registerChinchiroGame(ctx context.Context, input gqlmodel.NewChinchiroGame) (*gqlmodel.RegisterChinchiroGamePayload, error) {
	user := auth.GetUser(ctx)
	if user == nil {
		return nil, fmt.Errorf("user not found")
	}
	roomID, err := idToUint32(input.RoomID)
	if err != nil {
		return nil, err
	}
	game, err := r.chinchiroGameUsecase.RegisterChinchiroGame(ctx, *user, roomID, model.ChinchiroGame{})
	if err != nil {
		return nil, err
	}
	return &gqlmodel.RegisterChinchiroGamePayload{
		ChinchiroGame: MapToChinchiroGame(game),
	}, nil
}

func (r *mutationResolver) updateChinchiroGameTurnStatus(ctx context.Context, input gqlmodel.UpdateChinchiroGameTurnStatus) (*gqlmodel.UpdateChinchiroGameTurnStatusPayload, error) {
	user := auth.GetUser(ctx)
	if user == nil {
		return nil, fmt.Errorf("user not found")
	}
	turnID, err := idToUint32(input.TurnID)
	if err != nil {
		return nil, err
	}
	err = r.chinchiroGameUsecase.UpdateChinchiroGameTurnStatus(ctx, *user, turnID, *model.ChinchiroGameTurnStatusValueOf(input.Status.String()))
	if err != nil {
		return nil, err
	}
	return &gqlmodel.UpdateChinchiroGameTurnStatusPayload{
		Ok: true,
	}, nil
}

func (r *mutationResolver) betChinchiroGameTurnParticipant(ctx context.Context, input gqlmodel.BetChinchiroGameTurnParticipant) (*gqlmodel.BetChinchiroGameTurnParticipantPayload, error) {
	user := auth.GetUser(ctx)
	if user == nil {
		return nil, fmt.Errorf("user not found")
	}
	turnID, err := idToUint32(input.TurnID)
	if err != nil {
		return nil, err
	}
	err = r.chinchiroGameUsecase.BetChinchiroGameTurn(ctx, *user, turnID, input.BetAmount)
	if err != nil {
		return nil, err
	}
	return &gqlmodel.BetChinchiroGameTurnParticipantPayload{
		Ok: true,
	}, nil
}

func (r *mutationResolver) rollChinchiroGameTurnParticipant(ctx context.Context, input gqlmodel.RollChinchiroGameTurnParticipant) (*gqlmodel.RollChinchiroGameTurnParticipantPayload, error) {
	user := auth.GetUser(ctx)
	if user == nil {
		return nil, fmt.Errorf("user not found")
	}
	turnID, err := idToUint32(input.TurnID)
	if err != nil {
		return nil, err
	}
	err = r.chinchiroGameUsecase.RollChinchiroGameTurn(ctx, *user, turnID)
	if err != nil {
		return nil, err
	}
	return &gqlmodel.RollChinchiroGameTurnParticipantPayload{
		Ok: true,
	}, nil
}

func (r *queryResolver) chinchiroGames(_ context.Context, query gqlmodel.ChinchiroGamesQuery) ([]*gqlmodel.ChinchiroGame, error) {
	queryIDs, err := idsToUint32s(query.Ids)
	if err != nil {
		return nil, err
	}
	roomID, err := idToUint32(*query.RoomID)
	if err != nil {
		return nil, err
	}
	var statuses *[]model.ChinchiroGameStatus
	if query.Statuses != nil {
		ses := array.Map(query.Statuses, func(s gqlmodel.ChinchiroGameStatus) model.ChinchiroGameStatus {
			return *model.ChinchiroGameStatusValueOf(s.String())
		})
		statuses = &ses
	}

	games, err := r.chinchiroGameUsecase.FindChinchiroGames(model.ChinchiroGamesQuery{
		IDs:      &queryIDs,
		RoomID:   &roomID,
		Name:     query.Name,
		Statuses: statuses,
		Paging:   query.Paging.MapToPagingQuery(),
	})
	if err != nil {
		return nil, err
	}
	return array.Map(games.List, func(g model.ChinchiroGame) *gqlmodel.ChinchiroGame {
		return MapToChinchiroGame(&g)
	}), nil
}

func (r *queryResolver) chinchiroGame(_ context.Context, gameID string) (*gqlmodel.ChinchiroGame, error) {
	gid, err := idToUint32(gameID)
	if err != nil {
		return nil, err
	}
	game, err := r.chinchiroGameUsecase.FindChinchiroGame(gid)
	if err != nil {
		return nil, err
	}
	return MapToChinchiroGame(game), nil
}

func (r *queryResolver) myChinchiroGameParticipant(ctx context.Context, gameID string) (*gqlmodel.ChinchiroGameParticipant, error) {
	user := auth.GetUser(ctx)
	if user == nil {
		return nil, nil
	}
	gid, err := idToUint32(gameID)
	if err != nil {
		return nil, err
	}
	participant, err := r.chinchiroGameUsecase.FindMyChinchiroGameParticipant(gid, *user)
	if err != nil {
		return nil, err
	}
	return MapToChinchiroGameParticipant(participant), nil
}

func (r *queryResolver) chinchiroGameTurns(_ context.Context, query gqlmodel.ChinchiroGameTurnsQuery) ([]*gqlmodel.ChinchiroGameTurn, error) {
	queryIDs, err := idsToUint32s(query.Ids)
	if err != nil {
		return nil, err
	}
	gameID, err := idToUint32(*query.GameID)
	if err != nil {
		return nil, err
	}
	var statuses *[]model.ChinchiroGameTurnStatus
	if query.Statuses != nil {
		ses := array.Map(query.Statuses, func(s gqlmodel.ChinchiroGameTurnStatus) model.ChinchiroGameTurnStatus {
			return *model.ChinchiroGameTurnStatusValueOf(s.String())
		})
		statuses = &ses
	}

	turns, err := r.chinchiroGameUsecase.FindChinchiroGameTurns(model.ChinchiroGameTurnsQuery{
		IDs:      &queryIDs,
		GameID:   &gameID,
		Statuses: statuses,
		Paging:   query.Paging.MapToPagingQuery(),
	})
	if err != nil {
		return nil, err
	}
	return array.Map(turns.List, func(t model.ChinchiroGameTurn) *gqlmodel.ChinchiroGameTurn {
		return MapToChinchiroGameTurn(&t)
	}), nil
}

func (r *queryResolver) chinchiroGameTurn(_ context.Context, turnID string) (*gqlmodel.ChinchiroGameTurn, error) {
	tid, err := idToUint32(turnID)
	if err != nil {
		return nil, err
	}
	turn, err := r.chinchiroGameUsecase.FindChinchiroGameTurn(tid)
	if err != nil {
		return nil, err
	}
	return MapToChinchiroGameTurn(turn), nil
}

func (r *queryResolver) chinchiroGameTurnRolls(_ context.Context, query *gqlmodel.ChinchiroGameTurnRollsQuery) ([]*gqlmodel.ChinchiroGameTurnParticipantRoll, error) {
	queryIDs, err := idsToUint32s(query.Ids)
	if err != nil {
		return nil, err
	}
	turnID, err := idToUint32(*query.TurnID)
	if err != nil {
		return nil, err
	}
	rolls, err := r.chinchiroGameUsecase.FindChinchiroGameTurnRolls(model.ChinchiroGameTurnRollsQuery{
		IDs:    &queryIDs,
		TurnID: &turnID,
		Paging: query.Paging.MapToPagingQuery(),
	})
	if err != nil {
		return nil, err
	}
	return array.Map(rolls.List, func(r model.ChinchiroGameTurnRoll) *gqlmodel.ChinchiroGameTurnParticipantRoll {
		return MapToChinchiroGameParticipantRoll(&r)
	}), nil
}

func (r *queryResolver) chinchiroGameTurnParticipantResults(_ context.Context, query gqlmodel.ChinchiroGameTurnParticipantResultsQuery) ([]*gqlmodel.ChinchiroGameTurnParticipantResult, error) {
	queryIDs, err := idsToUint32s(query.Ids)
	if err != nil {
		return nil, err
	}
	turnID, err := idToUint32(*query.TurnID)
	if err != nil {
		return nil, err
	}
	results, err := r.chinchiroGameUsecase.FindChinchiroGameTurnResults(model.ChinchiroGameTurnResultsQuery{
		IDs:    &queryIDs,
		TurnID: &turnID,
		Paging: query.Paging.MapToPagingQuery(),
	})
	if err != nil {
		return nil, err
	}
	return array.Map(results.List, func(r model.ChinchiroGameTurnResult) *gqlmodel.ChinchiroGameTurnParticipantResult {
		return MapToChinchiroGameParticipantResult(&r)
	}), nil
}
