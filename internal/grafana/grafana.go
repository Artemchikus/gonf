package grafana

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func ShowTemplate(w fyne.Window) {
	hello := widget.NewLabel("Создание granfana.conf")

	mainPage := w.Content()

	w.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Назад", func() {
			w.SetContent(mainPage)
		}),
		makeFormTab(w),
	))
}

func makeFormTab(w fyne.Window) fyne.CanvasObject {
	name := widget.NewEntry()
	name.SetPlaceHolder("John Smith")

	// icon := widget.NewIcon(theme.QuestionIcon())

	email := widget.NewEntry()
	email.SetPlaceHolder("test@example.com")

	password := widget.NewPasswordEntry()
	password.SetPlaceHolder("Password")

	disabled := widget.NewRadioGroup([]string{"Option 1", "Option 2"}, func(string) {})
	disabled.Horizontal = true
	disabled.Disable()
	largeText := widget.NewMultiLineEntry()

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Name", Widget: name, HintText: "Your full name"},
			{Text: "Email", Widget: email, HintText: "A valid email address"},
		},
		OnCancel: func() {
		},
		OnSubmit: func() {
			fyne.CurrentApp().SendNotification(&fyne.Notification{
				Title:   "Form for: " + name.Text,
				Content: largeText.Text,
			})
		},
	}
	form.Append("Password", password)
	form.Append("Disabled", disabled)
	form.Append("Message", largeText)
	return form
}
