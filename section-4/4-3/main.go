package main

import (
	"strconv"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	c := 0
	a := app.New()
	w := a.NewWindow("Hello")
	l := widget.NewLabel("Hello Fyne!")
	w.SetContent(
		// widget.NewVBox is API for Fyne v1, so use container.NewVBox() instead.
		container.NewVBox(
			l,
			widget.NewButton("Click me!", func() {
				c++
				l.SetText("count: " + strconv.Itoa(c))
			}),
		),
	)

	w.ShowAndRun()
}
