package kubernetes

import (
	"gonf/internal/kubernetes/resources/models"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func ShowTemplate(w fyne.Window) {
	hello := widget.NewLabel("Создание ресурсов Kubernetes")

	mainPage := w.Content()

	resources := getResources(w)

	w.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Назад", func() {
			w.SetContent(mainPage)
		}),
		resources.Pod,
	))
}

func getResources(w fyne.Window) *models.Resources {
	return &models.Resources{Pod: widget.NewButton("Модуль", func() {
		hello := widget.NewLabel("Создание конфига Модуля")

		resoucePage := w.Content()

		w.SetContent(container.NewVBox(
			hello,
			widget.NewButton("Назад", func() {
				w.SetContent(resoucePage)
			}),
		))
	})}
}
