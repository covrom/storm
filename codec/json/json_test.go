package json

import (
	"testing"

	"github.com/covrom/storm/codec/internal"
)

func TestJSON(t *testing.T) {
	internal.RoundtripTester(t, Codec)
}
