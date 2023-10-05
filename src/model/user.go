package model


type User struct {
	UserId             string    `json:"userId,omitempty" bson:"userId,omitempty"`
	Name               string    `json:"name,omitempty" bson:"name,omitempty"`
	Email              string    `json:"email,omitempty" bson:"email,omitempty"`
	PassWord           string    `json:"passWord,omitempty" bson:"passWord,omitempty"`
	IdAcess     	   int       `json:"idAcess,omitempty" bson:"idAcess,omitempty"`
}

type UserRequest struct {
	Name               string    `json:"name,omitempty" bson:"name"`
	Email              string    `json:"email,omitempty" bson:"email"`
	PassWord           string    `json:"passWord,omitempty" bson:"passWord"`
	IdAcess     	   int       `json:"idAcess,omitempty" bson:"idAcess"`
}