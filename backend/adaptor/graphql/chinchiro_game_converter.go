package graphql

import (
	"chat-role-play/domain/model"
	"chat-role-play/middleware/graph/gqlmodel"
	"chat-role-play/util/array"
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
		Participants: array.Map(p.Participants.List, func(p model.ChinchiroGameParticipant) *gqlmodel.ChinchiroGameParticipant {
			return MapToChinchiroGameParticipant(&p)
		}),
		Turns: array.Map(p.Turns.List, func(t model.ChinchiroGameTurn) *gqlmodel.ChinchiroGameTurn {
			return MapToChinchiroGameTurn(&t)
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
	return &gqlmodel.ChinchiroGameTurn{
		ID:         intIdToBase64(p.ID, "ChinchiroGameTurn"),
		DealerID:   intIdToBase64(p.DealerID, "ChinchiroRoomParticipant"),
		Status:     gqlmodel.ChinchiroGameTurnStatus(p.Status.String()),
		TurnNumber: p.TurnNumber,
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
		Winnings:      p.Winnings,
	}
}
