package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/sonrhq/sonr/x/identity/keeper"
	"github.com/sonrhq/sonr/x/identity/types"
)

func SimulateMsgRegisterIdentity(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgRegisterIdentity{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the RegisterIdentity simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "RegisterIdentity simulation not implemented"), nil, nil
	}
}
