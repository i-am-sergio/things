package db

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
)

type CloudinaryClient interface {
	InitCloudinary(envLoader EnvLoader) error
	UploadImage(file *multipart.FileHeader) (string, error)
}

type Cloudinary struct {
	Instance  *cloudinary.Cloudinary
	Context   context.Context
}

func (c *Cloudinary) InitCloudinary(envLoader EnvLoader) error {
    err := envLoader.LoadEnv(".env")
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
    c.Instance = cld
	c.Context = context.Background()
    return nil
}

func (c *Cloudinary) UploadImage(file *multipart.FileHeader) (string, error){
	src, err := file.Open()
    if err != nil {
        return "", err
    }
    defer src.Close()
    
    filename := filepath.Base(file.Filename)
	filenameWithoutExt := filename[:len(filename)-len(filepath.Ext(filename))]
    currentTime := time.Now().Format("20060102150405")
    publicID := fmt.Sprintf("%s_%s", currentTime, filenameWithoutExt)
	uploadParams := uploader.UploadParams{
        Folder:             "products/",
        PublicID:           publicID,
        UniqueFilename:     api.Bool(true),
        Overwrite:          api.Bool(true),
    }
	uploadResult, err := c.Instance.Upload.Upload(c.Context, src, uploadParams)
    if err != nil {
        return "", err
    }
	return uploadResult.SecureURL, nil
}