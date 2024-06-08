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

func (r *chinchiroRoomResolver) games(ctx context.Context, obj *gqlmodel.ChinchiroRoom) ([]*gqlmodel.ChinchiroGame, error) {
	if obj.GameIDs == nil || len(obj.GameIDs) == 0 {
		return nil, nil
	}
	thunk := r.loaders.ChinchiroGameLoader.LoadMany(ctx, dataloader.NewKeysFromStrings(obj.GameIDs))
	c, errs := thunk()
	if errs != nil || len(errs) > 0 {
		return nil, errs[0]
	}
	return array.Map(c, func(c interface{}) *gqlmodel.ChinchiroGame {
		r := c.(*model.ChinchiroGame)
		return MapToChinchiroGame(r)
	}), nil
}

func (r *chinchiroRoomResolver) roomMasters(ctx context.Context, obj *gqlmodel.ChinchiroRoom) ([]*gqlmodel.ChinchiroRoomMaster, error) {
	if obj.RoomMasterIDs == nil || len(obj.RoomMasterIDs) == 0 {
		return nil, nil
	}
	thunk := r.loaders.ChinchiroRoomMasterLoader.LoadMany(ctx, dataloader.NewKeysFromStrings(obj.RoomMasterIDs))
	c, errs := thunk()
	if errs != nil || len(errs) > 0 {
		return nil, errs[0]
	}
	return array.Map(c, func(c interface{}) *gqlmodel.ChinchiroRoomMaster {
		r := c.(*model.ChinchiroRoomMaster)
		return MapToRoomMaster(r)
	}), nil
}

func (r *chinchiroRoomResolver) participants(ctx context.Context, obj *gqlmodel.ChinchiroRoom) ([]*gqlmodel.ChinchiroRoomParticipant, error) {
	if obj.ParticipantIDs == nil || len(obj.ParticipantIDs) == 0 {
		return nil, nil
	}
	thunk := r.loaders.ChinchiroRoomParticipantLoader.LoadMany(ctx, dataloader.NewKeysFromStrings(obj.ParticipantIDs))
	c, errs := thunk()
	if errs != nil || len(errs) > 0 {
		return nil, errs[0]
	}
	return array.Map(c, func(c interface{}) *gqlmodel.ChinchiroRoomParticipant {
		r := c.(*model.ChinchiroRoomParticipant)
		return MapToRoomParticipant(r)
	}), nil
}

func (r *mutationResolver) registerChinchiroRoom(ctx context.Context, input gqlmodel.NewChinchiroRoom) (*gqlmodel.RegisterChinchiroRoomPayload, error) {
	user := auth.GetUser(ctx)
	if user == nil {
		return nil, fmt.Errorf("user not found")
	}
	game, err := r.chinchiroRoomUsecase.RegisterChinchiroRoom(ctx, *user, model.ChinchiroRoom{
		Name: input.Name,
	})
	if err != nil {
		return nil, err
	}
	return &gqlmodel.RegisterChinchiroRoomPayload{
		ChinchiroRoom: MapToChinchiroRoom(game),
	}, nil
}

func (r *mutationResolver) registerChinchiroRoomMaster(ctx context.Context, input gqlmodel.NewChinchiroRoomMaster) (*gqlmodel.RegisterChinchiroRoomMasterPayload, error) {
	user := auth.GetUser(ctx)
	if user == nil {
		return nil, fmt.Errorf("user not found")
	}
	roomID, err := idToUint32(input.RoomID)
	if err != nil {
		return nil, err
	}
	playerID, err := idToUint32(input.PlayerID)
	if err != nil {
		return nil, err
	}
	master, err := r.chinchiroRoomUsecase.RegisterChinchiroRoomMaster(ctx, *user, roomID, playerID)
	if err != nil {
		return nil, err
	}
	return &gqlmodel.RegisterChinchiroRoomMasterPayload{
		ChinchiroRoomMaster: MapToRoomMaster(&master),
	}, nil
}

