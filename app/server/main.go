package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	r.GET("/", Handler)
	r.GET("/readyz", HealthCheckHandler)
	r.Run(":8080")
}

func Handler(c *gin.Context) {
	randomWord := generateRandomWord(c)
	c.String(http.StatusOK, randomWord)
}

func generateRandomWord(c *gin.Context) string {

	words := []string{"apple", "banana", "cherry", "date", "elderberry"}

	seed := time.Now().UnixNano()
	rand.New(rand.NewSource(seed))

	word := words[rand.Intn(len(words))]

	fmt.Println(word)

	return word
}

func HealthCheckHandler(c *gin.Context) {
	c.String(http.StatusOK, "Healthy")
}
