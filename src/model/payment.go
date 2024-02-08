package model

type Sale struct {
	SaleId             string    `json:"saleId,omitempty" bson:"saleId,omitempty"`
	UserId             string    `json:"userId,omitempty" bson:"UserId,omitempty"`
    Payments 	   	   []Payment `json:"payments,omitempty" bson:"payments,omitempty"`
	Products           []Product `json:"products,omitempty" bson:"products,omitempty"`
}

type Payment struct {
	Value        string   `json:"value,omitempty" bson:"value,omitempty"`
	TypePayment  string   `json:"typePayment,omitempty" bson:"typePayment,omitempty"`
}

