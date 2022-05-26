package src

import (
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

	instruments = `	
	Instrument 1 Name "01"
		Volume 24 FineTune 0
		WaveFile "samples/01.wav"
		LoopStart 254 LoopLength 506
	Instrument 2 Name "02"
		Volume 32 FineTune 0
		WaveFile "samples/02.wav"
		LoopStart 334 LoopLength 444
	Instrument 3 Name "03"
		Volume 64 FineTune 0
		WaveFile "samples/03.wav"
	Instrument 4 Name "04"
		Volume 64 FineTune 0
		WaveFile "samples/04.wav"
	Instrument 5 Name "05"
		Volume 40 FineTune 0
		WaveFile "samples/05.wav"
		LoopStart 324 LoopLength 508
	Instrument 6 Name "06"
		Volume 64 FineTune 0
		WaveFile "samples/06.wav"
	Instrument 7 Name "07"
		Volume 64 FineTune 0
		WaveFile "samples/07.wav"
	Instrument 8 Name "08"
		Volume 64 FineTune 0
		WaveFile "samples/08.wav"
	Instrument 9 Name "09"
		Volume 64 FineTune 0
		WaveFile "samples/09.wav"
	Instrument 10 Name "10"
		Volume 64 FineTune 0
		WaveFile "samples/10.wav"`
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
