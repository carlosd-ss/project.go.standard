package handler

import (
	"encoding/json"
	"net/http"

	"github.com/jeffotoni/project.go.standard/project.web/standard.libray/crud.user.singleton/model/usermodel"
	"github.com/jeffotoni/project.go.standard/project.web/standard.libray/crud.user.singleton/repo/user"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var p usermodel.User
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		jsonmsg := "Error body!"
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(jsonmsg)
		return
	}

	if p.Age < 0 {
		jsonmsg := "Campo Age obrigatorio!"
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(jsonmsg)
		return
	}

	usr := user.InsertUser(p)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(usr)
	jsonmsg := "usuario cadastrado"
	json.NewEncoder(w).Encode(jsonmsg)
}
