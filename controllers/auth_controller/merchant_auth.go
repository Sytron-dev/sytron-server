package auth_controller

import (
	"sytron-server/constants"
	"sytron-server/resolvers"
)

// creates merchant credentials
var CreateMerchantAuth, LoginMerchant = GetGenericAuthCredentials(
	constants.USER_ROLE_MERCHANT,
	resolvers.MerchantAuthCredentialsResolver,
)
