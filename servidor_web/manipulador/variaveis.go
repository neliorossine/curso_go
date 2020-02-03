package manipulador

import "text/template"

//Modelos armazenam os modelos de página que serão executados pelos manipuladores
var Modelos = template.Must(template.ParseFiles("html/ola.html"))
