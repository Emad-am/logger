package httpserver

import "github.com/gin-gonic/gin"

type httpServer struct {
	host string
	port string
	Ge   *gin.Engine
}

func NewHttpServer(host string, port string) *httpServer {
	return &httpServer{
		host,
		port,
		gin.Default(),
	}
}

