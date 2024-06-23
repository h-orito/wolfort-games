package db

import (
	"context"
	"errors"
	"fmt"
	model "wolfort-games/domain/model"
	"wolfort-games/util/array"

	"gorm.io/gorm"
)

type ChinchiroGameRepository struct {
	db *DB
}

func NewChinchiroGameRepository(db *DB) *ChinchiroGameRepository {
	return &ChinchiroGameRepository{
		db: db,
	}
}

func (repo *ChinchiroGameRepository) FindGames(query model.ChinchiroGamesQuery) (model.ChinchiroGames, error) {
	return findChinchiroGames(repo.db.Connection, query)
}

func (repo *ChinchiroGameRepository) FindGame(ID uint32) (*model.ChinchiroGame, error) {
	return findChinchiroGame(repo.db.Connection, ID)
}

func (repo *ChinchiroGameRepository) RegisterGame(ctx context.Context, roomID uint32, game model.ChinchiroGame) (saved *model.ChinchiroGame, err error) {
	tx, ok := GetTx(ctx)
	if !ok {
		return nil, fmt.Errorf("failed to get tx from context")
	}
	// game
	g, err := registerRdbChinchiroGame(tx, roomID, game)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	return findChinchiroGame(tx, g.ID)
}

func (repo *ChinchiroGameRepository) UpdateGameStatus(ctx context.Context, gameID uint32, status model.ChinchiroGameStatus) (err error) {
	tx, ok := GetTx(ctx)
	if !ok {
		return fmt.Errorf("failed to get tx from context")
	}
	result := tx.Model(&ChinchiroGame{}).Where("id = ?", gameID).Update("game_status_code", status.String())
	if result.Error != nil {
		return fmt.Errorf("failed to update: %s", result.Error)
	}
	return nil
}

func (repo *ChinchiroGameRepository) FindGameParticipants(query model.ChinchiroGameParticipantsQuery) (participants model.ChinchiroGameParticipants, err error) {
	return findChinchiroGameParticipants(repo.db.Connection, query)
}

func (repo *ChinchiroGameRepository) FindGameParticipant(query model.ChinchiroGameParticipantQuery) (participant *model.ChinchiroGameParticipant, err error) {
	return findChinchiroGameParticipant(repo.db.Connection, query)
}

func (repo *ChinchiroGameRepository) RegisterGameParticipant(ctx context.Context, gameID uint32, participant model.ChinchiroGameParticipant) (saved *model.ChinchiroGameParticipant, err error) {
	tx, ok := GetTx(ctx)
	if !ok {
		return nil, fmt.Errorf("failed to get tx from context")
	}
	// turn
	p, err := registerRdbChinchiroGameParticipant(tx, gameID, participant)
	if err != nil {
		return nil, err
	}
	return findChinchiroGameParticipant(tx, model.ChinchiroGameParticipantQuery{ID: &p.ID})
}

func (repo *ChinchiroGameRepository) FindGameTurns(query model.ChinchiroGameTurnsQuery) (turns model.ChinchiroGameTurns, err error) {
	return findChinchiroGameTurns(repo.db.Connection, query)
}

func (repo *ChinchiroGameRepository) FindGameTurn(ID uint32) (turn *model.ChinchiroGameTurn, err error) {
	return findChinchiroGameTurn(repo.db.Connection, ID)
}

func (repo *ChinchiroGameRepository) RegisterGameTurn(ctx context.Context, gameID uint32, turn model.ChinchiroGameTurn) (saved *model.ChinchiroGameTurn, err error) {
	tx, ok := GetTx(ctx)
	if !ok {
		return nil, fmt.Errorf("failed to get tx from context")
	}
	// turn
	t, err := registerRdbChinchiroGameTurn(tx, gameID, turn)
	if err != nil {
		return nil, err
	}
	return findChinchiroGameTurn(tx, t.ID)
}

func (repo *ChinchiroGameRepository) UpdateGameTurnStatus(ctx context.Context, turnID uint32, status model.ChinchiroGameTurnStatus) (err error) {
	tx, ok := GetTx(ctx)
	if !ok {
		return fmt.Errorf("failed to get tx from context")
	}
	result := tx.Model(&ChinchiroGameTurn{}).Where("id = ?", turnID).Update("turn_status_code", status.String())
	if result.Error != nil {
		return fmt.Errorf("failed to update: %s", result.Error)
	}
	return nil
}

