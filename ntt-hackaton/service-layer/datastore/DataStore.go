package datastore

import (
	"log"
	"time"
)

type SmartParkDataStore struct {
	platesEntered          map[string]time.Time
	lotStatus              map[int32]string
	parkingDurationHistory []ParkingSession
	lotReservation         map[string]int32
}

type ParkingSession struct {
	Plate    string
	duration int
}

func New() *SmartParkDataStore {
	return &SmartParkDataStore{
		platesEntered:          map[string]time.Time{},
		lotStatus:              map[int32]string{},
		parkingDurationHistory: make([]ParkingSession, 0),
		lotReservation:         map[string]int32{},
	}
}

func (ds *SmartParkDataStore) GetReservedParkingLot(plate string) int32 {
	if lot, ok := ds.lotReservation[plate]; ok {
		return lot
	}

	return -1
}

func (ds *SmartParkDataStore) IsPlateInParkingLot(plate string) bool {
	if _, ok := ds.platesEntered[plate]; ok {
		return true
	}

	return false
}

func (ds *SmartParkDataStore) IsPlateInHistory(plate string) bool {
	return ds.LatestParkDuration(plate) != -1
}

func (ds *SmartParkDataStore) LatestParkDuration(plate string) int {
	latestDuration := -1

	for _, session := range ds.parkingDurationHistory {
		if session.Plate == plate {
			latestDuration = session.duration
		}
	}

	return latestDuration
}

func (ds *SmartParkDataStore) CurrentParkingLot(plate string) int32 {
	for lot := range ds.lotStatus {
		if ds.lotStatus[lot] == plate {
			return lot
		}
	}

	return -1
}

func (ds *SmartParkDataStore) ParkingLotOccupied(lot int32, plate string) {
	log.Printf("Parkinglot '%d' occupied by '%s'\n", lot, plate)

	for lot := range ds.lotStatus {
		if ds.lotStatus[lot] == plate {
			delete(ds.lotStatus, lot)
		}
	}

	ds.lotStatus[lot] = plate
}

func (ds *SmartParkDataStore) ParkingLotEmpty(lot int32) {
	log.Printf("Parkinglot '%d' empty\n", lot)

	delete(ds.lotStatus, lot)
}

func (ds *SmartParkDataStore) LicensePlateEnter(plate string) {
	if ds.IsPlateInParkingLot(plate) {
		log.Printf("License plate '%s' entered, but ignored because it is alredy entered\n", plate)
	} else {
		log.Printf("License plate '%s' entered\n", plate)
		ds.lotReservation[plate] = ds.getNextFreeParkingSpot()
		ds.platesEntered[plate] = time.Now()
	}
}

func (ds *SmartParkDataStore) LicensePlateExit(plate string) {
	if ds.IsPlateInParkingLot(plate) {
		delta := time.Now().Sub(ds.platesEntered[plate])

		session := ParkingSession{
			Plate:    plate,
			duration: int(delta.Seconds()),
		}
		ds.parkingDurationHistory = append(ds.parkingDurationHistory, session)
		delete(ds.platesEntered, plate)
		delete(ds.lotReservation, plate)

		log.Printf("License plate '%s' left. Parked for %d seconds", plate, session.duration)
	} else {
		log.Printf("License plate '%s' left, but ignored because already left", plate)
	}
}

func (ds *SmartParkDataStore) CurrentParkingDuration(plate string) int {
	if ds.IsPlateInParkingLot(plate) {
		delta := time.Now().Sub(ds.platesEntered[plate])
		return int(delta.Seconds())
	}

	return -1
}

func (ds *SmartParkDataStore) getNextFreeParkingSpot() int32 {
	var lotNumber int32
	for lotNumber = 1; lotNumber <= 4; lotNumber++ {
		if _, ok := ds.lotStatus[lotNumber]; ok {

		} else {
			return lotNumber
		}
	}

	return -1
}
