-install - protoc (download it from https://github.com/protocolbuffers/protobuf/releases)
-update the path in windows control panel
- install protoc-gen-go (go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest)
- use protoc to generage the output files for protobuf
create a invoice folder then run
protoc --go_out=invoicer --go_opt=paths=source_relative --go-grpc_out=invoicer --go-grpc_opt=paths=source_relative invoicer.proto

run go mod tidy.

