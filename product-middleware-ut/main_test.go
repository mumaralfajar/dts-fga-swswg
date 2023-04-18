package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_IsValidName_Success(t *testing.T) {
	result, err := IsValidName("John Doe")

	assert.Nil(t, err)

	assert.Equal(t, "valid name", result)
}

func Test_IsValidName_Error(t *testing.T) {
	result, err := IsValidName("John")

	assert.NotNil(t, err)

	assert.Empty(t, result)

	assert.Equal(t, "invalid name", err.Error())
}
