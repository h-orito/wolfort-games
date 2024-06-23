package model

import (
	"context"
	"wolfort-games/util/array"
)

type ChinchiroGames struct {
	Count int
	List  []ChinchiroGame
}

type ChinchiroGame struct {
	ID             uint32
	RoomID         uint32
	Status         ChinchiroGameStatus
	ParticipantIDs []uint32
	TurnIDs        []uint32
}

type ChinchiroGameStatus int

const (
	ChinchiroGameStatusProgress ChinchiroGameStatus = iota
	ChinchiroGameStatusFinished
)

func (gs ChinchiroGameStatus) String() string {
	switch gs {
	case ChinchiroGameStatusProgress:
		return "Progress"
	case ChinchiroGameStatusFinished:
		return "Finished"
	default:
		return ""
	}
}

func ChinchiroGameStatusValues() []ChinchiroGameStatus {
	return []ChinchiroGameStatus{
		ChinchiroGameStatusProgress,
		ChinchiroGameStatusFinished,
	}
}

func ChinchiroGameStatusValueOf(s string) *ChinchiroGameStatus {
	return array.Find(ChinchiroGameStatusValues(), func(gs ChinchiroGameStatus) bool {
		return gs.String() == s
	})
}

type ChinchiroGamesQuery struct {
	IDs      *[]uint32
	RoomID   *uint32
	Name     *string
	Statuses *[]ChinchiroGameStatus
	Paging   *PagingQuery
}

type ChinchiroGameParticipants struct {
	Count int
	List  []ChinchiroGameParticipant
}

type ChinchiroGameParticipantsQuery struct {
	IDs     *[]uint32
	GameID  *uint32
	GameIDs *[]uint32
}

type ChinchiroGameParticipant struct {
	ID                uint32
	RoomParticipantID uint32
	Balance           int
	TurnOrder         int
}

var ChinchiroGameParticipantInitialBalance = 100000

type ChinchiroGameParticipantQuery struct {
	ID                *uint32
	GameID            *uint32
	RoomParticipantID *uint32
}

type ChinchiroGameTurns struct {
	Count int
	List  []ChinchiroGameTurn
}

type ChinchiroGameTurnsQuery struct {
	IDs      *[]uint32
	GameID   *uint32
	GameIDs  *[]uint32
	Statuses *[]ChinchiroGameTurnStatus
	Paging   *PagingQuery
}

type ChinchiroGameTurn struct {
	ID           uint32
	GameID       uint32
	DealerID     uint32
	NextRollerID *uint32
	Status       ChinchiroGameTurnStatus
	TurnNumber   int
	RollIDs      []uint32
	ResultIDs    []uint32
}

type ChinchiroGameTurnStatus int

const (
	ChinchiroGameTurnStatusBetting ChinchiroGameTurnStatus = iota
	ChinchiroGameTurnStatusRolling
	ChinchiroGameTurnStatusFinished
)

func (s ChinchiroGameTurnStatus) String() string {
	switch s {
	case ChinchiroGameTurnStatusBetting:
		return "Betting"
	case ChinchiroGameTurnStatusRolling:
		return "Rolling"
	case ChinchiroGameTurnStatusFinished:
		return "Finished"
	default:
		return ""
	}
}

func ChinchiroGameTurnStatusValues() []ChinchiroGameTurnStatus {
	return []ChinchiroGameTurnStatus{
		ChinchiroGameTurnStatusBetting,
		ChinchiroGameTurnStatusRolling,
		ChinchiroGameTurnStatusFinished,
	}
}

func ChinchiroGameTurnStatusValueOf(s string) *ChinchiroGameTurnStatus {
	return array.Find(ChinchiroGameTurnStatusValues(), func(ts ChinchiroGameTurnStatus) bool {
		return ts.String() == s
	})
}

type ChinchiroGameTurnRolls struct {
	Count int
	List  []ChinchiroGameTurnRoll
}

type ChinchiroGameTurnRollsQuery struct {
	IDs           *[]uint32
	TurnID        *uint32
	TurnIDs       *[]uint32
	ParticipantID *uint32
	Paging        *PagingQuery
}

type ChinchiroGameTurnRoll struct {
	ID            uint32
	TurnID        uint32
	ParticipantID uint32
	RollNumber    int
	DiceRoll      ChinchiroDiceRoll
}

type ChinchiroGameTurnResults struct {
	Count int
	List  []ChinchiroGameTurnResult
}

func (r ChinchiroGameTurnResults) IsRollFinished() bool {
	return array.All(r.List, func(result ChinchiroGameTurnResult) bool {
		return result.Winnings != nil
	})
}

type ChinchiroGameTurnResultsQuery struct {
	IDs     *[]uint32
	TurnID  *uint32
	TurnIDs *[]uint32
	Paging  *PagingQuery
}

type ChinchiroGameTurnResult struct {
	ID            uint32
	TurnID        uint32
	ParticipantID uint32
	BetAmount     int
	DiceRoll      *ChinchiroDiceRoll
	Winnings      *int
}

type ChinchiroGameTurnResultQuery struct {
	ID            *uint32
	TurnID        *uint32
	ParticipantID *uint32
}

type ChinchiroCombination int

const (
	ChinchiroCombinationHifumi ChinchiroCombination = iota
	ChinchiroCombinationMenashi
	ChinchiroCombinationIchinome
	ChinchiroCombinationNinome
	ChinchiroCombinationSannome
	ChinchiroCombinationYonnome
	ChinchiroCombinationGonome
	ChinchiroCombinationRokunome
	ChinchiroCombinationPinzoro
	ChinchiroCombinationNizoro
	ChinchiroCombinationSanzoro
	ChinchiroCombinationYonzoro
	ChinchiroCombinationGozoro
	ChinchiroCombinationRokuzoro
	ChinchiroCombinationShigoro
)

