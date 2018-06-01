go test -coverprofile=.coverage ./components/awesome
go tool cover -html=.coverage -o coverage.html
