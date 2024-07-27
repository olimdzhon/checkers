package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

// x/checkers module sentinel errors
var (
	ErrInvalidSigner = sdkerrors.Register(ModuleName, 1100, "expected gov account as only signer for proposal message")
	ErrSample        = sdkerrors.Register(ModuleName, 1101, "sample error")
	ErrInvalidBlack     = sdkerrors.Register(ModuleName, 1100, "black address is invalid: %s")
    ErrInvalidRed       = sdkerrors.Register(ModuleName, 1101, "red address is invalid: %s")
    ErrGameNotParseable = sdkerrors.Register(ModuleName, 1102, "game cannot be parsed")
)
