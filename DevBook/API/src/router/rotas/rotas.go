package rotas

import (
	"net/http"

	"github.com/gorilla/mux"
)

//Rota representa todas as rotas da AO
type Rota struct {
	URI                string
	Metodo             string
	Funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}

// Retorna Router com todas as rotas configuradas
// Recebe parâmetro mux.Router (que não tem nenhuma rota dentro) e vai retornar o mux.Router
// Vai configurar todas as  rotas e devolver o Router pronto

// Configurar coloca todas as rotas dentro do Router
func Configurar(r *mux.Router) *mux.Router {
	rotas := rotaUsuarios

	for _, rota := range rotas {
		r.HandleFunc(rota.URI, rota.Funcao).Methods(rota.Metodo)
	}

	return r
}
