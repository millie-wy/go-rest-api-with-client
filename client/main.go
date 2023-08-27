package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-rest-api/lets-go/api"
	"io"
	"log"
	"net/http"
	"time"
)

type APIClient struct {
	BaseURL    string
	HTTPClient *http.Client
}

func newAPIClient(baseURL string, timeout time.Duration) *APIClient {
	client := &http.Client{
		Timeout: timeout,
	}
	return &APIClient{
		BaseURL:    baseURL,
		HTTPClient: client,
	}
}

func (c *APIClient) doRequest(method, endpoint string, body io.Reader) ([]byte, error) {
	url := c.BaseURL + endpoint

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	response, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return responseBody, nil
}

func (c *APIClient) GetEvents() ([]api.Event, error) {
	res, err := c.doRequest(http.MethodGet, "/events", nil)
	if err != nil {
		return nil, err
	}

	var events []api.Event
	if err := json.Unmarshal(res, &events); err != nil {
		return nil, err
	}
	return events, nil
}

func (c *APIClient) CreateEvent(event map[string]interface{}) error {
	eventJSON, _ := json.Marshal(event)
	_, err := c.doRequest(http.MethodPost, "/events", bytes.NewReader(eventJSON))
	return err
}

func (c *APIClient) UpdateEvent(id string, event map[string]interface{}) error {
	eventJSON, _ := json.Marshal(event)
	_, err := c.doRequest(http.MethodPut, "/events/"+id, bytes.NewReader(eventJSON))
	return err
}

func (c *APIClient) DeleteEvent(id string) error {
	_, err := c.doRequest(http.MethodDelete, "/events/"+id, nil)
	return err
}

func main() {

	// Fetch events using GetEvents method
	defaultTimeout := time.Second * 10
	client := newAPIClient("http://localhost:8080", defaultTimeout)
	events, err := client.GetEvents()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(events)

	// Example: Create a new event using the CreateEvent method
	newEvent := map[string]interface{}{
		"ID":          "2",
		"Title":       "New Event",
		"Description": "This is a new event",
	}
	err = client.CreateEvent(newEvent)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Event created successfully")

	// Example: Update an event using the UpdateEvent method
	eventToUpdate := map[string]interface{}{
		"ID":          "2",
		"Title":       "Updated Event",
		"Description": "This is the updated event",
	}
	err = client.UpdateEvent("1", eventToUpdate)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Event updated successfully")

	// Example: Delete an event using the DeleteEvent method
	err = client.DeleteEvent("1")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Event deleted successfully")

}
