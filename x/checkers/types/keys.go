package types

import "time"

const (
	// ModuleName defines the module name
	ModuleName = "checkers"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_checkers"
)

var (
	ParamsKey = []byte("p_checkers")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	SystemInfoKey = "SystemInfo/value/"
)

const (
    GameCreatedEventType      = "new-game-created" // Indicates what event type to listen to
    GameCreatedEventCreator   = "creator"          // Subsidiary information
    GameCreatedEventGameIndex = "game-index"       // What game is relevant
    GameCreatedEventBlack     = "black"            // Is it relevant to me?
    GameCreatedEventRed       = "red"              // Is it relevant to me?
	GameCreatedEventWager     = "wager"
)

const (
	MovePlayedEventType      = "move-played"
	MovePlayedEventCreator   = "creator"
	MovePlayedEventGameIndex = "game-index"
	MovePlayedEventCapturedX = "captured-x"
	MovePlayedEventCapturedY = "captured-y"
	MovePlayedEventWinner    = "winner"
	MovePlayedEventBoard     = "board"
)

const (
	MaxTurnDuration = time.Duration(5 * 60 * 1000_000_000) // 5 minutes
	DeadlineLayout  = "2006-01-02 15:04:05.999999999 +0000 UTC"
)

const (
    NoFifoIndex = "-1"
)

const (
	GameForfeitedEventType      = "game-forfeited"
	GameForfeitedEventGameIndex = "game-index"
	GameForfeitedEventWinner    = "winner"
	GameForfeitedEventBoard     = "board"
)
