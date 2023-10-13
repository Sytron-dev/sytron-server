package auth_controller

import (
	"sytron-server/constants"
)

// creates backofficer credentials
var CreateBackOfficer, LoginBackOfficer = GetGenericAuthCredentials(constants.USER_ROLE_BACKOFFICER)
