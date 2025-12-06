package models

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAPIKeyJSON(t *testing.T) {
	// Test APIKey marshaling and unmarshaling
	expiration := time.Now().Add(24 * time.Hour)
	apiKey := APIKey{
		ID:        1,
		Prefix:    "testprefix",
		Expiration: expiration,
		CreatedAt: time.Now(),
	}

	// Marshal to JSON
	data, err := json.Marshal(apiKey)
	assert.NoError(t, err)
	assert.Contains(t, string(data), "testprefix")

	// Unmarshal from JSON
	var parsedAPIKey APIKey
	err = json.Unmarshal(data, &parsedAPIKey)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), parsedAPIKey.ID)
	assert.Equal(t, "testprefix", parsedAPIKey.Prefix)
}

func TestNodeJSON(t *testing.T) {
	// Test Node marshaling and unmarshaling
	node := Node{
		ID:         1,
		Name:       "test-node",
		IPAddresses: []string{"10.0.0.1", "10.0.0.2"},
		Online:     true,
	}

	// Marshal to JSON
	data, err := json.Marshal(node)
	assert.NoError(t, err)
	assert.Contains(t, string(data), "test-node")

	// Unmarshal from JSON
	var parsedNode Node
	err = json.Unmarshal(data, &parsedNode)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), parsedNode.ID)
	assert.Equal(t, "test-node", parsedNode.Name)
	assert.Equal(t, []string{"10.0.0.1", "10.0.0.2"}, parsedNode.IPAddresses)
	assert.True(t, parsedNode.Online)
}

func TestUserJSON(t *testing.T) {
	// Test User marshaling and unmarshaling
	user := User{
		ID:          1,
		Name:        "test-user",
		DisplayName: "Test User",
	}

	// Marshal to JSON
	data, err := json.Marshal(user)
	assert.NoError(t, err)
	assert.Contains(t, string(data), "test-user")

	// Unmarshal from JSON
	var parsedUser User
	err = json.Unmarshal(data, &parsedUser)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), parsedUser.ID)
	assert.Equal(t, "test-user", parsedUser.Name)
	assert.Equal(t, "Test User", parsedUser.DisplayName)
}

func TestPreAuthKeyJSON(t *testing.T) {
	// Test PreAuthKey marshaling and unmarshaling
	user := User{
		ID:   1,
		Name: "test-user",
	}
	
	preAuthKey := PreAuthKey{
		ID:       1,
		Key:      "test-key",
		User:     user,
		Reusable: true,
		ACLTags:  []string{"tag1", "tag2"},
	}

	// Marshal to JSON
	data, err := json.Marshal(preAuthKey)
	assert.NoError(t, err)
	assert.Contains(t, string(data), "test-key")

	// Unmarshal from JSON
	var parsedPreAuthKey PreAuthKey
	err = json.Unmarshal(data, &parsedPreAuthKey)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), parsedPreAuthKey.ID)
	assert.Equal(t, "test-key", parsedPreAuthKey.Key)
	assert.Equal(t, true, parsedPreAuthKey.Reusable)
	assert.Equal(t, []string{"tag1", "tag2"}, parsedPreAuthKey.ACLTags)
}

func TestCreateApiKeyRequestJSON(t *testing.T) {
	// Test CreateApiKeyRequest marshaling and unmarshaling
	expiration := time.Now().Add(24 * time.Hour)
	req := CreateApiKeyRequest{
		Expiration: &expiration,
	}

	// Marshal to JSON
	data, err := json.Marshal(req)
	assert.NoError(t, err)
	assert.Contains(t, string(data), "expiration")

	// Unmarshal from JSON
	var parsedReq CreateApiKeyRequest
	err = json.Unmarshal(data, &parsedReq)
	assert.NoError(t, err)
	assert.NotNil(t, parsedReq.Expiration)
}

func TestSetTagsBodyJSON(t *testing.T) {
	// Test SetTagsBody marshaling and unmarshaling
	req := SetTagsBody{
		Tags: []string{"tag1", "tag2"},
	}

	// Marshal to JSON
	data, err := json.Marshal(req)
	assert.NoError(t, err)
	assert.Contains(t, string(data), "tag1")

	// Unmarshal from JSON
	var parsedReq SetTagsBody
	err = json.Unmarshal(data, &parsedReq)
	assert.NoError(t, err)
	assert.Equal(t, []string{"tag1", "tag2"}, parsedReq.Tags)
}

func TestCreatePreAuthKeyRequestJSON(t *testing.T) {
	// Test CreatePreAuthKeyRequest marshaling and unmarshaling
	req := CreatePreAuthKeyRequest{
		User:      "test-user",
		Reusable:  true,
		Ephemeral: false,
		ACLTags:   []string{"tag1", "tag2"},
	}

	// Marshal to JSON
	data, err := json.Marshal(req)
	assert.NoError(t, err)
	assert.Contains(t, string(data), "test-user")

	// Unmarshal from JSON
	var parsedReq CreatePreAuthKeyRequest
	err = json.Unmarshal(data, &parsedReq)
	assert.NoError(t, err)
	assert.Equal(t, "test-user", parsedReq.User)
	assert.Equal(t, true, parsedReq.Reusable)
	assert.Equal(t, false, parsedReq.Ephemeral)
	assert.Equal(t, []string{"tag1", "tag2"}, parsedReq.ACLTags)
}

func TestHeadscaleServiceMoveNodeBodyJSON(t *testing.T) {
	// Test HeadscaleServiceMoveNodeBody marshaling and unmarshaling
	req := HeadscaleServiceMoveNodeBody{
		User: "test-user",
	}

	// Marshal to JSON
	data, err := json.Marshal(req)
	assert.NoError(t, err)
	assert.Contains(t, string(data), "test-user")

	// Unmarshal from JSON
	var parsedReq HeadscaleServiceMoveNodeBody
	err = json.Unmarshal(data, &parsedReq)
	assert.NoError(t, err)
	assert.Equal(t, "test-user", parsedReq.User)
}