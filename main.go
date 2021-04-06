package main

import (
	"github.com/r4lly99/echo-snack/db"
	"github.com/r4lly99/echo-snack/handler"
	"github.com/r4lly99/echo-snack/router"
	"github.com/r4lly99/echo-snack/service"
)

func main() {

	r := router.New()

	v1 := r.Group("/api")

	d := db.New()
	db.AutoMigrate(d)

	us := service.NewUserService(d)

	h := handler.NewHandler(us)
	h.Register(v1)
	r.Logger.Fatal(r.Start("127.0.0.1:8080"))

}
