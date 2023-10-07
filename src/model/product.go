package model


type Product struct {
	BarCodeNumber   string       `json:"barCodeNumber,omitempty" bson:"BarCodeNumber"`
	Name            string       `json:"name,omitempty" bson:"Name"`
	NCM             string       `json:"ncm,omitempty" bson:"NCM"`
	CFOP     	    string       `json:"cfop,omitempty" bson:"CFOP"`
	Desconto     	float64          `json:"desconto,omitempty" bson:"Desconto"`
	OutrosDesconto  float64          `json:"outrosDesconto,omitempty" bson:"OutrosDesconto"`
	IndRegra     	string       `json:"indRegra,omitempty" bson:"IndRegra"`
	UCom     	   	float64       	 `json:"uCom,omitempty" bson:"UCom"`
	QCom     	   	float64       `json:"qCom,omitempty" bson:"QCom"`
	VUnCom     	   	float64       `json:"vUnCom,omitempty" bson:"VUnCom"`
	Imposto     	map[string]interface{}   `json:"imposto,omitempty" bson:"Imposto"`
}
