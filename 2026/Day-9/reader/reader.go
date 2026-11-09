// Package reader provides ReadInput utility
package reader

import (
	"fmt"
	"iter"
	"os"
	"strings"
)

func ReadInput(path string) iter.Seq[string] {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}

	return strings.SplitSeq(strings.TrimSuffix(string(data), "\n"), "\n")
}
