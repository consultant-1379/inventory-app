package models

type Menu struct {
	Instance []MenuInstance `json:"instances,"`
}

type MenuInstance struct {
	Name     string        `json:"name,omitempty"`
	Vpods    []MenuVpod    `json:"vpods,omitempty"`
	Servers  []MenuServer  `json:"servers,omitempty"`
	Clusters []MenuCluster `json:"clusters,omitempty"`
}

type MenuCluster struct {
	Name        string           `json:"name,omitempty" `
	Deployments []MenuDeployment `json:"deployments,omitempty" `
}

type MenuServer struct {
	Name        string           `json:"name,omitempty" `
	Deployments []MenuDeployment `json:"deployments,omitempty" `
}

type MenuVpod struct {
	Name        string           `json:"name,omitempty"`
	Deployments []MenuDeployment `json:"deployments,omitempty" `
	Clusters    []MenuCluster    `json:"clusters,omitempty" `
}

type MenuDeployment struct {
	Name string `json:"name,omitempty"`
}
