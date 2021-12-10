package util


import (
	"context"
	"log"
	"os"
	"strings"

	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/cloudinary/cloudinary-go"
)

var CLD *cloudinary.Cloudinary

func ConnectCloudinary() {
	c, err := cloudinary.NewFromParams(
		os.Getenv("CLOUDAINARY_CLOUD_NAME"),
		os.Getenv("CLOUDAINARY_API_KEY"),
		os.Getenv("CLOUDAINARY_API_SECRET"),
	)
	if err != nil {
		log.Fatalf("Failed to intialize Cloudinary, %v", err)
		panic(err)
	}

	CLD = c
}

func UploadImage(img_uri string) (url string, err error) {
	img_uri = strings.TrimSpace(img_uri)
	if len(img_uri) == 0 {
		return "", err
	}

	var ctx = context.Background()
	uploadResult, err := CLD.Upload.Upload(
		ctx,
		img_uri,
		uploader.UploadParams{Folder: "images/"},
	)
	if err != nil {
		return "", err
	}

	return uploadResult.SecureURL, nil
}