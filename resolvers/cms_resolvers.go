package resolvers

import (
	"sytron-server/database"
	"sytron-server/models"
)

type cmsResolver[T any] struct {
	*collectionResolver[T]
}

var DestinationResolver = cmsResolver[models.Destination]{
	&collectionResolver[models.Destination]{
		collectionName: database.DESTINATIONS_COLLECTION,
	},
}
