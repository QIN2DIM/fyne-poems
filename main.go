package main

import "github.com/QIN2DIM/fyne-poems/poems"

func main() {
	pane := poems.NewPane()
	pane.Startup(poems.NewSelector())
}
