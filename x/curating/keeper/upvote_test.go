package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stakebird/testdata"
	"github.com/stretchr/testify/require"
)

func TestDelegation(t *testing.T) {
	// _, app, ctx := testdata.CreateTestInput()

	// // create fake addresses
	// delAddrs := testdata.AddTestAddrsIncremental(app, ctx, 3, sdk.NewInt(100000))
	// valAddrs := testdata.ConvertAddrsToValAddrs(delAddrs)

	// // create validator with 50% commission
	// commission := staking.NewCommissionRates(sdk.NewDecWithPrec(5, 1), sdk.NewDecWithPrec(5, 1), sdk.NewDec(0))
	// msg := staking.NewMsgCreateValidator(
	// 	valAddrs[0], testdata.ValConsPk1, sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(100)), staking.Description{}, commission, sdk.OneInt(),
	// )
	// sh := staking.NewHandler(app.StakingKeeper)
	// res, err := sh(ctx, msg)
	// require.NoError(t, err)
	// require.NotNil(t, res)
	// // end block to bond validator
	// staking.EndBlocker(ctx, app.StakingKeeper)
	// // next block
	// ctx = ctx.WithBlockHeight(ctx.BlockHeight() + 1)
	// // historical count should be 2 (once for validator init, once for delegation init)
	// require.Equal(t, uint64(2), app.DistrKeeper.GetValidatorHistoricalReferenceCount(ctx))

	// // perform end-user delegation
	// vendorID := uint64(1)
	// postID := uint64(1)
	// votingPeriod := time.Hour * 24 * 7
	// amount := sdk.NewInt64Coin("ufuel", 10000)
	// err = app.CuratingKeeper.Delegate(ctx, vendorID, postID, delAddrs[0], valAddrs[0], amount)
	// require.NoError(t, err)

	// // check if delegation is stored in staking store
	// delegations := app.StakingKeeper.GetAllDelegations(ctx)
	// require.Len(t, delegations, 1)

	// // check if delegation is in voting delegation queue
	// // endTime1 := ctx.BlockTime().Add(votingPeriod * 5 * time.Hour) // after block time
	// endTime := ctx.BlockTime().Add(votingPeriod * -5 * time.Hour) // before block time
	// spew.Dump(endTime)
	// app.CuratingKeeper.IterateVotingDelegationQueue(ctx, endTime, func(endTime time.Time, vendorID, postID, stakeID uint64, delegation stakingtypes.Delegation) bool {
	// 	require.True(t, delegation.GetShares().Equal(amount.Amount.ToDec()))
	// 	return true
	// })

	// // test end blocker, should remove delegation from voting delegation queue
	// ctx = ctx.WithBlockTime(endTime)
	// curating.EndBlocker(ctx, app.CuratingKeeper)

	// app.CuratingKeeper.IterateVotingDelegationQueue(ctx, endTime, func(endTime time.Time, vendorID, postID, stakeID uint64, delegation stakingtypes.Delegation) bool {
	// 	require.Fail(t, "queue should be empty")
	// 	return true
	// })
}

func TestCreateUpvote(t *testing.T) {
	_, app, ctx := testdata.CreateTestInput()

	postID := "500"
	vendorID := uint32(100)
	deposit := sdk.NewInt64Coin("ufuel", 100000)
	addrs := testdata.AddTestAddrsIncremental(app, ctx, 3, sdk.NewInt(100000))

	_, err := app.CuratingKeeper.CreateUpvote(ctx, vendorID, postID, addrs[0], addrs[0], 5, deposit)
	require.NoError(t, err)

	_, found := app.CuratingKeeper.GetPost(ctx, vendorID, postID)
	require.True(t, found, "post should be found")

	upvote, found := app.CuratingKeeper.GetUpvote(ctx, vendorID, postID, addrs[0])
	require.True(t, found, "upvote should be found")

	require.Equal(t, "25000000ufuel", upvote.VoteAmount.String())
}
