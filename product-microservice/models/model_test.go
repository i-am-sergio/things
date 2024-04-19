package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductModel(t *testing.T) {
	product := Product{}
    tableName := product.TableName()
    assert.Equal(t, "products", tableName, "Expected table name to be 'products'")
}

func TestCommentModel(t *testing.T) {
	comment := Comment{}
	tableName := comment.TableName()
	assert.Equal(t, "comments", tableName, "Expected table name to be 'comments'")
}