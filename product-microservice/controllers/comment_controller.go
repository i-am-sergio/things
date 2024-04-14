package controllers

import (
	"net/http"
	"product-microservice/models"
	"product-microservice/services"
	"strconv"

	"github.com/labstack/echo/v4"
)

const invalidCommentIDError = "Invalid comment ID"

type CommentController interface {
    CreateComment(c echo.Context) error
    GetComments(c echo.Context) error
    GetCommentByID(c echo.Context) error
    GetCommentsByProductID(c echo.Context) error
    UpdateComment(c echo.Context) error
    DeleteComment(c echo.Context) error
}
type CommentControllerImpl struct {
    CommentService services.CommentService
}
func NewCommentController(commentService services.CommentService) *CommentControllerImpl {
    return &CommentControllerImpl{CommentService: commentService}
}

func (cx *CommentControllerImpl) CreateComment(c echo.Context) error {
    var comment models.Comment
    if err := c.Bind(&comment); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
    }
    if err := cx.CommentService.CreateCommentService(comment); err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }
    return c.JSON(http.StatusCreated, comment)
}

func (cx *CommentControllerImpl) GetComments(c echo.Context) error {
    comments, err := cx.CommentService.GetCommentsService()
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }
    return c.JSON(http.StatusOK, comments)
}

func (cx *CommentControllerImpl) GetCommentByID(c echo.Context) error {
    ID, err := strconv.Atoi(c.Param("id")) 
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": invalidCommentIDError})
    }
    getComment, err := cx.CommentService.GetCommentByIDService(uint(ID))
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }
    return c.JSON(http.StatusOK, getComment)
}

func (cx *CommentControllerImpl) GetCommentsByProductID(c echo.Context) error {
    ID, err := strconv.Atoi(c.Param("id")) 
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": invalidCommentIDError})
    }
    comments, err := cx.CommentService.GetCommentsByProductIDService(uint(ID))
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }
    return c.JSON(http.StatusOK, comments)
}

func (cx *CommentControllerImpl) UpdateComment(c echo.Context) error {
    commentID, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": invalidCommentIDError})
    }
    comment, err := cx.CommentService.GetCommentByIDService(uint(commentID))
    if err != nil {
        return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
    }
    befProduct := comment.ProductID
    if err := c.Bind(&comment); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
    }
    if err := cx.CommentService.UpdateCommentService(comment, uint(befProduct)); err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }
    return c.JSON(http.StatusOK, comment)
}

func (cx *CommentControllerImpl) DeleteComment(c echo.Context) error {
    commentID, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": invalidCommentIDError})
    }
    if err := cx.CommentService.DeleteCommentService(uint(commentID)); err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }
    return c.JSON(http.StatusOK, map[string]string{"message": "Comment deleted"})
}