// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type App struct {
	Name        string                 `json:"name"`
	Namespace   string                 `json:"namespace"`
	Containers  []*Container           `json:"containers"`
	ExposePorts map[string]interface{} `json:"expose_ports"`
	Replicas    *int                   `json:"replicas"`
	State       *State                 `json:"state"`
	Status      map[string]interface{} `json:"status"`
}

type AppInput struct {
	Name        string                 `json:"name"`
	Namespace   string                 `json:"namespace"`
	Containers  []*ContainerInput      `json:"containers"`
	ExposePorts map[string]interface{} `json:"expose_ports"`
	Replicas    int                    `json:"replicas"`
	State       *StateInput            `json:"state"`
}

type Container struct {
	Name   string                 `json:"name"`
	Image  string                 `json:"image"`
	Env    map[string]interface{} `json:"env"`
	Ports  map[string]interface{} `json:"ports"`
	Memory *string                `json:"memory"`
}

type ContainerInput struct {
	Name   string                 `json:"name"`
	Image  string                 `json:"image"`
	Env    map[string]interface{} `json:"env"`
	Ports  map[string]interface{} `json:"ports"`
	Memory *string                `json:"memory"`
}

type State struct {
	Statefulset bool   `json:"statefulset"`
	StoragePath string `json:"storage_path"`
	StorageSize string `json:"storage_size"`
}

type StateInput struct {
	Statefulset bool   `json:"statefulset"`
	StoragePath string `json:"storage_path"`
	StorageSize string `json:"storage_size"`
}