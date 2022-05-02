package prometheus

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func ShowTemplate(w fyne.Window) {
	hello := widget.NewLabel("Создание prometheus.conf")
	
	hello.Move(fyne.NewPos(w.Canvas().Size().Width/2, 0))

	mainPage := w.Content()

	w.SetContent(container.NewWithoutLayout(
		hello,
		widget.NewButton("Назад", func() {
			w.SetContent(mainPage)
		}),
	))
}
