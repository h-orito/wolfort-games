package db_test

import (
	"chat-role-play/domain/model"
	db "chat-role-play/infrastructure/rdb"
	"context"
	"testing"
)

func newPlayerRepository() model.PlayerRepository {
	database := NewTestDB()
	return db.NewPlayerRepository(&database)
}

func TestFind(t *testing.T) {
	repo := newPlayerRepository()
	got, err := repo.Find(1)
	if err != nil {
		t.Errorf("failed to find player: %s", err)
	}
	if got == nil {
		t.Errorf("got nil, want player")
	}
}

func TestFindByName(t *testing.T) {
	repo := newPlayerRepository()
	got, err := repo.FindByName("player name 1")
	if err != nil {
		t.Errorf("failed to find player: %s", err)
	}
	if got == nil {
		t.Errorf("got nil, want player")
	}
}

func TestFindByUserName(t *testing.T) {
	repo := newPlayerRepository()
	got, err := repo.FindByUserName("user name 1")
	if err != nil {
		t.Errorf("failed to find player: %s", err)
	}
	if got == nil {
		t.Errorf("got nil, want player")
	}
}

func TestSave(t *testing.T) {
	database := NewTestDB()
	transaction := NewTestTransaction(database.Connection)
	transaction.DoInTx(context.Background(), func(ctx context.Context) (interface{}, error) {
		repo := db.NewPlayerRepository(&database)
		player := model.Player{
			ID:   1,
			Name: "test player 2",
		}
		got, err := repo.Save(ctx, &player)
		if err != nil {
			t.Errorf("failed to save player: %s", err)
		}
		if got == nil {
			t.Errorf("got nil, want player")
		}
		if got.Name != player.Name {
			t.Errorf("got %s, want %s", got.Name, player.Name)
		}
		return got, nil
	})
}
