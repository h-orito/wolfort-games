package graphql

import (
	"chat-role-play/adaptor/auth"
	"chat-role-play/domain/model"
	"chat-role-play/middleware/graph/gqlmodel"
	"chat-role-play/util/array"
	"context"
	"fmt"

	"github.com/graph-gophers/dataloader"
)

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

func (r *mutationResolver) registerChinchiroGameParticipant(ctx context.Context, input gqlmodel.NewChinchiroGameParticipant) (*gqlmodel.RegisterChinchiroGameParticipantPayload, error) {
	user := auth.GetUser(ctx)
	if user == nil {
		return nil, fmt.Errorf("user not found")
	}
	gameID, err := idToUint32(input.GameID)
	if err != nil {
		return nil, err
	}
	participant, err := r.chinchiroGameUsecase.RegisterChinchiroGameParticipant(ctx, *user, gameID, model.ChinchiroGameParticipant{})
	if err != nil {
		return nil, err
	}
	return &gqlmodel.RegisterChinchiroGameParticipantPayload{
		ChinchiroGameParticipant: MapToChinchiroGameParticipant(participant),
	}, nil
}

func (r *mutationResolver) deleteChinchiroGameParticipant(ctx context.Context, input gqlmodel.DeleteChinchiroGameParticipant) (*gqlmodel.DeleteChinchiroGameParticipantPayload, error) {
	user := auth.GetUser(ctx)
	if user == nil {
		return nil, fmt.Errorf("user not found")
	}
	participantID, err := idToUint32(input.ParticipantID)
	if err != nil {
		return nil, err
	}
	err = r.chinchiroGameUsecase.DeleteChinchiroGameParticipant(ctx, *user, participantID)
	if err != nil {
		return nil, err
	}
	return &gqlmodel.DeleteChinchiroGameParticipantPayload{
		Ok: true,
	}, nil
}

func (r *mutationResolver) registerChinchiroGameTurn(ctx context.Context, input gqlmodel.NewChinchiroGameTurn) (*gqlmodel.RegisterChinchiroGameTurnPayload, error) {
	user := auth.GetUser(ctx)
	if user == nil {
		return nil, fmt.Errorf("user not found")
	}
	gameID, err := idToUint32(input.GameID)
	if err != nil {
		return nil, err
	}
	turn, err := r.chinchiroGameUsecase.RegisterChinchiroGameTurn(ctx, *user, gameID)
	if err != nil {
		return nil, err
	}
	return &gqlmodel.RegisterChinchiroGameTurnPayload{
		ChinchiroGameTurn: MapToChinchiroGameTurn(turn),
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

// ChinchiroGames is the resolver for the chinchiroGames field.
func (r *queryResolver) chinchiroGames(ctx context.Context, query gqlmodel.ChinchiroGamesQuery) ([]*gqlmodel.ChinchiroGame, error) {
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
	return array.Map(games, func(g model.ChinchiroGame) *gqlmodel.ChinchiroGame {
		return MapToChinchiroGame(&g)
	}), nil
}

// ChinchiroGame is the resolver for the chinchiroGame field.
func (r *queryResolver) chinchiroGame(ctx context.Context, gameID string) (*gqlmodel.ChinchiroGame, error) {
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

// MyChinchiroGameParticipant is the resolver for the myChinchiroGameParticipant field.
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

// ChinchiroGameTurns is the resolver for the chinchiroGameTurns field.
func (r *queryResolver) chinchiroGameTurns(ctx context.Context, query gqlmodel.ChinchiroGameTurnsQuery) ([]*gqlmodel.ChinchiroGameTurn, error) {
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

// ChinchiroGameTurn is the resolver for the chinchiroGameTurn field.
func (r *queryResolver) chinchiroGameTurn(ctx context.Context, turnID string) (*gqlmodel.ChinchiroGameTurn, error) {
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

// ChinchiroGameTurnRolls is the resolver for the chinchiroGameTurnRolls field.
func (r *queryResolver) chinchiroGameTurnRolls(ctx context.Context, query *gqlmodel.ChinchiroGameTurnRollsQuery) ([]*gqlmodel.ChinchiroGameTurnParticipantRoll, error) {
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

// ChinchiroGameTurnParticipantResults is the resolver for the chinchiroGameTurnParticipantResults field.
func (r *queryResolver) chinchiroGameTurnParticipantResults(ctx context.Context, query gqlmodel.ChinchiroGameTurnParticipantResultsQuery) ([]*gqlmodel.ChinchiroGameTurnParticipantResult, error) {
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
