package resolvers

import (
	"sytron-server/database"
	"sytron-server/models"
)

type cmsResolver[T any] struct {
	*collectionResolver[T]
}

var LocationResolver = cmsResolver[models.Location]{
	&collectionResolver[models.Location]{
		collectionName: database.CMS_COLLECTION_LOCATIONS,
	},
}

var DestinationResolver = cmsResolver[models.Destination]{
	&collectionResolver[models.Destination]{
		collectionName: database.CMS_COLLECTION_DESTINATIONS,
	},
}