func (repo *ChinchiroGameRepository) UpdateGameTurnRoller(ctx context.Context, turnID uint32, rollerID *uint32) (err error) {
	tx, ok := GetTx(ctx)
	if !ok {
		return fmt.Errorf("failed to get tx from context")
	}
	result := tx.Model(&ChinchiroGameTurn{}).Where("id = ?", turnID).Update("next_roller_participant_id", rollerID)
	if result.Error != nil {
		return fmt.Errorf("failed to update: %s", result.Error)
	}
	return nil
}

func (repo *ChinchiroGameRepository) FindGameTurnRolls(query model.ChinchiroGameTurnRollsQuery) (rolls model.ChinchiroGameTurnRolls, err error) {
	return findChinchiroGameTurnRolls(repo.db.Connection, query)
}

func (repo *ChinchiroGameRepository) RegisterGameTurnRoll(ctx context.Context, gameID uint32, turnID uint32, roll model.ChinchiroGameTurnRoll) (saved *model.ChinchiroGameTurnRoll, err error) {
	tx, ok := GetTx(ctx)
	if !ok {
		return nil, fmt.Errorf("failed to get tx from context")
	}
	// roll
	r, err := registerRdbChinchiroGameTurnRoll(tx, gameID, turnID, roll)
	if err != nil {
		return nil, err
	}
	return r.ToModel(), nil
}

func (repo *ChinchiroGameRepository) FindGameTurnResults(query model.ChinchiroGameTurnResultsQuery) (results model.ChinchiroGameTurnResults, err error) {
	return findChinchiroGameTurnResults(repo.db.Connection, query)
}

func (repo *ChinchiroGameRepository) FindGameTurnResult(query model.ChinchiroGameTurnResultQuery) (result *model.ChinchiroGameTurnResult, err error) {
	return findChinchiroGameTurnResult(repo.db.Connection, query)
}

func (repo *ChinchiroGameRepository) RegisterGameTurnResult(ctx context.Context, gameID uint32, turnID uint32, result model.ChinchiroGameTurnResult) (saved *model.ChinchiroGameTurnResult, err error) {
	tx, ok := GetTx(ctx)
	if !ok {
		return nil, fmt.Errorf("failed to get tx from context")
	}
	// result
	r, err := registerRdbChinchiroGameTurnResult(tx, gameID, turnID, result)
	if err != nil {
		return nil, err
	}
	return findChinchiroGameTurnResult(tx, model.ChinchiroGameTurnResultQuery{ID: &r.ID})
}

func (repo *ChinchiroGameRepository) UpdateGameTurnResult(ctx context.Context, resultID uint32, result model.ChinchiroGameTurnResult) (saved *model.ChinchiroGameTurnResult, err error) {
	tx, ok := GetTx(ctx)
	if !ok {
		return nil, fmt.Errorf("failed to get tx from context")
	}
	// result
	_, err = updateRdbChinchiroGameTurnResult(tx, resultID, result)
	if err != nil {
		return nil, err
	}
	return findChinchiroGameTurnResult(tx, model.ChinchiroGameTurnResultQuery{ID: &resultID})
}

// -----------------------

func findChinchiroGames(db *gorm.DB, query model.ChinchiroGamesQuery) (model.ChinchiroGames, error) {
	rdbs, err := findRdbChinchiroGames(db, query)
	if err != nil {
		return model.ChinchiroGames{}, err
	}
	if rdbs == nil {
		return model.ChinchiroGames{}, nil
	}

	ids := array.Map(rdbs, func(rdb ChinchiroGame) uint32 {
		return rdb.ID
	})
	if len(ids) == 0 {
		return model.ChinchiroGames{}, nil
	}

	participants, err := findChinchiroGameParticipants(db, model.ChinchiroGameParticipantsQuery{
		GameIDs: &ids,
	})
	if err != nil {
		return model.ChinchiroGames{}, err
	}

	turns, err := findChinchiroGameTurns(db, model.ChinchiroGameTurnsQuery{
		GameIDs: &ids,
	})
	if err != nil {
		return model.ChinchiroGames{}, err
	}

	games := array.Map(rdbs, func(rdb ChinchiroGame) model.ChinchiroGame {
		return *rdb.ToModel(participants, turns)
	})
	return model.ChinchiroGames{
		List:  games,
		Count: len(games),
	}, nil
}

