package service

import (
	//"fmt"
	"os"
	"encoding/json"
	kafka "github.com/Eli15x/ZCOM/src/client/kafka"
	//"go.mongodb.org/mongo-driver/bson"
	"context"
	"errors"
	"sync"


	"github.com/Eli15x/ZCOM/src/model"
	"github.com/Eli15x/ZCOM/src/client"
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
	SaveProduct(ctx context.Context) error
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

	productJson, err := json.Marshal(product)

	err = kafka.GetInstanceKafka().SendMessage(productJson, "createProduct")
	if err != nil {
		return err
	}

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

	return nil
}

func (p *product) GetProduct(ctx context.Context, id string) (model.Product, error) {
	var product model.Product

	/*if err := client.GetInstance().Initialize(context.Background()); err == nil {
		fmt.Println("entrou")
		barCode := map[string]interface{}{"BarCodeNumber": id}
		product, err := repository.GetInstanceProduct().FindOne(ctx, "product", barCode)
		if err != nil {
			return product, errors.New("Get user: problem to Find Id into MongoDB")
		}
	} else*/
		namefile := id + ".txt" 
		data, err := os.ReadFile(os.Getenv("SaveProduct")+ namefile )
		if err != nil {
			return product, err
		}
		json.Unmarshal([]byte(data), &product)
	

	return product, nil
}

func (p *product) DeleteProduct(ctx context.Context, id string) error{

	productExist, _ := p.GetProduct(ctx, id)
	if productExist.BarCodeNumber == "" {
		return errors.New("Delete Product: doesn't have any match for this barCode")
	}

	productJson, err := json.Marshal(productExist)

	err = kafka.GetInstanceKafka().SendMessage(productJson, "deleteProduct")
	if err != nil {
		return err
	}

	return nil
}

func (p *product) GetProductByName(ctx context.Context, name string) (model.Product,error){
	var product model.Product
	Name := map[string]interface{}{"Name": name}
	if err := client.GetInstance().Initialize(context.Background()); err == nil {
		product, err := repository.GetInstanceProduct().FindOne(ctx, "product", Name)
		if err != nil {
		
			return product, errors.New("Get product by name: problem to Find name into MongoDB")
		}
	}

	return product, nil
}


func (p *product) GetProducts(ctx context.Context)([]model.Product, error){

	products := []model.Product{}
	all := map[string]interface{}{}

	if err := client.GetInstance().Initialize(context.Background()); err == nil {
		products, err = repository.GetInstanceProduct().Find(ctx, "product", all)
		if err != nil {
			return nil, errors.New("Get Products: problem to Find Id into MongoDB")
		}
		
	}

	return products, nil
}

func (p *product) SaveProduct(ctx context.Context) error{

	products, err := p.GetProducts(ctx)
	if err != nil {
		return errors.New("Get Products: problem to Find Id into MongoDB")
	}

	//remove all files for not be duplicated or for not exists more files that are not related to what have on mongo
	err = os.RemoveAll(os.Getenv("SaveProduct"))
    if err != nil {
        return err
    }

	err = os.Mkdir(os.Getenv("SaveProduct"), 0755) //create a directory and give it required permissions
	if err != nil {
	   return err
	}

	for _, product := range products {
		barCodeNumber := product.BarCodeNumber
		productJson, err := json.Marshal(product)
		if err != nil {
			return err
		}

		namefile := barCodeNumber + ".txt" 
	
		if err = os.WriteFile(os.Getenv("SaveProduct") + namefile , productJson, 0666); err != nil {
			return err
		}
	}

	return nil
}