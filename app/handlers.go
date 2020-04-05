package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

type Response struct {
	Key, Value, Error string
	Success           bool
}

func GetKV(c *gin.Context, db *redis.Client) {
	req := struct {
		Key string
	}{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Error: "invalid request",
		})
	}
	var value string
	if err := db.Get(req.Key).Scan(&value); err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Error: "internal server error",
		})
	}
	c.JSON(http.StatusOK, Response{
		Key:     req.Key,
		Value:   value,
		Success: true,
	})
}

func SetKV(c *gin.Context, db *redis.Client) {
	req := struct {
		Key, Value string
	}{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Error: "invalid request",
		})
	}
	if err := db.Set(req.Key, req.Value, 0).Err(); err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Error: "internal server error",
		})
	}
	c.JSON(http.StatusOK, Response{
		Key:     req.Key,
		Value:   req.Value,
		Success: true,
	})
}

func DeleteKV(c *gin.Context, db *redis.Client) {
	req := struct {
		Key string
	}{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Error: "invalid request",
		})
	}
	if err := db.Del(req.Key).Err(); err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Error: "internal server error",
		})
	}
	c.JSON(http.StatusOK, Response{
		Key:     req.Key,
		Success: true,
	})
}
