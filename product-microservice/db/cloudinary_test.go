package db

import (
	"context"
	"errors"
	"mime/multipart"
	"os"
	"testing"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

const image = "example.jpg"
const example = "https://example.com/image.jpg"

type MockCloudinaryAPI struct {
	mock.Mock
}
func(m *MockCloudinaryAPI) NewFromParams(cloudName, apiKey, apiSecret string) (*cloudinary.Cloudinary, error) {
	args := m.Called(cloudName, apiKey, apiSecret)
	return args.Get(0).(*cloudinary.Cloudinary), args.Error(1)
}

type MockUploader struct {
	mock.Mock
}
func (mu *MockUploader) Upload(ctx context.Context, file multipart.File, params uploader.UploadParams) (*uploader.UploadResult, error) {
	args := mu.Called(ctx, file, params)
	return args.Get(0).(*uploader.UploadResult), args.Error(1)
}

type MockFileHeaderWrapper struct {
    mock.Mock
}
func (m *MockFileHeaderWrapper) Open() (multipart.File, error) {
    args := m.Called()
    return args.Get(0).(multipart.File), args.Error(1)
}
func (m *MockFileHeaderWrapper) Filename() string {
    args := m.Called()
    return args.String(0)
}

func TestNewFromParams(t *testing.T) {
	cloud := CloudinaryService{}
	_, err := cloud.NewFromParams("testCloudName", "testAPIKey", "testAPISecret")
	require.NoError(t, err)
}

func TestUpload(t *testing.T) {
	adapter := &CloudinaryUploaderAdapter{
		Cld: &cloudinary.Cloudinary{},
	}
	file := &os.File{}
	params := uploader.UploadParams{
		Folder: "testFolder",
	}
	_, err := adapter.Upload(context.Background(), file, params)
	require.Error(t, err)
}

func TestOpen(t *testing.T) {
	adapter := &MultipartFileHeaderAdapter{
		&multipart.FileHeader{},
	}
	_, err := adapter.Open()
	require.Error(t, err)
}

func TestFilename(t *testing.T) {
	adapter := &MultipartFileHeaderAdapter{
		&multipart.FileHeader{Filename: "test.jpg"},
	}
	filename := adapter.Filename()
	assert.Equal(t, "test.jpg", filename)
}

func TestInitCloudinarySuccess(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		os.Setenv("CLOUDINARY_CLOUD_NAME", "testCloudName")
		os.Setenv("CLOUDINARY_API_KEY", "testAPIKey")
		os.Setenv("CLOUDINARY_API_SECRET", "testAPISecret")
		defer func() {
			os.Unsetenv("CLOUDINARY_CLOUD_NAME")
			os.Unsetenv("CLOUDINARY_API_KEY")
			os.Unsetenv("CLOUDINARY_API_SECRET")
		}()
		mockEnvLoader := new(MockEnvLoader)
		mockEnvLoader.On("LoadEnv", ".env").Return(nil)
		mockCloudinaryAPI := new(MockCloudinaryAPI)
		mockCloudinaryAPI.On("NewFromParams", "testCloudName", "testAPIKey", "testAPISecret").Return(&cloudinary.Cloudinary{}, nil)
		cloudinary := Cloudinary{
			API: mockCloudinaryAPI,
		}
		err := cloudinary.InitCloudinary(mockEnvLoader)
		require.NoError(t, err)
		mockEnvLoader.AssertExpectations(t)
		mockCloudinaryAPI.AssertExpectations(t)
	})
	t.Run("Failure DotEnv Setup", func(t *testing.T) {
		mockEnvLoader := new(MockEnvLoader)
		mockEnvLoader.On("LoadEnv", ".env").Return(errors.New("error loading .env file"))
		cloudinary := Cloudinary{}
		err := cloudinary.InitCloudinary(mockEnvLoader)
		require.Error(t, err)
		expectedError := "error loading .env file: error loading .env file"
		assert.Equal(t, expectedError, err.Error())
		mockEnvLoader.AssertExpectations(t)
	})
	t.Run("Failure Cloudinary Setup", func(t *testing.T) {
		os.Setenv("CLOUDINARY_CLOUD_NAME", "testCloudName")
		os.Setenv("CLOUDINARY_API_KEY", "testAPIKey")
		os.Setenv("CLOUDINARY_API_SECRET", "testAPISecret")
		defer func() {
			os.Unsetenv("CLOUDINARY_CLOUD_NAME")
			os.Unsetenv("CLOUDINARY_API_KEY")
			os.Unsetenv("CLOUDINARY_API_SECRET")
		}()
		mockEnvLoader := new(MockEnvLoader)
		mockEnvLoader.On("LoadEnv", ".env").Return(nil)
		mockCloudinaryAPI := new(MockCloudinaryAPI)
		mockCloudinaryAPI.On("NewFromParams", "testCloudName", "testAPIKey", "testAPISecret").
			Return(&cloudinary.Cloudinary{}, errors.New("error creating cloudinary instance"))
		cloudinary := Cloudinary{
			API: mockCloudinaryAPI,
		}
		err := cloudinary.InitCloudinary(mockEnvLoader)
		require.Error(t, err)
		expectedError := "error cloudinary setup: error creating cloudinary instance"
		assert.Equal(t, expectedError, err.Error())
		mockEnvLoader.AssertExpectations(t)
		mockCloudinaryAPI.AssertExpectations(t)
	})
}

