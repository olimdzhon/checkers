package types

import (
	"strconv"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/olimdzhon/checkers/x/checkers/rules"
)

var _ sdk.Msg = &MsgPlayMove{}

func NewMsgPlayMove(creator string, gameIndex string, fromX uint64, fromY uint64, toX uint64, toY uint64) *MsgPlayMove {
	return &MsgPlayMove{
		Creator:   creator,
		GameIndex: gameIndex,
		FromX:     fromX,
		FromY:     fromY,
		ToX:       toX,
		ToY:       toY,
	}
}

func (msg *MsgPlayMove) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	gameIndex, err := strconv.ParseInt(msg.GameIndex, 10, 64)
	if err != nil {
		return errorsmod.Wrapf(ErrInvalidGameIndex, "not parseable (%s)", err)
	}
	if uint64(gameIndex) < DefaultIndex {
		return errorsmod.Wrapf(ErrInvalidGameIndex, "number too low (%d)", gameIndex)
	}
	boardChecks := []struct {
		value uint64
		err   string
	}{
		{
			value: msg.FromX,
			err:   "fromX out of range (%d)",
		},
		{
			value: msg.ToX,
			err:   "toX out of range (%d)",
		},
		{
			value: msg.FromY,
			err:   "fromY out of range (%d)",
		},
		{
			value: msg.ToY,
			err:   "toY out of range (%d)",
		},
	}
	for _, situation := range boardChecks {
		if situation.value < 0 || rules.BOARD_DIM <= situation.value {
			return errorsmod.Wrapf(ErrInvalidPositionIndex, situation.err, situation.value)
		}
	}
	if msg.FromX == msg.ToX && msg.FromY == msg.ToY {
		return errorsmod.Wrapf(ErrMoveAbsent, "x (%d) and y (%d)", msg.FromX, msg.FromY)
	}
	return nil
}
