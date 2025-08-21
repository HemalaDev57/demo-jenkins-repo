package mathutil

import "testing"

func TestAdd(t *testing.T) {
    if Add(5, 3) != 8 {
        t.Fatalf("expected 8")
    }
}

func TestSubtract(t *testing.T) {
    if Subtract(5, 3) != 2 {
        t.Fatalf("expected 2")
    }
}
