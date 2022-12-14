API_PROTO_FILES=$(shell find plugins -name *.proto)

# build server
# .PHONY: build
# build:
# 	go build -o ./test/plugins ./bootstrap/

# build proto file to go file
.PHONY: protoc
protoc:
	protoc --proto_path=./plugins \
           --go_out=paths=source_relative:./plugins \
		   --go-http_out=paths=source_relative:./plugins \
		   --go-grpc_out=paths=source_relative:./plugins \
		   $(API_PROTO_FILES)
#		   --openapi_out=fq_schema_naming=true,default_response=false:. \

.PHONY: ent
ent:
	ent generate ./db/schema --target ./db/gen