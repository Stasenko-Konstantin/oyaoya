package src

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"os"
	"os/exec"
	"strconv"
)

var (
	isPlay = false
)

func takeContent(pattern, row, col int) string { // ужас
	switch col {
	case 0:
		col = 2
	case 1:
		col = 8
	case 2:
		col = 14
	case 3:
		col = 20
	}
	var r string
	for i := 0; i < 5; i++ {
		t := channelsSelect[pattern].channel.Objects[0].(*container.Scroll).Content.(*fyne.Container).Objects[col+i].(*fyne.Container).Objects[row].(*widget.Entry).Text
		if i%2 == 0 {
			if len(t) == 0 {
				t = "--"
			} else if len(t) == 1 {
				switch i {
				case 0:
					t += "-"
				case 2:
					t = "-" + t
				case 4:
					t = "0" + t
				}
			}
		} else {
			if len(t) == 0 {
				t = "-"
			}
		}
		r += t
	}
	return r
}

func playTrack(play *widget.Button, current bool) {
	needPlay := false
	play.SetText("||")
	sequenceStr, sequenceSlice := getSequence(current)
	temp := title + sequenceStr + instruments
	for _, s := range sequenceSlice {
		temp += "\n\tPattern " + strconv.Itoa(s) + " (play)"
		for row := 0; row < 64; row++ {
			row := "\n\t\tRow \"" + takeContent(s, row, 0) + " " + takeContent(s, row, 1) +
				" " + takeContent(s, row, 2) + " " + takeContent(s, row, 3) + "\""
			if row != "\n\t\tRow \"-------- -------- -------- --------\"" {
				needPlay = true
			}
			temp += row
		}
	}
	if !needPlay {
		play.SetText(">")
		return
	}
	err := os.WriteFile("temp.mt", []byte(temp+"\n(End)"), 0644)
	if err != nil {
		cathcer <- err
	}
	cmd := exec.Command("java", "-jar", "micromod.jar", "temp.mt")
	stdout, _ := cmd.Output()
	fmt.Println(string(stdout))
	play.SetText(">")
	isPlay = false
}

func makePlay() fyne.CanvasObject {
	play := widget.NewButton(">", func() {})
	play.OnTapped = func() {
		fmt.Println(isPlay)
		if !isPlay {
			go playTrack(play, false)
		}
	}
	play.Resize(fyne.NewSize(33, 37))
	play.Move(fyne.NewPos(975, 10))
	playCurr := widget.NewButton(">", func() {})
	playCurr = widget.NewButton(">", func() {
		if !isPlay {
			go playTrack(playCurr, true)
		}
	})
	playCurr.Resize(fyne.NewSize(33, 37))
	playCurr.Move(fyne.NewPos(0, 210))
	return container.NewWithoutLayout(play, playCurr)
}
