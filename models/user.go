package models

type User struct {
	Name      string   `json:"name" bson:"user_name"`
	Age       int      `json:"age" bson:"user_age"`
	Gender    string   `json:"gender" bson:"user_gender"`
	Pin       string   `json:"pin" bson:"user_pin"`
	AccountNo int      `json:"no" bson:"user_no"`
	Balance   string   `json:"balance" bson:"user_balance"`
	Statement []string `json:"statement" bson:"user_statement"`
}
