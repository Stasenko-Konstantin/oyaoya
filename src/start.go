package src

import (
	"errors"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func Start() {
	readConf()
	readLocale()

	a := app.New()
	w := a.NewWindow("oyaoya")
	w.Resize(fyne.NewSize(1200, 700))

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

	mainMenu := fyne.NewMainMenu(fyne.NewMenu("Меню",
		fyne.NewMenuItem(locale["license"], func() { dialog.ShowInformation(locale["license"], license, w) })))
	w.SetMainMenu(mainMenu)

	label := widget.NewLabel(locale["text"])

	w.SetContent(container.NewVBox(
		label,
	))

	if !confOk {
		dialog.ShowError(errors.New("Конфигурационный файл не найден!\nВсе настройки выставлены по умолчанию."), w)
	}
	if !localeOk {
		dialog.ShowError(errors.New("Файл локализации не найден!\nЯзыковые настройки выставлены по умолчанию."), w)
	}

	w.ShowAndRun()
}