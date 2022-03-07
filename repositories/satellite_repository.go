package repositories

import (
	"errors"
	"github.com/abbul/operacion-fuego-quasar/domain/top_secret"
	"log"
)

const (
	KENOBI    string = "kenobi"
	SKYWALKER string = "skywalker"
	SATO      string = "sato"
)

var SATELLITES = []top_secret.SatellitePosition{
	{Name: KENOBI, Position: top_secret.Position{X: -500, Y: -200}},
	{Name: SKYWALKER, Position: top_secret.Position{X: 100, Y: 100}},
	{Name: SATO, Position: top_secret.Position{X: 500, Y: 100}},
}

func GetSatelliteByName(name string) (satellite top_secret.SatellitePosition, err error) {
	for _, satellite := range SATELLITES {
		if name == satellite.Name {
			log.Println(satellite.Name)
			return satellite, nil
		}
	}
	log.Println(name)
	return top_secret.SatellitePosition{}, errors.New("satellite-not_found")
}

func GetSatellites() (satellite []top_secret.SatellitePosition, err error) {
	return SATELLITES, nil
}
