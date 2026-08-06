package models

// Each struct in this file represents one MongoDB collection or a document
// that is stored inside another collection.

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Struct tags (json/bson) tell Go how each field should appear while
// encoding (marshalling) and decoding (unmarshalling).

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"` //omit empty removes these feilds from the struct
	FirstName string             `json:"first_name"`
	LastName  string             `json:"last_name"`
	Password  string             `json:"password"`
	Email     string             `json:"email"`

	// Phone stores the user's contact number.
	Phone string `json:"phone"`

	//stores the currently assigned token to the user
	Token string `json:"token"`

	RefreshToken string `json:"refresh_token"`

	CreatedAt time.Time `json:"created_at"`

	// UpdatedAt stores the last time this document was modified.
	UpdatedAt time.Time `json:"updated_at"`

	UserID string `json:"user_id"`

	// UserCart stores every product currently present
	// in the user's shopping cart.
	//

	UserCart []ProductUser `json:"user_cart" bson:"user_cart"`

	AddressDetails []Address `json:"address" bson:"address"`

	// OrderStatus stores every order placed by this user.
	//
	// Since one user can place multiple orders,
	// this is also represented as a slice.
	OrderStatus []Order `json:"orders" bson:"orders"`
}

type Product struct {
	ProductID primitive.ObjectID `json:"_id" bson:"_id"`

	ProductName string `json:"product_name" bson:"product_name"`

	ProductPrice int `json:"product_price" bson:"product_price"`

	ProductRating int `json:"product_rating" bson:"product_rating"`

	Image string `json:"image" bson:"image"`
}

type ProductUser struct {
	ProductID primitive.ObjectID `json:"_id" bson:"_id"`

	ProductName string `json:"product_name" bson:"product_name"`

	ProductPrice int `json:"product_price" bson:"product_price"`

	ProductRating int `json:"product_rating" bson:"product_rating"`

	Image string `json:"image" bson:"image"`
}

type Address struct {

	// Unique identifier for this address.
	AddressID primitive.ObjectID `bson:"_id"`

	House string `json:"house_name"`

	Street string `json:"house_street"`

	City string `json:"house_city"`

	Pincode int `json:"house_pincode"`
}

// Order represents one completed purchase.
//
// Every order contains its own list of products,
// payment details, timestamps and pricing information.
type Order struct {

	// Unique identifier for the order.
	OrderID primitive.ObjectID `json:"_id" bson:"_id"`

	// OrderCart stores all products purchased
	// in this particular order.
	OrderCart []ProductUser `json:"order_cart" bson:"order_cart"`

	// OrderedAt stores the exact time
	// when the order was placed.
	OrderedAt time.Time `json:"ordered_at" bson:"ordered_at"`

	// Total order price.
	Price int `json:"price" bson:"price"`

	// Discount applied on the order.
	Discount int `json:"discount" bson:"discount"`

	// PaymentMethod stores how
	// the customer paid.
	PaymentMethod Payment `json:"payment_method" bson:"payment_method"`
}

type Payment struct {

	// true if paid online
	Digital bool `json:"digital"`

	// true if Cash On Delivery
	COD bool `json:"cod"`
}
