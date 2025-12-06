package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/wold9168/validate_headscale_client/client"
	"github.com/wold9168/validate_headscale_client/models"
)

func main() {
	// Define command line flags
	baseURL := flag.String("base-url", "", "Headscale server base URL (default: HEADSCALE_BASE_URL environment variable)")
	apiKey := flag.String("api-key", "", "Headscale API key (default: HEADSCALE_API_KEY environment variable)")

	// Parse command line flags
	flag.Parse()

	// Use environment variables if flags are not provided
	if *baseURL == "" {
		if envBaseURL := os.Getenv("HEADSCALE_BASE_URL"); envBaseURL != "" {
			*baseURL = envBaseURL
		} else {
			*baseURL = "http://localhost:8080" // default fallback
		}
	}

	if *apiKey == "" {
		if envAPIKey := os.Getenv("HEADSCALE_API_KEY"); envAPIKey != "" {
			*apiKey = envAPIKey
		} else {
			*apiKey = "your-api-key" // default fallback
		}
	}

	// Create a new client instance
	c := client.NewClient(*baseURL, *apiKey)

	fmt.Printf("Using Headscale server: %s\n", *baseURL)

	// Example: Check health
	health, err := c.GetHealth()
	if err != nil {
		log.Printf("Error getting health: %v", err)
	} else {
		fmt.Printf("Health check result: %+v\n", health)
	}

	// Example: List API keys
	apiKeys, err := c.ListApiKeys()
	if err != nil {
		log.Printf("Error listing API keys: %v", err)
	} else {
		fmt.Printf("Found %d API keys\n", len(apiKeys.APIKeys))
	}

	// Example: Create a new API key
	createReq := &models.CreateApiKeyRequest{}
	newKey, err := c.CreateApiKey(createReq)
	if err != nil {
		log.Printf("Error creating API key: %v", err)
	} else {
		fmt.Printf("Created new API key: %s\n", newKey.APIKey)
	}

	// Example: List users
	users, err := c.ListUsers(nil, nil, nil)
	if err != nil {
		log.Printf("Error listing users: %v", err)
	} else {
		fmt.Printf("Found %d users\n", len(users.Users))

		// If there are users, try to get the first one's nodes
		if len(users.Users) > 0 {
			nodes, err := c.ListNodes(&users.Users[0].Name)
			if err != nil {
				log.Printf("Error listing nodes for user %s: %v", users.Users[0].Name, err)
			} else {
				fmt.Printf("Found %d nodes for user %s\n", len(nodes.Nodes), users.Users[0].Name)
			}

			// Example: List pre-authentication keys for the first user if available
			// Note: Using the user's Name field instead of ID since that's what the API expects
			// The API endpoints that need user ID typically expect the user ID as a string
			preAuthKeys, err := c.ListPreAuthKeys(users.Users[0].Name)
			if err != nil {
				log.Printf("Error listing pre-auth keys: %v", err)
			} else {
				fmt.Printf("Found %d pre-auth keys for user %s\n", len(preAuthKeys.PreAuthKeys), users.Users[0].Name)
			}
		}
	}

	fmt.Println("Client operations completed successfully!")
}