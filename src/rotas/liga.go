package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasLiga = []Rota{
	{
		URI:                "/ligas",
		Metodo:             http.MethodGet,
		Funcao:             controllers.ListaTodasLigas,
		RequerAutenticacao: false,
	},

	{
		URI:                "/ligas/{ligaID}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.ListaLiga,
		RequerAutenticacao: false,
	},

	{
		URI:                "/liga",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarLiga,
		RequerAutenticacao: false,
	},

	{
		URI:                "/ligas/{ligaID}",
		Metodo:             http.MethodPost,
		Funcao:             controllers.AutalizarLiga,
		RequerAutenticacao: false,
	},

	{
		URI:                "/liga_del/{ligaID}",
		Metodo:             http.MethodPost,
		Funcao:             controllers.DeletarLiga,
		RequerAutenticacao: false,
	},
}
