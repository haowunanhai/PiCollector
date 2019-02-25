package log

// brush is a color join function
type brush func(string) string

// newBrush return a fix color Brush
func newBrush(color string) brush {
	pre := "\033["
	reset := "\033[0m"
	return func(text string) string {
		return pre + color + "m" + text + reset
	}
}

var colors = []brush{
	newBrush("1;34"), // Debug              blue
	newBrush("1;36"), // Trace              cyan
	newBrush("1;32"), // Informational      green
	newBrush("1;33"), // Warning            yellow
	newBrush("1;31"), // Error              red
	newBrush("1;37"), // Fatal          	white
}
