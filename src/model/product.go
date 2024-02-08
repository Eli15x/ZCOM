package model


type Product struct {
	CODIGO_CEST                string `json:"CODIGO_CEST,omitempty" bson:"CODIGO_CEST"`
	GTIN            		   string `json:"GTIN,omitempty" bson:"GTIN"`
	IAT                        string `json:"IAT,omitempty" bson:"IAT"`
	ID_PRODUTO_GRUPO           int `json:"ID_PRODUTO_GRUPO,omitempty" bson:"ID_PRODUTO_GRUPO"`
	ID_PRODUTO_MARCA           int `json:"ID_PRODUTO_MARCA,omitempty" bson:"ID_PRODUTO_MARCA"`
	ID_PRODUTO_UNIDADE         int `json:"ID_PRODUTO_UNIDADE,omitempty" bson:"ID_PRODUTO_UNIDADE"`
	ID_TRIBUT_GRUPO_TRIBUTARIO int `json:"ID_TRIBUT_GRUPO_TRIBUTARIO,omitempty" bson:"ID_TRIBUT_GRUPO_TRIBUTARIO"`
	IPPT  					   string `json:"IPPT,omitempty" bson:"IPPT"`
	PESO  					   int `json:"PESO,omitempty" bson:"PESO"`
	QUANTIDADE_EMBALAGEM       int `json:"QUANTIDADE_EMBALAGEM,omitempty" bson:"QUANTIDADE_EMBALAGEM"`
	QUANTIDADE_ESTOQUE         int  `json:"QUANTIDADE_ESTOQUE,omitempty" bson:"QUANTIDADE_ESTOQUE"`
	NAME            string       `json:"NAME,omitempty" bson:"NAME"`
	CODIGO_NCM      string       `json:"CODIGO_NCM,omitempty" bson:"CODIGO_NCM"`
	CFOP     	    string       `json:"CFOP,omitempty" bson:"CFOP"`
	Desconto     	float64          `json:"desconto,omitempty" bson:"Desconto"`
	OutrosDesconto  float64          `json:"outrosDesconto,omitempty" bson:"OutrosDesconto"`
	IndRegra     	string       `json:"indRegra,omitempty" bson:"IndRegra"`
	//UCom     	   	float64       	 `json:"uCom,omitempty" bson:"UCom"`
	//QCom     	   	float64       `json:"qCom,omitempty" bson:"QCom"`
	//VUnCom     	   	float64       `json:"vUnCom,omitempty" bson:"VUnCom"`
	Imposto     	map[string]interface{}   `json:"imposto,omitempty" bson:"Imposto"`
}

type Unidade struct {
	Descricao string 		`json:"DESCRICAO,omitempty" bson:"DESCRICAO"`
    Id        int    		`json:"ID,omitempty" bson:"ID"`
    Pode_Fracionar string   `json:"PODE_FRACIONAR,omitempty" bson:"PODE_FRACIONAR"`
    Sigla string  			`json:"SIGLA,omitempty" bson:"SIGLA"` 
}

type Marca struct {
	Id   int    		`json:"ID,omitempty" bson:"ID"`
	Nome string 		`json:"NOME,omitempty" bson:"NOME"`
}

type Grupo struct {
	Id   int    		`json:"ID,omitempty" bson:"ID"`
	Grupo string 		`json:"GRUPO,omitempty" bson:"GRUPO"`
}


