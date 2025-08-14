// Package client provides HTTP client functionality for interacting with MCPJungle API.
package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/mcpjungle/mcpjungle/internal/api"
	"github.com/mcpjungle/mcpjungle/internal/model"
)

// Client represents a client for interacting with the MCPJungle HTTP API
type Client struct {
	baseURL     string
	accessToken string
	httpClient  *http.Client
}

func NewClient(baseURL, accessToken string, httpClient *http.Client) *Client {
	return &Client{
		baseURL:     baseURL,
		accessToken: accessToken,
		httpClient:  httpClient,
	}
}

type InitServerResponse struct {
	AdminAccessToken string `json:"admin_access_token"`
}

// InitServer sends a request to initialize the server in production mode
func (c *Client) InitServer() (*InitServerResponse, error) {
	u, _ := url.JoinPath(c.baseURL, "/init")

	payload := struct {
		Mode string `json:"mode"`
	}{
		Mode: string(model.ModeProd),
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Post(u, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("failed to send request to %s: %w", u, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("request failed with status: %d, message: %s", resp.StatusCode, body)
	}

	var initResp InitServerResponse
	if err := json.NewDecoder(resp.Body).Decode(&initResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &initResp, nil
}

// constructAPIEndpoint constructs the full API endpoint URL where a request must be sent
func (c *Client) constructAPIEndpoint(suffixPath string) (string, error) {
	return url.JoinPath(c.baseURL, api.V0PathPrefix, suffixPath)
}

// newRequest creates a new HTTP request with the specified method, URL, and body.
// It automatically adds the Authorization header if an access token is present.
func (c *Client) newRequest(method, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	if c.accessToken != "" {
		req.Header.Set("Authorization", "Bearer "+c.accessToken)
	}

	return req, nil
}
