package simulation

import (
	"math/rand"

	"fmt"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/simapp/helpers"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	proto "github.com/gogo/protobuf/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	regentypes "github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/x/group"
	"github.com/regen-network/regen-ledger/x/group/exported"
)

// Simulation operation weights constants
const (
	OpMsgCreateGroupRequest               = "op_weight_msg_create_group"
	OpMsgUpdateGroupAdminRequest          = "op_weight_msg_update_group_admin"
	OpMsgUpdateGroupMetadata              = "op_wieght_msg_update_group_metadata"
	OpMsgUpdateGroupMembers               = "op_weight_msg_update_group_members"
	OpMsgCreateGroupAccountRequest        = "op_weight_msg_create_group_account"
	OpMsgUpdateGroupAccountAdmin          = "op_weight_msg_update_group_account_admin"
	OpMsgUpdateGroupAccountDecisionPolicy = "op_weight_msg_update_group_account_decision_policy"
	OpMsgUpdateGroupAccountComment        = "op_weight_msg_update_group_account_comment"
	OpMsgCreateProposal                   = "op_weight_msg_create_proposal"
	OpMsgVote                             = "op_weight_msg_vote"
	OpMsgExec                             = "ops_weight_msg_exec"
)

const (
	WeightCreateGroup        = 100
	WeightUpdateGroup        = 50
	WeightCreateProposal     = 100
	WeightMsgVote            = 100
	WeightUpdateGroupAccount = 30
)

// WeightedOperations returns all the operations from the module with their respective weights
func WeightedOperations(
	appParams simtypes.AppParams, cdc codec.JSONMarshaler, ak exported.AccountKeeper,
	bk exported.BankKeeper, qryClient group.QueryClient) simulation.WeightedOperations {
	var (
		weightMsgCreateGroup                      int
		weightMsgUpdateGroupAdmin                 int
		weightMsgUpdateGroupMetadata              int
		weightMsgUpdateGroupMembers               int
		weightMsgCreateGroupAccount               int
		weightMsgUpdateGroupAccountAdmin          int
		weightMsgUpdateGroupAccountDecisionPolicy int
		weightMsgUpdateGroupAccountComment        int
		weightMsgCreateProposal                   int
		weightMsgVote                             int
		weightMsgExec                             int
	)

	appParams.GetOrGenerate(cdc, OpMsgCreateGroupRequest, &weightMsgCreateGroup, nil,
		func(_ *rand.Rand) {
			weightMsgCreateGroup = WeightCreateGroup
		},
	)
	appParams.GetOrGenerate(cdc, OpMsgCreateGroupAccountRequest, &weightMsgCreateGroupAccount, nil,
		func(_ *rand.Rand) {
			weightMsgCreateGroupAccount = WeightCreateGroup
		},
	)
	appParams.GetOrGenerate(cdc, OpMsgUpdateGroupAdminRequest, &weightMsgUpdateGroupAdmin, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateGroupAdmin = WeightUpdateGroup
		},
	)
	appParams.GetOrGenerate(cdc, OpMsgUpdateGroupMetadata, &weightMsgUpdateGroupMetadata, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateGroupMetadata = WeightUpdateGroupAccount
		},
	)
	appParams.GetOrGenerate(cdc, OpMsgCreateProposal, &weightMsgCreateProposal, nil,
		func(_ *rand.Rand) {
			weightMsgCreateProposal = WeightCreateProposal
		},
	)
	appParams.GetOrGenerate(cdc, OpMsgUpdateGroupMembers, &weightMsgUpdateGroupMembers, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateGroupMembers = WeightUpdateGroupAccount
		},
	)
	appParams.GetOrGenerate(cdc, OpMsgUpdateGroupAccountAdmin, &weightMsgUpdateGroupAccountAdmin, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateGroupAccountAdmin = WeightUpdateGroupAccount
		},
	)
	appParams.GetOrGenerate(cdc, OpMsgUpdateGroupAccountDecisionPolicy, &weightMsgUpdateGroupAccountDecisionPolicy, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateGroupAccountDecisionPolicy = WeightUpdateGroupAccount
		},
	)
	appParams.GetOrGenerate(cdc, OpMsgUpdateGroupAccountComment, &weightMsgUpdateGroupAccountComment, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateGroupAccountComment = WeightUpdateGroupAccount
		},
	)
	appParams.GetOrGenerate(cdc, OpMsgVote, &weightMsgVote, nil,
		func(_ *rand.Rand) {
			weightMsgVote = WeightMsgVote
		},
	)
	appParams.GetOrGenerate(cdc, OpMsgExec, &weightMsgExec, nil,
		func(_ *rand.Rand) {
			weightMsgExec = WeightMsgVote
		},
	)

	return simulation.WeightedOperations{
		simulation.NewWeightedOperation(
			weightMsgCreateGroup,
			SimulateMsgCreateGroup(ak, bk),
		),
		simulation.NewWeightedOperation(
			weightMsgUpdateGroupAdmin,
			SimulateMsgUpdateGroupAdmin(ak, bk, qryClient),
		),
		simulation.NewWeightedOperation(
			weightMsgCreateGroupAccount,
			SimulateMsgCreateGroupAccount(ak, bk, qryClient),
		),
		simulation.NewWeightedOperation(
			weightMsgUpdateGroupMetadata,
			SimulateMsgUpdateGroupMetadata(ak, bk, qryClient),
		),
		simulation.NewWeightedOperation(
			weightMsgCreateProposal,
			SimulateMsgCreateProposal(ak, bk, qryClient),
		),
		simulation.NewWeightedOperation(
			weightMsgUpdateGroupMembers,
			SimulateMsgUpdateGroupMembers(ak, bk, qryClient),
		),
		simulation.NewWeightedOperation(
			weightMsgUpdateGroupAccountAdmin,
			SimulateMsgUpdateGroupAccountAdmin(ak, bk, qryClient),
		),
		simulation.NewWeightedOperation(
			weightMsgUpdateGroupAccountDecisionPolicy,
			SimulateMsgUpdateGroupAccountDecisionPolicy(ak, bk, qryClient),
		),
		simulation.NewWeightedOperation(
			weightMsgUpdateGroupAccountComment,
			SimulateMsgUpdateGroupAccountComment(ak, bk, qryClient),
		),
	}
}

