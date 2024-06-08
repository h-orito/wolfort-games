package graphql

import (
	"chat-role-play/application/usecase"
	"chat-role-play/domain/model"
	"chat-role-play/util/array"
	"context"

	"github.com/graph-gophers/dataloader"
)

type Loaders struct {
	PlayerLoader                   *dataloader.Loader
	ChinchiroRoomParticipantLoader *dataloader.Loader
	ChinchiroRoomMasterLoader      *dataloader.Loader
	ChinchiroGameLoader            *dataloader.Loader
	ChinchiroGameParticipantLoader *dataloader.Loader
	ChinchiroGameTurnLoader        *dataloader.Loader
	ChinchiroGameTurnRollLoader    *dataloader.Loader
	ChinchiroGameTurnResultLoader  *dataloader.Loader
}

func NewLoaders(
	playerUsecase usecase.PlayerUsecase,
	chinchiroRoomUsecase usecase.ChinchiroRoomUsecase,
	chinchiroGameUsecase usecase.ChinchiroGameUsecase,
) *Loaders {
	chinchiroBatcher := &chinchiroBatcher{
		chinchiroRoomUsecase: chinchiroRoomUsecase,
		chinchiroGameUsecase: chinchiroGameUsecase,
	}
	playerBatcher := &playerBatcher{playerUsecase: playerUsecase}
	return &Loaders{
		PlayerLoader:                   dataloader.NewBatchedLoader(playerBatcher.batchLoadPlayer),
		ChinchiroRoomParticipantLoader: dataloader.NewBatchedLoader(chinchiroBatcher.batchLoadRoomParticipant),
		ChinchiroRoomMasterLoader:      dataloader.NewBatchedLoader(chinchiroBatcher.batchLoadRoomMaster),
		ChinchiroGameLoader:            dataloader.NewBatchedLoader(chinchiroBatcher.batchLoadGame),
		ChinchiroGameParticipantLoader: dataloader.NewBatchedLoader(chinchiroBatcher.batchLoadGameParticipant),
		ChinchiroGameTurnLoader:        dataloader.NewBatchedLoader(chinchiroBatcher.batchLoadGameTurn),
		ChinchiroGameTurnRollLoader:    dataloader.NewBatchedLoader(chinchiroBatcher.batchLoadGameTurnRoll),
		ChinchiroGameTurnResultLoader:  dataloader.NewBatchedLoader(chinchiroBatcher.batchLoadGameTurnResult),
	}
}

type chinchiroBatcher struct {
	chinchiroRoomUsecase usecase.ChinchiroRoomUsecase
	chinchiroGameUsecase usecase.ChinchiroGameUsecase
}

func NewChinchiroBatcher(chinchiroRoomUsecase usecase.ChinchiroRoomUsecase, chinchiroGameUsecase usecase.ChinchiroGameUsecase) *chinchiroBatcher {
	return &chinchiroBatcher{
		chinchiroRoomUsecase: chinchiroRoomUsecase,
		chinchiroGameUsecase: chinchiroGameUsecase,
	}
}

type playerBatcher struct {
	playerUsecase usecase.PlayerUsecase
}

func NewPlayerBatcher(playerUsecase usecase.PlayerUsecase) *playerBatcher {
	return &playerBatcher{
		playerUsecase: playerUsecase,
	}
}

func (c *chinchiroBatcher) batchLoadRoomParticipant(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	var err error
	intids := array.Map(keys, func(ID dataloader.Key) uint32 {
		intid, e := idToUint32(ID.String())
		if e != nil {
			err = e
		}
		return intid
	})
	if err != nil {
		return nil
	}
	participants, err := c.chinchiroRoomUsecase.FindChinchiroRoomParticipants(model.ChinchiroRoomParticipantsQuery{
		IDs: &intids,
	})
	if err != nil {
		return nil
	}
	return array.Map(intids, func(id uint32) *dataloader.Result {
		ps := array.Find(participants.List, func(p model.ChinchiroRoomParticipant) bool {
			return p.ID == id
		})
		return &dataloader.Result{
			Data:  ps,
			Error: nil,
		}
	})
}

func (c *chinchiroBatcher) batchLoadRoomMaster(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	var err error
	intids := array.Map(keys, func(ID dataloader.Key) uint32 {
		intid, e := idToUint32(ID.String())
		if e != nil {
			err = e
		}
		return intid
	})
	if err != nil {
		return nil
	}
	masters, err := c.chinchiroRoomUsecase.FindChinchiroRoomMasters(model.ChinchiroRoomMastersQuery{
		IDs: &intids,
	})
	if err != nil {
		return nil
	}
	return array.Map(intids, func(id uint32) *dataloader.Result {
		ps := array.Find(masters, func(p model.ChinchiroRoomMaster) bool {
			return p.ID == id
		})
		return &dataloader.Result{
			Data:  ps,
			Error: nil,
		}
	})
}

