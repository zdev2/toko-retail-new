package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Base struct{
	CreatedAt time.Time  `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" bson:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" bson:"deleted_at,omitempty"`
}

type Admin struct {
	ID 		  primitive.ObjectID `json:"_id" bson:"_id"`
	Username  string 			 `json:"username"   bson:"username"`
	Password  string  			 `json:"password"   bson:"password"`
	LastLogin string             `json:"last_login" bson:"last_login"`
	Base
}

type Product struct {
	ID 		  	primitive.ObjectID 	`json:"_id"          bson:"_id"`
	ProductName string 				`json:"product_name" bson:"product_name"`
	Stock 		int 				`json:"stock"        bson:"stock"`
}