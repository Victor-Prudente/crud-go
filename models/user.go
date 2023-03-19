package models

import "database/sql"

type User struct {
	id        int
	nome      string
	sobrenome string
}

func createUser(db *sql.DB, user User) (int64, error) {
	stmt, err := db.Prepare("INSERT INTO users(nome, email) VALUES(?, ?)")
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(user.nome, user.sobrenome)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func updateUser(db *sql.DB, user User) error {
	stmt, err := db.Prepare("UPDATE users SET nome=?, email=? WHERE id=?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(user.nome, user.sobrenome, user.id)
	if err != nil {
		return err
	}

	return nil
}

func deleteUser(db *sql.DB, id int) error {
	stmt, err := db.Prepare("DELETE FROM users WHERE id=?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}

func getUsers(db *sql.DB) ([]User, error) {
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []User

	for rows.Next() {
		var user User

		err := rows.Scan(&user.id, &user.nome, &user.sobrenome)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func getUserById(db *sql.DB, id int) (User, error) {
	var user User

	err := db.QueryRow("SELECT * FROM users WHERE id=?", id).Scan(&user.id, &user.nome, &user.sobrenome)
	if err != nil {
		return user, err
	}

	return user, nil
}
