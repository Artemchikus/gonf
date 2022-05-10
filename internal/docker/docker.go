package docker

import (
	"gonf/internal/docker/models"
	"gonf/internal/docker/models/creator"
	"io/ioutil"
	"log"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"gopkg.in/yaml.v2"
)

func ShowTemplate(w fyne.Window) {
	instructions := getInstructions()

	form := newForm(w)

	instList := newInstructionList(instructions, w)

	creator := creator.New(&w, form, instList, instructions)

	hello := widget.NewLabel("Создание Dockerfile")
	hello.Alignment = fyne.TextAlignCenter
	hello.TextStyle = fyne.TextStyle{Bold: true}

	form.AppendItem(creator.CreateNewFromItem(instructions[0]))

	form.AppendItem(creator.CreateEmptyFormItem())

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

func newForm(w fyne.Window) *widget.Form {
	cancelDialog := dialog.NewConfirm("Вы уверены?", "Все поля будут сброшены", func(b bool) {
		if b {
			ShowTemplate(w)
		}
	}, w)

	form := &widget.Form{OnCancel: func() { cancelDialog.Show() }}

	var filePath string

	fileDialog := dialog.NewFileSave(func(lu fyne.URIWriteCloser, err error) {
		filePath = lu.URI().Path()

		parseToFile(form, filePath)

		filePath = "Путь до файла: " + filePath

		finishDialog := dialog.NewInformation("Конфиг создан", filePath, w)

		finishDialog.Show()
	}, w)

	form.OnSubmit = func() {
		fileDialog.Show()
	}
	return form
}

func parseToFile(form *widget.Form, filePath string) {
	newFile, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	lines := []byte("")
	var fromCount int
	for _, item := range form.Items {
		if item.Text != "" {
			str := item.Text + " " + item.Widget.(*fyne.Container).Objects[0].(*widget.Entry).Text + "\n\n"

			if item.Text == "FROM" {
				fromCount++
				if fromCount > 1 {
					lines = append(lines, "\n\n"...)
				}
			}

			lines = append(lines, str...)
		}
	}
	lines = lines[:len(lines)-3]

	err = ioutil.WriteFile(filePath, lines, 0666)
	if err != nil {
		log.Fatal(err)
	}
}

func newInstructionList(instructions []*models.Instruction, window fyne.Window) *widget.List {
	list := &widget.List{
		Length: func() int {
			return len(instructions)
		},
		CreateItem: func() fyne.CanvasObject {
			name := widget.NewLabel("Name")

			selectImage := widget.NewIcon(theme.ConfirmIcon())
			selectImage.Hide()

			askButtom := widget.NewButtonWithIcon("", theme.QuestionIcon(), func() {})

			hbox := container.NewHBox(selectImage, askButtom)

			border := container.NewBorder(nil, nil, nil, hbox, name)
			return border
		},
		UpdateItem: func(id widget.ListItemID, co fyne.CanvasObject) {
			co.(*fyne.Container).Objects[0].(*widget.Label).SetText(instructions[id].Name)

			co.(*fyne.Container).Objects[1].(*fyne.Container).Objects[1].(*widget.Button).OnTapped = func() {
				description := widget.NewLabel(instructions[id].Description)
				description.Wrapping = fyne.TextWrapWord

				scroll := container.NewHScroll(description)

				info := dialog.NewCustom(instructions[id].Name, "OK", scroll, window)
				info.Show()
				info.Resize(fyne.Size{Width: window.Canvas().Size().Width * 0.85})
			}

			if instructions[id].IsSelected {
				co.(*fyne.Container).Objects[1].(*fyne.Container).Objects[0].(*widget.Icon).Show()
			} else {
				co.(*fyne.Container).Objects[1].(*fyne.Container).Objects[0].(*widget.Icon).Hide()
			}
		},
	}

	list.OnSelected = func(id widget.ListItemID) {
		instructions[id].IsSelected = true
		list.Refresh()
	}

	list.OnUnselected = func(id widget.ListItemID) {
		instructions[id].IsSelected = false
		list.Refresh()
	}

	return list
}
