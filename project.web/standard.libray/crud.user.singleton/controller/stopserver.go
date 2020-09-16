// Go Api server
// @jeffotoni

package controller

import (
	"context"
	"time"

	"github.com/jeffotoni/project.go.standard/project.web/standard.libray/crud.user.singleton/pkg/util"
	"github.com/jeffotoni/project.go.standard/project.web/standard.libray/crud.user.singleton/pkg/zerolog"
)

// Stop turns off the HTML Server
func (GoServerHttp *GoServerHttp) StopServer() error {

	// Create a context to attempt a graceful 6 second shutdown.
	const timeout = 6 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// optimized version
	util.Printnew("\n\033[0;31mServer Goapiproduto:\033[0m \033[0;32mService stopped\033[0m\n")

	// Attempt the graceful shutdown by closing the listener
	// and completing all inflight requests
	if err := GoServerHttp.server.Shutdown(ctx); err != nil {
		zerolog.Error(
			"1.0.0",
			"stopserver.go",
			32,
			"api.crud.user.singleton.com.br",
			"GoServerHttp StopServer",
			err.Error())
		// Looks like we timed out on the graceful shutdown. Force close.
		if err := GoServerHttp.server.Close(); err != nil {
			zerolog.Error(
				"1.0.0",
				"stopserver.go",
				39,
				"api.crud.user.singleton.com.br",
				"GoServerHttp Close",
				err.Error())
			util.Printnew("\n\033[0;31mServer Goapiproduto:\033[0m \033[0;32mService Error Close()", err.Error(), "\033[0m\n")
			return err
		}
	}

	// Wait for the listener to report that it is closed.
	GoServerHttp.wg.Wait()
	util.Printnew("\n\033[0;31mServer Goapiproduto:\033[0m \033[0;32mService Stopped")
	return nil
}
