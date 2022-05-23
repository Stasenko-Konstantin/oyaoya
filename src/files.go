package src

import (
	"errors"
	"fyne.io/fyne/v2"
	"os"
	"strconv"
	"strings"
)

func setNewSong() {
	
}

func openSong(path fyne.URIReadCloser) {
	name := strings.Split(path.URI().Name(), ".")
	if len(name) < 2 || checkFormat(name[1]) {
		cathcer <- errors.New(locale["wrong format"])
		return
	}
	if name[1] == "mod" {
		openMod(path)
	} else {
		openMt(path)
	}
}

func openMod(path fyne.URIReadCloser) {

}

func openMt(path fyne.URIReadCloser) {
	defer func() {
		if r := recover(); r != nil {
			cathcer <- r.(error)
			setNewSong()
		}
	}()
	defer path.Close()
	data, err := os.ReadFile(path.URI().Path())
	if err != nil {
		cathcer <- err
		return
	}
	content := string(data)
	// contentSlice := strings.Fields(content)
	module := func() string {
		var r string
		for i := 0; true; i++ {
			if content[i] == 'M' {
				if content[i:i+6] == "Module" {
					r = strings.Split(content[i+8:], "\n")[0]
					break
				}
			}
		}
		return r[:len(r)-1]
	}()
	sequence := func() []int {
		var sequence []int
		var s string
		for i := 0; true; i++ {
			if content[i] == 'S' {
				if content[i:i+8] == "Sequence" {
					s = strings.Split(content[i+8:], "\n")[0]
					break
				}
			}
		}
		s = strings.ReplaceAll(s, " ", "")
		for _, p := range strings.Split(s, ",") {
			t, err := strconv.Atoi(p)
			if err != nil {
				cathcer <- err
			}
			sequence = append(sequence, t)
		}
		if len(sequence) > 10 {
			sequence = sequence[9:]
		}
		return sequence
	}

	(*window).SetTitle(module)
	setSequence(sequence())
}

func setSequence(sequence []int) {

}
