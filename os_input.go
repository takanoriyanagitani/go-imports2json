package imports2json

import (
	"bufio"
	gt "go/token"
	"io"
	"os"
)

// ParseStdin parses Go source from `os.Stdin`.
func (m ParseMode) ParseStdin(fset *gt.FileSet, name string) (GoFile, error) {
	var br io.Reader = bufio.NewReader(os.Stdin)
	return m.ParseReader(fset, name, br)
}

// ParseStdinDefault is a convenience function to parse from `os.Stdin` with
// a default fileset and name.
func (m ParseMode) ParseStdinDefault() (GoFile, error) {
	var fset *gt.FileSet = gt.NewFileSet()
	return m.ParseStdin(fset, "stdin")
}
