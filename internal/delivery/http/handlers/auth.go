package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/go_web/internal/models"
	"log"
	"net/http"
)

// SignInGet отправка шаблона для аутентификации, get запрос
func (h *Handler) SignInGet(c *gin.Context) {
	c.HTML(http.StatusOK, "signin.html", gin.H{
		"error": "",
	})
}

// SignInPost аутентификация пользователя и генерация jwt токена, post запрос
func (h *Handler) SignInPost(c *gin.Context) {
	token, err := h.serv.GenerateJWTToken(c.PostForm("username"), c.PostForm("password"))
	if err != nil {
		c.HTML(http.StatusOK, "signin.html", gin.H{"error": err.Error()})
		return
	}
	c.SetCookie(
		"token",
		token,
		24*3600,
		"/",
		"",
		true,
		true,
	)
	c.Redirect(http.StatusMovedPermanently, "/id")
}

// SignUpGet отправка шаблона для создания пользователя, get запрос
func (h *Handler) SignUpGet(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.html", gin.H{})
}

// SignUpPost парсинг запроса и создание пользователя, post запрос
func (h *Handler) SignUpPost(c *gin.Context) {
	var client models.User

	if err := c.Bind(&client); err != nil {
		c.HTML(http.StatusOK, "signup.html", gin.H{"error": err.Error()})
		log.Println("json:", client.Password, client.Name, client.Password)
		return
	}
	err := h.serv.CreateUser(client)
	if err != nil {
		c.HTML(http.StatusOK, "signup.html", gin.H{"error": err.Error()})
		return
	}
	c.Redirect(http.StatusMovedPermanently, "/id")
}

func (h *Handler) idGet(c *gin.Context) {
	c.HTML(http.StatusOK, "welcome.html", gin.H{
		"IsAuthorized": true,
	})
}

func (h *Handler) welcome(c *gin.Context) {
	c.HTML(http.StatusOK, "welcome.html", gin.H{
		"IsAuthorized": false,
	})
}
