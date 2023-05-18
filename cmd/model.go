package cmd

type GetProjectResponse struct {
	Data []ProjectData `json:"data"`
}

type ProjectData struct {
	Name  string      `json:"name"`
	Links ProjectLink `json:"links"`
}

type ProjectLink struct {
	Deployments string `json:"deployments"`
}

type GetDeploymentResponse struct {
	Data []DeploymentData `json:"data"`
}

type DeploymentData struct {
	Name        string            `json:"name"`
	NamespaceId string            `json:"namespaceId"`
	Actions     DeploymentActions `json:"actions"`
	Links       DeploymentLinks   `json:"links"`
}

type DeploymentActions struct {
	Pause    string `json:"pause"`
	Redeploy string `json:"redeploy"`
	Resume   string `json:"resume"`
	RollBack string `json:"rollBack"`
}

type DeploymentLinks struct {
	Update string `json:"update"`
}
