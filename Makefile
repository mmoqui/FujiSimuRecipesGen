BINARY=FujiSimuRecipesGen

build: clean linux windows macosx

linux:
	GOOS=linux GOARCH=amd64 go build -o build/linux/${BINARY}

windows:
	GOOS=windows GOARCH=amd64 go build -o build/windows/${BINARY}

macosx:
	GOOS=darwin GOARCH=amd64 go build -o build/macosx/${BINARY}

clean:
	go clean
	rm -f build/linux/${BINARY}
	rm -f build/windows/${BINARY}
	rm -f build/macosx/${BINARY}
