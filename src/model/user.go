package model


type User struct {
	UserId             string    `json:"userId,omitempty" bson:"UserId,omitempty"`
	Name               string    `json:"name,omitempty" bson:"Name,omitempty"`
	Email              string    `json:"email,omitempty" bson:"Email,omitempty"`
	PassWord           string    `json:"passWord,omitempty" bson:"PassWord,omitempty"`
	IdAcess     	   int       `json:"idAcess,omitempty" bson:"IdAcess,omitempty"`
}

type UserRequest struct {
	Name               string    `json:"name,omitempty" bson:"name"`
	Email              string    `json:"email,omitempty" bson:"email"`
	PassWord           string    `json:"passWord,omitempty" bson:"passWord"`
	IdAcess     	   int       `json:"idAcess,omitempty" bson:"idAcess"`
}

//ver redundancia e se realmente é necessario o userRequest