package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"time"

	"context"

	"./configuration"
	"./listener"
	"./protos"
	"./recognizer"
	"./smartpark"
	grpc "google.golang.org/grpc"
)

var configEntries []configuration.ConfigEntry

func generateSubImages() []*recognizer.SubImage {
	var subImages []*recognizer.SubImage
	for _, configEntry := range configEntries {
		subImages = append(subImages, recognizer.NewSubImage(configEntry.ID, configEntry.Direction, configEntry.X, configEntry.Y, configEntry.Width, configEntry.Height))
	}
	// TODO change to real logic
	// subImages = append(subImages, recognizer.NewSubImage(2, smartpark.LicensePlateEvent_ENTRY, 300, 000, 300, 768))
	// subImages = append(subImages, recognizer.NewSubImage(1, -1, 700, 000, 330, 768))
	// smartpark.LicensePlateEvent_EXIT
	return subImages
}

var portNumber int
var configFilePath string
var serviceLayerUrl string

func init() {
	flag.IntVar(&portNumber, "portNumber", 9999, "number of the port to be listen")
	flag.StringVar(&configFilePath, "config", "configuration.xml", "path to config file ")
	flag.StringVar(&serviceLayerUrl, "serviceLayerUrl", "service.smartpark.stritzke.me:8081", "url to ServiceLayer component")

}

type logicExecutor struct {
}

func createTempFile() (*os.File, error) {
	tempFile, err := ioutil.TempFile("", "jpg")
	if err != nil {
		return nil, err
	}
	defer os.Remove(tempFile.Name())
	return os.OpenFile(tempFile.Name()+".jpg", os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)
}

func writeAnswerIntoTempFile(values []byte) (*os.File, error) {
	var tempFile *os.File
	var err error
	if tempFile, err = createTempFile(); err != nil {
		log.Panic("can't create temp file ")
		return nil, err
	}
	defer tempFile.Close()

	tempFile.Write(values)
	tempFile.Sync()
	return tempFile, err
}

func sendToServiceLayer(url string, id int32, number string, direction smartpark.LicensePlateEvent_Direction) error {
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()

	protoClient := smartpark.NewSmartParkProtoClient(conn)
	// TODO reconsider "cancel"
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if direction >= 0 {
		_, err = protoClient.SendLicensePlateEvent(ctx, &smartpark.LicensePlateEvent{
			Plate:     number,
			Direction: direction,
		})
	} else {
		_, err = protoClient.SendParkingLotEvent(ctx, &smartpark.ParkingLotEvent{
			ParkingLotId: int32(id),
			Plate:        number,
		})
	}
	return err
}

func (logicExecutor logicExecutor) Process(values []byte) {
	if tempFile, err := writeAnswerIntoTempFile(values); err != nil {
		log.Panic("can't save data into temp file ", err)
	} else {
		defer tempFile.Close()
		log.Printf("saved file:  %v \n", tempFile.Name())
		// output, err := RecognizeImage(file.Name())
		var subImages = generateSubImages()
		tempFile2, _ := os.OpenFile(tempFile.Name(), os.O_RDONLY, 400)
		defer tempFile2.Close()
		recognizer.CutSubImages(tempFile2, subImages...)
		defer os.Remove(tempFile2.Name())

		for _, image := range subImages {
			image.RecognizedValue, err = recognizer.RecognizeImage(image.PathToImage)
			log.Printf("send: %v %v \n", image, err)
			defer os.Remove(image.PathToImage)
			if err := sendToServiceLayer(serviceLayerUrl, image.ID, image.RecognizedValue, image.Direction); err != nil {
				log.Printf("can't send data to server: %v \n", err)
			}

		}

	}
}

func createAndRegisterServer() *grpc.Server {
	grpcServer := grpc.NewServer()
	protos.RegisterChunkerServer(grpcServer, listener.NewServer(logicExecutor{}))
	return grpcServer
}

func serve(grpcServer *grpc.Server) {
	tcpListener, err := net.Listen("tcp", fmt.Sprintf(":%v", portNumber))
	if err != nil {
		panic(err)
	}
	// not necessary defer tcpListener.Close()

	log.Printf("Serving on localhost:%v", portNumber)
	log.Fatalln(grpcServer.Serve(tcpListener))
}

func readConfiguration() {
	configEntries = configuration.ReadConfiguration(configFilePath)
}

func main() {
	flag.Parse()

	readConfiguration()

	grpcServer := createAndRegisterServer()

	serve(grpcServer)
}
