export GOPROXY=https://goproxy.cn,direct
export GO111MODULE=on

OBJ = server.exe client.exe

default: $(OBJ)

$(OBJ):
	go build -gcflags "-N -l" -o $@ ./$(@:.exe=)

clean:
	rm -fr $(OBJ)

-include .deps

generate:
	protoc --go_out=./ --go_opt=paths=source_relative --go-grpc_out=./ --go-grpc_opt=paths=source_relative proto/*.proto

dep:
	echo -n "$(OBJ):" > .deps
	find . -name '*.go' | awk '{print $$0 " \\"}' >> .deps
	echo "" >> .deps