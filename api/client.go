package api

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
)

// Client holds the parameters for making requests to an API endpoint.
type Client struct {
	BaseURL string
}

// Get is a goroutine which performs a GET request to the given endpoint.
func (client *Client) Get(endpoint string, result chan<- Response) {

	fullURL, err := client.getURL(endpoint)
	if err != nil {
		result <- Response{Error: err}
		return
	}

	resp, err := http.Get(fullURL)
	if err != nil {
		result <- Response{Error: err, Code: resp.StatusCode}
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	result <- Response{Body: string(body), Error: err, Code: resp.StatusCode}
}

func (client *Client) getURL(endpoint string) (string, error) {

	u, err := url.Parse(client.BaseURL)

	if err != nil {
		return "", err
	}

	u.Path = path.Join(u.Path, endpoint)
	return u.String(), nil
}
