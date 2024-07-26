

this code not working***************
look in the grpc2 folder for working example


-install - protoc (download it from https://github.com/protocolbuffers/protobuf/releases)
-update the path in windows control panel
- install protoc-gen-go (go install google.golang.org/protobuf/cmd/protoc-gen-go)
- update go.mode file 
    require (
    google.golang.org/protobuf v1.5.4 // or your desired version
    google.golang.org/protobuf/cmd/protoc-gen-go v0.0.0-20240725165638-a5025b7db78d // Latest version at the time of writing
    )
-