// SimulateMsgCreateGroup generates a MsgCreateGroupRequest with random values
func SimulateMsgCreateGroup(ak exported.AccountKeeper, bk exported.BankKeeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accounts []simtypes.Account, chainID string) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		acc := accounts[0]

		account := ak.GetAccount(ctx, acc.Address)

		spendableCoins := bk.SpendableCoins(ctx, account.GetAddress())
		fees, err := simtypes.RandomFees(r, ctx, spendableCoins)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, group.TypeMsgCreateGroup, "fee error"), nil, err
		}

		members := []group.Member{
			{
				Address:  acc.Address.String(),
				Weight:   fmt.Sprintf("%d", simappparams.DefaultWeightMsgCreateValidator),
				Metadata: []byte(simtypes.RandStringOfLength(r, 10)),
			},
		}

		msg := &group.MsgCreateGroupRequest{Admin: acc.Address.String(), Members: members, Metadata: []byte(simtypes.RandStringOfLength(r, 10))}

		txGen := simappparams.MakeTestEncodingConfig().TxConfig

		tx, err := helpers.GenTx(
			txGen,
			[]sdk.Msg{msg},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{account.GetAccountNumber()},
			[]uint64{account.GetSequence()},
			acc.PrivKey,
		)

		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, group.TypeMsgCreateGroup, "unable to generate mock tx"), nil, err
		}

		_, _, err = app.Deliver(txGen.TxEncoder(), tx)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, msg.Type(), "unable to deliver tx"), nil, err
		}
		return simtypes.NewOperationMsg(msg, true, ""), nil, err
	}
}

