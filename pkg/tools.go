package pkg

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"time"
)

func GenerateRequestID() string {
	// Generate a random 16-byte array
	bytes := make([]byte, 16)
	_, err := rand.Read(bytes)
	if err != nil {
		// If there was an error generating random bytes, panic
		panic(err)
	}

	// Convert the byte array to a hex string
	hexString := hex.EncodeToString(bytes)

	// Insert hyphens at the appropriate positions
	return fmt.Sprintf("1-%s-%s-%s-%s", hexString[:8], hexString[8:12], hexString[12:16], hexString[16:])
}

func GenerateRequestTime() string {
	// Load the Asia/Shanghai timezone
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		// If there was an error loading the timezone, panic
		panic(err)
	}

	// Get the current time in the Asia/Shanghai timezone
	now := time.Now().In(loc)

	// Format the time as a string in the desired format
	return fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d.%03dZ", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), now.Nanosecond()/1000000)
}
