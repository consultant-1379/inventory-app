package main

import (
	"crypto/tls"
	"log"
	"net/http"
	"time"

	"gerrit.ericsson.se/a/DETES/com.ericsson.de.stsoss/inventory-app/internal/config"
	"gerrit.ericsson.se/a/DETES/com.ericsson.de.stsoss/inventory-app/internal/connection"
	"gerrit.ericsson.se/a/DETES/com.ericsson.de.stsoss/inventory-app/internal/getters"
	"gerrit.ericsson.se/a/DETES/com.ericsson.de.stsoss/inventory-app/internal/handlers"
	"gerrit.ericsson.se/a/DETES/com.ericsson.de.stsoss/inventory-app/internal/initialiser"
	"gerrit.ericsson.se/a/DETES/com.ericsson.de.stsoss/inventory-app/internal/render"
	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	app.InProduction = false
	app.InitDB = false
	app.MongoHost = "10.120.151.117"
	app.MongoPort = "8082"
	app.MongoUser = "root"
	app.MongoPassword = "example"
	app.DBName = "DETSIT"
	app.MenuJson = "menu.json"
	app.HydraTokenPath = "token"
	app.HydraToken = getters.GetHydraToken(app.HydraTokenPath)

	connection.ConnectDatabase("mongodb://" + app.MongoUser + ":" + app.MongoPassword + "@" + app.MongoHost + ":" + app.MongoPort)
	if app.InitDB {
		initialiser.InitDB(&app)
	}

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseChache = false
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	log.Println("Running on", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
