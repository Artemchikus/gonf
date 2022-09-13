package creator

import (
	"gonf/internal/docker/models"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/validation"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type formCreator struct {
	window          *fyne.Window
	form            *widget.Form
	instructionList *widget.List
	instructions    []*models.Instruction
}

func (c *formCreator) CreateNewFromItem(instruction *models.Instruction) *widget.FormItem {
	entry := widget.NewEntry()
	entry.PlaceHolder = instruction.PlaceHolder
	SetValidation(entry, instruction)

	if instruction.Name == "RUN" {
		entry.MultiLine = true
	}

	hbox := container.NewHBox()

	label := widget.NewLabel(instruction.Description)
	label.Wrapping = fyne.TextWrapWord

	con := container.NewVScroll(label)
	con.SetMinSize(fyne.Size{Width: c.form.Size().Width * 0.85, Height: c.form.Size().Height * 0.85})

	info := dialog.NewCustom(instruction.Name, "OK", con, *c.window)

	askButtom := widget.NewButtonWithIcon("", theme.QuestionIcon(), func() {
		info.Show()
		info.Resize(fyne.Size{Width: c.form.Size().Width * 0.85})
	})
	hbox.Add(askButtom)

	addButton := widget.NewButtonWithIcon("", theme.ContentAddIcon(), func() {
		con := container.NewBorder(nil, nil, nil, nil, c.instructionList)

		window := *c.window

		choose := dialog.NewCustomConfirm("Добавление инструкции", "ADD", "CANCEL", con, func(ch bool) {
			if ch {
				for index, inst := range c.instructions {
					if inst.IsSelected {
						c.AddFormItemAfter(*inst, instruction)
						c.instructionList.Unselect(index)
						break
					}
				}
			}
		}, window)

		choose.Show()
		choose.Resize(fyne.Size{Width: c.form.Size().Width * 0.85, Height: window.Canvas().Size().Height})
	})
	hbox.Add(addButton)

	border := container.NewBorder(nil, nil, nil,
		hbox,
		entry,
	)

	formItem := &widget.FormItem{
		Text:     instruction.Name,
		Widget:   border,
		HintText: instruction.HintText,
	}

	delButton := widget.NewButtonWithIcon("", theme.ContentClearIcon(), func() {
		c.DelFormItem(formItem)
	})
	hbox.Add(delButton)

	return formItem
}

func (c *formCreator) AddFormItemAfter(add models.Instruction, after *models.Instruction) {
	for index := len(c.form.Items) - 1; index >= 0; index-- {
		if c.form.Items[index].Text == after.Name {
			newEntry := c.CreateNewFromItem(&add)

			aft := make([]*widget.FormItem, len(c.form.Items)-1-index)
			copy(aft, c.form.Items[index+1:])

			c.form.Items = c.form.Items[:index+1]
			c.form.Refresh()

			c.form.Items = append(c.form.Items, newEntry)
			c.form.Items = append(c.form.Items, aft...)
			c.form.Refresh()
			break
		}
	}
}

func (c *formCreator) DelFormItem(formItem *widget.FormItem) {
	for index, item := range c.form.Items {
		if item == formItem {
			oldItems := c.form.Items[index+1:]
			c.form.Items = c.form.Items[:index]
			c.form.Refresh()
			c.form.Items = append(c.form.Items, oldItems...)
			c.form.Refresh()
			break
		}
	}
}

func NewFormCreator(window *fyne.Window, form *widget.Form, instructionList *widget.List, instructions []*models.Instruction) *formCreator {
	return &formCreator{
		window:          window,
		form:            form,
		instructionList: instructionList,
		instructions:    instructions,
	}
}

func (c *formCreator) CreateEmptyFormItem() *widget.FormItem {
	entry := widget.NewEntry()
	entry.PlaceHolder = "Пустое поле"
	entry.Disable()

	instruction := &models.Instruction{}

	hbox := container.NewHBox()

	addButton := widget.NewButtonWithIcon("", theme.ContentAddIcon(), func() {
		con := container.NewBorder(nil, nil, nil, nil, c.instructionList)
		window := *c.window
		choose := dialog.NewCustomConfirm("Добавление инструкции", "ADD", "CANCEL", con, func(ch bool) {
			if ch {
				for index, inst := range c.instructions {
					if inst.IsSelected {
						c.AddFormItemAfter(*inst, instruction)
						c.instructionList.Unselect(index)
						break
					}
				}
			}
		}, window)
		choose.Show()
		choose.Resize(fyne.Size{Width: c.form.Size().Width * 0.85, Height: window.Canvas().Size().Height})
	})
	hbox.Add(addButton)

	border := container.NewBorder(nil, nil, nil,
		hbox,
		entry,
	)

	formItem := &widget.FormItem{
		Text:     instruction.Name,
		Widget:   border,
		HintText: instruction.HintText,
	}

	return formItem
}

