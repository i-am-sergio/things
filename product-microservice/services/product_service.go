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
	if err := db.Client.Create(&product); err != nil {
		return product, err
	}
	return product, nil
}

func UpdateProductService(productID uint, form *multipart.Form, image *multipart.FileHeader) (models.Product, error){
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
		cloudinaryURL, err := db.UploadImage(image)
		if err != nil {
			return product, err
		}
		product.Image = cloudinaryURL
	}
	if err := db.Client.Save(&product); err != nil {
		return product, err
	}
	return product, nil
}

func GetProductsService() ([]models.Product, error) {
	var products []models.Product
	if err := db.Client.Find(&products); err != nil {
		return nil, err
	}
	return products, nil
}

func GetProductByIDService(id uint) (models.Product, error) {
	var product models.Product
	if err := db.Client.First(&product, id); err != nil {
		return product, err
	}
	return product, nil
}

func GetProductsByCategoryService(category string) ([]models.Product, error) {
	var products []models.Product
	if category != "" {
		err := db.Client.FindWithCondition(&products, "category = ?", category)
		if err != nil {
			return nil, err
		}
	} else {
		err := db.Client.Find(&products)
		if err != nil {
			return nil, err
		}
	}
	return products, nil
}

func SearchProductsService(searchTerm string) ([]models.Product, error) {
	var products []models.Product
	err := db.Client.FindWithCondition(&products, "name LIKE ? OR description LIKE ?", "%"+searchTerm+"%", "%"+searchTerm+"%")
	if err != nil {
		return nil, err
	}
	return products, nil
}

func DeleteProductService(productID uint) error {
	if err := db.Client.DeleteWithCondition(&models.Comment{}, "product_id = ?", productID); err != nil {
		return err
	}
	if err := db.Client.DeleteByID(&models.Product{}, productID); err != nil {
		return err
	}
	return nil
}

func PremiumService(productID uint) (models.Product, error) {
	var product models.Product
	if err := db.Client.First(&product, productID); err != nil {
		return product, err
	}
	product.Status = !product.Status
	if err := db.Client.Save(&product); err != nil {
		return product, err
	}
	return product, nil
}

func GetProductsPremiumService() ([]models.Product, error) {
	var products []models.Product
	err := db.Client.FindWithCondition(&products, "status = ?", true)
	if err != nil {
		return nil, err
	}
	return products, nil
}