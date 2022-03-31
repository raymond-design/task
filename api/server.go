package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/raymond-design/task/api/controllers"
	"github.com/raymond-design/task/api/models"
	"github.com/raymond-design/task/api/services"
)

type Backend struct {
	server  *http.Server
	service *services.Reminders
}

// initializes and creates a new server
func New(addr string, service *services.Reminders) *Backend {
	cfg := controllers.RouterConfig{Service: service}
	router := controllers.NewRouter(cfg)
	return &Backend{
		server: &http.Server{
			Addr:    addr,
			Handler: router,
		},
		service: service,
	}
}

// starts the initialized
func (b *Backend) Start() error {
	log.Printf("application started on address %s\n", b.server.Addr)
	err := b.service.Populate()
	if err != nil {
		return models.WrapError("could not initialize reminders service", err)
	}

	err = b.server.ListenAndServe()
	if err == http.ErrServerClosed {
		log.Println("http server is closed")
		return nil
	}
	return err
}

// gracefully stops the server
func (b *Backend) Stop() error {
	timeout := 2 * time.Second
	done, err := make(chan struct{}), make(chan error)

	go func() {
		log.Println("shutting down the http server")
		if e := b.server.Shutdown(context.Background()); e != nil {
			err <- models.WrapError("error on server shutdown", e)
		}
		close(done)
	}()

	select {
	case <-done:
		log.Println("application was shut down")
		return nil
	case e := <-err:
		return e
	case <-time.After(timeout):
		return fmt.Errorf("shudown timeout of %v", timeout)
	}
}
