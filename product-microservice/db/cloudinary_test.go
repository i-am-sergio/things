package db

import (
	"errors"
	"mime/multipart"
	"testing"

	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type MockCloudinaryUploader struct {
	mock.Mock
}

func (m *MockCloudinaryUploader) Upload(uploadParams interface{}, src *multipart.FileHeader) (uploader.UploadResult, error) {
	args := m.Called(uploadParams, src)
	return args.Get(0).(uploader.UploadResult), args.Error(1)
}


func TestInitCloudinarySuccess(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockEnvLoader := new(MockEnvLoader)
		mockEnvLoader.On("LoadEnv", ".env").Return(nil)
		cloudinary := Cloudinary{}
		err := cloudinary.InitCloudinary(mockEnvLoader)
		require.NoError(t, err)
		mockEnvLoader.AssertExpectations(t)
	})
	t.Run("Failure", func(t *testing.T) {
		mockEnvLoader := new(MockEnvLoader)
		mockEnvLoader.On("LoadEnv", ".env").Return(errors.New("error loading .env file"))
		cloudinary := Cloudinary{}
		err := cloudinary.InitCloudinary(mockEnvLoader)
		require.Error(t, err)
		expectedError := "error loading .env file: error loading .env file"
		assert.Equal(t, expectedError, err.Error())
		mockEnvLoader.AssertExpectations(t)
	})
}

// func TestUploadImage(t *testing.T) {
// 	mockUploader := new(MockCloudinaryUploader)
// 	mockUploadResult := uploader.UploadResult{SecureURL: "http://example.com/image.jpg"}
// 	mockUploader.On("Upload", mock.Anything, mock.Anything).Return(mockUploadResult, nil)
// 	cloudinary := Cloudinary{Instance: MockCloudinaryUploader}
// 	file := &multipart.FileHeader{
// 		Filename: "example.jpg",
// 		Size:     1234,
// 	}
// 	url, err := cloudinary.UploadImage(file)
// 	require.NoError(t, err)
// 	assert.NotEmpty(t, url)
// }