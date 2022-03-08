package services

import (
	"github.com/abbul/operacion-fuego-quasar/models"
	"github.com/abbul/operacion-fuego-quasar/repositories"
	"github.com/abbul/operacion-fuego-quasar/services"
	"github.com/abbul/operacion-fuego-quasar/tests/mock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetMessage(t *testing.T) {

	var message, err = services.GetMessage(mock.Messages...)

	assert.EqualValues(t, message, mock.MessageResult)
	assert.EqualValues(t, nil, err)
}

func TestNormalizedMessages(t *testing.T) {
	var expectResult = [][]string{{"", "world", ""}, {"hello", "", ""}, {"", "", "!"}}
	var messages = services.NormalizedMessages(mock.TopSecrets)

	assert.EqualValues(t, expectResult, messages)
}

func TestGetMessageSplit(t *testing.T) {

	for _, satellite := range repositories.SATELLITES {
		topSecretSplit := models.RequestSplit{
			Name:     satellite.Name,
			Distance: 123.123,
			Message:  []string{"hello", "world", "!"},
		}
		services.AddTopSecret(topSecretSplit)
	}

	var message, err = services.GetMessageSplit()
	assert.EqualValues(t, "hello world !", message)
	assert.EqualValues(t, nil, err)
}
