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

	creator := creator.NewFormCreator(&w, form, instList, instructions)

	hello := widget.NewLabel("Создание Dockerfile")
	hello.Alignment = fyne.TextAlignCenter
	hello.TextStyle = fyne.TextStyle{Bold: true}

	form.AppendItem(creator.CreateEmptyFormItem())

	mainPage := w.Content()

	backButton := widget.NewButton("Назад", func() {
		w.SetContent(mainPage)
	})

	backBox := container.NewGridWithRows(1, &widget.Label{}, &widget.Label{}, backButton)

	samples := getSamples()

	sampleList := newSampleList(samples, w)

	sampleButton := widget.NewButton("Выбрать шаблон", func() {
		con := container.NewBorder(nil, nil, nil, nil, sampleList)

		choose := dialog.NewCustomConfirm("Выбор шаблона", "CHOOSE", "CANCEL", con, func(ch bool) {
			if ch {
				form.Items = form.Items[:1]
				for index, smp := range samples {
					if smp.IsSelected {
						for _, instName := range smp.InstructionNames {
							for _, inst := range instructions {
								if instName == inst.Name {
									form.AppendItem(creator.CreateNewFromItem(inst))
									break
								}
							}
						}
						sampleList.Unselect(index)
						form.Refresh()
						break
					}
				}
			}
		}, w)

		choose.Show()
		choose.Resize(fyne.Size{Width: form.Size().Width * 0.85, Height: w.Canvas().Size().Height})
	})

	sampleBox := container.NewGridWithRows(1, &widget.Label{}, &widget.Label{}, sampleButton)

	vBox := container.NewVBox(
		hello,
		sampleBox,
		form,
		backBox,
	)

	border := container.NewBorder(nil, nil, nil, container.NewHBox(&widget.Label{}, &widget.Label{}), vBox)

	scrollbar := container.NewVScroll(border)

	w.SetContent(scrollbar)
}

func getInstructions() []*models.Instruction {
	var instMass []*models.Instruction
	instructions := &models.Instructions{InstructMass: instMass}

	instructionCPath := os.Getenv("CONFIGPATH") + "/docker/instructions.yaml"

	sampleYaml, err := ioutil.ReadFile(instructionCPath)
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(sampleYaml, instructions)
	if err != nil {
		log.Fatal(err)
	}

	rezult := instructions.InstructMass

	return rezult
}

func newForm(w fyne.Window) *widget.Form {
	emprtyPage := w.Content()
	cancelDialog := dialog.NewConfirm("Вы уверены?", "Все поля будут сброшены", func(b bool) {
		if b {
			w.Canvas().SetContent(emprtyPage)
			ShowTemplate(w)
		}
	}, w)

	form := &widget.Form{OnCancel: func() { cancelDialog.Show() }}

	var filePath string

	fileDialog := dialog.NewFileSave(func(lu fyne.URIWriteCloser, err error) {
		if lu == nil {
			return
		}

		filePath = lu.URI().Path()

		parseToFile(form, filePath)

		filePath = "Путь до файла: " + filePath

		finishDialog := dialog.NewInformation("Конфиг создан", filePath, w)

		finishDialog.Show()
	}, w)

	fileDialog.SetFileName("Dockerfile.yaml")

	form.OnSubmit = func() {
		for _, item := range form.Items {
			err := item.Widget.(*fyne.Container).Objects[0].(*widget.Entry).Validate()
			if err != nil {
				errorDialog := dialog.NewError(err, w)
				errorDialog.Show()
				return
			}
		}
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
		if item.Text == "" {
			continue
		}

		str := item.Text + " " + item.Widget.(*fyne.Container).Objects[0].(*widget.Entry).Text + "\n\n"

		if item.Text == "FROM" {
			fromCount++
			if fromCount > 1 {
				lines = append(lines, "\n\n"...)
			}
		}

		lines = append(lines, str...)
	}
	lines = lines[:len(lines)-2]

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

				scroll := container.NewVScroll(description)
				scroll.SetMinSize(fyne.Size{Height: description.MinSize().Height * 4})

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

func getSamples() []*models.Sample {
	var sampleMass []*models.Sample
	samples := &models.Samples{SampleMass: sampleMass}

	instructionCPath := os.Getenv("CONFIGPATH") + "/docker/samples.yaml"

	sampleYaml, err := ioutil.ReadFile(instructionCPath)
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(sampleYaml, samples)
	if err != nil {
		log.Fatal(err)
	}

	rezult := samples.SampleMass

	return rezult
}

func newSampleList(samples []*models.Sample, window fyne.Window) *widget.List {
	list := &widget.List{
		Length: func() int {
			return len(samples)
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
			co.(*fyne.Container).Objects[0].(*widget.Label).SetText(samples[id].Name)

			co.(*fyne.Container).Objects[1].(*fyne.Container).Objects[1].(*widget.Button).OnTapped = func() {
				description := widget.NewLabel(samples[id].Description)
				description.Wrapping = fyne.TextWrapWord

				scroll := container.NewVScroll(description)
				scroll.SetMinSize(fyne.Size{Height: window.Canvas().Size().Height * 0.1})

				info := dialog.NewCustom(samples[id].Name, "OK", scroll, window)
				info.Show()
				info.Resize(fyne.Size{Width: window.Canvas().Size().Width * 0.85})
			}

			if samples[id].IsSelected {
				co.(*fyne.Container).Objects[1].(*fyne.Container).Objects[0].(*widget.Icon).Show()
			} else {
				co.(*fyne.Container).Objects[1].(*fyne.Container).Objects[0].(*widget.Icon).Hide()
			}
		},
	}

	list.OnSelected = func(id widget.ListItemID) {
		samples[id].IsSelected = true
		list.Refresh()
	}

	list.OnUnselected = func(id widget.ListItemID) {
		samples[id].IsSelected = false
		list.Refresh()
	}

	return list
}
