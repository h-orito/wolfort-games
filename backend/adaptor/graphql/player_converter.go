package graphql

import (
	"wolfort-games/domain/model"
	"wolfort-games/middleware/graph/gqlmodel"
	"wolfort-games/util/array"
)

func MapToPlayer(
	p *model.Player,
	authorities []model.PlayerAuthority,
) *gqlmodel.Player {
	if p == nil {
		return nil
	}
	return &gqlmodel.Player{
		ID:   intIdToBase64(p.ID, "Player"),
		Name: p.Name,
		AuthorityCodes: array.Map(authorities, func(a model.PlayerAuthority) string {
			return a.String()
		}),
	}
}
