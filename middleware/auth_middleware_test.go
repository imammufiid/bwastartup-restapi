package middleware

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	token string = "123123123123123"
	validToken string = "Bearer 123123123123123"
	invalidToken string = "Bearer123123123123123"
)

func TestSplitBearerTokenSuccess(t *testing.T) {
	result := splitBearerToken(validToken)
	assert.Equal(t, token, result, "Success split bearer token")
}

func TestSplitBearerTokenFailed(t *testing.T) {
	result := splitBearerToken(invalidToken)
	assert.Equal(t, "", result, "Failed split bearer token")
}