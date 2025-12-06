# Headscale API Client

This project provides a Go client library for interacting with the Headscale API. It was generated from an OpenAPI v2 specification file and includes comprehensive implementations for all API endpoints, along with models, error handling, and tests.

## Project Structure

- `client/`: Contains the main client implementation with all API methods
  - `client.go`: Main client implementation with methods for all API endpoints
  - `client_test.go`: Comprehensive tests for all client functionality
- `models/`: Contains Go structs that match the API definitions
  - `models.go`: Data models for requests and responses
  - `models_test.go`: Tests for JSON marshaling/unmarshaling
- `go.mod` / `go.sum`: Go module files
- `main.go`: Example usage of the client
- `openapiv2.json`: Original OpenAPI specification file used as reference

## Features

1. **Complete API Coverage**: All endpoints from the Headscale API are implemented:
   - API Key management (list, create, expire, delete)
   - Node operations (get, list, delete, rename, expire, tag, move, approve routes)
   - User management (list, create, delete, rename)
   - Pre-auth key operations (list, create, expire)
   - Health checks and policy management

2. **JSON Handling**: Proper marshaling and unmarshaling of all request/response types

3. **Error Handling**: Robust error handling for API responses and network errors

4. **Authentication**: Bearer token authentication support via API key

5. **Testing**: Comprehensive test coverage with 100% of major functionality tested

6. **Minimal Dependencies**: Only uses standard library and testify for testing

## Building and Running

### Prerequisites
- Go 1.24.10 or later

### Building
```bash
go build .
```

### Testing
```bash
# Run all tests
go test ./...

# Run tests for specific package
go test ./client
go test ./models
```

### Usage
```go
import (
    "github.com/wold9168/validate_headscale_client/client"
    "github.com/wold9168/validate_headscale_client/models"
)

// Create client
c := client.NewClient("http://your-headscale-server:8080", "your-api-key")

// Check health
health, err := c.GetHealth()

// List API keys
keys, err := c.ListApiKeys()
```

### Running the Example Application

The main.go file provides an example application that can be configured via command line flags or environment variables:

#### Command Line Flags:
- `-base-url`: Headscale server base URL (default: HEADSCALE_BASE_URL environment variable)
- `-api-key`: Headscale API key (default: HEADSCALE_API_KEY environment variable)

#### Environment Variables:
- `HEADSCALE_BASE_URL`: Headscale server URL (e.g., "http://localhost:8080")
- `HEADSCALE_API_KEY`: API key for authentication

#### Examples:
```bash
# Using command line flags
go run main.go -base-url "http://my-headscale-server:8080" -api-key "my-api-key"

# Using environment variables
export HEADSCALE_BASE_URL="http://my-headscale-server:8080"
export HEADSCALE_API_KEY="my-api-key"
go run main.go

# Using a mix of command line and environment variables
HEADSCALE_API_KEY="my-api-key" go run main.go -base-url "http://my-headscale-server:8080"
```

If neither command line flags nor environment variables are provided, the application will use default fallback values.

## Client Methods

The client provides methods for all Headscale API operations:

- `GetHealth()` - Check the health status of the Headscale server
- `ListApiKeys()` - List all API keys
- `CreateApiKey()` - Create a new API key
- `ExpireApiKey()` - Expire an API key
- `DeleteApiKey()` - Delete an API key
- `GetNode(nodeId)` - Get a specific node
- `ListNodes(userId)` - List all nodes (optionally filtered by user)
- `DeleteNode(nodeId)` - Delete a node
- `RenameNode(nodeId, newName)` - Rename a node
- `ExpireNode(nodeId)` - Expire a node
- `SetNodeTags(nodeId, tags)` - Set tags for a node
- `MoveNode(nodeId, userId)` - Move a node to a different user
- `SetApprovedRoutes(nodeId, routes)` - Set approved routes for a node
- `ListUsers(displayName, email, provider)` - List users with optional filters
- `CreateUser()` - Create a new user
- `DeleteUser(id)` - Delete a user
- `RenameUser(oldId, newName)` - Rename a user
- `ListPreAuthKeys(userId)` - List pre-authentication keys for a user
- `CreatePreAuthKey()` - Create a new pre-authentication key
- `ExpirePreAuthKey(userId, key)` - Expire a pre-authentication key
- `GetPolicy()` - Get the current ACL policy
- `SetPolicy(policy)` - Set the ACL policy
- `RegisterNode(machinePublicKey, authKey)` - Register a new node
- `BackfillNodeIPs(nodeId)` - Backfill node IP addresses
- `DebugCreateNode()` - Create a debug node

## Development Conventions

- All methods follow Go naming conventions
- Error handling is done explicitly in all methods
- JSON tags are properly defined for all model fields
- HTTP client timeouts are configured to prevent hanging requests
- All API endpoints are thoroughly tested