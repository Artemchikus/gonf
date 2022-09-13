package kubernetes

import (
	"gonf/internal/kubernetes/models"
	"io/ioutil"
	"log"
	"os"

	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"gopkg.in/yaml.v2"
)

func ShowTemplate(w fyne.Window) {
	// resources := getResources()

	hello := widget.NewLabel("Создание ресурсов Kubernetes")
	hello.Alignment = fyne.TextAlignCenter
	hello.TextStyle = fyne.TextStyle{Bold: true}

	mainPage := w.Content()

	backButton := widget.NewButton("Назад", func() {
		w.SetContent(mainPage)
	})

	backBox := container.NewGridWithRows(1, &widget.Label{}, &widget.Label{}, backButton)

	vBox := container.NewVBox(
		hello,
		backBox,
	)

	border := container.NewBorder(nil, nil, nil, container.NewHBox(&widget.Label{}, &widget.Label{}), vBox)

	scrollbar := container.NewVScroll(border)

	w.SetContent(scrollbar)
}

func getResources() []*models.Resource {
	var resMass []*models.Resource
	resources := &models.Resources{ResourceMass: resMass}

	resourcesCPath := os.Getenv("CONFIGPATH") + "/kubernetes/resources.yaml"

	sampleYaml, err := ioutil.ReadFile(resourcesCPath)
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(sampleYaml, resources)
	if err != nil {
		log.Fatal(err)
	}

	rezult := resources.ResourceMass

	return rezult
}

func setButtons(vBox *fyne.Container, w fyne.Window) {
	button := widget.NewButton("Pod", func(){
		
	})

	askButtom := widget.NewButtonWithIcon("", theme.QuestionIcon(), func() {
		description := widget.NewLabel(resources[0].Description)
		description.Wrapping = fyne.TextWrapWord

		scroll := container.NewVScroll(description)
		scroll.SetMinSize(fyne.Size{Height: description.MinSize().Height * 4})

		info := dialog.NewCustom(instructions[id].Name, "OK", scroll, w)
		info.Show()
		info.Resize(fyne.Size{Width: w.Canvas().Size().Width * 0.85})
	})

	border := container.NewBorder(nil, nil, nil, askButtom, button)
}