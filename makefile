API_PROTO_FILES=$(shell find plugins -name *.proto)

# build server
# .PHONY: build
# build:
# 	go build -o ./test/plugins ./bootstrap/

# build proto file to go file
.PHONY: protoc
protoc:
	python -m grpc_tools.protoc \
		   -I ./plugins/ \
           --python_out=./plugins \
		   --grpc_python_out=./plugins \
		   $(API_PROTO_FILES)