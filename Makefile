build:
	# for windows
	env GOOS=windows GOARCH=386 go build -o bin/main.exe main.go
	# for mac
	env GOOS=darwin GOARCH=amd64 go build -o bin/main-mac main.go