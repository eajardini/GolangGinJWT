package mensagem

import (
	seguranca "eajardini/vue/vuejwt/back/seguranca"

	"github.com/gin-gonic/gin"
)

type login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

//Ola : zz
func Ola(c *gin.Context) {
	claims := seguranca.RetornaClaims(c)
	user, _ := c.Get(seguranca.IdentityKey)
	c.JSON(200, gin.H{
		"userID": claims[seguranca.IdentityKey],
		"text":   "Olá Usuário",
		"user":   user.(*seguranca.User).UserName,
	})
}

//Olasemclains : zz
func Olasemclains(c *gin.Context) {
	c.JSON(200, gin.H{
		"text": "Olá sem Clains",
	})
}

//Olapost : zz
func Olapost(c *gin.Context) {
	var log login
	c.ShouldBind(&log)

	c.JSON(200, gin.H{
		"nomeusuario": log.Username,
		"senha":       log.Password,
	})
}
