package main

import(
	"fmt"
	"github.com/google/uuid"
	"database/sql"
	"log"
	_"github.com/go-sql-driver/mysql"
)

type Pessoa struct{
	ID string
	Nome string
	Idade float64
} 

func main(){

	fmt.Println("trabalhando com DB")

	p := Pessoa{
 
		ID: uuid.New().String(),
		Nome: "Joao",
	    Idade: 10.1,
 	}
	fmt.Println(p)

	//Abrir nonexão com mysql
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/aula")
	if  err != nil{
		log.Println("erro ao fazer conexão", err)
	}

	log.Println("Deu certo")

	defer db.Close()


	ponhar, err := db.Prepare("insert into aula(id, nome, idade) values(?,?,?)")
	if err != nil{
		log.Println("Erro ao insserir dados no banco de dados", err)
	}

	_, err = ponhar.Exec(p.ID, p.Nome, p.Idade)
	if err != nil{
		log.Println(err)
	}

	fmt.Println("Deu certo")
	
	defer ponhar.Close()



}