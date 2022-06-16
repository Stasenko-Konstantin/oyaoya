package src

import (
	"errors"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

func setNewSong() {
	namer = 0
	needSave = true
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
	widgets.Objects[2] = makeInstruments("samples")
	widgets.Refresh()
}

func openSong(path, file string) {
	clean()
	name := strings.Split(file, ".")
	if len(name) < 2 || name[1] != "mod" {
		cathcer <- errors.New(locale["wrong format"])
		return
	}
	fmt.Println(path)
	needSave = false
	savePath = path
	openMod(path)
}

func openMod(path string) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		cathcer <- err
		return
	}
	f, err := os.Create("temp.mod")
	if err != nil {
		cathcer <- err
		return
	}
	_, err = f.Write(data)
	if err != nil {
		cathcer <- err
		return
	}

	cmd := exec.Command("java", "-jar", "micromod.jar", "-mod", "temp.mod", "-dir", "temp")
	stdout, _ := cmd.Output()
	fmt.Println(string(stdout))
	normalize()
	openMt("temp/module.mt")
}

func normalize() {
	data, err := os.ReadFile("temp/module.mt")
	if err != nil {
		cathcer <- err
	}
	song := string(data)
	var t string
	for _, s := range strings.Split(song, "\n")[3:] {
		if strings.Contains(s, "Pattern") {
			break
		}
		t += s + "\n"
	}
	song = t
	instrs := strings.Split(song, "Instrument")
	for i := len(instrs) - 1; i >= 0; i-- {
		if len(strings.Split(instrs[i], "\n")) < 3 {
			if i >= len(instrs) {
				instrs = instrs[:i-1]
			} else {
				instrs = append(instrs[:i], instrs[i+1:]...)
			}
		}
	}
	for idx, i := range instrs {
		var (
			name   string
			number string
			start  string
			length string
			file   string
		)
		for j := range i {
			if i[j:j+4] == "Name" {
				for _, c := range i[j+6:] {
					if c == '"' {
						break
					}
					name += string(c)
				}
				break
			}
		}
		name = strings.ReplaceAll(name, "/", "-")
		name = strings.ReplaceAll(name, "\\", "-")
		if name == "" {
			name = genName()
		}
		number = strconv.Itoa(idx + 1)
		if len(number) < 2 {
			number = "0" + number
		}
		if strings.Contains(i, "LoopStart") {
			for j := range i {
				if i[j:j+9] == "LoopStart" {
					for _, c := range i[j+10:] {
						if c == ' ' {
							break
						}
						start += string(c)
					}
					break
				}
			}
		} else {
			start = "***"
		}
		if len(start) < 3 {
			start = strings.Repeat("-", 3-len(start)) + start
		}
		if strings.Contains(i, "LoopLength") {
			for j := range i {
				if i[j:j+10] == "LoopLength" {
					for _, c := range i[j+11:] {
						length += string(c)
					}
					break
				}
			}
		} else {
			length = "***"
		}
		length = strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(length[:3], " ", ""), "\t", ""), "\n", ""), "\r", "")
		if len(length) < 3 {
			length = strings.Repeat("-", 3-len(length)) + length
		}
		if strings.Contains(i, "WaveFile") {
			for j := range i {
				if i[j:j+8] == "WaveFile" {
					for _, c := range i[j+10:] {
						if c == '"' {
							break
						}
						file += string(c)
					}
					break
				}
			}
		}
		data, err := ioutil.ReadFile("temp/" + file)
		if err != nil {
			cathcer <- err
			return
		}
		f, err := os.Create(filepath.Join("temp/", number+"|"+name+"|"+start+length+".wav"))
		if err != nil {
			cathcer <- err
			return
		}
		_, err = f.Write(data)
		if err != nil {
			cathcer <- err
			return
		}
	}
	widgets.Objects[2] = makeInstruments("temp")
	widgets.Refresh()
}

