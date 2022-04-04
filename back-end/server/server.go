package server

import (
	"log"

	"github.com/NicolasSales0101/ultiVidros-project/back-end/routes"
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	port   string
	server *fiber.App
}

func NewServer() Server {
	fiber.New()
	return Server{
		port:   "5000",
		server: fiber.New(),
	}
}

func (s *Server) Run() {
	router := routes.ConfigRoutes(s.server)

	log.Println("Server is running at port:", s.port)
	log.Fatal(router.Listen(":" + s.port))
}
