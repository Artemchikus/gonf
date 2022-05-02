package docker

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func ShowTemplate(w fyne.Window) {
	hello := widget.NewLabel("Создание Dockerfile")
	hello.Alignment = fyne.TextAlignCenter

	fileName := widget.NewEntry()
	fileName.SetPlaceHolder("Название Dockerfile")
	fileName.SetText("Dockerfile.yaml")

	mainPage := w.Content()

	backButton := widget.NewButton("Назад", func() {
		w.SetContent(mainPage)
	})

	vBox := container.NewVBox(
		hello,
		widget.NewLabel("Название Dockerfile"),
		fileName,
		backButton,
	)

	w.SetContent(vBox)
}