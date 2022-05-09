package docker

import (
	"gonf/internal/docker/models"
	"gonf/internal/docker/models/creator"
	"io/ioutil"
	"log"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"gopkg.in/yaml.v2"
)

func ShowTemplate(w fyne.Window) {
	instructions := getInstructions()

	form := &widget.Form{
		OnSubmit: func() {},
		OnCancel: func() {},
	}

	creator := creator.New(&w, form)

	hello := widget.NewLabel("Создание Dockerfile")
	hello.Alignment = fyne.TextAlignCenter
	hello.TextStyle = fyne.TextStyle{Bold: true}

	for index, inst := range instructions {
		if index < 3 {
		instEnrty := creator.CreateNewFromItem(inst)
		form.AppendItem(instEnrty)
		}
	}

	mainPage := w.Content()

	backButton := widget.NewButton("Назад", func() {
		w.SetContent(mainPage)
	})

	buttonBox := container.NewGridWithRows(1, &widget.Label{}, &widget.Label{}, backButton)

	vBox := container.NewVBox(
		hello,
		form,
		buttonBox,
	)

	border := container.NewBorder(nil, nil, nil, container.NewHBox(&widget.Label{}, &widget.Label{}), vBox)

	scrollbar := container.NewVScroll(border)

	w.SetContent(scrollbar)
}

func getInstructions() []*models.Instruction {
	var instMass []*models.Instruction
	instructions := &models.Instructions{InstructMass: instMass}

	fileName := &models.Instruction{
		Name:        "Название файла",
		IsMany:      false,
		Description: "Dockerfile",
		PlaceHolder: "Dockerfile.yaml",
		HintText:    "Обязательное поле",
	}

	instructions.InstructMass = append(instructions.InstructMass, fileName)

	instructionCPath := os.Getenv("CONFIGPATH") + "/docker/instructions.yaml"

	instructionYaml, err := ioutil.ReadFile(instructionCPath)
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(instructionYaml, instructions)
	if err != nil {
		log.Fatal(err)
	}

	rezult := instructions.InstructMass

	return rezult
}
