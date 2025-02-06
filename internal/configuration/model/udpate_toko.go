package model

type UpdateTokoRequest struct {
	Tokopedia string `json:"tokopedia"`
	Shopee    string `json:"shopee"`
	Blibli    string `json:"blibli"`
	Lazada    string `json:"lazada"`
	Bukalapak string `json:"bukalapak"`
}
