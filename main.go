package main

import (
	"log"
	"net/http"
	"time"

	bolt "go.etcd.io/bbolt"
)

func main() {
	db, err := bolt.Open("my.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	server := NewPlayerServer(NewBoltPlayerStore(db))
	log.Fatal(http.ListenAndServe(":5001", server))
}
