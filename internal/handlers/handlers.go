package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	"gerrit.ericsson.se/a/DETES/com.ericsson.de.stsoss/inventory-app/internal/config"
	"gerrit.ericsson.se/a/DETES/com.ericsson.de.stsoss/inventory-app/internal/getters"
	"gerrit.ericsson.se/a/DETES/com.ericsson.de.stsoss/inventory-app/internal/helpers"
	"gerrit.ericsson.se/a/DETES/com.ericsson.de.stsoss/inventory-app/internal/models"
	"gerrit.ericsson.se/a/DETES/com.ericsson.de.stsoss/inventory-app/internal/render"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the handler for the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, r, "home.page.tmpl.html", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	// send data to the template
	render.RenderTemplate(w, r, "about.page.tmpl.html", &models.TemplateData{
		StringMap: stringMap,
	})
}

func (m *Repository) Instance(w http.ResponseWriter, r *http.Request) {

	instance := strings.Replace(r.URL.Path, "/instance/", "", 1)

	var hydraResults models.HydraInstanceResults
	err := json.Unmarshal([]byte(getters.GetInstanceFromHydra(instance, m.App.HydraToken)), &hydraResults)
	helpers.CheckError(err)

	htmlMap := make(map[string]template.HTML)
	stringMap := make(map[string]string)
	stringMap["instance"] = instance
	stringMap["sub-title1"] = instance + " Summary"
	htmlMap["content1"] = template.HTML("name:" + hydraResults.HydraInstanceResult[0].Name + "<br>Id:" + fmt.Sprint(hydraResults.HydraInstanceResult[0].Id) + "<br>Description:" + hydraResults.HydraInstanceResult[0].Description)
	stringMap["sub-title2"] = "Vpods"
	htmlMap["content2"] = template.HTML("Some Vpod Info")
	stringMap["sub-title3"] = "Additional info"
	htmlMap["content3"] = template.HTML("<br>" + "<br>" + getters.GetInstanceFromMonitoring())

	render.RenderTemplate(w, r, "instance.page.tmpl.html", &models.TemplateData{
		StringMap: stringMap,
		HtmlMap:   htmlMap,
	})
}

func (m *Repository) Server(w http.ResponseWriter, r *http.Request) {

	server := strings.Replace(r.URL.Path, "/server/", "", 1)

	stringMap := make(map[string]string)
	htmlMap := make(map[string]template.HTML)
	stringMap["instance"] = server
	stringMap["sub-title1"] = server + " Summary"

	stringMap["sub-title2"] = "VMs"
	stringMap["sub-title3"] = "Additional info"

	render.RenderTemplate(w, r, "instance.page.tmpl.html", &models.TemplateData{
		StringMap: stringMap,
		HtmlMap:   htmlMap,
	})
}

func (m *Repository) Cluster(w http.ResponseWriter, r *http.Request) {
	cluster_name := strings.Replace(r.URL.Path, "/cluster/", "", 1)
	fmt.Println(cluster_name)
	cluster_bson := helpers.GetOneByName("Clusters", cluster_name, m.App)
	fmt.Println(cluster_bson)
	hwInfo_bson := helpers.GetByForeginKey("cluster_id", cluster_bson[0]["_id"].(primitive.ObjectID), "HwInfo", m.App)
	if len(hwInfo_bson) == 0 {
		hwInfo_bson = append(hwInfo_bson, primitive.M{})
	}
	fmt.Println(hwInfo_bson)
	deployments_bson := helpers.GetByForeginKey("cluster_id", cluster_bson[0]["_id"].(primitive.ObjectID), "Deployments", m.App)

	var deployments_html string
	for dep := range helpers.FieldFromSliceOfM(deployments_bson, "name") {
		deployments_html = helpers.FieldFromSliceOfM(deployments_bson, "name")[dep] + "<br>" + deployments_html
	}

	stringMap := make(map[string]string)
	htmlMap := make(map[string]template.HTML)
	stringMap["instance"] = cluster_name
	stringMap["sub-title1"] = cluster_name + " Summary"

	stringMap["sub-title2"] = "Deployments"
	htmlMap["content2"] = template.HTML(deployments_html)
	stringMap["sub-title3"] = "HW info"
	htmlMap["content3"] = template.HTML("cpu total: " + helpers.FieldFromM(hwInfo_bson[0], "cpu_total") + "</br>" +
		"cpu used: " + helpers.FieldFromM(hwInfo_bson[0], "cpu_used") + "</br>" +
		"ram total: " + helpers.FieldFromM(hwInfo_bson[0], "ram_total") + "</br>" +
		"ram used: " + helpers.FieldFromM(hwInfo_bson[0], "ram_used") + "</br>")

	render.RenderTemplate(w, r, "cluster.page.tmpl.html", &models.TemplateData{
		StringMap: stringMap,
		HtmlMap:   htmlMap,
	})
}

