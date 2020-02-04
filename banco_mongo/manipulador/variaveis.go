package manipulador

import "text/template"

//ModeloOla armazenam os modelos de página ola
var ModeloOla = template.Must(template.ParseFiles("html/ola.html"))

//ModeloLocal armazenam os modelos de página Local
var ModeloLocal = template.Must(template.ParseFiles("html/local.html"))
