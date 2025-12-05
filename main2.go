package main

import (
	"fmt"
	"log"
	"my-go-project/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main1() {
	r := gin.Default()
	// r.Use(gin.LoggerWithFormatter(logFormat()))
	r.Use(errorHandler())
	r.Use(logger())
	r.Static("/assets", "./assets")
	authorized := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		"foo":    "bar",
		"austin": "1234",
		"lena":   "hello",
	}))
	authorized.GET("/user", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)
		c.String(http.StatusOK, user)
	})
	r.GET("/handleerror", func(c *gin.Context) {
		somethingWrong := c.Query("wrong") == "true"
		if somethingWrong {
			c.Error(fmt.Errorf("something went wrong"))
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "all good!"})
	})
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
	r.GET("/albums", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, models.GetAllAlbums())
	})
	r.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, models.TestFunc())
	})
	r.GET("/albums/:id", models.GetAlbumByID)
	r.GET("/albums/search", models.GetAlbumByTitle)
	r.POST("/albums", models.PostAlbum)

	stc := r.Group("/struct")
	{
		stc.GET("/sb", models.GetDataB)
		stc.GET("/sc", models.GetDataC)
		stc.GET("/sd", models.GetDataD)
	}
	r.GET("/cookie", models.CookieExample)
	r.GET("/jsonp", func(c *gin.Context) {
		data := map[string]interface{}{
			"foo": "bar",
		}
		c.JSONP(http.StatusOK, data)
	})
	r.GET("/securejson", func(c *gin.Context) {
		names := []string{"nihao", "hello", "welldone"}
		c.SecureJSON(http.StatusOK, names)
	})
	r.GET("/basicauth", customBasicAuth(), func(c *gin.Context) {
		c.String(http.StatusOK, "login success")
	})

	r.Run("localhost:8080")
}

func main2() {
	router := gin.Default()
	http.ListenAndServe(":8080", router)
	// s := &http.Server{
	// 	Addr:           ":8080",
	// 	Handler:        router,
	// 	ReadTimeout:    10 * time.Second,
	// 	WriteTimeout:   10 * time.Second,
	// 	MaxHeaderBytes: 1 << 20,
	// }
	// s.ListenAndServe()
}

func logFormat() func(param gin.LogFormatterParams) string {
	return func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s\" \"%s\" %s\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}
}
func logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		c.Set("startTime:", t)
		c.Next()
		latency := time.Since(t)
		log.Println(latency)
		status := c.Writer.Status()
		log.Println(status)
	}
}
func errorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err
			c.JSON(http.StatusInternalServerError, map[string]string{
				"success": "false",
				"message": err.Error(),
			})
		}
	}
}
func customBasicAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, pass, hasAuth := c.Request.BasicAuth()
		if !hasAuth {
			c.AbortWithStatusJSON(401, gin.H{"error": "require authrization"})
			return
		}
		if user != "admin" || pass != "password" {
			c.AbortWithStatusJSON(401, gin.H{"error": "user or password wrong"})
			return
		}
		c.Set(gin.AuthUserKey, user)
		c.Next()
	}
}
