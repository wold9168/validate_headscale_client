package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wold9168/validate_headscale_client/models"
)

func TestNewClient(t *testing.T) {
	client := NewClient("http://localhost:8080", "test-api-key")
	
	assert.Equal(t, "http://localhost:8080", client.BaseURL)
	assert.Equal(t, "test-api-key", client.APIKey)
	assert.NotNil(t, client.HTTPClient)
}

func TestGetHealth(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v1/health", r.URL.Path)
		assert.Equal(t, "GET", r.Method)
		
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		response := `{"databaseConnectivity": true}`
		w.Write([]byte(response))
	}))
	defer server.Close()

	// Create client with test server URL
	client := NewClient(server.URL, "")

	// Call the method
	health, err := client.GetHealth()
	
	assert.NoError(t, err)
	assert.NotNil(t, health)
	assert.Equal(t, true, health.DatabaseConnectivity)
}

func TestListApiKeys(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v1/apikey", r.URL.Path)
		assert.Equal(t, "GET", r.Method)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		response := `{"apiKeys": [{"id": "1", "prefix": "testprefix"}]}`
		w.Write([]byte(response))
	}))
	defer server.Close()

	// Create client with test server URL
	client := NewClient(server.URL, "")

	// Call the method
	apiKeys, err := client.ListApiKeys()

	assert.NoError(t, err)
	assert.NotNil(t, apiKeys)
	assert.Len(t, apiKeys.APIKeys, 1)
	assert.Equal(t, "1", apiKeys.APIKeys[0].ID)
	assert.Equal(t, "testprefix", apiKeys.APIKeys[0].Prefix)
}

func TestCreateApiKey(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v1/apikey", r.URL.Path)
		assert.Equal(t, "POST", r.Method)
		
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		response := `{"apiKey": "generated-api-key"}`
		w.Write([]byte(response))
	}))
	defer server.Close()

	// Create client with test server URL
	client := NewClient(server.URL, "")

	// Call the method
	req := &models.CreateApiKeyRequest{}
	apiKey, err := client.CreateApiKey(req)
	
	assert.NoError(t, err)
	assert.NotNil(t, apiKey)
	assert.Equal(t, "generated-api-key", apiKey.APIKey)
}

func TestExpireApiKey(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v1/apikey/expire", r.URL.Path)
		assert.Equal(t, "POST", r.Method)
		
		// Parse request body
		var req models.ExpireApiKeyRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		assert.NoError(t, err)
		assert.Equal(t, "test-prefix", req.Prefix)
		
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		response := `{}`
		w.Write([]byte(response))
	}))
	defer server.Close()

	// Create client with test server URL
	client := NewClient(server.URL, "")

	// Call the method
	expireResp, err := client.ExpireApiKey("test-prefix")
	
	assert.NoError(t, err)
	assert.NotNil(t, expireResp)
}

func TestDeleteApiKey(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v1/apikey/test-prefix", r.URL.Path)
		assert.Equal(t, "DELETE", r.Method)
		
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		response := `{}`
		w.Write([]byte(response))
	}))
	defer server.Close()

	// Create client with test server URL
	client := NewClient(server.URL, "")

	// Call the method
	deleteResp, err := client.DeleteApiKey("test-prefix")
	
	assert.NoError(t, err)
	assert.NotNil(t, deleteResp)
}

func TestGetNode(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v1/node/123", r.URL.Path)
		assert.Equal(t, "GET", r.Method)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		response := `{"node": {"id": "123", "name": "test-node"}}`
		w.Write([]byte(response))
	}))
	defer server.Close()

	// Create client with test server URL
	client := NewClient(server.URL, "")

	// Call the method
	node, err := client.GetNode("123")

	assert.NoError(t, err)
	assert.NotNil(t, node)
	assert.Equal(t, "123", node.Node.ID)
	assert.Equal(t, "test-node", node.Node.Name)
}

func TestListNodes(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v1/node", r.URL.Path)
		assert.Equal(t, "GET", r.Method)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		response := `{"nodes": [{"id": "1", "name": "node1"}, {"id": "2", "name": "node2"}]}`
		w.Write([]byte(response))
	}))
	defer server.Close()

	// Create client with test server URL
	client := NewClient(server.URL, "")

	// Call the method
	nodes, err := client.ListNodes(nil)

	assert.NoError(t, err)
	assert.NotNil(t, nodes)
	assert.Len(t, nodes.Nodes, 2)
	assert.Equal(t, "1", nodes.Nodes[0].ID)
	assert.Equal(t, "node1", nodes.Nodes[0].Name)
	assert.Equal(t, "2", nodes.Nodes[1].ID)
	assert.Equal(t, "node2", nodes.Nodes[1].Name)
}

