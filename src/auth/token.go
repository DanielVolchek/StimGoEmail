package auth

// generates auth token to use for receiving emails

import "github.com/google/uuid"

func GenerateToken(envArgs) uuid.UUID {
	uuid := uuid.New()

	//

	return uuid
}
