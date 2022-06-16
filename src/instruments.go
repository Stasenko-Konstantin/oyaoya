package src

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	_ "github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
	"io/fs"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

func makeInstruments(path string) fyne.CanvasObject {
	var (
		instrumentsArr []fyne.CanvasObject
		instr          string
	)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		cathcer <- err
	}
	for _, i := range files {
		slice := strings.Split(i.Name(), ".")
		if slice[len(slice)-1] == "wav" {
			parts := strings.Split(i.Name()[:len(i.Name())-4], "#")
			if len(parts) != 3 {
				continue
			}
			instr += "\n\tInstrument "
			_, err := strconv.Atoi(parts[0])
			if err != nil {
				continue
			}
			instr += parts[0] + " Name \"" + parts[1] + "\"\n\t\tVolume 64 FineTune 0\n\t\tWaveFile \"" + path + "/" + i.Name() + "\""
			if cont, start, length := checkStars(parts[2]); !cont {
				if strings.Contains(start, "-") {
					start = strings.ReplaceAll(start, "-", "")
				}
				if strings.Contains(length, "-") {
					length = strings.ReplaceAll(length, "-", "")
				}
				instr += "\n\t\tLoopStart " + start + " LoopLength " + length
			}
			name := parts[0] + " " + parts[1]
			instrumentsArr = append(instrumentsArr, container.NewHBox(
				widget.NewButton(">", func(i fs.FileInfo) func() {
					return func() {
						go func() {
							f, err := os.Open(path + "/" + i.Name())
							if err != nil {
								cathcer <- err
							}
							streamer, format, err := wav.Decode(f)
							if err != nil {
								cathcer <- err
							}
							defer streamer.Close()
							speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
							speaker.Play(streamer)
							select {}
						}()
					}
				}(i)),
				widget.NewLabel(name),
			))
		}
	}
	instruments = instr
	instruments := container.NewVScroll(container.NewVBox(instrumentsArr...))
	instruments.SetMinSize(fyne.NewSize(440, 150))
	label := widget.NewLabel(locale["instruments"])
	r := container.NewVBox(
		label,
		instruments,
	)
	r.Move(fyne.NewPos(490, r.Position().Y))
	return r
}
