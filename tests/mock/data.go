package mock

import "github.com/abbul/operacion-fuego-quasar/models"

var TopSecrets = []models.TopSecret{
	{Name: "a", Distance: 100.0, Message: []string{"", "world", ""}},
	{Name: "b", Distance: 115.5, Message: []string{"hello", "", ""}},
	{Name: "c", Distance: 142.7, Message: []string{"", "", "!"}},
}

var Distances = []float64{TopSecrets[0].Distance, TopSecrets[1].Distance, TopSecrets[2].Distance}

var Messages = [][]string{TopSecrets[0].Message, TopSecrets[1].Message, TopSecrets[2].Message}

var PositionResult = models.Position{
	X: 31.47972271378469,
	Y: -4.255193579550551,
}

var MessageResult = "hello world !"
