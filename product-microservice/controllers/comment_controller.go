package controllers

import (
	"net/http"
	"product-microservice/db"
	"product-microservice/models"
	"product-microservice/services"
	"strconv"

	"github.com/labstack/echo/v4"
)

const invalidCommentIDError = "Invalid comment ID"

func CreateComment(c echo.Context) error {
    var comment models.Comment
    if err := c.Bind(&comment); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
    }
    if err := services.CreateCommentService(comment); err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }
    return c.JSON(http.StatusCreated, comment)
}

func GetComments(c echo.Context) error {
    comments, err := services.GetCommentsService()
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }
    return c.JSON(http.StatusOK, comments)
}

func GetCommentByID(c echo.Context) error {
    ID, err := strconv.Atoi(c.Param("id")) 
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": invalidCommentIDError})
    }
    getComment, err := services.GetCommentByIDService(uint(ID))
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }
    return c.JSON(http.StatusOK, getComment)
}

func GetCommentsByProductID(c echo.Context) error {
    ID, err := strconv.Atoi(c.Param("id")) 
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": invalidCommentIDError})
    }
    comments, err := services.GetCommentsByProductIDService(uint(ID))
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }
    return c.JSON(http.StatusOK, comments)
}

func UpdateComment(c echo.Context) error {
    commentID, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": invalidCommentIDError})
    }
    var comment models.Comment
    if result := db.DB.First(&comment, commentID); result.Error != nil {
        return c.JSON(http.StatusNotFound, map[string]string{"error": "Comment not found"})
    }
    befProduct := comment.ProductID
    if err := c.Bind(&comment); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
    }
    if err := services.UpdateCommentService(comment, uint(befProduct)); err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }
    return c.JSON(http.StatusOK, comment)
}

func DeleteComment(c echo.Context) error {
    commentID, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": invalidCommentIDError})
    }
    if err := services.DeleteCommentService(uint(commentID)); err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }
    return c.JSON(http.StatusOK, map[string]string{"message": "Comment deleted"})
}

