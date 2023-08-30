package model


type Procut struct {
	BarCodeNumber   string       `json:"barCodeNumber,omitempty" bson:"barCodeNumber"`
	Name            string       `json:"name,omitempty" bson:"name"`
	NCM             string       `json:"ncm,omitempty" bson:"ncm"`
	CFOP     	    string       `json:"cfop,omitempty" bson:"cfop"`
	Desconto     	int          `json:"desconto,omitempty" bson:"desconto"`
	OutrosDesconto  int          `json:"outrosDesconto",omitempty" bson:"outrosDesconto"`
	IndRegra     	string       `json:"indRegra,omitempty" bson:"indRegra"`
	UCom     	   	int       	 `json:"uCom,omitempty" bson:"uCom"`
	QCom     	   	string       `json:"qCom,omitempty" bson:"qCom"`
	VUnCom     	   	string       `json:"vUnCom,omitempty" bson:"vUnCom"`
	Imposto     	map[string]interface{}   `json:"imposto,omitempty" bson:"imposto"`
}
