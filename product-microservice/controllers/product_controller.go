package controllers

import (
	"mime/multipart"
	"net/http"
	"product-microservice/services"
	"strconv"

	"github.com/labstack/echo/v4"
)
type ProductController interface {
	CreateProduct(c echo.Context) error
	UpdateProduct(c echo.Context) error
	GetProducts(c echo.Context) error
	GetProductsById(c echo.Context) error
	GetProductsByCategory(c echo.Context) error
	SearchProducts(c echo.Context) error
	DeleteProduct(c echo.Context) error
	Premium(c echo.Context) error
	GetProductsPremium(c echo.Context) error
}
type ProductControllerImpl struct {
	ProductService services.ProductService
}
func NewProductController(productService services.ProductService) *ProductControllerImpl {
	return &ProductControllerImpl{ProductService: productService}
}

func (cx *ProductControllerImpl) CreateProduct(c echo.Context) error {
	form, err := c.MultipartForm()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	file, err := c.FormFile("image")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	product,err := cx.ProductService.CreateProductService(form, file)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, product)
}

func (cx *ProductControllerImpl) UpdateProduct(c echo.Context)  error {
	id, err := strconv.Atoi(c.Param("id")) 
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": invalidCommentIDError})
	}
	form, err := c.MultipartForm()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	var image *multipart.FileHeader
	file, err := c.FormFile("image")
	if err == nil {
		image = file
	}
	product, err := cx.ProductService.UpdateProductService(uint(id), form, image)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, product)
}

func (cx *ProductControllerImpl) GetProducts(c echo.Context) error {
	products, err := cx.ProductService.GetProductsService()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, products)
}

func (cx *ProductControllerImpl) GetProductsById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id")) 
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": invalidCommentIDError})
    }
	product, err := cx.ProductService.GetProductByIDService(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, product)
}

func (cx *ProductControllerImpl) GetProductsByCategory(c echo.Context) error {
    category := c.QueryParam("category")
	products, err := cx.ProductService.GetProductsByCategoryService(category)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, products)
}

func (cx *ProductControllerImpl) SearchProducts(c echo.Context) error {
    searchTerm := c.QueryParam("q")
	products, err := cx.ProductService.SearchProductsService(searchTerm)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, products)
}

func (cx *ProductControllerImpl) DeleteProduct(c echo.Context) error {
	productID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": invalidCommentIDError})
	}
	if err := cx.ProductService.DeleteProductService(uint(productID)); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Product deleted successfully"})
}

func (cx *ProductControllerImpl) Premium(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": invalidCommentIDError})
	}
	product, err := cx.ProductService.PremiumService(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, product)
}

func (cx *ProductControllerImpl) GetProductsPremium(c echo.Context) error {
	products, err := cx.ProductService.GetProductsPremiumService()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, products)
}