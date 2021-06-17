package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
)

// Greetings struct
// string
// type Greetings string
type Greetings struct {
	Message string `json:"message"`
}

// ShowMessage POST
// @tags Show message
// @Summary Show message
// @Description Show message
// @Accept  json
// @Produce  json
// @Param data body Greetings true "Show Body message"
// @Param Authorization header string true "Basic token"
// @Success 200 {object} string
// @SecuritySchemes BasicAuth
// @Router /  [post]
// 	func ShowMessage(w http.ResponseWriter, r *http.Request) {
// 		data := &Greetings{}
// 		err := json.NewDecoder(r.Body).Decode(&data)
// 		f err != nil {
// 			http.Error(w, err.Error(), http.StatusBadRequest)
// 			return
// 		}
// 		json.NewEncoder(w).Encode(data.Message)
// 	}
func ShowMessage(w http.ResponseWriter, r *http.Request) {

	data := &Greetings{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(strings.ToLower(reqToken), "basic")

	data.Message = data.Message + " - TOKEN: " + strings.TrimSpace(splitToken[1])
	json.NewEncoder(w).Encode(data.Message)
}
