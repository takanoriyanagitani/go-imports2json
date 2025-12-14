// package imports2json provides utilities to parse Go source code and extract
// import declarations.
package imports2json

import (
	"bytes"
	ga "go/ast"
	gp "go/parser"
	gt "go/token"
	"io"
	"strings"
)

// GoFile wraps a `go/ast.File` to provide convenient access to import data.
type GoFile struct{ *ga.File }

// ImportSpecs returns the raw `go/ast.ImportSpec` slice from the parsed file.
func (f GoFile) ImportSpecs() []*ga.ImportSpec { return f.File.Imports }

// GoImport wraps a `go/ast.ImportSpec` for easier access to its components.
type GoImport struct{ *ga.ImportSpec }

// DocComments returns the doc comment group associated with the import.
func (i GoImport) DocComments() []*ga.Comment {
	var doc *ga.CommentGroup = i.ImportSpec.Doc

	switch doc {
	case nil:
		return nil
	default:
		return doc.List
	}
}

// Name returns the local package name (alias) of the import. It returns an
// empty string if no alias is used.
func (i GoImport) Name() string {
	var name *ga.Ident = i.ImportSpec.Name
	switch name {
	case nil:
		return ""
	default:
		return name.Name
	}
}

// Path returns the raw import path string, including quotes.
func (i GoImport) Path() string {
	var path *ga.BasicLit = i.ImportSpec.Path
	switch path {
	case nil:
		return ""
	default:
		return path.Value
	}
}

// PathTrim returns the import path with surrounding quotes removed.
func (i GoImport) PathTrim() string {
	var path string = i.Path()
	return strings.Trim(path, `"`)
}

// Comments returns the line comment group associated with the import.
func (i GoImport) Comments() []*ga.Comment {
	var comment *ga.CommentGroup = i.ImportSpec.Comment

	switch comment {
	case nil:
		return nil
	default:
		return comment.List
	}
}

// GoComment wraps a `go/ast.Comment` for easier access.
type GoComment struct{ *ga.Comment }

// Text returns the text of the comment.
func (c GoComment) Text() string {
	switch c.Comment {
	case nil:
		return ""
	default:
		return c.Comment.Text
	}
}

// ParseMode is a wrapper for `go/parser.Mode`.
type ParseMode gp.Mode

// ParseModeDefault is the default mode for parsing, focusing only on imports.
const ParseModeDefault ParseMode = ParseMode(gp.ImportsOnly)

// ParseReader parses Go source from an `io.Reader`.
func (m ParseMode) ParseReader(fset *gt.FileSet, name string, src io.Reader) (GoFile, error) {
	f, e := gp.ParseFile(fset, name, src, gp.Mode(m))
	return GoFile{File: f}, e
}

// ParseString parses Go source from a string.
func (m ParseMode) ParseString(fset *gt.FileSet, name string, src string) (GoFile, error) {
	var sr io.Reader = strings.NewReader(src)
	return m.ParseReader(fset, name, sr)
}

// ParseBytes parses Go source from a byte slice.
func (m ParseMode) ParseBytes(fset *gt.FileSet, name string, src []byte) (GoFile, error) {
	var br io.Reader = bytes.NewReader(src)
	return m.ParseReader(fset, name, br)
}

// Imports processes the parsed file and returns a slice of `ImportInfo`,
// which is a simplified, serializable representation of the import declarations.
func (f GoFile) Imports() []ImportInfo {
	var specs []*ga.ImportSpec = f.ImportSpecs()
	var infos []ImportInfo = make([]ImportInfo, 0, len(specs))
	for _, spec := range specs {
		var imp GoImport = GoImport{ImportSpec: spec}
		infos = append(infos, ImportInfo{
			Path: imp.PathTrim(),
			Name: imp.Name(),
		})
	}
	return infos
}
