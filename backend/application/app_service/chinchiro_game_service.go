package app_service

import (
	"chat-role-play/domain/model"
	"context"
)

type ChinchiroGameService interface {
	// game
	FindChinchiroGames(query model.ChinchiroGamesQuery) ([]model.ChinchiroGame, error)
	FindChinchiroGame(ID uint32) (*model.ChinchiroGame, error)
	RegisterChinchiroGame(ctx context.Context, roomID uint32, game model.ChinchiroGame) (*model.ChinchiroGame, error)
	UpdateChinchiroGameStatus(ctx context.Context, gameID uint32, game model.ChinchiroGame) (*model.ChinchiroGame, error)
	// game participant
	FindChinchiroGameParticipants(query model.ChinchiroGameParticipantsQuery) (model.ChinchiroGameParticipants, error)
	FindMyChinchiroGameParticipant(gameID uint32, roomParticipantID uint32) (*model.ChinchiroGameParticipant, error)
	// game turn
	FindChinchiroGameTurns(query model.ChinchiroGameTurnsQuery) (model.ChinchiroGameTurns, error)
	FindChinchiroGameTurn(ID uint32) (*model.ChinchiroGameTurn, error)
	RegisterChinchiroGameTurn(ctx context.Context, gameID uint32, dealerID uint32) (*model.ChinchiroGameTurn, error)
	UpdateChinchiroGameTurn(ctx context.Context, turnID uint32, status model.ChinchiroGameStatus) (*model.ChinchiroGameTurn, error)
	// game turn roll
	FindChinchiroGameTurnRolls(query model.ChinchiroGameTurnRollsQuery) (model.ChinchiroGameTurnRolls, error)
	RegisterChinchiroGameTurnBet(ctx context.Context, turnID uint32, participantID uint32, betAmount int) error
	RegisterChinchiroGameTurnRoll(ctx context.Context, turnID uint32, participantID uint32, diceRoll model.ChinchiroDiceRoll) error
	// game turn result
	FindChinchiroGameTurnResults(query model.ChinchiroGameTurnResultsQuery) (model.ChinchiroGameTurnResults, error)
}

type chinchiroGameService struct {
	chinchiroGameRepository model.ChinchiroGameRepository
}

// FindChinchiroGameParticipants implements ChinchiroGameService.
func (*chinchiroGameService) FindChinchiroGameParticipants(query model.ChinchiroGameParticipantsQuery) (model.ChinchiroGameParticipants, error) {
	panic("unimplemented")
}

// FindMyChinchiroGameParticipant implements ChinchiroGameService.
func (*chinchiroGameService) FindMyChinchiroGameParticipant(gameID uint32, roomParticipantID uint32) (*model.ChinchiroGameParticipant, error) {
	panic("unimplemented")
}

// FindChinchiroGame implements ChinchiroGameService.
func (*chinchiroGameService) FindChinchiroGame(ID uint32) (*model.ChinchiroGame, error) {
	panic("unimplemented")
}

// FindChinchiroGameTurn implements ChinchiroGameService.
func (*chinchiroGameService) FindChinchiroGameTurn(ID uint32) (*model.ChinchiroGameTurn, error) {
	panic("unimplemented")
}

// FindChinchiroGameTurnResults implements ChinchiroGameService.
func (*chinchiroGameService) FindChinchiroGameTurnResults(query model.ChinchiroGameTurnResultsQuery) (model.ChinchiroGameTurnResults, error) {
	panic("unimplemented")
}

// FindChinchiroGameTurnRolls implements ChinchiroGameService.
func (*chinchiroGameService) FindChinchiroGameTurnRolls(query model.ChinchiroGameTurnRollsQuery) (model.ChinchiroGameTurnRolls, error) {
	panic("unimplemented")
}

// FindChinchiroGameTurns implements ChinchiroGameService.
func (*chinchiroGameService) FindChinchiroGameTurns(query model.ChinchiroGameTurnsQuery) (model.ChinchiroGameTurns, error) {
	panic("unimplemented")
}

// FindChinchiroGames implements ChinchiroGameService.
func (*chinchiroGameService) FindChinchiroGames(query model.ChinchiroGamesQuery) ([]model.ChinchiroGame, error) {
	panic("unimplemented")
}

// RegisterChinchiroGame implements ChinchiroGameService.
func (*chinchiroGameService) RegisterChinchiroGame(ctx context.Context, roomID uint32, game model.ChinchiroGame) (*model.ChinchiroGame, error) {
	panic("unimplemented")
}

// RegisterChinchiroGameTurn implements ChinchiroGameService.
func (*chinchiroGameService) RegisterChinchiroGameTurn(ctx context.Context, gameID uint32, dealerID uint32) (*model.ChinchiroGameTurn, error) {
	panic("unimplemented")
}

// RegisterChinchiroGameTurnBet implements ChinchiroGameService.
func (*chinchiroGameService) RegisterChinchiroGameTurnBet(ctx context.Context, turnID uint32, participantID uint32, betAmount int) error {
	panic("unimplemented")
}

// RegisterChinchiroGameTurnRoll implements ChinchiroGameService.
func (*chinchiroGameService) RegisterChinchiroGameTurnRoll(ctx context.Context, turnID uint32, participantID uint32, diceRoll model.ChinchiroDiceRoll) error {
	panic("unimplemented")
}

// UpdateChinchiroGameStatus implements ChinchiroGameService.
func (*chinchiroGameService) UpdateChinchiroGameStatus(ctx context.Context, gameID uint32, game model.ChinchiroGame) (*model.ChinchiroGame, error) {
	panic("unimplemented")
}

// UpdateChinchiroGameTurn implements ChinchiroGameService.
func (*chinchiroGameService) UpdateChinchiroGameTurn(ctx context.Context, turnID uint32, status model.ChinchiroGameStatus) (*model.ChinchiroGameTurn, error) {
	panic("unimplemented")
}

func NewChinchiroGameService(
	chinchiroGameRepository model.ChinchiroGameRepository,
) ChinchiroGameService {
	return &chinchiroGameService{
		chinchiroGameRepository: chinchiroGameRepository,
	}
}
