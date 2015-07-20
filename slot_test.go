package slot

import (
	"bytes"
	"reflect"
	"testing"
)

func TestSlotStart(t *testing.T) {
	w := bytes.NewBuffer([]byte{})
	expected := []string{"h", "o", "g", "e"}
	s := New(expected, w)
	s.Start()

	b := w.Bytes()
	var got []string
	for _, v := range b[len(b)-4:] {
		got = append(got, string(v))
	}
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Output should be end with %v, but got %v", expected, got)
	}
}
