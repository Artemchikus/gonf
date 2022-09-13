package creator

import (
	"gonf/internal/docker/models"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type formCreator struct {
	window *fyne.Window
	form   *widget.Form
}

func (c *formCreator) CreateNewFromItem(instruction *models.Instruction) *widget.FormItem {
	entry := widget.NewEntry()
	entry.PlaceHolder = instruction.PlaceHolder

	hbox := container.NewHBox()

	if !instruction.IsAdded {
		info := dialog.NewInformation(instruction.Name, instruction.Description, *c.window)

		label := widget.NewLabel(instruction.Description)

		label.Wrapping = fyne.TextWrapWord

		con := container.NewVScroll(label)
		con.SetMinSize(fyne.Size{Width: c.form.Size().Width * 0.85})

		info = dialog.NewCustom(instruction.Name, "OK", con, *c.window)

		askButtom := widget.NewButtonWithIcon("", theme.QuestionIcon(), func() {
			info.Show()
			info.Resize(fyne.Size{Width: c.form.Size().Width * 0.85})
		})
		hbox.Add(askButtom)

		if instruction.IsMany {
			addButton := widget.NewButtonWithIcon("", theme.ContentAddIcon(), func() {
				c.AddSameFormItem(*instruction)
			})
			hbox.Add(addButton)
		}
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

func (c *formCreator) AddSameFormItem(instruction models.Instruction) {
	for index := len(c.form.Items) - 1; index >= 0; index-- {
		if c.form.Items[index].Text == instruction.Name {
			instruction.IsAdded = true
			newEntry := c.CreateNewFromItem(&instruction)

			after := make([]*widget.FormItem, len(c.form.Items)-1-index)
			copy(after, c.form.Items[index+1:])

			c.form.Items = c.form.Items[:index+1]
			c.form.Refresh()

			c.form.Items = append(c.form.Items, newEntry)
			c.form.Items = append(c.form.Items, after...)
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

func New(window *fyne.Window, form *widget.Form) *formCreator {
	return &formCreator{
		window: window,
		form:   form,
	}
}
