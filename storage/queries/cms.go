package queries

import (
	"sytron-server/storage/tables"
	"sytron-server/types/models"
)

func GetDestinations() (destinations models.Destination, err error) {
	err = db.From(tables.DESTINATIONS).
		Select("_id,name,one_liner,hero").
		Execute(&destinations)
	return
}
