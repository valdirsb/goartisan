package controllers

import "net/http"

type {{.Name}}Controller struct {}

func (c *{{.Name}}Controller) Handle(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Handling {{.Name}}..."))
}
