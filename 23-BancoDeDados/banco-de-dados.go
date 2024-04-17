package main

import (
	"database/sql"
	"fmt"
	"log"

	// importa com _ na frente mas pra indicar que vai ser usado se forma implicita
	_"github.com/go-sql-driver/mysql"
)

func main() {
	// Formato: "usuario:senha@/banco"
	// charset: reconhecer caracteres latinos
	stringConexao := "golang:golang@/devbook?charset=utf8&parseTime=True&loc=Local"
	db, erro := sql.Open("mysql", stringConexao)
	if erro != nil {
		fmt.Println("Dentro do SQL.Open!")
		log.Fatal(erro)
	}
	defer db.Close() 
	
	if erro = db.Ping(); erro != nil {
		fmt.Println("Dentro do ping")
		log.Fatal(erro)
	}

	fmt.Println("Conexão está aberta!")

	linhas, erro := db.Query("select * from usuarios")
	if erro != nil {
		log.Fatal(erro)
	}
	defer linhas.Close()

	fmt.Println(linhas)

}
