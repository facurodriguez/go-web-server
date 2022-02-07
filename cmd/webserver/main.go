package main

import (
	"github.com/facurodriguez/go-web-server"
	"log"
	"net/http"
	"os"
)

const boltDbFileName = "my.db"
const dbFileName = "game.db.json"

func main() {
	// db, err := bolt.Open(boltDbFileName, 0600, &bolt.Options{Timeout: 1 * time.Second})
	// if err != nil {
	// 	log.Fatalf("problem opening %s %v", boltDbFileName, err)
	// }
	// defer db.Close()

	// server := NewPlayerServer(NewBoltPlayerStore(db))

	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		log.Fatalf("problem opening %s %v", dbFileName, err)
	}

	store, err := poker.NewFileSystemPlayerStore(db)

	if err != nil {
		log.Fatalf("problem creating file system player store, %v ", err)
	}

	server := poker.NewPlayerServer(store)

	log.Fatal(http.ListenAndServe(":5001", server))
}
