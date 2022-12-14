package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authTypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	stakingTypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

type DistrKeeper interface {
	// Methods imported from distr should be defined here
}

type SlashingKeeper interface {
	// Methods imported from slashing should be defined here
}

type StakingKeeper interface {
	// Methods imported from staking should be defined here
	GetValidator(ctx sdk.Context, addr sdk.ValAddress) (validator stakingTypes.Validator, found bool)
}

// AccountKeeper defines the expected account keeper used for simulations (noalias)
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) authTypes.AccountI
	// Methods imported from account should be defined here
}

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	// Methods imported from bank should be defined here
}
