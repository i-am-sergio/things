package controllers

import (
	"fmt"
	"net/http"
	"product-microservice/db"
	"product-microservice/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

func UpdateProductRating(productID uint) error {
    var product models.Product
    if err := db.DB.Preload("Comments").First(&product, productID).Error; err != nil {
        return err
    }
    if len(product.Comments) == 0 {
        product.Rate = 0
    } else {
        var totalRating float64
        for _, comment := range product.Comments {
            totalRating += comment.Rating
        }
        product.Rate = totalRating / float64(len(product.Comments))
    }
    if err := db.DB.Save(&product).Error; err != nil {
        return err
    }
    return nil
}

func CreateComment(c echo.Context) error {
    var comment models.Comment
    if err := c.Bind(&comment); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
    }
    fmt.Printf("Comment Received: %+v\n", comment)
    var product models.Product
    if err := db.DB.First(&product, comment.ProductID).Error; err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Product not found"})
    }
    if result := db.DB.Create(&comment); result.Error != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": result.Error.Error()})
    }
    if err := UpdateProductRating(comment.ProductID); err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }
    return c.JSON(http.StatusCreated, comment)
}

func GetComments(c echo.Context) error {
    var comments []models.Comment
    if result := db.DB.Find(&comments); result.Error != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": result.Error.Error()})
    }
    return c.JSON(http.StatusOK, comments)
}

func GetCommentByID(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))
    var comment models.Comment
    if result := db.DB.First(&comment, id); result.Error != nil {
        return c.JSON(http.StatusNotFound, map[string]string{"error": result.Error.Error()})
    }
    return c.JSON(http.StatusOK, comment)
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
    if result := db.DB.Delete(&comment); result.Error != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": result.Error.Error()})
    }
    if err := UpdateProductRating(comment.ProductID); err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
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
    befProduct := comment.ProductID
    if err := c.Bind(&comment); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
    }
    if result := db.DB.Save(&comment); result.Error != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": result.Error.Error()})
    }
    if(befProduct != comment.ProductID) {
        if err := UpdateProductRating(befProduct); err != nil {
            return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
        }
    }
    if err := UpdateProductRating(comment.ProductID); err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }
    return c.JSON(http.StatusOK, comment)
}