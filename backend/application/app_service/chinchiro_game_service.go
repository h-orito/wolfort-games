package app_service

import (
	"context"
	"fmt"
	"wolfort-games/domain/dom_service"
	"wolfort-games/domain/model"
	"wolfort-games/util/array"
)

type ChinchiroGameService interface {
	// game
	FindChinchiroGames(query model.ChinchiroGamesQuery) (model.ChinchiroGames, error)
	FindChinchiroGame(ID uint32) (*model.ChinchiroGame, error)
	RegisterChinchiroGame(ctx context.Context, roomID uint32, game model.ChinchiroGame) (*model.ChinchiroGame, error)
	UpdateChinchiroGameStatus(ctx context.Context, gameID uint32, status model.ChinchiroGameStatus) error
	// game participant
	FindChinchiroGameParticipants(query model.ChinchiroGameParticipantsQuery) (model.ChinchiroGameParticipants, error)
	FindChinchiroGameParticipant(query model.ChinchiroGameParticipantQuery) (*model.ChinchiroGameParticipant, error)
	// game turn
	FindChinchiroGameTurns(query model.ChinchiroGameTurnsQuery) (model.ChinchiroGameTurns, error)
	FindChinchiroGameTurn(ID uint32) (*model.ChinchiroGameTurn, error)
	UpdateChinchiroGameTurn(ctx context.Context, turnID uint32, status model.ChinchiroGameTurnStatus) error
	// game turn roll
	FindChinchiroGameTurnRolls(query model.ChinchiroGameTurnRollsQuery) (model.ChinchiroGameTurnRolls, error)
	RegisterChinchiroGameTurnBet(ctx context.Context, turnID uint32, participantID uint32, betAmount int) error
	RegisterChinchiroGameTurnRoll(ctx context.Context, turnID uint32, participantID uint32, diceRoll model.ChinchiroDiceRoll) error
	// game turn result
	FindChinchiroGameTurnResults(query model.ChinchiroGameTurnResultsQuery) (model.ChinchiroGameTurnResults, error)
}

type chinchiroGameService struct {
	chinchiroRoomParticipantRepository model.ChinchiroRoomParticipantRepository
	chinchiroGameRepository            model.ChinchiroGameRepository
	chinchiroGameDomainService         dom_service.ChinchiroGameDomainService
}

func (s *chinchiroGameService) FindChinchiroGames(query model.ChinchiroGamesQuery) (model.ChinchiroGames, error) {
	return s.chinchiroGameRepository.FindGames(query)
}

func (s *chinchiroGameService) FindChinchiroGame(ID uint32) (*model.ChinchiroGame, error) {
	return s.chinchiroGameRepository.FindGame(ID)
}

func (s *chinchiroGameService) RegisterChinchiroGame(ctx context.Context, roomID uint32, game model.ChinchiroGame) (*model.ChinchiroGame, error) {
	g, err := s.chinchiroGameRepository.RegisterGame(ctx, roomID, game)
	if err != nil {
		return nil, err
	}
	// 参加者を登録
	gameParticipants, err := s.registerGameParticipants(ctx, g.ID, roomID)
	if err != nil {
		return nil, err
	}

	// 最初のターンを開始
	_, err = s.registerChinchiroGameTurn(ctx, g.ID, gameParticipants[0].ID)
	if err != nil {
		return nil, err
	}

	return g, err
}

func (s *chinchiroGameService) UpdateChinchiroGameStatus(ctx context.Context, gameID uint32, status model.ChinchiroGameStatus) error {
	return s.chinchiroGameRepository.UpdateGameStatus(ctx, gameID, status)
}

func (s *chinchiroGameService) FindChinchiroGameParticipants(query model.ChinchiroGameParticipantsQuery) (model.ChinchiroGameParticipants, error) {
	return s.chinchiroGameRepository.FindGameParticipants(query)
}

func (s *chinchiroGameService) FindChinchiroGameParticipant(query model.ChinchiroGameParticipantQuery) (*model.ChinchiroGameParticipant, error) {
	return s.chinchiroGameRepository.FindGameParticipant(query)
}

