package ProductModel

import (
	"fmt"

	"github.com/alihassan1177/ecom-backend/pkg"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Product struct {
	gorm.Model
  Title string `gorm:"title" json:"title"`
  Thumbnail string `gorm:"thumbnail" json:"thumbnail"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Product{})
}

func GetAllProducts() []Product {
	var Products []Product
	db.Find(&Products)
	return Products
}

func CreateNewProduct(b *Product) []Product{
  fmt.Println(b)
  db.NewRecord(b)
  db.Create(&b)
  return GetAllProducts()
} 

func GetProductByID(ID int64) (*Product, *gorm.DB){
  var book Product
  db :=db.Where("ID=?", ID).Find(&book)
  return &book, db
}

func DeleteProductByID(ID int64) []Product{
  var book Product
  db.Where("ID=?", ID).Delete(book)
  return GetAllProducts()
}


