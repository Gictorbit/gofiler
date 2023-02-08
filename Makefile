.PHONY:build
build:
	@echo "build server..."
	go build -ldflags="-s -w" -o bin/server github.com/Gictorbit/gofiler/server/cmd
	GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o bin/server.exe github.com/Gictorbit/gofiler/server/cmd

	@echo "build client..."
	go build -ldflags="-s -w" -o bin/client github.com/Gictorbit/gofiler/client/cmd
	GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o bin/client.exe github.com/Gictorbit/gofiler/client/cmd
.PHONE:clean
clean:
	@rm -r bin

.PHONY:proto
proto:
	@echo "run proto linter..."
	@cd proto && buf lint && cd -

	@echo "format proto..."
	@cd proto && buf format -w && cd -

	@echo "generate proto..."
	@cd proto && buf generate && cd -