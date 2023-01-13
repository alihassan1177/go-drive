package fileupload

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"

	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/jwt"
	"google.golang.org/api/drive/v3"
)

// Use Service account
func ServiceAccount(secretFile string) *http.Client {
	b, err := ioutil.ReadFile(secretFile)
	if err != nil {
		log.Fatal("error while reading the credential file", err)
	}
	var s = struct {
		Email      string `json:"client_email"`
		PrivateKey string `json:"private_key"`
	}{}
	json.Unmarshal(b, &s)
	config := &jwt.Config{
		Email:      s.Email,
		PrivateKey: []byte(s.PrivateKey),
		Scopes: []string{
			drive.DriveScope,
		},
		TokenURL: google.JWTTokenURL,
	}
	client := config.Client(context.Background())
	return client
}

func createFile(service *drive.Service, name string, mimeType string, content io.Reader, parentId string) (*drive.File, error) {
	f := &drive.File{
		MimeType: mimeType,
		Name:     name,
		Parents:  []string{parentId},
	}
	file, err := service.Files.Create(f).Media(content).Do()

	if err != nil {
		log.Println("Could not create file: " + err.Error())
		return nil, err
	}

	return file, nil
}

func SaveFile(file *multipart.FileHeader) string {
	contentType := file.Header.Get("Content-Type")
	fmt.Println(contentType)
	data, err := file.Open()
	if err != nil {
		panic(err)
	}
	defer data.Close()
	var osFile *os.File
	if contentType == "image/png" {
		osFile, err = ioutil.TempFile("images", "*.png")
		fmt.Println("PNG FILE :", osFile, err)
	} else {
		osFile, err = ioutil.TempFile("images", "*.jpg")
		fmt.Println("JPG FILE :", osFile, err)
	}
	defer osFile.Close()
	fileBytes, err := io.ReadAll(data)
	if err != nil {
		panic(err)
	}
	osFile.Write(fileBytes)
  fmt.Println("Your file is saved :", osFile.Name())
  return osFile.Name()
}

func UploadFileToGoogleDrive(data *multipart.FileHeader) {
	// Step 1: Open  file
  filename := SaveFile(data)
	f, err := os.Open(filename)

	if err != nil {
		panic(fmt.Sprintf("cannot open file: %v", err))
	}

	defer f.Close()

	// Step 2: Get the Google Drive service
	client := ServiceAccount("service.json")

	srv, err := drive.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve drive Client %v", err)
	}

	//give your folder id here in which you want to upload or create new directory
	folderId := "17yMwcYmaVQ7DwlLQuYWZnZTKOAStms8_"

	// Step 4: create the file and upload
	file, err := createFile(srv, f.Name(), "application/octet-stream", f, folderId)

	if err != nil {
		panic(fmt.Sprintf("Could not create file: %v\n", err))
	}

	fmt.Printf("File '%s' successfully uploaded", file.Name)
  fmt.Println("File is :", file)
	fmt.Printf("\nFile Id: '%s' ", file.Id)
}
