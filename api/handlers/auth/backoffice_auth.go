package auth

import (
	"sytron-server/constants"
	"sytron-server/storage/resolvers"
)

// creates backoffice auth credentials
var CreateBackOfficeAuth, LoginBackOfficer = GetGenericAuthCredentials(
	constants.USER_ROLE_BACKOFFICER,
	resolvers.BackOfficerAuthCredentialsResolver,
)
