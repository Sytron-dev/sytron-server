package resolvers

import (
	"sytron-server/constants"
	"sytron-server/types/models"
)

var EventsResolver = cmsResolver[models.Event]{
	CollectionResolver[models.Event]{
		collectionName: constants.EVENTS_COLLECTION,
	},
}
