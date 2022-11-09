package test

import (
	"restful-api/simple"
	"testing"

	_ "github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/assert"
)

func TestSimpleServiceError(t *testing.T) {
	simpleService, err := simple.IntializedService(true)
	assert.NotNil(t, err)
	assert.Nil(t, simpleService)
}

func TestSimpleServiceTrue(t *testing.T) {
	simpleService, err := simple.IntializedService(false)
	assert.Nil(t, err)
	assert.NotNil(t, simpleService)
}
