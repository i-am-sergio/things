package services

import (
	"fmt"
	"mime/multipart"
	"product-microservice/db"
	"product-microservice/models"
	"strconv"
)

func validateRequiredFields(form *multipart.Form) error {
    requiredFields := []string{"UserID", "Price", "Name", "Description", "Category", "Ubication"}
    for _, field := range requiredFields {
        if len(form.Value[field]) == 0 {
            return fmt.Errorf("missing required field: %s", field)
        }
    }
    return nil
}

func CreateProductService(form *multipart.Form, image *multipart.FileHeader) (models.Product, error){
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
	cloudinaryURL, err := db.UploadImage(image)
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
	if result := db.DB.Create(&product); result.Error != nil {
		return product, result.Error
	}
	return product, nil
}

func UpdateProductService(productID uint, form *multipart.Form, image *multipart.FileHeader) (models.Product, error){
	var product models.Product
    if result := db.DB.First(&product, productID); result.Error != nil {
        return product, result.Error
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
		cloudinaryURL, err := db.UploadImage(image)
		if err != nil {
			return product, err
		}
		product.Image = cloudinaryURL
	}
	if result := db.DB.Save(&product); result.Error != nil {
		return product, result.Error
	}
	return product, nil
}

func GetProductsService() ([]models.Product, error) {
	var products []models.Product
	if result := db.DB.Find(&products); result.Error != nil {
		return nil, result.Error
	}
	return products, nil
}

func GetProductByIDService(id uint) (models.Product, error) {
	var product models.Product
	if result := db.DB.First(&product, id); result.Error != nil {
		return product, result.Error
	}
	return product, nil
}

func GetProductsByCategoryService(category string) ([]models.Product, error) {
	var products []models.Product
	if category != "" {
		result := db.DB.Where("category = ?", category).Find(&products)
		if result.Error != nil {
			return nil, result.Error
		}
	} else {
		result := db.DB.Find(&products)
		if result.Error != nil {
			return nil, result.Error
		}
	}
	return products, nil
}

func SearchProductsService(searchTerm string) ([]models.Product, error) {
	var products []models.Product
	result := db.DB.Where("name LIKE ? OR description LIKE ?", "%"+searchTerm+"%", "%"+searchTerm+"%").Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}
	return products, nil
}

func DeleteProductService(productID uint) error {
	if err := db.DB.Where("product_id = ?", productID).Delete(&models.Comment{}).Error; err != nil {
		return err
	}
	if err := db.DB.Delete(&models.Product{}, productID).Error; err != nil {
		return err
	}
	return nil
}

func PremiumService(productID uint) (models.Product, error) {
	var product models.Product
	if result := db.DB.First(&product, productID); result.Error != nil {
		return product, result.Error
	}
	product.Status = !product.Status
	if result := db.DB.Save(&product); result.Error != nil {
		return product, result.Error
	}
	return product, nil
}

func GetProductsPremiumService() ([]models.Product, error) {
	var products []models.Product
	result := db.DB.Where("status = ?", true).Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}
	return products, nil
}