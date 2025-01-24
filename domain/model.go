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
	Price       float64 			`json:"price"        bson:"price"`
	Stock 		int 				`json:"stock"        bson:"stock"`
	CreatedBy	string				`json:"created_by"   bson:"created_by"`
	Base
}

type History struct {
	ID 		  	primitive.ObjectID 	`json:"_id"          bson:"_id"`
	ProductID	primitive.ObjectID  `json:"product_id"   bson:"product_id"`
	Amount		int					`json:"amount"       bson:"amount"`
	Status		string				`json:"status"       bson:"status"`
	Description string				`json:"description"  bson:"description"`
	Base
}

type Sale struct {
	ID 		  		primitive.ObjectID 	`json:"_id"           bson:"_id"`
	InvoiceCode		string				`json:"invoice_code"  bson:"invoice_code"`
	BuyerName		string				`json:"buyer_name"    bson:"buyer_name"`
	Subtotal		float64				`json:"subtotal"      bson:"subtotal"`
	DiscountCode	string				`json:"discount_code" bson:"discount_code"`
	Discount 		float64				`json:"discount"      bson:"discount"`
	Total			float64				`json:"total"         bson:"total"`
	CreatedBy		string				`json:"created_by"    bson:"created_by"`
	Base
	Items 			[]SaleItems			`json:"items"         bson:"items"`
}

type SaleItems struct {
	ProductID	primitive.ObjectID `json:"product_id" bson:"product_id"`
	SaleID		primitive.ObjectID `json:"sale_id"    bson:"sale_id"`
	Amount		int				   `json:"amount"     bson:"amount"`
	Subtotal	float64			   `json:"subtotal"   bson:"subtotal"`
	Base
}