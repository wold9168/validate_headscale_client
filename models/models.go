package models

import (
	"time"
)

// APIKey represents the API key structure
type APIKey struct {
	ID        string    `json:"id"`
	Prefix    string    `json:"prefix"`
	Expiration time.Time `json:"expiration"`
	CreatedAt time.Time `json:"createdAt"`
	LastSeen  *time.Time `json:"lastSeen,omitempty"`
}

// BackfillNodeIPsResponse represents the response for backfilling node IPs
type BackfillNodeIPsResponse struct {
	Changes []string `json:"changes"`
}

// CreateApiKeyRequest represents the request to create an API key
type CreateApiKeyRequest struct {
	Expiration *time.Time `json:"expiration,omitempty"`
}

// CreateApiKeyResponse represents the response for creating an API key
type CreateApiKeyResponse struct {
	APIKey string `json:"apiKey"`
}

// PreAuthKey represents a pre-authentication key
type PreAuthKey struct {
	User        User       `json:"user"`
	ID          string     `json:"id"`
	Key         string     `json:"key"`
	Reusable    bool       `json:"reusable"`
	Ephemeral   bool       `json:"ephemeral"`
	Used        bool       `json:"used"`
	Expiration  *time.Time `json:"expiration,omitempty"`
	CreatedAt   time.Time  `json:"createdAt"`
	ACLTags     []string   `json:"aclTags"`
}

// CreatePreAuthKeyRequest represents the request to create a pre-auth key
type CreatePreAuthKeyRequest struct {
	User        string    `json:"user"`
	Reusable    bool      `json:"reusable"`
	Ephemeral   bool      `json:"ephemeral"`
	Expiration  *time.Time `json:"expiration,omitempty"`
	ACLTags     []string  `json:"aclTags"`
}

// CreatePreAuthKeyResponse represents the response for creating a pre-auth key
type CreatePreAuthKeyResponse struct {
	PreAuthKey PreAuthKey `json:"preAuthKey"`
}

// User represents a user in the system
type User struct {
	ID            string    `json:"id"`
	Name          string    `json:"name"`
	CreatedAt     time.Time `json:"createdAt"`
	DisplayName   string    `json:"displayName,omitempty"`
	Email         string    `json:"email,omitempty"`
	ProviderID    string    `json:"providerId,omitempty"`
	Provider      string    `json:"provider,omitempty"`
	ProfilePicURL string    `json:"profilePicUrl,omitempty"`
}

// CreateUserRequest represents the request to create a user
type CreateUserRequest struct {
	Name        string `json:"name"`
	DisplayName string `json:"displayName,omitempty"`
	Email       string `json:"email,omitempty"`
	PictureURL  string `json:"pictureUrl,omitempty"`
}

// CreateUserResponse represents the response for creating a user
type CreateUserResponse struct {
	User User `json:"user"`
}

// DebugCreateNodeRequest represents the request to create a debug node
type DebugCreateNodeRequest struct {
	User  string   `json:"user"`
	Key   string   `json:"key"`
	Name  string   `json:"name"`
	Routes []string `json:"routes"`
}

// Node represents a node in the system
type Node struct {
	ID                string    `json:"id"`
	MachineKey        string    `json:"machineKey"`
	NodeKey           string    `json:"nodeKey"`
	DiscoKey          string    `json:"discoKey"`
	IPAddresses       []string  `json:"ipAddresses"`
	Name              string    `json:"name"`
	User              User      `json:"user"`
	LastSeen          time.Time `json:"lastSeen"`
	Expiry            time.Time `json:"expiry"`
	PreAuthKey        *PreAuthKey `json:"preAuthKey,omitempty"`
	CreatedAt         time.Time `json:"createdAt"`
	RegisterMethod    string    `json:"registerMethod"`
	ForcedTags        []string  `json:"forcedTags"`
	InvalidTags       []string  `json:"invalidTags"`
	ValidTags         []string  `json:"validTags"`
	GivenName         string    `json:"givenName"`
	Online            bool      `json:"online"`
	ApprovedRoutes    []string  `json:"approvedRoutes"`
	AvailableRoutes   []string  `json:"availableRoutes"`
	SubnetRoutes      []string  `json:"subnetRoutes"`
}

