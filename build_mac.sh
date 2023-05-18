go env -w GOARCH=amd64
go env -w GOOS=darwin
go build -ldflags "-w -s" -o rancher-tool main.go
