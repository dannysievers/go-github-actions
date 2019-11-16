package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetSongs(t *testing.T) {
	req, err := http.NewRequest("GET", "/songs", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getSongs)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code : got %v and expected %v", status, http.StatusOK)
	}
}

func TestCreateSong(t *testing.T) {
	jsonStr := []byte(`{"id": "123", "title": "testing", "artist": "tester"}`)
	req, err := http.NewRequest("PUT", "/songs", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(createSong)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handle rreturned wrong status code : got %v and expected %v", status, http.StatusOK)
	}
}

// func TestTestyMctesterson(t *testing.T) {
// 	expect := "YOU DIDN'T TEST ME"
// 	actual := testyMctesterson()
// 	if actual != expect {
// 		t.Errorf("Got %v and expected %v", actual, expect)
// 	}
// }
