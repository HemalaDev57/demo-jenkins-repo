package main

import (
    "fmt"
    "demo-jenkins-repo/internal/mathutil"
    "demo-jenkins-repo/internal/crypto"
    "demo-jenkins-repo/internal/runner"
)

func main() {
    a, b := 5, 3
    fmt.Printf("Add(%d,%d) = %d\n", a, b, mathutil.Add(a, b))
    fmt.Printf("MD5 of 'demo' (insecure) = %x\n", crypto.InsecureMD5([]byte("demo")))
    fmt.Printf("Token (insecure RNG) = %s\n", runner.GenerateToken(8))
}
