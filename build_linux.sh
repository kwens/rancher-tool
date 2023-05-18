go env -w GOARCH=amd64
go env -w GOOS=linux
go build -ldflags "-w -s" -o rancher-tool main.go