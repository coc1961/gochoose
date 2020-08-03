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
	ch.term.Close()
}

func (ch Choose) Choose(options []string) (string, error) {
	ind := 0
	exit := make(chan bool)
	r1, c1 := ch.term.CursorPos()
	r, c := 1, 1

	print := func() {
		ch.term.Cls()
		for i := 0; i < len(options); i++ {
			if i == ind {
				ch.term.Goto(r+i, c).Yellow().Print(options[i] + "            ")
			} else {
				ch.term.Goto(r+i, c).Reset().Print(options[i] + "            ")
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
	return options[ind], nil
}
