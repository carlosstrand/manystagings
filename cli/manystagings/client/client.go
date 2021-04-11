package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/carlosstrand/manystagings/models"
)

type Client struct {
	baseURL string
}

type EnvironmentList struct {
	Data  []*models.Environment `json:"data"`
	Count int64                 `json:"count"`
}

func NewClient(baseURL string) *Client {
	return &Client{baseURL: baseURL}
}

func (c *Client) withBaseURL(endpoint string) string {
	return c.baseURL + endpoint
}

func (c *Client) GetEnvironments(ctx context.Context) (*EnvironmentList, error) {
	var envList EnvironmentList
	path := c.withBaseURL("/environments")
	fmt.Println(path)
	res, err := http.Get(path)
	if err != nil {
		return nil, err
	}
	json.NewDecoder(res.Body).Decode(&envList)
	return &envList, nil
}
