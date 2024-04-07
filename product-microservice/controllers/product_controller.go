package controllers

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"product-microservice/db"
	"product-microservice/models"
	"product-microservice/service"
	"strconv"

	"github.com/labstack/echo/v4"
)

var errorMessage string = "Internal Server Error"
var notFoundMessage string = "Parameter Invalid"

func validateRequiredFields(form *multipart.Form) error {
    requiredFields := []string{"UserID", "Price", "Name", "Description", "Category", "Ubication"}
    for _, field := range requiredFields {
        if len(form.Value[field]) == 0 {
            return fmt.Errorf("missing required field: %s", field)
        }
    }
    return nil
}

func CreateProduct(c echo.Context) error {
    form, err := c.MultipartForm()
	if err != nil {
		return err
	}
	if err := validateRequiredFields(form); err != nil {
        return err
    }
    userID, err := strconv.ParseUint(form.Value["UserID"][0], 10, 32)
	if err != nil {
		return err
	}
	price, err := strconv.ParseFloat(form.Value["Price"][0], 64)
	if err != nil {
		return err
	}
    product := models.Product{
		UserID:      uint(userID),
		State:       true,
		Status:      false,
		Name:        form.Value["Name"][0],
		Description: form.Value["Description"][0],
		Category:    form.Value["Category"][0],
		Price:       float64(price),
		Rate:        0.0,
		Ubication:   form.Value["Ubication"][0],
	}
	file, err := c.FormFile("image")
	if err != nil {
		return err
	}
    cloudinaryURL, err := service.UploadImage(file)
	if err != nil {
		return err
	}
    product.Image = cloudinaryURL
    if result := db.DB.Create(&product); result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": result.Error.Error()})
	}
    return c.JSON(http.StatusCreated, product)
}

func UpdateProduct(c echo.Context) error {
    id := c.Param("id")
    var product models.Product
    if result := db.DB.First(&product, id); result.Error != nil {
        return c.JSON(http.StatusNotFound, map[string]string{"error": notFoundMessage})
    }
	form, err := c.MultipartForm()
    if err != nil {
        return err
    }
	if err := validateRequiredFields(form); err != nil {
        return err
    }
    userID, err := strconv.ParseUint(form.Value["UserID"][0], 10, 32)
	if err != nil {
		return err
	}
	price, err := strconv.ParseFloat(form.Value["Price"][0], 64)
	if err != nil {
		return err
	}
	product.UserID = uint(userID)
	product.Name = form.Value["Name"][0]
	product.Description = form.Value["Description"][0]
	product.Category = form.Value["Category"][0]
	product.Price = float64(price)
	product.Ubication = form.Value["Ubication"][0]
	if file, err := c.FormFile("image"); err == nil {
        cloudinaryURL, err := service.UploadImage(file)
        if err != nil {
            return err
        }
        product.Image = cloudinaryURL
    }
	if result := db.DB.Save(&product); result.Error != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": result.Error.Error()})
    }
    return c.JSON(http.StatusOK, product)
}

func GetProducts(c echo.Context) error {
	var products []models.Product
	result := db.DB.Find(&products)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": errorMessage})
	}
	return c.JSON(http.StatusOK, products)
}

func GetProductsById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	product := new(models.Product)
	result := db.DB.First(&product, id)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": notFoundMessage})
	}
	return c.JSON(http.StatusOK, product)
}

func GetProductsByCategory(c echo.Context) error {
    var products []models.Product
    category := c.QueryParam("category")
    if category != "" {
        result := db.DB.Where("category = ?", category).Find(&products)
        if result.Error != nil {
            return c.JSON(http.StatusInternalServerError, map[string]string{"error": notFoundMessage})
        }
    } else {
        result := db.DB.Find(&products)
        if result.Error != nil {
            return c.JSON(http.StatusInternalServerError, map[string]string{"error": errorMessage})
        }
    }
    return c.JSON(http.StatusOK, products)
}

func SearchProducts(c echo.Context) error {
    searchTerm := c.QueryParam("q")
    var products []models.Product
    result := db.DB.Where("name LIKE ? OR description LIKE ?", "%"+searchTerm+"%", "%"+searchTerm+"%").Find(&products)
    if result.Error != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": errorMessage})
    }
    return c.JSON(http.StatusOK, products)
}



func DeleteProduct(c echo.Context) error {
	productID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": notFoundMessage})
	}
	if err := db.DB.Where("product_id = ?", productID).Delete(&models.Comment{}).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": errorMessage})
	}
	if err := db.DB.Delete(&models.Product{}, productID).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": errorMessage})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Product and its comments deleted successfully"})
}

func Premium(c echo.Context) error {
	id := c.Param("id")
    var product models.Product
    if result := db.DB.First(&product, id); result.Error != nil {
        return c.JSON(http.StatusNotFound, map[string]string{"error": notFoundMessage})
    }
	product.Status = !product.Status
	if result := db.DB.Save(&product); result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": result.Error.Error()})
	}
	return c.JSON(http.StatusOK, product)
}

func GetProductsPremium(c echo.Context) error {
	var products []models.Product
	result := db.DB.Where("status = ?", true).Find(&products)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": errorMessage})
	}
	return c.JSON(http.StatusOK, products)
}