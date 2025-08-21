# Demo Go Jenkins Repo

Minimal Go project for Jenkins with build, tests, artifact publishing, and security scan.

## What it Shows
- Go build producing `dist/app`
- Unit tests with JUnit XML via `gotestsum`
- Security scan via `gosec` (intentionally includes minor issues: MD5 usage and insecure RNG)
- Jenkinsfile to run all stages and publish artifacts and test reports

## Local Run
```bash
go mod tidy
go build -o dist/app ./cmd/app
go test ./...
```
