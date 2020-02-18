package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"eajardini/vue/vuejwt/back/controler/mensagem"
	seguranca "eajardini/vue/vuejwt/back/seguranca"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

/*
type login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

var identityKey = "id"

func helloHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	user, _ := c.Get(identityKey)
	c.JSON(200, gin.H{
		"userID":   claims[identityKey],
		"userName": user.(*User).UserName,
		"text":     "Hello World.",
	})
}

// User demo
type User struct {
	UserName  string
	FirstName string
	LastName  string
}
*/

// CORSMiddleware : zz
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Origin, Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		// c.Writer.Header().Set("Access-Control-Allow-Credentials", "false")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true") // original
		fmt.Println("Metodo:", c.Request.Method)
		c.Next()
	}
}

// https://stackoverflow.com/questions/59516369/angular-8-cors-when-connecting-to-go-gin

func aberto(c *gin.Context) {
	c.JSON(200, gin.H{"mensagem": "ok"})
}

func main() {
	var userseg seguranca.User

	userseg.FirstName = "mama"

	port := os.Getenv("PORT")

	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	// r.Use(CORSMiddleware())
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH"},
		AllowHeaders:     []string{"Access-Control-Allow-Origin, Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "*"
		},
		MaxAge: 12 * time.Hour,
	}))

	if port == "" {
		port = "8000"
	}

	autenticadorMiddleware := seguranca.RetornaMiddleware()

	r.NoRoute(autenticadorMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	r.POST("/login", autenticadorMiddleware.LoginHandler)

	auth := r.Group("/auth")
	// Refresh time can be longer than token timeout
	auth.GET("/refresh_token", autenticadorMiddleware.RefreshHandler)
	auth.Use(autenticadorMiddleware.MiddlewareFunc())
	{
		auth.GET("/hello", seguranca.HelloHandler)
	}

	mens := r.Group("/mens")
	mens.Use(autenticadorMiddleware.MiddlewareFunc())
	{
		mens.GET("/ola", mensagem.Ola)
		mens.GET("/olasemclains", mensagem.Olasemclains)
		mens.POST("/olapost", mensagem.Olapost)
	}
	r.POST("/aberto", aberto)

	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
}
