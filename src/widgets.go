package src

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type chann struct {
	entry *widget.Entry
	num   int
}

var (
	channs    []chann
	currChann = 0
)

func makeChannels() *fyne.Container {
	channels := container.NewWithoutLayout()
	var diff float32 = 0
	for i := 0; i < 4; i++ {
		channel := widget.NewMultiLineEntry()
		channel.Resize(fyne.NewSize(300, 900))
		channel.Move(fyne.NewPos(diff, 300))
		channels.Add(channel)
		channs = append(channs, chann{channel, i})
		diff += 305
	}
	return channels
}

func unfocusChanns() {
	for _, c := range channs {
		c.entry.FocusLost()
	}
}
