package shortid

import (
	"crypto/rand"
	"encoding/base64"
)

// URL-safe, tanpa padding, panjang n (≈ n*6 bit entropi)
func Generate(n int) string {
	b := make([]byte, n)
	_, _ = rand.Read(b)
	// RawURLEncoding menghindari karakter yang tidak aman di URL
	return base64.RawURLEncoding.EncodeToString(b)[:n]
}
