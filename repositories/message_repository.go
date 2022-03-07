package repositories

var MESSAGE []string

func GetMessage() (message []string) {
	return MESSAGE
}

func PushWordToMessage(message string) {
	MESSAGE = append(MESSAGE, message)
}

func InsertionWordToMessage(i int, message string) {
	MESSAGE[i] = message
}
