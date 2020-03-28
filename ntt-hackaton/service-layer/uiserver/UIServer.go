package uiserver

import (
	"net/http"
	"log"
	"service-layer/datastore"
	"encoding/json"
)

var dataStoreInstance *datastore.SmartParkDataStore

func ListenUI(store *datastore.SmartParkDataStore) {
	dataStoreInstance = store

	http.HandleFunc("/vehicleParkingInfo", vehicleParkingInfo)
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}

func vehicleParkingInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	plate := r.URL.Query().Get("plate")

	if dataStoreInstance.IsPlateInParkingLot(plate) {
		send(VehicleStatus{
			Entered:     true,
			ReservedLot: dataStoreInstance.GetReservedParkingLot(plate),
			EnteredLot:  dataStoreInstance.CurrentParkingLot(plate),
			DwellTime:   dataStoreInstance.CurrentParkingDuration(plate),
		}, w)
	} else if dataStoreInstance.IsPlateInHistory(plate) {
		send(VehicleStatus{
			Entered:     false,
			ReservedLot: -1,
			EnteredLot:  -1,
			DwellTime:   dataStoreInstance.LatestParkDuration(plate),
		}, w)
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}

func send(status VehicleStatus, w http.ResponseWriter) {
	responseJson, err := json.Marshal(status)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Write(responseJson)
}

type VehicleStatus struct {
	Entered     bool  `json:"entered"`
	ReservedLot int32 `json:"reservedLot"`
	EnteredLot  int32 `json:"enteredLot"`
	DwellTime   int   `json:"dwellTime"`
}
