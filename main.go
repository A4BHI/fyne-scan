package main

import (
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

	inputlabel := widget.NewLabel("Target IP/Domain : ")
	inputlabel.TextStyle = fyne.TextStyle{Monospace: true}

	input := widget.NewEntry()
	input.SetPlaceHolder("Target IP/Domain")

	hbox := container.NewGridWithColumns(2, inputlabel, input)
	mainWindow.SetContent(container.NewVBox(heading, hbox))
	mainWindow.Resize(fyne.NewSize(400, 300))

	mainWindow.ShowAndRun()

}
