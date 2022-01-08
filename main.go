package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

const LoremPicsum = "https://picsum.photos/500"

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Random Pics")

	// main menu
	fileMenu := fyne.NewMenu("File",
		fyne.NewMenuItem("Quit", func() { myApp.Quit() }),
	)

	helpMenu := fyne.NewMenu("Help",
		fyne.NewMenuItem("About", func() {
			dialog.ShowCustom("About", "Close", container.NewVBox(
				widget.NewLabel("Welcome to Gopher Pics"),
				widget.NewLabel("Version: v0.1"),
				widget.NewLabel("Author: Tajwar Rahman"),
			), myWindow)
		}))

	mainMenu := fyne.NewMainMenu(
		fileMenu,
		helpMenu,
	)

	myWindow.SetMainMenu(mainMenu)

	// Welcome text
	text := canvas.NewText("Display a random Gopher", color.White)
	text.Alignment = fyne.TextAlignCenter

	// Gopher image
	var resource, _ = fyne.LoadResourceFromURLString(LoremPicsum)
	gopherImg := canvas.NewImageFromResource(resource)
	gopherImg.SetMinSize(fyne.Size{Width: 500, Height: 500})

	// create a button to get random gopher images
	button := widget.NewButton("Get Random Pic", func() {
		resource, _ = fyne.LoadResourceFromURLString(LoremPicsum)
		gopherImg.Resource = resource

		// redraw img
		gopherImg.Refresh()
	})

	button.Importance = widget.HighImportance

	// a box containing the text and img and button
	box := container.NewVBox(
		text,
		gopherImg,
		button,
	)

	// display
	myWindow.SetContent(box)

	// close window when esc is pressed
	myWindow.Canvas().SetOnTypedKey(func(ke *fyne.KeyEvent) {
		if ke.Name == fyne.KeyEscape {
			myApp.Quit()
		}
	})

	myWindow.ShowAndRun()
}
