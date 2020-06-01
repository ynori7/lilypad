package handler

import "net/http"

// Handler is a method for handling an http request
type Handler func(*http.Request) Response
