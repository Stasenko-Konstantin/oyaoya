package src

import (
	"errors"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"os"
	"strconv"
	"strings"
)

func setNewSong() {
	(*window).SetTitle("oyaoya")
	for _, c := range channelsSelect {
		for _, col := range c.channel.Objects[0].(*container.Scroll).Content.(*fyne.Container).Objects {
			for _, cell := range col.(*fyne.Container).Objects {
				switch v := cell.(type) {
				case *widget.Entry:
					v.SetText("")
				}
			}
		}
	}
	for _, p := range songPatterns.Content.(*fyne.Container).Objects {
		p.(*widget.Button).SetText(" + ")
	}
}

func openSong(path fyne.URIReadCloser) {
	if path == nil {
		return
	}
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
	i := 0
	defer func() {
		if r := recover(); r != nil {
			setNewSong()
			i = 1
			return
		}
	}()
	fmt.Println(i)
	for i < 1 {
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
		i++
	}
}

func setSequence(sequence []int) {

}
