package json_domain

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPostJson(t *testing.T) {
	request := Post{
		Title:  "Awesome Title",
		Body:   "Great Body",
		UserId: 1,
	}
	bytes, err := json.Marshal(request)
	assert.Nil(t, err)
	assert.NotNil(t, bytes)

	var result Post
	err = json.Unmarshal(bytes, &result)

	assert.Nil(t, err)
	assert.EqualValues(t,result.Title, request.Title)
	assert.EqualValues(t,result.Body, request.Body)
	assert.EqualValues(t,result.UserId, request.UserId)
}