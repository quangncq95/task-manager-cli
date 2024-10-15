package utils

import (
	"crypto/rand"
	"fmt"
	"time"
)

func GenerateTimestampBasedID() string {
    timestamp := time.Now().UnixNano()

    randomBytes := make([]byte, 8)
    _, _ = rand.Read(randomBytes)

    return fmt.Sprintf("%x-%x", timestamp, randomBytes)
}