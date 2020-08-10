package responses

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// JSON stringify response
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

// ERROR returns an error
func ERROR(w http.ResponseWriter, statusCode int, err error) {
	if err != nil {
		w.WriteHeader(statusCode)
		JSON(w, statusCode, struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),
		})
	}
	JSON(w, http.StatusBadRequest, nil)
}
