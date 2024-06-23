package graphql

import (
	"wolfort-games/domain/model"
	"wolfort-games/middleware/graph/gqlmodel"
	"wolfort-games/util/array"
)

func MapToChinchiroRoom(
	r *model.ChinchiroRoom,
) *gqlmodel.ChinchiroRoom {
	if r == nil {
		return nil
	}
	return &gqlmodel.ChinchiroRoom{
		ID:     intIdToBase64(r.ID, "ChinchiroRoom"),
		Name:   r.Name,
		Status: gqlmodel.ChinchiroRoomStatus(r.Status.String()),
		RoomMasterIDs: array.Map(r.RoomMasterIDs, func(id uint32) string {
			return intIdToBase64(id, "ChinchiroRoomMaster")
		}),
		ParticipantIDs: array.Map(r.ParticipantIDs, func(id uint32) string {
			return intIdToBase64(id, "ChinchiroRoomParticipant")
		}),
		GameIDs: array.Map(r.GameIDs, func(id uint32) string {
			return intIdToBase64(id, "ChinchiroGame")
		}),
		Settings: MapToChinchiroRoomSettings(r.Settings),
	}
}

func MapToSimpleChinchiroRoom(
	r *model.ChinchiroRoom,
) *gqlmodel.SimpleChinchiroRoom {
	if r == nil {
		return nil
	}
	return &gqlmodel.SimpleChinchiroRoom{
		ID:                intIdToBase64(r.ID, "ChinchiroRoom"),
		Name:              r.Name,
		Status:            gqlmodel.ChinchiroRoomStatus(r.Status.String()),
		ParticipantCounts: len(r.ParticipantIDs),
		Settings:          MapToChinchiroRoomSettings(r.Settings),
	}
}

func MapToChinchiroRoomSettings(
	s model.ChinchiroRoomSettings,
) *gqlmodel.ChinchiroRoomSettings {
	return &gqlmodel.ChinchiroRoomSettings{
		Dummy: "dummy",
	}
}

func MapToRoomMaster(
	p *model.ChinchiroRoomMaster,
) *gqlmodel.ChinchiroRoomMaster {
	if p == nil {
		return nil
	}
	return &gqlmodel.ChinchiroRoomMaster{
		ID:       intIdToBase64(p.ID, "ChinchiroRoomMaster"),
		PlayerID: intIdToBase64(p.PlayerID, "Player"),
	}
}

func MapToRoomParticipant(
	p *model.ChinchiroRoomParticipant,
) *gqlmodel.ChinchiroRoomParticipant {
	if p == nil {
		return nil
	}
	return &gqlmodel.ChinchiroRoomParticipant{
		ID:       intIdToBase64(p.ID, "ChinchiroRoomParticipant"),
		Name:     p.Name,
		PlayerID: intIdToBase64(p.PlayerID, "Player"),
		IsGone:   p.IsGone,
	}
}
