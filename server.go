package main

import (
  "github.com/codegangsta/martini"
	"github.com/martini-contrib/render"
  "github.com/thresholderio/go-processing/config/routes"
  "net/http"
)

func main() {
  server := martini.Classic()
  routes.Route(server)

  server.Use(render.Renderer(render.Options{ IndentJSON: true }))

  http.ListenAndServe(":8080", server)
  server.Run()
}
