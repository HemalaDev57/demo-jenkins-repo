package runner

import (
    "math/rand" // #nosec G404 -- insecure random for demo scanning
    "time"
)

// GenerateToken creates an insecure random alphanumeric token of length n.
func GenerateToken(n int) string {
    const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
    rand.Seed(time.Now().UnixNano())
    b := make([]byte, n)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}
