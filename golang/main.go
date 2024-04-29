package main

import (
	_ "logger/config"
	_ "logger/internal/context"
	db "logger/internal/db"
	httpserver "logger/internal/http_server"
	_ "logger/internal/redis"
	"logger/src/routes"
	"logger/src/services"
)

func main() {
	hs := httpserver.NewHttpServer("localhost", "8080")
	r := hs.Ge
	routes.DefineMiddlewares(r)
	routes.DefineRoutes(r)
	services.StartLoggerService()
	defer func() {
		db.Disconnect()
	}()

	r.Run()
}
