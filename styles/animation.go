package styles

import (
	"fmt"
	"time"
)

// CuteSpinnerFrames are simple, friendly frames for a spinner animation.
var CuteSpinnerFrames = []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}

// SuccessFrames are a minimal success pulse animation.
var SuccessFrames = []string{"[=   ]", "[==  ]", "[=== ]", "[====]"}

// Spinner handles a lightweight terminal spinner.
type Spinner struct {
	message string
	stopCh  chan struct{}
	doneCh  chan struct{}
}

// StartSpinner starts a spinner with a message and returns a handle to stop it.
func StartSpinner(message string) *Spinner {
	s := &Spinner{
		message: message,
		stopCh:  make(chan struct{}),
		doneCh:  make(chan struct{}),
	}
	go s.run()
	return s
}

func (s *Spinner) run() {
	defer close(s.doneCh)

	frame := 0
	ticker := time.NewTicker(120 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-s.stopCh:
			fmt.Print("\r\033[K")
			return
		case <-ticker.C:
			glyph := InfoStyle().Render(CuteSpinnerFrames[frame%len(CuteSpinnerFrames)])
			msg := MutedStyle().Render(s.message)
			fmt.Printf("\r%s %s", glyph, msg)
			frame++
		}
	}
}

// Stop stops the spinner and clears the line.
func (s *Spinner) Stop() {
	close(s.stopCh)
	<-s.doneCh
}

// PlaySuccessAnimation shows a short success pulse and prints a success line.
func PlaySuccessAnimation(message string) {
	for _, frame := range SuccessFrames {
		fmt.Printf("\r%s %s", SuccessStyle().Render(frame), message)
		time.Sleep(70 * time.Millisecond)
	}

	symbol := SuccessStyle().Render(SuccessSymbol)
	fmt.Printf("\r%s %s\n", symbol, message)
}
