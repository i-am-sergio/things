package services

import (
	"product-microservice/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestUpdateProductRating(t *testing.T) {
	t.Run("Sucess", func(t *testing.T) {
		product := models.Product{
			ID: 1,
			Rate: 0.0,
			Comments: []models.Comment{
				{Rating: 5},
				{Rating: 3},
			},
		}
		mockDB := new(MockDBClient)
		mockDB.On("FindPreloaded", "Comments", mock.AnythingOfType("*models.Product"), []interface{}{uint(1)}).Run(func(args mock.Arguments) {
			arg := args.Get(1).(*models.Product)
			*arg = product
		}).Return(nil)
		mockDB.On("Save", mock.AnythingOfType("*models.Product")).Run(func(args mock.Arguments) {
			arg := args.Get(0).(*models.Product)
			product.Rate = arg.Rate
		}).Return(nil)
		service := NewCommentService(mockDB)
		err := service.UpdateProductRating(1)
		require.NoError(t, err)
		assert.Equal(t, 4.0, product.Rate)
		mockDB.AssertExpectations(t)
	})
	t.Run("No comments", func(t *testing.T) {
		product := models.Product{
			ID: 1,
			Rate: 0.0,
		}
		mockDB := new(MockDBClient)
		mockDB.On("FindPreloaded", "Comments", mock.AnythingOfType("*models.Product"), []interface{}{uint(1)}).Run(func(args mock.Arguments) {
			arg := args.Get(1).(*models.Product)
			*arg = product
		}).Return(nil)
		mockDB.On("Save", mock.AnythingOfType("*models.Product")).Run(func(args mock.Arguments) {
			arg := args.Get(0).(*models.Product)
			product.Rate = arg.Rate
		}).Return(nil)
		service := NewCommentService(mockDB)
		err := service.UpdateProductRating(1)
		require.NoError(t, err)
		assert.Equal(t, 0.0, product.Rate)
		mockDB.AssertExpectations(t)
	})
	t.Run("Error FindPreloaded", func(t *testing.T) {
		mockDB := new(MockDBClient)
		mockDB.On("FindPreloaded", "Comments", mock.AnythingOfType("*models.Product"), []interface{}{uint(1)}).Return(assert.AnError)
		service := NewCommentService(mockDB)
		err := service.UpdateProductRating(1)
		require.Error(t, err)
		mockDB.AssertExpectations(t)
	})
	t.Run("Error Save", func(t *testing.T) {
		product := models.Product{
			ID: 1,
			Rate: 0.0,
			Comments: []models.Comment{
				{Rating: 5},
				{Rating: 3},
			},
		}
		mockDB := new(MockDBClient)
		mockDB.On("FindPreloaded", "Comments", mock.AnythingOfType("*models.Product"), []interface{}{uint(1)}).Run(func(args mock.Arguments) {
			arg := args.Get(1).(*models.Product)
			*arg = product
		}).Return(nil)
		mockDB.On("Save", mock.AnythingOfType("*models.Product")).Return(assert.AnError)
		service := NewCommentService(mockDB)
		err := service.UpdateProductRating(1)
		require.Error(t, err)
		mockDB.AssertExpectations(t)
	})
}

