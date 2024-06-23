package gqlmodel

// custom modelを定義

type ChinchiroRoom struct {
	ID             string                      `json:"id"`
	Name           string                      `json:"name"`
	Status         ChinchiroRoomStatus         `json:"status"`
	RoomMasterIDs  []string                    `json:"roomMasterIds"`
	RoomMasters    []*ChinchiroRoomMaster      `json:"roomMasters"`
	ParticipantIDs []string                    `json:"participantIds"`
	Participants   []*ChinchiroRoomParticipant `json:"participants"`
	GameIDs        []string                    `json:"gameIds"`
	Games          []*ChinchiroGame            `json:"games"`
	Settings       *ChinchiroRoomSettings      `json:"settings"`
}

type ChinchiroRoomParticipant struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	PlayerID string  `json:"playerId"`
	Player   *Player `json:"player"`
	IsGone   bool    `json:"isGone"`
}

type ChinchiroRoomMaster struct {
	ID       string  `json:"id"`
	PlayerID string  `json:"playerId"`
	Player   *Player `json:"player"`
}

type ChinchiroGame struct {
	ID             string                      `json:"id"`
	Status         ChinchiroGameStatus         `json:"status"`
	ParticipantIDs []string                    `json:"participantIds"`
	Participants   []*ChinchiroGameParticipant `json:"participants"`
	TurnIDs        []string                    `json:"turnIds"`
	Turns          []*ChinchiroGameTurn        `json:"turns"`
}

type ChinchiroGameParticipant struct {
	ID                string                    `json:"id"`
	RoomParticipantID string                    `json:"roomParticipantId"`
	RoomParticipant   *ChinchiroRoomParticipant `json:"roomParticipant"`
	Balance           int                       `json:"balance"`
	TurnOrder         int                       `json:"turnOrder"`
}

type ChinchiroGameTurn struct {
	ID           string                                `json:"id"`
	DealerID     string                                `json:"dealerId"`
	Dealer       *ChinchiroGameParticipant             `json:"dealer"`
	NextRollerID *string                               `json:"nextRollerId"`
	NextRoller   *ChinchiroGameParticipant             `json:"nextRoller"`
	Status       ChinchiroGameTurnStatus               `json:"status"`
	TurnNumber   int                                   `json:"turnNumber"`
	RollIDs      []string                              `json:"rollIds"`
	Rolls        []*ChinchiroGameTurnParticipantRoll   `json:"rolls"`
	ResultIDs    []string                              `json:"resultIds"`
	Results      []*ChinchiroGameTurnParticipantResult `json:"results"`
}

type ChinchiroGameTurnParticipantRoll struct {
	ID            string                    `json:"id"`
	TurnID        string                    `json:"turnId"`
	Turn          *ChinchiroGameTurn        `json:"turn"`
	ParticipantID string                    `json:"participantId"`
	Participant   *ChinchiroGameParticipant `json:"participant"`
	RollNumber    int                       `json:"rollNumber"`
	Dice1         int                       `json:"dice1"`
	Dice2         int                       `json:"dice2"`
	Dice3         int                       `json:"dice3"`
}

type ChinchiroGameTurnParticipantResult struct {
	ID            string                    `json:"id"`
	GameTurnID    string                    `json:"gameTurnId"`
	Turn          *ChinchiroGameTurn        `json:"turn"`
	ParticipantID string                    `json:"participantId"`
	Participant   *ChinchiroGameParticipant `json:"participant"`
	Dice1         int                       `json:"dice1"`
	Dice2         int                       `json:"dice2"`
	Dice3         int                       `json:"dice3"`
	Combination   ChinchiroCombination      `json:"combination"`
	Winnings      int                       `json:"winnings"`
}
