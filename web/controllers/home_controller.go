package controllers

import (
	"github.com/kataras/iris/mvc"
)

// HomeController is the main controller
//
// Referene: https://github.com/joefitzgerald/gometalinter-linter/issues/34
type HomeController struct{}

var homeView = mvc.View{
	Name: "home/index.html",
	Data: map[string]interface{}{
		"Title":       "Welcome to Iris!",
		"MainMessage": "The fastest backend web framework for Go.",
	},
}

// Get will return a predefined view with bind data.
//
// `mvc.Result` is just an interface with a `Dispatch` function.
// `mvc.Response` and `mvc.View` are the built'n result type dispatchers
// you can even create custom response dispatchers by
// implementing the `github.com/kataras/iris/hero#Result` interface.
func (c *HomeController) Get() mvc.Result {
	return homeView
}
