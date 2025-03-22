package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")
	l := widget.NewLabel("Hello Fyne!")
	ck1 := widget.NewCheck("check 1", nil)
	ck2 := widget.NewCheck("check 2", nil)
	// widget.NewGroup is API for Fyne v1, so use grp variable instead.
	grp := container.NewVBox(
		ck1, ck2,
	)
	w.SetContent(
		// widget.NewVBox is API for Fyne v1, so use container.NewVBox() instead.
		container.NewVBox(
			l,
			grp,
			widget.NewButton("OK", func() {
				re := "result: "
				if ck1.Checked {
					re += "Check-1"
				}
				if ck2.Checked {
					re += "Check-2"
				}
				l.SetText(re)
			}),
		),
	)
	w.ShowAndRun()
}
