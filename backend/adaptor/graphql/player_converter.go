package graphql

import (
	"chat-role-play/domain/model"
	"chat-role-play/middleware/graph/gqlmodel"
	"chat-role-play/util/array"
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