func TestCreateCommentsService(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		comment := models.Comment{
			ID: 1,
			ProductID: 1,
			Rating: 5,
		}
		product := models.Product{
			ID: 1,
			Rate: 0.0,
			Comments: []models.Comment{},
		}
		mockDB := new(MockDBClient)
		mockDB.On("First", mock.AnythingOfType("*models.Product"), []interface{}{uint(1)}).Run(func(args mock.Arguments) {
			arg := args.Get(0).(*models.Product)
			*arg = product
		}).Return(nil)
		mockDB.On("Create", mock.AnythingOfType("*models.Comment")).Return(nil)
		mockDB.On("FindPreloaded", "Comments", mock.AnythingOfType("*models.Product"), []interface{}{uint(1)}).Run(func(args mock.Arguments) {
			arg := args.Get(1).(*models.Product)
			*arg = product
			arg.Comments = append(arg.Comments, comment)
		}).Return(nil)
		mockDB.On("Save", mock.AnythingOfType("*models.Product")).Return(nil)
		service := NewCommentService(mockDB)
		err := service.CreateCommentService(comment)
		require.NoError(t, err)
		mockDB.AssertExpectations(t)
	})
	t.Run("Error First", func(t *testing.T) {
		comment := models.Comment{
			ID: 1,
			ProductID: 1,
			Rating: 5,
		}
		mockDB := new(MockDBClient)
		mockDB.On("First", mock.AnythingOfType("*models.Product"), []interface{}{uint(1)}).Return(assert.AnError)
		service := NewCommentService(mockDB)
		err := service.CreateCommentService(comment)
		require.Error(t, err)
		mockDB.AssertExpectations(t)
	})
	t.Run("Error Create", func(t *testing.T) {
		comment := models.Comment{
			ID: 1,
			ProductID: 1,
			Rating: 5,
		}
		product := models.Product{
			ID: 1,
			Rate: 0.0,
			Comments: []models.Comment{},
		}
		mockDB := new(MockDBClient)
		mockDB.On("First", mock.AnythingOfType("*models.Product"), []interface{}{uint(1)}).Run(func(args mock.Arguments) {
			arg := args.Get(0).(*models.Product)
			*arg = product
		}).Return(nil)
		mockDB.On("Create", mock.AnythingOfType("*models.Comment")).Return(assert.AnError)
		service := NewCommentService(mockDB)
		err := service.CreateCommentService(comment)
		require.Error(t, err)
		mockDB.AssertExpectations(t)
	})
}

func TestGetCommentsService(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		comments := []models.Comment{
			{ID: 1},
			{ID: 2},
		}
		mockDB := new(MockDBClient)
		mockDB.On("Find", mock.AnythingOfType("*[]models.Comment"), []interface{}(nil)).Run(func(args mock.Arguments) {
			arg := args.Get(0).(*[]models.Comment)
			*arg = comments
		}).Return(nil)
		service := NewCommentService(mockDB)
		result, err := service.GetCommentsService()
		require.NoError(t, err)
		assert.Equal(t, comments, result)
		mockDB.AssertExpectations(t)
	})
	t.Run("Error Find", func(t *testing.T) {
		mockDB := new(MockDBClient)
		mockDB.On("Find", mock.AnythingOfType("*[]models.Comment"), []interface{}(nil)).Return(assert.AnError)
		service := NewCommentService(mockDB)
		result, err := service.GetCommentsService()
		require.Error(t, err)
		assert.Nil(t, result)
		mockDB.AssertExpectations(t)
	})
}

func TestGetCommentByIDService(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		comment := models.Comment{ID: 1}
		mockDB := new(MockDBClient)
		mockDB.On("First", mock.AnythingOfType("*models.Comment"), []interface{}{uint(1)}).Run(func(args mock.Arguments) {
			arg := args.Get(0).(*models.Comment)
			*arg = comment
		}).Return(nil)
		service := NewCommentService(mockDB)
		result, err := service.GetCommentByIDService(1)
		require.NoError(t, err)
		assert.Equal(t, comment, result)
		mockDB.AssertExpectations(t)
	})
	t.Run("Error First", func(t *testing.T) {
		mockDB := new(MockDBClient)
		mockDB.On("First", mock.AnythingOfType("*models.Comment"), []interface{}{uint(1)}).Return(assert.AnError)
		service := NewCommentService(mockDB)
		result, err := service.GetCommentByIDService(1)
		require.Error(t, err)
		assert.Equal(t, models.Comment{}, result)
		mockDB.AssertExpectations(t)
	})
}

func TestGetCommentsByProductIDService(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		comments := []models.Comment{
			{ID: 1, ProductID: 1},
			{ID: 2, ProductID: 1},
		}
		mockDB := new(MockDBClient)
		mockDB.On("FindWithCondition", mock.AnythingOfType("*[]models.Comment"), "product_id = ?", []interface{}{uint(1)}).Run(func(args mock.Arguments) {
			arg := args.Get(0).(*[]models.Comment)
			*arg = comments
		}).Return(nil)
		service := NewCommentService(mockDB)
		result, err := service.GetCommentsByProductIDService(1)
		require.NoError(t, err)
		assert.Equal(t, comments, result)
		mockDB.AssertExpectations(t)
	})
	t.Run("Error FindWithCondition", func(t *testing.T) {
		mockDB := new(MockDBClient)
		mockDB.On("FindWithCondition", mock.AnythingOfType("*[]models.Comment"), "product_id = ?", []interface{}{uint(1)}).Return(assert.AnError)
		service := NewCommentService(mockDB)
		result, err := service.GetCommentsByProductIDService(1)
		require.Error(t, err)
		assert.Nil(t, result)
		mockDB.AssertExpectations(t)
	})
}

