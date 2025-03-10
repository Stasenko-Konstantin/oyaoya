package src

import (
	"errors"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	theme2 "fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"os"
)

var (
	cathcer  = make(chan error)
	window   *fyne.Window
	fon      *canvas.Image
	widgets  *fyne.Container
	needSave = true
	savePath string
)

func errorCatcher(handler chan error, w fyne.Window) {
	dialog.ShowError(<-handler, w)
}

func Start() {
	var (
		width  float32 = 1025
		height float32 = 653
	)

	logo, _ := fyne.LoadResourceFromPath("etc/logo.png")

	a := app.New()
	a.SetIcon(logo)
	a.Settings().SetTheme(theme2.DarkTheme())
	w := a.NewWindow("oyaoya")
	w.Resize(fyne.NewSize(width, height))
	w.SetFixedSize(true)
	w.CenterOnScreen()
	window = &w

	fon = canvas.NewImageFromFile("etc/fon.jpg")
	fon.FillMode = canvas.ImageFillStretch
	fon.Move(fyne.NewPos(-4, -5))
	fon.Resize(fyne.NewSize(width, height))
	fon.Show()

	go errorCatcher(cathcer, w)

	readConf()
	readLocale()
	clean()

	instructionC := container.NewVScroll(widget.NewLabel(instruction))
	instructionC.SetMinSize(fyne.NewSize(500, 400))

	mainMenu := fyne.NewMainMenu(
		fyne.NewMenu(locale["menu"],
			fyne.NewMenuItem(locale["new"], func() { setNewSong() }),
			fyne.NewMenuItem(locale["open"], func() {
				dialog.ShowFileOpen(func(closer fyne.URIReadCloser, err error) {
					if closer == nil {
						return
					}
					openSong(closer.URI().Path(), closer.URI().Name())
				}, w)
			}),
			fyne.NewMenuItem(locale["save"], func() {
				if needSave {
					saveAs()
				} else {
					saveSong(savePath)
				}
			}),
			fyne.NewMenuItem(locale["save as"], func() { saveAs() }),
			fyne.NewMenuItem(locale["settings"], func() {
				dialog.ShowCustom(locale["settings"], locale["cancel"], container.NewVBox(
					container.NewHScroll(container.NewHBox(
						makeLangs()...,
					)),
					container.NewHBox(
						widget.NewButton(locale["dark theme"], func() { a.Settings().SetTheme(theme2.DarkTheme()); fon.Hide(); setTheme("dark") }),
						widget.NewButton(locale["light theme"], func() { a.Settings().SetTheme(theme2.LightTheme()); fon.Hide(); setTheme("light") }),
						widget.NewButton(locale["color theme"], func() { a.Settings().SetTheme(theme2.DarkTheme()); fon.Show(); setTheme("color") }),
					),
				), w)
			}),
		),
		fyne.NewMenu(locale["help"],
			fyne.NewMenuItem(locale["about"], func() { dialog.ShowInformation(locale["about"], locale["about text"], w) }),
			fyne.NewMenuItem(locale["instruction"], func() {
				dialog.ShowCustom(locale["instruction"], "OK", instructionC, w)
			}),
			fyne.NewMenuItem(locale["license"], func() { dialog.ShowInformation(locale["license"], license, w) }),
		),
	)
	w.SetMainMenu(mainMenu)

	widgets = container.NewWithoutLayout(
		fon,
		makePatterns(),
		makeInstruments("samples"),
		makePlay(),
		makeNames(),
		makeChannels(),
	)
	w.SetContent(widgets)
	hideAll()
	channelsSelect[0].channel.Show()

	if !confOk {
		cathcer <- errors.New("Конфигурационный файл не найден!\nВсе настройки выставлены по умолчанию.")
	}
	if !localeOk {
		cathcer <- errors.New("Файл локализации не найден!\nЯзыковые настройки выставлены по умолчанию.")
	}

	instr, err := getInstruction(locale["instruction text"])
	instruction = instr
	if err != nil {
		cathcer <- err
	}
	instructionC.Content.(*widget.Label).SetText(instruction)

	go func() {
		openSong("test.mod", "test.mod")
		setNewSong()
	}()

	w.ShowAndRun()
}

func setTheme(theme string) {
	setConfigField("theme", theme)
}

func clean() {
	if _, err := os.Stat("temp.mod"); !errors.Is(err, os.ErrNotExist) {
		os.Remove("temp.mod")
	}
	if _, err := os.Stat("temp.mt"); !errors.Is(err, os.ErrNotExist) {
		os.Remove("temp.mt")
	}
	if _, err := os.Stat("temp"); !errors.Is(err, os.ErrNotExist) {
		os.RemoveAll("temp")
		os.Remove("temp")
	}
}
