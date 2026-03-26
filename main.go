package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	// Generate a new UUID
	id := uuid.New()
	fmt.Printf("Generated UUID: %s\n", id.String())
}
