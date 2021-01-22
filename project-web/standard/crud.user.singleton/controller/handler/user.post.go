package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/jeffotoni/project.go.standard/project.web/standard.libray/crud.user.singleton/internal/fmts"
	mUser "github.com/jeffotoni/project.go.standard/project-web/standard/crud.user.singleton/model/user"
	repo "github.com/jeffotoni/project.go.standard/project-web/standard/crud.user.singleton/repo/user"
)

//UserPost ..
func UserPost(w http.ResponseWriter, r *http.Request) {
	if strings.ToUpper(r.Method) == http.MethodPost {
		handleruserCreate(w, r)
	} else {
		w.Header().Add("Content-Type", "application/json")
		jsonstr := `{"msg":"O método permitido é POST!"}`
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(jsonstr))
		return
	}
}

//UserCreate ..
func handleruserCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var p mUser.User
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		//bad request
		w.WriteHeader(400)
		w.Write([]byte(fmts.Concat(`{"erro":"`, err.Error(), `"}`)))
		return
	}
	//checando se veio o nome
	if len(p.Name) < 0 {
		w.WriteHeader(400)
		w.Write([]byte(`{"erro":"Campo Nome obrigatorio!"}`))
		return
	}

	if err := repo.InsertUser(p); err != nil {
		//erro interno
		w.WriteHeader(500)
		w.Write([]byte(fmts.Concat(`{"erro":"`, err.Error(), `"}`)))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"msg":"Cadastrado com sucesso!"}`))
	return
}
