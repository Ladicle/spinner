package main

import (
	"time"

	"github.com/Ladicle/spinner"
)

func main() {
	lines := []string{
		"first",
		"second",
		"third",
	}
	icon := []string{".", "..", "..."}
	spin := spinner.NewSpinner()
	for i := 0; i < 10; i++ {
		for _, l := range lines {
			_, _ = spin.Writer.Write([]byte("\033[2K\r"))
			_, _ = spin.Writer.Write([]byte(l + icon[i%3] + "\n"))
		}
		_, _ = spin.Writer.Write([]byte("\033[A\033[A\033[A"))
		time.Sleep(1 * time.Second)
	}
}
