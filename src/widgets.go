package src

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"strconv"
)

type chann struct {
	entry *widget.Table
	num   int
}

var (
	channs    []chann
	currChann = 0
)

func makeChannels() *widget.Table {
	channels := widget.NewTable(
		func() (int, int) {
			return 64, 25
		},
		func() fyne.CanvasObject {
			entry := container.NewVBox(widget.NewEntry())
			return entry
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			if i.Col == 0 {
				o.(*fyne.Container).Objects[0] = widget.NewLabel(strconv.Itoa(i.Row))
			} else if i.Col == 1 || (i.Col-1)%6 == 0 {
				o.(*fyne.Container).Objects[0].(*widget.Entry).SetPlaceHolder("___")
			} else {
				o.(*fyne.Container).Objects[0].(*widget.Entry).SetPlaceHolder("  .")
			}
		})
	channels.Resize(fyne.NewSize(1200, 400))
	channels.Move(fyne.NewPos(0, 300))
	return channels
}
