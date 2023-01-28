.PHONY:build
build:
	@echo "build server..."
	go build -o bin/server github.com/Gictorbit/gofiler/server/cmd

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