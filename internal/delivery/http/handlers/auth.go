package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/go_web/internal/models"
	"log"
	"net/http"
)

func (h *Handler) SignInGet(c *gin.Context) {
	c.HTML(http.StatusOK, "signin.html", gin.H{
		"error": "",
	})
}

func (h *Handler) SignInPost(c *gin.Context) {
	_, err := h.serv.GenerateJWTToken(c.PostForm("username"), c.PostForm("password"))
	if err != nil {
		c.HTML(http.StatusOK, "signin.html", gin.H{"error": "неправильный пароль или логин"})
		return
	}
	c.Redirect(http.StatusMovedPermanently, "/id")
}

func (h *Handler) SignUpGet(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.html", gin.H{})
}

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
