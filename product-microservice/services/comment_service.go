package services

import (
	"product-microservice/models"
	"product-microservice/repository"
)

type CommentService interface {
	UpdateProductRating(productID uint) error
	CreateCommentService(comment models.Comment) error
	GetCommentsService() ([]models.Comment, error)
	GetCommentByIDService(id uint) (models.Comment, error)
	GetCommentsByProductIDService(id uint) ([]models.Comment, error)
	UpdateCommentService(comment models.Comment, id uint) error
	DeleteCommentService(id uint) error
}
type CommentServiceImpl struct {
	dbClient repository.DBInterface
}
func NewCommentService(client repository.DBInterface) *CommentServiceImpl {
	return &CommentServiceImpl{dbClient: client}
}

func (c *CommentServiceImpl) UpdateProductRating(productID uint) error {
    var product models.Product
    if err := c.dbClient.FindPreloaded("Comments",&product,productID); err != nil {
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
    if err := c.dbClient.Save(&product); err != nil {
        return err
    }
    return nil
}

func (c *CommentServiceImpl) CreateCommentService(comment models.Comment) error {
	var product models.Product
	if err := c.dbClient.First(&product, comment.ProductID); err != nil {
		return err
	}
	if err := c.dbClient.Create(&comment); err != nil {
		return err
	}
	if err := c.UpdateProductRating(comment.ProductID); err != nil {
		return err
	}
	return nil
}

func (c *CommentServiceImpl) GetCommentsService() ([]models.Comment, error) {
	var comments []models.Comment
	if err := c.dbClient.Find(&comments); err != nil {
		return nil, err
	}
	return comments, nil
}

func (c *CommentServiceImpl) GetCommentByIDService(id uint) (models.Comment, error) {
	var comment models.Comment
	if err := c.dbClient.First(&comment, id); err != nil {
		return comment, err
	}
	return comment, nil
}

func (c *CommentServiceImpl) GetCommentsByProductIDService(id uint) ([]models.Comment, error) {
	var comments []models.Comment
	if err := c.dbClient.FindWithCondition(&comments,"product_id = ?", id); err != nil {
		return nil, err
	}
	return comments, nil
}

func (c *CommentServiceImpl) UpdateCommentService(comment models.Comment, id uint) error {
	if err := c.dbClient.Save(&comment); err != nil {
		return err
	}
	if(id != uint(comment.ProductID)){
		if err := c.UpdateProductRating(uint(id)); err != nil {
			return err
		}
	}
	if err := c.UpdateProductRating(comment.ProductID); err != nil {
		return err
	}
	return nil
}

func (c *CommentServiceImpl) DeleteCommentService(id uint) error {
	var comment models.Comment
	if err := c.dbClient.First(&comment, id); err != nil {
		return err
	}
	if err := c.dbClient.Delete(&comment); err != nil {
		return err
	}
	if err := c.UpdateProductRating(comment.ProductID); err != nil {
		return err
	}
	return nil
}