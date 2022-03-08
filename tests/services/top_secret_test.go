package services

import (
	"github.com/abbul/operacion-fuego-quasar/models"
	"github.com/abbul/operacion-fuego-quasar/services"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddTopSecret(t *testing.T) {
	var topSecretSplit = models.RequestSplit{
		Name:     "kenobi",
		Distance: 123.123,
		Message:  []string{"hello", "world", "!"},
	}
	var result, err = services.AddTopSecret(topSecretSplit)
	assert.EqualValues(t, result, true)
	assert.EqualValues(t, nil, err)
}
