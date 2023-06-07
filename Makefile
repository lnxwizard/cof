run:
	go run cmd/cof/main.go go.mod

build-win-32:
	echo "Setting Environments"
	set GOOS=windows
	set GOARCH=386

	echo "Building cof"
	go build -o .out/cof32.exe cmd/cof/main.go go.mod

build-win-64:
	echo "Setting Environments"
	set GOOS=windows
	set GOARCH=amd64

	echo "Building cof"
	go build -o .out/cof32.exe cmd/cof/main.go go.mod

clean:
	go clean