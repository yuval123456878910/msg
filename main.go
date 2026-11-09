package main

import (
	"msg/ui"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()
	w := a.NewWindow("homepage")
	w.Resize(fyne.NewSize(300, 600))
	w.SetContent(ui.Homepage(w))
	w.ShowAndRun()
}
