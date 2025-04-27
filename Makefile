run:
	rm -rf ./bin/*
	GOOS=linux GOARCH=arm64 go build -o ./bin/test-module ./test-module
	GOOS=linux GOARCH=arm64 go build -o ./bin/runtime .
buf:
	rm -rf ./proto/**/*.pb.go && buf generate