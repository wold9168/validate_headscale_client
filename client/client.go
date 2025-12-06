package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/wold9168/validate_headscale_client/models"
)

// Client represents the Headscale API client
type Client struct {
	BaseURL    string
	HTTPClient *http.Client
	APIKey     string
}

// NewClient creates a new Headscale API client
func NewClient(baseURL, apiKey string) *Client {
	return &Client{
		BaseURL: baseURL,
		HTTPClient: &http.Client{
			Timeout: 30 * time.Second,
		},
		APIKey: apiKey,
	}
}

// doRequest performs an HTTP request to the API
func (c *Client) doRequest(method, endpoint string, body interface{}) (*http.Response, error) {
	url := c.BaseURL + endpoint

	var reqBody io.Reader
	if body != nil {
		jsonData, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
		reqBody = bytes.NewBuffer(jsonData)
	}

	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	if c.APIKey != "" {
		req.Header.Set("Authorization", "Bearer "+c.APIKey)
	}

	// Perform the request
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to perform request: %w", err)
	}

	return resp, nil
}

// parseResponse parses the HTTP response into the provided interface
func (c *Client) parseResponse(resp *http.Response, out interface{}) error {
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(body))
	}

	if out != nil {
		if err := json.Unmarshal(body, out); err != nil {
			return fmt.Errorf("failed to unmarshal response: %w", err)
		}
	}

	return nil
}

// GetHealth checks the health of the Headscale API
func (c *Client) GetHealth() (*models.HealthResponse, error) {
	resp, err := c.doRequest("GET", "/api/v1/health", nil)
	if err != nil {
		return nil, err
	}

	var healthResponse models.HealthResponse
	err = c.parseResponse(resp, &healthResponse)
	if err != nil {
		return nil, err
	}

	return &healthResponse, nil
}

// ListApiKeys lists all API keys
func (c *Client) ListApiKeys() (*models.ListApiKeysResponse, error) {
	resp, err := c.doRequest("GET", "/api/v1/apikey", nil)
	if err != nil {
		return nil, err
	}

	var apiKeysResponse models.ListApiKeysResponse
	err = c.parseResponse(resp, &apiKeysResponse)
	if err != nil {
		return nil, err
	}

	return &apiKeysResponse, nil
}

// CreateApiKey creates a new API key
func (c *Client) CreateApiKey(req *models.CreateApiKeyRequest) (*models.CreateApiKeyResponse, error) {
	resp, err := c.doRequest("POST", "/api/v1/apikey", req)
	if err != nil {
		return nil, err
	}

	var apiKeyResponse models.CreateApiKeyResponse
	err = c.parseResponse(resp, &apiKeyResponse)
	if err != nil {
		return nil, err
	}

	return &apiKeyResponse, nil
}

// ExpireApiKey expires an API key
func (c *Client) ExpireApiKey(prefix string) (*models.ExpireApiKeyResponse, error) {
	req := &models.ExpireApiKeyRequest{
		Prefix: prefix,
	}
	
	resp, err := c.doRequest("POST", "/api/v1/apikey/expire", req)
	if err != nil {
		return nil, err
	}

	var expireResponse models.ExpireApiKeyResponse
	err = c.parseResponse(resp, &expireResponse)
	if err != nil {
		return nil, err
	}

	return &expireResponse, nil
}

// DeleteApiKey deletes an API key by prefix
func (c *Client) DeleteApiKey(prefix string) (*models.DeleteApiKeyResponse, error) {
	resp, err := c.doRequest("DELETE", fmt.Sprintf("/api/v1/apikey/%s", prefix), nil)
	if err != nil {
		return nil, err
	}

	var deleteResponse models.DeleteApiKeyResponse
	err = c.parseResponse(resp, &deleteResponse)
	if err != nil {
		return nil, err
	}

	return &deleteResponse, nil
}

// ListNode gets a specific node
func (c *Client) GetNode(nodeId int64) (*models.GetNodeResponse, error) {
	resp, err := c.doRequest("GET", fmt.Sprintf("/api/v1/node/%d", nodeId), nil)
	if err != nil {
		return nil, err
	}

	var nodeResponse models.GetNodeResponse
	err = c.parseResponse(resp, &nodeResponse)
	if err != nil {
		return nil, err
	}

	return &nodeResponse, nil
}

// ListNodes lists all nodes
func (c *Client) ListNodes(userId *string) (*models.ListNodesResponse, error) {
	url := "/api/v1/node"
	if userId != nil {
		url += "?user=" + *userId
	}
	
	resp, err := c.doRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var nodesResponse models.ListNodesResponse
	err = c.parseResponse(resp, &nodesResponse)
	if err != nil {
		return nil, err
	}

	return &nodesResponse, nil
}

