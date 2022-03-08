package services

import (
	"github.com/abbul/operacion-fuego-quasar/models"
	"github.com/abbul/operacion-fuego-quasar/repositories"
	"github.com/abbul/operacion-fuego-quasar/services"
	"github.com/abbul/operacion-fuego-quasar/tests/mock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetLocation(t *testing.T) {

	var x, y, err = services.GetLocation(mock.Distances...)

	assert.EqualValues(t, mock.PositionResult.X, x)
	assert.EqualValues(t, mock.PositionResult.Y, y)
	assert.EqualValues(t, nil, err)
}

func TestNormalizedLocations(t *testing.T) {
	var distances = services.NormalizedLocations(mock.TopSecrets)

	assert.EqualValues(t, mock.Distances, distances)
}

func TestGetLocationSplit(t *testing.T) {

	for _, satellite := range repositories.SATELLITES {
		topSecretSplit := models.RequestSplit{
			Name:     satellite.Name,
			Distance: 123.123,
			Message:  []string{"hello", "world", "!"},
		}
		services.AddTopSecret(topSecretSplit)
	}

	var x, y, err = services.GetLocationSplit()
	assert.IsType(t, 1.1, x)
	assert.IsType(t, 1.1, y)
	assert.EqualValues(t, nil, err)
}
