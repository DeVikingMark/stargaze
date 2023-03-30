package types

import sdkErrors "github.com/cosmos/cosmos-sdk/types/errors"

var (
	DefaultCodespace  = ModuleName
	ErrInvalidMethods = sdkErrors.Register(DefaultCodespace, 0, "invalid method in code/contract authorization") // Code or Contract Authorizations have invalid methods configured
)