func TestDeleteNode(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v1/node/123", r.URL.Path)
		assert.Equal(t, "DELETE", r.Method)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		response := `{}`
		w.Write([]byte(response))
	}))
	defer server.Close()

	// Create client with test server URL
	client := NewClient(server.URL, "")

	// Call the method
	deleteResp, err := client.DeleteNode("123")

	assert.NoError(t, err)
	assert.NotNil(t, deleteResp)
}

func TestCreateUser(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v1/user", r.URL.Path)
		assert.Equal(t, "POST", r.Method)
		
		// Parse request body
		var req models.CreateUserRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		assert.NoError(t, err)
		assert.Equal(t, "testuser", req.Name)
		
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		response := `{"user": {"id": "1", "name": "testuser"}}`
		w.Write([]byte(response))
	}))
	defer server.Close()

	// Create client with test server URL
	client := NewClient(server.URL, "")

	// Call the method
	req := &models.CreateUserRequest{
		Name: "testuser",
	}
	user, err := client.CreateUser(req)

	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "1", user.User.ID)
	assert.Equal(t, "testuser", user.User.Name)
}

func TestListUsers(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v1/user", r.URL.Path)
		assert.Equal(t, "GET", r.Method)
		
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		response := `{"users": [{"id": "1", "name": "user1"}, {"id": "2", "name": "user2"}]}`
		w.Write([]byte(response))
	}))
	defer server.Close()

	// Create client with test server URL
	client := NewClient(server.URL, "")

	// Call the method
	users, err := client.ListUsers(nil, nil, nil)

	assert.NoError(t, err)
	assert.NotNil(t, users)
	assert.Len(t, users.Users, 2)
	assert.Equal(t, "1", users.Users[0].ID)
	assert.Equal(t, "user1", users.Users[0].Name)
	assert.Equal(t, "2", users.Users[1].ID)
	assert.Equal(t, "user2", users.Users[1].Name)
}

func TestDoRequestError(t *testing.T) {
	// Test error case when creating request
	client := NewClient("invalid://url", "")
	
	_, err := client.doRequest("GET", "invalid-path", nil)
	assert.Error(t, err)
}

func TestParseResponseError(t *testing.T) {
	// Create a test server that returns an error status
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
	}))
	defer server.Close()

	client := NewClient(server.URL, "")
	
	resp, err := client.doRequest("GET", "/api/v1/health", nil)
	assert.NoError(t, err)
	
	err = client.parseResponse(resp, nil)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "API request failed with status 500")
}

func TestParseResponseJSONError(t *testing.T) {
	// Create a test server that returns invalid JSON
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("{invalid json}"))
	}))
	defer server.Close()

	client := NewClient(server.URL, "")
	
	resp, err := client.doRequest("GET", "/api/v1/health", nil)
	assert.NoError(t, err)
	
	var health models.HealthResponse
	err = client.parseResponse(resp, &health)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to unmarshal response")
}

func TestAPIKeyAuth(t *testing.T) {
	// Create a test server to check if API key is sent correctly
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		assert.Equal(t, "Bearer test-api-key", authHeader)
		
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		response := `{"databaseConnectivity": true}`
		w.Write([]byte(response))
	}))
	defer server.Close()

	client := NewClient(server.URL, "test-api-key")
	
	health, err := client.GetHealth()
	assert.NoError(t, err)
	assert.NotNil(t, health)
}

func TestListNodesWithUserId(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Contains(t, r.URL.Path, "/api/v1/node")
		assert.Contains(t, r.URL.RawQuery, "user=testuser")
		assert.Equal(t, "GET", r.Method)
		
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		response := `{"nodes": []}`
		w.Write([]byte(response))
	}))
	defer server.Close()

	// Create client with test server URL
	client := NewClient(server.URL, "")

	// Call the method
	userId := "testuser"
	nodes, err := client.ListNodes(&userId)
	
	assert.NoError(t, err)
	assert.NotNil(t, nodes)
}

func TestRenameNode(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, fmt.Sprintf("/api/v1/node/%s/rename/newname", "123"), r.URL.Path)
		assert.Equal(t, "POST", r.Method)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		response := `{"node": {"id": "123", "name": "newname"}}`
		w.Write([]byte(response))
	}))
	defer server.Close()

	// Create client with test server URL
	client := NewClient(server.URL, "")

	// Call the method
	resp, err := client.RenameNode("123", "newname")

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "123", resp.Node.ID)
	assert.Equal(t, "newname", resp.Node.Name)
}