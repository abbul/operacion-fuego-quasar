package repositories

import "github.com/abbul/operacion-fuego-quasar/domain/top_secret"

var SECRETS []top_secret.TopSecret

func GetTopSecrets() []top_secret.TopSecret {
	return SECRETS
}

func PushTopSecrets(topSecret top_secret.TopSecret) []top_secret.TopSecret {
	SECRETS = append(SECRETS, topSecret)
	return SECRETS
}

func InsertionTopSecrets(i int, topSecret top_secret.TopSecret) {
	SECRETS[i] = topSecret
}
