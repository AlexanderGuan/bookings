package main

import (
	"log"
	"net/http"
	"time"

	"github.com/AlexanderGuan/bookings.git/pkg/config"
	"github.com/AlexanderGuan/bookings.git/pkg/handlers"
	"github.com/AlexanderGuan/bookings.git/pkg/render"
	"github.com/alexedwards/scs/v2"
)

var app config.AppConfig
var session *scs.SessionManager

func main() {

	// change this to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = time.Hour * 24
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("can not create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	// http.HandleFunc("/", repo.Home)
	// http.HandleFunc("/about", repo.About)

	// err = http.ListenAndServe("0.0.0.0:8080", nil)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	portNumber := "0.0.0.0:8080"
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
