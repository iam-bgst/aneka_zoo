package infrastructures

import (
	"fmt"
	"net/http"
	"time"
	"zoo/middleware"
	
	"github.com/gin-gonic/gin"
	"zoo/application"
	"zoo/libraries/ginResponse"
	"zoo/runner"
)

type (
	HTTPServer struct {
		Router *gin.Engine
	}
)

func NewHTTPServer(app *application.App) {
	httpServer := &HTTPServer{}
	
	httpServer.Router = gin.New()
	
	server := &http.Server{
		Addr:           fmt.Sprintf(":%s", app.Configurations.Const.HtppPort),
		Handler:        httpServer.Router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	
	// requestId Middleware
	requestIdMiddleware := middleware.NewRequestID()
	httpServer.Router.Use(requestIdMiddleware.Use())
	
	httpServer.setupRouter(app)
	
	runner.ServerRunner(server)
}

func (hs *HTTPServer) setupRouter(app *application.App) {
	dep := application.NewDependency(app.Configurations, app.Logger, app.Postgres)
	
	v1 := hs.Router.Group("/v1")
	{
		dep.Deliveries.Animal.SetupRoute(v1.Group("/animal"))
	}
	hs.Router.NoRoute(func(c *gin.Context) {
		ginResponse.SendResponseWithoutMeta(c, "PAGE_NOT_FOUND", gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"}, http.StatusNotFound)
	})
	
}
