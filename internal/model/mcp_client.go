// Package model provides data models for MCPJungle.
package model

import (
	"encoding/json"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// McpClient represents MCP clients and their access to the MCP Servers provided MCPJungle MCP server
type McpClient struct {
	gorm.Model

	Name        string `gorm:"uniqueIndex;not null" json:"name"`
	Description string `json:"description"`

	AccessToken string `gorm:"unique; not null" json:"access_token"`

	// AllowList contains a list of MCP Server names that this client is allowed to view and call
	// storing the list of server names as a JSON array is a convenient way for now.
	// In the future, this will be removed in favor of a separate table for ACLs.
	AllowList datatypes.JSON `gorm:"type:jsonb; not null" json:"allow_list"`
}

// CheckHasServerAccess returns true if this client has access to the specified MCP server.
// If not, it returns false.
func (c *McpClient) CheckHasServerAccess(serverName string) bool {
	if c.AllowList == nil {
		return false
	}

	var allowedServers []string

	err := json.Unmarshal(c.AllowList, &allowedServers)
	if err != nil {
		return false
	}

	for _, allowed := range allowedServers {
		if allowed == serverName {
			return true
		}
	}

	return false
}
