package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	mUser "github.com/jeffotoni/project.go.standard/project.web/standard.libray/crud.user.singleton/model/user"
	"github.com/jeffotoni/project.go.standard/project.web/standard.libray/crud.user.singleton/pkg/fmts"
	repo "github.com/jeffotoni/project.go.standard/project.web/standard.libray/crud.user.singleton/repo/user"
)

//UserPut ..
func UserPut(id string, w http.ResponseWriter, r *http.Request) {
	if strings.ToUpper(r.Method) == http.MethodPut {
		handlerupdateUser(id, w, r)
	} else {
		w.Header().Add("Content-Type", "application/json")
		jsonstr := `{"msg":"O método permitido é PUT!"}`
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(jsonstr))
		return
	}
}

//UpdateUser ..
func handlerupdateUser(id string, w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var p mUser.User
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		jsonmsg := "Error body!"
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(jsonmsg)
		return
	}
	//checando se veio o nome
	if len(p.Name) < 0 {
		w.WriteHeader(400)
		w.Write([]byte(`{"erro":"Campo Nome obrigatorio!"}`))
		return
	}

	if err := repo.UpdateUser(id, p); err != nil {
		w.WriteHeader(500)
		w.Write([]byte(fmts.Concat(`{"erro":"`, err.Error(), `"}`)))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"msg":"Dados atualizados com sucesso!"}`))
	return
}