func (s *chinchiroGameService) FindChinchiroGameTurns(query model.ChinchiroGameTurnsQuery) (model.ChinchiroGameTurns, error) {
	return s.chinchiroGameRepository.FindGameTurns(query)
}

func (s *chinchiroGameService) FindChinchiroGameTurn(ID uint32) (*model.ChinchiroGameTurn, error) {
	return s.chinchiroGameRepository.FindGameTurn(ID)
}

func (s *chinchiroGameService) UpdateChinchiroGameTurn(ctx context.Context, turnID uint32, status model.ChinchiroGameTurnStatus) error {
	return s.chinchiroGameRepository.UpdateGameTurnStatus(ctx, turnID, status)
}

func (s *chinchiroGameService) FindChinchiroGameTurnRolls(query model.ChinchiroGameTurnRollsQuery) (model.ChinchiroGameTurnRolls, error) {
	return s.chinchiroGameRepository.FindGameTurnRolls(query)
}

func (s *chinchiroGameService) RegisterChinchiroGameTurnBet(ctx context.Context, turnID uint32, participantID uint32, betAmount int) error {
	turn, err := s.assertBet(turnID, participantID)
	if err != nil {
		return err
	}
	// 掛け金を登録
	_, err = s.chinchiroGameRepository.RegisterGameTurnResult(ctx, turnID, participantID, model.ChinchiroGameTurnResult{
		TurnID:        turnID,
		ParticipantID: participantID,
		BetAmount:     betAmount,
	})
	if err != nil {
		return err
	}
	// 全員が賭けたらロール開始
	return s.updateGameTurnRollingIfNeeded(ctx, turn)
}

func (s *chinchiroGameService) RegisterChinchiroGameTurnRoll(ctx context.Context, turnID uint32, participantID uint32, diceRoll model.ChinchiroDiceRoll) error {
	turn, err := s.assertRoll(turnID, participantID)
	if err != nil {
		return err
	}
	// ロール結果を登録
	rolls, err := s.chinchiroGameRepository.FindGameTurnRolls(model.ChinchiroGameTurnRollsQuery{
		TurnID:        &turnID,
		ParticipantID: &participantID,
	})
	if err != nil {
		return err
	}

	_, err = s.chinchiroGameRepository.RegisterGameTurnRoll(ctx, turnID, participantID, model.ChinchiroGameTurnRoll{
		TurnID:        turnID,
		ParticipantID: participantID,
		RollNumber:    rolls.Count + 1,
		DiceRoll:      diceRoll,
	})
	if err != nil {
		return err
	}

	// 役が確定したら結果を登録
	diceRolls := array.Map(rolls.List, func(r model.ChinchiroGameTurnRoll) model.ChinchiroDiceRoll {
		return r.DiceRoll
	})
	diceRolls = append(diceRolls, diceRoll)
	return s.registerChinchiroGameTurnResultIfNeeded(ctx, *turn, participantID, diceRolls)
}

func (s *chinchiroGameService) FindChinchiroGameTurnResults(query model.ChinchiroGameTurnResultsQuery) (model.ChinchiroGameTurnResults, error) {
	return s.chinchiroGameRepository.FindGameTurnResults(query)
}

func NewChinchiroGameService(
	chinchiroRoomParticipantRepository model.ChinchiroRoomParticipantRepository,
	chinchiroGameRepository model.ChinchiroGameRepository,
	chinchiroGameDomainService dom_service.ChinchiroGameDomainService,
) ChinchiroGameService {
	return &chinchiroGameService{
		chinchiroRoomParticipantRepository: chinchiroRoomParticipantRepository,
		chinchiroGameRepository:            chinchiroGameRepository,
		chinchiroGameDomainService:         chinchiroGameDomainService,
	}
}

// -----------------------

func (s *chinchiroGameService) assertBet(
	turnID uint32,
	participantID uint32,
) (*model.ChinchiroGameTurn, error) {
	turn, err := s.chinchiroGameRepository.FindGameTurn(turnID)
	if err != nil {
		return nil, err
	}
	if turn.Status != model.ChinchiroGameTurnStatusBetting {
		return nil, fmt.Errorf("turn is not betting")
	}
	if turn.DealerID == participantID {
		return nil, fmt.Errorf("dealer cannot bet")
	}
	result, err := s.chinchiroGameRepository.FindGameTurnResult(model.ChinchiroGameTurnResultQuery{
		TurnID:        &turnID,
		ParticipantID: &participantID,
	})
	if err != nil {
		return nil, err
	}
	if result != nil {
		return nil, fmt.Errorf("result already exists")
	}
	return turn, nil
}

