package models

type JsonInstances struct {
	Instance []Instance `json:"instances,"`
}

type JsonInstance struct {
	Name     string    `json:"name,omitempty"`
	Vpods    []Vpod    `json:"vpods,omitempty"`
	Servers  []Server  `json:"servers,omitempty"`
	Clusters []Cluster `json:"clusters,omitempty"`
}

type JsonCluster struct {
	Name        string       `json:"name,omitempty" `
	Deployments []Deployment `json:"deployments,omitempty" `
}

type JsonServer struct {
	Name        string       `json:"name,omitempty" `
	Deployments []Deployment `json:"deployments,omitempty" `
}

type JsonVpod struct {
	Name        string       `json:"name,omitempty"`
	Deployments []Deployment `json:"deployments,omitempty" `
	Clusters    []Cluster    `json:"clusters,omitempty" `
}

type JsonDeployment struct {
	Name string `json:"name,omitempty"`
}

type HydraInstanceResults struct {
	HydraInstanceResult []HydraInstanceResult `json:"result"`
}

type HydraInstanceResult struct {
	Id                         int    `json:"id"`
	Name                       string `json:"name"`
	Consumer_team_id           int    `json:"consumer_team_id"`
	Consumer_tenant_id         int    `json:"consumer_tenant_id"`
	Description                string `json:"description"`
	Service_id                 int    `json:"service_id"`
	Life_cycle_status_id       int    `json:"life_cycle_status_id"`
	Logical_site_id            int    `json:"logical_site_id"`
	Service_version_id         int    `json:"service_version_id"`
	Service_level_agreement_id int    `json:"service_level_agreement_id"`
}

type DTTDeploymentResult struct {
	Id     string `json:"_id"`
	Name   string `json:"name"`
	Status string `josn:"status"`
}

type DTTBookingResult struct {
	Id          string `json:"_id"`
	Name        string `json:"name"`
	JiraIssue   string `josn:"jiraIssue"`
	IsStarted   bool   `json:"isStarted"`
	IsExpired   bool   `json:"isExpired"`
	StartTime   string `josn:"startTime"`
	EndTime     string `josn:"EndTime"`
	TestingType string `json:"testingType"`
}

type DITDeploymentResult struct {
	Id        string     `json:"_id"`
	Name      string     `json:"name"`
	Documents []Document `josn:"documents"`
}

type Document struct {
	Schema          string `json:"schema"`
	Id              string `json:"document_id"`
	Schema_category string `json:"schema_category"`
}

type DITDocumentResult struct {
	Id      string `json:"_id"`
	Name    string `json:"name"`
	Content struct {
		Ram string `json:"ram"`
		Cpu string `json:"cpu"`
	}
}
