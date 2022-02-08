.PHONY: build darwin windows

build:
	go build -o build/ChronoEditor cmd/chrono/main.go

test:
	go test ./... -v

cover:
	go test ./... -coverpkg=./... -coverprofile cp.out
	go tool cover -html=cp.out

package: build
	fyne package -executable ./build/ChronoEditor -os darwin -icon ./assets/icon.png

clean:
	rm -Rf build/*

darwin:
	GOOS=darwin GOARCH=amd64 go build -o build/darwin/ChronoEditor cmd/chrono/main.go
	cd ./build/darwin/ && fyne package -executable ChronoEditor -os darwin -icon ../../assets/icon.png -name ChronoEditor

windows:
	GOOS=windows CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc GOARCH=amd64 go build -o build/windows/ChronoEditor.exe cmd/chrono/main.go
	cd ./build/windows/ && GOOS=windows CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc GOARCH=amd64 fyne package -executable ChronoEditor.exe -os windows -icon ../../assets/icon.png -name ChronoEditor
