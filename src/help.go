package src

import (
	"os"
	"strconv"
	"strings"
)

var (
	license = "\n\toyaoya - tracker music editor\n" +
		"\tCopyright (C) 2021  Stasenko Konstantin\n" + "\n\n" +
		"\tThis program is free software: you can redistribute it and/or modify\n" +
		"\tit under the terms of the GNU General Public License as published by\n" +
		"\tthe Free Software Foundation, either version 3 of the License, or\n" +
		"\t(at your option) any later version.\n" + "\n\n" +
		"\tThis program is distributed in the hope that it will be useful,\n" +
		"\tbut WITHOUT ANY WARRANTY; without even the implied warranty of\n" +
		"\tMERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the\n" +
		"\tGNU General Public License for more details.\n" + "\n\n" +
		"\tYou should have received a copy of the GNU General Public License\n" +
		"\talong with this program.  If not, see <http://www.gnu.org/licenses/>.\n" + "\n\n" +
		"\tcontacts:\n" +
		"\t    mail - stasenko.ky@gmail.com\n" +
		"\t    github - Stasenko-Konstantin\n\n"
	instruction = "" // start.go getInstruction()
)

var (
	title = `
(temp)
Module "temp"
	Channels 4
	Sequence `

	instruments string
)

func wrapStr(str string) string {
	if len(str) == 1 {
		return " " + str + " "
	}
	return str
}

func strToNum(str string) int {
	str = strings.TrimSpace(str)
	n, _ := strconv.Atoi(str)
	return n
}

func checkFormat(format string) bool {
	r := true
	if format == "mod" || format == "mt" {
		r = false
	}
	return r
}

func addTxt(lang string) string {
	var r string
	for _, c := range lang {
		if c != '\r' {
			r += string(c)
		}
	}
	for _, c := range ".txt" {
		r += string(c)
	}
	return r
}

func findDir(dirs []os.DirEntry, dir string) bool {
	for _, file := range dirs {
		if file.IsDir() && file.Name() == dir {
			return true
		}
	}
	return false
}

func checkStars(end string) (cont bool, start string, length string) {
	start = end[:3]
	length = end[3:]
	_, err := strconv.Atoi(start)
	_, err = strconv.Atoi(length)
	if err != nil {
		cont = true
	}
	return
}