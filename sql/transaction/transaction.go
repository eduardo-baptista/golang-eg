package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:mysql@/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// comeca uma transacao
	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("insert into usuarios(id, nome) values(?,?)")

	stmt.Exec(2000, "Bia")
	stmt.Exec(2001, "Carlos")
	_, err = stmt.Exec(1, "Tiago") // chave duplicada

	if err != nil {
		// é obrigatorio tratar o erro caso contrario sera comitado as mudancas
		tx.Rollback()
		log.Fatal(err)
	}

	tx.Commit()
}
