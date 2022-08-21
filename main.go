package main

import "github.com/QIN2DIM/rainbow-fart/poems"

func main() {
	pane := poems.NewPane()
	pane.Startup(poems.NewSelector())
}
