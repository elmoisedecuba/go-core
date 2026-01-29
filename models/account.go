package models

type Account struct {
	UUID       string `json:"uuid" bson:"uuid"`
	Email      string `json:"email" bson:"email"`
	Token      string `json:"token" bson:"token"`
	Password   string `json:"password" bson:"password"`
	FirstName  string `json:"firstName" bson:"firstName"`
	LastName   string `json:"lastName" bson:"lastName"`
	Username   string `json:"username" bson:"username"`
	Identifier string `json:"identifier" bson:"identifier"`
}
