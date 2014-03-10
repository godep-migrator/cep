package root

import (
	"github.com/codegangsta/martini"
	"github.com/martini-contrib/render"
	"log"
	"net/http"
)

func Index(req *http.Request, res http.ResponseWriter, params martini.Params, logger *log.Logger, r render.Render) {
  r.JSON(200, map[string]bool{ "root": true })
}
