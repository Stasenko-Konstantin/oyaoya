package src

import (
	"fyne.io/fyne/v2"
	theme2 "fyne.io/fyne/v2/theme"
	"os"
	"strings"
)

var (
	confOk = true

	localeOk = true
	locale   = make(map[string]string)
	lang     string
	langs    []string
)

func readConf() {
	defer func() {
		if r := recover(); r != nil {
			setStdConf()
		}
	}()
	text, err := os.ReadFile("config.txt")
	if err != nil {
		setStdConf()
	}
	conf := strings.Split(string(text), "\n")
	for _, e := range conf {
		parts := strings.Split(e, ":")
		right := parts[0]
		left := parts[1]
		for _, l := range left {
			if l == ' ' {
				left = left[1:]
			}
			break
		}
		switch right {
		case "lang":
			lang = left
		case "langs":
			langs = strings.Split(left, ", ")
		case "theme":
			switch left {
			case "dark":
				fyne.CurrentApp().Settings().SetTheme(theme2.DarkTheme())
				fon.Hide()
			case "light":
				fyne.CurrentApp().Settings().SetTheme(theme2.LightTheme())
				fon.Hide()
			default:
				fyne.CurrentApp().Settings().SetTheme(theme2.DarkTheme())
				fon.Show()
			}
		}
	}
}

func readLocale() {
	defer func() {
		if r := recover(); r != nil {
			setStdLocale()
		}
	}()
	text, err := os.ReadFile("locale/" + lang + ".txt")
	if err != nil {
		setStdLocale()
	}
	local := strings.Split(string(text), "\n")
	for _, e := range local {
		parts := strings.Split(e, ":")
		right := parts[0]
		left := parts[1]
		for _, l := range left {
			if l == ' ' {
				left = left[1:]
			}
			break
		}
		locale[right] = left
	}
}

func setStdConf() {
	confOk = false
	conf := `langs: ru
lang: ru
theme: color
`
	os.WriteFile("config.txt", []byte(conf), 0644)
	readConf()
}

func setStdLocale() {
	localeOk = false
	locale = make(map[string]string)
	local := ""
	for k, e := range locale {
		local += k + ": " + e + "\n"
	}
	os.WriteFile("locale/ru.txt", []byte(local[:len(local)-1]), 0644)
}

func getInstruction() string {
	var r string
	data, err := os.ReadFile(locale["instruction text"])
	if err != nil {
		return err.Error()
	}
	r = string(data)
	return r
}
