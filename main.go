package main

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type myTheme struct{}

var _ fyne.Theme = (*myTheme)(nil)

func (m *myTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	if name == theme.ColorNameBackground {
		return color.NRGBA{R: 18, G: 18, B: 18, A: 255}
	}

	if name == theme.ColorNameButton {
		return color.NRGBA{R: 0, G: 163, B: 255, A: 255}
	}

	if name == theme.ColorNameInputBackground {
		return color.NRGBA{R: 200, G: 30, B: 30, A: 255}
	}

	if name == theme.ColorNameForeground {
		return color.NRGBA{R: 240, G: 240, B: 240, A: 255}
	}

	return theme.DefaultTheme().Color(name, variant)
}

func (m *myTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)

	// return theme.DefaultTheme().Icon(name)
}
func (m *myTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

func (m *myTheme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}

func main() {
	a := app.New()
	a.Settings().SetTheme(&myTheme{})
	mainWindow := a.NewWindow("fyne-scan")

	heading := canvas.NewText("Fyne-Scanner", color.NRGBA{R: 255, G: 170, B: 0, A: 255})
	heading.Alignment = fyne.TextAlignCenter
	heading.TextStyle = fyne.TextStyle{Bold: true, Italic: true}
	heading.TextSize = 50

	inputText := canvas.NewText("Enter IP/Domain : ", color.NRGBA{R: 160, G: 165, B: 175, A: 255})
	inputText.TextStyle.Bold = true

	input := widget.NewEntry()
	input.SetPlaceHolder("Target IP/Domain")

	hbox1 := container.NewGridWithColumns(2, inputText, input)

	spacer := canvas.NewRectangle(color.Transparent)
	spacer.SetMinSize(fyne.NewSize(0, 20))
	infinite := widget.NewProgressBarInfinite()
	infinite.Hide()

	scanButton := widget.NewButton("SCAN", func() {
		ip := input.Text
		fmt.Println(ip)
		infinite.Start()
		infinite.Show()
	})

	cancelButton := widget.NewButton("CANCEL", func() {

		fmt.Println("CANCELED")
	})
	cancelButton.Resize(fyne.NewSize(1, 20))

	hbox2 := container.NewGridWithColumns(2, scanButton, cancelButton)
	mainWindow.SetContent(container.NewVBox(heading, spacer, hbox1, spacer, hbox2, spacer, infinite))
	mainWindow.Resize(fyne.NewSize(400, 300))

	mainWindow.ShowAndRun()

}
