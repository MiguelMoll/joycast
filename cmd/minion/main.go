package main

import (
	"fmt"
	"os"

	"github.com/MiguelMoll/joycast/minion"
)

func main() {
	redisURL := os.Getenv("REDIS_URL")
	if redisURL == "" {
		redisURL = "redis://localhost:6379/"
	}

	m, err := minion.New(
		minion.RedisURL(redisURL),
		minion.RedisHandle("youtube_upload", minion.YoutubeHandler()),
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := m.Run(); err != nil {
		fmt.Println(err)
	}
}
