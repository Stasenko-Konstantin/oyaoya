package src

import (
	"strconv"
	"strings"
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
