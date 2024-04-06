package controllers

import (
	"net/http"
	"product-microservice/db"
	"product-microservice/models"
	"product-microservice/service"
	"strconv"

	"github.com/labstack/echo/v4"
)

var errorMessage string = "Internal Server Error"
var notFoundMessage string = "Parameter Invalid"

func CreateProduct(c echo.Context) error {
	
    form, err := c.MultipartForm()
	if err != nil {
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
    cld, ctx := service.Credentials()
    cloudinaryURL, err := service.UploadImage(cld,ctx,file)
	if err != nil {
		return err
	}
    product.Image = cloudinaryURL
    if result := db.DB.Create(&product); result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": result.Error.Error()})
	}
    return c.JSON(http.StatusCreated, product)
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

func UpdateProduct(c echo.Context) error {
    id := c.Param("id")
    var product models.Product
    if result := db.DB.First(&product, id); result.Error != nil {
        return c.JSON(http.StatusNotFound, map[string]string{"error": notFoundMessage})
    }
    if err := c.Bind(&product); err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": errorMessage})
    }
    db.DB.Save(&product)
    return c.JSON(http.StatusOK, product)
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