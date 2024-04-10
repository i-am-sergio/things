package services

import (
	"product-microservice/db"
	"product-microservice/models"
)

func UpdateProductRating(productID uint) error {
    var product models.Product
    if err := db.Client.FindPreloaded("Comments",&product,productID); err != nil {
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
    if err := db.Client.Save(&product); err != nil {
        return err
    }
    return nil
}

func CreateCommentService(comment models.Comment) error {
	var product models.Product
	if err := db.Client.First(&product, comment.ProductID); err != nil {
		return err
	}
	if err := db.Client.Create(&comment); err != nil {
		return err
	}
	if err := UpdateProductRating(comment.ProductID); err != nil {
		return err
	}
	return nil
}

func GetCommentsService() ([]models.Comment, error) {
	var comments []models.Comment
	if err := db.Client.Find(&comments); err != nil {
		return nil, err
	}
	return comments, nil
}

func GetCommentByIDService(id uint) (models.Comment, error) {
	var comment models.Comment
	if err := db.Client.First(&comment, id); err != nil {
		return comment, err
	}
	return comment, nil
}

func GetCommentsByProductIDService(id uint) ([]models.Comment, error) {
	var comments []models.Comment
	if err := db.Client.FindWithCondition(&comments,"product_id = ?", id); err != nil {
		return nil, err
	}
	return comments, nil
}

func UpdateCommentService(comment models.Comment, id uint) error {
	if err := db.Client.Save(&comment); err != nil {
		return err
	}
	if(id != uint(comment.ProductID)){
		if err := UpdateProductRating(uint(id)); err != nil {
			return err
		}
	}
	if err := UpdateProductRating(comment.ProductID); err != nil {
		return err
	}
	return nil
}

func DeleteCommentService(id uint) error {
	var comment models.Comment
	if err := db.Client.First(&comment, id); err != nil {
		return err
	}
	if err := db.Client.Delete(&comment); err != nil {
		return err
	}
	if err := UpdateProductRating(comment.ProductID); err != nil {
		return err
	}
	return nil
}