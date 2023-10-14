package auth_controller

import (
	"sytron-server/constants"
	"sytron-server/resolvers"
)

// creates backoffice auth credentials
var CreateBackOfficeAuth, LoginBackOfficer = GetGenericAuthCredentials(
	constants.USER_ROLE_BACKOFFICER,
	resolvers.BackOfficerAuthCredentialsResolver,
)
