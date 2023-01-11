1. `go test ./...`
2. `cd dog`
3. `go test -bench .`
4. `cd ..`
5. `go test ./... -coverprofile=c.out`
6. `go tool cover -html=c.out -o coverage.html`
7. `godoc -http=:8080`
