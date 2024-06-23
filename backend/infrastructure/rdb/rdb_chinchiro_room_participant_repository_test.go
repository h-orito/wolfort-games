package db_test

import (
	"context"
	"testing"
	"wolfort-games/domain/model"
	db "wolfort-games/infrastructure/rdb"
)

func newChinchiroGameParticipantRepository() model.ChinchiroRoomParticipantRepository {
	database := NewTestDB()
	return db.NewChinchiroRoomParticipantRepository(&database)
}

func TestFindChinchiroRoomParticipants(t *testing.T) {
	repo := newChinchiroGameParticipantRepository()
	got, err := repo.FindRoomParticipants(model.ChinchiroRoomParticipantsQuery{
		RoomIDs: &[]uint32{1},
	})
	if err != nil {
		t.Errorf("failed to find chinchiro room participants: %s", err)
	}
	wantSize := 2
	if len(got.List) != wantSize {
		t.Errorf("got %d chinchiro room participants, want %d", len(got.List), wantSize)
	}
}

func TestFindChinchiroRoomParticipant(t *testing.T) {
	repo := newChinchiroGameParticipantRepository()
	var id uint32 = 1
	got, err := repo.FindRoomParticipant(model.ChinchiroRoomParticipantQuery{
		ID: &id,
	})
	if err != nil {
		t.Errorf("failed to find chinchiro room participant: %s", err)
	}
	if got == nil {
		t.Errorf("got nil, want participant")
	}
}

func TestRegisterChinchiroRoomParticipant(t *testing.T) {
	database := NewTestDB()
	transaction := NewTestTransaction(database.Connection)
	transaction.DoInTx(context.Background(), func(ctx context.Context) (interface{}, error) {
		repo := db.NewChinchiroRoomParticipantRepository(&database)
		p := model.ChinchiroRoomParticipant{
			Name:     "test",
			PlayerID: 1,
			IsGone:   false,
		}
		got, err := repo.RegisterRoomParticipant(ctx, 1, p)
		if err != nil {
			t.Errorf("failed to register: %s", err)
		}
		if got == nil {
			t.Errorf("got nil, want game")
		}
		return nil, nil
	})
}

func TestUpdateChinchiroRoomParticipant(t *testing.T) {
	database := NewTestDB()
	transaction := NewTestTransaction(database.Connection)
	transaction.DoInTx(context.Background(), func(ctx context.Context) (interface{}, error) {
		repo := db.NewChinchiroRoomParticipantRepository(&database)
		p := model.ChinchiroRoomParticipant{
			ID:       1,
			Name:     "test",
			PlayerID: 1,
			IsGone:   true,
		}
		err := repo.UpdateRoomParticipant(ctx, p)
		if err != nil {
			t.Errorf("failed to update: %s", err)
		}
		return nil, nil
	})
}

func TestDeleteChinchiroRoomParticipant(t *testing.T) {
	database := NewTestDB()
	transaction := NewTestTransaction(database.Connection)
	transaction.DoInTx(context.Background(), func(ctx context.Context) (interface{}, error) {
		repo := db.NewChinchiroRoomParticipantRepository(&database)
		err := repo.DeleteRoomParticipant(ctx, 1)
		if err != nil {
			t.Errorf("failed to delete: %s", err)
		}
		return nil, nil
	})
}
