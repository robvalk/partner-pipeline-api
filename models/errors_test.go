package models

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewErrorResponse_toReturnExpectedStructForJSONMarshaling(t *testing.T) {
	// given
	err := errors.New("this is an error")

	// when
	errResponse := NewErrorResponse(err)

	// then
	bytes, _ := json.Marshal(errResponse)
	assert.Equal(t, `{"error":{"message":"this is an error"}}`, string(bytes))
}
