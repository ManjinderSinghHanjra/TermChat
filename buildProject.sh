protoc -I . services.proto --go_out=plugins=grpc:common/generatedCode

go build server