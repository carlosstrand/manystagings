package client

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/carlosstrand/manystagings/core/service"
	"github.com/carlosstrand/manystagings/models"
	"github.com/go-zepto/zepto/plugins/auth/authcore"
)

var ErrUnauthorized = errors.New("unauthorized")

type Client struct {
	baseURL   string
	client    *http.Client
	authToken string
}

type EnvironmentList struct {
	Data  []*models.Environment `json:"data"`
	Count int64                 `json:"count"`
}

func NewClient(baseURL string) *Client {
	return &Client{
		baseURL: baseURL,
		client:  &http.Client{},
	}
}

func (c *Client) withBaseURL(endpoint string) string {
	return c.baseURL + endpoint
}

func (c *Client) SetAuthToken(authToken string) {
	c.authToken = authToken
}

func errStatusFromRes(res *http.Response) error {
	return errors.New(http.StatusText(res.StatusCode))
}

func (c *Client) Auth(ctx context.Context, username string, password string) (*authcore.Token, error) {
	path := c.withBaseURL("/auth")
	fmt.Println(path)
	credentials := authcore.AuthCredentials{
		Username: username,
		Password: password,
	}
	authJson, err := json.Marshal(credentials)
	if err != nil {
		return nil, err
	}
	res, err := http.Post(path, "application/json", bytes.NewReader(authJson))
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, errStatusFromRes(res)
	}
	var tokenRes struct {
		Token authcore.Token `json:"token"`
	}
	err = json.NewDecoder(res.Body).Decode(&tokenRes)
	if err != nil {
		return nil, err
	}
	return &tokenRes.Token, nil
}

func (c *Client) GetInfo(ctx context.Context) (*service.Info, error) {
	var info service.Info
	path := c.withBaseURL("/api/info")
	req, err := http.NewRequest("GET", path, nil)
	if c.authToken != "" {
		req.Header.Add("Authorization", "Bearer "+c.authToken)
	}
	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, errStatusFromRes(res)
	}
	json.NewDecoder(res.Body).Decode(&info)
	return &info, nil
}

func (c *Client) GetEnvironments(ctx context.Context) (*EnvironmentList, error) {
	var envList EnvironmentList
	path := c.withBaseURL("/api/environments")
	req, err := http.NewRequest("GET", path, nil)
	if c.authToken != "" {
		req.Header.Add("Authorization", "Bearer "+c.authToken)
	}
	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, errStatusFromRes(res)
	}
	json.NewDecoder(res.Body).Decode(&envList)
	return &envList, nil
}