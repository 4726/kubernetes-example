package app

import (
	"log"

	"github.com/4726/kubernetes-example/config"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

type App struct {
	db     *redis.Client
	engine *gin.Engine
	conf   config.Config
}

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
		GetKV(c, a.db)
	})

	a.engine.POST("/kv", func(c *gin.Context) {
		SetKV(c, a.db)
	})

	a.engine.DELETE("/kv", func(c *gin.Context) {
		DeleteKV(c, a.db)
	})
}

func (a *App) Run() {
	log.Println("starting server on addr: ", a.conf.Addr)
	if err := a.engine.Run(a.conf.Addr); err != nil {
		log.Fatal(err)
	}
}

func (a *App) Close() error {
	return a.db.Close()
}
