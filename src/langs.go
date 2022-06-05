package src

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"os"
)

func makeLangs() []fyne.CanvasObject {
	var langsArr []fyne.CanvasObject
	var langs []string
	dir, _ := os.ReadDir("locale")
	for _, file := range dir {
		name := file.Name()
		if len(name) == 6 && name[2:] == ".txt" {
			langs = append(langs, name[:2])
		}
	}
	for _, lang := range langs {
		_, err := getInstruction(addTxt("locale/instruction-" + lang))
		if err != nil {
			continue
		}
		langsArr = append(langsArr, func(lang string) *widget.Button {
			return widget.NewButton(lang, func() {
				setLang(lang)
			})
		}(lang))
	}
	return langsArr
}

func setLang(lang string) {
	go func() {
		setConfigField("lang", lang)
		readConf()
		readLocale()
	}()
	dialog.ShowInformation(locale["change lang"], locale["change lang msg"], *window)
}
