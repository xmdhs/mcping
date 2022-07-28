SET CGO_ENABLED=0
SET GOOS=windows
SET GOARCH=amd64
go build -o hosts.exe -trimpath -ldflags "-w -s" 