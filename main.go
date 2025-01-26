package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/cdated/cycles/events"
	"github.com/jackc/pgx/v5"
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

	conn, err := pgx.Connect(context.Background(), os.Getenv("DB_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	// Add sample events
	eventOps()
}
