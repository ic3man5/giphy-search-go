// https://go.dev/doc/tutorial/getting-started
// https://gioui.org/doc/learn/get-started
package main

import (
	"fmt"
	"giphy-search-go/gif"
	"image/color"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/widget/material"
)

func main() {
	go func() {
		w := app.NewWindow()
		if err := run(w); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	results := gif.Search("Hello World")
	for i, d := range results {
		fmt.Println(i, d)
	}
	app.Main()
}

func run(w *app.Window) error {
	th := material.NewTheme(gofont.Collection())
	var ops op.Ops
	for {
		e := <-w.Events()
		switch e := e.(type) {
		case system.DestroyEvent:
			return e.Err
		case system.FrameEvent:
			gtx := layout.NewContext(&ops, e)
			title := material.H1(th, "Hello, Gio")
			maroon := color.NRGBA{R: 127, G: 0, B: 0, A: 255}
			title.Color = maroon
			title.Alignment = text.Middle
			title.Layout(gtx)

			e.Frame(gtx.Ops)

		}
	}
}
