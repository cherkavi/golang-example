package listener

import (
	"io"
	"log"

	"../protos"
)

// const chunkSize = 64 * 1024 // 64 KiB

type LogicProcessor interface {
	Process(v []byte)
}

type ChunkerServer struct {
	logicProcessor LogicProcessor
}

func NewServer(processor LogicProcessor) *ChunkerServer {
	s := &ChunkerServer{}
	s.logicProcessor = processor
	return s
}

func (chunkerServer ChunkerServer) Chunker(server protos.Chunker_ChunkerServer) error {
	var err error
	for {
		c, err := server.Recv()
		// fmt.Printf(">>> Received next data chunk \n")
		if err != nil {
			if err == io.EOF {
				log.Printf("Received data size:  %d", len(c.Chunk))
				chunkerServer.logicProcessor.Process(c.Chunk)
				continue
			}
			goto endReading
		} else {
			log.Printf("Received next data size:  %d", len(c.Chunk))
			chunkerServer.logicProcessor.Process(c.Chunk)
		}
	}
endReading:
	return err
}
