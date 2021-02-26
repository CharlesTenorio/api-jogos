package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasTime = []Rota{
	{
	  URI: "/times",
	  Metodo: http.MethodGet,
	  Funcao: controllers.ListaTodosTimes,
	  RequerAutenticacao:false,
	},

	  {
		URI: "/times/{timeID}",
		Metodo: http.MethodGet,
		Funcao: controllers.ListaTtime,
		RequerAutenticacao:false,
	  },
	  
	  {
			URI: "/time/",
			Metodo: http.MethodPost,
			Funcao: controllers.CriarTime,
			RequerAutenticacao:false,
	  },
	  
	  

},