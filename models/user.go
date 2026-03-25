package models

type User struct {
	UserId   string `bson:"user_id,omitempty"`
	Name     string `bson:"name,omitempty"`
	Age      int32  `bson:"age,omitempty"`
	Email    string `bson:"email,omitempty"`
	Password string `bson:"password,omitempty"`
}
