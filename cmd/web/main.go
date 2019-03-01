package main

import (
	"fmt"
	"log"
	"os"

	"github.com/MiguelMoll/joycast/realm"
	"github.com/MiguelMoll/joycast/storage/db"
	"github.com/MiguelMoll/joycast/web"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = "postgres://jc@localhost:5432/jc?sslmode=disable"
	}

	db, err := db.New(dbURL)
	if err != nil {
		// TODO: Handle this error better!
		log.Fatal(err)
	}
	defer db.Close()

	us, err := realm.NewUserService(
		realm.UserRepo(db),
	)
	if err != nil {
		log.Fatal(err)
	}

	server, err := web.New(
		web.UserService(us),
	)
	if err != nil {
		log.Fatal(err)
	}

	server.Start(fmt.Sprintf(":%s", port))
}
