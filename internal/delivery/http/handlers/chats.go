package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) idGet(c *gin.Context) {
	token, err := c.Cookie("token")
	if err != nil {
		c.HTML(http.StatusUnauthorized, "index.html", gin.H{
			"IsAuthorized": false,
		})
		return
	}
	if _, err := h.Serv.ParsingJWTToken(token); err != nil {
		c.HTML(http.StatusUnauthorized, "index.html", gin.H{
			"IsAuthorized": false,
		})
		return
	}
	c.HTML(http.StatusOK, "index.html", gin.H{
		"IsAuthorized": true,
	})
}
