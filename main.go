package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

const ChpAPI = "https://api.shadiao.pro/chp"

func main() {
	a := app.New()

	win := a.NewWindow("彩虹屁生成器")
	win.Resize(fyne.NewSize(480, 192))
	win.CenterOnScreen()

	win.SetContent(widget.NewLabel("content"))

	win.ShowAndRun()
}
