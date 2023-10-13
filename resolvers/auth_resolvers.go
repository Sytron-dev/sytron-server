package resolvers

import (
	"sytron-server/constants"
	"sytron-server/models"
)

var UserAuthCredentialsResolver = CollectionResolver[models.AuthCredential]{
	collectionName: constants.CREDENTIALS_COLLECTION_USERS,
	model:          models.AuthCredential{},
}

var BackOfficerAuthCredentialsResolver = CollectionResolver[models.AuthCredential]{
	collectionName: constants.CREDENTIALS_COLLECTION_BACK_OFFICERS,
	model:          models.AuthCredential{},
}

var MerchantAuthCredentialsResolver = CollectionResolver[models.AuthCredential]{
	collectionName: constants.CREDENTIALS_COLLECTION_MERCHANTS,
	model:          models.AuthCredential{},
}
