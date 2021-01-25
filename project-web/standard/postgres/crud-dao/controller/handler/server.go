package handler

import (
	"database/sql"
	"net/http"
)

type Server struct {
	App *http.Request
	Db  *sql.DB
}
