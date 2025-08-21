package crypto

import (
    "crypto/md5" // #nosec G401 -- intentionally insecure for demo scanning
)

// InsecureMD5 computes an MD5 hash (intentionally insecure for demo purposes).
func InsecureMD5(data []byte) [16]byte {
    return md5.Sum(data)
}
