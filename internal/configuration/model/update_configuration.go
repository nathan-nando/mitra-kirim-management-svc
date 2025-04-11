package model

type UpdateConfigurationRequest struct {
	AppName        string `json:"appName,omitempty"`
	AppDescription string `json:"appDescription,omitempty"`
	Tokopedia      string `json:"tokopedia,omitempty"`
	Shopee         string `json:"shopee,omitempty"`
	Blibli         string `json:"blibli,omitempty"`
	Lazada         string `json:"lazada,omitempty"`
	Bukalapak      string `json:"bukalapak,omitempty"`
	Instagram      string `json:"instagram,omitempty"`
	Facebook       string `json:"facebook,omitempty"`
	Tiktok         string `json:"tiktok,omitempty"`
	Whatsapp       string `json:"whatsapp,omitempty"`
}
