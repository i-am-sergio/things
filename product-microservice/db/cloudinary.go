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

type CloudinaryAPI interface {
	NewFromParams(cloudName, apiKey, apiSecret string) (*cloudinary.Cloudinary, error)
}
type CloudinaryService struct{}
func (cs *CloudinaryService) NewFromParams(cloudName, apiKey, apiSecret string) (*cloudinary.Cloudinary, error) {
	return cloudinary.NewFromParams(cloudName, apiKey, apiSecret)
}

type CloudinaryUploader interface {
	Upload(ctx context.Context, file multipart.File, params uploader.UploadParams) (*uploader.UploadResult, error)
}
type CloudinaryUploaderAdapter struct {
	Cld *cloudinary.Cloudinary
}
func (cua *CloudinaryUploaderAdapter) Upload(ctx context.Context, file multipart.File, params uploader.UploadParams) (*uploader.UploadResult, error) {
	uploadResult, err := cua.Cld.Upload.Upload(ctx, file, params)
	if err != nil {
		return nil, err
	}
	return uploadResult, nil
}

type FileHeaderWrapper interface {
	Open() (multipart.File, error)
    Filename() string
}
type MultipartFileHeaderAdapter struct {
    *multipart.FileHeader
}
func (fha *MultipartFileHeaderAdapter) Open() (multipart.File, error) {
    return fha.FileHeader.Open()
}
func (fha *MultipartFileHeaderAdapter) Filename() string {
    return fha.FileHeader.Filename
}


type CloudinaryClient interface {
	InitCloudinary(envLoader EnvLoader) error
	UploadImage(file FileHeaderWrapper) (string, error)
}

type Cloudinary struct {
	Uploader  CloudinaryUploader
	Context   context.Context
	API       CloudinaryAPI
}

func (c *Cloudinary) InitCloudinary(envLoader EnvLoader) error {
    err := envLoader.LoadEnv(".env")
	if err != nil {
		return fmt.Errorf("error loading .env file: %w", err)
	}
    cloudName := os.Getenv("CLOUDINARY_CLOUD_NAME")
	apiKey := os.Getenv("CLOUDINARY_API_KEY")
	apiSecret := os.Getenv("CLOUDINARY_API_SECRET")
    cld, err := c.API.NewFromParams(cloudName, apiKey, apiSecret)
    if err != nil {
        return fmt.Errorf("error cloudinary setup: %w", err)
    }
    cld.Config.URL.Secure = true
    c.Uploader = &CloudinaryUploaderAdapter{Cld: cld}
	c.Context = context.Background()
    return nil
}

func (c *Cloudinary) UploadImage(file FileHeaderWrapper) (string, error){
	src, err := file.Open()
    if err != nil {
        return "", err
    }
    defer src.Close()
    filename := filepath.Base(file.Filename())
	filenameWithoutExt := filename[:len(filename)-len(filepath.Ext(filename))]
    currentTime := time.Now().Format("20060102150405")
    publicID := fmt.Sprintf("%s_%s", currentTime, filenameWithoutExt)
	uploadParams := uploader.UploadParams{
        Folder:             "products/",
        PublicID:           publicID,
        UniqueFilename:     api.Bool(true),
        Overwrite:          api.Bool(true),
    }
	uploadResult, err := c.Uploader.Upload(c.Context, src, uploadParams)
    if err != nil {
        return "", err
    }
	return uploadResult.SecureURL, nil
}