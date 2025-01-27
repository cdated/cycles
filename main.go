package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/cdated/cycles/db"
	"github.com/cdated/cycles/events"
	"github.com/joho/godotenv"
)

func eventOps() {
	fmt.Println("Create event")
	e1 := events.NewEventRecurs("Use device a", time.Second)
	fmt.Println(e1)

	fmt.Println("Complete event")
	e1.Complete()
	fmt.Println(e1)

	fmt.Println("Attempt reopen event")
	e1.Reopen()
	fmt.Println(e1)

	fmt.Println("Wait...attempt reopen event")
	time.Sleep(time.Second * 2)
	e1.Reopen()
	fmt.Println(e1)

	e2 := events.NewEvent("Use device b")
	fmt.Println(e2)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	connString := os.Getenv("DB_URL")
	pg, err := db.NewPG(context.Background(), connString)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Get users")
	users, err := pg.GetUsers(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to query database: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Users: %v\n", users)

	// Add sample events
	eventOps()
}
