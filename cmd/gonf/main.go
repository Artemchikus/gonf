package main

import (
	"gonf/assets"
	"gonf/internal/docker"
	"gonf/internal/grafana"
	"gonf/internal/kubernetes"
	"gonf/internal/prometheus"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("GONF")
	w.Resize(fyne.Size{Width: assets.WindowWidth, Height: assets.WindowHeight})

	res, _ := fyne.LoadResourceFromPath("assets/icon.png")
	w.SetIcon(res)
	SetMainPage(w)
}

func SetMainPage(w fyne.Window) {
	hello := widget.NewLabel("Создание файлов конфигурации")
	hello.Alignment = fyne.TextAlignCenter
	hello.TextStyle = fyne.TextStyle{Bold: true}

	w.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Dockerfile", func() {
			docker.ShowTemplate(w)
		}),
		widget.NewButton("Kubernetes", func() {
			kubernetes.ShowTemplate(w)
		}),
		widget.NewButton("Prometheus", func() {
			prometheus.ShowTemplate(w)
		}),
		widget.NewButton("Grafana", func() {
			grafana.ShowTemplate(w)
		}),
	))
	w.ShowAndRun()
}
