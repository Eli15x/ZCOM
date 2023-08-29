package model

type user struct {
	UserId             string    `json:"userId,omitempty" bson:"userId,omitempty"`
	Name               string    `json:"name,omitempty" bson:"name"`
	Email              string    `json:"email,omitempty" bson:"email"`
	PassWord           string    `json:"passWord,omitempty" bson:"passWord"`
	IdAcess     	   int       `json:"idAcess,omitempty" bson:"IdAcess"`
}

type userRequest struct {
	Name               string    `json:"name,omitempty" bson:"name"`
	Email              string    `json:"email,omitempty" bson:"email"`
	PassWord           string    `json:"passWord,omitempty" bson:"passWord"`
	IdAcess     	   int       `json:"idAcess,omitempty" bson:"idAcess"`
}