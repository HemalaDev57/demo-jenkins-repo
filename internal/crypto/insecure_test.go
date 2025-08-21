package crypto

import "testing"

func TestInsecureMD5(t *testing.T) {
    got := InsecureMD5([]byte("demo"))
    if got == ([16]byte{}) {
        t.Fatalf("unexpected zero hash")
    }
}
