package util

import (
	"log"
	"os"

	"github.com/cloudinary/cloudinary-go"
)

var CLD *cloudinary.Cloudinary

func ConnectCloudinary() {
	c, err := cloudinary.NewFromParams(os.Getenv("CLOUDAINARY_CLOUD_NAME"), os.Getenv("CLOUDAINARY_API_KEY"), os.Getenv("CLOUDAINARY_API_SECRET"))
	if err != nil {
		log.Fatalf("Failed to intialize Cloudinary, %v", err)
		panic(err)
	}

	CLD = c
}
