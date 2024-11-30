package stresstest

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ValidateHeaders(t *testing.T) {
	assert.True(t, ValidateHeaders([]string{
		"Authorization:Bearer 12381923812983",
		"Content-Type:application/json",
		"X-Tenant:xpto.com",
		"API_KEY:acb123",
	}))
}

func Test_ValidateHeaders_Invalid(t *testing.T) {
	assert.False(t, ValidateHeaders([]string{"Authorization:"}))
	assert.False(t, ValidateHeaders([]string{":application/json"}))
	assert.False(t, ValidateHeaders([]string{"A=B"}))
	assert.False(t, ValidateHeaders([]string{"API_KEY:acb123:"}))
}
