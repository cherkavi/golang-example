package imageprocessingserver

import (
	"net"
	"google.golang.org/grpc/reflection"
	pb "service-layer/smartpark"
	"context"
	"log"
	"google.golang.org/grpc"
	"service-layer/datastore"
)

const (
	port = ":8081"
)

var dataStoreInstance *datastore.SmartParkDataStore
type server struct{}

func (s *server) SendParkingLotEvent(ctx context.Context, in *pb.ParkingLotEvent) (*pb.Ack, error) {
	if in.Plate == "" {
		dataStoreInstance.ParkingLotEmpty(in.ParkingLotId)
	} else {
		// If plate is not in parking lot, we didn't recognise that the car entered. Just put it in now.
		// Thats better than a technical error.
		if !dataStoreInstance.IsPlateInParkingLot(in.Plate) {
			dataStoreInstance.LicensePlateEnter(in.Plate)
		}
		dataStoreInstance.ParkingLotOccupied(in.ParkingLotId, in.Plate)
	}

	return &pb.Ack{}, nil
}

func (s *server) SendLicensePlateEvent(ctx context.Context, in *pb.LicensePlateEvent) (*pb.Ack, error) {
	if in.Direction == pb.LicensePlateEvent_ENTRY {
		dataStoreInstance.LicensePlateEnter(in.Plate)
	} else if in.Direction == pb.LicensePlateEvent_EXIT {
		dataStoreInstance.LicensePlateExit(in.Plate)
	}

	return &pb.Ack{}, nil
}

func ListenImageProcessing(store *datastore.SmartParkDataStore) {
	dataStoreInstance = store
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterSmartParkProtoServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}