package handler

import (
	"net/http"
	"strings"

	repo "github.com/jeffotoni/project.go.standard/project-web/standard/postgres/crud-dao/repo/user"
)

//UserDelete ..
func UserDelete(id string, w http.ResponseWriter, r *http.Request) {
	if strings.ToUpper(r.Method) == http.MethodDelete {
		handlerDeleteUser(id, w, r)
	} else {
		w.Header().Add("Content-Type", "application/json")
		jsonstr := `{"msg":"O método permitido é Delete!"}`
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(jsonstr))
		return
	}
}

func handlerDeleteUser(id string, w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	contentType := strings.ToLower(r.Header.Get("Content-Type"))
	if contentType != "application/json" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		jsonstr := `{"msg":"Content-type é obrigatório!"}`
		w.Write([]byte(jsonstr))
		return
	}

	if err := repo.Delete(id); err != nil {
		w.WriteHeader(http.StatusOK)
		return
	}

	jsonstr := `{"msg":"Usuário não pode ser deletado!"}`
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(jsonstr))
	return
}
