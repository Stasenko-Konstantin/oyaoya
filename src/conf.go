package src

import (
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
	conf := "langs: ru\n" +
		"lang: ru"
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
