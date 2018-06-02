```
go test -coverprofile=.coverage -v ./...
go tool cover -html=.coverage -o coverage.html
```