func (c *chinchiroBatcher) batchLoadGame(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	var err error
	intids := array.Map(keys, func(ID dataloader.Key) uint32 {
		intid, e := idToUint32(ID.String())
		if e != nil {
			err = e
		}
		return intid
	})
	if err != nil {
		return nil
	}
	games, err := c.chinchiroGameUsecase.FindChinchiroGames(model.ChinchiroGamesQuery{
		IDs: &intids,
	})
	if err != nil {
		return nil
	}
	return array.Map(intids, func(id uint32) *dataloader.Result {
		ps := array.Find(games, func(p model.ChinchiroGame) bool {
			return p.ID == id
		})
		return &dataloader.Result{
			Data:  ps,
			Error: nil,
		}
	})
}

func (c *chinchiroBatcher) batchLoadGameParticipant(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	var err error
	intids := array.Map(keys, func(ID dataloader.Key) uint32 {
		intid, e := idToUint32(ID.String())
		if e != nil {
			err = e
		}
		return intid
	})
	if err != nil {
		return nil
	}
	participants, err := c.chinchiroGameUsecase.FindChinchiroGameParticipants(model.ChinchiroGameParticipantsQuery{
		IDs: &intids,
	})
	if err != nil {
		return nil
	}
	return array.Map(intids, func(id uint32) *dataloader.Result {
		ps := array.Find(participants.List, func(p model.ChinchiroGameParticipant) bool {
			return p.ID == id
		})
		return &dataloader.Result{
			Data:  ps,
			Error: nil,
		}
	})
}

func (c *chinchiroBatcher) batchLoadGameTurn(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	var err error
	intids := array.Map(keys, func(ID dataloader.Key) uint32 {
		intid, e := idToUint32(ID.String())
		if e != nil {
			err = e
		}
		return intid
	})
	if err != nil {
		return nil
	}
	participants, err := c.chinchiroGameUsecase.FindChinchiroGameTurns(model.ChinchiroGameTurnsQuery{
		IDs: &intids,
	})
	if err != nil {
		return nil
	}
	return array.Map(intids, func(id uint32) *dataloader.Result {
		ps := array.Find(participants.List, func(p model.ChinchiroGameTurn) bool {
			return p.ID == id
		})
		return &dataloader.Result{
			Data:  ps,
			Error: nil,
		}
	})
}

func (c *chinchiroBatcher) batchLoadGameTurnRoll(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	var err error
	intids := array.Map(keys, func(ID dataloader.Key) uint32 {
		intid, e := idToUint32(ID.String())
		if e != nil {
			err = e
		}
		return intid
	})
	if err != nil {
		return nil
	}
	participants, err := c.chinchiroGameUsecase.FindChinchiroGameTurnRolls(model.ChinchiroGameTurnRollsQuery{
		IDs: &intids,
	})
	if err != nil {
		return nil
	}
	return array.Map(intids, func(id uint32) *dataloader.Result {
		ps := array.Find(participants.List, func(p model.ChinchiroGameTurnRoll) bool {
			return p.ID == id
		})
		return &dataloader.Result{
			Data:  ps,
			Error: nil,
		}
	})
}

func (c *chinchiroBatcher) batchLoadGameTurnResult(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	var err error
	intids := array.Map(keys, func(ID dataloader.Key) uint32 {
		intid, e := idToUint32(ID.String())
		if e != nil {
			err = e
		}
		return intid
	})
	if err != nil {
		return nil
	}
	participants, err := c.chinchiroGameUsecase.FindChinchiroGameTurnResults(model.ChinchiroGameTurnResultsQuery{
		IDs: &intids,
	})
	if err != nil {
		return nil
	}
	return array.Map(intids, func(id uint32) *dataloader.Result {
		ps := array.Find(participants.List, func(p model.ChinchiroGameTurnResult) bool {
			return p.ID == id
		})
		return &dataloader.Result{
			Data:  ps,
			Error: nil,
		}
	})
}

func (p *playerBatcher) batchLoadPlayer(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	var err error
	intids := array.Map(keys, func(ID dataloader.Key) uint32 {
		intid, e := idToUint32(ID.String())
		if e != nil {
			err = e
		}
		return intid
	})
	if err != nil {
		return nil
	}
	players, err := p.playerUsecase.FindPlayers(model.PlayersQuery{
		IDs: &intids,
	})
	if err != nil {
		return nil
	}
	return array.Map(intids, func(id uint32) *dataloader.Result {
		ps := array.Find(players, func(p model.Player) bool {
			return p.ID == id
		})
		return &dataloader.Result{
			Data:  ps,
			Error: nil,
		}
	})
}
