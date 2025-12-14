package main

import (
	"encoding/json"
	"log"
	"os"

	ij "github.com/takanoriyanagitani/go-imports2json"
)

func main() {
	log.SetFlags(0)

	var mode ij.ParseMode = ij.ParseModeDefault
	file, e := mode.ParseStdinDefault()
	if nil != e {
		log.Fatalf("failed to parse stdin: %v", e)
	}

	var imports []ij.ImportInfo = file.Imports()

	var encoder *json.Encoder = json.NewEncoder(os.Stdout)
	e = encoder.Encode(imports)
	if nil != e {
		log.Fatalf("failed to encode imports to json: %v", e)
	}
}
