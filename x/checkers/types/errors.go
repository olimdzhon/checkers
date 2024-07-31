package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

// x/checkers module sentinel errors
var (
	ErrInvalidSigner    = sdkerrors.Register(ModuleName, 1100, "expected gov account as only signer for proposal message")
	ErrSample           = sdkerrors.Register(ModuleName, 1101, "sample error")
	ErrInvalidBlack     = sdkerrors.Register(ModuleName, 1102, "black address is invalid: %s")
	ErrInvalidRed       = sdkerrors.Register(ModuleName, 1103, "red address is invalid: %s")
	ErrGameNotParseable = sdkerrors.Register(ModuleName, 1104, "game cannot be parsed")
	ErrInvalidGameIndex     = sdkerrors.Register(ModuleName, 1105, "game index is invalid")
    ErrInvalidPositionIndex = sdkerrors.Register(ModuleName, 1106, "position index is invalid")
    ErrMoveAbsent           = sdkerrors.Register(ModuleName, 1107, "there is no move")
	ErrGameNotFound     = sdkerrors.Register(ModuleName, 1108, "game by id not found")
    ErrCreatorNotPlayer = sdkerrors.Register(ModuleName, 1109, "message creator is not a player")
    ErrNotPlayerTurn    = sdkerrors.Register(ModuleName, 1110, "player tried to play out of turn")
    ErrWrongMove        = sdkerrors.Register(ModuleName, 1111, "wrong move")
    
)
