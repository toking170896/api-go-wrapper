package products

import (
	"context"
	"encoding/json"
	"github.com/erply/api-go-wrapper/pkg/common"
	erro "github.com/erply/api-go-wrapper/pkg/errors"
	"net/http"
	"strconv"
)

type (
	GetProductsResponse struct {
		Status   common.Status `json:"status"`
		Products []Product     `json:"records"`
	}

	getProductCategoriesResponse struct {
		Status            common.Status     `json:"status"`
		ProductCategories []ProductCategory `json:"records"`
	}

	getProductBrandsResponse struct {
		Status        common.Status  `json:"status"`
		ProductBrands []ProductBrand `json:"records"`
	}

	getProductGroupsResponse struct {
		Status        common.Status  `json:"status"`
		ProductGroups []ProductGroup `json:"records"`
	}

	ProductDimensions struct {
		Name             string `json:"name"`
		Value            string `json:"value"`
		Order            int    `json:"order"`
		DimensionID      int    `json:"dimensionID"`
		DimensionValueID int    `json:"dimensionValueID"`
	}

	ProductVariaton struct {
		ProductID  string              `json:"productID"`
		Name       string              `json:"name"`
		Code       string              `json:"code"`
		Code2      string              `json:"code2"`
		Dimensions []ProductDimensions `json:"dimensions"`
	}

	//Product ...
	Product struct {
		ProductID          int                `json:"productID"`
		ParentProductID    int                `json:"parentProductID"`
		Type               string             `json:"type"`
		Name               string             `json:"name"`
		Description        string             `json:"description"`
		DescriptionLong    string             `json:"longdesc"`
		Status             string             `json:"status"`
		Code               string             `json:"code"`
		Code2              string             `json:"code2"`
		Code3              *string            `json:"code3"`
		Price              float64            `json:"price"`
		PriceWithVat       float32            `json:"priceWithVat"`
		UnitName           *string            `json:"unitName"`
		Images             []ProductImage     `json:"images"`
		DisplayedInWebshop byte               `json:"displayedInWebshop"`
		CategoryId         uint               `json:"categoryID"`
		CategoryName       string             `json:"categoryName"`
		BrandID            uint               `json:"brandID"`
		BrandName          string             `json:"brandName"`
		GroupID            uint               `json:"groupID"`
		GroupName          string             `json:"groupName"`
		Warehouses         map[uint]StockInfo `json:"warehouses"`
		RelatedProducts    []string           `json:"relatedProducts"`
		Vatrate            float64            `json:"vatrate"`
		ProductVariations  []string           `json:"productVariations"` // Variations of matrix product
		VariationList      []ProductVariaton  `json:"variationList"`
	}

	StockInfo struct {
		WarehouseID   uint    `json:"warehouseID"`
		Free          int     `json:"free"`
		OrderPending  int     `json:"orderPending"`
		ReorderPoint  int     `json:"reorderPoint"`
		RestockLevel  int     `json:"restockLevel"`
		FifoCost      float32 `json:"FIFOCost"`
		PurchasePrice float32 `json:"purchasePrice"`
	}

	ProductImage struct {
		PictureID       string  `json:"pictureID"`
		Name            string  `json:"name"`
		ThumbURL        string  `json:"thumbURL"`
		SmallURL        string  `json:"smallURL"`
		LargeURL        string  `json:"largeURL"`
		FullURL         string  `json:"fullURL"`
		External        byte    `json:"external"`
		HostingProvider string  `json:"hostingProvider"`
		Hash            *string `json:"hash"`
		Tenant          *string `json:"tenant"`
	}

	ProductCategory struct {
		ProductCategoryID   uint   `json:"productCategoryID"`
		ParentCategoryID    uint   `json:"parentCategoryID"`
		ProductCategoryName string `json:"productCategoryName"`
		Added               uint64 `json:"added"`
		LastModified        uint64 `json:"lastModified"`
	}

	ProductBrand struct {
		ID           uint   `json:"brandID"`
		Name         string `json:"name"`
		Added        uint64 `json:"added"`
		LastModified uint64 `json:"lastModified"`
	}

	ProductGroup struct {
		ID              uint           `json:"productGroupID"`
		Name            string         `json:"name"`
		ShowInWebshop   string         `json:"showInWebshop"`
		NonDiscountable byte           `json:"nonDiscountable"`
		PositionNo      int            `json:"positionNo"`
		ParentGroupID   string         `json:"parentGroupID"`
		Added           uint64         `json:"added"`
		LastModified    uint64         `json:"lastModified"`
		SubGroups       []ProductGroup `json:"subGroups"`
	}

	//GetProductUnitsResponse ...
	GetProductUnitsResponse struct {
		Status       common.Status `json:"status"`
		ProductUnits []ProductUnit `json:"records"`
	}

	//ProductUnit ...
	ProductUnit struct {
		UnitID string `json:"unitID"`
		Name   string `json:"name"`
	}

	ProductManager interface {
		GetProducts(ctx context.Context, filters map[string]string) ([]Product, error)
		GetProductUnits(ctx context.Context, filters map[string]string) ([]ProductUnit, error)
		GetProductCategories(ctx context.Context, filters map[string]string) ([]ProductCategory, error)
		GetProductBrands(ctx context.Context, filters map[string]string) ([]ProductBrand, error)
		GetProductGroups(ctx context.Context, filters map[string]string) ([]ProductGroup, error)
	}
	Client struct {
		*common.Client
	}
)

