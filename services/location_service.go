package services

import (
	"errors"
	"github.com/abbul/operacion-fuego-quasar/models"
	"github.com/abbul/operacion-fuego-quasar/repositories"
	"github.com/wilhelmstoll/trilaterate"
)

func GetLocation(distances ...float64) (x, y float64, err error) {
	var satellites, _ = repositories.GetSatellites()
	if len(distances) != len(satellites) {
		return 0, 0, errors.New("location-not_enough_information")
	}
	b1 := trilaterate.Beacon{Lat: satellites[0].Position.X, Lon: satellites[0].Position.Y, Dist: distances[0]}
	b2 := trilaterate.Beacon{Lat: satellites[1].Position.X, Lon: satellites[1].Position.Y, Dist: distances[1]}
	b3 := trilaterate.Beacon{Lat: satellites[2].Position.X, Lon: satellites[2].Position.Y, Dist: distances[2]}

	lat, lon := trilaterate.Solve(b1, b2, b3)
	return lat, lon, nil
}

func NormalizedLocations(topSecrets []models.TopSecret) (distances []float64) {
	distances = []float64{}
	for _, ts := range topSecrets {
		distances = append(distances, ts.Distance)
	}
	return distances
}

func GetLocationSplit() (x, y float64, err error) {
	var locations = NormalizedLocations(repositories.GetTopSecrets())
	return GetLocation(locations...)
}