func (s *chinchiroGameService) assertRoll(
	turnID uint32,
	participantID uint32,
) (*model.ChinchiroGameTurn, error) {
	turn, err := s.chinchiroGameRepository.FindGameTurn(turnID)
	if err != nil {
		return nil, err
	}
	if turn.Status != model.ChinchiroGameTurnStatusRolling {
		return nil, fmt.Errorf("turn is not rolling")
	}

	if turn.NextRollerID != nil && *turn.NextRollerID != participantID {
		return nil, fmt.Errorf("not your turn")
	}

	// 既にロール結果が登録されていたらNG
	result, err := s.chinchiroGameRepository.FindGameTurnResult(model.ChinchiroGameTurnResultQuery{
		TurnID:        &turnID,
		ParticipantID: &participantID,
	})
	if err != nil {
		return nil, err
	}
	if result != nil && result.DiceRoll != nil {
		return nil, fmt.Errorf("result already exists")
	}

	return turn, nil
}

func (s *chinchiroGameService) registerGameParticipants(
	ctx context.Context,
	gameID uint32,
	roomID uint32,
) ([]model.ChinchiroGameParticipant, error) {
	isExcludeGone := true
	roomParticipants, err := s.chinchiroRoomParticipantRepository.FindRoomParticipants(model.ChinchiroRoomParticipantsQuery{
		RoomIDs:       &[]uint32{roomID},
		IsExcludeGone: &isExcludeGone,
	})
	if err != nil {
		return nil, err
	}
	var gameParticipants []model.ChinchiroGameParticipant
	for index, p := range array.Shuffle(roomParticipants.List) {
		gp, err := s.chinchiroGameRepository.RegisterGameParticipant(ctx, gameID, model.ChinchiroGameParticipant{
			RoomParticipantID: p.ID,
			Balance:           model.ChinchiroGameParticipantInitialBalance,
			TurnOrder:         index + 1,
		})
		if err != nil {
			return nil, err
		}
		gameParticipants = append(gameParticipants, *gp)
	}
	return gameParticipants, nil
}

func (s *chinchiroGameService) registerChinchiroGameTurn(ctx context.Context, gameID uint32, dealerID uint32) (*model.ChinchiroGameTurn, error) {
	turns, err := s.chinchiroGameRepository.FindGameTurns(model.ChinchiroGameTurnsQuery{
		GameID: &gameID,
	})
	if err != nil {
		return nil, err
	}
	var maxTurnNumber int = 0
	if max := array.MaxOrNil(turns.List, func(t model.ChinchiroGameTurn) int {
		return t.TurnNumber
	}); max != nil {
		maxTurnNumber = *max
	}
	return s.chinchiroGameRepository.RegisterGameTurn(ctx, gameID, model.ChinchiroGameTurn{
		ID:         gameID,
		DealerID:   dealerID,
		Status:     model.ChinchiroGameTurnStatusBetting,
		TurnNumber: maxTurnNumber + 1,
	})
}

func (s *chinchiroGameService) updateGameTurnRollingIfNeeded(
	ctx context.Context,
	turn *model.ChinchiroGameTurn,
) error {
	participants, err := s.chinchiroGameRepository.FindGameParticipants(model.ChinchiroGameParticipantsQuery{
		GameID: &turn.GameID,
	})
	if err != nil {
		return err
	}
	results, err := s.chinchiroGameRepository.FindGameTurnResults(model.ChinchiroGameTurnResultsQuery{
		TurnID: &turn.ID,
	})
	if err != nil {
		return err
	}
	if s.chinchiroGameDomainService.ShouldTurnRolling(*turn, participants, results) {
		// 全員が賭けたのでロール開始
		err = s.chinchiroGameRepository.UpdateGameTurnStatus(ctx, turn.ID, model.ChinchiroGameTurnStatusRolling)
		if err != nil {
			return err
		}
		// 親が最初に振る
		return s.chinchiroGameRepository.UpdateGameTurnRoller(ctx, turn.ID, &turn.DealerID)
	}
	return nil
}

