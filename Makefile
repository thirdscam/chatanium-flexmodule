run:
	rm -rf ./sample/runtime
	rm -rf ./sample/module
	go build -o ./sample/runtime ./sample
	go build -o ./sample/module ./sample/plugin-go-grpc
	./sample/runtime

buf_generate:
	cd sample && buf generate