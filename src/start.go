package src

import (
	"errors"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

var (
	inMainWindow = true
	cathcer      = make(chan error)
	window       *fyne.Window
)

func errorCatcher(handler chan error, w fyne.Window) {
	dialog.ShowError(<-handler, w)
}

func Start() {
	readConf()
	readLocale()

	a := app.New()
	w := a.NewWindow("oyaoya")
	w.Resize(fyne.NewSize(1025, 700))
	w.SetFixedSize(true)
	w.CenterOnScreen()
	window = &w

	mainMenu := fyne.NewMainMenu(
		fyne.NewMenu(locale["menu"],
			fyne.NewMenuItem(locale["new"], func() { dialog.ShowInformation(locale["new"], locale["new"], w) }),
			fyne.NewMenuItem(locale["open"], func() { dialog.ShowFileOpen(func(closer fyne.URIReadCloser, err error) { openSong(closer) }, w) }),
			fyne.NewMenuItem(locale["save"], func() { dialog.ShowInformation(locale["save"], locale["save"], w) }),
			fyne.NewMenuItem(locale["save as"], func() { dialog.ShowInformation(locale["save as"], locale["save as"], w) }),
			fyne.NewMenuItem(locale["settings"], func() { dialog.ShowInformation(locale["settings"], locale["settings"], w) }),
		),
		fyne.NewMenu(locale["help"],
			fyne.NewMenuItem(locale["about"], func() { dialog.ShowInformation(locale["about"], locale["about"], w) }),
			fyne.NewMenuItem(locale["instruction"], func() { dialog.ShowCustom(locale["instruction"], "OK", widget.NewLabel(instruction), w) }),
			fyne.NewMenuItem(locale["license"], func() { dialog.ShowInformation(locale["license"], license, w) }),
		),
	)
	w.SetMainMenu(mainMenu)

	w.SetContent(container.NewWithoutLayout(
		makePatterns(),
		makePlay(),
		makeNames(),
		makeChannels(),
	))
	hideAll()
	channelsSelect[0].channel.Show()

	go errorCatcher(cathcer, w)

	if !confOk {
		cathcer <- errors.New("Конфигурационный файл не найден!\nВсе настройки выставлены по умолчанию.")
	}
	if !localeOk {
		cathcer <- errors.New("Файл локализации не найден!\nЯзыковые настройки выставлены по умолчанию.")
	}

	w.ShowAndRun()
}
