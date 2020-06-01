package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestHealthHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(healthHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := apiResponse{
		Message:     "The server is alive",
		Description: "No Errors detected",
	}

	var response apiResponse
	json.NewDecoder(rr.Body).Decode(&response)
	if response != expected {
		t.Errorf("handler returned unexpected body: got %v but want %v",
			response, expected)
	}
}

/*These tests are not working */

// func TestEntryExistsSearchHandler(t *testing.T) {
// 	req, err := http.NewRequest("GET", "/search/drake", nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	rr := httptest.NewRecorder()
// 	handler := http.HandlerFunc(searchHandler)
// 	handler.ServeHTTP(rr, req)
// 	if status := rr.Code; status != http.StatusOK {
// 		t.Errorf("handler returned wrong status code: got %v want %v",
// 			status, http.StatusOK)
// 	}
// 	//Response are "Randomized", unsure of how to check content
// }

// func TestEntryNoExistsSearchHandler(t *testing.T) {
// 	searchFor := "jalkdjlajdl"
// 	req, err := http.NewRequest("GET", "/search/"+searchFor, nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	rr := httptest.NewRecorder()
// 	handler := http.HandlerFunc(searchHandler)
// 	handler.ServeHTTP(rr, req)
// 	if status := rr.Code; status != http.StatusNotFound {
// 		t.Errorf("handler returned wrong status code: got %v want %v",
// 			status, http.StatusNotFound)
// 	}

// 	// Check the response body is what we expect.
// 	expected := apiResponse{
// 		Message:     "Search Term Not Found",
// 		Description: "Unable to find " + searchFor,
// 	}

// 	var response apiResponse
// 	json.NewDecoder(rr.Body).Decode(&response)
// 	if response != expected {
// 		t.Errorf("handler returned unexpected body: got %v but want %v",
// 			response, expected)
// 	}
// }
