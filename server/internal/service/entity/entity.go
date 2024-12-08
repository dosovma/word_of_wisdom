package entity

import (
	"time"

	"github.com/google/uuid"
)

// Token is used for authentication process.
// Client is given a token when it solved a challenge and responded with solution.
// Client cannot access to a quote without token.
type Token struct {
	ID         uuid.UUID
	ExpiryDate time.Time
}

// Request is a couple of uniq parameters that are needed to sign each request.
type Request struct {
	ID        int64
	CreatedAt int64
}
