package main

import "testing"

func TestAdd(t *testing.T) {
    result := Add(5, 3)
    if result != 8 {
        t.Errorf("Expected 8, got %d", result)
    }
}

func TestSubtract(t *testing.T) {
    result := Subtract(5, 3)
    if result != 2 { // Intentional failing test to simulate issue
        t.Errorf("Expected 2, got %d", result)
    }
}
