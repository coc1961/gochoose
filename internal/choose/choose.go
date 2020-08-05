package choose

import (
	"github.com/coc1961/gochoose/internal/gui"
	"github.com/coc1961/gochoose/internal/gui/keyboard"
)

func New() Choose {
	return Choose{
		term: gui.NewTerminal(),
	}
}

type Choose struct {
	term *gui.Terminal
}

func (ch Choose) Close() {
	ch.term.Cls()
	ch.term.Flush()
	ch.term.Close()
}

func (ch Choose) Choose(options []string) (string, error) {
	ind := 0
	exit := make(chan bool)
	r1, c1 := ch.term.CursorPos()
	rows, _, _ := ch.term.TerminalSize()
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
			exit <- true
		case keyboard.KeyEsc:
			exit <- true
		}

		if ind < 0 {
			ind = 0
		}
		if ind >= len(options) {
			ind = len(options) - 1
		}
		print()
	})

	<-exit

	ch.term.Reset().Cls()
	ch.term.Flush()

	return options[ind], nil
}
