package main

import (
	"go-calculator/calculator"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	fyneApp := app.New()

	window := fyneApp.NewWindow("Calculator")
	window.SetTitle("Calculator")
	window.Resize(fyne.NewSize(250, 250))

	formula := widget.NewLabel("")
	formula.TextStyle = fyne.TextStyle{Bold: true, Monospace: true}
	formula.Alignment = fyne.TextAlignTrailing

	calc := calculator.Calculator{CurrentCalculatedValue: 0.0, CurrentFormula: "", UpdateFormula: formula.SetText}

	buttonsContainer := calc.RenderButtons()
	content := container.NewVBox(layout.NewSpacer(), formula, layout.NewSpacer(), buttonsContainer)

	window.SetContent(content)

	window.ShowAndRun()
}
