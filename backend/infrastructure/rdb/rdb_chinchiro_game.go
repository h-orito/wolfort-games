package db

import (
	"time"
	model "wolfort-games/domain/model"
	"wolfort-games/util/array"
)

type ChinchiroGame struct {
	ID             uint32
	RoomID         uint32
	GameStatusCode string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func (r ChinchiroGame) ToModel(
	participants model.ChinchiroGameParticipants,
	turns model.ChinchiroGameTurns,
) *model.ChinchiroGame {
	return &model.ChinchiroGame{
		ID:             r.ID,
		RoomID:         r.RoomID,
		Status:         *model.ChinchiroGameStatusValueOf(r.GameStatusCode),
		ParticipantIDs: array.Map(participants.List, func(p model.ChinchiroGameParticipant) uint32 { return p.ID }),
		TurnIDs:        array.Map(turns.List, func(t model.ChinchiroGameTurn) uint32 { return t.ID }),
	}
}

type ChinchiroGameParticipant struct {
	ID                uint32
	GameID            uint32
	RoomParticipantID uint32
	Balance           int
	TurnOrder         int
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

func (p ChinchiroGameParticipant) ToModel() *model.ChinchiroGameParticipant {
	return &model.ChinchiroGameParticipant{
		ID:                p.ID,
		RoomParticipantID: p.RoomParticipantID,
		Balance:           p.Balance,
		TurnOrder:         p.TurnOrder,
	}
}

type ChinchiroGameTurn struct {
	ID                      uint32
	GameID                  uint32
	DealerParticipantID     uint32
	NextRollerParticipantID *uint32
	TurnStatusCode          string
	TurnNumber              int
	CreatedAt               time.Time
	UpdatedAt               time.Time
}

func (t ChinchiroGameTurn) ToModel(
	rolls []model.ChinchiroGameTurnRoll,
	results []model.ChinchiroGameTurnResult,
) *model.ChinchiroGameTurn {
	return &model.ChinchiroGameTurn{
		ID:           t.ID,
		GameID:       t.GameID,
		DealerID:     t.DealerParticipantID,
		NextRollerID: t.NextRollerParticipantID,
		Status:       *model.ChinchiroGameTurnStatusValueOf(t.TurnStatusCode),
		TurnNumber:   t.TurnNumber,
		RollIDs:      array.Map(rolls, func(r model.ChinchiroGameTurnRoll) uint32 { return r.ID }),
		ResultIDs:    array.Map(results, func(r model.ChinchiroGameTurnResult) uint32 { return r.ID }),
	}
}

type ChinchiroGameTurnParticipantRoll struct {
	ID            uint32
	GameID        uint32
	TurnID        uint32
	ParticipantID uint32
	RollNumber    int
	Dice1         int
	Dice2         int
	Dice3         int
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (r ChinchiroGameTurnParticipantRoll) ToModel() *model.ChinchiroGameTurnRoll {
	return &model.ChinchiroGameTurnRoll{
		ID:            r.ID,
		TurnID:        r.TurnID,
		ParticipantID: r.ParticipantID,
		RollNumber:    r.RollNumber,
		DiceRoll: model.NewChinchiroDiceRoll(
			r.Dice1,
			r.Dice2,
			r.Dice3,
		),
	}
}

type ChinchiroGameTurnParticipantResult struct {
	ID              uint32
	GameID          uint32
	TurnID          uint32
	ParticipantID   uint32
	BetAmount       int
	Dice1           int
	Dice2           int
	Dice3           int
	CombinationCode string
	Winnings        int
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func (r ChinchiroGameTurnParticipantResult) ToModel() *model.ChinchiroGameTurnResult {
	diceRoll := model.NewChinchiroDiceRoll(
		r.Dice1,
		r.Dice2,
		r.Dice3,
	)
	return &model.ChinchiroGameTurnResult{
		ID:            r.ID,
		TurnID:        r.TurnID,
		ParticipantID: r.ParticipantID,
		BetAmount:     r.BetAmount,
		DiceRoll:      &diceRoll,
		Winnings:      &r.Winnings,
	}
}
