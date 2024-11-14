package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type UserCreatedEvent struct {
	Data   UserCreatedEventData `json:"data"`
	Object string               `json:"object"`
	Type   string               `json:"type"`
}

type UserCreatedEventData struct {
	Birthday              string                         `json:"birthday"`
	CreatedAt             int64                          `json:"created_at"`
	EmailAddresses        []UserCreatedEventEmailAddress `json:"email_addresses"`
	ExternalAccounts      []interface{}                  `json:"external_accounts"`
	ExternalID            string                         `json:"external_id"`
	FirstName             string                         `json:"first_name"`
	Gender                string                         `json:"gender"`
	ID                    string                         `json:"id"`
	ImageURL              string                         `json:"image_url"`
	LastName              string                         `json:"last_name"`
	LastSignInAt          int64                          `json:"last_sign_in_at"`
	Object                string                         `json:"object"`
	PasswordEnabled       bool                           `json:"password_enabled"`
	PhoneNumbers          []interface{}                  `json:"phone_numbers"`
	PrimaryEmailAddressID string                         `json:"primary_email_address_id"`
	PrimaryPhoneNumberID  interface{}                    `json:"primary_phone_number_id"`
	PrimaryWeb3WalletID   interface{}                    `json:"primary_web3_wallet_id"`
	PrivateMetadata       UserCreatedEventMetadata       `json:"private_metadata"`
	ProfileImageURL       string                         `json:"profile_image_url"`
	PublicMetadata        UserCreatedEventMetadata       `json:"public_metadata"`
	TwoFactorEnabled      bool                           `json:"two_factor_enabled"`
	UnsafeMetadata        UserCreatedEventMetadata       `json:"unsafe_metadata"`
	UpdatedAt             int64                          `json:"updated_at"`
	Username              interface{}                    `json:"username"`
	Web3Wallets           []interface{}                  `json:"web3_wallets"`
}

type UserCreatedEventEmailAddress struct {
	EmailAddress string                       `json:"email_address"`
	ID           string                       `json:"id"`
	LinkedTo     []interface{}                `json:"linked_to"`
	Object       string                       `json:"object"`
	Verification UserCreatedEventVerification `json:"verification"`
}

type UserCreatedEventVerification struct {
	Status   string `json:"status"`
	Strategy string `json:"strategy"`
}

type UserCreatedEventMetadata struct {
}

func handleSignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var event UserCreatedEvent
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		http.Error(w, "Error parsing JSON body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	data := event.Data

	// Extract the primary email address
	var primaryEmail string
	for _, email := range data.EmailAddresses {
		if email.ID == data.PrimaryEmailAddressID {
			primaryEmail = email.EmailAddress
			break
		}
	}

	// Generate a new ULID for the user
	userID := GenerateULID()

	// Prepare the SQL statement
	stmt, err := db.Prepare(`
		INSERT INTO users (id, clerk_id, email, created_at)
		VALUES (?, ?, ?, ?)
	`)
	if err != nil {
		log.Printf("Error preparing SQL statement: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	// Execute the SQL statement
	_, err = stmt.Exec(
		userID,
		data.ID,
		primaryEmail,
		time.Unix(data.CreatedAt, 0),
	)
	if err != nil {
		log.Printf("Error executing SQL statement: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User created successfully"))
}
