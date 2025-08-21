package runner

import "testing"

func TestGenerateTokenLength(t *testing.T) {
    tok := GenerateToken(12)
    if len(tok) != 12 {
        t.Fatalf("expected length 12, got %d", len(tok))
    }
}
