build:
	go build -o bin/lockerr-linux -ldflags '-w -s' -a cmd/lockerr/main.go

test:
	go test

compile:
	# Cross compile for every major platform.
	GOOS=linux GOARCH=amd64 go build -o bin/lockerr-linux -ldflags '-w -s' -a cmd/lockerr/main.go
	GOOS=darwin GOARCH=amd64 go build -o bin/lockerr-macos -ldflags '-w -s' -a cmd/lockerr/main.go
	GOOS=windows GOARCH=amd64 go build -o bin/lockerr-windows.exe -ldflags '-w -s' -a cmd/lockerr/main.go

release: compile

run:
	go run cmd/lockerr/main.go