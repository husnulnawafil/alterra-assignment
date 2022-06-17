package helper

func ResponseSuccess(message string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"status":  "Success",
		"message": message,
		"data":    data,
	}
}

func ResponseSuccessWithoutData(message string) map[string]interface{} {
	return map[string]interface{}{
		"status":  "Success",
		"message": message,
	}
}

func ResponseFailed(message string) map[string]interface{} {
	return map[string]interface{}{
		"status":  "Failed",
		"message": message,
	}
}
