package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()

	w := a.NewWindow("Hello")
	w.SetContent(
		// widget.NewVBox is API for Fyne v1, so use container.NewVBox() instead.
		container.NewVBox(
			widget.NewLabel("Hello Fyne!"),
			widget.NewLabel("This is sample application!"),
		),
	)

	w.ShowAndRun()
}
