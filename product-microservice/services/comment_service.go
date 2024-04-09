package services

import (
	"product-microservice/db"
	"product-microservice/models"
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

func CreateCommentService(comment models.Comment) error {
	var product models.Product
	if err := db.DB.First(&product, comment.ProductID).Error; err != nil {
		return err
	}
	if result := db.DB.Create(&comment); result.Error != nil {
		return result.Error
	}
	if err := UpdateProductRating(comment.ProductID); err != nil {
		return err
	}
	return nil
}

func GetCommentsService() ([]models.Comment, error) {
	var comments []models.Comment
	if result := db.DB.Find(&comments); result.Error != nil {
		return nil, result.Error
	}
	return comments, nil
}

func GetCommentByIDService(id uint) (models.Comment, error) {
	var comment models.Comment
	if result := db.DB.First(&comment, id); result.Error != nil {
		return comment, result.Error
	}
	return comment, nil
}

func GetCommentsByProductIDService(id uint) ([]models.Comment, error) {
	var comments []models.Comment
	if result := db.DB.Where("product_id = ?", id).Find(&comments); result.Error != nil {
		return nil, result.Error
	}
	return comments, nil
}

func UpdateCommentService(comment models.Comment, id uint) error {
	if result := db.DB.Save(&comment); result.Error != nil {
		return result.Error
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
	if result := db.DB.First(&comment, id); result.Error != nil {
		return result.Error
	}
	if result := db.DB.Delete(&comment); result.Error != nil {
		return result.Error
	}
	if err := UpdateProductRating(comment.ProductID); err != nil {
		return err
	}
	return nil
}