func NewClient(sk, cc, partnerKey string, httpCli *http.Client) *Client {

	cli := &Client{
		common.NewClient(sk, cc, partnerKey, httpCli),
	}
	return cli
}

func (cli *Client) GetProductUnits(ctx context.Context, filters map[string]string) ([]ProductUnit, error) {

	resp, err := cli.SendRequest(ctx, "getProductUnits", filters)
	if err != nil {
		return nil, erro.NewFromError("GetProductUnits request failed", err)
	}

	res := &GetProductUnitsResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erro.NewFromError("unmarshaling GetProductUnitsResponse failed", err)
	}

	if !common.IsJSONResponseOK(&res.Status) {
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}

	return res.ProductUnits, nil
}

func (cli *Client) GetProducts(ctx context.Context, filters map[string]string) ([]Product, error) {
	resp, err := cli.SendRequest(ctx, "getProducts", filters)
	if err != nil {
		return nil, err
	}
	var res GetProductsResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erro.NewFromError("failed to unmarshal GetProductsResponse", err)
	}
	if !common.IsJSONResponseOK(&res.Status) {
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}
	return res.Products, nil
}

func (cli *Client) GetProductCategories(ctx context.Context, filters map[string]string) ([]ProductCategory, error) {
	resp, err := cli.SendRequest(ctx, "getProductCategories", filters)
	if err != nil {
		return nil, err
	}
	var res getProductCategoriesResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erro.NewFromError("failed to unmarshal getProductCategoriesResponse", err)
	}
	if !common.IsJSONResponseOK(&res.Status) {
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}
	return res.ProductCategories, nil
}

func (cli *Client) GetProductBrands(ctx context.Context, filters map[string]string) ([]ProductBrand, error) {
	resp, err := cli.SendRequest(ctx, "getProductBrands", filters)
	if err != nil {
		return nil, err
	}
	var res getProductBrandsResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erro.NewFromError("failed to unmarshal getProductBrandsResponse", err)
	}
	if !common.IsJSONResponseOK(&res.Status) {
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}
	return res.ProductBrands, nil
}

func (cli *Client) GetProductGroups(ctx context.Context, filters map[string]string) ([]ProductGroup, error) {
	resp, err := cli.SendRequest(ctx, "getProductGroups", filters)
	if err != nil {
		return nil, err
	}
	var res getProductGroupsResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erro.NewFromError("failed to unmarshal getProductGroupsResponse", err)
	}
	if !common.IsJSONResponseOK(&res.Status) {
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}
	return res.ProductGroups, nil
}
