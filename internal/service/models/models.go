package models

import "fmt"

type Create struct {
	Type        Type   `json:"type"`
	Image       string `json:"image"`
	Replicas    int    `json:"replicas"`
	Description string `json:"description"`
}

func (c Create) String() string {
	str := fmt.Sprintf("Image: %s replicas: %d", c.Image, c.Replicas)

	if c.Description != "" {
		str += " " + c.Description
	}

	return str
}

type Delete struct {
}

type Update struct {
}

type List struct {
	Namespace string `json:"namespace"`
	Type      Type   `json:"type"`
}

const (
	TypeDeployment Type = iota
	TypePersistentVolumeClaim
	TypeService
)

const (
	TypeDB Type = iota
	TypeApp
)

type Type int

func (t Type) String() string {
	switch t {
	case 0:
		return "Deployment"
	default:
		return ""
	}
}

type Deployment struct {
	Namespace  string     `json:"namespace"`
	ObjectMeta ObjectMeta `json:"objectMeta"`
	Spec       Spec       `json:"spec"`
	Template   Template   `json:"template"`
}

type PersistentVolumeClaim struct {
}

type Service struct {
}

type ObjectMeta struct {
	Name string `json:"name"`
}

type Spec struct {
	Replicas int `json:"replicas"`
	Selector struct {
		MatchLabels map[string]string `json:"match_labels"`
	} `json:"selector"`
}

type Template struct {
	Spec struct {
		Containers []struct {
			Name  string `json:"name"`
			Image string `json:"image"`
		}
	}
}
