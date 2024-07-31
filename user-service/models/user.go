package models

type User struct {
    UserID   string `bson:"_id,omitempty"`
    Username string `bson:"username,omitempty"`
    Email    string `bson:"email,omitempty"`
    Password string `bson:"password,omitempty"`
}
