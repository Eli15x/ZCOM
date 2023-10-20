package service

import (
	"encoding/json"
	kafka "github.com/Eli15x/ZCOM/src/client/kafka"
	//"go.mongodb.org/mongo-driver/bson"
	"context"
	"errors"
	"sync"


	"github.com/Eli15x/ZCOM/src/model"
	//"github.com/Eli15x/ZCOM/src/client"
	"github.com/Eli15x/ZCOM/src/repository"
	//"github.com/fatih/structs"
)

var (
	instanceServiceProduct ServiceProduct
	onceServiceProduct    sync.Once
)

type ServiceProduct interface {
	CreateProduct(ctx context.Context, product model.Product) error
	EditProduct(ctx context.Context, product model.Product) error
	GetProduct(ctx context.Context, id string) (model.Product, error)
	GetProductByName(ctx context.Context, name string) (model.Product,error)
	GetProducts(ctx context.Context) ([]model.Product, error)
	DeleteProduct(ctx context.Context, id string) error
}

type product struct{}

func GetInstanceProduct() ServiceProduct {
	onceServiceProduct.Do(func() {
		instanceServiceProduct = &product{}
	})
	return instanceServiceProduct
}

func (p *product) CreateProduct(ctx context.Context, product model.Product) error {

	productExist, _ := p.GetProduct(ctx, product.BarCodeNumber)
	if productExist.BarCodeNumber != "" {
		return errors.New("Product: this barcode exists")
	}
	
	//vejo se o kafka roda normalmente e está funcionando.
	//vejo se o banco está funcionando
	//se o banco nao estiver funcionando e nem o kafka
	//neste caso eu mando um erro

	productJson, err := json.Marshal(product)

	err = kafka.GetInstanceKafka().SendMessage(productJson, "createProduct")
	if err != nil {
		return err
	}

	/*productInsert := structs.Map(product)

	_, err := client.GetInstance().Insert(ctx, "product", productInsert)
	if err != nil {
		return errors.New("Product: problem to insert into MongoDB")
	}*/

	return nil
}

func (p *product) EditProduct(ctx context.Context, product model.Product) error{

	productExist, _ := p.GetProduct(ctx, product.BarCodeNumber)
	if productExist.BarCodeNumber == "" {
		return errors.New("Edit Product: doesn't have any match for this barCode")
	}

	productJson, err := json.Marshal(product) 

	err = kafka.GetInstanceKafka().SendMessage(productJson, "editProduct")
	if err != nil {
		return err
	}
	
	/*productUpdate:= structs.Map(product)
	barCode := map[string]interface{}{"BarCodeNumber": product.BarCodeNumber}
	change := bson.M{"$set": productUpdate}

	_, err := client.GetInstance().UpdateOne(ctx, "product", barCode, change)
	if err != nil {
		return errors.New("Edit product: problem to update into MongoDB")
	}*/

	return nil
}

func (p *product) GetProduct(ctx context.Context, id string) (model.Product, error) {
	var product model.Product

	//vejo se o banco está funcionando
	//se o banco nao estiver funcionando 
	//eu busco no arquivo texto o produto

	barCode := map[string]interface{}{"BarCodeNumber": id}
	product, err := repository.GetInstanceProduct().FindOne(ctx, "product", barCode)
	if err != nil {
		return product, errors.New("Get user: problem to Find Id into MongoDB")
	}

	return product, nil
}

func (p *product) DeleteProduct(ctx context.Context, id string) error{

	productExist, _ := p.GetProduct(ctx, id)
	if productExist.BarCodeNumber == "" {
		return errors.New("Delete Product: doesn't have any match for this barCode")
	}

	//vejo se o kafka roda normalmente e está funcionando.
	//vejo se o banco está funcionando
	//se o banco nao estiver funcionando e nem o kafka
	//neste caso eu mando um erro

	productJson, err := json.Marshal(productExist)

	err = kafka.GetInstanceKafka().SendMessage(productJson, "deleteProduct")
	if err != nil {
		return err
	}

	/*barCode := map[string]interface{}{"BarCodeNumber": id}

	err := client.GetInstance().Remove(ctx, "product", barCode)
	if err != nil {
		return errors.New("Delete Product: problem to delete into MongoDB")
	}*/

	return nil
}

func (p *product) GetProductByName(ctx context.Context, name string) (model.Product,error){
	var product model.Product
	Name := map[string]interface{}{"Name": name}
	product, err := repository.GetInstanceProduct().FindOne(ctx, "product", Name)
	if err != nil {
	
		return product, errors.New("Get product by name: problem to Find name into MongoDB")
	}

	return product, nil
}


func (p *product) GetProducts(ctx context.Context)([]model.Product, error){

	all := map[string]interface{}{}

	products, err := repository.GetInstanceProduct().Find(ctx, "product", all)
	if err != nil {
		return nil, errors.New("Get Products: problem to Find Id into MongoDB")
	}

	return products, nil
}