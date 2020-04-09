package app

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

// Response is the http response for requests
type Response struct {
	Key, Value, Error string
	Success           bool
}

func getKV(c *gin.Context, db *redis.Client) {
	req := struct {
		Key string
	}{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Error: "invalid request",
		})
		log.Println("json error: ", err)
		return
	}
	var value string
	if err := db.Get(req.Key).Scan(&value); err != nil {
		if err != redis.Nil {
			c.JSON(http.StatusInternalServerError, Response{
				Error: "internal server error",
			})
			log.Println("internal server error: ", err)
			return
		}
	}
	c.JSON(http.StatusOK, Response{
		Key:     req.Key,
		Value:   value,
		Success: true,
	})
}

func setKV(c *gin.Context, db *redis.Client) {
	req := struct {
		Key, Value string
	}{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Error: "invalid request",
		})
		log.Println("json error: ", err)
		return
	}
	if err := db.Set(req.Key, req.Value, 0).Err(); err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Error: "internal server error",
		})
		log.Println("internal server error: ", err)
		return
	}
	c.JSON(http.StatusOK, Response{
		Key:     req.Key,
		Value:   req.Value,
		Success: true,
	})
}

func deleteKV(c *gin.Context, db *redis.Client) {
	req := struct {
		Key string
	}{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Error: "invalid request",
		})
		log.Println("json error: ", err)
		return
	}
	if err := db.Del(req.Key).Err(); err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Error: "internal server error",
		})
		log.Println("internal server error: ", err)
		return
	}
	c.JSON(http.StatusOK, Response{
		Key:     req.Key,
		Success: true,
	})
}
