package models

type User struct {
	*CollectionDocument

	First_name   string `json:"first_name"    validate:"required,min=2,max=100"`
	Last_name    string `json:"last_name"     validate:"required,min=2,max=100"`
	Password     string `                     validate:"required"               bson:"password"`
	Email        string `json:"email"         validate:"email,required"`
	Phone        string `json:"phone"         validate:"required"`
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}