// SimulateMsgCreateGroupAccount generates a NewMsgCreateGroupAccountRequest with random values
func SimulateMsgCreateGroupAccount(ak exported.AccountKeeper, bk exported.BankKeeper, qryClient group.QueryClient) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accounts []simtypes.Account, chainID string) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		acc := accounts[0]

		account := ak.GetAccount(ctx, acc.Address)

		spendableCoins := bk.SpendableCoins(ctx, account.GetAddress())
		fees, err := simtypes.RandomFees(r, ctx, spendableCoins)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, group.TypeMsgCreateGroupAccount, "fee error"), nil, err
		}

		groupAdmin, groupID, op, err := getGroupDetails(ctx, qryClient, acc)
		if err != nil {
			return op, nil, err
		}

		if groupAdmin == "" {
			return op, nil, nil
		}

		addr, err := sdk.AccAddressFromBech32(groupAdmin)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, group.TypeMsgCreateGroupAccount, "fail to decode acc address"), nil, err
		}

		msg, err := group.NewMsgCreateGroupAccountRequest(
			addr,
			groupID,
			[]byte(simtypes.RandStringOfLength(r, 10)),
			&group.ThresholdDecisionPolicy{
				Threshold: "50",
				Timeout:   proto.Duration{Seconds: int64(30 * 24 * 60 * 60)},
			},
		)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, group.TypeMsgCreateGroupAccount, err.Error()), nil, err
		}
		txGen := simappparams.MakeTestEncodingConfig().TxConfig
		tx, err := helpers.GenTx(
			txGen,
			[]sdk.Msg{msg},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{account.GetAccountNumber()},
			[]uint64{account.GetSequence()},
			acc.PrivKey,
		)

		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, group.TypeMsgCreateGroupAccount, "unable to generate mock tx"), nil, err
		}

		_, _, err = app.Deliver(txGen.TxEncoder(), tx)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, msg.Type(), "unable to deliver tx"), nil, err
		}
		return simtypes.NewOperationMsg(msg, true, ""), nil, err
	}
}

func getGroupDetails(ctx sdk.Context, qryClient group.QueryClient, acc simtypes.Account) (string, uint64, simtypes.OperationMsg, error) {
	ctx1 := regentypes.Context{Context: ctx}

	result, err := qryClient.Groups(ctx1, &group.QueryGroupsRequest{})
	if err != nil {
		return "", 0, simtypes.NoOpMsg(group.ModuleName, group.TypeMsgCreateGroupAccount, "fail to query group info"), err
	}

	groupAdmin := ""
	groupID := uint64(0)
	for _, item := range result.Info {
		if item.Admin == acc.Address.String() {
			groupAdmin = acc.Address.String()
			groupID = item.GroupId
			break
		}
	}

	return groupAdmin, groupID, simtypes.NoOpMsg(group.ModuleName, group.TypeMsgCreateGroupAccount, ""), nil
}

