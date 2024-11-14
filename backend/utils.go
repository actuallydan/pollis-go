package main

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"math/rand"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/oklog/ulid/v2"
)

func GenerateULID() string {
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	id := ulid.MustNew(ulid.Timestamp(time.Now()), entropy).String()

	return id
}

func GetHashStringFromFile(getResult *s3.GetObjectOutput) (string, error) {
	hash := sha256.New()

	_, err := io.Copy(hash, getResult.Body)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}

// Helper function to check if a user has permission to access an organization
func UserHasOrgPermission(userID, orgID string) bool {
	var count int
	err := db.QueryRow(`
		SELECT COUNT(*) FROM user_organizations
		WHERE user_id = (SELECT id FROM users WHERE clerk_id = ?) AND organization_id = ?
	`, userID, orgID).Scan(&count)
	if err != nil {
		log.Printf("Error checking user organization permission: user `%s` doesn't have permission to access organization `%s` %v", userID, orgID, err)
		return false
	}
	return count > 0
}

// Helper function to get URL parameters
func GetURLParam(r *http.Request, param string) string {
	parts := strings.Split(r.URL.Path, "/")
	for i, part := range parts {
		if part == param && i+1 < len(parts) {
			return parts[i+1]
		}
	}
	return ""
}
