package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"

	"personal-projects/iris-01/web/controllers"
)

func main() {
	app := iris.New()

	app.Logger().SetLevel("debug")

	// Load the template files.
	app.RegisterView(iris.HTML("./web/views", ".html"))

	// Serve our controllers.
	mvc.New(app.Party("/home")).Handle(new(controllers.HomeController))

	app.StaticWeb("/public", "./web/wwwroot")

	app.Run(
		// Start the web server at localhost:8080
		iris.Addr("localhost:8080"),
		// disables updates:
		iris.WithoutVersionChecker,
		// skip err server closed when CTRL/CMD+C pressed:
		iris.WithoutServerError(iris.ErrServerClosed),
		// enables faster json serialization and more:
		iris.WithOptimizations,
	)
}
