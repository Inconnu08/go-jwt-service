package handlers

import (
	"log"
	"net/http"
)

func respond(w http.ResponseWriter, v string, statusCode int) {
	//b, err := json.Marshal(v)
	//if err != nil {
	//	respondErr(w, fmt.Errorf("could not marshal response: %v", err))
	//	return
	//}

	w.Header().Set("Content-Type", "text/plain") // as plain text instead of "application/json" type
	w.WriteHeader(statusCode)
	w.Write([]byte(v))
}

func respondErr(w http.ResponseWriter, err error) {
	log.Println(err)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}