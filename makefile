# Input directory where proto files exist.
PROTO_INPUT_DIR=proto

# Output directories for protobuf-generated code for RPCs.
PROTO_TS_OUTPUT_DIR=ts_proto
PROTO_GO_OUTPUT_DIR=go_proto

# There is NOT native support for TS in protoc, so specify a plugin for TS that we've installed with NPM.
# To do this, "npm install protoc-gen-ts" and reference the local plugin.
PROTOC_GEN_TS_PATH=node_modules/.bin/protoc-gen-ts

job-application-go:
	protoc \
		--proto_path=proto \
		--go_out=${PROTO_GO_OUTPUT_DIR} \
		--go_opt=paths=source_relative \
		--go-grpc_out=${PROTO_GO_OUTPUT_DIR} \
		--go-grpc_opt=paths=source_relative \
		${PROTO_INPUT_DIR}/job_application.proto
	
job-application-node:
	protoc \
		--proto_path=proto \
		--plugin="protoc-gen-ts=${PROTOC_GEN_TS_PATH}" \
		--js_out="import_style=commonjs,binary:${PROTO_TS_OUTPUT_DIR}" \
		--ts_out="service=grpc-web:${PROTO_TS_OUTPUT_DIR}" \
		${PROTO_INPUT_DIR}/job_application.proto

job-application: job-application-go job-application-node

job-posting-go:
	protoc \
		--proto_path=proto \
		--go_out=${PROTO_GO_OUTPUT_DIR} \
		--go_opt=paths=source_relative \
		--go-grpc_out=${PROTO_GO_OUTPUT_DIR} \
		--go-grpc_opt=paths=source_relative \
		${PROTO_INPUT_DIR}/job_posting.proto

job-posting-node:
	protoc \
		--proto_path=proto \
		--plugin="protoc-gen-ts=${PROTOC_GEN_TS_PATH}" \
		--js_out="import_style=commonjs,binary:${PROTO_TS_OUTPUT_DIR}" \
		--ts_out="service=grpc-web:${PROTO_TS_OUTPUT_DIR}" \
		${PROTO_INPUT_DIR}/job_posting.proto

job-posting: job-posting-go job-posting-node

gen: job-application job-posting



