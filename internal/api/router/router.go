package router

import (
	"backend-food/pkg/infrastucture/db"
	"backend-food/pkg/share/middleware"
	"backend-food/pkg/share/validators"
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
	r.DB.MigrateDBWithGorm()
	validators.SetUpValidator()
	// h := handler.NewHTTPHandler(r.DB)

	webAPI := r.Engine.Group("/app")
	{

		musicAPI := webAPI.Group("/account")
		{
			musicAPI.Use(middleware.AuthMiddleware(r.DB))
			{
			}
		}

	}
}
func NewRouter() Router {
	var r Router
	r.Setup()
	return r
}
