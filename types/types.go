package types

type Coordinates struct {
	Latitude  float64 `bson:"lat" json:"lat"`
	Longitude float64 `bson:"lon" json:"lon"`
}
