package belajar_golang_embed

import (
	_ "embed"
	"fmt"
	"testing"
)

//go:embed version.txt
var version string

func TestEmbedToString(t *testing.T) {
	fmt.Println(version)
}