func (m *Repository) Deployment(w http.ResponseWriter, r *http.Request) {

	deployment := strings.Replace(r.URL.Path, "/deployment/", "", 1)
	dttResults := make([]models.DTTDeploymentResult, 0)
	err := json.Unmarshal([]byte(getters.GetDeploymentFromDTT(deployment)), &dttResults)
	helpers.CheckError(err)

	dttBookingResults := make([]models.DTTBookingResult, 0)
	err2 := json.Unmarshal([]byte(getters.GetBookingFromDTT(dttResults[0].Id)), &dttBookingResults)
	helpers.CheckError(err2)

	ditDeploymentResults := make([]models.DITDeploymentResult, 0)
	err3 := json.Unmarshal([]byte(getters.GetDeploymentFromDIT(deployment)), &ditDeploymentResults)
	helpers.CheckError(err3)

	// HERE Add mechanism to loop through docs
	var ditDocumentResults models.DITDocumentResult
	err4 := json.Unmarshal([]byte(getters.GetDocumentFromDIT(ditDeploymentResults[0].Documents[0].Id)), &ditDocumentResults)
	helpers.CheckError(err4)

	htmlMap := make(map[string]template.HTML)
	stringMap := make(map[string]string)
	stringMap["instance"] = deployment
	stringMap["sub-title1"] = deployment + " Summary"
	htmlMap["content1"] = template.HTML("name:" + dttResults[0].Name + "<br>Id:" + dttResults[0].Id + "<br>Status:" + dttResults[0].Status)
	stringMap["sub-title2"] = "Bookings"
	htmlMap["content2"] = template.HTML("name:" + dttBookingResults[0].Name + "<br>Id:" + dttBookingResults[0].Id + "<br>Start:" +
		dttBookingResults[0].StartTime +
		"<br>End:" + dttBookingResults[0].EndTime)
	stringMap["sub-title3"] = "HW info"
	htmlMap["content3"] = template.HTML("cpu:" + ditDocumentResults.Content.Cpu + "<br>ram:" + ditDocumentResults.Content.Ram)

	render.RenderTemplate(w, r, "deployment.page.tmpl.html", &models.TemplateData{
		StringMap: stringMap,
		HtmlMap:   htmlMap,
	})
}

func (m *Repository) Vpod(w http.ResponseWriter, r *http.Request) {

	instance := strings.Replace(r.URL.Path, "/vpod/", "", 1)
	var hydraResults models.HydraInstanceResults
	json.Unmarshal([]byte(getters.GetInstanceFromHydra(instance, m.App.HydraToken)), &hydraResults)
	htmlMap := make(map[string]template.HTML)
	stringMap := make(map[string]string)
	stringMap["instance"] = instance
	stringMap["sub-title1"] = instance + " Summary"
	htmlMap["content1"] = template.HTML("name:" + hydraResults.HydraInstanceResult[0].Name + "<br>Id:" + fmt.Sprint(hydraResults.HydraInstanceResult[0].Id) + "<br>Description:" + hydraResults.HydraInstanceResult[0].Description)
	stringMap["sub-title2"] = "Vpods"
	htmlMap["content2"] = template.HTML("Some Vpod Info")
	stringMap["sub-title3"] = "Additional info"
	htmlMap["content3"] = template.HTML("<br>" + "<br>" + getters.GetInstanceFromMonitoring())

	render.RenderTemplate(w, r, "vpod.page.tmpl.html", &models.TemplateData{
		StringMap: stringMap,
		HtmlMap:   htmlMap,
	})
}

func (m *Repository) PostInstance(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}

	instance := r.Form.Get("instance")
	vpod := r.Form.Get("vpod")
	server := r.Form.Get("server")
	log.Println("element :" + instance + vpod + server)
	stringMap := make(map[string]string)
	stringMap["instance"] = instance
	stringMap["vpod"] = vpod
	stringMap["server"] = server
	stringMap["title1"] = instance
	stringMap["title2"] = vpod
	stringMap["title3"] = server
	stringMap["sub-title1"] = instance
	stringMap["sub-title2"] = vpod
	stringMap["sub-title3"] = server
	stringMap["content1"] = "1"
	stringMap["content2"] = "2"
	stringMap["content3"] = "3"
	stringMap["page-title"] = "SYSTEMS"
	stringMap["page-sub-title"] = "INFO"
	m.App.Session.Put(r.Context(), "instance", instance)
	m.App.Session.Put(r.Context(), "vpod", server)
	m.App.Session.Put(r.Context(), "server", server)

	render.RenderTemplate(w, r, "instance.page.tmpl.html", &models.TemplateData{
		StringMap: stringMap,
	})
}
