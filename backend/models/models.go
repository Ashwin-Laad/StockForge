package models

// Each struct in this file represents one MongoDB collection or a document
// that is stored inside another collection.
//
// These structs act as "blueprints" for the data. Whenever data is received
// from MongoDB or sent back to it, the MongoDB driver converts between BSON
// documents and these Go structs.
//
// Similarly, when data is sent or received through an HTTP API, the encoding/json
// package converts between JSON and these structs using the json tags.

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Struct tags (json/bson) tell Go how each field should appear while
// encoding (marshalling) and decoding (unmarshalling).
//
// Example:
//
// Go Field:
//     FirstName
//
// JSON:
//     {
//         "first_name": "Ashwin"
//     }
//
// BSON (MongoDB):
//     {
//         "first_name": "Ashwin"
//     }
//
// Without these tags, Go would use the field names themselves,
// which usually don't match API or database naming conventions.

type User struct {

	// ID is MongoDB's unique identifier (_id).
	//
	// primitive.ObjectID is MongoDB's special type for uniquely identifying
	// every document inside a collection.
	//
	// Every user stored in MongoDB gets one automatically.
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`

	// FirstName stores the user's first name.
	//
	// Type: string
	// Example: "Ashwin"
	FirstName string `json:"first_name"`

	// LastName stores the surname of the user.
	LastName string `json:"last_name"`

	// Password stores the user's password.
	//
	// IMPORTANT:
	// In production this should NEVER store the plain password.
	// Instead it should contain a hashed password (bcrypt, Argon2, etc.).
	Password string `json:"password"`

	// Email stores the user's email address.
	Email string `json:"email"`

	// Phone stores the user's contact number.
	Phone string `json:"phone"`

	// Token stores the JWT access token currently assigned
	// to this user after authentication.
	Token string `json:"token"`

	// RefreshToken stores the refresh token used
	// to generate new access tokens.
	RefreshToken string `json:"refresh_token"`

	// CreatedAt stores when the account was first created.
	//
	// time.Time is Go's built-in type for representing
	// dates and timestamps.
	CreatedAt time.Time `json:"created_at"`

	// UpdatedAt stores the last time this document was modified.
	UpdatedAt time.Time `json:"updated_at"`

	// UserID is another unique identifier stored as a string.
	//
	// Unlike MongoDB's ObjectID, this may be generated manually
	// and used throughout the application.
	UserID string `json:"user_id"`

	// UserCart stores every product currently present
	// in the user's shopping cart.
	//
	// []ProductUser means:
	// - [] represents a slice (dynamic array).
	// - ProductUser is the type stored inside that slice.
	//
	// Therefore this field can hold zero, one, or many products.
	UserCart []ProductUser `json:"user_cart" bson:"user_cart"`

	// AddressDetails stores every address belonging to the user.
	//
	// A user may save multiple addresses like:
	// Home
	// Office
	// Hostel
	AddressDetails []Address `json:"address" bson:"address"`

	// OrderStatus stores every order placed by this user.
	//
	// Since one user can place multiple orders,
	// this is also represented as a slice.
	OrderStatus []Order `json:"orders" bson:"orders"`
}

// Product represents a product available in the store.
//
// These products are typically stored in the Products collection
// and are shown to users for purchase.
type Product struct {

	// Unique MongoDB identifier for this product.
	ProductID primitive.ObjectID `json:"_id" bson:"_id"`

	// Product name.
	ProductName string `json:"product_name" bson:"product_name"`

	// Price of the product.
	ProductPrice int `json:"product_price" bson:"product_price"`

	// Product rating.
	//
	// Example:
	// 5 means five-star rating.
	ProductRating int `json:"product_rating" bson:"product_rating"`

	// URL pointing to the product image.
	Image string `json:"image" bson:"image"`
}

// ProductUser represents one product that the user
// has selected or added to their cart.
//
// Although it looks similar to Product,
// keeping it separate allows adding user-specific
// information later (quantity, selected size, etc.).
type ProductUser struct {
	ProductID primitive.ObjectID `json:"_id" bson:"_id"`

	ProductName string `json:"product_name" bson:"product_name"`

	ProductPrice int `json:"product_price" bson:"product_price"`

	ProductRating int `json:"product_rating" bson:"product_rating"`

	Image string `json:"image" bson:"image"`
}

// Address represents one saved delivery address.
//
// Since users may have multiple saved addresses,
// this struct is stored inside a slice.
type Address struct {

	// Unique identifier for this address.
	AddressID primitive.ObjectID `bson:"_id"`

	// House or apartment number.
	
	House string `json:"house_name"`

	// Street name.
	Street string `json:"house_street"`

	// City name.
	City string `json:"house_city"`

	// Postal PIN code.
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

// Payment represents the payment mode
// chosen by the customer.
//
// Only one of these should normally be true.
type Payment struct {

	// true if paid online
	Digital bool `json:"digital"`

	// true if Cash On Delivery
	COD bool `json:"cod"`
}