func (r *mutationResolver) deleteChinchiroRoomMaster(ctx context.Context, input gqlmodel.DeleteChinchiroRoomMaster) (*gqlmodel.DeleteChinchiroRoomMasterPayload, error) {
	user := auth.GetUser(ctx)
	if user == nil {
		return nil, fmt.Errorf("user not found")
	}
	roomID, err := idToUint32(input.RoomID)
	if err != nil {
		return nil, err
	}
	masterID, err := idToUint32(input.MasterID)
	if err != nil {
		return nil, err
	}
	err = r.chinchiroRoomUsecase.DeleteChinchiroRoomMaster(ctx, *user, roomID, masterID)
	if err != nil {
		return nil, err
	}
	return &gqlmodel.DeleteChinchiroRoomMasterPayload{
		Ok: true,
	}, nil
}

func (r *mutationResolver) updateChinchiroRoomStatus(ctx context.Context, input gqlmodel.UpdateChinchiroRoomStatus) (*gqlmodel.UpdateChinchiroRoomStatusPayload, error) {
	user := auth.GetUser(ctx)
	if user == nil {
		return nil, fmt.Errorf("user not found")
	}
	roomID, err := idToUint32(input.RoomID)
	if err != nil {
		return nil, err
	}
	status := model.ChinchiroRoomStatusValueOf(input.Status.String())
	err = r.chinchiroRoomUsecase.UpdateChinchiroRoomStatus(ctx, *user, roomID, *status)
	if err != nil {
		return nil, err
	}
	return &gqlmodel.UpdateChinchiroRoomStatusPayload{
		Ok: true,
	}, nil
}

func (r *mutationResolver) updateChinchiroRoomSettings(ctx context.Context, input gqlmodel.UpdateChinchiroRoomSettings) (*gqlmodel.UpdateChinchiroRoomSettingsPayload, error) {
	user := auth.GetUser(ctx)
	if user == nil {
		return nil, fmt.Errorf("user not found")
	}
	roomID, err := idToUint32(input.RoomID)
	if err != nil {
		return nil, err
	}
	settings := model.ChinchiroRoomSettings{
		Password: model.ChinchiroRoomPasswordSettings{
			Password: nil, // TODO: input.Password,
		},
	}
	err = r.chinchiroRoomUsecase.UpdateChinchiroRoomSettings(ctx, *user, roomID, settings)
	if err != nil {
		return nil, err
	}
	return &gqlmodel.UpdateChinchiroRoomSettingsPayload{
		Ok: true,
	}, nil
}

func (r *mutationResolver) registerChinchiroRoomParticipant(ctx context.Context, input gqlmodel.NewChinchiroRoomParticipant) (*gqlmodel.RegisterChinchiroRoomParticipantPayload, error) {
	user := auth.GetUser(ctx)
	if user == nil {
		return nil, fmt.Errorf("user not found")
	}
	roomID, err := idToUint32(input.RoomID)
	if err != nil {
		return nil, err
	}
	participant, err := r.chinchiroRoomUsecase.ParticipateChinchiroRoom(ctx, *user, roomID, model.ChinchiroRoomParticipant{
		Name: input.Name,
	}, input.Password)
	if err != nil {
		return nil, err
	}
	return &gqlmodel.RegisterChinchiroRoomParticipantPayload{
		ChinchiroRoomParticipant: MapToRoomParticipant(participant),
	}, nil
}

func (r *mutationResolver) updateChinchiroRoomParticipant(ctx context.Context, input gqlmodel.UpdateChinchiroRoomParticipant) (*gqlmodel.UpdateChinchiroRoomParticipantPayload, error) {
	user := auth.GetUser(ctx)
	if user == nil {
		return nil, fmt.Errorf("user not found")
	}
	participantID, err := idToUint32(input.ParticipantID)
	if err != nil {
		return nil, err
	}
	err = r.chinchiroRoomUsecase.UpdateChinchiroRoomParticipant(ctx, *user, participantID, model.ChinchiroRoomParticipant{
		Name: input.Name,
	})
	if err != nil {
		return nil, err
	}
	return &gqlmodel.UpdateChinchiroRoomParticipantPayload{
		Ok: true,
	}, nil
}

