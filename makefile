SHELL := /bin/bash

# Input directory where proto files exist.
PROTO_INPUT_DIR=proto

# Output directories for gRPC stubs generated by protoc.
PROTO_TS_OUTPUT_DIR=typescript/generated
PROTO_GO_OUTPUT_DIR=golang

# There is NOT native support for TS in protoc, so specify a plugin for TS that we've installed with NPM.
# To do this, "npm install protoc-gen-ts" and reference the local plugin.
PROTOC_GEN_TS_PATH=node_modules/.bin/protoc-gen-ts

# Generate typescript and golang gRPC stubs for each proto file in the proto directory.
buf:
	chmod +x buf-clean.sh && ./buf-clean.sh $(PROTO_GO_OUTPUT_DIR) && ./buf-clean.sh $(PROTO_TS_OUTPUT_DIR)
	@for file in $(PROTO_INPUT_DIR)/*.proto; do \
		if [ -f "$$file" ]; then \
			echo "Generating protos for $$file"; \
			protoc \
				--proto_path=$(PROTO_INPUT_DIR) \
				--go_out=${PROTO_GO_OUTPUT_DIR} \
				--go_opt=paths=source_relative \
				--go-grpc_out=${PROTO_GO_OUTPUT_DIR} \
				--go-grpc_opt=paths=source_relative \
				--plugin="protoc-gen-ts=${PROTOC_GEN_TS_PATH}" \
				--js_out="import_style=commonjs,binary:${PROTO_TS_OUTPUT_DIR}" \
				--ts_out="service=grpc-web:${PROTO_TS_OUTPUT_DIR}" \
				$$file; \
		fi; \
	done
	chmod +x publish-npm.sh && ./publish-npm.sh