func findChinchiroGame(db *gorm.DB, ID uint32) (*model.ChinchiroGame, error) {
	rdb, err := findRdbChinchiroGame(db, ID)
	if err != nil {
		return nil, err
	}
	if rdb == nil {
		return nil, nil
	}

	participants, err := findChinchiroGameParticipants(db, model.ChinchiroGameParticipantsQuery{
		GameID: &ID,
	})
	if err != nil {
		return nil, err
	}

	turns, err := findChinchiroGameTurns(db, model.ChinchiroGameTurnsQuery{
		GameID: &ID,
	})
	if err != nil {
		return nil, err
	}

	return rdb.ToModel(participants, turns), nil
}

func findChinchiroGameParticipants(db *gorm.DB, query model.ChinchiroGameParticipantsQuery) (participants model.ChinchiroGameParticipants, err error) {
	rdbs, err := findRdbChinchiroGameParticipants(db, query)
	if err != nil {
		return model.ChinchiroGameParticipants{}, err
	}
	if rdbs == nil {
		return model.ChinchiroGameParticipants{}, nil
	}
	pts := array.Map(rdbs, func(rdb ChinchiroGameParticipant) model.ChinchiroGameParticipant {
		return *rdb.ToModel()
	})
	return model.ChinchiroGameParticipants{
		List:  pts,
		Count: len(pts),
	}, nil
}

func findChinchiroGameParticipant(db *gorm.DB, query model.ChinchiroGameParticipantQuery) (*model.ChinchiroGameParticipant, error) {
	rdb, err := findRdbChinchiroGameParticipant(db, query)
	if err != nil {
		return nil, err
	}
	if rdb == nil {
		return nil, nil
	}
	return rdb.ToModel(), nil
}

func findChinchiroGameTurns(db *gorm.DB, query model.ChinchiroGameTurnsQuery) (turns model.ChinchiroGameTurns, err error) {
	rdbs, err := findRdbChinchiroGameTurns(db, query)
	if err != nil {
		return model.ChinchiroGameTurns{}, err
	}
	if rdbs == nil {
		return model.ChinchiroGameTurns{}, nil
	}
	ids := array.Map(rdbs, func(rdb ChinchiroGameTurn) uint32 {
		return rdb.ID
	})
	if len(ids) == 0 {
		return model.ChinchiroGameTurns{}, nil
	}
	rolls, err := findChinchiroGameTurnRolls(db, model.ChinchiroGameTurnRollsQuery{
		TurnIDs: &ids,
	})
	if err != nil {
		return model.ChinchiroGameTurns{}, err
	}
	results, err := findChinchiroGameTurnResults(db, model.ChinchiroGameTurnResultsQuery{
		TurnIDs: &ids,
	})
	if err != nil {
		return model.ChinchiroGameTurns{}, err
	}

	ts := array.Map(rdbs, func(rdb ChinchiroGameTurn) model.ChinchiroGameTurn {
		turnRolls := array.Filter(rolls.List, func(r model.ChinchiroGameTurnRoll) bool {
			return r.TurnID == rdb.ID
		})
		turnResults := array.Filter(results.List, func(r model.ChinchiroGameTurnResult) bool {
			return r.TurnID == rdb.ID
		})
		return *rdb.ToModel(turnRolls, turnResults)
	})
	return model.ChinchiroGameTurns{
		List:  ts,
		Count: len(ts),
	}, nil
}

func findChinchiroGameTurn(db *gorm.DB, ID uint32) (*model.ChinchiroGameTurn, error) {
	rdb, err := findRdbChinchiroGameTurn(db, ID)
	if err != nil {
		return nil, err
	}
	if rdb == nil {
		return nil, nil
	}
	rolls, err := findChinchiroGameTurnRolls(db, model.ChinchiroGameTurnRollsQuery{
		TurnID: &ID,
	})
	if err != nil {
		return nil, err
	}
	results, err := findChinchiroGameTurnResults(db, model.ChinchiroGameTurnResultsQuery{
		TurnID: &ID,
	})
	if err != nil {
		return nil, err
	}
	return rdb.ToModel(
		rolls.List,
		results.List,
	), nil
}