// DeleteNode deletes a node
func (c *Client) DeleteNode(nodeId int64) (*models.DeleteNodeResponse, error) {
	resp, err := c.doRequest("DELETE", fmt.Sprintf("/api/v1/node/%d", nodeId), nil)
	if err != nil {
		return nil, err
	}

	var deleteResponse models.DeleteNodeResponse
	err = c.parseResponse(resp, &deleteResponse)
	if err != nil {
		return nil, err
	}

	return &deleteResponse, nil
}

// RenameNode renames a node
func (c *Client) RenameNode(nodeId int64, newName string) (*models.RenameNodeResponse, error) {
	resp, err := c.doRequest("POST", fmt.Sprintf("/api/v1/node/%d/rename/%s", nodeId, newName), nil)
	if err != nil {
		return nil, err
	}

	var renameResponse models.RenameNodeResponse
	err = c.parseResponse(resp, &renameResponse)
	if err != nil {
		return nil, err
	}

	return &renameResponse, nil
}

// ExpireNode expires a node
func (c *Client) ExpireNode(nodeId int64) (*models.ExpireNodeResponse, error) {
	resp, err := c.doRequest("POST", fmt.Sprintf("/api/v1/node/%d/expire", nodeId), nil)
	if err != nil {
		return nil, err
	}

	var expireResponse models.ExpireNodeResponse
	err = c.parseResponse(resp, &expireResponse)
	if err != nil {
		return nil, err
	}

	return &expireResponse, nil
}

// SetNodeTags sets tags for a node
func (c *Client) SetNodeTags(nodeId int64, tags []string) (*models.SetTagsResponse, error) {
	req := &models.SetTagsBody{
		Tags: tags,
	}
	
	resp, err := c.doRequest("POST", fmt.Sprintf("/api/v1/node/%d/tags", nodeId), req)
	if err != nil {
		return nil, err
	}

	var tagsResponse models.SetTagsResponse
	err = c.parseResponse(resp, &tagsResponse)
	if err != nil {
		return nil, err
	}

	return &tagsResponse, nil
}

// MoveNode moves a node to a different user
func (c *Client) MoveNode(nodeId int64, userId string) (*models.MoveNodeResponse, error) {
	req := &models.HeadscaleServiceMoveNodeBody{
		User: userId,
	}

	resp, err := c.doRequest("POST", fmt.Sprintf("/api/v1/node/%d/user", nodeId), req)
	if err != nil {
		return nil, err
	}

	var moveResponse models.MoveNodeResponse
	err = c.parseResponse(resp, &moveResponse)
	if err != nil {
		return nil, err
	}

	return &moveResponse, nil
}

// SetApprovedRoutes sets approved routes for a node
func (c *Client) SetApprovedRoutes(nodeId int64, routes []string) (*models.SetApprovedRoutesResponse, error) {
	req := &models.SetApprovedRoutesBody{
		Routes: routes,
	}
	
	resp, err := c.doRequest("POST", fmt.Sprintf("/api/v1/node/%d/approve_routes", nodeId), req)
	if err != nil {
		return nil, err
	}

	var routesResponse models.SetApprovedRoutesResponse
	err = c.parseResponse(resp, &routesResponse)
	if err != nil {
		return nil, err
	}

	return &routesResponse, nil
}

// RegisterNode registers a new node
func (c *Client) RegisterNode(machinePublicKey, authKey string) (*models.RegisterNodeResponse, error) {
	req := &models.RegisterNodeRequest{
		MachinePublicKey: machinePublicKey,
		AuthKey:          authKey,
	}
	
	resp, err := c.doRequest("POST", "/api/v1/node/register", req)
	if err != nil {
		return nil, err
	}

	var registerResponse models.RegisterNodeResponse
	err = c.parseResponse(resp, &registerResponse)
	if err != nil {
		return nil, err
	}

	return &registerResponse, nil
}

// BackfillNodeIPs backfills node IP addresses
func (c *Client) BackfillNodeIPs(nodeId int64) (*models.BackfillNodeIPsResponse, error) {
	resp, err := c.doRequest("POST", fmt.Sprintf("/api/v1/node/%d/backfillips", nodeId), nil)
	if err != nil {
		return nil, err
	}

	var backfillResponse models.BackfillNodeIPsResponse
	err = c.parseResponse(resp, &backfillResponse)
	if err != nil {
		return nil, err
	}

	return &backfillResponse, nil
}

// GetPolicy gets the current ACL policy
func (c *Client) GetPolicy() (*models.GetPolicyResponse, error) {
	resp, err := c.doRequest("GET", "/api/v1/policy", nil)
	if err != nil {
		return nil, err
	}

	var policyResponse models.GetPolicyResponse
	err = c.parseResponse(resp, &policyResponse)
	if err != nil {
		return nil, err
	}

	return &policyResponse, nil
}

