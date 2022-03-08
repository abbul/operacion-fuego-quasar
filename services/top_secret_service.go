package services

import (
	"github.com/abbul/operacion-fuego-quasar/models"
	"github.com/abbul/operacion-fuego-quasar/repositories"
)

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
