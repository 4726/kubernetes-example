package app

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/4726/kubernetes-example/config"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

type App struct {
	db     *redis.Client
	engine *gin.Engine
	conf   config.Config
	srv    *http.Server
}

//New returns a new App
func New(conf config.Config) *App {
	db := redis.NewClient(&redis.Options{
		Addr:     conf.DB.Addr,
		Password: conf.DB.Password,
		DB:       conf.DB.DB,
	})
	log.Println("using redis config: ")
	log.Println("Addr: ", conf.DB.Addr)
	log.Println("Password: ", conf.DB.Password)
	log.Println("DB: ", conf.DB.DB)

	r := gin.Default()
	app := &App{
		db:     db,
		engine: r,
		conf:   conf,
	}
	app.initRoutes()
	return app
}

func (a *App) initRoutes() {
	a.engine.GET("/kv", func(c *gin.Context) {
		getKV(c, a.db)
	})

	a.engine.POST("/kv", func(c *gin.Context) {
		setKV(c, a.db)
	})

	a.engine.DELETE("/kv", func(c *gin.Context) {
		deleteKV(c, a.db)
	})
}

//Run runs the app and blocks until an error occurs
func (a *App) Run() {
	a.srv = &http.Server{
		Addr:    a.conf.Addr,
		Handler: a.engine,
	}
	log.Println("starting server on addr: ", a.conf.Addr)
	if err := a.srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

//Close gracefully shuts down the server
func (a *App) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	a.srv.Shutdown(ctx)
	return a.db.Close()
}