func findChinchiroGameTurnRolls(db *gorm.DB, query model.ChinchiroGameTurnRollsQuery) (model.ChinchiroGameTurnRolls, error) {
	rdbs, err := findRdbChinchiroGameTurnRolls(db, query)
	if err != nil {
		return model.ChinchiroGameTurnRolls{}, err
	}
	if rdbs == nil {
		return model.ChinchiroGameTurnRolls{}, nil
	}
	rolls := array.Map(rdbs, func(rdb ChinchiroGameTurnParticipantRoll) model.ChinchiroGameTurnRoll {
		return *rdb.ToModel()
	})
	return model.ChinchiroGameTurnRolls{
		List:  rolls,
		Count: len(rolls),
	}, nil
}

func findChinchiroGameTurnResults(db *gorm.DB, query model.ChinchiroGameTurnResultsQuery) (model.ChinchiroGameTurnResults, error) {
	rdbs, err := findRdbChinchiroGameTurnResults(db, query)
	if err != nil {
		return model.ChinchiroGameTurnResults{}, err
	}
	if rdbs == nil {
		return model.ChinchiroGameTurnResults{}, nil
	}
	results := array.Map(rdbs, func(rdb ChinchiroGameTurnParticipantResult) model.ChinchiroGameTurnResult {
		return *rdb.ToModel()
	})
	return model.ChinchiroGameTurnResults{
		List:  results,
		Count: len(results),
	}, nil
}

func findChinchiroGameTurnResult(db *gorm.DB, query model.ChinchiroGameTurnResultQuery) (*model.ChinchiroGameTurnResult, error) {
	rdb, err := findRdbChinchiroGameTurnResult(db, query)
	if err != nil {
		return nil, err
	}
	if rdb == nil {
		return nil, nil
	}
	return rdb.ToModel(), nil
}

// -----------------------

func findRdbChinchiroGames(db *gorm.DB, query model.ChinchiroGamesQuery) (games []ChinchiroGame, err error) {
	var rdbs []ChinchiroGame
	result := db.Model(&ChinchiroGame{})
	if query.Paging != nil {
		result = result.Scopes(Paginate(query.Paging))
	} else {
		result = result.Scopes(Paginate(&model.PagingQuery{
			PageSize:   10000,
			PageNumber: 1,
			Desc:       false,
		}))
	}
	if query.IDs != nil {
		result = result.Where("id in (?)", *query.IDs)
	}
	if query.RoomID != nil {
		result = result.Where("room_id = ?", *query.RoomID)
	}
	if query.Name != nil {
		result = result.Where("game_name like ?", fmt.Sprintf("%%%s%%", *query.Name))
	}
	if query.Statuses != nil {
		statuses := array.Map(*query.Statuses, func(s model.ChinchiroGameStatus) string {
			return s.String()
		})
		result = result.Where("game_status_code in (?)", statuses)
	}
	result = result.Find(&rdbs)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if result.Error != nil {
		return nil, fmt.Errorf("failed to find: %s", result.Error)
	}

	return rdbs, nil
}

func findRdbChinchiroGame(db *gorm.DB, ID uint32) (*ChinchiroGame, error) {
	var rdb ChinchiroGame
	result := db.Model(&ChinchiroGame{}).First(&rdb, ID)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if result.Error != nil {
		return nil, fmt.Errorf("failed to find: %s", result.Error)
	}
	return &rdb, nil
}

func registerRdbChinchiroGame(tx *gorm.DB, roomID uint32, _ model.ChinchiroGame) (*ChinchiroGame, error) {
	r := ChinchiroGame{
		RoomID:         roomID,
		GameStatusCode: model.ChinchiroGameStatusProgress.String(),
	}
	result := tx.Model(&ChinchiroGame{}).Create(&r)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to create: %s", result.Error)
	}
	return &r, nil
}

func findRdbChinchiroGameParticipants(db *gorm.DB, query model.ChinchiroGameParticipantsQuery) ([]ChinchiroGameParticipant, error) {
	var rdbs []ChinchiroGameParticipant
	result := db.Model(&ChinchiroGameParticipant{})
	if query.IDs != nil {
		result = result.Where("id in (?)", *query.IDs)
	}
	if query.GameID != nil {
		result = result.Where("game_id = ?", *query.GameID)
	}
	if query.GameIDs != nil {
		result = result.Where("game_id in (?)", *query.GameIDs)
	}
	result = result.Find(&rdbs)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if result.Error != nil {
		return nil, fmt.Errorf("failed to find: %s", result.Error)
	}

	return rdbs, nil
}

