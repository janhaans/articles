package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllArticles(t *testing.T) {
	allarticles, err := GetAllArticles()

	assert.Nil(t, err)
	assert.NotNil(t, allarticles)
	assert.EqualValues(t, len(articles), len(allarticles))
	for k, v := range allarticles {
		assert.EqualValues(t, articles[k].ID, v.ID)
		assert.EqualValues(t, articles[k].Title, v.Title)
		assert.EqualValues(t, articles[k].Content, v.Content)
	}
}
