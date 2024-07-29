package model

import infrastructure "apiRest/infraestructure"

type User struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type Users []User

const UserSchema string = `CREATE TABLE users_golang (
    id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(50) NOT NULL,
    password VARCHAR(100) NOT NULL,
    email VARCHAR(100),
    created_at TIMESTAMP NOT NULL)`

func Newuser(username string, password string, email string) *User {
	return &User{Username: username, Password: password, Email: email}
}

func CreateUser(username string, password string, email string) *User {
	user := Newuser(username, password, email)
	user.insert()
	return user
}

func (user *User) insert() {
	sql := "INSERT INTO users_golang (username, password, email) VALUES(?,?,?)"
	result, _ := infrastructure.Execute(sql, user.Username, user.Password, user.Email)
	user.Id, _ = result.LastInsertId()
}

func ListUser() (Users, error) {
	sql := "SELECT id, username, password, email FROM users_golang"
	users := Users{}
	rows, err := infrastructure.Query(sql)
	for rows.Next() {
		user := User{}
		//
		rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email)
		users = append(users, user)
	}
	return users, err
}

func GetuserById(id int) (*User, error) {
	sql := "SELECT id, username, password, email FROM users_golang WHERE id=?"
	newUser := Newuser("", "", "")
	rows, err := infrastructure.Query(sql, id)
	for rows.Next() {
		rows.Scan(&newUser.Id, &newUser.Username, &newUser.Password, &newUser.Email)
	}
	return newUser, err
}

func (user *User) updateUser() {
	sql := "UPDATE users_golang SET username=?, password=?, email=? WHERE id=?"
	infrastructure.Execute(sql, user.Username, user.Password, user.Email, user.Id)
}

func (user *User) Save() {
	if user.Id == 0 {
		user.insert()
	} else {
		user.updateUser()
	}
}
