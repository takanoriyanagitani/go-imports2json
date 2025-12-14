package imports2json_test

import (
	_ "embed"
	ga "go/ast"
	gt "go/token"
	"testing"

	ij "github.com/takanoriyanagitani/go-imports2json"
)

//go:embed testdata.d/import1.go
var import1 string

func TestImportsToJson(t *testing.T) {
	t.Parallel()

	t.Run("noimport", func(t *testing.T) {
		t.Parallel()

		var fset *gt.FileSet = gt.NewFileSet()

		var mode ij.ParseMode = ij.ParseModeDefault
		file, e := mode.ParseString(fset, "", "package noimp")
		if nil != e {
			t.Fatalf("unexpected error: %v\n", e)
		}

		var imports []*ga.ImportSpec = file.ImportSpecs()

		if 0 != len(imports) {
			t.Fatalf("unexpected import got\n")
		}
	})

	t.Run("single import", func(t *testing.T) {
		t.Parallel()

		var fset *gt.FileSet = gt.NewFileSet()

		var mode ij.ParseMode = ij.ParseModeDefault
		file, e := mode.ParseString(fset, "import1.go", import1)
		if nil != e {
			t.Fatalf("unexpected error: %v\n", e)
		}

		var imports []*ga.ImportSpec = file.ImportSpecs()

		if 1 != len(imports) {
			t.Fatalf("expected single import")
		}

		var imp *ga.ImportSpec = imports[0]

		gi := ij.GoImport{ImportSpec: imp}

		if 0 != len(gi.DocComments()) {
			t.Fatalf("unexpected doc comments got\n")
		}

		if "_" != gi.Name() {
			t.Fatalf("unexpected import found: %s\n", gi.Name())
		}

		if `"log"` != gi.Path() {
			t.Fatalf("unexpected import found: %s\n", gi.Path())
		}

		if 0 != len(gi.Comments()) {
			t.Fatalf("unexpected comments got\n")
		}
	})

	t.Run("parse error", func(t *testing.T) {
		t.Parallel()

		var fset *gt.FileSet = gt.NewFileSet()
		var mode ij.ParseMode = ij.ParseModeDefault

		// Malformed Go code with two import paths in one line
		var invalidSrc string = `package main; import "fmt" "log"`

		_, e := mode.ParseString(fset, "invalid.go", invalidSrc)
		if nil == e {
			t.Fatalf("expected a parse error but got nil")
		}
	})
}
