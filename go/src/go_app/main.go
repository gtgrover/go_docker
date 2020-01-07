package main

import (
  "github.com/gin-gonic/gin"
  "net/http"
)

func getting(c *gin.Context ) {
   name := c.DefaultQuery("name", "Friend")
   action := c.DefaultQuery("action","waiting")
   message := "GET " + name + " is " + action

   c.JSON(200, gin.H{
   "status":  "OK",
   "message": message,
   })
//   c.String(http.StatusOK, message)
 }

 func posting(c *gin.Context ) {
    name := c.PostForm("name")
    action := c.PostForm("action")
    message := "POST " + name + " is " + action
    c.String(http.StatusOK, message)
  }

func putting(c *gin.Context ) {
     name := c.Param("name")
     action := c.Param("action")
     message := "PUT " + name + " is " + action
     c.String(http.StatusOK, message)
   }

 func deleting(c *gin.Context ) {
      name := c.Param("name")
      action := c.Param("action")
      message := "DELETE " + name + " is " + action
      c.String(http.StatusOK, message)
    }

func patching(c *gin.Context ) {
     name := c.Param("name")
     action := c.Param("action")
     message := "PATCH " + name + " is " + action
     c.String(http.StatusOK, message)
   }

func head(c *gin.Context ) {
    name := c.Param("name")
    action := c.Param("action")
    message := "HEAD " + name + " is " + action
    c.String(http.StatusOK, message)
  }

func options(c *gin.Context ) {
    name := c.Param("name")
    action := c.Param("action")
    message := "OPTIONS " + name + " is " + action
    c.String(http.StatusOK, message)
  }

func main() {
	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

  router.GET("/", getting)
  router.POST("/some", posting)
	router.PUT("/some", putting)
	router.DELETE("/some", deleting)
	router.PATCH("/some", patching)
	router.HEAD("/some", head)
	router.OPTIONS("/some", options)

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	router.Run()
	// router.Run(":3000") for a hard coded port
}
