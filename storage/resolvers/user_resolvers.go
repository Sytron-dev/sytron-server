package resolvers

import (
	"sytron-server/constants"
	"sytron-server/types/models"
)

var BackOfficersResolver = CollectionResolver[models.BackOfficer]{
	collectionName: constants.CREDENTIALS_COLLECTION_BACK_OFFICERS,
	model:          models.BackOfficer{},
}
