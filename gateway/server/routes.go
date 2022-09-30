package server

import "github.com/gofiber/fiber/v2"

func (s *server) Routes() {
	route := s.app.Group("/api/v1")
	user := route.Group("/user")
	s.userRoutes(user)
}

func (s *server) userRoutes(user fiber.Router) {
	user.Post("/withdraw/request", s.handler.RequestWithdraw)
	user.Post("/withdraw/confirm", s.handler.ConfirmWithdraw)
}
