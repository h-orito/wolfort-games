package db_test

import (
	"context"
	"testing"
	"wolfort-games/domain/model"
	db "wolfort-games/infrastructure/rdb"
)

func newChinchiroGameRepository() model.ChinchiroGameRepository {
	database := NewTestDB()
	return db.NewChinchiroGameRepository(&database)
}

func TestFindChinchiroGames(t *testing.T) {
	repo := newChinchiroGameRepository()
	got, err := repo.FindGames(model.ChinchiroGamesQuery{
		IDs: &[]uint32{1, 2},
	})
	if err != nil {
		t.Errorf("failed to find chinchiro games: %s", err)
	}
	wantSize := 2
	if len(got.List) != wantSize {
		t.Errorf("got %d chinchiro games, want %d", len(got.List), wantSize)
	}
}

func TestFindChinchiroGame(t *testing.T) {
	repo := newChinchiroGameRepository()
	got, err := repo.FindGame(1)
	if err != nil {
		t.Errorf("failed to find chinchiro game: %s", err)
	}
	if got == nil {
		t.Errorf("got nil, want game")
	}
}

func TestRegisterChinchiroGame(t *testing.T) {
	database := NewTestDB()
	transaction := NewTestTransaction(database.Connection)
	transaction.DoInTx(context.Background(), func(ctx context.Context) (interface{}, error) {
		repo := db.NewChinchiroGameRepository(&database)
		g := model.ChinchiroGame{
			Status: model.ChinchiroGameStatusProgress,
		}
		got, err := repo.RegisterGame(ctx, 1, g)
		if err != nil {
			t.Errorf("failed to register: %s", err)
		}
		if got == nil {
			t.Errorf("got nil, want game")
		}
		return nil, nil
	})
}

func TestUpdateChinchiroGameStatus(t *testing.T) {
	database := NewTestDB()
	transaction := NewTestTransaction(database.Connection)
	transaction.DoInTx(context.Background(), func(ctx context.Context) (interface{}, error) {
		repo := db.NewChinchiroGameRepository(&database)
		err := repo.UpdateGameStatus(ctx, 1, model.ChinchiroGameStatusFinished)
		if err != nil {
			t.Errorf("failed to update: %s", err)
		}
		return nil, nil
	})
}

func TestFindChinchiroGameParticipants(t *testing.T) {
	repo := newChinchiroGameRepository()
	got, err := repo.FindGameParticipants(model.ChinchiroGameParticipantsQuery{
		GameIDs: &[]uint32{1},
	})
	if err != nil {
		t.Errorf("failed to find chinchiro game participants: %s", err)
	}
	wantSize := 2
	if len(got.List) != wantSize {
		t.Errorf("got %d chinchiro game participants, want %d", len(got.List), wantSize)
	}
}

func TestFindChinchiroGameParticipant(t *testing.T) {
	repo := newChinchiroGameRepository()
	var pID uint32 = 1
	got, err := repo.FindGameParticipant(model.ChinchiroGameParticipantQuery{
		ID: &pID,
	})
	if err != nil {
		t.Errorf("failed to find chinchiro game participant: %s", err)
	}
	if got == nil {
		t.Errorf("got nil, want participant")
	}
}

func TestRegisterChinchiroGameParticipant(t *testing.T) {
	database := NewTestDB()
	transaction := NewTestTransaction(database.Connection)
	transaction.DoInTx(context.Background(), func(ctx context.Context) (interface{}, error) {
		repo := db.NewChinchiroGameRepository(&database)
		p := model.ChinchiroGameParticipant{
			RoomParticipantID: 1,
			Balance:           1000,
			TurnOrder:         1,
		}
		got, err := repo.RegisterGameParticipant(ctx, 3, p)
		if err != nil {
			t.Errorf("failed to register: %s", err)
		}
		if got == nil {
			t.Errorf("got nil, want participant")
		}
		return nil, nil
	})
}

func TestFindChinchiroGameTurns(t *testing.T) {
	repo := newChinchiroGameRepository()
	got, err := repo.FindGameTurns(model.ChinchiroGameTurnsQuery{
		GameIDs: &[]uint32{1},
	})
	if err != nil {
		t.Errorf("failed to find chinchiro game turns: %s", err)
	}
	wantSize := 2
	if len(got.List) != wantSize {
		t.Errorf("got %d chinchiro game turns, want %d", len(got.List), wantSize)
	}
}

