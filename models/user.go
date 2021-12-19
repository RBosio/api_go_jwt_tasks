package models

import "packages/db"

type User struct {
	Id       int64  `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
}

func GetUserByEmail(email string) (*User, error) {
	sql := "SELECT * FROM user WHERE email = ?"
	rows, err := db.Query(sql, email)
	if err != nil {
		return nil, err
	} else {
		user := User{}
		for rows.Next() {
			rows.Scan(&user.Id, &user.Email, &user.Password, &user.Name, &user.Surname)
		}
		return &user, nil
	}
}

func (user *User) NewUser() {
	sql := "INSERT INTO user (email, password, name, surname) VALUES (?, ?, ?, ?)"
	result, _ := db.Exec(sql, user.Email, user.Password, user.Name, user.Surname)
	user.Id, _ = result.LastInsertId()
}
