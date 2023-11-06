package auth

import (
	"sytron-server/constants"
	"sytron-server/storage/resolvers"
)

// creates merchant credentials
var CreateMerchantAuth, LoginMerchant = GetGenericAuthCredentials(
	constants.USER_ROLE_MERCHANT,
	resolvers.MerchantAuthCredentialsResolver,
)
