package choose

import (
	gc "github.com/coc1961/gochoose/internal/goncurses"
	"github.com/coc1961/gochoose/internal/keyboard"
)

func New() Choose {
	return Choose{}
}

type Choose struct {
}

func (ch Choose) Choose(options []string) (string, error) {
	if len(options) == 0 {
		return "", nil
	}
	stdscr, err := gc.Init()
	if err != nil {
		return "", err
	}
	gc.Echo(false)
	gc.CBreak(true)
	gc.Cursor(0)
	stdscr.NoutRefresh()
	_ = gc.UseDefaultColors()
	_ = gc.StartColor()

	_ = gc.InitPair(1, gc.C_WHITE, -1)
	_ = gc.InitPair(2, gc.C_YELLOW, -1)

	rows, cols := stdscr.MaxYX()
	var win *gc.Window
	win, err = gc.NewWindow(rows, cols, 0, 0)
	if err != nil {
		return "", err
	}

	_ = gc.Update()
	ind := 0
	print := func() {
		_ = win.Clear()
		w := 0
		r, c := 1, 1
		for i := 0; i < len(options); i++ {
			if i == ind {
				win.Move(r, c)
				_ = win.ColorOn(int16(2))
				win.Print(options[i])
			} else {
				win.Move(r, c)
				_ = win.ColorOn(1)
				win.Print(options[i])
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
		win.Refresh()
	}
	print()

	run := true
	keysEvents, _ := keyboard.GetKeys(10)
	for run {
		event := <-keysEvents
		switch event.Key {
		case keyboard.KeyEnter:
			run = false
			continue
		case keyboard.KeyEsc:
			ind = -1
			run = false
			continue
		case keyboard.KeyArrowDown:
			ind++
		case keyboard.KeyArrowUp:
			ind--
		default:
			continue
		}
		if ind < 0 {
			ind = 0
		}
		if ind >= len(options) {
			ind = len(options) - 1
		}
		print()
	}
	gc.End()
	if ind < 0 {
		return "", nil
	}
	return options[ind], nil
}
