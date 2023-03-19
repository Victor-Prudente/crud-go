package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	id        int
	nome      string
	sobrenome string
}

func main() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/crudgojava")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Testando a função createUser
	user1 := User{nome: "João", sobrenome: "Silva"}
	id, err := createUser(db, user1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Usuário criado com ID %d\n", id)

	// Testando a função getUsers
	users, err := getUsers(db)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Lista de usuários: %v\n", users)

	// Testando a função updateUser
	user2 := User{id: 1, nome: "Maria", sobrenome: "Silveira"}
	err = updateUser(db, user2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Usuário atualizado com sucesso")

	// Testando a função getUserById
	user3, err := getUserById(db, 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Usuário encontrado: %v\n", user3)

	// Testando a função deleteUser
	err = deleteUser(db, 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Usuário deletado com sucesso")
}
