package infrastructure

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

//username:password@tcp(localhost:3306)/database

const url = "root:1234@tcp(localhost:3306)/golang"

var db *sql.DB

func Connect() {
	connection, err := sql.Open("mysql", url)
	if err != nil {
		panic(err)
	}
	fmt.Println("conexion exitosa")
	db = connection
}

func Close() {
	db.Close()
}

func VerifyConnect() {
	if err := db.Ping(); err != nil {
		panic(err)
	}
}

/*
	se utilizan 2 diferentes metodos para volver el lenguaje SQL en comandos que se pueden ejecutar en Go
	Query -> se utiliza principalmente para hacer consultas de la bd ya que se retorna un objeto tipo *sql.Rows
		que basicamente son las filas de la data para iterar y manipular los datos
	Exec -> este segundo metodo acepta de igual manera codigo SQL pero en este caso no retorna columnas
		se utiliza mayormente para hacer operaciones de modificacion creacion o eliminacion
*/

func Execute(query string, args ...interface{}) (sql.Result, error) {
	return db.Exec(query, args...)
}

func Query(query string, args ...interface{}) (*sql.Rows, error) {
	return db.Query(query, args...)
}

func CreateTable(schema string, nameTable string) {
	if !GetTable(nameTable) {
		//esto lo que hace es ejecutar alguna sentencia sql
		_, err := db.Exec(schema)
		if err != nil {
			fmt.Println("Error creating table", err)
		}
		fmt.Println("Create table success")
	}
}

func GetTable(name string) bool {
	sql := fmt.Sprintf("SHOW TABLES LIKE '%s'", name)
	// lo que hace el query es obtener sentencias SQL
	rows, err := db.Query(sql)
	if err != nil {
		fmt.Println("Error", err.Error())
	}
	return rows.Next()
}

func DeleteRows(nameTable string) {
	//sPrint funciona para formatear el texto forma de concatenar
	sql := fmt.Sprintf("TRUNCATE %s", nameTable)
	_, err := db.Exec(sql)
	if err != nil {
		fmt.Println("error deleting rows")
	}
}
