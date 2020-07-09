package main

import (
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/driver/desktop"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

func main() {
	//Win Init and congfig
	myApp := app.New()
	myWindow := myApp.NewWindow("Proyecto LuisFlahan4051 Registro de usuarios")

	myDriverApp := fyne.CurrentApp().Driver()
	driverDesktop, _ := myDriverApp.(desktop.Driver)
	windowSplash := driverDesktop.CreateSplashWindow()

	//widgets
	btn1 := widget.NewButton("btn1", func() {
		w := fyne.CurrentApp().Driver().CreateWindow("Hello2")
		w.Show()
	})

	//Layout
	row1 := layout.NewVBoxLayout()
	contenedor := fyne.NewContainerWithLayout(row1, btn1)

	//Win Config

	myWindow.Resize(fyne.NewSize(1000, 580))
	myWindow.SetFixedSize(true)
	myWindow.SetContent(contenedor)

	windowSplash.Resize(fyne.NewSize(400, 200))
	windowSplash.SetContent(
		widget.NewLabelWithStyle(
			"Bienvenido!\n\n*~ Luisflahan4051 APPS ~*",
			fyne.TextAlignCenter, fyne.TextStyle{Bold: true},
		),
	)

	//Show and run
	windowSplash.Show()

	go func() {
		time.Sleep(time.Second * 3)
		windowSplash.Hide()
		myWindow.Show()
		windowSplash.Close()
	}()

	myApp.Run()
}
