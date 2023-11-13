package resolvers

import (
	"sytron-server/constants"
	"sytron-server/types/models"
)

type cmsResolver[T any] struct {
	CollectionResolver[T]
}

var LocationResolver = cmsResolver[models.Location]{
	CollectionResolver[models.Location]{
		collectionName: constants.CMS_COLLECTION_LOCATIONS,
	},
}

var DestinationResolver = cmsResolver[models.Destination]{
	CollectionResolver[models.Destination]{
		collectionName: constants.CMS_COLLECTION_DESTINATIONS,
	},
}
