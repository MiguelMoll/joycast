package main

import (
	"fmt"
	"log"
	"os"

	"github.com/MiguelMoll/joycast/minion"
	"github.com/MiguelMoll/joycast/realm"
	"github.com/MiguelMoll/joycast/storage/db"
)

func main() {
	redisURL := os.Getenv("REDIS_URL")
	if redisURL == "" {
		redisURL = "redis://localhost:6379/"
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

	yt, err := minion.YoutubeHandler(
		minion.YoutubeUsers(us),
	)
	if err != nil {
		log.Fatal(err)
	}

	m, err := minion.New(
		minion.RedisURL(redisURL),
		minion.RedisHandle("youtube_upload", yt),
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := m.Run(); err != nil {
		fmt.Println(err)
	}
}
