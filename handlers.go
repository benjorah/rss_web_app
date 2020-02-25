package main

import (
	"encoding/json"
	"log"
	"net/http"
)

//handleFeedSearch is a handler for the `feeds` route of the API
//It only accepts one method GET as is required by the API
func (app *App) handleFeedSearch(w http.ResponseWriter, r *http.Request) {

	log.Printf("[INFO] Recived a %s request from %s \n", r.Method, r.Host)

	switch r.Method {
	case "GET":
		searchParam, ok := r.URL.Query()["search"]

		//check if the correct query oarameter is included in the request
		if !ok || len(searchParam[0]) < 1 {
			log.Println("[ERROR] App.handleFeedSearch(): Url Param 'search' is missing")
			returnResponseJSON(nil, http.StatusBadRequest, false, "Url Param 'search' is missing", w)
			return
		}

		//search the database for records matching the search string
		data, err := app.DBAdapter.SearchRecords(searchParam[0])

		if err != nil {
			log.Printf("[ERROR] App.handleFeedSearch() <= %s", err.Error())
		}

		returnResponseJSON(data, http.StatusOK, true, "", w)

	default:
		returnResponseJSON(nil, http.StatusBadRequest, false, "METHOD "+r.Method+" not allowed", w)

	}

}

//returnResponseJSON builds appropriate the response JSON object and sends it to the client.
//All handlers should use this response while commmunicating with the client in order to keep things consistent
func returnResponseJSON(data []RSSData, statusCode int, success bool, errorMessage string, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	var httpResponse HTTPResponse

	//build the appropriate HTTPResponse depending on if the response is to be an error or otherwise
	if !success {

		httpResponse = HTTPResponse{
			Success: success,

			ResponseError: ResponseError{
				statusCode,
				errorMessage,
			},
		}

	} else {

		httpResponse = HTTPResponse{
			Success: success,
			ResponsePayload: ResponsePayload{
				data,
				len(data),
			},
		}

	}

	w.WriteHeader(statusCode)

	//convert HTTPResponse to JSON
	JSONresponse, err := json.Marshal(httpResponse)
	if err != nil {
		JSONresponse = nil
		log.Println("[ERROR] returnResponseJSON() : encountered and error while converting data to JSON")
		w.WriteHeader(http.StatusInternalServerError)

	}

	w.Write(JSONresponse)

}
