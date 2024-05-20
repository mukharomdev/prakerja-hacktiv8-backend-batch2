package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidName(t *testing.T) {
	err := ValidName("john")

	assert.NotNil(t, err)

	assert.Equal(t, "invalid name", err.Error())

	err = ValidName("andre")

	assert.Nil(t, err)
}
