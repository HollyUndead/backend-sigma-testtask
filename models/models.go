package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id          primitive.ObjectID `bson:"_id" json:"id"`
	FirstName   string             `json:"firstname" validate:"required,min=3,max=20"`
	LastName    string             `json:"lastname" validate:"required,min=3,max=35"`
	Age         int64              `json:"age,string" validate:"required,numeric,min=16"`
	PhoneNumber string             `json:"phonenumber" validate:"required,len=10"`
}
