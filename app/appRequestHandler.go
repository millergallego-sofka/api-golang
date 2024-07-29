package app

import (
	infrastructure "apiRest/infraestructure"
	model "apiRest/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func GetUsers(rw http.ResponseWriter, r *http.Request) {
	infrastructure.Connect()
	defer infrastructure.Close()
	//users := model.ListUser()
	//infrastructure.Close()
	//marshal lo que hace es volver datos en formato json
	//output, _ := json.Marshal(users)
	//forma de responder un string a la peticion que se hace
	//fmt.Fprint(rw, string(output))

	//responder en json
	//rw.Header().Set("Content-Type", "application/json")
	//rw.Write(output)

	if users, err := model.ListUser(); err != nil {
		model.SendNoFound(rw)
	} else {
		model.SenData(rw, users)
	}
}

func GetUser(rw http.ResponseWriter, r *http.Request) {
	infrastructure.Connect()
	defer infrastructure.Close()

	/*

		vars := mux.Vars(r)
		userId, _ := strconv.Atoi(vars["id"])
		user := model.GetuserById(userId)
		output, _ := json.Marshal(user)
		rw.Header().Set("Content-Type", "application/json")
		rw.Write(output)

	*/

	if user, err := getUserByRequest(r); err != nil {
		model.SendNoFound(rw)
	} else {
		model.SenData(rw, user)
	}

}

func CreateUser(rw http.ResponseWriter, r *http.Request) {
	infrastructure.Connect()
	defer infrastructure.Close()

	user := model.User{}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		/*
			rw.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(rw, "{\"error\":\"%v\"}", err)
			return
		*/
		model.SendNoFound(rw)
	} else {
		user.Save()
		model.SenData(rw, user)
	}

	/*
		user.Save()
		output, _ := json.Marshal(user)
		rw.Header().Set("Content-Type", "application/json")
		rw.Write(output)
	*/

}

func UpdateUser(rw http.ResponseWriter, r *http.Request) {
	infrastructure.Connect()
	defer infrastructure.Close()
	user := model.User{}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(rw, "{\"error\":\"%v\"}", err)
		return
	}

	user.Save()
	output, _ := json.Marshal(user)
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(output)

}

func DeleteUser(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "Hello World")
}

func getUserByRequest(r *http.Request) (model.User, error) {
	vars := mux.Vars(r)
	userId, _ := strconv.Atoi(vars["id"])

	if user, err := model.GetuserById(userId); err != nil {
		return model.User{}, err
	} else {
		return *user, nil
	}
}
