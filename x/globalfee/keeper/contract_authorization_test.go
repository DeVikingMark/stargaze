package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stargaze/v9/testutil/keeper"
	"github.com/public-awesome/stargaze/v9/testutil/sample"
	"github.com/public-awesome/stargaze/v9/x/globalfee/types"
	"github.com/stretchr/testify/require"
)

func Test_ContractAuthorization(t *testing.T) {
	k, ctx := keeper.GlobalFeeKeeper(t)
	ca := types.ContractAuthorization{
		ContractAddress: "cosmos1qyqszqgpqyqszqgpqyqszqgpqyqszqgpjnp7du",
		Methods:         []string{"mint", "list"},
	}

	t.Run("store invalid ca", func(t *testing.T) {
		err := k.SetContractAuthorization(ctx, types.ContractAuthorization{
			ContractAddress: "👻",
			Methods:         []string{"mint", "list"},
		})
		require.Error(t, err)
	})

	t.Run("non existing contract address", func(t *testing.T) {
		err := k.SetContractAuthorization(ctx, types.ContractAuthorization{
			ContractAddress: sample.AccAddress().String(),
			Methods:         []string{"mint", "list"},
		})
		require.Error(t, err)
	})

	t.Run("authorization doesnt exist", func(t *testing.T) {
		_, found := k.GetContractAuthorization(ctx, sdk.MustAccAddressFromBech32(ca.ContractAddress))
		require.False(t, found)
	})

	t.Run("store authorization", func(t *testing.T) {
		err := k.SetContractAuthorization(ctx, ca)
		require.NoError(t, err)

		_, found := k.GetContractAuthorization(ctx, sdk.MustAccAddressFromBech32(ca.ContractAddress))
		require.True(t, found)
	})

	t.Run("delete authorization", func(t *testing.T) {
		_, found := k.GetContractAuthorization(ctx, sdk.MustAccAddressFromBech32(ca.ContractAddress))
		require.True(t, found)

		k.DeleteContractAuthorization(ctx, sdk.MustAccAddressFromBech32(ca.ContractAddress))

		_, found = k.GetContractAuthorization(ctx, sdk.MustAccAddressFromBech32(ca.ContractAddress))
		require.False(t, found)
	})
}