// SimulateMsgCreateProposal generates a NewMsgCreateProposalRequest with random values
func SimulateMsgCreateProposal(ak exported.AccountKeeper, bk exported.BankKeeper, queryClient group.QueryClient) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accounts []simtypes.Account, chainID string) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		acc := accounts[0]

		account := ak.GetAccount(ctx, acc.Address)

		spendableCoins := bk.SpendableCoins(ctx, account.GetAddress())
		fees, err := simtypes.RandomFees(r, ctx, spendableCoins)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, group.TypeMsgCreateProposal, "fee error"), nil, err
		}

		ctx1 := regentypes.Context{Context: ctx}
		groupAdmin, _, op, err := getGroupDetails(ctx, queryClient, acc)
		if err != nil {
			return op, nil, err
		}
		if groupAdmin == "" {
			return op, nil, nil
		}

		result, err := queryClient.GroupAccountsByAdmin(ctx1, &group.QueryGroupAccountsByAdminRequest{Admin: groupAdmin})
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, group.TypeMsgCreateProposal, "fail to query group info"), nil, err
		}
		groupAccounts := result.GetGroupAccounts()

		if len(groupAccounts) == 0 {
			return simtypes.NoOpMsg(group.ModuleName, group.TypeMsgCreateProposal, "no group account found"), nil, nil
		}

		msg := group.MsgCreateProposalRequest{
			GroupAccount: groupAccounts[0].GroupAccount,
			Proposers:    []string{acc.Address.String()},
			Metadata:     []byte(simtypes.RandStringOfLength(r, 10)),
		}
		txGen := simappparams.MakeTestEncodingConfig().TxConfig

		tx, err := helpers.GenTx(
			txGen,
			[]sdk.Msg{&msg},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{account.GetAccountNumber()},
			[]uint64{account.GetSequence()},
			acc.PrivKey,
		)

		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, group.TypeMsgCreateProposal, "unable to generate mock tx"), nil, err
		}

		_, _, err = app.Deliver(txGen.TxEncoder(), tx)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, msg.Type(), "unable to deliver tx"), nil, err
		}
		return simtypes.NewOperationMsg(&msg, true, ""), nil, err
	}
}

// SimulateMsgUpdateGroupAdmin generates a MsgUpdateGroupAccountAdminRequest with random values
func SimulateMsgUpdateGroupAdmin(ak exported.AccountKeeper, bk exported.BankKeeper, queryClient group.QueryClient) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accounts []simtypes.Account, chainID string) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		acc1 := accounts[0]
		acc2 := accounts[1]

		account := ak.GetAccount(ctx, acc1.Address)

		spendableCoins := bk.SpendableCoins(ctx, account.GetAddress())
		fees, err := simtypes.RandomFees(r, ctx, spendableCoins)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, group.TypeMsgUpdateGroupAdmin, "fee error"), nil, err
		}

		ctx1 := regentypes.Context{Context: ctx}

		groupAdmin, groupID, op, err := getGroupDetails(ctx, queryClient, acc1)
		if err != nil {
			return op, nil, err
		}
		if groupAdmin == "" {
			return op, nil, nil
		}
		result1, err := queryClient.GroupAccountsByAdmin(ctx1, &group.QueryGroupAccountsByAdminRequest{Admin: groupAdmin})
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, group.TypeMsgUpdateGroupAdmin, "update-group-admin: fail to query group info"), nil, err
		}

		groupAccounts := result1.GetGroupAccounts()

		if len(groupAccounts) == 0 {
			return simtypes.NoOpMsg(group.ModuleName, group.TypeMsgUpdateGroupAdmin, "no group account found"), nil, nil
		}

		msg := group.MsgUpdateGroupAdminRequest{
			GroupId:  groupID,
			Admin:    groupAccounts[0].Admin,
			NewAdmin: acc2.Address.String(),
		}
		txGen := simappparams.MakeTestEncodingConfig().TxConfig

		tx, err := helpers.GenTx(
			txGen,
			[]sdk.Msg{&msg},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{account.GetAccountNumber()},
			[]uint64{account.GetSequence()},
			acc1.PrivKey,
		)

		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, group.TypeMsgUpdateGroupAdmin, "unable to generate mock tx"), nil, err
		}

		_, _, err = app.Deliver(txGen.TxEncoder(), tx)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, msg.Type(), "unable to deliver tx"), nil, err
		}
		return simtypes.NewOperationMsg(&msg, true, ""), nil, err
	}
}