func TestUploadImage(t *testing.T) {
	t.Run("Successful upload", func(t *testing.T) {
		mockUploader := new(MockUploader)
		uploadResult := &uploader.UploadResult{SecureURL: example}
		mockUploader.On("Upload", mock.Anything, mock.Anything, mock.Anything).Return(uploadResult, nil)
		mockFileHeaderWrapper := new(MockFileHeaderWrapper)
        file := &os.File{}
        mockFileHeaderWrapper.On("Open").Return(file, nil)
        mockFileHeaderWrapper.On("Filename").Return(image)
		cloudinary := Cloudinary{
			Uploader: mockUploader,
			Context:  context.Background(),
			API:      &CloudinaryService{},
		}
		url, err := cloudinary.UploadImage(mockFileHeaderWrapper)
		require.NoError(t, err)
		assert.Equal(t, example, url)
		mockUploader.AssertExpectations(t)
		mockFileHeaderWrapper.AssertExpectations(t)
	})
	t.Run("File open error", func(t *testing.T) {
		mockUploader := new(MockUploader)
		mockFileHeaderWrapper := new(MockFileHeaderWrapper)
		file := &os.File{}
        mockFileHeaderWrapper.On("Open").Return(file, errors.New("file open error"))
		cloudinary := Cloudinary{
			Uploader: mockUploader,
			Context:  context.Background(),
			API:      &CloudinaryService{},
		}
		_, err := cloudinary.UploadImage(mockFileHeaderWrapper)
		require.Error(t, err)
		assert.Equal(t, "file open error", err.Error())
		mockUploader.AssertExpectations(t)
		mockFileHeaderWrapper.AssertExpectations(t)
	})
	t.Run("Upload Cloudinary error", func(t *testing.T) {
		mockUploader := new(MockUploader)
		uploadResult := &uploader.UploadResult{SecureURL: example}
		mockUploader.On("Upload", mock.Anything, mock.Anything, mock.Anything).Return(uploadResult, errors.New("upload error"))
		mockFileHeaderWrapper := new(MockFileHeaderWrapper)
        file := &os.File{}
        mockFileHeaderWrapper.On("Open").Return(file, nil)
        mockFileHeaderWrapper.On("Filename").Return(image)
		cloudinary := Cloudinary{
			Uploader: mockUploader,
			Context:  context.Background(),
			API:      &CloudinaryService{},
		}
		_, err := cloudinary.UploadImage(mockFileHeaderWrapper)
		require.Error(t, err)
		assert.Equal(t, "upload error", err.Error())
		mockUploader.AssertExpectations(t)
		mockFileHeaderWrapper.AssertExpectations(t)
	})
}