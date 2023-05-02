package main

import (
	"os"
	"path/filepath"
	"verifier/pkg/cmd/verifier"
)

func main() {
	basename := filepath.Base(os.Args[0])
	verifier.Execute(basename)
}