// SimulateMsgUpdateGroupMetadata generates a MsgUpdateGroupMetadataRequest with random values
func SimulateMsgUpdateGroupMetadata(ak exported.AccountKeeper, bk exported.BankKeeper, queryClient group.QueryClient) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accounts []simtypes.Account, chainID string) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		acc := accounts[0]

		account := ak.GetAccount(ctx, acc.Address)

		spendableCoins := bk.SpendableCoins(ctx, account.GetAddress())
		fees, err := simtypes.RandomFees(r, ctx, spendableCoins)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, group.TypeMsgUpdateGroupComment, "fee error"), nil, err
		}

		groupAdmin, groupID, op, err := getGroupDetails(ctx, queryClient, acc)
		if err != nil {
			return op, nil, err
		}

		if groupAdmin == "" {
			return op, nil, nil
		}

		msg := group.MsgUpdateGroupMetadataRequest{
			GroupId:  groupID,
			Admin:    groupAdmin,
			Metadata: []byte(simtypes.RandStringOfLength(r, 10)),
		}
		txGen := simappparams.MakeTestEncodingConfig().TxConfig

		tx, err := helpers.GenTx(
			txGen,
			[]sdk.Msg{&msg},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{account.GetAccountNumber()},
			[]uint64{account.GetSequence()},
			acc.PrivKey,
		)

		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, group.TypeMsgUpdateGroupComment, "unable to generate mock tx"), nil, err
		}

		_, _, err = app.Deliver(txGen.TxEncoder(), tx)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, msg.Type(), "unable to deliver tx"), nil, err
		}
		return simtypes.NewOperationMsg(&msg, true, ""), nil, err
	}
}

// SimulateMsgUpdateGroupMembers generates a MsgUpdateGroupMembersRequest with random values
func SimulateMsgUpdateGroupMembers(ak exported.AccountKeeper,
	bk exported.BankKeeper, queryClient group.QueryClient) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accounts []simtypes.Account, chainID string) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		acc1 := accounts[0]
		acc2 := accounts[1]
		acc3 := accounts[2]

		account := ak.GetAccount(ctx, acc1.Address)

		spendableCoins := bk.SpendableCoins(ctx, account.GetAddress())
		fees, err := simtypes.RandomFees(r, ctx, spendableCoins)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, group.TypeMsgUpdateGroupMembers, "fee error"), nil, err
		}

		ctx1 := regentypes.Context{Context: ctx}
		groupAdmin, groupID, op, err := getGroupDetails(ctx, queryClient, acc1)
		if err != nil {
			return op, nil, err
		}

		if groupAdmin == "" {
			return op, nil, nil
		}

		result, err := queryClient.GroupAccountsByAdmin(ctx1, &group.QueryGroupAccountsByAdminRequest{Admin: groupAdmin})
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, group.TypeMsgUpdateGroupMembers, "fail to query group info"), nil, err
		}
		groupAccounts := result.GetGroupAccounts()

		if len(groupAccounts) == 0 {
			return simtypes.NoOpMsg(group.ModuleName, group.TypeMsgUpdateGroupMembers, "no group account found"), nil, nil
		}

		members := []group.Member{
			{
				Address:  acc2.Address.String(),
				Weight:   fmt.Sprintf("%d", simappparams.DefaultWeightMsgCreateValidator),
				Metadata: []byte(simtypes.RandStringOfLength(r, 10)),
			},
			{
				Address:  acc3.Address.String(),
				Weight:   fmt.Sprintf("%d", simappparams.DefaultWeightCommunitySpendProposal),
				Metadata: []byte(simtypes.RandStringOfLength(r, 10)),
			},
		}

		msg := group.MsgUpdateGroupMembersRequest{
			GroupId:       groupID,
			Admin:         groupAdmin,
			MemberUpdates: members,
		}
		txGen := simappparams.MakeTestEncodingConfig().TxConfig

		tx, err := helpers.GenTx(
			txGen,
			[]sdk.Msg{&msg},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{account.GetAccountNumber()},
			[]uint64{account.GetSequence()},
			acc1.PrivKey,
		)

		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, group.TypeMsgUpdateGroupMembers, "unable to generate mock tx"), nil, err
		}

		_, _, err = app.Deliver(txGen.TxEncoder(), tx)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, msg.Type(), "unable to deliver tx"), nil, err
		}
		return simtypes.NewOperationMsg(&msg, true, ""), nil, err
	}
}

