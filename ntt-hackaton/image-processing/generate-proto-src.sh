# https://gist.github.com/sofyanhadia/37787e5ed098c97919b8c593f0ec44d8
#protoc --go_out=. ./protos/camdevice-to-image.proto
# protoc -I protos/ --go_out=. ./protos/camdevice-to-image.proto
# protoc --go_out=plugins=grpc:$$GOPATH/src/ ./protos/camdevice-to-image.proto
# protoc -I${GOPATH}/src --go_out=plugins=grpc:api ./protos/camdevice-to-image.proto

protoc -I protos protos/camdevice-to-image.proto --go_out=plugins=grpc:protos
protoc -I smartpark smartpark/smartpark.proto --go_out=plugins=grpc:smartpark