func openMt(path string) {
	i := 0
	defer func() {
		if r := recover(); r != nil {
			setNewSong()
			cathcer <- r.(error)
			i = 1
			return
		}
	}()
	for i < 1 {
		data, err := os.ReadFile(path)
		if err != nil {
			cathcer <- err
			return
		}
		content := string(data)
		for i := 0; true; i++ {
			if content[i:i+8] == "Channels" {
				channels, _ := strconv.Atoi(string(content[i+9]))
				if channels != 4 {
					dialog.ShowInformation(locale["error"], locale["channels error"], *window)
					setNewSong()
					return
				} else {
					break
				}
			}
		}
		module := func() string {
			var r string
			for i := 0; true; i++ {
				if content[i:i+6] == "Module" {
					r = strings.Split(content[i+8:], "\n")[0]
					break
				}
			}
			return r[:len(r)-1]
		}()
		sequence := func() []string {
			var sequence []string
			var s string
			for i := 0; true; i++ {
				if content[i:i+8] == "Sequence" {
					s = strings.Split(content[i+8:], "\n")[0]
					break
				}
			}
			s = strings.ReplaceAll(s[:len(s)-1], " ", "")[1:]
			for _, p := range strings.Split(s, ",") {
				if p == "" {
					continue
				}
				_, err := strconv.Atoi(p)
				if err != nil {
					cathcer <- err
				}
				sequence = append(sequence, p)
			}
			return sequence
		}

		(*window).SetTitle(module)
		seq := sequence()
		if len(seq) > 10 {
			dialog.ShowConfirm(locale["transform"], locale["transform msg"], func(b bool) {
				if !b {
					setNewSong()
					return
				}
			}, *window)
			seq = seq[:10]
		}
		var patterns []string
		for i := range content {
			if content[i:i+7] == "Pattern" {
				patterns = strings.Split(content[i:], "Pattern ")
				break
			}
		}
		if patterns[0] == "" {
			patterns = patterns[1:]
		}
		for i := 0; i < len(patterns); i++ {
			patterns[i] = "Pattern " + patterns[i]
		}
		seq = normilizeSeq(seq, patterns)
		for i := 0; i < 10; i++ {
			songPatterns.Content.(*fyne.Container).Objects[i].(*widget.Button).SetText(" + ")
		}
		for i, p := range seq {
			p := p
			if p == "" {
				p = "+"
			}
			songPatterns.Content.(*fyne.Container).Objects[i].(*widget.Button).SetText(wrapStr(p))
		}
		for _, p := range patterns {
			rows := strings.Split(p, "\n")
			pattern, _ := strconv.Atoi(strings.TrimPrefix(rows[0], "Pattern "))
			pattern, err = findPattern(pattern)
			if err != nil {
				continue
			}
			rows = rows[1:]
			for row, r := range rows {
				if len(r) < 20 {
					break
				}
				r = strings.TrimPrefix(r, "\t\tRow \"")
				r = strings.TrimSuffix(r, "\"")
				channels := strings.Split(r, " ")
				if len(channels) > 4 {
					channels = channels[1:]
				}
				var (
					note   string
					octave string
					sample string
					effect string
					param  string
				)
				r := strings.Split(r, " ")
				if len(r) == 5 {
					r = r[1:]
				}
				for _, col := range []int{2, 8, 14, 20} {
					var channel int
					switch col {
					case 2:
						channel = 0
					case 8:
						channel = 1
					case 14:
						channel = 2
					case 20:
						channel = 3
					}
					note = strings.ReplaceAll(r[channel][:2], "-", "")
					octave = strings.ReplaceAll(string(r[channel][2]), "-", "")
					sample = strings.ReplaceAll(r[channel][3:5], "-", "")
					effect = strings.ReplaceAll(string(r[channel][5]), "-", "")
					param = strings.ReplaceAll(r[channel][6:8], "-", "")
					channelsSelect[pattern].channel.Objects[0].(*container.Scroll).Content.(*fyne.Container).Objects[col].(*fyne.Container).Objects[row].(*widget.Entry).SetText(note)
					channelsSelect[pattern].channel.Objects[0].(*container.Scroll).Content.(*fyne.Container).Objects[col+1].(*fyne.Container).Objects[row].(*widget.Entry).SetText(octave)
					channelsSelect[pattern].channel.Objects[0].(*container.Scroll).Content.(*fyne.Container).Objects[col+2].(*fyne.Container).Objects[row].(*widget.Entry).SetText(sample)
					channelsSelect[pattern].channel.Objects[0].(*container.Scroll).Content.(*fyne.Container).Objects[col+3].(*fyne.Container).Objects[row].(*widget.Entry).SetText(effect)
					channelsSelect[pattern].channel.Objects[0].(*container.Scroll).Content.(*fyne.Container).Objects[col+4].(*fyne.Container).Objects[row].(*widget.Entry).SetText(param)
				}
			}
		}
		i++
	}
}

