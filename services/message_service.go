package services

import (
	"errors"
	"github.com/abbul/operacion-fuego-quasar/models"
	"github.com/abbul/operacion-fuego-quasar/repositories"
	"strings"
)

func GetLocation(distances float32) (x, y float32) {
	return distances, distances * 2
}

func GetMessage(messages []string) (msg string) {
	for i, s := range messages {
		if i >= len(repositories.GetMessage()) {
			repositories.PushWordToMessage(s)
			continue
		}
		if s == "" {
			continue
		}
		repositories.InsertionWordToMessage(i, s)
	}

	return strings.Join(repositories.GetMessage(), " ")
}

func JoinLocations(topSecrets []models.TopSecret) (x, y float32, err error) {
	var satellites, _ = repositories.GetSatellites()
	if len(topSecrets) != len(satellites) {
		return 0, 0, errors.New("location__not enough information")
	}
	for _, topSecret := range topSecrets {
		var satellite models.TopSecret = topSecret
		x, y = GetLocation(satellite.Distance)
	}
	return x, y, nil
}

func JoinMessages(topSecrets []models.TopSecret) (msg string, err error) {
	var satellites, _ = repositories.GetSatellites()
	if len(topSecrets) != len(satellites) {
		return "", errors.New("message-not_enough_information")
	}
	for _, topSecret := range topSecrets {
		var satellite models.TopSecret = topSecret
		GetMessage(satellite.Message)
	}
	return strings.Join(repositories.GetMessage(), " "), nil
}

func GetMessageSplit() (msg string, err error) {
	return JoinMessages(repositories.GetTopSecrets())
}

func GetLocationSplit() (x, y float32, err error) {
	return JoinLocations(repositories.GetTopSecrets())
}

func AddTopSecret(topSecretSplit models.RequestSplit) (result bool, err error) {
	satellite, err := repositories.GetSatelliteByName(topSecretSplit.Name)
	if err != nil {
		return false, err
	}

	var topSecret = models.TopSecret{Name: topSecretSplit.Name, Distance: topSecretSplit.Distance, Message: topSecretSplit.Message}
	var topSecrets = repositories.GetTopSecrets()
	for i, secret := range topSecrets {
		if satellite.Name == secret.Name {
			repositories.InsertionTopSecrets(i, topSecret)
			return true, nil
		}
	}
	repositories.PushTopSecrets(topSecret)
	return true, nil
}
