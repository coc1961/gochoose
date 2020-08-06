package choose

import (
	"errors"

	"github.com/coc1961/gochoose/internal/gui"
	"github.com/coc1961/gochoose/internal/gui/keyboard"
)

func New() *Choose {
	return &Choose{
		term: gui.NewTerminal(),
	}
}

type Choose struct {
	term     *gui.Terminal
	selected string
}

func (ch *Choose) SetSelected(s string) {
	ch.selected = s
}

func (ch *Choose) Choose(options []string) (string, error) {
	ind := 0
	exit := make(chan error)
	r1, c1 := ch.term.CursorPos()
	rows, _, _ := ch.term.TerminalSize()

	if ch.selected != "" {
		for i, o := range options {
			if o == ch.selected {
				ind = i
			}
		}
	}
	print := func() {
		ch.term.Cls()
		w := 0
		r, c := 1, 1
		for i := 0; i < len(options); i++ {
			if i == ind {
				ch.term.Goto(r, c).Yellow().Print(options[i] + "            ")
			} else {
				ch.term.Goto(r, c).Reset().Print(options[i] + "            ")
			}
			if len(options[i]) > w {
				w = len(options[i])
			}
			r++
			if r >= rows-1 {
				r = 1
				c += w + 10
				w = 0
			}

		}
		ch.term.Goto(r1, c1).Reset().Flush()
	}
	print()

	ch.term.Keyboard(func(key keyboard.KeyEvent) {
		switch key.Key {
		case keyboard.KeyArrowUp:
			ind--
		case keyboard.KeyArrowDown:
			ind++
		case keyboard.KeyEnter:
			exit <- nil
		case keyboard.KeyEsc:
			exit <- errors.New("esc pressed")
		}

		if ind < 0 {
			ind = 0
		}
		if ind >= len(options) {
			ind = len(options) - 1
		}
		print()
	})

	err := <-exit

	ch.term.Reset().Cls()
	ch.term.Flush()
	ch.term.Close()

	return options[ind], err
}
