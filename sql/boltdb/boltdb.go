package main

import (
	"fmt"
	"time"

	"github.com/boltdb/bolt"
)

const (
	bucketName = "technik_bucket"
)

func addRecord(db *bolt.DB, key, value string) {
	updateFunction := func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		if err != nil {
			return err
		}
		return bucket.Put([]byte(key), []byte(value))
	}
	if err := db.Update(updateFunction); err != nil {
		fmt.Println("can't add record into db: ", err)
		return
	}
}

func readRecord(db *bolt.DB, key string) {
	viewFunction := func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		v := bucket.Get([]byte(key))
		fmt.Println("data from Database is: ", string(v))
		return nil
	}
	if err := db.View(viewFunction); err != nil {
		fmt.Println("can't execute operation of reading data: ", err)
		return
	}
}

func main() {
	fmt.Println("---- working with BoltDB database (single file DB ) ----")
	db, err := bolt.Open("test.db", 0700, &bolt.Options{Timeout: 2 * time.Second}) // to prevent waiting infinite time
	if err != nil {
		fmt.Println(" error create/open file: ", err)
		return
	}
	defer db.Close()

	addRecord(db, "solely", time.Now().Format(time.RFC822))
	readRecord(db, "solely")
}
