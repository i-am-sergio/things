package controllers

import (
	"mime/multipart"
	"net/http"
	"product-microservice/db"
	"product-microservice/services"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CreateProduct(cloudinary *db.Cloudinary) echo.HandlerFunc {
	return func(c echo.Context) error {
		form, err := c.MultipartForm()
		if err != nil {
			return err
		}
		file, err := c.FormFile("image")
		if err != nil {
			return err
		}
		product,err := services.CreateProductService(cloudinary, form, file)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		return c.JSON(http.StatusCreated, product)
	}
}

func UpdateProduct(cloudinary *db.Cloudinary) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id")) 
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": invalidCommentIDError})
		}
		form, err := c.MultipartForm()
		if err != nil {
			return err
		}
		var image *multipart.FileHeader
		file, err := c.FormFile("image")
		if err == nil {
			image = file
		}
		product, err := services.UpdateProductService(cloudinary, uint(id), form, image)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		return c.JSON(http.StatusOK, product)
	}
}

func GetProducts(c echo.Context) error {
	products, err := services.GetProductsService()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, products)
}

func GetProductsById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id")) 
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": invalidCommentIDError})
    }
	product, err := services.GetProductByIDService(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, product)
}

func GetProductsByCategory(c echo.Context) error {
    category := c.QueryParam("category")
	products, err := services.GetProductsByCategoryService(category)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, products)
}

func SearchProducts(c echo.Context) error {
    searchTerm := c.QueryParam("q")
	products, err := services.SearchProductsService(searchTerm)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, products)
}

func DeleteProduct(c echo.Context) error {
	productID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": invalidCommentIDError})
	}
	if err := services.DeleteProductService(uint(productID)); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Product deleted successfully"})
}

func Premium(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": invalidCommentIDError})
	}
	product, err := services.PremiumService(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, product)
}

func GetProductsPremium(c echo.Context) error {
	products, err := services.GetProductsPremiumService()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, products)
}