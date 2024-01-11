package model

type Payment struct {
	UserId             string    `json:"userId,omitempty" bson:"UserId,omitempty"`
	Name               string    `json:"name,omitempty" bson:"Name,omitempty"`
	Email              string    `json:"email,omitempty" bson:"Email,omitempty"`
	PassWord           string    `json:"passWord,omitempty" bson:"PassWord,omitempty"`
	IdAcess     	   int       `json:"idAcess,omitempty" bson:"IdAcess,omitempty"`
}