// SimulateMsgUpdateGroupAccountAdmin generates a MsgUpdateGroupAccountAdminRequest with random values
func SimulateMsgUpdateGroupAccountAdmin(ak exported.AccountKeeper, bk exported.BankKeeper, queryClient group.QueryClient) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accounts []simtypes.Account, chainID string) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		acc1 := accounts[0]
		acc2 := accounts[1]

		account := ak.GetAccount(ctx, acc1.Address)

		spendableCoins := bk.SpendableCoins(ctx, account.GetAddress())
		fees, err := simtypes.RandomFees(r, ctx, spendableCoins)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, group.TypeMsgUpdateGroupAccountAdmin, "fee error"), nil, err
		}

		ctx1 := regentypes.Context{Context: ctx}

		groupAdmin, _, op, err := getGroupDetails(ctx, queryClient, acc1)
		if err != nil {
			return op, nil, err
		}

		if groupAdmin == "" {
			return op, nil, nil
		}

		result, err := queryClient.GroupAccountsByAdmin(ctx1, &group.QueryGroupAccountsByAdminRequest{Admin: groupAdmin})
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, group.TypeMsgUpdateGroupAccountAdmin, "fail to query group info"), nil, err
		}

		groupAccounts := result.GetGroupAccounts()

		if len(groupAccounts) == 0 {
			return simtypes.NoOpMsg(group.ModuleName, group.TypeMsgUpdateGroupAccountAdmin, "no group account found"), nil, nil
		}

		msg := group.MsgUpdateGroupAccountAdminRequest{
			Admin:        groupAccounts[0].Admin,
			GroupAccount: groupAccounts[0].GroupAccount,
			NewAdmin:     acc2.Address.String(),
		}
		txGen := simappparams.MakeTestEncodingConfig().TxConfig

		tx, err := helpers.GenTx(
			txGen,
			[]sdk.Msg{&msg},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{account.GetAccountNumber()},
			[]uint64{account.GetSequence()},
			acc1.PrivKey,
		)

		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, group.TypeMsgUpdateGroupAccountAdmin, "unable to generate mock tx"), nil, err
		}

		_, _, err = app.Deliver(txGen.TxEncoder(), tx)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, msg.Type(), "unable to deliver tx"), nil, err
		}
		return simtypes.NewOperationMsg(&msg, true, ""), nil, err
	}
}

// SimulateMsgUpdateGroupAccountDecisionPolicy generates a NewMsgUpdateGroupAccountDecisionPolicyRequest with random values
func SimulateMsgUpdateGroupAccountDecisionPolicy(ak exported.AccountKeeper,
	bk exported.BankKeeper, queryClient group.QueryClient) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accounts []simtypes.Account, chainID string) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		acc1 := accounts[0]

		account := ak.GetAccount(ctx, acc1.Address)

		spendableCoins := bk.SpendableCoins(ctx, account.GetAddress())
		fees, err := simtypes.RandomFees(r, ctx, spendableCoins)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, group.TypeMsgUpdateGroupAccountDecisionPolicy, "fee error"), nil, err
		}

		ctx1 := regentypes.Context{Context: ctx}
		groupAdmin, _, op, err := getGroupDetails(ctx, queryClient, acc1)
		if err != nil {
			return op, nil, err
		}

		if groupAdmin == "" {
			return op, nil, nil
		}

		result, err := queryClient.GroupAccountsByAdmin(ctx1, &group.QueryGroupAccountsByAdminRequest{Admin: groupAdmin})
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, group.TypeMsgUpdateGroupAccountDecisionPolicy, "fail to query group info"), nil, err
		}
		groupAccounts := result.GetGroupAccounts()

		if len(groupAccounts) == 0 {
			return simtypes.NoOpMsg(group.ModuleName, group.TypeMsgUpdateGroupAccountDecisionPolicy, "no group account found"), nil, nil
		}

		adminBech32, err := sdk.AccAddressFromBech32(groupAccounts[0].Admin)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, group.TypeMsgUpdateGroupAccountDecisionPolicy, fmt.Sprintf("fail to decide bech32 address: %s", err.Error())), nil, nil
		}

		groupAccountBech32, err := sdk.AccAddressFromBech32(groupAccounts[0].GroupAccount)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, group.TypeMsgUpdateGroupAccountDecisionPolicy, fmt.Sprintf("fail to decide bech32 address: %s", err.Error())), nil, nil
		}

		msg, err := group.NewMsgUpdateGroupAccountDecisionPolicyRequest(adminBech32, groupAccountBech32, &group.ThresholdDecisionPolicy{
			Threshold: fmt.Sprintf("%d", simtypes.RandIntBetween(r, 1, 100)),
			Timeout:   proto.Duration{Seconds: int64(simtypes.RandIntBetween(r, 100, 1000))},
		})

		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, group.TypeMsgUpdateGroupAccountDecisionPolicy, err.Error()), nil, err
		}

		txGen := simappparams.MakeTestEncodingConfig().TxConfig

		tx, err := helpers.GenTx(
			txGen,
			[]sdk.Msg{msg},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{account.GetAccountNumber()},
			[]uint64{account.GetSequence()},
			acc1.PrivKey,
		)

		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, group.TypeMsgUpdateGroupAccountDecisionPolicy, "unable to generate mock tx"), nil, err
		}

		_, _, err = app.Deliver(txGen.TxEncoder(), tx)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, msg.Type(), "unable to deliver tx"), nil, err
		}
		return simtypes.NewOperationMsg(msg, true, ""), nil, err
	}
}

