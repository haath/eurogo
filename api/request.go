package api

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

type Request interface {
	Endpoint(endpoint string) Request
	Set(key string, value string) Request
	Get(result chan<- Response)
}

// Request holds the parameters for making requests to an API endpoint.
type request struct {
	query *url.URL
}

func NewRequest(baseURL string) (Request, error) {

	req, err := url.Parse(baseURL)

	return &request{
		query: req,
	}, err
}

func (this *request) Endpoint(endpoint string) Request {

	this.query.Path = endpoint
	return this
}

func (this *request) Set(key string, value string) Request {

	query := this.query.Query()
	query.Set(key, value)
	this.query.RawQuery = query.Encode()
	return this
}

// Get is a goroutine which performs a GET request to the given endpoint.
func (this *request) Get(result chan<- Response) {

	fullURL := this.query.String()

	if os.Getenv("EUROGO_DEBUG") == "true" {
		log.Println(fullURL)
	}

	resp, err := http.Get(fullURL)
	if err != nil {
		result <- Response{Error: err}
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	result <- Response{Body: string(body), Error: err, Code: resp.StatusCode}
}
