SET CGO_ENABLED=0
SET GOOS=windows
SET GOARCH=386
go build -o hosts.exe -ldflags "-w -s" main.go