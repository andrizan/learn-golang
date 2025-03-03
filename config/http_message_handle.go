package config

import (
	"encoding/json"
	"net/http"
)

// JsonMarkup formats the response data
func JsonMarkup(data interface{}, success bool) map[string]interface{} {
	if success {
		return map[string]interface{}{"success": true, "data": data}
	}
	return map[string]interface{}{"success": false, "error": data}
}

// WriteJSONResponse simplifies writing JSON responses
func WriteJSONResponse(w http.ResponseWriter, data interface{}, statusCode int,) {
	if statusCode == 0 {
		statusCode = http.StatusOK
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(JsonMarkup(data, statusCode == http.StatusOK))
}