func SetValidation(entry *widget.Entry, instruction *models.Instruction) {
	switch instruction.PlaceHolder {
	case "<образ>":
		entry.Validator = validation.NewRegexp(
			`[${}a-z0-9-A-Z_]+(:[${}a-z0-9-A-Z_.]+@sha256:[${}a-z0-9-A-Z_]+|:[${}a-z0-9-A-Z_.]+|@sha256:[${}a-z0-9-A-Z_]+|)( AS [a-z0-9-A-Z]+)*`,
			"Введенные данные не соотетсвтуют формату ввода")
	case "<команда>":
		entry.Validator = validation.NewRegexp(
			`( ?\w+ ?\\?)+|(\[(("[a-z0-9-A-Z.]+"(, |,|))|(\${[a-z0-9-A-Z_]+})(, |,|))+\])`,
			"Введенные данные не соотетсвтуют формату ввода")
	case "<ключ>=<значение>":
		entry.Validator = validation.NewRegexp(
			`[${}a-z-A-Z_]+(=[${}a-z0-9-A-Z_"/]+)`,
			"Введенные данные не соотетсвтуют формату ввода")
	case "<имя>=<значение по умолчанию>":
		entry.Validator = validation.NewRegexp(
			`[${}a-z-A-Z_]+(=[${}a-z0-9-A-Z _"/]*)?`,
			"Введенные данные не соотетсвтуют формату ввода")
	case "<откуда> <куда>":
		entry.Validator = validation.NewRegexp(
			`([${}a-z0-9-A-Z/._*?]+ )+([${}a-z0-9-A-Z/._]+\/)|([${}a-z0-9-A-Z/._]+ [${}a-z0-9-A-Z/._]+)`,
			"Введенные данные не соотетсвтуют формату ввода")
	case "[\"<исполняемый файл>\", \"<параметры>\"...]":
		entry.Validator = validation.NewRegexp(
			`\[(("[a-z0-9-A-Z.]+"(, |,|))|(\${[a-z0-9-A-Z_]+})(, |,|))+\]|((\w+) ?)+`,
			"Введенные данные не соотетсвтуют формату ввода")
	case "<порт>":
		entry.Validator = validation.NewRegexp(
			`((\d{1,6}(\/tcp|\/udp)? ?)*|\${[a-z-A-Z_]+}|)`,
			"Введенные данные не соотетсвтуют формату ввода")
	case "<инструкция>":
		entry.Validator = validation.NewRegexp(
			`(RUN|WORKDIR|ADD|COPY|VOLUME|USER|ARG|CMD|HEALTHCHECK|LABEL|ENTRYPOINT|ENV|SHELL|STOPSIGNAL|EXPOSE)(.|\n)+`,
			"Введенные данные не соотетсвтуют формату ввода")
	case "<сигнал>":
		entry.Validator = validation.NewRegexp(
			`(\d{1,2}|[A-Z]+)`,
			"Введенные данные не соотетсвтуют формату ввода")
	case "[опции...] <CMD команда>":
		entry.Validator = validation.NewRegexp(
			`(--interval=.+ |--timeout=.+ |--start-period=.+ |--retries=.+ |)( \\\n)?CMD (.|\n)+|NONE`,
			"Введенные данные не соотетсвтуют формату ввода")
	case "<имя>":
		entry.Validator = validation.NewRegexp(
			`(\w+|\${[a-z-A-Z_]+})`,
			"Введенные данные не соотетсвтуют формату ввода")
	case "<ключ>=\"<значение>\"":
		entry.Validator = validation.NewRegexp(
			`[${}a-z-A-Z_]+=("[a-z0-9-A-Z]+"|\${[a-z-A-Z_]+})`,
			"Введенные данные не соотетсвтуют формату ввода")
	case "/<рабочий каталог>":
		entry.Validator = validation.NewRegexp(
			`(\/?[a-z-A-Z0-9/]+|(|\/)\${[a-z-A-Z_]+})`,
			"Введенные данные не соотетсвтуют формату ввода")
	case "/<точка монтирования>":
		entry.Validator = validation.NewRegexp(
			`\[(("[a-z0-9-A-Z./]+"(, |,|))|(\${[a-z0-9-A-Z_/]+})(, |,|))+\]|(\/[a-z-A-Z0-9/]+|(|\/)\${[a-z-A-Z_]+})`,
			"Введенные данные не соотетсвтуют формату ввода")
	}
}
