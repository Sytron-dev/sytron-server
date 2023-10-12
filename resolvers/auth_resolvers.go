package resolvers

import (
	"sytron-server/database"
	"sytron-server/models"
)

var UserAuthCredentialsResolver = collectionResolver[models.AuthCredential]{
	collectionName: database.CREDENTIALS_COLLECTION_USERS,
	model:          models.AuthCredential{},
}

var BackOfficerAuthCredentialsResolver = collectionResolver[models.AuthCredential]{
	collectionName: database.CREDENTIALS_COLLECTION_USERS,
	model:          models.AuthCredential{},
}
var MerchantAuthCredentialsResolver = collectionResolver[models.AuthCredential]{
	collectionName: database.CREDENTIALS_COLLECTION_USERS,
	model:          models.AuthCredential{},
}
