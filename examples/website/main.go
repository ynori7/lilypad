package main

import (
	"html/template"
	"net/http"
	"strings"
	"unicode"

	"github.com/ynori7/lilypad/errors"
	"github.com/ynori7/lilypad/handler"
	"github.com/ynori7/lilypad/log"
	"github.com/ynori7/lilypad/routing"
	"github.com/ynori7/lilypad/view"
)

/*
 * This is a simple website which demos most of the functionality available within this framework
 */

func main() {
	// Register the http routes
	routing.RegisterRoutes(routing.Route{
		Path:    "/hello/{name}",
		Handler: Hello,
	})

	// Add a simple template function to be used within templates
	view.RegisterGlobalTemplateFuncs(template.FuncMap{
		"UppercaseFirstLetter": func(s string) string {
			return strings.Title(strings.ToLower(s))
		},
	})

	// Indicate where the base HTML templates are located
	view.SetLayoutDirectory("examples/website/view/layout")

	// Error responses should be turned into HTML
	errors.UseMarkupErrorsWithLayout("layout", "examples/website/view/error.gohtml")

	// Start the server
	log.Info("Starting service")
	routing.ServeHttp(":8080")
}

// This is the handler which will receive requests
func Hello(r *http.Request) handler.Response {
	logger := log.WithRequest(r).WithFields(log.Fields{"logger": "Hello"})
	logger.Info("Handling request")

	// Validate the input
	name := routing.GetVar(r, "name")
	if !isValidName(name) {
		logger.Debug("Invalid name sent")
		return handler.ErrorResponse(errors.BadRequestError("Names should be non-empty and contain only letters"))
	}

	// Render the view
	resp, err := view.New("layout", "examples/website/view/hello.gohtml").Render(HelloTemplateData{Name: name})
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("Error rendering view")
		return handler.ErrorResponse(errors.InternalServerError("something went wrong"))
	}

	return handler.SuccessResponse(resp).WithMaxAge(300) //cache for 5 minutes
}

func isValidName(name string) bool {
	if len(name) == 0 {
		return false
	}
	for _, r := range name {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

type HelloTemplateData struct {
	Name string
}
