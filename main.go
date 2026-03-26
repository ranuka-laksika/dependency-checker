package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/modelcontextprotocol/go-sdk/pkg/protocol"
)

func main() {
	// Generate a new UUID
	id := uuid.New()
	fmt.Printf("Generated UUID: %s\n", id.String())

	// Example MCP SDK usage
	version := protocol.LATEST_PROTOCOL_VERSION
	fmt.Printf("MCP Protocol Version: %s\n", version)
}
