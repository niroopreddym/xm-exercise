
build:	winbuild macbuild linuxbuild
winbuild:
	GOOS=windows GOARCH=amd64 go build -o bin/dss-amd64-win.exe .

macbuild:
	GOOS=darwin GOARCH=amd64 go build -o bin/dss-amd64-darwin .

linuxbuild:
	GOOS=linux GOARCH=amd64 go build -o bin/dss-amd64-linux .
