package prometheus

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func ShowTemplate(w fyne.Window) {
	hello := widget.NewLabel("Создание prometheus.conf")

	hello.Move(fyne.NewPos(w.Canvas().Size().Width/4, 0))

	mainPage := w.Content()

	returnButton := widget.NewButton("Назад", func() {
		w.SetContent(mainPage)
	})
	returnButton.Resize(fyne.NewSize(50, 50))
	returnButton.Move(fyne.NewPos(w.Canvas().Size().Width/2, hello.Size().Height+50))

	w.SetContent(container.NewWithoutLayout(
		hello,
		returnButton,
	))
}
