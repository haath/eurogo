package api

import "testing"

func TestGet(t *testing.T) {

	client := Client{BaseURL: "https://jsonplaceholder.typicode.com"}

	ch := make(chan Response)

	go client.Get("/todos/1", ch)

	response := <-ch

	if response.Error != nil {
		t.Error(response.Error)
	}

	if response.Code != 200 {
		t.Fatalf("ApiClient.Get status code: %v", response.Code)
	}
}
