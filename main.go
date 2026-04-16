package main

import (
	"fmt"
	"image/color"
	"net"
	"strconv"
	"sync"
	"time"

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
		return color.NRGBA{R: 30, G: 35, B: 45, A: 255}
	}

	if name == theme.ColorNameForeground {
		return color.NRGBA{R: 230, G: 235, B: 240, A: 255}
	}

	if name == theme.ColorNamePrimary {

		return color.NRGBA{R: 255, G: 170, B: 0, A: 255}
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

type IPStatus struct {
	port   string
	status string
}

func scanner(ip string, in <-chan string, results chan IPStatus, scanWg *sync.WaitGroup) {

	scanWg.Go(func() {
		for port := range in {

			conn, err := net.DialTimeout("tcp", ip+":"+port, 2000*time.Millisecond)
			if err != nil {
				// fmt.Println("\x1b[31m", err)
				continue
			}

			results <- IPStatus{port: port, status: "OPEN"}
			conn.Close()
		}

	})

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
	data := []IPStatus{}

	list := widget.NewList(
		func() int {
			return len(data)
		},
		func() fyne.CanvasObject {

			t := canvas.NewText("", color.NRGBA{R: 0, G: 255, B: 195, A: 255})
			t.Alignment = fyne.TextAlignLeading
			t.TextStyle.Bold = true
			t.TextSize = 20
			return t
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {

			result := o.(*canvas.Text)

			result.Text = "[+]Port = " + data[i].port + " Status = " + data[i].status
			// result.Color = color.NRGBA{R: 0, G: 255, B: 0, A: 255}
		},
	)

	list.Resize(fyne.NewSize(20, 100))

	scanButton := widget.NewButton("SCAN", func() {
		// ip := input.Text
		// fmt.Println(ip)

		go func() {
			fyne.Do(func() {
				infinite.Show()
				infinite.Start()
			})

		}()

		// data = append(data, 5)
		list.Refresh()
		ip := input.Text
		data = []IPStatus{}
		// list.Refresh()
		go func() {
			var wg sync.WaitGroup
			var scanWg sync.WaitGroup
			var ports []string
			portChan := make(chan string, 100)
			resultsChan := make(chan IPStatus)
			ports = append(ports, "8080", "3389", "1443", "3306", "3389", "5900", "9050", "5432")

			wg.Go(func() {
				for i := 1; i <= 1024; i++ {
					portChan <- strconv.Itoa(i)
				}
				for _, p := range ports {
					portChan <- p
				}
				defer close(portChan)
			})

			for i := 0; i < 100; i++ {
				scanner(ip, portChan, resultsChan, &scanWg)
			}

			go func() {
				scanWg.Wait()
				close(resultsChan)
			}()

			for res := range resultsChan {
				data = append(data, res)
				go func() {
					fyne.Do(func() {
						list.Refresh()
					})
				}()

			}
			fyne.Do(func() {
				infinite.Stop()
				infinite.Hide()
			})

		}()

	})

	cancelButton := widget.NewButton("CANCEL", func() {
		infinite.Stop()
		infinite.Hide()
		fmt.Println("CANCELED")
	})
	cancelButton.Resize(fyne.NewSize(1, 20))

	hbox2 := container.NewGridWithColumns(2, scanButton, cancelButton)
	controls := container.NewVBox(heading, spacer, hbox1, spacer, hbox2, spacer, infinite, spacer)

	footer := canvas.NewText("Made by A4BHI", color.NRGBA{R: 255, G: 170, B: 0, A: 255})
	footer.TextSize = 10
	footer.TextStyle.Italic = true
	footer.Alignment = fyne.TextAlignCenter

	mainLayout := container.NewBorder(controls, footer, nil, nil, list)
	mainWindow.SetContent(mainLayout)
	mainWindow.Resize(fyne.NewSize(400, 300))

	mainWindow.ShowAndRun()

}
