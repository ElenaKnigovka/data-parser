package data_parser

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// HTTPResponse represents an HTTP response.
type HTTPResponse struct {
	Body       string
	StatusCode int
}

// GetHTTPResponse makes an HTTP GET request to the provided URL and returns the response.
func GetHTTPResponse(url string) (*HTTPResponse, error) {
	resp, err := http.Get(url)
	if err!= nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode!= http.StatusOK {
		return nil, fmt.Errorf("HTTP error: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err!= nil {
		return nil, err
	}

	return &HTTPResponse{
		Body:       string(body),
		StatusCode: resp.StatusCode,
	}, nil
}

// ParseJSON parses the provided JSON string into a map.
func ParseJSON(jsonString string) (map[string]interface{}, error) {
	var data map[string]interface{}
	err := json.Unmarshal([]byte(jsonString), &data)
	return data, err
}

// ParseJSONBytes parses the provided JSON bytes into a map.
func ParseJSONBytes(jsonBytes []byte) (map[string]interface{}, error) {
	var data map[string]interface{}
	err := json.Unmarshal(jsonBytes, &data)
	return data, err
}

// TrimTrailingNewline removes trailing newlines from the provided string.
func TrimTrailingNewline(s string) string {
	return strings.TrimSuffix(s, "\n")
}