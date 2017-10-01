package spinner

import "time"

// Theme is theme of the spinner.
type Theme struct {
	Icon  []string
	Speed time.Duration // Speed is the interval time to display icon.
}

// NewTheme generate spinner theme
func NewTheme(speed int64, icon []string) Theme {
	return Theme{
		Speed: time.Duration(speed) * time.Second,
		Icon:  icon,
	}
}

// ThemeMap has theme structs
var ThemeMap = map[int]Theme{
	1: NewTheme(2, []string{"▁", "▃", "▄", "▅", "▆", "▇", "█", "▇", "▆", "▅", "▄", "▃", "▁"}),
	2: NewTheme(2, []string{"⣾", "⣽", "⣻", "⢿", "⡿", "⣟", "⣯", "⣷"}),
	3: NewTheme(2, []string{"⠁", "⠂", "⠄", "⡀", "⢀", "⠠", "⠐", "⠈"}),
	4: NewTheme(2, []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}),
	5: NewTheme(2, []string{">>--->", " >>--->", "  >>--->", "   >>--->", "    >>--->", "    <---<<", "   <---<<", "  <---<<", " <---<<", "<---<<"}),
	6: NewTheme(2, []string{"|", "||", "|||", "||||", "|||||", "|||||||", "||||||||", "|||||||", "||||||", "|||||", "||||", "|||", "||", "|"}),
}
