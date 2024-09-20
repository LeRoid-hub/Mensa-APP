package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.NewWithID("mensa.barfuss.email")

	w := a.NewWindow("Hello")

	w.SetMainMenu(fyne.NewMainMenu())
	w.SetMaster()

	w.Resize(fyne.NewSize(800, 600))
	w.ShowAndRun()

}