// SimulateMsgUpdateGroupAccountComment generates a MsgUpdateGroupAccountMetadataRequest with random values
func SimulateMsgUpdateGroupAccountComment(ak exported.AccountKeeper,
	bk exported.BankKeeper, queryClient group.QueryClient) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accounts []simtypes.Account, chainID string) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		acc1 := accounts[0]

		account := ak.GetAccount(ctx, acc1.Address)

		spendableCoins := bk.SpendableCoins(ctx, account.GetAddress())
		fees, err := simtypes.RandomFees(r, ctx, spendableCoins)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, group.TypeMsgUpdateGroupAccountComment, "fee error"), nil, err
		}

		ctx1 := regentypes.Context{Context: ctx}
		groupAdmin, _, op, err := getGroupDetails(ctx, queryClient, acc1)
		if err != nil {
			return op, nil, err
		}

		if groupAdmin == "" {
			return op, nil, nil
		}

		result, err := queryClient.GroupAccountsByAdmin(ctx1, &group.QueryGroupAccountsByAdminRequest{Admin: groupAdmin})
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, group.TypeMsgUpdateGroupAccountComment, "fail to query group info"), nil, err
		}
		groupAccounts := result.GetGroupAccounts()

		if len(groupAccounts) == 0 {
			return simtypes.NoOpMsg(group.ModuleName, group.TypeMsgUpdateGroupAccountComment, "no group account found"), nil, nil
		}

		msg := group.MsgUpdateGroupAccountMetadataRequest{
			Admin:        groupAccounts[0].Admin,
			GroupAccount: groupAccounts[0].GroupAccount,
			Metadata:     []byte(simtypes.RandStringOfLength(r, 10)),
		}

		txGen := simappparams.MakeTestEncodingConfig().TxConfig
		tx, err := helpers.GenTx(
			txGen,
			[]sdk.Msg{&msg},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{account.GetAccountNumber()},
			[]uint64{account.GetSequence()},
			acc1.PrivKey,
		)

		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, group.TypeMsgUpdateGroupAccountComment, "unable to generate mock tx"), nil, err
		}

		_, _, err = app.Deliver(txGen.TxEncoder(), tx)
		if err != nil {
			return simtypes.NoOpMsg(group.ModuleName, msg.Type(), "unable to deliver tx"), nil, err
		}
		return simtypes.NewOperationMsg(&msg, true, ""), nil, err
	}
}
