package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"cosmossdk.io/math"
	sdktypes "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	ibcclienttypes "github.com/cosmos/ibc-go/v8/modules/core/02-client/types" //nolint:staticcheck

	feetypes "github.com/neutron-org/neutron/v3/x/feerefunder/types"
	"github.com/neutron-org/neutron/v3/x/transfer/types"
)

const TestAddress = "cosmos10h9stc5v6ntgeygf5xf945njqq5h32r53uquvw"

func TestMsgTransferValidate(t *testing.T) {
	tests := []struct {
		name        string
		msg         types.MsgTransfer
		expectedErr error
	}{
		// We can check only fee validity because we didn't change original ValidateBasic call
		{
			"valid",
			types.MsgTransfer{
				SourcePort:    "port_id",
				SourceChannel: "channel_id",
				Token:         sdktypes.NewCoin("denom", math.NewInt(100)),
				Sender:        TestAddress,
				Receiver:      TestAddress,
				TimeoutHeight: ibcclienttypes.Height{
					RevisionNumber: 100,
					RevisionHeight: 100,
				},
				TimeoutTimestamp: 10000,
				Fee: feetypes.Fee{
					RecvFee:    sdktypes.NewCoins(),
					AckFee:     sdktypes.NewCoins(sdktypes.NewCoin("denom", math.NewInt(100))),
					TimeoutFee: sdktypes.NewCoins(sdktypes.NewCoin("denom", math.NewInt(100))),
				},
			},
			nil,
		},
		{
			"non-zero recv fee",
			types.MsgTransfer{
				SourcePort:    "port_id",
				SourceChannel: "channel_id",
				Token:         sdktypes.NewCoin("denom", math.NewInt(100)),
				Sender:        TestAddress,
				Receiver:      TestAddress,
				TimeoutHeight: ibcclienttypes.Height{
					RevisionNumber: 100,
					RevisionHeight: 100,
				},
				TimeoutTimestamp: 10000,
				Fee: feetypes.Fee{
					RecvFee:    sdktypes.NewCoins(sdktypes.NewCoin("denom", math.NewInt(100))),
					AckFee:     sdktypes.NewCoins(sdktypes.NewCoin("denom", math.NewInt(100))),
					TimeoutFee: sdktypes.NewCoins(sdktypes.NewCoin("denom", math.NewInt(100))),
				},
			},
			sdkerrors.ErrInvalidCoins,
		},
		{
			"zero ack fee",
			types.MsgTransfer{
				SourcePort:    "port_id",
				SourceChannel: "channel_id",
				Token:         sdktypes.NewCoin("denom", math.NewInt(100)),
				Sender:        TestAddress,
				Receiver:      TestAddress,
				TimeoutHeight: ibcclienttypes.Height{
					RevisionNumber: 100,
					RevisionHeight: 100,
				},
				TimeoutTimestamp: 10000,
				Fee: feetypes.Fee{
					RecvFee:    sdktypes.NewCoins(),
					AckFee:     sdktypes.NewCoins(),
					TimeoutFee: sdktypes.NewCoins(sdktypes.NewCoin("denom", math.NewInt(100))),
				},
			},
			sdkerrors.ErrInvalidCoins,
		},
		{
			"zero timeout fee",
			types.MsgTransfer{
				SourcePort:    "port_id",
				SourceChannel: "channel_id",
				Token:         sdktypes.NewCoin("denom", math.NewInt(100)),
				Sender:        TestAddress,
				Receiver:      TestAddress,
				TimeoutHeight: ibcclienttypes.Height{
					RevisionNumber: 100,
					RevisionHeight: 100,
				},
				TimeoutTimestamp: 10000,
				Fee: feetypes.Fee{
					RecvFee:    sdktypes.NewCoins(),
					AckFee:     sdktypes.NewCoins(sdktypes.NewCoin("denom", math.NewInt(100))),
					TimeoutFee: sdktypes.NewCoins(),
				},
			},
			sdkerrors.ErrInvalidCoins,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic(true)
			if tt.expectedErr != nil {
				require.ErrorIs(t, err, tt.expectedErr)
				return
			}
			require.NoError(t, err)
		})
	}
}
