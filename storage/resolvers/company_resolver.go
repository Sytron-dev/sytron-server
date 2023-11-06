package resolvers

import (
	"sytron-server/constants"
	"sytron-server/storage/models"
)

var CompanyResolver = CollectionResolver[models.Company]{
	model:          models.Company{},
	collectionName: constants.COMPANIES_COLLECTION,
}
