# Lilypad [![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/ynori7/lilypad) [![Build Status](https://travis-ci.org/ynori7/lilypad.svg?branch=master)](https://travis-ci.com/github/ynori7/lilypad) [![Coverage Status](https://coveralls.io/repos/github/ynori7/lilypad/badge.svg?branch=master)](https://coveralls.io/github/ynori7/lilypad?branch=master) [![Go Report Card](https://goreportcard.com/badge/ynori7/lilypad)](https://goreportcard.com/report/github.com/ynori7/lilypad)
A lightweight web application framework in Go.

## Features
Here is a brief overview of the feature set provided by the Lilypad framework.
A detailed example can be found in [examples](/examples).

### Routing
The Lilypad framework provides a simple wrapper around the Gorilla mux router. This
allows you to use Lilypad's handlers to simplfy response writing and error handling.

The framework keeps a global router which allows you to register routes close to the
handlers which control them without having to pass a router around everywhere.

To add a route, simply register it like this:

```go
import "github.com/ynori7/lilypad/http"

...

http.RegisterRoutes(http.Route{
    Path:    "/hello/{name}",
    Handler: Hello,
})
```

RegisterRoutes is variardic, so it's possible to register multiple routes at once:

```go
http.RegisterRoutes([]http.Route{
    {
        Path:    "/hello/{name}",
        Handler: Hello,
    },
    {
        Path:    "/goodbye/{name}",
        Handler: Goodbye,
    },
}...)
```

### Handlers
Lilypad's handler definition makes it easier to build your http handlers. You no longer
need to care about writing responses for ever place where the method exists. You simply
return either a `SuccessResponse` or an `ErrorResponse`.

For example:

```go
import (
	"github.com/ynori7/lilypad/http"
)

func Hello(r http.Request) http.Response {
    name := http.GetVar(r, "name")

    if !isValidName(name) {
        return http.ErrorResponse(errors.BadRequestError("Names should be non-empty and contain only letters"))
    }

    ...

    return SuccessResponse(resp)
```

It's also possible to return a `RedirectResponse` to redirect to another page.

You can also specify the cache duration like this:

```go
return SuccessResponse(resp).WithMaxAge(300)
```

Or any other arbitrary headers using the `WithHeaders()` method.

### Errors
The errors package provides a definition for HttpErros as well as convenient wrappers
for getting some of the most common error types.

One key attribute of the errors package is the HttpError's `Write()` method. This
method will present the error in a useful format depending on the configuration. In your
main.go, you can set errors to be presented in a JSON format, HTML format, or plaintext format.

```go
errors.UseMarkupErrors(errorHtmlTemplate)
//or
errors.UseJsonErrors()
//or
errors.UsePlaintextErrors() //the default
```

### Logging
The log package wraps the Sirupsen Logrus library and adds a few useful additions such as
the `WithRequest` method which returns a logger with some fields populated from the HTTP
request such as the client's IP address.

### Templating
The view package provides a minimalistic method to `RenderTemplate` which accepts a template 
body as a string and the data used to render it.

Alternatively, you can use nested templates by registering your base layouts like this:

```go
view.SetLayoutDirectory("path/to/layouts")
```

This will configure the framework to load all files with the gohtml extension. Next you can 
render a view like this:

```go
out, err := view.New("layoutName", "templates/specificTemplate.gohtml").Render(myData)
```

Here is an example layout.gohtml:
```gotemplate
{{ define "layoutName" }}
<html>
<body>
{{ template "body" . }}
</body>
</html>
```

And an example specificTemplate.gohtml:
```gotemplate
{{ define "body" }}
<h1>My Page!</h1>
<article>blahblah</article>
{{ end }}
```

It also has a global register of template functions which can be easily registered like this:

```go
view.RegisterGlobalTemplateFuncs(template.FuncMap{
    "UppercaseFirstLetter": func(s string) string {
        return strings.Title(strings.ToLower(s))
    },
})
```

And from any template, the function can be used like this:

```gotemplate
{{ UppercaseFirstLetter .Str }}
```
