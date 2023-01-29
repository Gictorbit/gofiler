.PHONY:build
build:
	@echo "build server..."
	go build -o bin/server github.com/Gictorbit/gofiler/server/cmd

	@echo "build client..."
	go build -o bin/client github.com/Gictorbit/gofiler/client/cmd
.PHONE:clean
clean:
	@rm -r bin
release:
	@echo "release server..."
	go build -o bin/server -ldflags '-s' github.com/Gictorbit/gofiler/server/cmd

	@echo "release client..."
	go build -o bin/client -ldflags '-s' github.com/Gictorbit/gofiler/client/cmd

.PHONY:compile
compile:
	@echo "Compiling for every OS and Platform"
	GOOS=linux GOARCH=arm go build -o bin/main-linux-arm github.com/Gictorbit/gofiler/server/cmd
	GOOS=linux GOARCH=arm64 go build -o bin/main-linux-arm64 github.com/Gictorbit/gofiler/server/cmd

.PHONY:run-server
run-server:
	@echo "running server"
	go build -o bin/server github.com/Gictorbit/gofiler/server/cmd
	./bin/server start

.PHONY:proto
proto:
	@echo "run proto linter..."
	@cd proto && buf lint && cd -

	@echo "format proto..."
	@cd proto && buf format -w && cd -

	@echo "generate proto..."
	@cd proto && buf generate && cd -