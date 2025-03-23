package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")
	l := widget.NewLabel("Hello Fyne!")
	e := widget.NewEntry()
	b := widget.NewButton("click", func() {
		dialog.ShowCustomConfirm("Enter message.", "OK",
			"Cancel", e, func(f bool) {
				if f {
					l.SetText("typed: '" + e.Text + "'.")
				} else {
					l.SetText("no message...")
				}
			}, w)
	})
	w.SetContent(
		container.New(
			layout.NewBorderLayout(
				nil, b, nil, nil,
			),
			l, b,
		),
	)
	w.Resize(fyne.NewSize(300, 200))
	w.ShowAndRun()
}
