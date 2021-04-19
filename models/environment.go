package models

type Environment struct {
	Base
	Name         string        `json:"name"`
	Namespace    string        `json:"namespace"`
	Description  string        `json:"description"`
	Applications []Application `json:"applications"`
}
