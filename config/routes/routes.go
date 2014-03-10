package routes

import (
  "github.com/codegangsta/martini"
  "github.com/thresholderio/go-processing/handlers/root"
)

func Route(server *martini.ClassicMartini) {
  server.Get("/", root.Index)
}
