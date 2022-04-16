package internal

import (
	"io"
	"strings"

	"github.com/Code-Hex/dd/p"
	"github.com/alecthomas/chroma/formatters"
	"github.com/alecthomas/chroma/styles"
)

var printer = p.New(p.WithStyle(styles.DoomOne2), p.WithFormatter(formatters.TTY16m))

func Dump(a ...interface{}) {
	printer.P(a...)
}

func Fdump(w io.Writer, a ...interface{}) {
	printer.Fp(w, a...)
}

func Sdump(a ...interface{}) string {
	var sb strings.Builder
	printer.Fp(&sb, a...)

	return sb.String()
}
