package main

import (
	"gee"
	"log"
	"net/http"
)

func main() {
	e := gee.New()
	e.GET("/", indexHandler)
	e.GET("/hello", helloHandler)
	e.POST("/login", loginHandler)
	if err := e.Run(":9999"); err != nil {
		log.Fatal(err)
	}
}

func indexHandler(c *gee.Context) {
	c.HTML(http.StatusOK, "<h1>Hello</h1>")
}

func helloHandler(c *gee.Context) {
	c.String(http.StatusOK, "hello %s, you're now at path [%s]\n", c.Query("name"), c.Path)
}

func loginHandler(c *gee.Context) {
	c.JSON(http.StatusOK, gee.H{
		"username": c.PostForm("username"),
		"password": c.PostForm("password"),
	})
}
