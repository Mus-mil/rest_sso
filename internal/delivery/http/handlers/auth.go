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
	var user models.UserSignIn

	if err := c.ShouldBindJSON(&user); err != nil {
		log.Println(err)
		log.Println(user)
		c.HTML(http.StatusOK, "signin.html", gin.H{"error": "вы ввели не все параметры"})
		return
	}

	token, err := h.Serv.GenerateJWTToken(user.Username, user.Password)
	if err != nil {
		c.HTML(http.StatusOK, "signin.html", gin.H{"error": "неправильный пароль или логин"})
		return
	}
	c.SetCookie("token", token, 24*3600, "/", "", true, true)

	id, err := h.Serv.GetID(user.Username, user.Password)
	if err != nil {
		c.HTML(http.StatusOK, "signin.html", gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "Успешный вход",
		"redirect": "http://localhost:8080/" + string(id), // Динамический маршрут
	})
}

// SignUpGet отправка шаблона для создания пользователя, get запрос
func (h *Handler) SignUpGet(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.html", gin.H{})
}

// SignUpPost парсинг запроса и создание пользователя, post запрос
func (h *Handler) SignUpPost(c *gin.Context) {
	var user models.User

	if err := c.Bind(&user); err != nil {
		log.Println("my_err:", err.Error())
		c.HTML(http.StatusOK, "signup.html", gin.H{"error": "пароль должен содержать минимум 8 символов"})
		return
	}
	err := h.Serv.CreateUser(user)
	if err != nil {
		c.HTML(http.StatusOK, "signup.html", gin.H{"error": "создайте другого пользователя"})
		return
	}

	id, err := h.Serv.GetID(user.Username, user.Password)
	c.Redirect(http.StatusMovedPermanently, "/"+id)
}

func (h *Handler) idGet(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"IsAuthorized": true,
	})
}

func (h *Handler) welcome(c *gin.Context) {
	c.HTML(http.StatusOK, "welcome.html", gin.H{
		"IsAuthorized": false,
	})
}
