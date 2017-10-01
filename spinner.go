package spinner

import (
	"io"
	"os"
	"time"

	"github.com/fatih/color"
)

const (
	ANSI_CURSOR_UP   = "\033[A"
	ANSI_CURSOR_DOWN = "\033[B"
	ANSI_NEXT_LINE   = "\033[E"
	ANSI_PREV_LINE   = "\033[F"
	ANSI_DEL_LINE    = "\033[2K"
)

// NewSpinner create pinner instance
func NewSpinner() Spinner {
	spinner := Spinner{
		Writer: os.Stdout,
	}
	return spinner
}

// Spinner is spinning icon.
type Spinner struct {
	Theme       Theme
	Color       color.Attribute
	Prefix      string
	Suffix      string
	CompleteMSG string
	complete    bool
	Writer      io.Writer
	Interval    time.Duration
}

func (s Spinner) write(msg []byte) error {
	_, err := s.Writer.Write(msg)
	return err
}

// Start is function to start spinner
func (s Spinner) Start() error {
	return nil
}
