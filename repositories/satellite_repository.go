package repositories

import (
	"errors"
	"github.com/abbul/operacion-fuego-quasar/models"
)

const (
	KENOBI    string = "kenobi"
	SKYWALKER string = "skywalker"
	SATO      string = "sato"
)

var SATELLITES = []models.SatellitePosition{
	{Name: KENOBI, Position: models.Position{X: -500, Y: -200}},
	{Name: SKYWALKER, Position: models.Position{X: 100, Y: 100}},
	{Name: SATO, Position: models.Position{X: 500, Y: 100}},
}

func GetSatelliteByName(name string) (satellite models.SatellitePosition, err error) {
	for _, satellite := range SATELLITES {
		if name == satellite.Name {
			return satellite, nil
		}
	}
	return models.SatellitePosition{}, errors.New("satellite-not_found")
}

func GetSatellites() (satellite []models.SatellitePosition, err error) {
	return SATELLITES, nil
}
