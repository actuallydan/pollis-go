package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/clerk/clerk-sdk-go/v2"
	"github.com/clerk/clerk-sdk-go/v2/user"
)

// GetUserFromContext retrieves the Clerk user from the request context
func GetUserFromContext(r *http.Request) (*clerk.User, error) {
	claims, ok := clerk.SessionClaimsFromContext(r.Context())
	if !ok {
		return nil, fmt.Errorf("session claims not found in context")
	}

	usr, err := user.Get(r.Context(), claims.Subject)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve user: %v", err)
	}

	return usr, nil
}

func initClerk() {
	clerk.SetKey(os.Getenv("CLERK_SECRET_KEY"))
}
