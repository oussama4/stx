package crypto

import (
	"crypto/rand"
	"strings"
)

func tokenBytes(size int) []byte {
	bytes := make([]byte, size)
	rand.Read(bytes)

	return bytes
}

// Token generates a securely generated random token
func Token(size int) string {
	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	bytes := tokenBytes(size)

	var b strings.Builder
	for i := 0; i < size; i++ {
		b.WriteByte(alphabet[bytes[i]%61])
	}

	return b.String()
}
