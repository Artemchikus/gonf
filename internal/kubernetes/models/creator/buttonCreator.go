package creator

import (
	"gonf/internal/kubernetes/models"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type buttonCreator struct {
	resources []*models.Resource
	window    *fyne.Window
}

func (c *buttonCreator) SetButtons(box *fyne.Container) {
	podButton := setPodButton()
	box.Add(podButton)
}

func (c *buttonCreator) SetAskButton(resource *models.Resource) *widget.Button {
	askButton := widget.NewButtonWithIcon("", theme.QuestionIcon(), func() {
		description := widget.NewLabel(resource.Description)
		description.Wrapping = fyne.TextWrapWord

		scroll := container.NewVScroll(description)
		scroll.SetMinSize(fyne.Size{Height: description.MinSize().Height * 4})

		w := *c.window

		info := dialog.NewCustom(resource.Name, "OK", scroll, w)
		info.Show()
		info.Resize(fyne.Size{Width: w.Canvas().Size().Width * 0.85})
	})

	return askButton
}

func setPodButton() *fyne.Container {
	return nil
}

