package controllers

import (
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"api/src/respostas"
	"encoding/json"
	"io"
	"net/http"
)

//CriarUsuario insere um usuario no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request){
	corpoRequest, erro := io.ReadAll(r.Body)
	// Erro caso não consiga ler o corpo da requisição
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var usuario modelos.Usuario
	// Erro caso não consiga jogar na struct de usuário
	if erro = json.Unmarshal(corpoRequest, &usuario); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = usuario.Preparar(); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	// Erro caso não consiga abrir conexão com o banco
	// tem a ver com servidor,\ não com reoquisição
	db, erro := banco.Conectar() 
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	// Quando a função do controller terminar oque a conexao com o banco seja fechada
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	usuario.ID, erro = repositorio.Criar(usuario)
	// Erro ao não conseguir criar
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	// Devolve o usuário que foi inserido
	respostas.JSON(w, http.StatusCreated, usuario)
	
}

//BuscarUsuarios busca todos os usuarios salvos no banco
func BuscarUsuarios(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Buscando todos os usuários"))
}

//BuscarUsuarios busca um usuario salvo no banco
func BuscarUsuario(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Buscando um Usuário!"))
}

//AtualizarUsuario altera as informacoes de um usuario no banco
func AtualizarUsuario(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Atualizando Usuário!"))
}

//DeletarUsuario exclui as informações de um usuário no banco
func DeletarUsuario(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Deletando Usuário!"))
}