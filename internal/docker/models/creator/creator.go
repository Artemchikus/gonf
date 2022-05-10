package creator

import (
	"gonf/internal/docker/models"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type creator struct {
	window          *fyne.Window
	form            *widget.Form
	instructionList *widget.List
	instructions    []*models.Instruction
}

func (c *creator) CreateNewFromItem(instruction *models.Instruction) *widget.FormItem {
	entry := widget.NewEntry()
	entry.PlaceHolder = instruction.PlaceHolder

	hbox := container.NewHBox()

	label := widget.NewLabel(instruction.Description)
	label.Wrapping = fyne.TextWrapWord

	con := container.NewHScroll(label)

	info := dialog.NewCustom(instruction.Name, "OK", con, *c.window)

	askButtom := widget.NewButtonWithIcon("", theme.QuestionIcon(), func() {
		info.Show()
		info.Resize(fyne.Size{Width: c.form.Size().Width * 0.85})
	})
	hbox.Add(askButtom)

	if instruction.IsMany {
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
	}

	border := container.NewBorder(nil, nil, nil,
		hbox,
		entry,
	)

	formItem := &widget.FormItem{
		Text:     instruction.Name,
		Widget:   border,
		HintText: instruction.HintText,
	}

	if instruction.IsAdded {
		delButton := widget.NewButtonWithIcon("", theme.ContentClearIcon(), func() {
			c.DelFormItem(formItem)
		})
		hbox.Add(delButton)
	}

	return formItem
}

func (c *creator) AddFormItemAfter(add models.Instruction, after *models.Instruction) {
	for index := len(c.form.Items) - 1; index >= 0; index-- {
		if c.form.Items[index].Text == after.Name {
			add.IsAdded = true
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

func (c *creator) DelFormItem(formItem *widget.FormItem) {
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

func New(window *fyne.Window, form *widget.Form, instructionList *widget.List, instructions []*models.Instruction) *creator {
	return &creator{
		window:          window,
		form:            form,
		instructionList: instructionList,
		instructions:    instructions,
	}
}

func (c *creator) CreateEmptyFormItem() *widget.FormItem {
	entry := widget.NewEntry()
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
						c.AddFormItemBefore(*inst, instruction)
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

	if instruction.IsAdded {
		delButton := widget.NewButtonWithIcon("", theme.ContentClearIcon(), func() {
			c.DelFormItem(formItem)
		})
		hbox.Add(delButton)
	}

	return formItem
}


func (c *creator) AddFormItemBefore(add models.Instruction, before *models.Instruction) {
	for index := len(c.form.Items) - 1; index >= 0; index-- {
		if c.form.Items[index].Text == before.Name {
			add.IsAdded = true
			newEntry := c.CreateNewFromItem(&add)

			after := make([]*widget.FormItem, len(c.form.Items)-index)
			copy(after, c.form.Items[index:])

			c.form.Items = c.form.Items[:index]
			c.form.Refresh()

			c.form.Items = append(c.form.Items, newEntry)
			c.form.Items = append(c.form.Items, after...)
			c.form.Refresh()
			break
		}
	}
}