package utils

func JsonError(message string) string {
	return `{"message": "` + message + `"}`
}
