package model


type User struct {
	UserId             string    `json:"UserId,omitempty" bson:"UserId,omitempty"`
	Name               string    `json:"Name,omitempty" bson:"Name,omitempty"`
	Email              string    `json:"Email,omitempty" bson:"Email,omitempty"`
	PassWord           string    `json:"PassWord,omitempty" bson:"PassWord,omitempty"`
	IdAcess     	   int       `json:"IdAcess,omitempty" bson:"IdAcess,omitempty"`
}

type UserRequest struct {
	Name               string    `json:"name,omitempty" bson:"name"`
	Email              string    `json:"email,omitempty" bson:"email"`
	PassWord           string    `json:"passWord,omitempty" bson:"passWord"`
	IdAcess     	   int       `json:"idAcess,omitempty" bson:"idAcess"`
}