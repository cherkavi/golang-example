package main

import (
	"fmt"
	"log"
	"os"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	databaseName   = "technik_db"
	collectionName = "technik_collection"
)

func openSession() *mgo.Session {
	// https://docs.mongodb.com/manual/reference/connection-string/
	// mongodb://[username:password@]host1[:port1][,host2[:port2],...[,hostN[:portN]]][/[database][?options]]
	session, err := mgo.Dial(fmt.Sprintf("mongodb://127.0.0.1:27017"))
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return session
}

// all elements should be public
type Record struct {
	ID   int    `bson: "id"`
	Name string `bson: "name"`
}

func (r Record) String() string {
	return fmt.Sprintf("id: %v  name: %v \n", r.ID, r.Name)
}

func printDatabases(session *mgo.Session) {
	dbNames, err := session.DatabaseNames()
	if err != nil {
		fmt.Println("can't read data from db")
		return
	}
	for _, name := range dbNames {
		fmt.Printf("> db: %v\n", name)
	}
}

func pingDatabase(session *mgo.Session) {
	if err := session.Ping(); err != nil {
		fmt.Println(" error during 'ping' Database ", err)
	} else {
		fmt.Println("> ping OK ")
	}
}

func getCollection(session *mgo.Session) *mgo.Collection {
	return session.DB(databaseName).C(collectionName)
}

func insertData(session *mgo.Session) {
	// if database not exists - just select it - will be enough
	// if collection not existsn - it will be created
	collection := getCollection(session)
	record := Record{}
	record.ID = 100
	record.Name = time.Now().Format(time.RFC822)
	if err := collection.Insert(record); err != nil {
		fmt.Println("error during trying to insert data ", err)
		return
	}
}

func countRecords(session *mgo.Session) {
	collection := getCollection(session)
	if count, err := collection.Count(); err == nil {
		fmt.Println("amount of records: ", count)
	} else {
		fmt.Println("error during trying to count records into collection : ", err)
	}
}

func printAllRecords(session *mgo.Session) {
	collection := getCollection(session)
	findMap := bson.M{"name": bson.M{"$ne": "null"}}
	if count, err := collection.Find(findMap).Count(); err != nil {
		fmt.Println("Error during count record into Collection : ", err)
	} else {
		fmt.Println("Record count : ", count)
	}

	iter := collection.Find(findMap).Iter()
	var result Record
	for iter.Next(&result) {
		fmt.Println(result)
	}
	iter.Close()
}

func deleteData(session *mgo.Session) {
	collection := getCollection(session)
	// selector := bson.M{"id": bson.M{"$eq": nil}}
	selector := bson.M{"id": bson.M{"$eq": 100}}
	if err := collection.Remove(selector); err != nil { // RemoveAll
		fmt.Println("error removing records: ", err)
	}
}

func serverStatus(session *mgo.Session) {
	var result bson.M
	session.Run("serverStatus", &result)
	fmt.Println("one of the parameter from MongoDB: ", result["process"])
}

// docker pull mysql
// docker run --detach --env MYSQL_ROOT_PASSWORD=root --env MYSQL_USER=root --env MYSQL_PASSWORD=root --env MYSQL_DATABASE=technik_db --name golang_mysql --publish 3306:3306 mysql;
func main() {
	fmt.Println("---- open session with MongoDB --- ")
	session := openSession()
	defer session.Close()

	serverStatus(session)
	pingDatabase(session)
	printDatabases(session)

	insertData(session)

	countRecords(session)
	printAllRecords(session)
	deleteData(session)

	// var result interface{}
	// var db *mgo.Database
	// err = session.Run("use technik_db", &result)
	// if err != nil {
	// 	fmt.Println("create DB error ", err)
	// }
	// names, err := db.CollectionNames()
	// if err != nil {
	// 	fmt.Println("can't read collections ", err)
	// 	return
	// }
	// for _, name := range names {
	// 	fmt.Println(name)
	// }

	// session.Run(bson.D{"use", databaseName}, &result)
	// session.DB("local").C(collectionName).Create()
}
