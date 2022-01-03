package src

import (
	"os"
	"strings"
)

var (
	confOk   bool = true
	micromod string

	localeOk bool = true
	locale        = make(map[string]string)
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
		case "micromod":
			micromod = left
		case "lang":
			lang = left
		case "langs":
			langs = strings.Split(left, ", ")
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
		switch right {
		case "license":
			locale["license"] = left
		case "text":
			locale["text"] = left
		}
	}
}

func setStdConf() {
	confOk = false
	conf := "micromod: micromod.jar\n" +
		"langs: ru\n" +
		"lang: ru"
	os.WriteFile("config.txt", []byte(conf), 0644)
	readConf()
}

func setStdLocale() {
	localeOk = false
	locale = stdLocale
	local := ""
	for k, e := range locale {
		local += k + ": " + e + "\n"
	}
	os.WriteFile("locale/ru.txt", []byte(local[:len(local)-1]), 0644)
}

var stdLocale = map[string]string{
	"license": "Лицензия",
	"text":    "скоро",
}
