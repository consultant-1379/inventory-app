package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
)

// Holds the application conifg
type AppConfig struct {
	TemplateCache  map[string]*template.Template
	UseChache      bool
	InfoLog        *log.Logger
	InProduction   bool
	Session        *scs.SessionManager
	MenuJson       string
	HydraTokenPath string
	HydraToken     string
	MongoHost      string
	MongoPort      string
	MongoUser      string
	MongoPassword  string
	InitDB         bool
	DBName         string
}
