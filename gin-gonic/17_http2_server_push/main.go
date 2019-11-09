package main

import (
	"html/template"
	"log"

	"github.com/gin-gonic/gin"
)

var html = template.Must(template.New("https").Parse(`
<html>
<head>
	<meta charset="utf-8" />
    <link rel="icon" href="/assets/favicon.ico" />
    <meta name="viewport" content="width=device-width, initial-scale=1" />
    <meta name="theme-color" content="#000000" />
	<script src="/assets/app.js"></script>
  	<link rel="stylesheet" type="text/css" href="/assets/style.css" />
	<title>Https Test</title>
</head>
	<body>
	  <h1 style="color:red;">Welcome, Ginner!</h1>
	</body>
</html>
`))

func main() {
	r := gin.Default()
	r.Static("/assets", "./assets")
	r.SetHTMLTemplate(html)

	r.GET("/", func(c *gin.Context) {
		if pusher := c.Writer.Pusher(); pusher != nil {
			// use pusher.Push() to do server push
			if err := pusher.Push("/assets/app.js", nil); err != nil {
				log.Printf("Failed to push: %v", err)
			}
			if err := pusher.Push("/assets/style.css", nil); err != nil {
				log.Printf("Failed to push: %v", err)
			}
		}
		c.HTML(200, "https", gin.H{
			"status":"Success",
		})
	})
	r.RunTLS(":8080", "./testdata/server.pem", "./testdata/server.key")
}

// Note: run at  https://localhost:8080   or https://127.0.0.1:8080
// https  not http