package poems

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	th "github.com/QIN2DIM/fyne-poems/theme"
	"image/color"
	"log"
)

const (
	Weight = 720
	Height = 192
)

type Pane struct {
	app    fyne.App
	window fyne.Window
}

func NewPane() *Pane {
	pane := &Pane{app: app.New()}
	pane.init()
	return pane
}

func (pane *Pane) init() {
	pane.app.Settings().SetTheme(&th.MyTheme{})

	// Setup window
	if dev, ok := fyne.CurrentApp().Driver().(desktop.Driver); ok {
		pane.window = dev.CreateSplashWindow()
		pane.window.Resize(fyne.NewSize(Weight, Height))
	} else {
		log.Fatal("暂不支持在移动端部署应用")
	}
}

func (pane *Pane) Startup(s *Selector) {
	pane.loadGUI(s)
	pane.window.ShowAndRun()
}

func (pane *Pane) loadGUI(s *Selector) {
	// Display area of corpus 18
	corpusBoard := canvas.NewText(s.GetCurrentCorpus(), color.Black)
	corpusBoard.Alignment = fyne.TextAlignCenter
	corpusBoard.TextSize = 20
	corpusBoard.TextStyle = fyne.TextStyle{Bold: true}

	// Render previous corpus
	preButton := widget.NewButton("Previous", func() {
		corpusBoard.Text = s.GetPreviousCorpus()
		corpusBoard.Refresh()
		pane.window.Resize(fyne.NewSize(Weight, Height))
	})

	// Quit process
	closeButton := widget.NewButton("Close", func() {
		pane.window.Close()
	})

	// Render Next corpus
	nextButton := widget.NewButton("Next", func() {
		corpusBoard.Text = s.GetNextCorpus()
		corpusBoard.Refresh()
		pane.window.Resize(fyne.NewSize(Weight, Height))
	})

	// Line up three buttons
	hButtonBox := container.NewHBox(
		layout.NewSpacer(),
		preButton,
		layout.NewSpacer(),
		closeButton,
		layout.NewSpacer(),
		nextButton,
		layout.NewSpacer(),
	)
	hSeparatorBox := container.NewHBox(
		layout.NewSpacer(),
		corpusBoard,
		layout.NewSpacer(),
	)
	vBox := container.NewVBox(
		layout.NewSpacer(),
		hSeparatorBox,
		layout.NewSpacer(),
		hButtonBox,
		layout.NewSpacer(),
	)

	pane.window.SetContent(vBox)
}
