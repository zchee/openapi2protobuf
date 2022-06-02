package dump

import (
	"strings"

	"github.com/Code-Hex/dd/p"
	"github.com/alecthomas/chroma/formatters"
	"github.com/alecthomas/chroma/styles"
)

// Dump dumps a for debug.
func Dump(a ...interface{}) {
	p.New(p.WithStyle(styles.DoomOne2), p.WithFormatter(formatters.TTY16m)).P(a...)
}

// Sdump dumps a string for debug.
func Sdump(a ...interface{}) string {
	var s strings.Builder
	p.New(p.WithStyle(styles.DoomOne2), p.WithFormatter(formatters.TTY16m)).Fp(&s, a...)

	return s.String()
}
