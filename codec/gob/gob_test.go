package gob

import (
	"testing"

	"github.com/covrom/storm/codec/internal"
)

func TestGob(t *testing.T) {
	internal.RoundtripTester(t, Codec)
}
