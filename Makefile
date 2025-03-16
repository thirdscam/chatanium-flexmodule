run:
	rm -rf ./bin/*
	go build -o ./bin/test-module ./test-module
	go build -o ./bin/runtime .
	./bin/runtime

buf:
	rm -rf ./proto/**/*.pb.go && buf generate