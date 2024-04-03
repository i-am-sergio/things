package controllers

import (
	"fmt"
	"math"
	"net/http"
	"product-microservice/db"
	"product-microservice/models"
	"strconv"

	"github.com/labstack/echo/v4"
)


func CreateComment(c echo.Context) error {
    var comment models.Comment
    if err := c.Bind(&comment); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
    }
    if result := db.DB.Create(&comment); result.Error != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": result.Error.Error()})
    }
    var product models.Product
    if result := db.DB.First(&product, comment.ProductID); result.Error != nil {
        return c.JSON(http.StatusNotFound, map[string]string{"error": "Product not found"})
    }
    product.Comments = append(product.Comments, comment)
    if result := db.DB.Save(&product); result.Error != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": result.Error.Error()})
    }
    return c.JSON(http.StatusCreated, comment)
}

func GetCommentsByProductID(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))
    var comments []models.Comment
    if result := db.DB.Where("product_id = ?", id).Find(&comments); result.Error != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": result.Error.Error()})
    }
    return c.JSON(http.StatusOK, comments)
}

func DeleteComment(c echo.Context) error {
    commentID, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid comment ID"})
    }
    var comment models.Comment
    if result := db.DB.First(&comment, commentID); result.Error != nil {
        return c.JSON(http.StatusNotFound, map[string]string{"error": "Comment not found"})
    }
    var product models.Product
    if result := db.DB.First(&product, comment.ProductID); result.Error != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": result.Error.Error()})
    }
    for i, c := range product.Comments {
        if c.ID == comment.ID {
            product.Comments = append(product.Comments[:i], product.Comments[i+1:]...)
            break
        }
    }
    if result := db.DB.Save(&product); result.Error != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": result.Error.Error()})
    }
    if result := db.DB.Delete(&comment); result.Error != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": result.Error.Error()})
    }
    return c.JSON(http.StatusOK, map[string]string{"message": "Comment deleted successfully"})
}

func UpdateComment(c echo.Context) error {
    commentID, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid comment ID"})
    }
    var comment models.Comment
    if result := db.DB.First(&comment, commentID); result.Error != nil {
        return c.JSON(http.StatusNotFound, map[string]string{"error": "Comment not found"})
    }
    if err := c.Bind(&comment); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
    }
    if result := db.DB.Save(&comment); result.Error != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": result.Error.Error()})
    }
    return c.JSON(http.StatusOK, comment)
}

func UpdateProductRating(c echo.Context) error {
    productID, err := strconv.Atoi(c.Param("id"))
    fmt.Println(productID)
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid product ID"})
    }
    var product models.Product
    if result := db.DB.First(&product, productID); result.Error != nil {
        return c.JSON(http.StatusNotFound, map[string]string{"error": "Product not found"})
    }
    var comments []models.Comment
    if result := db.DB.Where("product_id = ?", productID).Find(&comments); result.Error != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": result.Error.Error()})
    }
    totalRatings := 0.0
    for _, comment := range comments {
        totalRatings += comment.Rating
    }
    if len(comments) > 0 {
        averageRating := totalRatings / float64(len(comments))
        averageRating = math.Round(averageRating*100) / 100
        product.Ratings = averageRating
    } else {
        product.Ratings = 0.0
    }
    if result := db.DB.Save(&product); result.Error != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": result.Error.Error()})
    }
    return c.JSON(http.StatusOK, product)
}