package models

//Here each struct is a representation of a collection
import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//annotation tells golang how would it look in our json
//we would also need marshaling and unmarshaling to actually assign the particular values

type User struct {
	//for struct the item needs to be captilalized for export but in db we have small case
	//go lang does not understand json directly hence we look for
	ID            primitive.ObjectID `json:_id bson:_id`
	First_Name    string             `json:first_name ` //binds this name for the struct when sending and receving this would be used
	Last_Name     string             `json:last_name`
	Password      string             `json:password`
	Email         string             `json:email`
	Phone         string             `json:phone`
	Token         string             `json:token`
	Refresh_Token string             `json:refresh_token`
	Created_At    time.Time          `json:created_at`
	Updated_At    time.Time          `json:updated_at`
	UserID        string             `json:user_id`
	Usercart      []ProductUser      `json:user_cart bson:user_cart` //bson also used as referring to many structs
	//user can have many order
	Address_details []Address `json:address bson:address`
	OrderStatus     []Order   `json:address bson:address`
}
type Product struct {
	Product_ID     primitive.ObjectID `json:_id bson:_id`
	Product_name   string             `json:product_name`
	Product_Price  int                `json:product_price`
	Product_rating int                `json:product_rating`
	Image          string             `json:image` //using a url
}

//Individual product items user has selected

type ProductUser struct {
	Product_ID     primitive.ObjectID `json:_id bson:_id`
	Product_name   string             `json:product_name`
	Product_Price  int                `json:product_price`
	Product_rating int                `json:product_rating`
	Image          string             `json:image` //using a url
}
type Address struct {
	Address_id primitive.ObjectID `bson:_id`
	House      string             `json:house_name`
	Street     string             `json:house_street`
	City       string             `json:house_city`
	Pincode    int                `json:house_pincode`
}
type Order struct {
	Order_ID       primitive.ObjectID `json:_id bson:_id`
	Order_Cart     []ProductUser      `json:order_cart`
	Ordered_At     time.Time          `json:ordered_at`
	Price          int                `json:price`
	Discount       int                `json:discount`
	Payment_Method Paymemt            `json:payment_method`
}
type Paymemt struct {
	Digital bool `json:digital`
	COD     bool `json:cod`
}
