package servicediscovery

import (
	common2 "github.com/erply/api-go-wrapper/pkg/api/common"
)

type getServiceEndpointsResponse struct {
	Status  common2.Status
	Records []ServiceEndpoints `json:"records"`
}
type Endpoint struct {
	IsSandbox     bool   `json:"isSandbox"`
	Url           string `json:"url"`
	Documentation string `json:"documentation"`
}

type ServiceEndpoints struct {
	Cafa         Endpoint `json:"cafa"`
	Pim          Endpoint `json:"pim"`
	Wms          Endpoint `json:"wms"`
	Promotion    Endpoint `json:"promotion"`
	Reports      Endpoint `json:"reports"`
	JSON         Endpoint `json:"json"`
	Assignments  Endpoint `json:"assignments"`
	AccountAdmin Endpoint `json:"account-admin"`
	VisitorQueue Endpoint `json:"visitor-queue"`
	Loyalty      Endpoint `json:"loyalty"`
	Cdn          Endpoint `json:"cdn"`
	Tasks        Endpoint `json:"tasks"`
	Webhook      Endpoint `json:"webhook"`
	User         Endpoint `json:"user"`
	Import       Endpoint `json:"import"`
	Ems          Endpoint `json:"ems"`
	Clockin      Endpoint `json:"clockin"`
	Ledger       Endpoint `json:"ledger"`
	Auth         Endpoint `json:"auth"`
	Crm          Endpoint `json:"crm"`
	Buum         Endpoint `json:"buum"`
	Sales        Endpoint `json:"sales"`
	Pricing      Endpoint `json:"pricing"`
	Inventory    Endpoint `json:"inventory"`
	Chair        Endpoint `json:"chair"`
	PosAPI       Endpoint `json:"pos-api"`
	Erply        Endpoint `json:"erply"`
}
