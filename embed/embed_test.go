package embed

import (
	_ "embed"
	"testing"
)

//go:embedversion.txt
var version string

//go:embedversion.txt
var version2 string

func TestString(t *testing.T) {

}
