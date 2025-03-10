package handlers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (h *Handler) idGet(c *gin.Context) {
	token, err := c.Cookie("token")
	if err != nil {
		log.Println("error from token: ", err)
		c.HTML(http.StatusUnauthorized, "index.html", gin.H{
			"IsAuthorized": false,
		})
		return
	}
	if _, err := h.Serv.ParsingJWTToken(token); err != nil {
		log.Println("error from Parsing token: ", err)
		c.HTML(http.StatusUnauthorized, "index.html", gin.H{
			"IsAuthorized": false,
		})
		return
	}
	c.HTML(http.StatusOK, "index.html", gin.H{
		"IsAuthorized": true,
	})
}
