package router

import (
	"backend-food/internal/pkg/handler"
	"backend-food/pkg/infrastucture/db"
	"fmt"

	"github.com/gin-gonic/gin"
)

type Router struct {
	Engine *gin.Engine
	DB     db.Database
}

func (r *Router) Setup() {
	var err error
	r.Engine = gin.Default()
	r.DB, err = db.NewDB()
	if err != nil {
		fmt.Println(err)
	}

	h := handler.NewHTTPHandler(r.DB)
	webAPI := r.Engine.Group("/app")
	{
		webAPI.POST("/query", h.Handle)
	}
}
func NewRouter() Router {
	var r Router
	r.Setup()
	return r
}
