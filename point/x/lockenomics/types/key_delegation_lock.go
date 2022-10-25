package types

import (
	"encoding/binary"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
)

var _ binary.ByteOrder

const (
	// DelegationLockKeyPrefix is the prefix to retrieve all DelegationLock
	DelegationLockKeyPrefix = "DelegationLock/value/"
)

/* TODO Create function to make a bytes index from string indes and a function to make a string index to bytes index.
 it is actually hex.DecodeString(indexString), hex.EncodeToString(indesBytes), so we don't have a logic duplication everywhere
*/
 */

// GetDelegationKey creates the key for delegator bond with validator
// VALUE: staking/Delegation
func GetDelegationLockKey(delAddr sdk.AccAddress, valAddr sdk.ValAddress) []byte {
	return append(GetDelegationsLockKey(delAddr), address.MustLengthPrefix(valAddr)...)
}

// GetDelegationsKey creates the prefix for a delegator for all validators
func GetDelegationsLockKey(delAddr sdk.AccAddress) []byte {
	return append(KeyPrefix(DelegationLockKeyPrefix), address.MustLengthPrefix(delAddr)...)
}
