package route

import (
	"github.com/gocraft/web"

	"github.com/hello-world/app/context"
	ctrl "github.com/hello-world/app/controllers"
)

func Router() (root *web.Router) {
	root = web.New(context.Context{})
	root.Middleware(web.ShowErrorsMiddleware)

	user := root.Subrouter(context.Context{}, "/user")
	user.Get("/ping", ctrl.Ping)

	return
}
