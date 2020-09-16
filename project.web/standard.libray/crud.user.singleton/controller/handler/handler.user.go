package handler

import (
	"net/http"
	"strings"

	"github.com/jeffotoni/project.go.standard/project.web/standard.libray/crud.user/api.user/repo/user"
)

func DeleteUser(userUuid string, w http.ResponseWriter, r *http.Request) {
	if strings.ToUpper(r.Method) == http.MethodDelete {
		handlerDeleteUser(userUuid, w, r)
	} else {
		jsonstr := `{"msg":"O método permitido é Delete!"}`
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(jsonstr))
		return
	}
}

func handlerDeleteUser(userUuid string, w http.ResponseWriter, r *http.Request) {
	content_type := strings.ToLower(r.Header.Get("Content-Type"))
	if content_type != "application/json" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		jsonstr := `{"msg":"Content-type é obrigatório!"}`
		w.Write([]byte(jsonstr))
		return
	}

	if user.Delete(userUuid) {
		w.WriteHeader(http.StatusOK)
		return
	}

	jsonstr := `{"msg":"Usuário não pode ser deletado!"}`
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(jsonstr))
	return
}
