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
	"strings"
	"time"
)

func makeInstruments() fyne.CanvasObject {
	var instrumentsArr []fyne.CanvasObject
	files, err := ioutil.ReadDir("samples")
	if err != nil {
		cathcer <- err
	}
	for _, i := range files {
		slice := strings.Split(i.Name(), ".")
		if slice[len(slice)-1] == "wav" {
			instrumentsArr = append(instrumentsArr, container.NewHBox(
				widget.NewButton(">", func(i fs.FileInfo) func() {
					return func() {
						go func() {
							f, err := os.Open("samples/" + i.Name())
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
				widget.NewLabel(slice[0]),
			))
		}
	}
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