func normilizeSeq(sequence []string, patterns []string) []string {
	var (
		r       = make([]string, 10)
		trapped = make(map[int]string)
		free    = make(map[int]string)
		nums    = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	)
	for i, p := range sequence {
		n, _ := strconv.Atoi(p)
		if n < 10 {
			trapped[i] = p
		} else {
			free[i] = p
		}
	}
	for _, p := range trapped {
		n, _ := strconv.Atoi(p)
		if contains(nums, n) {
			nums = remove(nums, n)
		}
	}
	j := 0
	for i := range free {
		n := strconv.Itoa(nums[j])
		for j, pat := range patterns {
			if strings.Contains(pat, "Pattern "+free[i]) {
				for k, pat := range patterns {
					if strings.Contains(pat, "Pattern "+n) {
						rows := strings.Split(pat, "\n")
						patterns[k] = strings.ReplaceAll(patterns[k], rows[0], "Pattern 999")
					}
				}
				rows := strings.Split(pat, "\n")
				patterns[j] = strings.ReplaceAll(patterns[j], rows[0], "Pattern "+n)
				break
			}
		}
		free[i] = n
		j++
	}
	for i, p := range free {
		r[i] = p
	}
	for i, p := range trapped {
		r[i] = p
	}
	return r
}

func saveAs() {
	dialog.ShowFileSave(func(closer fyne.URIWriteCloser, err error) {
		if closer == nil {
			return
		}
		name := strings.Split(closer.URI().Name(), ".")
		if len(name) != 2 || name[1] != "mod" {
			closer.Close()
			dialog.ShowError(errors.New(locale["wrong format"]), *window)
			if _, err := os.Stat(closer.URI().Path()); !errors.Is(err, os.ErrNotExist) {
				os.Remove(closer.URI().Path())
			}
			return
		}
		saveSong(closer.URI().Path())
	}, *window)
}

func saveSong(path string) {
	if _, err := os.Stat("temp.mod"); !errors.Is(err, os.ErrNotExist) {
		os.Remove("temp.mod")
	}
	cmd := exec.Command("java", "-jar", "micromod.jar", "temp.mt", "-out", "temp.mod")
	stdout, _ := cmd.Output()
	fmt.Println(string(stdout))
	data, err := ioutil.ReadFile(path)
	if err != nil {
		cathcer <- err
		return
	}
	if _, err := os.Stat("old.mod"); !errors.Is(err, os.ErrNotExist) {
		os.Remove("old.mod")
	}
	f, err := os.Create("old.mod")
	if err != nil {
		cathcer <- err
		return
	}
	_, err = f.Write(data)
	if err != nil {
		cathcer <- err
		return
	}
	if _, err := os.Stat(path); !errors.Is(err, os.ErrNotExist) {
		os.Remove(path)
	}
	data, err = ioutil.ReadFile("temp.mod")
	if err != nil {
		cathcer <- err
		return
	}
	f, err = os.Create(path)
	if err != nil {
		cathcer <- err
		return
	}
	_, err = f.Write(data)
	if err != nil {
		cathcer <- err
		return
	}
	savePath = path
	needSave = false
}
