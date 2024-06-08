package db

import (
	model "chat-role-play/domain/model"
	"time"
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
		ID:           r.ID,
		Status:       *model.ChinchiroGameStatusValueOf(r.GameStatusCode),
		Participants: participants,
		Turns:        turns,
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
	ID                 uint32
	GameID             uint32
	DealerPartcipantID uint32
	TurnStatusCode     string
	TurnNumber         int
	CreatedAt          time.Time
	UpdatedAt          time.Time
}

func (t ChinchiroGameTurn) ToModel(
	rollIDs []uint32,
	resultIDs []uint32,
) *model.ChinchiroGameTurn {
	return &model.ChinchiroGameTurn{
		ID:         t.ID,
		DealerID:   t.DealerPartcipantID,
		Status:     *model.ChinchiroGameTurnStatusValueOf(t.TurnStatusCode),
		TurnNumber: t.TurnNumber,
		RollIDs:    rollIDs,
		ResultIDs:  resultIDs,
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
	return &model.ChinchiroGameTurnResult{
		ID:            r.ID,
		TurnID:        r.TurnID,
		ParticipantID: r.ParticipantID,
		BetAmount:     r.BetAmount,
		DiceRoll: model.NewChinchiroDiceRoll(
			r.Dice1,
			r.Dice2,
			r.Dice3,
		),
		Winnings: r.Winnings,
	}
}
