package video

import (
	"github.com/rs/zerolog"
)

type Uploader interface {
	Upload() error
}

// VideoSeer is high level type to interact and create
// video data
type VideoSeer struct {
	Config
}

// Config configures an VideoSeer
type Config struct {
	Log zerolog.Logger
}

// New creates a VideoSeer type
func New(config Config) *VideoSeer {
	as := &VideoSeer{
		Config: config,
	}

	as.Log = as.Log.With().Str("context", "videoseer").Logger()

	return as
}

func (vs *VideoSeer) UploadFromFile() error {
	return nil
}