func TestFindChinchiroGameTurn(t *testing.T) {
	repo := newChinchiroGameRepository()
	got, err := repo.FindGameTurn(1)
	if err != nil {
		t.Errorf("failed to find chinchiro game turn: %s", err)
	}
	if got == nil {
		t.Errorf("got nil, want turn")
	}
}

func TestRegisterChinchiroGameTurn(t *testing.T) {
	database := NewTestDB()
	transaction := NewTestTransaction(database.Connection)
	transaction.DoInTx(context.Background(), func(ctx context.Context) (interface{}, error) {
		repo := db.NewChinchiroGameRepository(&database)
		turn := model.ChinchiroGameTurn{
			DealerID:   1,
			Status:     model.ChinchiroGameTurnStatusBetting,
			TurnNumber: 3,
		}
		got, err := repo.RegisterGameTurn(ctx, 1, turn)
		if err != nil {
			t.Errorf("failed to register: %s", err)
		}
		if got == nil {
			t.Errorf("got nil, want turn")
		}
		return nil, nil
	})
}

func TestUpdateChinchiroGameTurnStatus(t *testing.T) {
	database := NewTestDB()
	transaction := NewTestTransaction(database.Connection)
	transaction.DoInTx(context.Background(), func(ctx context.Context) (interface{}, error) {
		repo := db.NewChinchiroGameRepository(&database)
		err := repo.UpdateGameTurnStatus(ctx, 1, model.ChinchiroGameTurnStatusFinished)
		if err != nil {
			t.Errorf("failed to update: %s", err)
		}
		return nil, nil
	})
}

func TestFindChinchiroGameTurnRolls(t *testing.T) {
	repo := newChinchiroGameRepository()
	got, err := repo.FindGameTurnRolls(model.ChinchiroGameTurnRollsQuery{
		TurnIDs: &[]uint32{1},
	})
	if err != nil {
		t.Errorf("failed to find chinchiro game turn rolls: %s", err)
	}
	wantSize := 3
	if len(got.List) != wantSize {
		t.Errorf("got %d chinchiro game turn rolls, want %d", got.Count, wantSize)
	}
}

func TestRegisterChinchiroGameTurnRoll(t *testing.T) {
	database := NewTestDB()
	transaction := NewTestTransaction(database.Connection)
	transaction.DoInTx(context.Background(), func(ctx context.Context) (interface{}, error) {
		repo := db.NewChinchiroGameRepository(&database)
		roll := model.ChinchiroGameTurnRoll{
			ParticipantID: 1,
			RollNumber:    1,
			DiceRoll: model.ChinchiroDiceRoll{
				Dice1: 1,
				Dice2: 2,
				Dice3: 3,
			},
		}
		got, err := repo.RegisterGameTurnRoll(ctx, 2, 3, roll)
		if err != nil {
			t.Errorf("failed to register: %s", err)
		}
		if got == nil {
			t.Errorf("got nil, want roll")
		}
		return nil, nil
	})
}

func TestFindChinchiroGameTurnResults(t *testing.T) {
	repo := newChinchiroGameRepository()
	got, err := repo.FindGameTurnResults(model.ChinchiroGameTurnResultsQuery{
		TurnIDs: &[]uint32{1},
	})
	if err != nil {
		t.Errorf("failed to find chinchiro game turn results: %s", err)
	}
	wantSize := 2
	if got.Count != wantSize {
		t.Errorf("got %d chinchiro game turn results, want %d", got.Count, wantSize)
	}
}

func TestUpdateChinchiroGameTurnResult(t *testing.T) {
	database := NewTestDB()
	transaction := NewTestTransaction(database.Connection)
	transaction.DoInTx(context.Background(), func(ctx context.Context) (interface{}, error) {
		repo := db.NewChinchiroGameRepository(&database)
		var winnings int = -20000
		res := model.ChinchiroGameTurnResult{
			DiceRoll: &model.ChinchiroDiceRoll{
				Dice1: 1,
				Dice2: 2,
				Dice3: 3,
			},
			Winnings: &winnings,
		}
		_, err := repo.UpdateGameTurnResult(ctx, 3, res)
		if err != nil {
			t.Errorf("failed to update: %s", err)
		}
		return nil, nil
	})
}