// DebugCreateNodeResponse represents the response for creating a debug node
type DebugCreateNodeResponse struct {
	Node Node `json:"node"`
}

// DeleteApiKeyResponse represents the response for deleting an API key
type DeleteApiKeyResponse struct{}

// DeleteNodeResponse represents the response for deleting a node
type DeleteNodeResponse struct{}

// DeleteUserResponse represents the response for deleting a user
type DeleteUserResponse struct{}

// ExpireApiKeyRequest represents the request to expire an API key
type ExpireApiKeyRequest struct {
	Prefix string `json:"prefix"`
}

// ExpireApiKeyResponse represents the response for expiring an API key
type ExpireApiKeyResponse struct{}

// ExpireNodeResponse represents the response for expiring a node
type ExpireNodeResponse struct {
	Node Node `json:"node"`
}

// ExpirePreAuthKeyRequest represents the request to expire a pre-auth key
type ExpirePreAuthKeyRequest struct {
	User string `json:"user"`
	Key  string `json:"key"`
}

// ExpirePreAuthKeyResponse represents the response for expiring a pre-auth key
type ExpirePreAuthKeyResponse struct{}

// GetNodeResponse represents the response for getting a node
type GetNodeResponse struct {
	Node Node `json:"node"`
}

// GetPolicyResponse represents the response for getting policy
type GetPolicyResponse struct {
	Policy    string    `json:"policy"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// HealthResponse represents the health check response
type HealthResponse struct {
	DatabaseConnectivity bool `json:"databaseConnectivity"`
}

// ListApiKeysResponse represents the response for listing API keys
type ListApiKeysResponse struct {
	APIKeys []APIKey `json:"apiKeys"`
}

// ListNodesResponse represents the response for listing nodes
type ListNodesResponse struct {
	Nodes []Node `json:"nodes"`
}

// ListPreAuthKeysResponse represents the response for listing pre-auth keys
type ListPreAuthKeysResponse struct {
	PreAuthKeys []PreAuthKey `json:"preAuthKeys"`
}

// ListUsersResponse represents the response for listing users
type ListUsersResponse struct {
	Users []User `json:"users"`
}

// MoveNodeResponse represents the response for moving a node
type MoveNodeResponse struct {
	Node Node `json:"node"`
}

// RegisterNodeResponse represents the response for registering a node
type RegisterNodeResponse struct {
	Node Node `json:"node"`
}

// RenameNodeResponse represents the response for renaming a node
type RenameNodeResponse struct {
	Node Node `json:"node"`
}

// RenameUserResponse represents the response for renaming a user
type RenameUserResponse struct {
	User User `json:"user"`
}

// SetApprovedRoutesResponse represents the response for setting approved routes
type SetApprovedRoutesResponse struct {
	Node Node `json:"node"`
}

// SetApprovedRoutesBody represents the body for setting approved routes
type SetApprovedRoutesBody struct {
	Routes []string `json:"routes"`
}

// SetPolicyRequest represents the request to set policy
type SetPolicyRequest struct {
	Policy string `json:"policy"`
}

// SetPolicyResponse represents the response for setting policy
type SetPolicyResponse struct {
	Policy    string    `json:"policy"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// SetTagsResponse represents the response for setting tags
type SetTagsResponse struct {
	Node Node `json:"node"`
}

// SetTagsBody represents the body for setting tags
type SetTagsBody struct {
	Tags []string `json:"tags"`
}

// RegisterNodeRequest represents the request to register a node
type RegisterNodeRequest struct {
	MachinePublicKey string `json:"machinePublicKey"`
	AuthKey          string `json:"authKey"`
}

// HeadscaleServiceMoveNodeBody represents the body for moving a node
type HeadscaleServiceMoveNodeBody struct {
	User string `json:"user"`
}