# SmartPark - Service Layer
The SmartPark Service Layer receives event data of the sensors, stores it and implements the logic to generate
meaningful information out of the data.

## Quick Start
```
protoc -I smartpark smartpark/smartpark.proto --go_out=plugins=grpc:smartpark
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

### Image Processing <--> Service Layer
- License Plate y at entry
- License Plate y at exit
- Parking space X occupied with plate y
- Parking space X free