func TestUpdateCommentService(t *testing.T) {
    t.Run("Success", func(t *testing.T) {
        comment := models.Comment{
            ID:        1,
            ProductID: 1,
            Rating:    5,
        }
        product := models.Product{
            ID: 1,
            Rate: 0.0,
            Comments: []models.Comment{},
        }
        mockDB := new(MockDBClient)
        mockDB.On("Save", mock.AnythingOfType("*models.Comment")).Return(nil)
        mockDB.On("FindPreloaded", "Comments", mock.AnythingOfType("*models.Product"), []interface{}{uint(1)}).Run(func(args mock.Arguments) {
            arg := args.Get(1).(*models.Product)
            *arg = product
            arg.Comments = append(arg.Comments, comment)
        }).Return(nil)
        mockDB.On("Save", mock.AnythingOfType("*models.Product")).Return(nil)
        service := NewCommentService(mockDB)
        err := service.UpdateCommentService(comment, 1)
        require.NoError(t, err)
        mockDB.AssertExpectations(t)
    })
	t.Run("Error Save", func(t *testing.T) {
		comment := models.Comment{
			ID: 1,
			ProductID: 1,
			Rating: 5,
		}
		mockDB := new(MockDBClient)
		mockDB.On("Save", mock.AnythingOfType("*models.Comment")).Return(assert.AnError)
		service := NewCommentService(mockDB)
		err := service.UpdateCommentService(comment, 1)
		require.Error(t, err)
		mockDB.AssertExpectations(t)
	})
	t.Run("Error FindPreloaded", func(t *testing.T) {
		comment := models.Comment{
			ID: 1,
			ProductID: 1,
			Rating: 5,
		}
		mockDB := new(MockDBClient)
		mockDB.On("Save", mock.AnythingOfType("*models.Comment")).Return(nil)
		mockDB.On("FindPreloaded", "Comments", mock.AnythingOfType("*models.Product"), []interface{}{uint(1)}).Return(assert.AnError)
		service := NewCommentService(mockDB)
		err := service.UpdateCommentService(comment, 1)
		require.Error(t, err)
		mockDB.AssertExpectations(t)
	})
	t.Run("Success with different product ID", func(t *testing.T) {
		comment := models.Comment{
			ID: 1,
			ProductID: 2,
			Rating: 5,
		}
		productForComment := models.Product{
			ID: 2,
			Rate: 0.0,
			Comments: []models.Comment{},
		}
		productForID := models.Product{
            ID: 1,
            Rate: 0.0,
            Comments: []models.Comment{},
        }
		mockDB := new(MockDBClient)
		mockDB.On("Save", mock.AnythingOfType("*models.Comment")).Return(nil)
		mockDB.On("FindPreloaded", "Comments", mock.AnythingOfType("*models.Product"), []interface{}{uint(2)}).Run(func(args mock.Arguments) {
			arg := args.Get(1).(*models.Product)
			*arg = productForComment
		}).Return(nil)
		mockDB.On("Save", mock.AnythingOfType("*models.Product")).Return(nil).Once()
		mockDB.On("FindPreloaded", "Comments", mock.AnythingOfType("*models.Product"), []interface{}{uint(1)}).Run(func(args mock.Arguments) {
            arg := args.Get(1).(*models.Product)
            *arg = productForID
        }).Return(nil)
		mockDB.On("Save", mock.AnythingOfType("*models.Product")).Return(nil).Once()
		service := NewCommentService(mockDB)
		err := service.UpdateCommentService(comment, 1)
		require.NoError(t, err)
		mockDB.AssertExpectations(t)
	})
}