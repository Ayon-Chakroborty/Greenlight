package data

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base32"
	"time"
)

const (
	ScopeActivation = "activation"
)

type Token struct {
	Plaintext string
	Hash      []byte
	UserID    int64
	Expiry    time.Time
	Scope     string
}

func generateToken(userID int64, ttl time.Duration, scope string) (*Token, error){
	token := &Token{
		UserID: userID,
		Expiry: time.Now().Add(ttl),
		Scope: scope,
	}

	randomBytes := make([]byte, 16)

	_, err := rand.Read(randomBytes) // fill randomBytes with random bytes
	if err != nil{
		return nil, err
	}

	// endcode with no padding
	token.Plaintext = base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(randomBytes)

	// encrypt 
	hash := sha256.Sum256([]byte(token.Plaintext))
	token.Hash = hash[:] // copy hash to token.Hash

	return token, nil
}
