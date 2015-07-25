package slot

import (
	"bytes"
	"reflect"
	"strings"
	"testing"
)

func TestSlotStart(t *testing.T) {
	assert := func(expected []string) {
		w := bytes.NewBuffer([]byte{})
		s := New(expected, w)
		s.Start()

		b := w.Bytes()
		var got []string
		for _, v := range b[len(b)-4:] {
			got = append(got, string(v))
		}
		if !reflect.DeepEqual(got, expected) {
			t.Fatalf("Output should be end with %v, but got %v", expected, got)
		}

		if c := strings.Count(string(b), strings.Join(expected, "")); c != 1 {
			t.Fatalf("Should be 1, but got %d", c)
		}
	}

	for i := 0; i < 50; i++ {
		assert([]string{"h", "o", "g", "e"})
	}
}
