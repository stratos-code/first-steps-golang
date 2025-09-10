# Primeros pasos con GO 
Basado en este video https://youtu.be/8uiZC0l4Ajw

## Modulos y paquetes
Un paquete es simplemente una carpeta que contiene código fuente. cada archivo de go debe pertenecer a un paquete. La colección de paquetes son los módulos

Para crear un nuevo módulo, se debe ejecutar 
`go mod init {nombe_modulo}`

Por conveniencia el módulo principal suele tener el nombre del repositorio pero se puede usar cualquier otro nombre

El comando nos crea un archivo con el nombre del módulo definido y la versión de go. Aquí también irían los paquetes externos.

## Primer paquete
Al crear unas carpetas y añadir un código fuente .go si se llega a tener un archivo main.go y dentro de el se define como el paquete principal al escribir `package main` le estamos diciendo al compilador que este es el entrypoint donde comenzar todo el programa

## Compilar
Usamos `go build {path}` para compilar un paquete

## Valor por defecto
Según el tipo de dato, go define valores por defecto

* Para todos los enteros y decimales, el valor por defecto es cero
* Para strings, es vacío
* para bool es false

## Inferencia
Go puede inferir el tipo de una variable.

`var myVar = "text"` Sabe que es un string

pero también se puede usar esta sintaxis corta

`myVar := "text"`

`var1, var2 := 1, 2`

## arrays
los arrays siempre son continuos en memoria


# Paquetes externos
se puede ejecutar un comando para que revise y actualice el archivo go.mod para añadir las dependencias que se necesiten en el código... si alguna dependencia ya no existe en el código, automáticamente las elimina

`go mod tidy`