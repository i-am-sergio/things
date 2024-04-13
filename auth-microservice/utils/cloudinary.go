package utils

import (
	"context"
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/joho/godotenv"
)

var CloudinaryInstance *cloudinary.Cloudinary
var CloudinaryContext context.Context

func Init() error {
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("error loading .env file: %w", err)

	}
	cloudName := os.Getenv("CLOUDINARY_CLOUD_NAME")
	apiKey := os.Getenv("CLOUDINARY_API_KEY")
	apiSecret := os.Getenv("CLOUDINARY_API_SECRET")

	cld, err := cloudinary.NewFromParams(cloudName, apiKey, apiSecret)
	if err != nil {
		return fmt.Errorf("error cloudinary setup: %w", err)
	}
	cld.Config.URL.Secure = true

	CloudinaryInstance = cld
	CloudinaryContext = context.Background()
	return nil
}

func UploadImage(file *multipart.FileHeader) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	filename := filepath.Base(file.Filename)
	filenameWithoutExt := filename[:len(filename)-len(filepath.Ext(filename))]
	currentTime := time.Now().Format("20060102150405")
	publicID := fmt.Sprintf("%s_%s", currentTime, filenameWithoutExt)
	fmt.Println(publicID)
	uploadParams := uploader.UploadParams{
		Folder:         "users/",
		PublicID:       publicID,
		UniqueFilename: api.Bool(true),
		Overwrite:      api.Bool(true),
	}

	uploadResult, err := CloudinaryInstance.Upload.Upload(CloudinaryContext, src, uploadParams)
	if err != nil {
		return "", err
	}
	return uploadResult.SecureURL, nil
}
