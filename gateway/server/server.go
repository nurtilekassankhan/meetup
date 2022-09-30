package server

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
	"os/signal"
)

type Server interface {
	Run() error
	Stop() error
	RunGracefulShutdown() error
}

type server struct {
	port    string
	app     *fiber.App
	handler *handlr
}

func NewServer(port string, app *fiber.App, handler *handlr) (Server, error) {
	fiberServer := server{
		app:     app,
		handler: handler,
		port:    port,
	}
	fiberServer.Routes()
	return &fiberServer, nil
}

func (s *server) Run() error {
	return s.app.Listen(s.port)
}

func (s *server) Stop() error {
	return s.app.Shutdown()
}

func (s *server) RunGracefulShutdown() error {
	go func() {
		// if somebody terminate
		sigInt := make(chan os.Signal, 1)
		signal.Notify(sigInt, os.Interrupt)
		<-sigInt

		if err := s.app.Shutdown(); err != nil {
			log.Println(err)
		}

		close(sigInt)
	}()

	return s.app.Listen(s.port)
}
