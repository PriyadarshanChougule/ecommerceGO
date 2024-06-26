package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID              primitive.ObjectID `json:"_id"`
	First_name      *string            `json:"first_name"`
	Last_name       *string            `json:"last_name"`
	Password        *string            `json:"password"`
	Email           *string            `json:"email"`
	Phone           *string            `json:"phone"`
	Token           *string            `json:"token"`
	Refresh_Token   *string            `json:"refresh_token"`
	Created_At      time.Time          `json:"created_at"`
	Updated_At      time.Time          `json:"updated_at"`
	User_ID         string             `json:"user_id"`
	UserCart        []ProductUser      `json:"user_cart" bson:"usercart"`
	Address_Details []Address          `json:"address_details" bson:"address"`
	Order_Status    []Order            `json:"order_status" bson:"orders"`
}

type Product struct {
	Product_ID   primitive.ObjectID `json:"_id"`
	Product_Name *string            `json:"product_name"`
	Price        *uint64            `json:"price"`
	Rating       *uint8             `json:"rating"`
	Image        *string            `json:"image"`
}

type ProductUser struct {
	Product_ID   primitive.ObjectID `json:"_id" bson:"_id"`
	Product_Name *string            `json:"product_name" bson:"product_name"`
	Price        int                `json:"price" bson:"price"`
	Rating       *uint8             `json:"rating" bson:"rating"`
	Image        *string            `json:"image" bson:"image"`
}

type Address struct {
	Address_id primitive.ObjectID `bson:"_id"`
	House      *string            `json:"house_name" bson:"house_name"`
	Street     *string            `json:"street_name" bson:"street_name"`
	City       *string            `json:"city_name" bson:"city_name"`
	Pincode    *string            `json:"pin_code" bson:"pin_code"`
}

type Order struct {
	Order_ID       primitive.ObjectID `bson:"_id"`
	Order_Cart     []ProductUser      `json:"order_list" bson:"order_list"`
	Ordered_At     time.Time          `json:"ordered_at" bson:"ordered_at"`
	Price          int                `json:"price" bson:"price"`
	Discount       *int               `json:"discount" bson:"discount"`
	Payment_Method Payment            `json:"payment_method" bson:"payment_method"`
}

type Payment struct {
	Digital bool
	COD     bool
}
