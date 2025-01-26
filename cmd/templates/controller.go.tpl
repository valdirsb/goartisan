package controllers

import "net/http"

type {{.Entity}}Controller struct {}

func (c *{{.Entity}}Controller) Handle(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Handling {{.Entity}}..."))
}
