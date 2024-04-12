package services

import (
	"fmt"
	"mime/multipart"
	"product-microservice/db"
	"product-microservice/models"
	"product-microservice/repository"
	"strconv"
)

type ProductService interface {
	CreateProductService(form *multipart.Form, image *multipart.FileHeader) (models.Product, error)
	UpdateProductService(productID uint, form *multipart.Form, image *multipart.FileHeader) (models.Product, error)
	GetProductsService() ([]models.Product, error)
	GetProductByIDService(id uint) (models.Product, error)
	GetProductsByCategoryService(category string) ([]models.Product, error)
	SearchProductsService(searchTerm string) ([]models.Product, error)
	DeleteProductService(productID uint) error
	PremiumService(productID uint) (models.Product, error)
	GetProductsPremiumService() ([]models.Product, error)
}
type ProductServiceImpl struct {
	dbClient repository.DBInterface
	cloudinaryClient db.CloudinaryClient
}
func NewProductService(client repository.DBInterface, cloudinary db.CloudinaryClient) *ProductServiceImpl {
	return &ProductServiceImpl{dbClient: client, cloudinaryClient: cloudinary}
}

func validateRequiredFields(form *multipart.Form) error {
    requiredFields := []string{"UserID", "Price", "Name", "Description", "Category", "Ubication"}
    for _, field := range requiredFields {
        if len(form.Value[field]) == 0 {
            return fmt.Errorf("missing required field: %s", field)
        }
    }
    return nil
}

func (c *ProductServiceImpl) CreateProductService(form *multipart.Form, image *multipart.FileHeader) (models.Product, error){
	var product models.Product
	if err := validateRequiredFields(form); err != nil {
        return product, err
    }
	userID, err := strconv.ParseUint(form.Value["UserID"][0], 10, 32)
	if err != nil {
		return product, err
	}
	price, err := strconv.ParseFloat(form.Value["Price"][0], 64)
	if err != nil {
		return product, err
	}
	cloudinaryURL, err := c.cloudinaryClient.UploadImage(&db.MultipartFileHeaderAdapter{FileHeader: image})
	if err != nil {
		return product, err
	}
	product = models.Product{
		UserID:      uint(userID),
		Name:        form.Value["Name"][0],
		Description: form.Value["Description"][0],
		Category:    form.Value["Category"][0],
		Price:       float64(price),
		Rate:        0.0,
		Ubication:   form.Value["Ubication"][0],
		Image:       cloudinaryURL,
	}
	if err := c.dbClient.Create(&product); err != nil {
		return product, err
	}
	return product, nil
}

func (c *ProductServiceImpl) UpdateProductService(productID uint, form *multipart.Form, image *multipart.FileHeader) (models.Product, error){
	var product models.Product
    if err := db.Client.First(&product, productID); err != nil {
        return product, err
    }
	if err := validateRequiredFields(form); err != nil {
		return product, err
	}
	userID, err := strconv.ParseUint(form.Value["UserID"][0], 10, 32)
	if err != nil {
		return product, err
	}
	price, err := strconv.ParseFloat(form.Value["Price"][0], 64)
	if err != nil {
		return product, err
	}
	product.UserID = uint(userID)
	product.Name = form.Value["Name"][0]
	product.Description = form.Value["Description"][0]
	product.Category = form.Value["Category"][0]
	product.Price = float64(price)
	product.Ubication = form.Value["Ubication"][0]
	if image != nil {
		cloudinaryURL, err := c.cloudinaryClient.UploadImage(&db.MultipartFileHeaderAdapter{FileHeader: image})
		if err != nil {
			return product, err
		}
		product.Image = cloudinaryURL
	}
	if err := c.dbClient.Save(&product); err != nil {
		return product, err
	}
	return product, nil
}

func (c *ProductServiceImpl) GetProductsService() ([]models.Product, error) {
	var products []models.Product
	if err := c.dbClient.Find(&products); err != nil {
		return nil, err
	}
	return products, nil
}

func (c *ProductServiceImpl) GetProductByIDService(id uint) (models.Product, error) {
	var product models.Product
	if err := c.dbClient.First(&product, id); err != nil {
		return product, err
	}
	return product, nil
}

func (c *ProductServiceImpl) GetProductsByCategoryService(category string) ([]models.Product, error) {
	var products []models.Product
	if category != "" {
		err := c.dbClient.FindWithCondition(&products, "category = ?", category)
		if err != nil {
			return nil, err
		}
	} else {
		err := c.dbClient.Find(&products)
		if err != nil {
			return nil, err
		}
	}
	return products, nil
}

func (c *ProductServiceImpl) SearchProductsService(searchTerm string) ([]models.Product, error) {
	var products []models.Product
	err := c.dbClient.FindWithCondition(&products, "name LIKE ? OR description LIKE ?", "%"+searchTerm+"%", "%"+searchTerm+"%")
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (c *ProductServiceImpl) DeleteProductService(productID uint) error {
	if err := c.dbClient.DeleteWithCondition(&models.Comment{}, "product_id = ?", productID); err != nil {
		return err
	}
	if err := c.dbClient.DeleteByID(&models.Product{}, productID); err != nil {
		return err
	}
	return nil
}

func (c *ProductServiceImpl) PremiumService(productID uint) (models.Product, error) {
	var product models.Product
	if err := c.dbClient.First(&product, productID); err != nil {
		return product, err
	}
	product.Status = !product.Status
	if err := c.dbClient.Save(&product); err != nil {
		return product, err
	}
	return product, nil
}

func (c *ProductServiceImpl) GetProductsPremiumService() ([]models.Product, error) {
	var products []models.Product
	err := c.dbClient.FindWithCondition(&products, "status = ?", true)
	if err != nil {
		return nil, err
	}
	return products, nil
}