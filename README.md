# Proyecto Simple Registradora

Proyecto escolar del instituto:
IPN - UPIICSA

Profesor: 
José Luis Lopez Goytia.

## Capturas
![alt text](captura1.png)
![alt text](captura2.png)
![alt text](captura3.png)
![alt text](captura4.png)

Video demostrativo:

https://drive.google.com/file/d/1FcKDgtt9bOPkCuHLCQSvExwZYBjXAFyx/view?usp=sharing

## Tecnologías usadas 🛠️

Lenguaje de programación **Golang**

Interfaz GUI **Fyne**

Base de Datos **MongoDB**

Controlador de versiones **Git**

### Pre-requisitos de compilación 📋

Golang

https://golang.org/doc/install

MongoDB

https://www.mongodb.com/try/download/community

Librería de Fyne-Go

https://fyne.io/

Driver de conexión mongodb-Go

https://docs.mongodb.com/drivers/go

**Despues de instalar Go:**
```
$ go get -u github.com/go-gl/glfw/v3.3/glfw
$ go get fyne.io/fyne
$ go get fyne.io/fyne/cmd/fyne
$ go get go.mongodb.org/mongo-driver/mongo
```

### Instalación 🔧

Para ejecutar en **windows**:

Descargar el repositorio.
```
$ cd C:\Users\%USERNAME%\Desktop
$ git clone https://github.com/LuisFlahan4051/proyecto-registros
$ cd proyecto-registros
```

Despues de instalar MongoDB y agregar al PATH los binarios correspondientes podremos ingresar las siguientes lineas para alzar el servidor:

**Nota:** En caso de que no exista la carpeta data en C:/ vamos a crearla.
```
$ cd C:\
$ mkdir data
$ cd C:\data
$ mkdir db
```

Ahora vamos a alzar el servidor en un CMD aparte con la siguiente linea:
```
$ mongod
```

Ahora podemos ejecutar el **goproj.exe** dentro de la carpeta del proyecto.

De la misma manera en linux podemos alzar el revidor y abrir el programa con la siguiente línea:
```
$ cd $HOME/Escritorio
$ git clone https://github.com/LuisFlahan4051/proyecto-registros
$ cd proyecto-registros
$ mongod
```
En otro shell:
```
$ cd $HOME/Escritorio/proyecto-registros
$ ./goproj
```
## Autores ✒️
Secuencia: **1CM12**

**Melendez Bustamante Luis Fernando.**
2020602568.
https://github.com/luisflahan4051

**Ávila Flores Ricardo.**
2020602143.
https://github.com/RichiePeek

**Pérez Cuevas Ivan.**
2020601813.
https://github.com/Ivanpc234

## Licencia 📄

GNU GPL