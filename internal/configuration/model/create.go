package model

type ConfigurationCreate struct {
	Key   string `json:"key"`
	Type  string `json:"type"`
	Value string `json:"value"`
}