func (c ChinchiroCombination) String() string {
	switch c {
	case ChinchiroCombinationHifumi:
		return "Hifumi"
	case ChinchiroCombinationMenashi:
		return "Menashi"
	case ChinchiroCombinationIchinome:
		return "Ichinome"
	case ChinchiroCombinationNinome:
		return "Ninome"
	case ChinchiroCombinationSannome:
		return "Sannome"
	case ChinchiroCombinationYonnome:
		return "Yonnome"
	case ChinchiroCombinationGonome:
		return "Gonome"
	case ChinchiroCombinationRokunome:
		return "Rokunome"
	case ChinchiroCombinationPinzoro:
		return "Pinzoro"
	case ChinchiroCombinationNizoro:
		return "Nizoro"
	case ChinchiroCombinationSanzoro:
		return "Sanzoro"
	case ChinchiroCombinationYonzoro:
		return "Yonzoro"
	case ChinchiroCombinationGozoro:
		return "Gozoro"
	case ChinchiroCombinationRokuzoro:
		return "Rokuzoro"
	case ChinchiroCombinationShigoro:
		return "Shigoro"
	default:
		return ""
	}
}

func (c ChinchiroCombination) Ratio() int {
	switch c {
	case ChinchiroCombinationHifumi:
		return -2
	case ChinchiroCombinationMenashi:
		return -1
	case ChinchiroCombinationIchinome,
		ChinchiroCombinationNinome,
		ChinchiroCombinationSannome,
		ChinchiroCombinationYonnome,
		ChinchiroCombinationGonome,
		ChinchiroCombinationRokunome:
		return 1
	case ChinchiroCombinationNizoro,
		ChinchiroCombinationSanzoro,
		ChinchiroCombinationYonzoro,
		ChinchiroCombinationGozoro,
		ChinchiroCombinationRokuzoro:
		return 3
	case ChinchiroCombinationPinzoro:
		return 5
	case ChinchiroCombinationShigoro:
		return 2
	default:
		return 0
	}
}

func ChinchiroCombinationValues() []ChinchiroCombination {
	return []ChinchiroCombination{
		ChinchiroCombinationHifumi,
		ChinchiroCombinationMenashi,
		ChinchiroCombinationIchinome,
		ChinchiroCombinationNinome,
		ChinchiroCombinationSannome,
		ChinchiroCombinationYonnome,
		ChinchiroCombinationGonome,
		ChinchiroCombinationRokunome,
		ChinchiroCombinationPinzoro,
		ChinchiroCombinationNizoro,
		ChinchiroCombinationSanzoro,
		ChinchiroCombinationYonzoro,
		ChinchiroCombinationGozoro,
		ChinchiroCombinationRokuzoro,
		ChinchiroCombinationShigoro,
	}
}

func ChinchiroCombinationValueOf(s string) *ChinchiroCombination {
	return array.Find(ChinchiroCombinationValues(), func(c ChinchiroCombination) bool {
		return c.String() == s
	})
}

// -----------------------

type ChinchiroGameRepository interface {
	// game
	FindGames(query ChinchiroGamesQuery) (games ChinchiroGames, err error)
	FindGame(ID uint32) (game *ChinchiroGame, err error)
	RegisterGame(ctx context.Context, roomID uint32, game ChinchiroGame) (saved *ChinchiroGame, err error)
	UpdateGameStatus(ctx context.Context, gameID uint32, status ChinchiroGameStatus) (err error)
	// game participant
	FindGameParticipants(query ChinchiroGameParticipantsQuery) (participants ChinchiroGameParticipants, err error)
	FindGameParticipant(query ChinchiroGameParticipantQuery) (participant *ChinchiroGameParticipant, err error)
	RegisterGameParticipant(ctx context.Context, gameID uint32, participant ChinchiroGameParticipant) (saved *ChinchiroGameParticipant, err error)
	// game turn
	FindGameTurns(query ChinchiroGameTurnsQuery) (turns ChinchiroGameTurns, err error)
	FindGameTurn(ID uint32) (turn *ChinchiroGameTurn, err error)
	RegisterGameTurn(ctx context.Context, gameID uint32, turn ChinchiroGameTurn) (saved *ChinchiroGameTurn, err error)
	UpdateGameTurnStatus(ctx context.Context, turnID uint32, status ChinchiroGameTurnStatus) (err error)
	UpdateGameTurnRoller(ctx context.Context, turnID uint32, rollerID *uint32) (err error)
	// game turn roll
	FindGameTurnRolls(query ChinchiroGameTurnRollsQuery) (rolls ChinchiroGameTurnRolls, err error)
	RegisterGameTurnRoll(ctx context.Context, gameID uint32, turnID uint32, roll ChinchiroGameTurnRoll) (saved *ChinchiroGameTurnRoll, err error)
	// game turn result
	FindGameTurnResults(query ChinchiroGameTurnResultsQuery) (results ChinchiroGameTurnResults, err error)
	FindGameTurnResult(query ChinchiroGameTurnResultQuery) (result *ChinchiroGameTurnResult, err error)
	RegisterGameTurnResult(ctx context.Context, gameID uint32, turnID uint32, result ChinchiroGameTurnResult) (saved *ChinchiroGameTurnResult, err error)
	UpdateGameTurnResult(ctx context.Context, resultID uint32, result ChinchiroGameTurnResult) (saved *ChinchiroGameTurnResult, err error)
}