// TODO: update next roller id
func (s *chinchiroGameService) registerChinchiroGameTurnResultIfNeeded(
	ctx context.Context,
	turn model.ChinchiroGameTurn,
	participantID uint32,
	rolls []model.ChinchiroDiceRoll,
) error {
	// 役が確定したら結果を登録
	combination := model.DetermineDiceCombination(rolls)
	if combination == nil {
		return nil // 役が確定していない場合はロール結果だけ登録して終了（3回振って目なしの場合はここには来ない）
	}

	turnID := turn.ID
	if turn.DealerID == participantID { // 振ったのが親の場合
		// 子も必ず振るので、親は結果を登録するだけで終わり
		err := s.registerGameTurnResult(ctx, turnID, participantID, 0)
		if err != nil {
			return err
		}
		// 次振るべき人を更新
		_, err = s.updateGameTurnNextRoller(ctx, turn)
		return err
	}

	// 以降は子の場合
	// 親の目と役
	dealerRolls, dealerCombination, err := s.findDealerRollsAndCombination(turn)
	if err != nil {
		return err
	}

	// 自身の掛け金を取得
	result, err := s.chinchiroGameRepository.FindGameTurnResult(model.ChinchiroGameTurnResultQuery{
		TurnID:        &turnID,
		ParticipantID: &participantID,
	})
	if err != nil {
		return err
	}
	if result == nil {
		return fmt.Errorf("result not found")
	}

	// 子の獲得額
	winnings := model.CalculateWinnings(*combination, *dealerCombination, result.BetAmount)
	// 子の結果を登録
	_, err = s.chinchiroGameRepository.UpdateGameTurnResult(ctx, result.ID, model.ChinchiroGameTurnResult{
		DiceRoll: &rolls[len(rolls)-1],
		Winnings: &winnings,
	})
	if err != nil {
		return err
	}

	// 親の結果を更新
	err = s.updateDealerResult(ctx, turn, dealerRolls, winnings)
	if err != nil {
		return err
	}

	// 次振るべき人を更新
	results, err := s.updateGameTurnNextRoller(ctx, turn)
	if err != nil {
		return err
	}

	// 子が全員振ったらターン終了
	return s.updateTurnFinishedIfNeeded(ctx, turn, *results)
}

// 次振るべき人を更新
func (s *chinchiroGameService) updateGameTurnNextRoller(
	ctx context.Context,
	turn model.ChinchiroGameTurn,
) (*model.ChinchiroGameTurnResults, error) {
	participants, err := s.chinchiroGameRepository.FindGameParticipants(model.ChinchiroGameParticipantsQuery{
		GameID: &turn.GameID,
	})
	if err != nil {
		return nil, err
	}
	results, err := s.chinchiroGameRepository.FindGameTurnResults(model.ChinchiroGameTurnResultsQuery{
		TurnID: &turn.ID,
	})
	if err != nil {
		return nil, err
	}
	roller := s.chinchiroGameDomainService.DetermineNextRoller(turn, participants, results)
	var rollerID *uint32
	if roller != nil {
		rollerID = &roller.ID
	}
	return &results, s.chinchiroGameRepository.UpdateGameTurnRoller(ctx, turn.ID, rollerID)
}

func (s *chinchiroGameService) registerGameTurnResult(
	ctx context.Context,
	turnID uint32,
	participantID uint32,
	betAmount int,
) error {
	_, err := s.chinchiroGameRepository.RegisterGameTurnResult(ctx, turnID, participantID, model.ChinchiroGameTurnResult{
		TurnID:        turnID,
		ParticipantID: participantID,
		BetAmount:     betAmount,
	})
	if err != nil {
		return err
	}
	return nil
}

