package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/starfederation/datastar-go/datastar"

	"tailwindtest/gintemplrenderer"
)

func main() {
	r := gin.Default()
	ginHtmlRenderer := r.HTMLRender
	r.HTMLRender = &gintemplrenderer.HTMLTemplRenderer{FallbackHtmlRenderer: ginHtmlRenderer}

	// Serve static files
	r.Static("/assets", "./assets")

	// main entry point
	r.GET("/", func(c *gin.Context) {
		r := gintemplrenderer.New(c.Request.Context(), http.StatusOK, Home())
		c.Render(http.StatusOK, r)
	})
	r.POST("/save", func(c *gin.Context) {
		sse := datastar.NewSSE(c.Writer, c.Request)
		p := Form()

		if err := sse.PatchElementTempl(p); err != nil {
			panic(">" + err.Error() + "<")
			return
		}
	})

	log.Println("Server running on http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
