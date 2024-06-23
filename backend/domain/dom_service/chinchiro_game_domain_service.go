package dom_service

import (
	"wolfort-games/domain/model"
	"wolfort-games/util/array"
)

type ChinchiroGameDomainService interface {
	ShouldGameFinished(participants model.ChinchiroGameParticipants, turns model.ChinchiroGameTurns) bool
	ShouldTurnRolling(
		turn model.ChinchiroGameTurn,
		participants model.ChinchiroGameParticipants,
		results model.ChinchiroGameTurnResults,
	) bool
	DetermineNextDealerID(
		participants model.ChinchiroGameParticipants,
		turns model.ChinchiroGameTurns,
		currentDealer model.ChinchiroGameParticipant,
	) uint32
	DetermineNextRoller(
		turn model.ChinchiroGameTurn,
		participants model.ChinchiroGameParticipants,
		results model.ChinchiroGameTurnResults,
	) *model.ChinchiroGameParticipant
}

type chinchiroGameDomainService struct {
}

// 現在誰が振る順番か
func (c *chinchiroGameDomainService) DetermineNextRoller(
	turn model.ChinchiroGameTurn,
	participants model.ChinchiroGameParticipants,
	results model.ChinchiroGameTurnResults,
) *model.ChinchiroGameParticipant {
	dealerTurnOrder := array.Find(participants.List, func(p model.ChinchiroGameParticipant) bool {
		return p.ID == turn.DealerID
	}).TurnOrder

	// 親から順に振る順番を並べて
	var arr []model.ChinchiroGameParticipant

	for turnOrder := dealerTurnOrder; turnOrder <= participants.Count; turnOrder++ {
		participant := array.Find(participants.List, func(p model.ChinchiroGameParticipant) bool {
			return p.TurnOrder == turnOrder
		})
		arr = append(arr, *participant)
	}
	for turnOrder := 1; turnOrder < dealerTurnOrder; turnOrder++ {
		participant := array.Find(participants.List, func(p model.ChinchiroGameParticipant) bool {
			return p.TurnOrder == turnOrder
		})
		arr = append(arr, *participant)
	}

	// まだ振っていない人を探す
	for _, p := range arr {
		result := array.Find(results.List, func(r model.ChinchiroGameTurnResult) bool {
			return r.ParticipantID == p.ID
		})
		if result == nil || result.DiceRoll == nil {
			return &p
		}
	}
	return nil
}

func (c *chinchiroGameDomainService) ShouldGameFinished(participants model.ChinchiroGameParticipants, turns model.ChinchiroGameTurns) bool {
	// 全員が2回親をやったらゲーム終了
	return array.All(participants.List, func(p model.ChinchiroGameParticipant) bool {
		return array.Count(turns.List, func(t model.ChinchiroGameTurn) bool {
			return t.DealerID == p.ID
		}) >= 2
	})
}

func (c *chinchiroGameDomainService) ShouldTurnRolling(
	turn model.ChinchiroGameTurn,
	participants model.ChinchiroGameParticipants,
	results model.ChinchiroGameTurnResults,
) bool {
	// 親以外の全員がbetしていたらロール開始
	return array.All(array.Filter(participants.List, func(p model.ChinchiroGameParticipant) bool {
		return p.ID != turn.DealerID
	}), func(p model.ChinchiroGameParticipant) bool {
		return array.Any(results.List, func(r model.ChinchiroGameTurnResult) bool {
			return r.ParticipantID == p.ID && r.BetAmount > 0
		})
	})
}

func (c *chinchiroGameDomainService) DetermineNextDealerID(
	participants model.ChinchiroGameParticipants,
	turns model.ChinchiroGameTurns,
	currentDealer model.ChinchiroGameParticipant,
) uint32 {
	if array.Count(turns.List, func(t model.ChinchiroGameTurn) bool {
		return t.DealerID == currentDealer.ID
	}) >= 2 {
		// 2回親をやったので次の人を親にする
		return array.Find(participants.List, func(p model.ChinchiroGameParticipant) bool {
			return p.TurnOrder == currentDealer.TurnOrder+1
		}).ID
	} else {
		// 1回しか親をやっていないので親を続行
		return currentDealer.ID
	}
}

func NewChinchiroGameDomainService() ChinchiroGameDomainService {
	return &chinchiroGameDomainService{}
}
