package repository

import (
	"testing"

	"github.com/muesli/cache2go"
	"github.com/praveeennagaraj97/gqlgen/model"
	"github.com/stretchr/testify/assert"
)

func TestArticleRepository_GetAll(t *testing.T) {
	db := cache2go.Cache("example")
	repo := NewArticleRepository(db)

	article, err := repo.Save(&model.Article{
		Description: "Some Description",
		Content:     "Some Content",
	})
	assert.NoError(t, err)
	assert.NotNil(t, article)
	articles, err := repo.GetAll()
	assert.Len(t, articles, 1)
	assert.Equal(t, article, articles[0])
}
