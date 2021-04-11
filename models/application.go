package models

type Config struct {
	Base
	Key   string `json:"key"`
	Value string `json:"value"`
}
