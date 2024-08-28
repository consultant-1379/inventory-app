package models

import (
	"html/template"

	"gerrit.ericsson.se/a/DETES/com.ericsson.de.stsoss/inventory-app/internal/forms"
)

type TemplateData struct {
	StringMap map[string]string
	HtmlMap   map[string]template.HTML
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
	Form      *forms.Form
	//MenuItemsStruct JsonInstances
	MenuItemsStruct Menu
}
