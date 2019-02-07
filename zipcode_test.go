package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGetAddress(t *testing.T) {
	expected := Address{
		Street:       "São Paulo",
		Neighborhood: "Lapa",
		City:         "R Roma",
		State:        "SP",
	}

	// Configures test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(&expected)
	}))
	defer server.Close()
	host = server.URL

	addr, err := GetAddress("05050090")
	if err != nil {
		t.Fatal(err)
	}
	if addr == nil {
		t.Fatalf("Unexpected <nil> address result")
	}

	if !reflect.DeepEqual(*addr, expected) {
		t.Errorf("Unexpected address result:\nexpected:%#v\ngot:%#v", expected, *addr)
	}
}

func TestGetAddressNotFound(t *testing.T) {
	// Configures test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Not found", http.StatusNotFound)
	}))
	defer server.Close()
	host = server.URL

	addr, err := GetAddress("12345678")
	if err != nil {
		t.Fatal(err)
	}
	if addr != nil {
		t.Fatalf("Unexpected non <nil> address result")
	}
}
