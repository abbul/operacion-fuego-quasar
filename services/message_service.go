package services

import (
	"errors"
	"github.com/abbul/operacion-fuego-quasar/models"
	"github.com/abbul/operacion-fuego-quasar/repositories"
	"strings"
)

func GetMessage(messages ...[]string) (msg string, err error) {
	var satellites, _ = repositories.GetSatellites()
	if len(messages) != len(satellites) {
		return "", errors.New("message-not_enough_information")
	}
	for _, message := range messages {
		for i, word := range message {
			if i >= len(repositories.GetMessage()) {
				repositories.PushWordToMessage(word)
				continue
			}
			if word == "" {
				continue
			}
			repositories.InsertionWordToMessage(i, word)
		}
	}

	return strings.Join(repositories.GetMessage(), " "), nil
}

func NormalizedMessages(topSecrets []models.TopSecret) (messages [][]string) {
	messages = [][]string{}
	for _, ts := range topSecrets {
		messages = append(messages, ts.Message)
	}
	return messages
}

func GetMessageSplit() (msg string, err error) {
	var messages = NormalizedMessages(repositories.GetTopSecrets())
	return GetMessage(messages...)
}
