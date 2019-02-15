package audio

import (
	"github.com/rs/zerolog"

	"github.com/MiguelMoll/joycast/exec"
)

// AudioSeer is high level type to interact and create
// audio data
type AudioSeer struct {
	Config
	exec exec.Executor
}

// Config configures an AudioSeer
type Config struct {
	Log zerolog.Logger
}

// New creates a AudioSeer type
func New(config Config) *AudioSeer {
	as := &AudioSeer{
		Config: config,
		exec:   exec.New(),
	}

	as.Log = as.Log.With().Str("context", "audioseer").Logger()

	return as
}

// FromVideoFile extracts the audio from a video file
func (as *AudioSeer) FromVideoFile(file string, dest string) error {
	output, err := as.exec.Run("ffmpeg", "-y", "-i", file, "-f", "mp3", "-ab", "192000", "-vn", dest)
	if err != nil {
		as.Log.Error().Err(err).Str("stderr", output.StdErr).Msg("ffmpeg failed to extract audio")
	}
	return err
}
