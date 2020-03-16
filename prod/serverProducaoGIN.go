package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//RetornaOla : retorna
func RetornaOla(c *gin.Context) {
	c.JSON(200, gin.H{"mensagem": "Olá, Mundo!"})
}

func main() {
	port := "8100"
	r := gin.Default()

	r.Static("/css", "./dist/css")
	r.Static("/img", "./dist/img")
	r.Static("/js", "./dist/js")

	r.LoadHTMLFiles("dist/index.html")
	r.StaticFile("/favicon.ico", "./dist/favicon.ico")

	r.GET("/ola", RetornaOla)
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	})

	r.Run(":" + port)
	// if err := http.ListenAndServe(":"+port, r); err != nil {
	// 	log.Fatal(err)
	// }
}
