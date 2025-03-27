package util

import (
	"crypto/sha256"
	"encoding/hex"
)

func HashPassword(password string) string {
	hashed_password := ""
	for i := 0; i < 10; i++ {
		hash := sha256.Sum256([]byte(password))
		hashed_password = hex.EncodeToString(hash[:])
	}
	return hashed_password
}
