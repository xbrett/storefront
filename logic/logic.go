package logic

import (
	"errors"
	"log"
	"strconv"

	"github.com/rue-brettadcock/storefront/database"
)

//Logic contains a pointer to a database instance
type logic struct {
	mydb database.SKUDataAccess
}

//Logic explicit interface for ioc
type Logic interface {
	AddProductSKU(SKU) error
	UpdateProductQuantity(SKU) error
	DeleteID(SKU) error
	PrintAllProductInfo() string
	GetProductInfo(SKU) (string, error)
}

//SKU for holding product information
type SKU struct {
	ID       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Vendor   string `json:"vendor,omitempty"`
	Quantity string `json:"quantity,omitempty"`
}

//New creates a new logic pointer to the database layer
func New() Logic {
	l := logic{mydb: database.NewInMemoryDB()}
	return &l
}

//AddProductSKU validates product info and Inserts into the db
func (l *logic) AddProductSKU(sku SKU) error {
	id, _ := strconv.Atoi(sku.ID)
	quant, _ := strconv.Atoi(sku.Quantity)
	if l.mydb.Get(id) != "[]" {
		return errors.New("Product id already exists")
	}
	if quant < 1 {
		return errors.New("Quantity must be at least 1")
	}
	if id < 0 {
		return errors.New("ID must be positive")
	}

	err := l.mydb.Insert(id, sku.Name, sku.Vendor, quant)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

//UpdateProductQuantity updates quantity for a given id
func (l *logic) UpdateProductQuantity(sku SKU) error {
	id, _ := strconv.Atoi(sku.ID)
	quant, _ := strconv.Atoi(sku.Quantity)
	if l.mydb.Get(id) == "[]" {
		return errors.New("Product id doesn't exist")
	}

	err := l.mydb.Update(id, quant)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

//DeleteID removes all product information for a given id
func (l *logic) DeleteID(sku SKU) error {
	id, _ := strconv.Atoi(sku.ID)
	if l.mydb.Get(id) == "[]" {
		return errors.New("Product id doesn't exist")
	}
	err := l.mydb.Delete(id)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

//PrintAllProductInfo returns all product SKUs
func (l *logic) PrintAllProductInfo() string {
	return l.mydb.Print()
}

//GetProductInfo returns product details for given id
func (l *logic) GetProductInfo(sku SKU) (string, error) {
	id, _ := strconv.Atoi(sku.ID)
	if l.mydb.Get(id) == "[]" {
		return "[]", errors.New("Product id doesn't exist")
	}
	return l.mydb.Get(id), nil
}
