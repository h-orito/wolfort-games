package usecase

import (
	"chat-role-play/application/app_service"
	"chat-role-play/domain/model"
	"context"
)

type ChinchiroGameUsecase interface {
	// game
	FindChinchiroGames(query model.ChinchiroGamesQuery) ([]model.ChinchiroGame, error)
	FindChinchiroGame(ID uint32) (*model.ChinchiroGame, error)
	RegisterChinchiroGame(ctx context.Context, user model.User, roomID uint32, game model.ChinchiroGame) (*model.ChinchiroGame, error)
	UpdateChinchiroGameStatus(ctx context.Context, user model.User, gameID uint32, status model.ChinchiroGameStatus) error
	// game participant
	FindChinchiroGameParticipants(query model.ChinchiroGameParticipantsQuery) (model.ChinchiroGameParticipants, error)
	FindMyChinchiroGameParticipant(gameID uint32, user model.User) (*model.ChinchiroGameParticipant, error)
	RegisterChinchiroGameParticipant(ctx context.Context, user model.User, gameID uint32, participant model.ChinchiroGameParticipant) (*model.ChinchiroGameParticipant, error)
	DeleteChinchiroGameParticipant(ctx context.Context, user model.User, participantID uint32) error
	// game turn
	FindChinchiroGameTurns(query model.ChinchiroGameTurnsQuery) (model.ChinchiroGameTurns, error)
	FindChinchiroGameTurn(ID uint32) (*model.ChinchiroGameTurn, error)
	RegisterChinchiroGameTurn(ctx context.Context, user model.User, gameID uint32) (*model.ChinchiroGameTurn, error)
	UpdateChinchiroGameTurnStatus(ctx context.Context, user model.User, turnID uint32, status model.ChinchiroGameTurnStatus) error
	// game turn roll
	FindChinchiroGameTurnRolls(query model.ChinchiroGameTurnRollsQuery) (model.ChinchiroGameTurnRolls, error)
	FindChinchiroGameTurnResults(query model.ChinchiroGameTurnResultsQuery) (model.ChinchiroGameTurnResults, error)
	BetChinchiroGameTurn(ctx context.Context, user model.User, turnID uint32, betAmount int) error
	RollChinchiroGameTurn(ctx context.Context, user model.User, turnID uint32) error
}

type chinchiroGameUsecase struct {
	chinchiroGameService app_service.ChinchiroGameService
}

// DeleteChinchiroGameParticipant implements ChinchiroGameUsecase.
func (*chinchiroGameUsecase) DeleteChinchiroGameParticipant(ctx context.Context, user model.User, participantID uint32) error {
	panic("unimplemented")
}

// RegisterChinchiroGameParticipant implements ChinchiroGameUsecase.
func (*chinchiroGameUsecase) RegisterChinchiroGameParticipant(ctx context.Context, user model.User, gameID uint32, participant model.ChinchiroGameParticipant) (*model.ChinchiroGameParticipant, error) {
	panic("unimplemented")
}

// FindChinchiroGameParticipants implements ChinchiroGameUsecase.
func (*chinchiroGameUsecase) FindChinchiroGameParticipants(query model.ChinchiroGameParticipantsQuery) (model.ChinchiroGameParticipants, error) {
	panic("unimplemented")
}

// BetChinchiroGameTurn implements ChinchiroGameUsecase.
func (*chinchiroGameUsecase) BetChinchiroGameTurn(ctx context.Context, user model.User, turnID uint32, betAmount int) error {
	panic("unimplemented")
}

// FindChinchiroGame implements ChinchiroGameUsecase.
func (*chinchiroGameUsecase) FindChinchiroGame(ID uint32) (*model.ChinchiroGame, error) {
	panic("unimplemented")
}

// FindChinchiroGameTurn implements ChinchiroGameUsecase.
func (*chinchiroGameUsecase) FindChinchiroGameTurn(ID uint32) (*model.ChinchiroGameTurn, error) {
	panic("unimplemented")
}

// FindChinchiroGameTurnResults implements ChinchiroGameUsecase.
func (*chinchiroGameUsecase) FindChinchiroGameTurnResults(query model.ChinchiroGameTurnResultsQuery) (model.ChinchiroGameTurnResults, error) {
	panic("unimplemented")
}

// FindChinchiroGameTurnRolls implements ChinchiroGameUsecase.
func (*chinchiroGameUsecase) FindChinchiroGameTurnRolls(query model.ChinchiroGameTurnRollsQuery) (model.ChinchiroGameTurnRolls, error) {
	panic("unimplemented")
}

// FindChinchiroGameTurns implements ChinchiroGameUsecase.
func (*chinchiroGameUsecase) FindChinchiroGameTurns(query model.ChinchiroGameTurnsQuery) (model.ChinchiroGameTurns, error) {
	panic("unimplemented")
}

// FindChinchiroGames implements ChinchiroGameUsecase.
func (*chinchiroGameUsecase) FindChinchiroGames(query model.ChinchiroGamesQuery) ([]model.ChinchiroGame, error) {
	panic("unimplemented")
}

// FindMyChinchiroGameParticipant implements ChinchiroGameUsecase.
func (*chinchiroGameUsecase) FindMyChinchiroGameParticipant(gameID uint32, user model.User) (*model.ChinchiroGameParticipant, error) {
	panic("unimplemented")
}

// RegisterChinchiroGame implements ChinchiroGameUsecase.
func (*chinchiroGameUsecase) RegisterChinchiroGame(ctx context.Context, user model.User, roomID uint32, game model.ChinchiroGame) (*model.ChinchiroGame, error) {
	panic("unimplemented")
}

// RegisterChinchiroGameTurn implements ChinchiroGameUsecase.
func (*chinchiroGameUsecase) RegisterChinchiroGameTurn(ctx context.Context, user model.User, gameID uint32) (*model.ChinchiroGameTurn, error) {
	panic("unimplemented")
}

// RollChinchiroGameTurn implements ChinchiroGameUsecase.
func (*chinchiroGameUsecase) RollChinchiroGameTurn(ctx context.Context, user model.User, turnID uint32) error {
	panic("unimplemented")
}

// UpdateChinchiroGameStatus implements ChinchiroGameUsecase.
func (*chinchiroGameUsecase) UpdateChinchiroGameStatus(ctx context.Context, user model.User, gameID uint32, status model.ChinchiroGameStatus) error {
	panic("unimplemented")
}

// UpdateChinchiroGameTurnStatus implements ChinchiroGameUsecase.
func (*chinchiroGameUsecase) UpdateChinchiroGameTurnStatus(ctx context.Context, user model.User, turnID uint32, status model.ChinchiroGameTurnStatus) error {
	panic("unimplemented")
}

// NewChinchiroGameUsecase creates a new ChinchiroGameUsecase.
func NewChinchiroGameUsecase(chinchiroGameService app_service.ChinchiroGameService) ChinchiroGameUsecase {
	return &chinchiroGameUsecase{
		chinchiroGameService: chinchiroGameService,
	}
}
