// Go Api server
// @jeffotoni

package controller

import (
	"net/http"
	"sync"
)

// HTMLServer represents
// the web service that
// serves up HTML
type GoServerHttp struct {
	server *http.Server
	wg     sync.WaitGroup
}