// SetPolicy sets the ACL policy
func (c *Client) SetPolicy(policy string) (*models.SetPolicyResponse, error) {
	req := &models.SetPolicyRequest{
		Policy: policy,
	}
	
	resp, err := c.doRequest("PUT", "/api/v1/policy", req)
	if err != nil {
		return nil, err
	}

	var setPolicyResponse models.SetPolicyResponse
	err = c.parseResponse(resp, &setPolicyResponse)
	if err != nil {
		return nil, err
	}

	return &setPolicyResponse, nil
}

// ListPreAuthKeys lists all pre-authentication keys
func (c *Client) ListPreAuthKeys(userId string) (*models.ListPreAuthKeysResponse, error) {
	resp, err := c.doRequest("GET", fmt.Sprintf("/api/v1/preauthkey?user=%s", userId), nil)
	if err != nil {
		return nil, err
	}

	var preAuthKeysResponse models.ListPreAuthKeysResponse
	err = c.parseResponse(resp, &preAuthKeysResponse)
	if err != nil {
		return nil, err
	}

	return &preAuthKeysResponse, nil
}

// CreatePreAuthKey creates a new pre-authentication key
func (c *Client) CreatePreAuthKey(req *models.CreatePreAuthKeyRequest) (*models.CreatePreAuthKeyResponse, error) {
	resp, err := c.doRequest("POST", "/api/v1/preauthkey", req)
	if err != nil {
		return nil, err
	}

	var preAuthKeyResponse models.CreatePreAuthKeyResponse
	err = c.parseResponse(resp, &preAuthKeyResponse)
	if err != nil {
		return nil, err
	}

	return &preAuthKeyResponse, nil
}

// ExpirePreAuthKey expires a pre-authentication key
func (c *Client) ExpirePreAuthKey(userId, key string) (*models.ExpirePreAuthKeyResponse, error) {
	req := &models.ExpirePreAuthKeyRequest{
		User: userId,
		Key:  key,
	}
	
	resp, err := c.doRequest("POST", "/api/v1/preauthkey/expire", req)
	if err != nil {
		return nil, err
	}

	var expireResponse models.ExpirePreAuthKeyResponse
	err = c.parseResponse(resp, &expireResponse)
	if err != nil {
		return nil, err
	}

	return &expireResponse, nil
}

// ListUsers lists all users
func (c *Client) ListUsers(displayName *string, email *string, provider *string) (*models.ListUsersResponse, error) {
	url := "/api/v1/user"
	hasParams := false
	
	if displayName != nil {
		url += "?displayName=" + *displayName
		hasParams = true
	}
	if email != nil {
		if hasParams {
			url += "&"
		} else {
			url += "?"
		}
		url += "email=" + *email
		hasParams = true
	}
	if provider != nil {
		if hasParams {
			url += "&"
		} else {
			url += "?"
		}
		url += "provider=" + *provider
	}
	
	resp, err := c.doRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var usersResponse models.ListUsersResponse
	err = c.parseResponse(resp, &usersResponse)
	if err != nil {
		return nil, err
	}

	return &usersResponse, nil
}

// CreateUser creates a new user
func (c *Client) CreateUser(req *models.CreateUserRequest) (*models.CreateUserResponse, error) {
	resp, err := c.doRequest("POST", "/api/v1/user", req)
	if err != nil {
		return nil, err
	}

	var userResponse models.CreateUserResponse
	err = c.parseResponse(resp, &userResponse)
	if err != nil {
		return nil, err
	}

	return &userResponse, nil
}

// DeleteUser deletes a user
func (c *Client) DeleteUser(id int64) (*models.DeleteUserResponse, error) {
	resp, err := c.doRequest("DELETE", fmt.Sprintf("/api/v1/user/%d", id), nil)
	if err != nil {
		return nil, err
	}

	var deleteResponse models.DeleteUserResponse
	err = c.parseResponse(resp, &deleteResponse)
	if err != nil {
		return nil, err
	}

	return &deleteResponse, nil
}

// RenameUser renames a user
func (c *Client) RenameUser(oldId int64, newName string) (*models.RenameUserResponse, error) {
	resp, err := c.doRequest("POST", fmt.Sprintf("/api/v1/user/%d/rename/%s", oldId, newName), nil)
	if err != nil {
		return nil, err
	}

	var renameResponse models.RenameUserResponse
	err = c.parseResponse(resp, &renameResponse)
	if err != nil {
		return nil, err
	}

	return &renameResponse, nil
}

// DebugCreateNode creates a debug node
func (c *Client) DebugCreateNode(req *models.DebugCreateNodeRequest) (*models.DebugCreateNodeResponse, error) {
	resp, err := c.doRequest("POST", "/api/v1/debug/node", req)
	if err != nil {
		return nil, err
	}

	var debugNodeResponse models.DebugCreateNodeResponse
	err = c.parseResponse(resp, &debugNodeResponse)
	if err != nil {
		return nil, err
	}

	return &debugNodeResponse, nil
}