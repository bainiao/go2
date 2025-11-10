package models

import (
	"github.com/gin-gonic/gin"
)

func CookieExample(c *gin.Context) {
	cookie, err := c.Cookie("gin_cookie")
	if err != nil {
		cookie = "No set"
		// cookie key, cookie value, max age(seconds), path, domain, secure - if true then https only, http only and forbidden javascript access
		c.SetCookie("gin_cookie", "here is cookie", 3600, "/", "localhost", false, true)
	}
	c.String(200, cookie)
}
