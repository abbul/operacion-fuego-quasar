package repositories

import (
	"github.com/abbul/operacion-fuego-quasar/models"
)

var SECRETS []models.TopSecret

func GetTopSecrets() []models.TopSecret {
	return SECRETS
}

func PushTopSecrets(topSecret models.TopSecret) []models.TopSecret {
	SECRETS = append(SECRETS, topSecret)
	return SECRETS
}

func InsertionTopSecrets(i int, topSecret models.TopSecret) {
	SECRETS[i] = topSecret
}
