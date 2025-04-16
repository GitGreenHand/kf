
build:
	go build  main.go

amdLinux:
	GOOS=linux GOARCH=amd64  go build


amdWindows:
	 GOOS=windows GOARCH=amd64 go build
