package src

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"strconv"
)

type Switcher struct {
	rows int
	cols int
}

func (s *Switcher) stepLeft() {
	fmt.Println("s")
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
	cols = 24
)

var (
	canvas   [][]fyne.CanvasObject
	switcher *Switcher
)

func makeChannels() *container.Scroll {
	canvas = make([][]fyne.CanvasObject, cols)
	for i := range canvas {
		canvas[i] = make([]fyne.CanvasObject, rows)
		for j := range canvas[i] {
			entry := widget.NewEntry()
			canvas[i][j] = entry
			entry.SetPlaceHolder("  .")
			if i%6 == 0 {
				entry.SetPlaceHolder("___")
			}
		}
	}
	var nums []fyne.CanvasObject
	for i := 0; i < rows; i++ {
		nums = append(nums, widget.NewLabel(strconv.Itoa(i)))
	}
	var spaces []fyne.CanvasObject
	for i := 0; i < rows; i++ {
		spaces = append(spaces, widget.NewLabel(" "))
	}
	var containers []fyne.CanvasObject
	containers = append(containers, container.NewVBox(nums...))
	containers = append(containers, container.NewVBox(spaces...))
	for i := 0; i < cols; i++ {
		entrys := canvas[i]
		if i%6 == 0 && i != 0 {
			containers = append(containers, container.NewVBox(spaces...))
		}
		containers = append(containers, container.NewVBox(entrys...))
	}
	channels := container.NewScroll(container.NewHBox(containers...))
	channels.Resize(fyne.NewSize(1200, 370))
	channels.Move(fyne.NewPos(0, 300))
	return channels
}
