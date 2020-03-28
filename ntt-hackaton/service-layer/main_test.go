package main

import (
	"sync"
	"time"
	"testing"
	"google.golang.org/grpc"
	pb "service-layer/smartpark"
	"context"
)

var once sync.Once

func Test_SendParkingLotOccupied(t *testing.T) {
	once.Do(startServer)

	conn, err := grpc.Dial("localhost:8081", grpc.WithInsecure())
	if err != nil {
		t.Errorf("did not connect: %v", err)
		return
	}
	defer conn.Close()

	c := pb.NewSmartParkProtoClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err = c.SendParkingLotEvent(ctx, &pb.ParkingLotEvent{
		ParkingLotId: 1,
		Plate: "SZ DS 3195",
	})
	if err != nil {
		t.Errorf("could not get response: %v", err)
	}
}

func Test_SendParkingLotEmpty(t *testing.T) {
	once.Do(startServer)

	conn, err := grpc.Dial("localhost:8081", grpc.WithInsecure())
	if err != nil {
		t.Errorf("did not connect: %v", err)
		return
	}
	defer conn.Close()

	c := pb.NewSmartParkProtoClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err = c.SendParkingLotEvent(ctx, &pb.ParkingLotEvent{
		ParkingLotId: 1,
		Plate: "",
	})
	if err != nil {
		t.Errorf("could not get response: %v", err)
	}
}

func Test_SendLicensePlateEntry(t *testing.T) {
	once.Do(startServer)
	conn, err := grpc.Dial("localhost:8081", grpc.WithInsecure())
	if err != nil {
		t.Errorf("did not connect: %v", err)
		return
	}
	defer conn.Close()

	c := pb.NewSmartParkProtoClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err = c.SendLicensePlateEvent(ctx, &pb.LicensePlateEvent{
		Plate: "SZ DS 3195",
		Direction: pb.LicensePlateEvent_ENTRY,
	})
	if err != nil {
		t.Errorf("could not get response: %v", err)
	}
}

func Test_SendLicensePlateExit(t *testing.T) {
	once.Do(startServer)
	conn, err := grpc.Dial("localhost:8081", grpc.WithInsecure())
	if err != nil {
		t.Errorf("did not connect: %v", err)
		return
	}
	defer conn.Close()

	c := pb.NewSmartParkProtoClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err = c.SendLicensePlateEvent(ctx, &pb.LicensePlateEvent{
		Plate: "SZ DS 3195",
		Direction: pb.LicensePlateEvent_EXIT,
	})
	if err != nil {
		t.Errorf("could not get response: %v", err)
	}
}

func startServer() {
	go main()
	time.Sleep(time.Second * 1)
}