func findRdbChinchiroGameParticipant(db *gorm.DB, query model.ChinchiroGameParticipantQuery) (*ChinchiroGameParticipant, error) {
	var rdb ChinchiroGameParticipant
	result := db.Model(&ChinchiroGameParticipant{})
	if query.ID != nil {
		result = result.Where("id = ?", *query.ID)
	}
	if query.GameID != nil {
		result = result.Where("game_id = ?", *query.GameID)
	}
	if query.RoomParticipantID != nil {
		result = result.Where("room_participant_id = ?", *query.RoomParticipantID)
	}
	result = result.First(&rdb)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if result.Error != nil {
		return nil, fmt.Errorf("failed to find: %s", result.Error)
	}
	return &rdb, nil
}

func registerRdbChinchiroGameParticipant(tx *gorm.DB, gameID uint32, participant model.ChinchiroGameParticipant) (*ChinchiroGameParticipant, error) {
	p := ChinchiroGameParticipant{
		GameID:            gameID,
		RoomParticipantID: participant.RoomParticipantID,
		Balance:           participant.Balance,
		TurnOrder:         participant.TurnOrder,
	}
	result := tx.Model(&ChinchiroGameParticipant{}).Create(&p)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to create: %s", result.Error)
	}
	return &p, nil
}

func findRdbChinchiroGameTurns(db *gorm.DB, query model.ChinchiroGameTurnsQuery) ([]ChinchiroGameTurn, error) {
	var rdbs []ChinchiroGameTurn
	result := db.Model(&ChinchiroGameTurn{})
	if query.Paging != nil {
		result = result.Scopes(Paginate(query.Paging))
	} else {
		result = result.Scopes(Paginate(&model.PagingQuery{
			PageSize:   10000,
			PageNumber: 1,
			Desc:       false,
		}))
	}
	if query.IDs != nil {
		result = result.Where("id in (?)", *query.IDs)
	}
	if query.GameID != nil {
		result = result.Where("game_id = ?", *query.GameID)
	}
	if query.GameIDs != nil {
		result = result.Where("game_id in (?)", *query.GameIDs)
	}
	if query.Statuses != nil {
		statuses := array.Map(*query.Statuses, func(s model.ChinchiroGameTurnStatus) string {
			return s.String()
		})
		result = result.Where("turn_status_code in (?)", statuses)
	}

	result = result.Find(&rdbs)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if result.Error != nil {
		return nil, fmt.Errorf("failed to find: %s", result.Error)
	}

	return rdbs, nil
}

func findRdbChinchiroGameTurn(db *gorm.DB, ID uint32) (*ChinchiroGameTurn, error) {
	var rdb ChinchiroGameTurn
	result := db.Model(&ChinchiroGameTurn{}).First(&rdb, ID)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if result.Error != nil {
		return nil, fmt.Errorf("failed to find: %s", result.Error)
	}
	return &rdb, nil
}

func registerRdbChinchiroGameTurn(tx *gorm.DB, gameID uint32, turn model.ChinchiroGameTurn) (*ChinchiroGameTurn, error) {
	t := ChinchiroGameTurn{
		GameID:              gameID,
		DealerParticipantID: turn.DealerID,
		TurnStatusCode:      model.ChinchiroGameTurnStatusBetting.String(),
		TurnNumber:          turn.TurnNumber,
	}
	result := tx.Model(&ChinchiroGameTurn{}).Create(&t)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to create: %s", result.Error)
	}
	return &t, nil
}

func findRdbChinchiroGameTurnRolls(db *gorm.DB, query model.ChinchiroGameTurnRollsQuery) ([]ChinchiroGameTurnParticipantRoll, error) {
	var rdbs []ChinchiroGameTurnParticipantRoll
	result := db.Model(&ChinchiroGameTurnParticipantRoll{})
	if query.Paging != nil {
		result = result.Scopes(Paginate(query.Paging))
	} else {
		result = result.Scopes(Paginate(&model.PagingQuery{
			PageSize:   10000,
			PageNumber: 1,
			Desc:       false,
		}))
	}
	if query.IDs != nil {
		result = result.Where("id in (?)", *query.IDs)
	}
	if query.TurnID != nil {
		result = result.Where("turn_id = ?", *query.TurnID)
	}
	if query.TurnIDs != nil {
		result = result.Where("turn_id in (?)", *query.TurnIDs)
	}
	if query.ParticipantID != nil {
		result = result.Where("participant_id = ?", *query.ParticipantID)
	}
	result = result.Find(&rdbs)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if result.Error != nil {
		return nil, fmt.Errorf("failed to find: %s", result.Error)
	}

	return rdbs, nil
}

