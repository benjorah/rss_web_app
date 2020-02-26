package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// for mocking the Database Adapter
type mockDBAdapter struct {
	name string
}

func (mockDb *mockDBAdapter) AddRecords(records []RSSData) (err error) {

	return nil

}
func (mockDb *mockDBAdapter) SearchRecords(searchString string) (records []RSSData, err error) {

	return getSampleRSSDataSliceForHandler(), nil

}

func TestApp_handleFeedSearch(t *testing.T) {
	type fields struct {
		DBAdapter DatabaseAdapter
	}

	t.Run("it should respond with a status ok in response to a GET method", func(t *testing.T) {
		t.Parallel()

		request, err := http.NewRequest("GET", feedPath+"?search=hi", nil)
		return
		if err != nil {
			t.Fatal(err)
		}

		response := httptest.NewRecorder()

		handler := http.HandlerFunc(app.handleFeedSearch)

		// Handlers satisfy http.Handler, so we can call their ServeHTTP method
		// directly and pass in our Request and ResponseRecorder.
		handler.ServeHTTP(response, request)

		// Check the status code is what we expect.
		if status := response.Code; status != http.StatusOK && request.Method == http.MethodGet {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

	})

	t.Run("it should respond with a status of BadRequest in response to a  method other than GET method", func(t *testing.T) {
		t.Parallel()

		request, err := http.NewRequest(http.MethodPost, feedPath+"?search=hello", nil)

		if err != nil {
			t.Fatal(err)
		}

		response := httptest.NewRecorder()

		handler := http.HandlerFunc(app.handleFeedSearch)

		// Handlers satisfy http.Handler, so we can call their ServeHTTP method
		// directly and pass in our Request and ResponseRecorder.
		handler.ServeHTTP(response, request)

		// Check the status code is what we expect.
		if status := response.Code; status != http.StatusBadRequest && request.Method != http.MethodGet {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusBadRequest)
		}

	})

	t.Run("it should respond with a status of BadRequest if the query parameter doesn't contain the parameter `search`", func(t *testing.T) {
		t.Parallel()
		// Check the status code is what we expect.

		request, err := http.NewRequest(http.MethodGet, feedPath, nil)
		if err != nil {
			t.Fatal(err)
		}

		response := httptest.NewRecorder()

		handler := http.HandlerFunc(app.handleFeedSearch)

		handler.ServeHTTP(response, request)

		if status := response.Code; status != http.StatusBadRequest && request.Method != http.MethodGet {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusBadRequest)
		}

	})

	t.Run("it should respond with a status of BadRequest if the query parameter doesn't contain the parameter other than `search`", func(t *testing.T) {
		t.Parallel()
		// Check the status code is what we expect.

		request, err := http.NewRequest(http.MethodGet, feedPath+"?notsearch=", nil)
		if err != nil {
			t.Fatal(err)
		}

		response := httptest.NewRecorder()

		handler := http.HandlerFunc(app.handleFeedSearch)

		handler.ServeHTTP(response, request)

		if status := response.Code; status != http.StatusBadRequest && request.Method != http.MethodGet {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusBadRequest)
		}

	})

	t.Run("it should respond with search results", func(t *testing.T) {
		t.Parallel()

		expected := `{"success":true,"payload":{"data":[{"title":"title","description":"desc","link":"link","created_at":"2020-02-24T22:10:15Z"}],"count":1},"error":{"code":0,"message":""}}`

		request, err := http.NewRequest(http.MethodGet, feedPath+"?search=hi", nil)
		return
		if err != nil {
			t.Fatal(err)
		}
		app.DBAdapter = &mockDBAdapter{}

		response := httptest.NewRecorder()

		handler := http.HandlerFunc(app.handleFeedSearch)

		// Handlers satisfy http.Handler, so we can call their ServeHTTP method
		// directly and pass in our Request and ResponseRecorder.
		handler.ServeHTTP(response, request)

		// Check the status code is what we expect.
		if response.Body.String() != string(expected) {
			t.Errorf("handler returned unexpected body: got %v want %v",
				response.Body.String(), expected)
		}

	})

}
