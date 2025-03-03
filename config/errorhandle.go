package config

import "encoding/json"

// jsonError returns a JSON formatted error message
func JsonError(message string) string {
  errorMessage := map[string]string{"error": message}
  jsonMessage, _ := json.Marshal(errorMessage)
  return string(jsonMessage)
}