func (r *mutationResolver) deleteChinchiroRoomParticipant(ctx context.Context, input gqlmodel.DeleteChinchiroRoomParticipant) (*gqlmodel.DeleteChinchiroRoomParticipantPayload, error) {
	user := auth.GetUser(ctx)
	if user == nil {
		return nil, fmt.Errorf("user not found")
	}
	participantID, err := idToUint32(input.ParticipantID)
	if err != nil {
		return nil, err
	}
	err = r.chinchiroRoomUsecase.LeaveChinchiroRoom(ctx, *user, participantID)
	if err != nil {
		return nil, err
	}
	return &gqlmodel.DeleteChinchiroRoomParticipantPayload{
		Ok: true,
	}, nil
}

// LeaveChinchiroRoom is the resolver for the leaveChinchiroRoom field.
func (r *mutationResolver) leaveChinchiroRoom(ctx context.Context, input gqlmodel.LeaveChinchiroRoom) (*gqlmodel.LeaveChinchiroRoomPayload, error) {
	panic(fmt.Errorf("not implemented: LeaveChinchiroRoom - leaveChinchiroRoom"))
}

// ChinchiroRooms is the resolver for the chinchiroRooms field.
func (r *queryResolver) chinchiroRooms(ctx context.Context, query gqlmodel.ChinchiroRoomsQuery) ([]*gqlmodel.SimpleChinchiroRoom, error) {
	queryIDs, err := idsToUint32s(query.Ids)
	if err != nil {
		return nil, err
	}
	var statuses *[]model.ChinchiroRoomStatus
	if query.Statuses != nil {
		ses := array.Map(query.Statuses, func(s gqlmodel.ChinchiroRoomStatus) model.ChinchiroRoomStatus {
			return *model.ChinchiroRoomStatusValueOf(s.String())
		})
		statuses = &ses
	}

	rooms, err := r.chinchiroRoomUsecase.FindChinchiroRooms(model.ChinchiroRoomsQuery{
		IDs:      &queryIDs,
		Name:     query.Name,
		Statuses: statuses,
		Paging:   query.Paging.MapToPagingQuery(),
	})
	if err != nil {
		return nil, err
	}
	return array.Map(rooms, func(r model.ChinchiroRoom) *gqlmodel.SimpleChinchiroRoom {
		return MapToSimpleChinchiroRoom(&r)
	}), nil
}

// ChinchiroRoom is the resolver for the chinchiroRoom field.
func (r *queryResolver) chinchiroRoom(ctx context.Context, roomID string) (*gqlmodel.ChinchiroRoom, error) {
	rID, err := idToUint32(roomID)
	if err != nil {
		return nil, err
	}
	room, err := r.chinchiroRoomUsecase.FindChinchiroRoom(rID)
	if err != nil {
		return nil, err
	}
	return MapToChinchiroRoom(room), nil
}

// MyChinchiroRoomParticipant is the resolver for the myChinchiroRoomParticipant field.
func (r *queryResolver) myChinchiroRoomParticipant(ctx context.Context, roomID string) (*gqlmodel.ChinchiroRoomParticipant, error) {
	user := auth.GetUser(ctx)
	if user == nil {
		return nil, nil
	}
	rID, err := idToUint32(roomID)
	if err != nil {
		return nil, err
	}
	participant, err := r.chinchiroRoomUsecase.FindMyChinchiroRoomParticipant(*user, rID)
	if err != nil {
		return nil, err
	}
	if participant == nil {
		return nil, nil
	}
	return MapToRoomParticipant(participant), nil
}
