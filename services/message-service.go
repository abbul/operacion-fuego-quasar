package services

import (
	"github.com/abbul/operacion-fuego-quasar/domain/top_secret"
	"github.com/abbul/operacion-fuego-quasar/utils/cache"
	"strings"
)

func GetLocation(distance float32) (x, y float32) {
	cache.POSITION_X = distance
	cache.POSITION_Y = distance
	return cache.POSITION_X, cache.POSITION_Y
}

func JoinLocations(satellites []top_secret.Satellite) (x, y float32) {
	for i := range satellites {
		var satellite top_secret.Satellite = satellites[i]
		GetLocation(satellite.Distance)
	}
	return cache.POSITION_X, cache.POSITION_Y
}

func GetMessage(messages []string) (msg string) {
	for i, s := range messages {
		if i >= len(cache.MESSAGE) {
			cache.MESSAGE = append(cache.MESSAGE, s)
			continue
		}
		if s == "" {
			continue
		}
		cache.MESSAGE[i] = s
	}

	return strings.Join(cache.MESSAGE, " ")
}

func JoinMessages(satellites []top_secret.Satellite) (msg string) {
	for i := range satellites {
		var satellite top_secret.Satellite = satellites[i]
		GetMessage(satellite.Message)
	}
	return strings.Join(cache.MESSAGE, " ")
}
