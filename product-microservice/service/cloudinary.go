package service

import (
	"context"
	"fmt"
	"mime/multipart"
	"path/filepath"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func Credentials() (*cloudinary.Cloudinary, context.Context) {
	//CLOUDINARY_CLOUD_NAME=ddto4lnfz
	//CLOUDINARY_API_KEY=321291374866932
	//CLOUDINARY_API_SECRET=G2gxrLA7VxiwNXpgUO_8MPDR_vs

	cloudName := "ddto4lnfz"
    apiKey := "321291374866932"
    apiSecret := "G2gxrLA7VxiwNXpgUO_8MPDR_vs"
    cld, err := cloudinary.NewFromParams(cloudName, apiKey, apiSecret)
    if err != nil {
        fmt.Println("Cloudinary setup error:", err)
        return nil, nil
    }
    cld.Config.URL.Secure = true
    ctx := context.Background()
    return cld, ctx
}

func UploadImage(cld *cloudinary.Cloudinary, ctx context.Context, file *multipart.FileHeader) (string, error){
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
        Folder:          "products/",
        PublicID:        publicID,
        UniqueFilename: api.Bool(true),
        Overwrite:      api.Bool(true),
    }

	uploadResult, err := cld.Upload.Upload(ctx, src, uploadParams)
    if err != nil {
        return "", err
    }
	return uploadResult.SecureURL, nil
}