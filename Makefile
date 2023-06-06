run:
	go run cmd/cof/main.go go.mod

build/386:
	echo "Setting Environments..."
	set GOOS=windows
	set GOARCH=386

	echo "Building x86"
	go build -o .out/cof_386.exe cmd/cof/main.go

build/amd64:
	echo "Setting Environments..."
	set GOOS=windows
	set GOARCH=amd64

	echo "Building x64"
	go build -o .out/cof_amd64.exe cmd/cof/main.go

clean:
	go clean