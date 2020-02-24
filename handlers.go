package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func (app *App) handleFeedSearch(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		searchParam, ok := r.URL.Query()["search"]

		if !ok || len(searchParam[0]) < 1 {
			log.Println("Url Param 'search' is missing")
			returnResponseJSON(nil, http.StatusBadRequest, w)
			return
		}

		data, _ := app.Database.SearchRecords(searchParam[0])

		dataJSON, err := json.Marshal(data)
		if err != nil {

			returnResponseJSON(nil, http.StatusInternalServerError, w)

		}
		returnResponseJSON(dataJSON, http.StatusOK, w)

	default:
		returnResponseJSON(nil, http.StatusBadRequest, w)

	}

}

func returnResponseJSON(json []byte, statusCode int, w http.ResponseWriter) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(json)

}
