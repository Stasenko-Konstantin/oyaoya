package src

import (
	"fmt"
	"fyne.io/fyne/v2"
	theme2 "fyne.io/fyne/v2/theme"
	"os"
	"os/exec"
	"strings"
)

var (
	confOk = true

	localeOk = true
	locale   = make(map[string]string)
	lang     string
)

func readConf() {
	defer func() {
		if r := recover(); r != nil {
			confOk = false
			fmt.Println(r.(error).Error())
			fmt.Println("readConf 1")
			setStdConf()
		}
	}()
	text, err := os.ReadFile("config.txt")
	if err != nil {
		confOk = false
		fmt.Println(err.Error())
		fmt.Println("readConf 2")
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
AGAIN:
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r.(error).Error())
			fmt.Println("readLocale 1")
			localeOk = false
			setStdLocale()
		}
	}()
	text, err := os.ReadFile(addTxt("locale/" + lang))
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("readLocale 2")
		localeOk = false
		setStdLocale()
		goto AGAIN
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
	conf := `lang: ru
theme: color
`
	os.WriteFile("config.txt", []byte(conf), 0644)
	readConf()
}

func setStdLocale() {
	localeOk = false
	locale := `menu: Меню
new: Новая композиция
open: Открыть
save: Сохранить
save as: Сохранить как
settings: Настройки
license: Лицензия
about: Об авторе
help: Помощь
instruction: Инструкция
patterns: Паттерны
choice pattern: Выберите паттерн
cancel: Отмена
wrong format: Неверный формат! Попробуйте .mt или .mod
instruction text: locale/instruction-ru.txt
dark theme: Темная тема
light theme: Светлая тема
color theme: Цветная тема
instruments: Инструменты
about text: Курсант 432 гр. Стасенко К.Ю.`
	instruction = `Ноты имеют форму 'KKOIIFPP', где:

    KK = Ключ (один из 'C-', 'C#', 'D-', 'D#', 'E-', 'F-', 'F#', 'G-', 'G#', 'A-', 'A#', 'B-', или '--').
    O  = Октава (от 0 до 6, или '-').
    II = Инструмент (десятичное число от 1 до 99, или '--').
    F  = Эффект (шестнадцатеричное число от 0 до F, или '-').
    PP = Параметр эффекта (шестнадцатеричное число от 00 до FF, или '--').
`
	dir, _ := os.ReadDir("./")
	if !findDir(dir, "locale") {
		cmd := exec.Command("mkdir", "locale")
		stdout, err := cmd.Output()
		if err != nil {
			cathcer <- err
		}
		fmt.Println(string(stdout))
	}
	os.WriteFile("locale/ru.txt", []byte(locale), 0644)
	os.WriteFile("locale/instruction-ru.txt", []byte(instruction), 0644)
	setLang("ru")
}

func getInstruction(file string) (string, error) {
	var r string
	var path string
	for _, c := range file {
		if c != '\r' {
			path += string(c)
		}
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	r = string(data)
	return r, nil
}

func setConfigField(field, right string) {
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
	for i, e := range conf {
		parts := strings.Split(e, ":")
		left := parts[0]
		for _, l := range left {
			if l == ' ' {
				left = (left)[1:]
			}
			break
		}
		if left == field {
			conf[i] = left + ": " + right
		}
	}
	err = os.WriteFile("config.txt", []byte(strings.Join(conf, "\n")), 0644)
	if err != nil {
		cathcer <- err
	}
}
