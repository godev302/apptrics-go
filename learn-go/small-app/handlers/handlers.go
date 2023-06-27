package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"small-app/data"
	"strconv"
)

// GetUser is entry point for /search endpoint
// think how would you handle the request when someone hit this endpoint
func GetUser(w http.ResponseWriter, r *http.Request) {

	// this line set your  ContentType as json
	w.Header().Set("Content-Type", "application/json")

	//fetching the variable from query
	userIdString := r.URL.Query().Get("user_id")

	//converting it to make sure it is a valid uint64
	userId, err := strconv.ParseUint(userIdString, 10, 64)
	if err != nil {
		log.Println(err)
		appErr := map[string]string{"Message": http.StatusText(http.StatusBadRequest)}

		w.WriteHeader(http.StatusBadRequest)
		jsonData, err := json.Marshal(appErr) // converting map to json and sending back to the client using responseWriter
		if err != nil {
			log.Println(err)
			return
		}
		w.Write(jsonData)
		return
	}

	//fetching the user with the userId provided
	u, err := data.FetchUser(userId)
	if err != nil {
		log.Println(err)
		appErr := map[string]string{"Message": "user not found"}

		w.WriteHeader(http.StatusBadRequest)
		jsonData, err := json.Marshal(appErr) // converting map to json and sending back to the client using responseWriter
		if err != nil {
			log.Println(err)
			return
		}
		w.Write(jsonData)
		return
	}

	json.NewEncoder(w).Encode(u)
}
