package msgpack

import (
	"testing"

	"github.com/covrom/storm/codec/internal"
)

func TestMsgpack(t *testing.T) {
	internal.RoundtripTester(t, Codec)
}
