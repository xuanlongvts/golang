package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var secrects = gin.H{
	"foo": gin.H{"email": "foo@bar.com", "phone": "123445"},
	"austin": gin.H{"email": "austin@gmail.com", "phone": "56789"},
	"lena": gin.H{"email": "lena", "phone": "8727273"},
}

func main() {
	r := gin.Default()
	// Group using gin.BasicAuth() middleware
	// gin.Accounts is a shortcut for map[string]string
	authorized := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		"foo":    "bar",
		"austin": "1234",
		"lena":   "hello2",
		"manu":   "4321",
	}))

	// /admin/secrets endpoint
	// hit "localhost:8080/admin/secrets
	authorized.GET("/secrets", func(c *gin.Context) {
		// get user, it was set by the BasicAuth middleware
		user := c.MustGet(gin.AuthUserKey).(string)
		if secrect, ok := secrects[user]; ok {
			c.JSON(http.StatusOK, gin.H{
				"user": user,
				"secrect": secrect,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"user": user,
				"secret": "NO SECRET :(",
			})
		}
	})

	r.Run()
}

/*
Note: Run by Incognito Chrome browswer
URL: http://localhost:8080/admin/secrets

username / password

"foo":    "bar",
"austin": "1234",
"lena":   "hello2",
"manu":   "4321",
*/