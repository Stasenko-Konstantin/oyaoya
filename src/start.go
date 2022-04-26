package src

import (
	"errors"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/driver/desktop"
)

func Start() {
	readConf()
	readLocale()

	a := app.New()
	w := a.NewWindow("oyaoya")
	w.Resize(fyne.NewSize(1200, 700))
	w.SetFixedSize(true)

	license := "\n\toyaoya - tracker music editor\n" +
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

	mainMenu := fyne.NewMainMenu(
		fyne.NewMenu(locale["menu"],
			fyne.NewMenuItem(locale["new"], func() { dialog.ShowInformation(locale["new"], locale["new"], w) }),
			fyne.NewMenuItem(locale["open"], func() { dialog.ShowInformation(locale["open"], locale["open"], w) }),
			fyne.NewMenuItem(locale["open recent"], func() { dialog.ShowInformation(locale["open recent"], locale["open recent"], w) }),
			fyne.NewMenuItem(locale["save"], func() { dialog.ShowInformation(locale["save"], locale["save"], w) }),
			fyne.NewMenuItem(locale["settings"], func() { dialog.ShowInformation(locale["settings"], locale["settings"], w) }),
		),
		fyne.NewMenu(locale["help"],
			fyne.NewMenuItem(locale["about"], func() { dialog.ShowInformation(locale["about"], locale["about"], w) }),
			fyne.NewMenuItem(locale["instruction"], func() { dialog.ShowInformation(locale["instruction"], locale["instruction"], w) }),
			fyne.NewMenuItem(locale["license"], func() { dialog.ShowInformation(locale["license"], license, w) }),
		),
	)
	w.SetMainMenu(mainMenu)

	channelsW := makeChannels()

	w.SetContent(container.NewVBox(
		container.NewWithoutLayout(
			channelsW,
		),
	))

	tab := desktop.CustomShortcut{fyne.KeyTab, desktop.ControlModifier}
	w.Canvas().AddShortcut(&tab, func(shortcut fyne.Shortcut) {
		if currChann == 3 {
			currChann = 0
		} else {
			currChann++
		}
	})

	if !confOk {
		dialog.ShowError(errors.New("Конфигурационный файл не найден!\nВсе настройки выставлены по умолчанию."), w)
	}
	if !localeOk {
		dialog.ShowError(errors.New("Файл локализации не найден!\nЯзыковые настройки выставлены по умолчанию."), w)
	}

	w.ShowAndRun()
}