func registerRdbChinchiroGameTurnRoll(tx *gorm.DB, gameID uint32, turnID uint32, roll model.ChinchiroGameTurnRoll) (*ChinchiroGameTurnParticipantRoll, error) {
	r := ChinchiroGameTurnParticipantRoll{
		GameID:        gameID,
		TurnID:        turnID,
		ParticipantID: roll.ParticipantID,
		RollNumber:    roll.RollNumber,
		Dice1:         roll.DiceRoll.Dice1,
		Dice2:         roll.DiceRoll.Dice2,
		Dice3:         roll.DiceRoll.Dice3,
	}
	result := tx.Model(&ChinchiroGameTurnParticipantRoll{}).Create(&r)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to create: %s", result.Error)
	}
	return &r, nil
}

func findRdbChinchiroGameTurnResults(db *gorm.DB, query model.ChinchiroGameTurnResultsQuery) ([]ChinchiroGameTurnParticipantResult, error) {
	var rdbs []ChinchiroGameTurnParticipantResult
	result := db.Model(&ChinchiroGameTurnParticipantResult{})
	if query.Paging != nil {
		result = result.Scopes(Paginate(query.Paging))
	} else {
		result = result.Scopes(Paginate(&model.PagingQuery{
			PageSize:   10000,
			PageNumber: 1,
			Desc:       false,
		}))
	}
	if query.IDs != nil {
		result = result.Where("id in (?)", *query.IDs)
	}
	if query.TurnID != nil {
		result = result.Where("turn_id = ?", *query.TurnID)
	}
	if query.TurnIDs != nil {
		result = result.Where("turn_id in (?)", *query.TurnIDs)
	}
	result = result.Find(&rdbs)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if result.Error != nil {
		return nil, fmt.Errorf("failed to find: %s", result.Error)
	}

	return rdbs, nil
}

func findRdbChinchiroGameTurnResult(db *gorm.DB, query model.ChinchiroGameTurnResultQuery) (*ChinchiroGameTurnParticipantResult, error) {
	var rdb ChinchiroGameTurnParticipantResult
	result := db.Model(&ChinchiroGameTurnParticipantResult{})

	if query.ID != nil {
		result = result.Where("id = ?", *query.ID)
	}
	if query.TurnID != nil {
		result = result.Where("turn_id = ?", *query.TurnID)
	}
	if query.ParticipantID != nil {
		result = result.Where("participant_id = ?", *query.ParticipantID)
	}
	result = result.First(&rdb)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if result.Error != nil {
		return nil, fmt.Errorf("failed to find: %s", result.Error)
	}
	return &rdb, nil
}

func registerRdbChinchiroGameTurnResult(tx *gorm.DB, gameID uint32, turnID uint32, res model.ChinchiroGameTurnResult) (*ChinchiroGameTurnParticipantResult, error) {
	r := ChinchiroGameTurnParticipantResult{
		GameID:        gameID,
		TurnID:        turnID,
		ParticipantID: res.ParticipantID,
		BetAmount:     res.BetAmount,
	}
	result := tx.Model(&ChinchiroGameTurnParticipantRoll{}).Create(&r)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to create: %s", result.Error)
	}
	return &r, nil
}

func updateRdbChinchiroGameTurnResult(tx *gorm.DB, resultID uint32, res model.ChinchiroGameTurnResult) (*ChinchiroGameTurnParticipantResult, error) {
	r := ChinchiroGameTurnParticipantResult{
		Dice1:           res.DiceRoll.Dice1,
		Dice2:           res.DiceRoll.Dice2,
		Dice3:           res.DiceRoll.Dice3,
		CombinationCode: res.DiceRoll.Combination.String(),
		Winnings:        *res.Winnings,
	}
	result := tx.Model(&ChinchiroGameTurnParticipantRoll{}).Where("id = ?", resultID).Updates(&r)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to create: %s", result.Error)
	}
	return &r, nil
}
