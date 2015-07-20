package slot

import (
	"fmt"
	"io"
	"math/rand"
	"time"
)

type Slot struct {
	symbols []string
	w       io.Writer
}

func New(syms []string, w io.Writer) *Slot {
	s := &Slot{
		symbols: syms,
		w:       w,
	}
	return s
}

func (s *Slot) Start() int {
	rand.Seed(time.Now().UnixNano())

	sym := 0
	for cnt := 1; ; cnt++ {
		i := rand.Intn(len(s.symbols))
		str := s.symbols[i]
		fmt.Fprint(s.w, str)

		if sym != i {
			sym = 0
			continue
		} else {
			sym++
			if sym == len(s.symbols) {
				return cnt
			}
		}
	}
}
