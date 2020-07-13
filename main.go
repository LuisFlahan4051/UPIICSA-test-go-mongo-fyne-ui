package main

import (
	"context"
	"log"
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/driver/desktop"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"golang.org/x/image/colornames"
)

func main() {
	//--------------Win Init------------------
	myApp := app.New()

	myWindow := myApp.NewWindow("Proyecto LuisFlahan4051 Registro de usuarios")

	myDriverApp := fyne.CurrentApp().Driver()
	driverDesktop, _ := myDriverApp.(desktop.Driver)
	windowSplash := driverDesktop.CreateSplashWindow()
	//------------ Base de datos init-----------------
	uri := "mongodb://localhost:27017"
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	//Ping
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}
	log.Println("Successfully connected and pinged.")
	//---
	coleccionUsuarios := client.Database("proyecto-registros_db").Collection("usuarios")

	//---------------Elementos-----------------
	inputNombre := widget.NewEntry()
	inputNombre.SetPlaceHolder("Nombre completo...")
	formItemNombre := widget.NewFormItem("Nombre:", inputNombre)

	inputApellidos := widget.NewEntry()
	inputApellidos.SetPlaceHolder("Ingrese sus dos apellidos...")
	formItemApellidos := widget.NewFormItem("Apellidos:", inputApellidos)

	inputEdad := widget.NewCheck("Soy mayor de edad.", func(value bool) {
		log.Println(value)
	})
	formItemEdad := widget.NewFormItem("Edad:", inputEdad)

	inputSexo := widget.NewRadio([]string{"Hombre", "Mujer"}, func(value string) {
		log.Println(value)
	})
	formItemSexo := widget.NewFormItem("Sexo:", inputSexo)

	inputEmail := widget.NewEntry()
	inputEmail.SetPlaceHolder("email_ejemplo@mail.com")
	formItemEmail := widget.NewFormItem("Email:", inputEmail)

	inputPassword := widget.NewPasswordEntry()
	formItemPassword := widget.NewFormItem("Contraseña:", inputPassword)

	formStatus := widget.NewLabel("")
	formStatus.TextStyle = fyne.TextStyle{Italic: true}

	form := widget.NewForm(
		formItemNombre,
		formItemApellidos,
		formItemEdad,
		formItemSexo,
		formItemEmail,
		formItemPassword,
	)
	form.OnSubmit = func() {
		log.Println("Form submited:", inputNombre.Text)
		formStatus.Text = "Datos guardados correctamente!"
		formStatus.Refresh()
		go func() {
			time.Sleep(time.Second * 2)
			formStatus.Text = ""
			formStatus.Refresh()
		}()
	}
	form.Resize(fyne.NewSize(200, 200))

	//--
	txtBienvenida := canvas.NewText("~* ¡Bienvenido! *~", colornames.Cyan)
	txtBienvenida.Alignment = fyne.TextAlignCenter
	txtBienvenida.TextStyle = fyne.TextStyle{Bold: true}
	txtBienvenida.TextSize = 38

	image := canvas.NewImageFromFile("src/luisflahan4051apps.png")
	image.FillMode = canvas.ImageFillOriginal

	barraProgreso := widget.NewProgressBarInfinite()

	//----------------Layouts init---------------------
	centrado := layout.NewCenterLayout()
	vertical := layout.NewVBoxLayout()
	grid := layout.NewGridLayout(2)

	//----------------Disposiciones--------------------
	contenedorFormulario := fyne.NewContainerWithLayout(centrado,
		widget.NewVBox(
			form,
			formStatus,
		),
	)

	tituloTabla := widget.NewLabel("Usuarios:")
	tituloTabla.TextStyle = fyne.TextStyle{Bold: true}
	cajaColumnasUsuarios := widget.NewHBox(
		tituloTabla,
	)
	cajaDatosUsuarios := widget.NewVBox()
	contenedorTabla := fyne.NewContainerWithLayout(vertical,
		cajaColumnasUsuarios,
		cajaDatosUsuarios,
	)

	contenedorPrincipal := fyne.NewContainerWithLayout(grid,
		contenedorFormulario,
		contenedorTabla,
	)

	//--
	contenedorBienvenida := fyne.NewContainerWithLayout(centrado,
		widget.NewVBox(
			txtBienvenida,
			image,
			barraProgreso,
		),
	)

	//---------------Consultas MongoDB-----------------------
	edad := ""
	puntero, err := coleccionUsuarios.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer puntero.Close(ctx)
	for puntero.Next(ctx) {
		// To decode into a struct, use cursor.Decode()
		resultado := struct {
			Nombre    string
			Apellidos string
			Edad      bool
			Sexo      string
			Email     string
			Password  string
		}{}
		err := puntero.Decode(&resultado)
		if err != nil {
			log.Fatal(err)
		}
		// do something with result...
		if resultado.Edad == true {
			edad = "Mayor de edad"
		} else {
			edad = "Menor de edad"
		}
		cajaDatosUsuarios.Append(
			widget.NewHScrollContainer(
				widget.NewHBox(
					widget.NewLabel(">   "+resultado.Nombre+" "+resultado.Apellidos+";"),
					widget.NewLabel("-   "+edad+";"),
					widget.NewLabel("-   "+resultado.Sexo+";"),
					widget.NewLabel("-   "+resultado.Email+";"),
					widget.NewLabel("-   "+resultado.Password+";"),
				),
			),
		)
		// To get the raw bson bytes use cursor.Current
		//raw := puntero.Current
		// do something with raw...
	}
	if err := puntero.Err(); err != nil {
		log.Fatal(err)
	}

	//---------------Win Config--------------------

	myWindow.Resize(fyne.NewSize(1000, 580))
	myWindow.SetFixedSize(true)
	myWindow.SetContent(contenedorPrincipal)

	//----
	windowSplash.Resize(fyne.NewSize(500, 400))
	windowSplash.CenterOnScreen()
	windowSplash.SetContent(contenedorBienvenida)

	//---------------Show and run, goroutines------------------

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
