package main

import (
	"log"
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/driver/desktop"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"golang.org/x/image/colornames"
)

func main() {
	//Win Init
	myApp := app.New()

	myWindow := myApp.NewWindow("Proyecto LuisFlahan4051 Registro de usuarios")

	myDriverApp := fyne.CurrentApp().Driver()
	driverDesktop, _ := myDriverApp.(desktop.Driver)
	windowSplash := driverDesktop.CreateSplashWindow()

	//Elementos
	inputNombre := widget.NewEntry()
	inputNombre.SetPlaceHolder("Nombre completo...")
	formItemNombre := widget.NewFormItem("Nombre:", inputNombre)

	inputApellidos := widget.NewEntry()
	inputApellidos.SetPlaceHolder("Ingrese sus dos apellidos...")
	formItemApellidos := widget.NewFormItem("Apellidos:", inputApellidos)

	form := widget.NewForm(
		formItemNombre,
		formItemApellidos,
	)
	form.OnSubmit = func() {
		log.Println("Form submited:", inputNombre.Text)
	}
	form.Resize(fyne.NewSize(200, 200))

	//--
	txtBienvenida := canvas.NewText("~* ¡Bienvenido! *~", colornames.Cyan)
	txtBienvenida.Alignment = fyne.TextAlignCenter
	txtBienvenida.TextStyle = fyne.TextStyle{Bold: true}
	txtBienvenida.TextSize = 38

	image := canvas.NewImageFromFile("src/luisflahan4051apps.png")
	image.FillMode = canvas.ImageFillOriginal

	//Layouts init
	centrado := layout.NewCenterLayout()
	horizontal := layout.NewHBoxLayout()
	grid := layout.NewGridLayout(2)

	//Disposiciones
	cajaCentradaFormulario := fyne.NewContainerWithLayout(centrado,
		widget.NewVBox(
			form,
			widget.NewLabel("Aquí abajo"),
		),
	)
	cajaHorizontalTabla := fyne.NewContainerWithLayout(horizontal,
		widget.NewLabel("Aquí va la tabla"),
		widget.NewLabel("Aquí al lado"),
	)
	contenedorPrincipal := fyne.NewContainerWithLayout(grid,
		cajaCentradaFormulario,
		cajaHorizontalTabla,
	)

	//--
	contenedorBienvenida := fyne.NewContainerWithLayout(centrado,
		widget.NewVBox(
			txtBienvenida,
			image,
		),
	)

	//Win Config

	myWindow.Resize(fyne.NewSize(1000, 580))
	myWindow.SetFixedSize(true)
	myWindow.SetContent(contenedorPrincipal)

	//----
	windowSplash.Resize(fyne.NewSize(500, 300))
	windowSplash.CenterOnScreen()
	windowSplash.SetContent(contenedorBienvenida)

	//Show and run
	windowSplash.Show()

	go func() {
		time.Sleep(time.Second * 3)
		windowSplash.Hide()
		myApp.Settings().SetTheme(theme.DarkTheme())
		myWindow.Show()
		windowSplash.Close()
	}()

	myApp.Settings().SetTheme(theme.LightTheme())
	myApp.Run()
}
