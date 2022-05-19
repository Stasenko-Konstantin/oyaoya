package src

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"strconv"
)

type Pattern string // !!!!!!!!!

type Patterns struct {
	scroll   *container.Scroll
	border   int
	patterns map[int]Pattern
}

var (
	patterns      Patterns
	patternsCount = 10
	songPatterns  *container.Scroll
)

func makePatterns() []fyne.CanvasObject {
	var channelsArr []fyne.CanvasObject
	for i := 0; i < patternsCount; i++ {
		button := func(i int) *widget.Button {
			return widget.NewButton(wrapStr(strconv.Itoa(i)), func() {
				hideAll()
				switch i {
				case channelsSelect[0].index:
					channelsSelect[0].channel.Show()
				case channelsSelect[1].index:
					channelsSelect[1].channel.Show()
				case channelsSelect[2].index:
					channelsSelect[2].channel.Show()
				case channelsSelect[3].index:
					channelsSelect[3].channel.Show()
				case channelsSelect[4].index:
					channelsSelect[4].channel.Show()
				case channelsSelect[5].index:
					channelsSelect[5].channel.Show()
				case channelsSelect[6].index:
					channelsSelect[6].channel.Show()
				case channelsSelect[7].index:
					channelsSelect[7].channel.Show()
				case channelsSelect[8].index:
					channelsSelect[8].channel.Show()
				case channelsSelect[9].index:
					channelsSelect[9].channel.Show()
				}
			})
		}(i)
		channelsArr = append(channelsArr, button)
	}
	patterns.scroll = container.NewHScroll(container.NewHBox(channelsArr...))
	var songPatternsArr []fyne.CanvasObject
	for i := 0; i < patternsCount; i++ {
		var button *widget.Button
		button = widget.NewButton(" + ", func() {
			var patternsChoice []fyne.CanvasObject
			for i := 0; i < patternsCount; i++ {
				patternsChoice = append(patternsChoice, func(i int) *widget.Button {
					return widget.NewButton(strconv.Itoa(i), func() {})
				}(i))
			}
			d := dialog.NewCustom(locale["choice pattern"], locale["cancel"], container.NewVBox(patternsChoice...),
				*window)
			for i := 0; i < patternsCount; i++ {
				patternsChoice[i] = func(i int) *widget.Button {
					return widget.NewButton(strconv.Itoa(i), func() {
						button.SetText(wrapStr(strconv.Itoa(i)))
						d.Hide()
					})
				}(i)
			}
			d.Resize(fyne.NewSize(100, 400))
			d.Show()
		})
		songPatternsArr = append(songPatternsArr, button)
	}
	songPatterns = container.NewHScroll(container.NewHBox(songPatternsArr...))
	label := widget.NewLabel(locale["patterns"])
	return []fyne.CanvasObject{
		label,
		patterns.scroll,
		songPatterns,
	}
}
