package algorithmia

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

func Query(algorithm string, api_key string, data io.Reader) ([]byte, error) {
	api_url := "http://api.algorithmia.com/api"
	req_url := strings.Join([]string{api_url, algorithm}, "/")

	req, err := http.NewRequest("POST", req_url, data)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", api_key)
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("Request was unsuccessful: %s", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
