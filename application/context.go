package application

import (
	"github.com/anditakaesar/uwa-back/log"
	"github.com/unrolled/render"
)

type Context struct {
	Log    log.Interface
	Render *render.Render
}
