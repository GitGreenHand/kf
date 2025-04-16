
build:
	go build -o ./bin/kf

amdLinux:
	GOOS=linux GOARCH=amd64  go build -o ./bin/kf

windows:
	GOOS=windows GOARCH=amd64 go build -o ./bin/kf.exe
