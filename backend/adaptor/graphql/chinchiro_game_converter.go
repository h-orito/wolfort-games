package graphql

import (
	"wolfort-games/domain/model"
	"wolfort-games/middleware/graph/gqlmodel"
	"wolfort-games/util/array"
)

func MapToChinchiroGame(
	p *model.ChinchiroGame,
) *gqlmodel.ChinchiroGame {
	if p == nil {
		return nil
	}
	return &gqlmodel.ChinchiroGame{
		ID:     intIdToBase64(p.ID, "ChinchiroGame"),
		Status: gqlmodel.ChinchiroGameStatus(p.Status.String()),
		ParticipantIDs: array.Map(p.ParticipantIDs, func(id uint32) string {
			return intIdToBase64(id, "ChinchiroGameParticipant")
		}),
		TurnIDs: array.Map(p.TurnIDs, func(id uint32) string {
			return intIdToBase64(id, "ChinchiroGameTurn")
		}),
	}
}

func MapToChinchiroGameParticipant(
	p *model.ChinchiroGameParticipant,
) *gqlmodel.ChinchiroGameParticipant {
	if p == nil {
		return nil
	}
	return &gqlmodel.ChinchiroGameParticipant{
		ID:                intIdToBase64(p.ID, "ChinchiroGameParticipant"),
		RoomParticipantID: intIdToBase64(p.RoomParticipantID, "ChinchiroRoomParticipant"),
		Balance:           p.Balance,
		TurnOrder:         p.TurnOrder,
	}
}

func MapToChinchiroGameTurn(
	p *model.ChinchiroGameTurn,
) *gqlmodel.ChinchiroGameTurn {
	if p == nil {
		return nil
	}
	var nextRollerID *string
	if p.NextRollerID != nil {
		nrID := intIdToBase64(*p.NextRollerID, "ChinchiroGameParticipant")
		nextRollerID = &nrID
	}
	return &gqlmodel.ChinchiroGameTurn{
		ID:           intIdToBase64(p.ID, "ChinchiroGameTurn"),
		DealerID:     intIdToBase64(p.DealerID, "ChinchiroRoomParticipant"),
		NextRollerID: nextRollerID,
		Status:       gqlmodel.ChinchiroGameTurnStatus(p.Status.String()),
		TurnNumber:   p.TurnNumber,
		RollIDs: array.Map(p.RollIDs, func(id uint32) string {
			return intIdToBase64(id, "ChinchiroGameRoll")
		}),
		ResultIDs: array.Map(p.ResultIDs, func(id uint32) string {
			return intIdToBase64(id, "ChinchiroGameResult")
		}),
	}
}

func MapToChinchiroGameParticipantRoll(
	p *model.ChinchiroGameTurnRoll,
) *gqlmodel.ChinchiroGameTurnParticipantRoll {
	if p == nil {
		return nil
	}
	return &gqlmodel.ChinchiroGameTurnParticipantRoll{
		ID:            intIdToBase64(p.ID, "ChinchiroGameRoll"),
		TurnID:        intIdToBase64(p.TurnID, "ChinchiroGameTurn"),
		ParticipantID: intIdToBase64(p.ParticipantID, "ChinchiroGameParticipant"),
		RollNumber:    p.RollNumber,
		Dice1:         p.DiceRoll.Dice1,
		Dice2:         p.DiceRoll.Dice2,
		Dice3:         p.DiceRoll.Dice3,
	}
}

func MapToChinchiroGameParticipantResult(
	p *model.ChinchiroGameTurnResult,
) *gqlmodel.ChinchiroGameTurnParticipantResult {
	if p == nil {
		return nil
	}
	return &gqlmodel.ChinchiroGameTurnParticipantResult{
		ID:            intIdToBase64(p.ID, "ChinchiroGameResult"),
		GameTurnID:    intIdToBase64(p.TurnID, "ChinchiroGameTurn"),
		ParticipantID: intIdToBase64(p.ParticipantID, "ChinchiroGameParticipant"),
		Dice1:         p.DiceRoll.Dice1,
		Dice2:         p.DiceRoll.Dice2,
		Dice3:         p.DiceRoll.Dice3,
		Combination:   gqlmodel.ChinchiroCombination(p.DiceRoll.Combination.String()),
		Winnings:      *p.Winnings,
	}
}
