package db_test

import (
	"context"
	"testing"
	"wolfort-games/domain/model"
	db "wolfort-games/infrastructure/rdb"
)

func newChinchiroRoomRepository() model.ChinchiroRoomRepository {
	database := NewTestDB()
	return db.NewChinchiroRoomRepository(&database)
}

func TestFindChinchiroRooms(t *testing.T) {
	repo := newChinchiroRoomRepository()
	got, err := repo.FindRooms(model.ChinchiroRoomsQuery{
		IDs: &[]uint32{1, 2},
	})
	if err != nil {
		t.Errorf("failed to find chinchiro rooms: %s", err)
	}
	wantSize := 1
	if len(got) != wantSize {
		t.Errorf("got %d chinchiro rooms, want %d", len(got), wantSize)
	}
}

func TestFindChinchiroRoom(t *testing.T) {
	repo := newChinchiroRoomRepository()
	got, err := repo.FindRoom(1)
	if err != nil {
		t.Errorf("failed to find chinchiro room: %s", err)
	}
	if got == nil {
		t.Errorf("got nil, want room")
	}
}

func TestRegisterChinchiroRoom(t *testing.T) {
	database := NewTestDB()
	transaction := NewTestTransaction(database.Connection)
	transaction.DoInTx(context.Background(), func(ctx context.Context) (interface{}, error) {
		repo := db.NewChinchiroRoomRepository(&database)
		password := "password"
		r := model.ChinchiroRoom{
			Name:   "name",
			Status: model.ChinchiroRoomStatusOpened,
			Settings: model.ChinchiroRoomSettings{
				Password: model.ChinchiroRoomPasswordSettings{
					Password: &password,
				},
			},
		}
		got, err := repo.RegisterRoom(ctx, r)
		if err != nil {
			t.Errorf("failed to register: %s", err)
		}
		if got == nil {
			t.Errorf("got nil, want room")
		}
		return nil, nil
	})
}

func TestUpdateChinchiroRoom(t *testing.T) {
	database := NewTestDB()
	transaction := NewTestTransaction(database.Connection)
	transaction.DoInTx(context.Background(), func(ctx context.Context) (interface{}, error) {
		repo := db.NewChinchiroRoomRepository(&database)
		r := model.ChinchiroRoom{
			ID:     1,
			Name:   "name",
			Status: model.ChinchiroRoomStatusOpened,
			Settings: model.ChinchiroRoomSettings{
				Password: model.ChinchiroRoomPasswordSettings{
					Password: nil,
				},
			},
		}
		err := repo.UpdateRoom(ctx, r)
		if err != nil {
			t.Errorf("failed to update: %s", err)
		}
		return nil, nil
	})
}

func TestFindChinchiroRoomMasters(t *testing.T) {
	repo := newChinchiroRoomRepository()
	got, err := repo.FindRoomMasters(model.ChinchiroRoomMastersQuery{
		IDs: &[]uint32{1},
	})
	if err != nil {
		t.Errorf("failed to find chinchiro room masters: %s", err)
	}
	wantSize := 1
	if len(got) != wantSize {
		t.Errorf("got %d chinchiro room masters, want %d", len(got), wantSize)
	}
}

func TestRegisterChinchiroRoomMaster(t *testing.T) {
	database := NewTestDB()
	transaction := NewTestTransaction(database.Connection)
	transaction.DoInTx(context.Background(), func(ctx context.Context) (interface{}, error) {
		repo := db.NewChinchiroRoomRepository(&database)
		r := model.ChinchiroRoomMaster{
			PlayerID: 2,
		}
		_, err := repo.RegisterRoomMaster(ctx, 1, r)
		if err != nil {
			t.Errorf("failed to register: %s", err)
		}
		return nil, nil
	})
}

func TestDeleteChinchiroRoomMaster(t *testing.T) {
	database := NewTestDB()
	transaction := NewTestTransaction(database.Connection)
	transaction.DoInTx(context.Background(), func(ctx context.Context) (interface{}, error) {
		repo := db.NewChinchiroRoomRepository(&database)
		err := repo.DeleteRoomMaster(ctx, 1)
		if err != nil {
			t.Errorf("failed to delete: %s", err)
		}
		return nil, nil
	})
}
