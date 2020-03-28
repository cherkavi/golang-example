# SmartPark - Image Processing layer
Waiting for images from "cam-device-layer" ( --portNumber 9999), 
parse image according configuration ( --config configuration.xml ), 
notify "service-layer" (--serviceLayerUrl "service.smartpark.stritzke.me:8081")  
about events what were found

## Quick Start
```
get-protobuf.sh
generate-proto-src.sh
go run main.go
```

## WebSocket Ping
```
go run main.go
# Connect to localhost:8080/parkingLotStatus
# Send something
# You'll receive pong
```

## Interface

### Cam Device Service --> Image Processing
- gRPC binary format ( JPEG file )

### Image Processing --> Service Layer
- License Plate y at entry
- License Plate y at exit
- Parking space X occupied with plate y
- Parking space X free