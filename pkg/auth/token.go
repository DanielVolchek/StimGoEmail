package auth

// generates auth token to use for receiving emails

import (
	"fmt"

	"github.com/google/uuid"
)

// generate token, write over previous token
func GenerateToken() uuid.UUID {
	uuid := uuid.New()

	fmt.Printf("token is\n%s\n", uuid)

	return uuid
}
