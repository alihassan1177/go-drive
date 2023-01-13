package ProductController

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/alihassan1177/ecom-backend/pkg/models"
	"github.com/alihassan1177/ecom-backend/pkg/utils"
	fileupload "github.com/alihassan1177/ecom-backend/pkg/utils/fileUpload"

	//	fileupload "github.com/alihassan1177/ecom-backend/pkg/utils/fileUpload"
	"github.com/gorilla/mux"
	"github.com/neox5/go-formdata"
)

func Index(w http.ResponseWriter, r *http.Request) {
	products := ProductModel.GetAllProducts()
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	data, _ := json.Marshal(products)
	w.Write(data)
}

func Create(w http.ResponseWriter, r *http.Request) {
	NewProduct := &ProductModel.Product{}
	utils.ParseBody(r, NewProduct)
	products := ProductModel.CreateNewProduct(NewProduct)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	data, _ := json.Marshal(products)
	w.Write(data)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 0, 0)
	products := ProductModel.DeleteProductByID(id)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	data, _ := json.Marshal(products)
	w.Write(data)
}

func Update(w http.ResponseWriter, r *http.Request) {
	UpdateProduct := &ProductModel.Product{}
	utils.ParseBody(r, UpdateProduct)
	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 0, 0)
	product, db := ProductModel.GetProductByID(id)

	if UpdateProduct.Title != "" && UpdateProduct.Thumbnail != "" {
		product.Title = UpdateProduct.Title
		product.Thumbnail = UpdateProduct.Thumbnail
	}
	db.Save(product)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	data, _ := json.Marshal(product)
	w.Write(data)
}

func CheckForm(w http.ResponseWriter, r *http.Request) {
	fd, err := formdata.Parse(r)
	if err != nil {
		panic(err)
	}
	fd.Validate("title").Required()
	if fd.HasErrors() {
		message := fmt.Sprintf("Validation Errors : %s", strings.Join(fd.Errors(), "; "))
		fmt.Fprintln(w, message)
		return
	}
  if fd.FileExists("thumbnail") {
    file := fd.GetFile("thumbnail")[0]
    fileupload.UploadFileToGoogleDrive(file)
  }
}
