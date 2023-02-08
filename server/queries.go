package server

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func CreateConnection() *sql.DB {
	db, err := sql.Open("postgres", "postgres://root@localhost:26257/zoomevents?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successful Connection")
	return db
}


func InsertEvent(event *Event) string {
	db := CreateConnection()
	fmt.Println("Inserting: ", event.Event)

	defer db.Close()

	sql := fmt.Sprintf("UPSERT INTO events (id, event, event_ts, accountid, email) VALUES ('%s', '%s', '%d', '%s', '%s');", event.Payload.Object.UUID, event.Event, event.EventTS, event.Payload.AccountID, event.Payload.Operator)

	if _, err := db.Exec(sql); err != nil {
		log.Fatalf("This is the error %s", err)
	}

	return "Event Successfully Inserted"
}
