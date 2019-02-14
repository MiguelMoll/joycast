package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/rs/zerolog"

	"github.com/MiguelMoll/joycast/audio"
)

func main() {
	as := audio.New(audio.Config{
		Log: zlog(),
	})

	tmpfile, err := ioutil.TempFile("", "audio")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(tmpfile.Name())
	err = as.FromVideo("/home/radioact1ve/Downloads/ocean.mp4", tmpfile.Name())
	fmt.Println(err)
}

func zlog() zerolog.Logger {
	return zerolog.New(os.Stdout).With().Str("app", "joycast").Logger()
}
