package prices

import (
	sharedCommon "github.com/erply/api-go-wrapper/pkg/api/common"
)

type PriceListRule struct {
	ProductID int     `json:"productID"`
	Price     float32 `json:"price,string"`
	Amount    int     `json:"amount"`
}

type PriceList struct {
	ID                     int                         `json:"supplierPriceListID"`
	SupplierID             int                         `json:"supplierID"`
	SupplierName           string                      `json:"supplierName"`
	Name                   string                      `json:"name"`
	ValidFrom              string                      `json:"startDate"`
	ValidTo                string                      `json:"endDate"`
	Active                 string                      `json:"active"`
	AddedTimestamp         int                         `json:"added"`
	LastModifiedTimestamp  int                         `json:"lastModified"`
	AddedByUserName        string                      `json:"addedByUserName"`
	LastModifiedByUserName string                      `json:"lastModifiedByUserName"`
	Rules                  []PriceListRule             `json:"pricelistRules"`
	Attributes             []sharedCommon.ObjAttribute `json:"attributes"`
}

type ProductPriceList struct {
	PriceID              int     `json:"supplierPriceListProductID"`
	ProductID            int     `json:"productID"`
	Price                float32 `json:"price,string"`
	Amount               int     `json:"amount"`
	CountryID            int     `json:"countryID"`
	ProductSupplierCode  string  `json:"supplierCode"`
	ImportCode           string  `json:"importCode"`
	MasterPackQuantity   int     `json:"masterPackQuantity"`
	MinimumOrderQuantity int     `json:"minimumOrderQuantity"`
}

type GetPriceListsResponseBulkItem struct {
	Status     sharedCommon.StatusBulk `json:"status"`
	PriceLists []PriceList             `json:"records"`
}

type GetPriceListsResponseBulk struct {
	Status    sharedCommon.Status             `json:"status"`
	BulkItems []GetPriceListsResponseBulkItem `json:"requests"`
}

type GetPriceListsResponse struct {
	Status     sharedCommon.Status `json:"status"`
	PriceLists []PriceList         `json:"records"`
}

type GetProductPriceListResponseBulkItem struct {
	Status           sharedCommon.StatusBulk `json:"status"`
	ProductPriceList []ProductPriceList      `json:"records"`
}

type GetProductPriceListResponseBulk struct {
	Status    sharedCommon.Status                   `json:"status"`
	BulkItems []GetProductPriceListResponseBulkItem `json:"requests"`
}

type GetProductPriceListResponse struct {
	Status            sharedCommon.Status `json:"status"`
	ProductPriceLists []ProductPriceList  `json:"records"`
}

type ChangeProductToSupplierPriceListResult struct {
	ProductID int `json:"supplierPriceListProductID"`
}

type ChangeProductToSupplierPriceListResponse struct {
	Status                                 sharedCommon.Status                      `json:"status"`
	ChangeProductToSupplierPriceListResult []ChangeProductToSupplierPriceListResult `json:"records"`
}

type ChangeProductToSupplierPriceListResultBulkItem struct {
	Status  sharedCommon.StatusBulk                  `json:"status"`
	Records []ChangeProductToSupplierPriceListResult `json:"records"`
}

type ChangeProductToSupplierPriceListResponseBulk struct {
	Status    sharedCommon.Status                              `json:"status"`
	BulkItems []ChangeProductToSupplierPriceListResultBulkItem `json:"requests"`
}
