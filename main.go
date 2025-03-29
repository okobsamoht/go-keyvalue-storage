package main

import (
	_ "fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/syndtr/goleveldb/leveldb"
	"log"
)

var db *leveldb.DB

func main() {
	var err error
	db, err = leveldb.OpenFile("mydb", nil)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	defer db.Close()

	router := gin.Default()

	router.POST("/kv", createKeyValue)
	router.GET("/kv/:key", getKeyValue)
	router.DELETE("/kv/:key", deleteKeyValue)

	router.Run(":8080")
}

func createKeyValue(c *gin.Context) {
	var json struct {
		Key   string `json:"key" binding:"required"`
		Value string `json:"value" binding:"required"`
	}

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := db.Put([]byte(json.Key), []byte(json.Value), nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to store value"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func getKeyValue(c *gin.Context) {
	key := c.Param("key")

	value, err := db.Get([]byte(key), nil)
	if err != nil {
		if err == leveldb.ErrNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Key not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve value"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"key": key, "value": string(value)})
}

func deleteKeyValue(c *gin.Context) {
	key := c.Param("key")

	err := db.Delete([]byte(key), nil)
	if err != nil {
		if err == leveldb.ErrNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Key not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete value"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "deleted"})
}
