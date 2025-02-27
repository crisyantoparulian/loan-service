package str

import (
	"math/rand"
	"time"
)

const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// GenerateRandomString generates a random string of given length using a local random generator.
func GenerateRandomString(length int) string {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src) // Use a new local random generator

	result := make([]byte, length)
	for i := range result {
		result[i] = charset[r.Intn(len(charset))]
	}
	return string(result)
}
