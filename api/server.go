package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/techschool/simplebank/db/sqlc"
)

type Server struct {
	store *db.Store
	router *gin.Engine
}
func errorResponse (err error) gin.H {
	return gin.H{"error":err.Error()}
}

func NewServer(store db.Store)*Server {
	server:=&Server{store:store}
	router:=gin.Default()
    router.POST("/accounts",server.createAccount)


	server.router=router
	return server
}