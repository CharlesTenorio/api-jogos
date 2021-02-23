package rotas

import "net/http"

var rotasTime = []Rota{
	{
	  URI: "/times",
	  Metodo: http.MethodGet,
	  Funcao: func(w http.ResponseWriter, r *http.Request){

	  },
	  RequerAutenticacao:false,

},