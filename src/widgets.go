package src

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"strconv"
)

type Chan struct {
	entry *widget.Table
	num   int
}

type Switcher struct {
	rows int
	cols int
}

func (s *Switcher) stepLeft() {
	for _, s := range canvas {
		for _, e := range s {
			e.FocusLost()
		}
	}
	if currCol == 0 {
		currCol = s.cols
		canvas[currRow][currCol].FocusGained()
	}
}

func (s *Switcher) stepRight() {
	fmt.Println("s")
}

func (s *Switcher) stepUp() {
	fmt.Println("s")
}

func (s *Switcher) stepDown() {
	fmt.Println("s")
}

const (
	rows = 64
	cols = 25
)

var (
	canvas   [][]*widget.Entry
	switcher *Switcher
	currRow  = 0
	currCol  = 0
)

func makeChannels() *widget.Table {
	canvas = make([][]*widget.Entry, rows)
	for i := range canvas {
		canvas[i] = make([]*widget.Entry, cols)
	}
	channels := widget.NewTable(
		func() (int, int) {
			switcher = &Switcher{rows, cols}
			return rows, cols
		},
		func() fyne.CanvasObject {
			entry := container.NewVBox(widget.NewEntry())
			return entry
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			cell := o.(*fyne.Container).Objects
			if i.Col == 0 {
				cell[0] = widget.NewLabel(strconv.Itoa(i.Row))
			} else {
				if i.Col == 1 || (i.Col-1)%6 == 0 { // каждый 6 начиная со 2
					cell[0].(*widget.Entry).SetPlaceHolder("___")
				} else {
					cell[0].(*widget.Entry).SetPlaceHolder("  .")
				}
				canvas[i.Row][i.Col] = cell[0].(*widget.Entry)
			}
		})
	channels.Resize(fyne.NewSize(1200, 370))
	channels.Move(fyne.NewPos(0, 300))
	return channels
}
