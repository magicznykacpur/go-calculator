package calculator

import (
	"log"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type Calculator struct {
	CurrentCalculatedValue float32
	CurrentFormula         string
	UpdateFormula          func(string)
}

func (c *Calculator) RenderButtons() *fyne.Container {
	buttonsContainer := container.New(layout.NewGridLayoutWithColumns(2))

	numbersContainer := c.getNumbersContainer()
	operatorsContainer := c.getOperatorsContainer()

	buttonsContainer.Add(numbersContainer)
	buttonsContainer.Add(operatorsContainer)

	return buttonsContainer
}

func (c *Calculator) getOperatorsContainer() *fyne.Container {
	operatorsContainer := container.New(layout.NewVBoxLayout())

	plus := widget.NewButton("+", func() {
		c.CurrentFormula = c.CurrentFormula + "+"
		c.UpdateFormula(c.CurrentFormula)
	})
	minus := widget.NewButton("-", func() {
		c.CurrentFormula = c.CurrentFormula + "-"
		c.UpdateFormula(c.CurrentFormula)
	})
	multiply := widget.NewButton("*", func() {
		c.CurrentFormula = c.CurrentFormula + "*"
		c.UpdateFormula(c.CurrentFormula)
	})
	divide := widget.NewButton("/", func() {
		c.CurrentFormula = c.CurrentFormula + "/"
		c.UpdateFormula(c.CurrentFormula)
	})

	operatorsContainer.Add(plus)
	operatorsContainer.Add(minus)
	operatorsContainer.Add(multiply)
	operatorsContainer.Add(divide)

	return operatorsContainer
}

func (c *Calculator) getNumbersContainer() *fyne.Container {
	numbersContainer := container.NewVBox()

	firstRow := container.New(layout.NewGridLayoutWithColumns(3))
	for i := 1; i < 4; i++ {
		button := widget.NewButton(strconv.Itoa(i), func() {
			c.CurrentFormula = c.CurrentFormula + strconv.Itoa(i)
			c.UpdateFormula(c.CurrentFormula)
		})
		firstRow.Add(button)
	}

	secondRow := container.New(layout.NewGridLayoutWithColumns(3))
	for i := 5; i < 8; i++ {
		button := widget.NewButton(strconv.Itoa(i), func() {
			c.CurrentFormula = c.CurrentFormula + strconv.Itoa(i)
			c.UpdateFormula(c.CurrentFormula)
		})
		secondRow.Add(button)
	}

	thirdRow := container.New(layout.NewGridLayoutWithColumns(3))
	for i := 7; i < 10; i++ {
		button := widget.NewButton(strconv.Itoa(i), func() {
			c.CurrentFormula = c.CurrentFormula + strconv.Itoa(i)
			c.UpdateFormula(c.CurrentFormula)
		})
		thirdRow.Add(button)
	}

	lastRow := container.New(layout.NewAdaptiveGridLayout(2))
	lastRow.Add(widget.NewButton("0", func() {
		c.CurrentFormula = c.CurrentFormula + "0"
		c.UpdateFormula(c.CurrentFormula)
	}))
	lastRow.Add(widget.NewButton("=", func() {
		log.Println("=")
	}))

	numbersContainer.Add(thirdRow)
	numbersContainer.Add(secondRow)
	numbersContainer.Add(firstRow)
	numbersContainer.Add(lastRow)

	return numbersContainer
}
