build:
	go build -o bin/lockerr-linux -ldflags '-w -s' -a cmd/lockerr/main.go
	upx --lzma bin/lockerr-linux

test:
	go test

compile:
	# Cross compile for every major platform.
	GOOS=linux GOARCH=amd64 go build -o bin/lockerr-linux -ldflags '-w -s' -a cmd/lockerr/main.go
	GOOS=darwin GOARCH=amd64 go build -o bin/lockerr-macos -ldflags '-w -s' -a cmd/lockerr/main.go
	GOOS=windows GOARCH=amd64 go build -o bin/lockerr-windows.exe -ldflags '-w -s' -a cmd/lockerr/main.go

compress:
	upx --lzma bin/lockerr-windows.exe
	upx --lzma bin/lockerr-macos
	upx --lzma bin/lockerr-linux

	# Test the compressed binaries.
	upx -t bin/*

release: compile compress

run:
	go run cmd/lockerr/main.go