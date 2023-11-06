package resolvers

import (
	"sytron-server/constants"
	"sytron-server/storage/models"
)

var EventsResolver = cmsResolver[models.Event]{
	CollectionResolver[models.Event]{
		collectionName: constants.EVENTS_COLLECTION,
	},
}
