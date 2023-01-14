package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello World")

	historyBtn := widget.NewButton("History", func() {})
	backBtn := widget.NewButton("Back", func() {})
	clearBtn := widget.NewButton("Clear", func() {})

	plusBtn := widget.NewButton("+", func() {})
	minusBtn := widget.NewButton("-", func() {})
	multiplyBtn := widget.NewButton("*", func() {})
	divideBtn := widget.NewButton("/", func() {})

	oneBtn := widget.NewButton("1", func() {})
	twoBtn := widget.NewButton("2", func() {})
	threeBtn := widget.NewButton("3", func() {})
	fourBtn := widget.NewButton("4", func() {})
	fiveBtn := widget.NewButton("5", func() {})
	sixBtn := widget.NewButton("6", func() {})
	sevenBtn := widget.NewButton("7", func() {})
	eightBtn := widget.NewButton("8", func() {})
	nineBtn := widget.NewButton("9", func() {})
	zeroBtn := widget.NewButton("0", func() {})

	history := widget.NewLabel("History")
	input := widget.NewLabel("Input")

	w.SetContent(container.NewVBox(
		input,
		history,
		container.NewGridWithColumns(2, historyBtn, backBtn),
		container.NewGridWithColumns(4, plusBtn, minusBtn, multiplyBtn, divideBtn),
		container.NewGridWithColumns(3, oneBtn, twoBtn, threeBtn, fourBtn, fiveBtn, sixBtn, sevenBtn, eightBtn, nineBtn),
		container.NewGridWithColumns(2, clearBtn, zeroBtn),
	))
  w.Resize(fyne.NewSize(400, 400))
	w.ShowAndRun()
	fmt.Println("OK I AM RUNNING")
}
