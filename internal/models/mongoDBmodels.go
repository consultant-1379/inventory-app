package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Instances struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Instance []Instance         `json:"instances," bson:"instances,omitempty"`
}

type Instance struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Name     string             `bson:"name,omitempty"`
	Vpods    []Vpod             `bson:"vpods,omitempty"`
	Servers  []Server           `bson:"servers,omitempty"`
	Clusters []Cluster          `bson:"clusters,omitempty"`
	HwInfo   HwInfo             `bson:"hwInfo,omitempty"`
}

type Cluster struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name,omitempty"`
	Deployments []Deployment       `bson:"deployments,omitempty"`
	HwInfo      HwInfo             `bson:"hwInfo,omitempty"`
}

type Server struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name,omitempty"`
	Deployments []Deployment       `bson:"deployments,omitempty"`
	HwInfo      HwInfo             `bson:"hwInfo,omitempty"`
}

type Vpod struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name,omitempty"`
	Deployments []Deployment       `bson:"deployments,omitempty"`
	Clusters    []Cluster          `bson:"clusters,omitempty"`
	HwInfo      HwInfo             `bson:"hwInfo,omitempty"`
}

type Deployment struct {
	ID     primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name   string             `bson:"name,omitempty" json:"name,omitempty"`
	InUse  string             `bson:"inUse,omitempty" json:"inUse,omitempty"`
	HwInfo HwInfo             `bson:"hwInfo,omitempty"`
}

type HwInfo struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	cpu_total string             `bson:"cpu_total,omitempty" json:"cpu_total,omitempty"`
	ram_total string             `bson:"ram_total,omitempty" json:"ram_total,omitempty"`
	cpu_used  string             `bson:"cpu_used,omitempty" json:"cpu_used,omitempty"`
	ram_used  string             `bson:"ram_used,omitempty" json:"ram_used,omitempty"`
}
