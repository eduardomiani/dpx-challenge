package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"time"
)

// Static token for this challange
const token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJkZXNmaW8uZm90b3JlZ2lzdHJvLmNvbS5iciIsImV4cCI6MTU3NzU1NDEzMywianRpIjoiNzBlODRlZmQtMGRmNC00ZmZhLTlmYTYtNTI1M2ZjNmFmMDgyIiwiaWF0IjoxNTQ2NDUwMTMzLCJpc3MiOiJodHRwczovL3NlcnZpY2UtaG9tb2xvZy5kaWdpcGl4LmNvbS5iciIsInN0b3JlSWQiOjc5LCJzdG9yZU5hbWUiOiJGb3RvcmVnaXN0cm8iLCJzdG9yZVVSTCI6ImRlc2Zpby5mb3RvcmVnaXN0cm8uY29tLmJyIn0.yPFKdRdc4jTAUuziqfkvJm74W5axDelkaH-Q6lBTE8k"

// Default service enpoint. Can be replaced for unit tests
var host = "https://service-homolog.digipix.com.br/v0b"

// Address represents a address json structure
type Address struct {
	Name         string `json:"name"`
	ZipCode      string `json:"zipcode"`
	Street       string `json:"street"`
	Neighborhood string `json:"neighborhood"`
	State        string `json:"state_short"`
	City         string `json:"city"`
}

// GetAddress makes a call to digipix api server, given a zipcode.
// Returns the address found or <nil>
func GetAddress(zipcode string) (*Address, error) {
	// Setup request
	endpoint := fmt.Sprintf("%s/shipments/zipcode/%s", host, zipcode)
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	logServer("Making request...")
	dump, _ := httputil.DumpRequest(req, false)
	logServer("%s", dump)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	dump, _ = httputil.DumpResponse(resp, false)
	logServer("%s", dump)

	// Reads the response result
	var body []byte
	if resp.Body != nil {
		body, err = ioutil.ReadAll(resp.Body)

	}
	logServer("%v: %s", resp.StatusCode, body)
	switch resp.StatusCode {
	case http.StatusOK:
		var addr Address
		if err := json.Unmarshal(body, &addr); err != nil {
			return nil, err
		}
		return &addr, nil
	case http.StatusNotFound:
		return nil, nil
	default:
		return nil, fmt.Errorf("Response Status %d: %s", resp.StatusCode, body)
	}
}

// logServer log default
func logServer(msg string, args ...interface{}) {
	serverTime := fmt.Sprintf("[%s]: %s", time.Now().Format("2006-01-02 15:04:05"), msg)
	log.Printf(serverTime, args...)
}
