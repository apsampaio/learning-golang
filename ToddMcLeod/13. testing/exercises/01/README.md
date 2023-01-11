1. `cd dog`
2. `go test -bench .`
3. `cd ..`
4. `go test ./... -coverprofile=c.out`
5. `go tool cover -html=c.out -o coverage.html`
6. `godoc -http=:8080`
