package getters

import (
	"bytes"
	"io"
	"net/http"
	"strings"

	"gerrit.ericsson.se/a/DETES/com.ericsson.de.stsoss/inventory-app/internal/helpers"
)

func GetInstanceFromHydra(instance string, token string) string {

	hydraUrl := "https://hydra.gic.ericsson.se/api/8.0/hql/instance"
	jsonString := "{\"query\":\"name='" + instance + "'\"}"
	//fmt.Println(jsonString)
	var jsonData = []byte(jsonString)

	request, err := http.NewRequest("POST", hydraUrl, bytes.NewBuffer(jsonData))
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	request.Header.Set("Authorization", strings.TrimSpace(token))

	helpers.CheckError(err)

	client := &http.Client{}
	response, err := client.Do(request)
	helpers.CheckError(err)
	defer response.Body.Close()

	// fmt.Println("response Status:", response.Status)
	// fmt.Println("response Headers:", response.Header)
	body, _ := io.ReadAll(response.Body)
	// fmt.Println("response Body:", string(body))

	return string(body)
}

func GetDeploymentFromDTT(deployment string) string {

	dttUrl := "https://atvdtt.athtem.eei.ericsson.se/api/deployments"
	query := "?q=name=" + deployment

	var reqUrl = dttUrl + query

	req, err := http.NewRequest("GET", reqUrl, nil)
	helpers.CheckError(err)

	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	helpers.CheckError(err)
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	helpers.CheckError(err)
	//fmt.Println("BODY: " + string(body))
	return string(body)
}

func GetBookingFromDTT(deployment_id string) string {

	dttUrl := "https://atvdtt.athtem.eei.ericsson.se/api/bookings"
	query := "?q=deployment_id=" + deployment_id

	var reqUrl = dttUrl + query

	req, err := http.NewRequest("GET", reqUrl, nil)
	helpers.CheckError(err)

	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	helpers.CheckError(err)
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	helpers.CheckError(err)
	//fmt.Println("BODY: " + string(body))
	return string(body)
}

func GetDeploymentFromDIT(deployment string) string {

	dttUrl := "https://atvdit.athtem.eei.ericsson.se/api/deployments"
	query := "?q=name=" + deployment

	var reqUrl = dttUrl + query

	req, err := http.NewRequest("GET", reqUrl, nil)
	helpers.CheckError(err)

	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	helpers.CheckError(err)
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	helpers.CheckError(err)
	//fmt.Println("BODY: " + string(body))
	return string(body)
}

func GetDocumentFromDIT(document_id string) string {

	dttUrl := "https://atvdit.athtem.eei.ericsson.se/api/documents"
	query := "/" + document_id

	var reqUrl = dttUrl + query

	req, err := http.NewRequest("GET", reqUrl, nil)
	helpers.CheckError(err)

	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	helpers.CheckError(err)
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	helpers.CheckError(err)
	//fmt.Println("BODY: " + string(body))
	return string(body)
}

func GetInstanceFromMonitoring() string {

	return "from Monitoring"
}
