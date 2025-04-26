package models

type GetPodsReq struct {
	Namespace string `json:"namespace"`
}

type Deployment struct {
	Namespace  string     `json:"namespace"`
	ObjectMeta ObjectMeta `json:"objectMeta"`
	Spec       Spec       `json:"spec"`
	Template   Template   `json:"template"`
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
