package graphql

import (
	"chat-role-play/adaptor/auth"
	"chat-role-play/domain/model"
	"chat-role-play/middleware/graph/gqlmodel"
	"chat-role-play/util/array"
	"context"
	"fmt"

	"github.com/graph-gophers/dataloader"
)

func (r *chinchiroRoomMasterResolver) player(ctx context.Context, obj *gqlmodel.ChinchiroRoomMaster) (*gqlmodel.Player, error) {
	thunk := r.loaders.PlayerLoader.Load(ctx, dataloader.StringKey(obj.PlayerID))
	p, err := thunk()
	if err != nil {
		return nil, err
	}
	player := p.(*model.Player)
	return MapToPlayer(player, []model.PlayerAuthority{}), nil
}

func (r *chinchiroRoomParticipantResolver) player(ctx context.Context, obj *gqlmodel.ChinchiroRoomParticipant) (*gqlmodel.Player, error) {
	thunk := r.loaders.PlayerLoader.Load(ctx, dataloader.StringKey(obj.PlayerID))
	p, err := thunk()
	if err != nil {
		return nil, err
	}
	player := p.(*model.Player)
	return MapToPlayer(player, []model.PlayerAuthority{}), nil
}

func (r *queryResolver) players(ctx context.Context, query gqlmodel.PlayersQuery) ([]*gqlmodel.Player, error) {
	var intids *[]uint32
	var err error
	if query.Ids != nil {
		ids := array.Map(query.Ids, func(id string) uint32 {
			intid, e := idToUint32(id)
			if e != nil {
				err = e
			}
			return intid
		})
		if err != nil {
			return nil, err
		}
		intids = &ids
	}
	var paging *model.PagingQuery
	if query.Paging != nil {
		paging = &model.PagingQuery{
			PageSize:   query.Paging.PageSize,
			PageNumber: query.Paging.PageNumber,
			Desc:       query.Paging.IsDesc,
			Latest:     query.Paging.IsLatest,
		}
	}
	players, err := r.playerUsecase.FindPlayers(model.PlayersQuery{
		IDs:    intids,
		Name:   query.Name,
		Paging: paging,
	})
	if err != nil {
		return nil, err
	}
	return array.Map(players, func(p model.Player) *gqlmodel.Player {
		return MapToPlayer(&p, []model.PlayerAuthority{})
	}), nil
}

func (r *queryResolver) player(ctx context.Context, id string) (*gqlmodel.Player, error) {
	playerID, err := idToUint32(id)
	if err != nil {
		return nil, err
	}
	p, err := r.playerUsecase.Find(playerID)
	if err != nil {
		return nil, err
	}
	authorities, err := r.playerUsecase.FindAuthorities(playerID)
	if err != nil {
		return nil, err
	}
	return MapToPlayer(p, authorities), nil
}

func (r *queryResolver) myPlayer(ctx context.Context) (*gqlmodel.Player, error) {
	user := auth.GetUser(ctx)
	if user == nil {
		return nil, fmt.Errorf("user not found")
	}
	p, err := r.playerUsecase.FindByUserName(user.UserName)
	if err != nil {
		return nil, err
	}
	authorities, err := r.playerUsecase.FindAuthorities(p.ID)
	if err != nil {
		return nil, err
	}
	return MapToPlayer(p, authorities), nil
}
