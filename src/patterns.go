package src

import (
	"errors"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"strconv"
	"strings"
)

type Patterns struct {
	scroll   *container.Scroll
	border   int
	patterns map[int]string
}

var (
	patterns      Patterns
	patternsCount = 10
	songPatterns  *container.Scroll
)

func getSequence(current bool) (string, map[int]int) {
	var sequenceStr string
	sequenceSet := make(map[int]int)
	if !current {
		for i := 0; i < 10; i++ {
			n := songPatterns.Content.(*fyne.Container).Objects[i].(*widget.Button).Text
			if !strings.Contains(n, "+") {
				num := strToNum(n)
				sequenceSet[num] = num
				sequenceStr += n + ","
			}
		}
	}
	if sequenceStr == "" {
		sequenceStr = strconv.Itoa(currPattern)
		sequenceSet[currPattern] = currPattern
		return sequenceStr, sequenceSet
	}
	sequenceStr = strings.ReplaceAll(sequenceStr, " ", "")
	return sequenceStr[:len(sequenceStr)-1], sequenceSet
}

func makePatterns() fyne.CanvasObject {
	var channelsArr []fyne.CanvasObject
	for i := 0; i < patternsCount; i++ {
		button := func(i int) *widget.Button {
			return widget.NewButton(wrapStr(strconv.Itoa(i)), func() {
				hideAll()
				channelsSelect[i].channel.Show()
				currPattern = i
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
			patternsChoice = append(patternsChoice, widget.NewButton("-", func() {}))
			d := dialog.NewCustom(locale["choose pattern"], locale["cancel"], container.NewVBox(patternsChoice...),
				*window)
			for i := 0; i < patternsCount; i++ {
				patternsChoice[i] = func(i int) *widget.Button {
					return widget.NewButton(strconv.Itoa(i), func() {
						button.SetText(wrapStr(strconv.Itoa(i)))
						d.Hide()
					})
				}(i)
			}
			patternsChoice[patternsCount] = widget.NewButton("-", func() {
				button.SetText(" + ")
				d.Hide()
			})
			d.Resize(fyne.NewSize(100, 400))
			d.Show()
		})
		songPatternsArr = append(songPatternsArr, button)
	}
	songPatterns = container.NewHScroll(container.NewHBox(songPatternsArr...))
	label := widget.NewLabel(locale["patterns"])
	patterns.scroll.SetMinSize(fyne.NewSize(500, 20))
	return container.NewVBox(
		label,
		patterns.scroll,
		songPatterns,
	)
}

func findPattern(pattern int) (int, error) {
	fmt.Println(pattern)
	for _, p := range songPatterns.Content.(*fyne.Container).Objects {
		if p.(*widget.Button).Text == wrapStr(strconv.Itoa(pattern)) {
			return pattern, nil
		}
	}
	return pattern, errors.New("")
}
