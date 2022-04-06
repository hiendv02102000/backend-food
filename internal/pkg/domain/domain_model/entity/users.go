package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var AutoUID int = 1

// Users struct
type Users struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	Username       string             `bson:"username,omitempty"`
	Email          string             `bson:"email,omitempty"`
	Password       string             `bson:"password,omitempty"`
	Token          *string            `bson:"token,omitempty"`
	TokenExpriedAt *time.Time         `bson:"token_expried_at,omitempty"`
}

func (u *Users) CollectionName() string {
	return "users"
}
func (u *Users) SetID(id primitive.ObjectID) {
	u.ID = id
}

// // Users struct
// type Users struct {
// 	ID             int        `gorm:"column:id;primary_key;auto_increment;not null"`
// 	Firstname      string     `gorm:"column:first_name;"`
// 	Lastname       string     `gorm:"column:last_name"`
// 	Email          string     `gorm:"column:email;not null"`
// 	Password       string     `gorm:"column:password;not null"`
// 	PhoneNumber    int        `gorm:"column:phone_number"`
// 	Decription     string     `gorm:"column:decription"`
// 	Token          *string    `gorm:"column:token"`
// 	TokenExpriedAt *time.Time `gorm:"column:token_expried_at"`
// }