func (s *chinchiroGameService) updateDealerResult(
	ctx context.Context,
	turn model.ChinchiroGameTurn,
	dealerRolls model.ChinchiroGameTurnRolls,
	winnings int,
) error {
	dealerResult, err := s.chinchiroGameRepository.FindGameTurnResult(model.ChinchiroGameTurnResultQuery{
		TurnID:        &turn.ID,
		ParticipantID: &turn.DealerID,
	})
	if err != nil {
		return err
	}
	if dealerResult == nil {
		return fmt.Errorf("dealer result not found")
	}
	var dealerWinnings int = -winnings
	if dealerResult.Winnings != nil {
		dealerWinnings += *dealerResult.Winnings
	}
	// 親の結果を登録
	s.chinchiroGameRepository.UpdateGameTurnResult(ctx, dealerResult.ID, model.ChinchiroGameTurnResult{
		DiceRoll: &dealerRolls.List[dealerRolls.Count-1].DiceRoll,
		Winnings: &dealerWinnings,
	})
	return nil
}

// 子が全員振ったらターン終了
func (s *chinchiroGameService) updateTurnFinishedIfNeeded(
	ctx context.Context,
	turn model.ChinchiroGameTurn,
	results model.ChinchiroGameTurnResults,
) error {
	if !results.IsRollFinished() {
		// まだ振っていない人がいる
		return nil
	}
	// ターン終了
	err := s.UpdateChinchiroGameTurn(ctx, turn.ID, model.ChinchiroGameTurnStatusFinished)
	if err != nil {
		return err
	}
	// 次のターン開始 or ゲーム終了
	return s.registerNextTurnOrFinishGame(ctx, turn)
}

// 親の目と役を取得
func (s *chinchiroGameService) findDealerRollsAndCombination(
	turn model.ChinchiroGameTurn,
) (model.ChinchiroGameTurnRolls, *model.ChinchiroCombination, error) {
	// 親の目を取得
	dealerRolls, err := s.chinchiroGameRepository.FindGameTurnRolls(model.ChinchiroGameTurnRollsQuery{
		TurnID:        &turn.ID,
		ParticipantID: &turn.DealerID,
	})
	if err != nil {
		return model.ChinchiroGameTurnRolls{}, nil, err
	}
	// 親の役
	dealerCombination := model.DetermineDiceCombination(array.Map(dealerRolls.List, func(r model.ChinchiroGameTurnRoll) model.ChinchiroDiceRoll {
		return r.DiceRoll
	}))

	return dealerRolls, dealerCombination, nil
}

// 次のターン開始 or ゲーム終了
func (s *chinchiroGameService) registerNextTurnOrFinishGame(
	ctx context.Context,
	turn model.ChinchiroGameTurn,
) error {
	turns, err := s.chinchiroGameRepository.FindGameTurns(model.ChinchiroGameTurnsQuery{
		GameID: &turn.GameID,
	})
	if err != nil {
		return err
	}
	participants, err := s.chinchiroGameRepository.FindGameParticipants(model.ChinchiroGameParticipantsQuery{
		GameID: &turn.GameID,
	})
	if err != nil {
		return err
	}
	// 全員が2回親をやったらゲーム終了
	if s.chinchiroGameDomainService.ShouldGameFinished(participants, turns) {
		// ゲーム終了
		return s.UpdateChinchiroGameStatus(ctx, turn.GameID, model.ChinchiroGameStatusFinished)
	} else {
		// 次のターンを開始させる
		return s.registerNextTurn(ctx, turn, participants, turns)
	}
}

// 次のターンを開始させる
func (s *chinchiroGameService) registerNextTurn(
	ctx context.Context,
	turn model.ChinchiroGameTurn, // 現在のターン
	participants model.ChinchiroGameParticipants,
	turns model.ChinchiroGameTurns,
) error {
	dealerParticipant := array.Find(participants.List, func(p model.ChinchiroGameParticipant) bool {
		return p.ID == turn.DealerID
	})
	if dealerParticipant == nil {
		return fmt.Errorf("dealer participant not found")
	}
	// 次の親
	nextDealerId := s.chinchiroGameDomainService.DetermineNextDealerID(participants, turns, *dealerParticipant)
	// 次のターンを開始
	_, err := s.registerChinchiroGameTurn(ctx, turn.GameID, nextDealerId)
	return err
}
