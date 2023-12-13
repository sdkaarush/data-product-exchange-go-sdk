/**
 * (C) Copyright IBM Corp. 2023.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

/*
 * IBM OpenAPI SDK Code Generator Version: 3.83.0-adaf0721-20231212-210453
 */

// Package dpxv1 : Operations and models for the DpxV1 service
package dpxv1

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"time"

	common "github.com/IBM/data-product-exchange-go-sdk/common"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/go-openapi/strfmt"
)

// DpxV1 : Data Product Exchange API Service
//
// API Version: 1.0.0
type DpxV1 struct {
	Service *core.BaseService
}

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "dpx"

// DpxV1Options : Service options
type DpxV1Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator
}

// NewDpxV1UsingExternalConfig : constructs an instance of DpxV1 with passed in options and external configuration.
func NewDpxV1UsingExternalConfig(options *DpxV1Options) (dpx *DpxV1, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			return
		}
	}

	dpx, err = NewDpxV1(options)
	if err != nil {
		return
	}

	err = dpx.Service.ConfigureService(options.ServiceName)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = dpx.Service.SetServiceURL(options.URL)
	}
	return
}

// NewDpxV1 : constructs an instance of DpxV1 with passed in options.
func NewDpxV1(options *DpxV1Options) (service *DpxV1, err error) {
	serviceOptions := &core.ServiceOptions{
		Authenticator: options.Authenticator,
	}

	baseService, err := core.NewBaseService(serviceOptions)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = baseService.SetServiceURL(options.URL)
		if err != nil {
			return
		}
	}

	service = &DpxV1{
		Service: baseService,
	}

	return
}

// GetServiceURLForRegion returns the service URL to be used for the specified region
func GetServiceURLForRegion(region string) (string, error) {
	return "", fmt.Errorf("service does not support regional URLs")
}

// Clone makes a copy of "dpx" suitable for processing requests.
func (dpx *DpxV1) Clone() *DpxV1 {
	if core.IsNil(dpx) {
		return nil
	}
	clone := *dpx
	clone.Service = dpx.Service.Clone()
	return &clone
}

// SetServiceURL sets the service URL
func (dpx *DpxV1) SetServiceURL(url string) error {
	return dpx.Service.SetServiceURL(url)
}

// GetServiceURL returns the service URL
func (dpx *DpxV1) GetServiceURL() string {
	return dpx.Service.GetServiceURL()
}

// SetDefaultHeaders sets HTTP headers to be sent in every request
func (dpx *DpxV1) SetDefaultHeaders(headers http.Header) {
	dpx.Service.SetDefaultHeaders(headers)
}

// SetEnableGzipCompression sets the service's EnableGzipCompression field
func (dpx *DpxV1) SetEnableGzipCompression(enableGzip bool) {
	dpx.Service.SetEnableGzipCompression(enableGzip)
}

// GetEnableGzipCompression returns the service's EnableGzipCompression field
func (dpx *DpxV1) GetEnableGzipCompression() bool {
	return dpx.Service.GetEnableGzipCompression()
}

// EnableRetries enables automatic retries for requests invoked for this service instance.
// If either parameter is specified as 0, then a default value is used instead.
func (dpx *DpxV1) EnableRetries(maxRetries int, maxRetryInterval time.Duration) {
	dpx.Service.EnableRetries(maxRetries, maxRetryInterval)
}

// DisableRetries disables automatic retries for requests invoked for this service instance.
func (dpx *DpxV1) DisableRetries() {
	dpx.Service.DisableRetries()
}

// GetInitializeStatus : Get the status of resources initialization in data product exchange
// Use this API to get the status of the resource initialization in data product exchange. <br/><br/>If the data product
// catalog exists but has never been initialized, the status will be "not_started".<br/>If the data product catalog
// exists and has been or is being initialized, the response will contain the status of the last or current
// initialization.If the initialization failed, the "errors" and the "trace" fields will contain the error(s)
// encountered during the initialization and the id to trace the error(s).<br/>If the data product catalog doesn't
// exist, a HTTP 404 response will be returned.
func (dpx *DpxV1) GetInitializeStatus(getInitializeStatusOptions *GetInitializeStatusOptions) (result *InitializeResource, response *core.DetailedResponse, err error) {
	return dpx.GetInitializeStatusWithContext(context.Background(), getInitializeStatusOptions)
}

// GetInitializeStatusWithContext is an alternate form of the GetInitializeStatus method which supports a Context parameter
func (dpx *DpxV1) GetInitializeStatusWithContext(ctx context.Context, getInitializeStatusOptions *GetInitializeStatusOptions) (result *InitializeResource, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getInitializeStatusOptions, "getInitializeStatusOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dpx.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dpx.Service.Options.URL, `/data_product_exchange/v1/configuration/initialize/status`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range getInitializeStatusOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("dpx", "V1", "GetInitializeStatus")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if getInitializeStatusOptions.ContainerID != nil {
		builder.AddQuery("container.id", fmt.Sprint(*getInitializeStatusOptions.ContainerID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dpx.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalInitializeResource)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// Initialize : Initialize resources in a data product exchange
// Use this API to initialize default assets for data product exchange. <br/><br/>You can initialize:
// <br/><ul><li>`delivery_methods` - Methods through which data product parts can be delivered to consumers of the data
// product exchange</li><li>`domains_multi_industry` - Taxonomy of domains and use cases applicable to multiple
// industries</li><li>`data_product_samples` - Sample data products used to illustrate capabilities of the data product
// exchange</li></ul><br/><br/>If a resource depends on resources that are not specified in the request, these dependent
// resources will be automatically initialized. E.g., initializing `data_product_samples` will also initialize
// `domains_multi_industry` and `delivery_methods` even if they are not specified in the request because it depends on
// them.<br/><br/>If initializing the data product exchange for the first time, do not specify a container. The default
// data product catalog will be created.<br/>For first time initialization, it is recommended that at least
// `delivery_methods` and `domains_multi_industry` is included in the initialize operation.<br/><br/>If the data product
// exchange has already been initialized, you may call this API again to initialize new resources, such as new delivery
// methods.In this case, specify the default data product catalog container information.
func (dpx *DpxV1) Initialize(initializeOptions *InitializeOptions) (result *InitializeResource, response *core.DetailedResponse, err error) {
	return dpx.InitializeWithContext(context.Background(), initializeOptions)
}

// InitializeWithContext is an alternate form of the Initialize method which supports a Context parameter
func (dpx *DpxV1) InitializeWithContext(ctx context.Context, initializeOptions *InitializeOptions) (result *InitializeResource, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(initializeOptions, "initializeOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(initializeOptions, "initializeOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dpx.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dpx.Service.Options.URL, `/data_product_exchange/v1/configuration/initialize`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range initializeOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("dpx", "V1", "Initialize")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if initializeOptions.Container != nil {
		body["container"] = initializeOptions.Container
	}
	if initializeOptions.Force != nil {
		body["force"] = initializeOptions.Force
	}
	if initializeOptions.Reinitialize != nil {
		body["reinitialize"] = initializeOptions.Reinitialize
	}
	if initializeOptions.Include != nil {
		body["include"] = initializeOptions.Include
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dpx.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalInitializeResource)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetDataProduct : Retrieve a data product identified by id
// Retrieve a data product identified by id.
func (dpx *DpxV1) GetDataProduct(getDataProductOptions *GetDataProductOptions) (result *DataProduct, response *core.DetailedResponse, err error) {
	return dpx.GetDataProductWithContext(context.Background(), getDataProductOptions)
}

// GetDataProductWithContext is an alternate form of the GetDataProduct method which supports a Context parameter
func (dpx *DpxV1) GetDataProductWithContext(ctx context.Context, getDataProductOptions *GetDataProductOptions) (result *DataProduct, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getDataProductOptions, "getDataProductOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getDataProductOptions, "getDataProductOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *getDataProductOptions.ID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dpx.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dpx.Service.Options.URL, `/data_product_exchange/v1/data_products/{id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getDataProductOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("dpx", "V1", "GetDataProduct")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dpx.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDataProduct)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListDataProducts : Retrieve a list of data products
// Retrieve a list of data products.
func (dpx *DpxV1) ListDataProducts(listDataProductsOptions *ListDataProductsOptions) (result *DataProductCollection, response *core.DetailedResponse, err error) {
	return dpx.ListDataProductsWithContext(context.Background(), listDataProductsOptions)
}

// ListDataProductsWithContext is an alternate form of the ListDataProducts method which supports a Context parameter
func (dpx *DpxV1) ListDataProductsWithContext(ctx context.Context, listDataProductsOptions *ListDataProductsOptions) (result *DataProductCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listDataProductsOptions, "listDataProductsOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dpx.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dpx.Service.Options.URL, `/data_product_exchange/v1/data_products`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range listDataProductsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("dpx", "V1", "ListDataProducts")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if listDataProductsOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*listDataProductsOptions.Limit))
	}
	if listDataProductsOptions.Start != nil {
		builder.AddQuery("start", fmt.Sprint(*listDataProductsOptions.Start))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dpx.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDataProductCollection)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListDataProductVersions : Retrieve a list of data product versions
// Retrieve a list of data product versions.
func (dpx *DpxV1) ListDataProductVersions(listDataProductVersionsOptions *ListDataProductVersionsOptions) (result *DataProductVersionCollection, response *core.DetailedResponse, err error) {
	return dpx.ListDataProductVersionsWithContext(context.Background(), listDataProductVersionsOptions)
}

// ListDataProductVersionsWithContext is an alternate form of the ListDataProductVersions method which supports a Context parameter
func (dpx *DpxV1) ListDataProductVersionsWithContext(ctx context.Context, listDataProductVersionsOptions *ListDataProductVersionsOptions) (result *DataProductVersionCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listDataProductVersionsOptions, "listDataProductVersionsOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dpx.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dpx.Service.Options.URL, `/data_product_exchange/v1/data_product_versions`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range listDataProductVersionsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("dpx", "V1", "ListDataProductVersions")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if listDataProductVersionsOptions.AssetContainerID != nil {
		builder.AddQuery("asset.container.id", fmt.Sprint(*listDataProductVersionsOptions.AssetContainerID))
	}
	if listDataProductVersionsOptions.DataProduct != nil {
		builder.AddQuery("data_product", fmt.Sprint(*listDataProductVersionsOptions.DataProduct))
	}
	if listDataProductVersionsOptions.State != nil {
		builder.AddQuery("state", fmt.Sprint(*listDataProductVersionsOptions.State))
	}
	if listDataProductVersionsOptions.Version != nil {
		builder.AddQuery("version", fmt.Sprint(*listDataProductVersionsOptions.Version))
	}
	if listDataProductVersionsOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*listDataProductVersionsOptions.Limit))
	}
	if listDataProductVersionsOptions.Start != nil {
		builder.AddQuery("start", fmt.Sprint(*listDataProductVersionsOptions.Start))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dpx.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDataProductVersionCollection)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateDataProductVersion : Create a new data product version
// Use this API to create a new data product version.<br/><br/>If the `state` is not specified, the data product version
// will be created in **draft** state.<br/><br/>**Create the first version of a data product**<br/><br/>Required
// fields:<br/><br/>- name<br/>- container<br/><br/>If `version` is not specified, the default version **1.0.0** will be
// used.<br/><br/>**Create a new version of an existing data product**<br/><br/>Required fields:<br/><br/>-
// container<br/>- data_product<br/>- version<br/><br/>The `domain` is required if state of data product is available.
// If no additional properties are specified, the values will be copied from the most recently available version of the
// data product.
func (dpx *DpxV1) CreateDataProductVersion(createDataProductVersionOptions *CreateDataProductVersionOptions) (result *DataProductVersion, response *core.DetailedResponse, err error) {
	return dpx.CreateDataProductVersionWithContext(context.Background(), createDataProductVersionOptions)
}

// CreateDataProductVersionWithContext is an alternate form of the CreateDataProductVersion method which supports a Context parameter
func (dpx *DpxV1) CreateDataProductVersionWithContext(ctx context.Context, createDataProductVersionOptions *CreateDataProductVersionOptions) (result *DataProductVersion, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createDataProductVersionOptions, "createDataProductVersionOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createDataProductVersionOptions, "createDataProductVersionOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dpx.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dpx.Service.Options.URL, `/data_product_exchange/v1/data_product_versions`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range createDataProductVersionOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("dpx", "V1", "CreateDataProductVersion")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if createDataProductVersionOptions.Container != nil {
		body["container"] = createDataProductVersionOptions.Container
	}
	if createDataProductVersionOptions.Version != nil {
		body["version"] = createDataProductVersionOptions.Version
	}
	if createDataProductVersionOptions.State != nil {
		body["state"] = createDataProductVersionOptions.State
	}
	if createDataProductVersionOptions.DataProduct != nil {
		body["data_product"] = createDataProductVersionOptions.DataProduct
	}
	if createDataProductVersionOptions.Name != nil {
		body["name"] = createDataProductVersionOptions.Name
	}
	if createDataProductVersionOptions.Description != nil {
		body["description"] = createDataProductVersionOptions.Description
	}
	if createDataProductVersionOptions.Tags != nil {
		body["tags"] = createDataProductVersionOptions.Tags
	}
	if createDataProductVersionOptions.UseCases != nil {
		body["use_cases"] = createDataProductVersionOptions.UseCases
	}
	if createDataProductVersionOptions.Domain != nil {
		body["domain"] = createDataProductVersionOptions.Domain
	}
	if createDataProductVersionOptions.Type != nil {
		body["type"] = createDataProductVersionOptions.Type
	}
	if createDataProductVersionOptions.PartsOut != nil {
		body["parts_out"] = createDataProductVersionOptions.PartsOut
	}
	if createDataProductVersionOptions.ContractTerms != nil {
		body["contract_terms"] = createDataProductVersionOptions.ContractTerms
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dpx.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDataProductVersion)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetDataProductVersion : Retrieve a data product version identified by ID
// Retrieve a data product version identified by a valid ID.
func (dpx *DpxV1) GetDataProductVersion(getDataProductVersionOptions *GetDataProductVersionOptions) (result *DataProductVersion, response *core.DetailedResponse, err error) {
	return dpx.GetDataProductVersionWithContext(context.Background(), getDataProductVersionOptions)
}

// GetDataProductVersionWithContext is an alternate form of the GetDataProductVersion method which supports a Context parameter
func (dpx *DpxV1) GetDataProductVersionWithContext(ctx context.Context, getDataProductVersionOptions *GetDataProductVersionOptions) (result *DataProductVersion, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getDataProductVersionOptions, "getDataProductVersionOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getDataProductVersionOptions, "getDataProductVersionOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *getDataProductVersionOptions.ID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dpx.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dpx.Service.Options.URL, `/data_product_exchange/v1/data_product_versions/{id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getDataProductVersionOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("dpx", "V1", "GetDataProductVersion")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dpx.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDataProductVersion)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteDataProductVersion : Delete a data product version identified by ID
// Delete a data product version identified by a valid ID. Delete can be performed only on data product versions in
// **draft** state. To retire a data product version which has already been published, use `PATCH
// /data_product_exchange/v1/data_product_versions` to change the data product version state to **retired**.
func (dpx *DpxV1) DeleteDataProductVersion(deleteDataProductVersionOptions *DeleteDataProductVersionOptions) (response *core.DetailedResponse, err error) {
	return dpx.DeleteDataProductVersionWithContext(context.Background(), deleteDataProductVersionOptions)
}

// DeleteDataProductVersionWithContext is an alternate form of the DeleteDataProductVersion method which supports a Context parameter
func (dpx *DpxV1) DeleteDataProductVersionWithContext(ctx context.Context, deleteDataProductVersionOptions *DeleteDataProductVersionOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteDataProductVersionOptions, "deleteDataProductVersionOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteDataProductVersionOptions, "deleteDataProductVersionOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *deleteDataProductVersionOptions.ID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dpx.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dpx.Service.Options.URL, `/data_product_exchange/v1/data_product_versions/{id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteDataProductVersionOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("dpx", "V1", "DeleteDataProductVersion")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = dpx.Service.Request(request, nil)

	return
}

// UpdateDataProductVersion : Update the data product version identified by ID
// Use this API to update the properties of a data product version identified by a valid ID.<br/><br/>Specify patch
// operations using http://jsonpatch.com/ syntax.<br/><br/>Supported patch operations include:<br/><br/>- Update the
// properties of a data product<br/><br/>- Add/Remove parts from a data product (up to 20 parts)<br/><br/>- Add/Remove
// use cases from a data product<br/><br/>- Update the data product state<br/><br/>.
func (dpx *DpxV1) UpdateDataProductVersion(updateDataProductVersionOptions *UpdateDataProductVersionOptions) (result *DataProductVersion, response *core.DetailedResponse, err error) {
	return dpx.UpdateDataProductVersionWithContext(context.Background(), updateDataProductVersionOptions)
}

// UpdateDataProductVersionWithContext is an alternate form of the UpdateDataProductVersion method which supports a Context parameter
func (dpx *DpxV1) UpdateDataProductVersionWithContext(ctx context.Context, updateDataProductVersionOptions *UpdateDataProductVersionOptions) (result *DataProductVersion, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateDataProductVersionOptions, "updateDataProductVersionOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateDataProductVersionOptions, "updateDataProductVersionOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *updateDataProductVersionOptions.ID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dpx.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dpx.Service.Options.URL, `/data_product_exchange/v1/data_product_versions/{id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateDataProductVersionOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("dpx", "V1", "UpdateDataProductVersion")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json-patch+json")

	_, err = builder.SetBodyContentJSON(updateDataProductVersionOptions.JSONPatchInstructions)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dpx.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDataProductVersion)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CompleteContractTermsDocument : Complete a contract document upload
// After uploading a file to the provided signed URL this endpoint is called to mark upload as complete, and make it
// available to download.
// - Once complete has been called on a document - returned URL will be using "url" field, which will contain a signed
// URL which can be used to download the document.
// - Calling complete on referential documents will result in an error.
// - Calling complete on attachment documents for which file has not been uploaded will result in an error.
//
// Contract terms documents can only be completed if the data product version which the contract terms are associated
// with is in a DRAFT state.
func (dpx *DpxV1) CompleteContractTermsDocument(completeContractTermsDocumentOptions *CompleteContractTermsDocumentOptions) (result *ContractTermsDocument, response *core.DetailedResponse, err error) {
	return dpx.CompleteContractTermsDocumentWithContext(context.Background(), completeContractTermsDocumentOptions)
}

// CompleteContractTermsDocumentWithContext is an alternate form of the CompleteContractTermsDocument method which supports a Context parameter
func (dpx *DpxV1) CompleteContractTermsDocumentWithContext(ctx context.Context, completeContractTermsDocumentOptions *CompleteContractTermsDocumentOptions) (result *ContractTermsDocument, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(completeContractTermsDocumentOptions, "completeContractTermsDocumentOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(completeContractTermsDocumentOptions, "completeContractTermsDocumentOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"data_product_version_id": *completeContractTermsDocumentOptions.DataProductVersionID,
		"contract_terms_id": *completeContractTermsDocumentOptions.ContractTermsID,
		"document_id": *completeContractTermsDocumentOptions.DocumentID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dpx.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dpx.Service.Options.URL, `/data_product_exchange/v1/data_product_versions/{data_product_version_id}/contract_terms/{contract_terms_id}/documents/{document_id}/complete`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range completeContractTermsDocumentOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("dpx", "V1", "CompleteContractTermsDocument")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dpx.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalContractTermsDocument)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateContractTermsDocument : Upload a contract document to the Data Product Version contract terms
// Upload a contract document to the Data Product Version identified by id.
//
// - If request object contains "url" parameter, a referential document will be created which will simply store provided
// url.
// - If request object does not contain "url" parameter, an attachment document will be created, and an "upload_url"
// parameter containing signed url will be returned. Client can upload the document using provided "upload_url". Once
// upload has been compeleted, "complete_contract_terms_document" for the given document needs to be called to mark
// attachment as completed. After completion of the attachment "get_contract_terms_document" for the given document will
// return signed "url" parameter that can be used to download perviously uploaded document.
//
// Contract terms documents can only be updated if the data product version which the contract terms are associated with
// is in a DRAFT state.
func (dpx *DpxV1) CreateContractTermsDocument(createContractTermsDocumentOptions *CreateContractTermsDocumentOptions) (result *ContractTermsDocument, response *core.DetailedResponse, err error) {
	return dpx.CreateContractTermsDocumentWithContext(context.Background(), createContractTermsDocumentOptions)
}

// CreateContractTermsDocumentWithContext is an alternate form of the CreateContractTermsDocument method which supports a Context parameter
func (dpx *DpxV1) CreateContractTermsDocumentWithContext(ctx context.Context, createContractTermsDocumentOptions *CreateContractTermsDocumentOptions) (result *ContractTermsDocument, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createContractTermsDocumentOptions, "createContractTermsDocumentOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createContractTermsDocumentOptions, "createContractTermsDocumentOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"data_product_version_id": *createContractTermsDocumentOptions.DataProductVersionID,
		"contract_terms_id": *createContractTermsDocumentOptions.ContractTermsID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dpx.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dpx.Service.Options.URL, `/data_product_exchange/v1/data_product_versions/{data_product_version_id}/contract_terms/{contract_terms_id}/documents`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range createContractTermsDocumentOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("dpx", "V1", "CreateContractTermsDocument")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if createContractTermsDocumentOptions.Type != nil {
		body["type"] = createContractTermsDocumentOptions.Type
	}
	if createContractTermsDocumentOptions.Name != nil {
		body["name"] = createContractTermsDocumentOptions.Name
	}
	if createContractTermsDocumentOptions.ID != nil {
		body["id"] = createContractTermsDocumentOptions.ID
	}
	if createContractTermsDocumentOptions.URL != nil {
		body["url"] = createContractTermsDocumentOptions.URL
	}
	if createContractTermsDocumentOptions.Attachment != nil {
		body["attachment"] = createContractTermsDocumentOptions.Attachment
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dpx.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalContractTermsDocument)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetContractTermsDocument : Get a contract document
// If document has an attachment that has been completed, - response will contain `url` which can be used to download
// the attachment. If document does not have an attachment, - the response will contain `url` which was submitted at
// document creation. If document has an attachment that has not been completed, - an error will be returned, prompting
// client to upload the document file and complete it, prior to retrieving it.
func (dpx *DpxV1) GetContractTermsDocument(getContractTermsDocumentOptions *GetContractTermsDocumentOptions) (result *ContractTermsDocument, response *core.DetailedResponse, err error) {
	return dpx.GetContractTermsDocumentWithContext(context.Background(), getContractTermsDocumentOptions)
}

// GetContractTermsDocumentWithContext is an alternate form of the GetContractTermsDocument method which supports a Context parameter
func (dpx *DpxV1) GetContractTermsDocumentWithContext(ctx context.Context, getContractTermsDocumentOptions *GetContractTermsDocumentOptions) (result *ContractTermsDocument, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getContractTermsDocumentOptions, "getContractTermsDocumentOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getContractTermsDocumentOptions, "getContractTermsDocumentOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"data_product_version_id": *getContractTermsDocumentOptions.DataProductVersionID,
		"contract_terms_id": *getContractTermsDocumentOptions.ContractTermsID,
		"document_id": *getContractTermsDocumentOptions.DocumentID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dpx.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dpx.Service.Options.URL, `/data_product_exchange/v1/data_product_versions/{data_product_version_id}/contract_terms/{contract_terms_id}/documents/{document_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getContractTermsDocumentOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("dpx", "V1", "GetContractTermsDocument")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dpx.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalContractTermsDocument)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteContractTermsDocument : Delete a contract document
// Delete an existing contract document.
//
// Contract terms documents can only be deleted if the data product version which the contract terms is associated with
// is in a DRAFT state.
func (dpx *DpxV1) DeleteContractTermsDocument(deleteContractTermsDocumentOptions *DeleteContractTermsDocumentOptions) (response *core.DetailedResponse, err error) {
	return dpx.DeleteContractTermsDocumentWithContext(context.Background(), deleteContractTermsDocumentOptions)
}

// DeleteContractTermsDocumentWithContext is an alternate form of the DeleteContractTermsDocument method which supports a Context parameter
func (dpx *DpxV1) DeleteContractTermsDocumentWithContext(ctx context.Context, deleteContractTermsDocumentOptions *DeleteContractTermsDocumentOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteContractTermsDocumentOptions, "deleteContractTermsDocumentOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteContractTermsDocumentOptions, "deleteContractTermsDocumentOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"data_product_version_id": *deleteContractTermsDocumentOptions.DataProductVersionID,
		"contract_terms_id": *deleteContractTermsDocumentOptions.ContractTermsID,
		"document_id": *deleteContractTermsDocumentOptions.DocumentID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dpx.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dpx.Service.Options.URL, `/data_product_exchange/v1/data_product_versions/{data_product_version_id}/contract_terms/{contract_terms_id}/documents/{document_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteContractTermsDocumentOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("dpx", "V1", "DeleteContractTermsDocument")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = dpx.Service.Request(request, nil)

	return
}

// UpdateContractTermsDocument : Update a contract document
// Use this API to update the properties of a contract document identified by a valid ID.
//
// Specify patch operations using http://jsonpatch.com/ syntax.
//
// Supported patch operations include:
// - Update the url of document if it does not have an attachment:
// - Update the type of the document Contract terms documents can only be updated if the data product version which the
// contract terms are associated with is in a DRAFT state.
func (dpx *DpxV1) UpdateContractTermsDocument(updateContractTermsDocumentOptions *UpdateContractTermsDocumentOptions) (result *ContractTermsDocument, response *core.DetailedResponse, err error) {
	return dpx.UpdateContractTermsDocumentWithContext(context.Background(), updateContractTermsDocumentOptions)
}

// UpdateContractTermsDocumentWithContext is an alternate form of the UpdateContractTermsDocument method which supports a Context parameter
func (dpx *DpxV1) UpdateContractTermsDocumentWithContext(ctx context.Context, updateContractTermsDocumentOptions *UpdateContractTermsDocumentOptions) (result *ContractTermsDocument, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateContractTermsDocumentOptions, "updateContractTermsDocumentOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateContractTermsDocumentOptions, "updateContractTermsDocumentOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"data_product_version_id": *updateContractTermsDocumentOptions.DataProductVersionID,
		"contract_terms_id": *updateContractTermsDocumentOptions.ContractTermsID,
		"document_id": *updateContractTermsDocumentOptions.DocumentID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dpx.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dpx.Service.Options.URL, `/data_product_exchange/v1/data_product_versions/{data_product_version_id}/contract_terms/{contract_terms_id}/documents/{document_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateContractTermsDocumentOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("dpx", "V1", "UpdateContractTermsDocument")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json-patch+json")

	_, err = builder.SetBodyContentJSON(updateContractTermsDocumentOptions.JSONPatchInstructions)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dpx.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalContractTermsDocument)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// AssetPartReference : The asset represented in this part.
type AssetPartReference struct {
	// The unique identifier of the asset.
	ID *string `json:"id" validate:"required"`

	// Data product exchange container.
	Container *ContainerReference `json:"container" validate:"required"`

	// The type of the asset.
	Type *string `json:"type,omitempty"`
}

// NewAssetPartReference : Instantiate AssetPartReference (Generic Model Constructor)
func (*DpxV1) NewAssetPartReference(id string, container *ContainerReference) (_model *AssetPartReference, err error) {
	_model = &AssetPartReference{
		ID: core.StringPtr(id),
		Container: container,
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalAssetPartReference unmarshals an instance of AssetPartReference from the specified map of raw messages.
func UnmarshalAssetPartReference(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AssetPartReference)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "container", &obj.Container, UnmarshalContainerReference)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AssetReference : AssetReference struct
type AssetReference struct {
	// The unique identifier of the asset.
	ID *string `json:"id" validate:"required"`

	// Data product exchange container.
	Container *ContainerReference `json:"container" validate:"required"`
}

// NewAssetReference : Instantiate AssetReference (Generic Model Constructor)
func (*DpxV1) NewAssetReference(id string, container *ContainerReference) (_model *AssetReference, err error) {
	_model = &AssetReference{
		ID: core.StringPtr(id),
		Container: container,
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalAssetReference unmarshals an instance of AssetReference from the specified map of raw messages.
func UnmarshalAssetReference(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AssetReference)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "container", &obj.Container, UnmarshalContainerReference)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CompleteContractTermsDocumentOptions : The CompleteContractTermsDocument options.
type CompleteContractTermsDocumentOptions struct {
	// Data product version id.
	DataProductVersionID *string `json:"data_product_version_id" validate:"required,ne="`

	// Contract terms id.
	ContractTermsID *string `json:"contract_terms_id" validate:"required,ne="`

	// Document id.
	DocumentID *string `json:"document_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCompleteContractTermsDocumentOptions : Instantiate CompleteContractTermsDocumentOptions
func (*DpxV1) NewCompleteContractTermsDocumentOptions(dataProductVersionID string, contractTermsID string, documentID string) *CompleteContractTermsDocumentOptions {
	return &CompleteContractTermsDocumentOptions{
		DataProductVersionID: core.StringPtr(dataProductVersionID),
		ContractTermsID: core.StringPtr(contractTermsID),
		DocumentID: core.StringPtr(documentID),
	}
}

// SetDataProductVersionID : Allow user to set DataProductVersionID
func (_options *CompleteContractTermsDocumentOptions) SetDataProductVersionID(dataProductVersionID string) *CompleteContractTermsDocumentOptions {
	_options.DataProductVersionID = core.StringPtr(dataProductVersionID)
	return _options
}

// SetContractTermsID : Allow user to set ContractTermsID
func (_options *CompleteContractTermsDocumentOptions) SetContractTermsID(contractTermsID string) *CompleteContractTermsDocumentOptions {
	_options.ContractTermsID = core.StringPtr(contractTermsID)
	return _options
}

// SetDocumentID : Allow user to set DocumentID
func (_options *CompleteContractTermsDocumentOptions) SetDocumentID(documentID string) *CompleteContractTermsDocumentOptions {
	_options.DocumentID = core.StringPtr(documentID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CompleteContractTermsDocumentOptions) SetHeaders(param map[string]string) *CompleteContractTermsDocumentOptions {
	options.Headers = param
	return options
}

// ContainerReference : Data product exchange container.
type ContainerReference struct {
	// Container identifier.
	ID *string `json:"id" validate:"required"`

	// Container type.
	Type *string `json:"type,omitempty"`
}

// Constants associated with the ContainerReference.Type property.
// Container type.
const (
	ContainerReference_Type_Catalog = "catalog"
)

// NewContainerReference : Instantiate ContainerReference (Generic Model Constructor)
func (*DpxV1) NewContainerReference(id string) (_model *ContainerReference, err error) {
	_model = &ContainerReference{
		ID: core.StringPtr(id),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalContainerReference unmarshals an instance of ContainerReference from the specified map of raw messages.
func UnmarshalContainerReference(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ContainerReference)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ContractTermsDocument : ContractTermsDocument struct
type ContractTermsDocument struct {
	// URL which can be used to retrieve the contract document.
	URL *string `json:"url,omitempty"`

	// Type of the contract document.
	Type *string `json:"type" validate:"required"`

	// Name of the contract document.
	Name *string `json:"name" validate:"required"`

	// Id uniquely identifying this document within contract terms instance.
	ID *string `json:"id" validate:"required"`

	// Attachment associated witht the document.
	Attachment *ContractTermsDocumentAttachment `json:"attachment,omitempty"`
}

// Constants associated with the ContractTermsDocument.Type property.
// Type of the contract document.
const (
	ContractTermsDocument_Type_Sla = "sla"
	ContractTermsDocument_Type_TermsAndConditions = "terms_and_conditions"
)

// NewContractTermsDocument : Instantiate ContractTermsDocument (Generic Model Constructor)
func (*DpxV1) NewContractTermsDocument(typeVar string, name string, id string) (_model *ContractTermsDocument, err error) {
	_model = &ContractTermsDocument{
		Type: core.StringPtr(typeVar),
		Name: core.StringPtr(name),
		ID: core.StringPtr(id),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalContractTermsDocument unmarshals an instance of ContractTermsDocument from the specified map of raw messages.
func UnmarshalContractTermsDocument(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ContractTermsDocument)
	err = core.UnmarshalPrimitive(m, "url", &obj.URL)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "attachment", &obj.Attachment, UnmarshalContractTermsDocumentAttachment)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

func (*DpxV1) NewContractTermsDocumentPatch(contractTermsDocument *ContractTermsDocument) (_patch []JSONPatchOperation) {
	if (contractTermsDocument.URL != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/url"),
			Value: contractTermsDocument.URL,
		})
	}
	if (contractTermsDocument.Type != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/type"),
			Value: contractTermsDocument.Type,
		})
	}
	if (contractTermsDocument.Name != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/name"),
			Value: contractTermsDocument.Name,
		})
	}
	if (contractTermsDocument.ID != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/id"),
			Value: contractTermsDocument.ID,
		})
	}
	if (contractTermsDocument.Attachment != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/attachment"),
			Value: contractTermsDocument.Attachment,
		})
	}
	return
}

// ContractTermsDocumentAttachment : Attachment associated witht the document.
type ContractTermsDocumentAttachment struct {
	// Id representing the corresponding attachment.
	ID *string `json:"id,omitempty"`
}

// UnmarshalContractTermsDocumentAttachment unmarshals an instance of ContractTermsDocumentAttachment from the specified map of raw messages.
func UnmarshalContractTermsDocumentAttachment(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ContractTermsDocumentAttachment)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CreateContractTermsDocumentOptions : The CreateContractTermsDocument options.
type CreateContractTermsDocumentOptions struct {
	// Data product version id.
	DataProductVersionID *string `json:"data_product_version_id" validate:"required,ne="`

	// Contract terms id.
	ContractTermsID *string `json:"contract_terms_id" validate:"required,ne="`

	// Type of the contract document.
	Type *string `json:"type" validate:"required"`

	// Name of the contract document.
	Name *string `json:"name" validate:"required"`

	// Id uniquely identifying this document within contract terms instance.
	ID *string `json:"id" validate:"required"`

	// URL which can be used to retrieve the contract document.
	URL *string `json:"url,omitempty"`

	// Attachment associated witht the document.
	Attachment *ContractTermsDocumentAttachment `json:"attachment,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the CreateContractTermsDocumentOptions.Type property.
// Type of the contract document.
const (
	CreateContractTermsDocumentOptions_Type_Sla = "sla"
	CreateContractTermsDocumentOptions_Type_TermsAndConditions = "terms_and_conditions"
)

// NewCreateContractTermsDocumentOptions : Instantiate CreateContractTermsDocumentOptions
func (*DpxV1) NewCreateContractTermsDocumentOptions(dataProductVersionID string, contractTermsID string, typeVar string, name string, id string) *CreateContractTermsDocumentOptions {
	return &CreateContractTermsDocumentOptions{
		DataProductVersionID: core.StringPtr(dataProductVersionID),
		ContractTermsID: core.StringPtr(contractTermsID),
		Type: core.StringPtr(typeVar),
		Name: core.StringPtr(name),
		ID: core.StringPtr(id),
	}
}

// SetDataProductVersionID : Allow user to set DataProductVersionID
func (_options *CreateContractTermsDocumentOptions) SetDataProductVersionID(dataProductVersionID string) *CreateContractTermsDocumentOptions {
	_options.DataProductVersionID = core.StringPtr(dataProductVersionID)
	return _options
}

// SetContractTermsID : Allow user to set ContractTermsID
func (_options *CreateContractTermsDocumentOptions) SetContractTermsID(contractTermsID string) *CreateContractTermsDocumentOptions {
	_options.ContractTermsID = core.StringPtr(contractTermsID)
	return _options
}

// SetType : Allow user to set Type
func (_options *CreateContractTermsDocumentOptions) SetType(typeVar string) *CreateContractTermsDocumentOptions {
	_options.Type = core.StringPtr(typeVar)
	return _options
}

// SetName : Allow user to set Name
func (_options *CreateContractTermsDocumentOptions) SetName(name string) *CreateContractTermsDocumentOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetID : Allow user to set ID
func (_options *CreateContractTermsDocumentOptions) SetID(id string) *CreateContractTermsDocumentOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetURL : Allow user to set URL
func (_options *CreateContractTermsDocumentOptions) SetURL(url string) *CreateContractTermsDocumentOptions {
	_options.URL = core.StringPtr(url)
	return _options
}

// SetAttachment : Allow user to set Attachment
func (_options *CreateContractTermsDocumentOptions) SetAttachment(attachment *ContractTermsDocumentAttachment) *CreateContractTermsDocumentOptions {
	_options.Attachment = attachment
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateContractTermsDocumentOptions) SetHeaders(param map[string]string) *CreateContractTermsDocumentOptions {
	options.Headers = param
	return options
}

// CreateDataProductVersionOptions : The CreateDataProductVersion options.
type CreateDataProductVersionOptions struct {
	// Data product exchange container.
	Container *ContainerReference `json:"container" validate:"required"`

	// The data product version number.
	Version *string `json:"version,omitempty"`

	// The state of the data product version. If not specified, the data product version will be created in `draft` state.
	State *string `json:"state,omitempty"`

	// Data product identifier.
	DataProduct *DataProductIdentity `json:"data_product,omitempty"`

	// The name to use to refer to the new data product version. If this is a new data product, this value must be
	// specified. If this is a new version of an existing data product, the name will default to the name of the previous
	// data product version. A name can contain letters, numbers, understores, dashes, spaces or periods. A name must
	// contain at least one non-space character.
	Name *string `json:"name,omitempty"`

	// Description of the data product version. If this is a new version of an existing data product, the description will
	// default to the description of the previous version of the data product.
	Description *string `json:"description,omitempty"`

	// Tags on the new data product version. If this is the first version of a data product, tags defaults to an empty
	// list. If this is a new version of an existing data product, tags will default to the list of tags on the previous
	// version of the data product.
	Tags []string `json:"tags,omitempty"`

	// Use cases that the data product version serves. If this is the first version of a data product, use cases defaults
	// to an empty list. If this is a new version of an existing data product, use cases will default to the list of use
	// cases on the previous version of the data product.
	UseCases []UseCase `json:"use_cases,omitempty"`

	// The business domain associated with the data product version.
	Domain *Domain `json:"domain,omitempty"`

	// The types of the parts included in this data product version. If this is the first version of a data product, this
	// field defaults to an empty list. If this is a new version of an existing data product, the types will default to the
	// types of the previous version of the data product.
	Type []string `json:"type,omitempty"`

	// The outgoing parts of this data product version to be delivered to consumers. If this is the first version of a data
	// product, this field defaults to an empty list. If this is a new version of an existing data product, the data
	// product parts will default to the parts list from the previous version of the data product.
	PartsOut []DataProductPart `json:"parts_out,omitempty"`

	// The contract terms that bind interactions with this data product version.
	ContractTerms []DataProductContractTerms `json:"contract_terms,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the CreateDataProductVersionOptions.State property.
// The state of the data product version. If not specified, the data product version will be created in `draft` state.
const (
	CreateDataProductVersionOptions_State_Available = "available"
	CreateDataProductVersionOptions_State_Draft = "draft"
	CreateDataProductVersionOptions_State_Retired = "retired"
)

// Constants associated with the CreateDataProductVersionOptions.Type property.
const (
	CreateDataProductVersionOptions_Type_Code = "code"
	CreateDataProductVersionOptions_Type_Data = "data"
)

// NewCreateDataProductVersionOptions : Instantiate CreateDataProductVersionOptions
func (*DpxV1) NewCreateDataProductVersionOptions(container *ContainerReference) *CreateDataProductVersionOptions {
	return &CreateDataProductVersionOptions{
		Container: container,
	}
}

// SetContainer : Allow user to set Container
func (_options *CreateDataProductVersionOptions) SetContainer(container *ContainerReference) *CreateDataProductVersionOptions {
	_options.Container = container
	return _options
}

// SetVersion : Allow user to set Version
func (_options *CreateDataProductVersionOptions) SetVersion(version string) *CreateDataProductVersionOptions {
	_options.Version = core.StringPtr(version)
	return _options
}

// SetState : Allow user to set State
func (_options *CreateDataProductVersionOptions) SetState(state string) *CreateDataProductVersionOptions {
	_options.State = core.StringPtr(state)
	return _options
}

// SetDataProduct : Allow user to set DataProduct
func (_options *CreateDataProductVersionOptions) SetDataProduct(dataProduct *DataProductIdentity) *CreateDataProductVersionOptions {
	_options.DataProduct = dataProduct
	return _options
}

// SetName : Allow user to set Name
func (_options *CreateDataProductVersionOptions) SetName(name string) *CreateDataProductVersionOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetDescription : Allow user to set Description
func (_options *CreateDataProductVersionOptions) SetDescription(description string) *CreateDataProductVersionOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetTags : Allow user to set Tags
func (_options *CreateDataProductVersionOptions) SetTags(tags []string) *CreateDataProductVersionOptions {
	_options.Tags = tags
	return _options
}

// SetUseCases : Allow user to set UseCases
func (_options *CreateDataProductVersionOptions) SetUseCases(useCases []UseCase) *CreateDataProductVersionOptions {
	_options.UseCases = useCases
	return _options
}

// SetDomain : Allow user to set Domain
func (_options *CreateDataProductVersionOptions) SetDomain(domain *Domain) *CreateDataProductVersionOptions {
	_options.Domain = domain
	return _options
}

// SetType : Allow user to set Type
func (_options *CreateDataProductVersionOptions) SetType(typeVar []string) *CreateDataProductVersionOptions {
	_options.Type = typeVar
	return _options
}

// SetPartsOut : Allow user to set PartsOut
func (_options *CreateDataProductVersionOptions) SetPartsOut(partsOut []DataProductPart) *CreateDataProductVersionOptions {
	_options.PartsOut = partsOut
	return _options
}

// SetContractTerms : Allow user to set ContractTerms
func (_options *CreateDataProductVersionOptions) SetContractTerms(contractTerms []DataProductContractTerms) *CreateDataProductVersionOptions {
	_options.ContractTerms = contractTerms
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateDataProductVersionOptions) SetHeaders(param map[string]string) *CreateDataProductVersionOptions {
	options.Headers = param
	return options
}

// DataProduct : Data Product.
type DataProduct struct {
	// Data product identifier.
	ID *string `json:"id" validate:"required"`

	// Data product exchange container.
	Container *ContainerReference `json:"container" validate:"required"`

	// Name to refer to the data product.
	Name *string `json:"name" validate:"required"`
}

// UnmarshalDataProduct unmarshals an instance of DataProduct from the specified map of raw messages.
func UnmarshalDataProduct(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DataProduct)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "container", &obj.Container, UnmarshalContainerReference)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DataProductCollection : A collection of data products.
type DataProductCollection struct {
	// Set a limit on the number of results returned.
	Limit *int64 `json:"limit" validate:"required"`

	// First page in the collection.
	First *FirstPage `json:"first" validate:"required"`

	// Next page in the collection.
	Next *NextPage `json:"next,omitempty"`

	// Collection of data products.
	DataProducts []DataProduct `json:"data_products" validate:"required"`
}

// UnmarshalDataProductCollection unmarshals an instance of DataProductCollection from the specified map of raw messages.
func UnmarshalDataProductCollection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DataProductCollection)
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "first", &obj.First, UnmarshalFirstPage)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "next", &obj.Next, UnmarshalNextPage)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "data_products", &obj.DataProducts, UnmarshalDataProduct)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Retrieve the value to be passed to a request to access the next page of results
func (resp *DataProductCollection) GetNextStart() (*string, error) {
	if core.IsNil(resp.Next) {
		return nil, nil
	}
	return resp.Next.Start, nil
}

// DataProductContractTerms : DataProductContractTerms struct
type DataProductContractTerms struct {
	ID *string `json:"id,omitempty"`

	Documents []ContractTermsDocument `json:"documents,omitempty"`

	Asset *AssetReference `json:"asset,omitempty"`
}

// UnmarshalDataProductContractTerms unmarshals an instance of DataProductContractTerms from the specified map of raw messages.
func UnmarshalDataProductContractTerms(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DataProductContractTerms)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "documents", &obj.Documents, UnmarshalContractTermsDocument)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "asset", &obj.Asset, UnmarshalAssetReference)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DataProductIdentity : Data product identifier.
type DataProductIdentity struct {
	// Data product identifier.
	ID *string `json:"id" validate:"required"`
}

// NewDataProductIdentity : Instantiate DataProductIdentity (Generic Model Constructor)
func (*DpxV1) NewDataProductIdentity(id string) (_model *DataProductIdentity, err error) {
	_model = &DataProductIdentity{
		ID: core.StringPtr(id),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalDataProductIdentity unmarshals an instance of DataProductIdentity from the specified map of raw messages.
func UnmarshalDataProductIdentity(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DataProductIdentity)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DataProductPart : DataProductPart struct
type DataProductPart struct {
	// The asset represented in this part.
	Asset *AssetPartReference `json:"asset" validate:"required"`

	// The revision number of the asset represented in this part.
	Revision *int64 `json:"revision,omitempty"`

	// The time for when the part was last updated.
	UpdatedAt *strfmt.DateTime `json:"updated_at,omitempty"`

	// Delivery methods describing the delivery options available for this part.
	DeliveryMethods []DeliveryMethod `json:"delivery_methods,omitempty"`
}

// NewDataProductPart : Instantiate DataProductPart (Generic Model Constructor)
func (*DpxV1) NewDataProductPart(asset *AssetPartReference) (_model *DataProductPart, err error) {
	_model = &DataProductPart{
		Asset: asset,
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalDataProductPart unmarshals an instance of DataProductPart from the specified map of raw messages.
func UnmarshalDataProductPart(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DataProductPart)
	err = core.UnmarshalModel(m, "asset", &obj.Asset, UnmarshalAssetPartReference)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "revision", &obj.Revision)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "updated_at", &obj.UpdatedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "delivery_methods", &obj.DeliveryMethods, UnmarshalDeliveryMethod)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DataProductVersion : Data Product version.
type DataProductVersion struct {
	// The data product version number.
	Version *string `json:"version" validate:"required"`

	// The state of the data product version.
	State *string `json:"state" validate:"required"`

	// Data product identifier.
	DataProduct *DataProductIdentity `json:"data_product" validate:"required"`

	// The name of the data product version. A name can contain letters, numbers, understores, dashes, spaces or periods.
	// Names are mutable and reusable.
	Name *string `json:"name" validate:"required"`

	// The description of the data product version.
	Description *string `json:"description" validate:"required"`

	// The identifier of the data product version.
	ID *string `json:"id" validate:"required"`

	Asset *AssetReference `json:"asset" validate:"required"`

	// Tags on the data product.
	Tags []string `json:"tags,omitempty"`

	// A list of use cases associated with the data product version.
	UseCases []UseCase `json:"use_cases,omitempty"`

	// The business domain associated with the data product version.
	Domain *Domain `json:"domain" validate:"required"`

	// Type of parts on the data product.
	Type []string `json:"type,omitempty"`

	// Outgoing parts of a data product used to deliver the data product to consumers.
	PartsOut []DataProductPart `json:"parts_out" validate:"required"`

	// The user who published this data product version.
	PublishedBy *string `json:"published_by,omitempty"`

	// The time when this data product version was published.
	PublishedAt *strfmt.DateTime `json:"published_at,omitempty"`

	// Contract terms binding various aspects of the data product.
	ContractTerms []DataProductContractTerms `json:"contract_terms,omitempty"`

	// The creator of this data product version.
	CreatedBy *string `json:"created_by" validate:"required"`

	// The time when this data product version was created.
	CreatedAt *strfmt.DateTime `json:"created_at" validate:"required"`
}

// Constants associated with the DataProductVersion.State property.
// The state of the data product version.
const (
	DataProductVersion_State_Available = "available"
	DataProductVersion_State_Draft = "draft"
	DataProductVersion_State_Retired = "retired"
)

// Constants associated with the DataProductVersion.Type property.
const (
	DataProductVersion_Type_Code = "code"
	DataProductVersion_Type_Data = "data"
)

// UnmarshalDataProductVersion unmarshals an instance of DataProductVersion from the specified map of raw messages.
func UnmarshalDataProductVersion(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DataProductVersion)
	err = core.UnmarshalPrimitive(m, "version", &obj.Version)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "state", &obj.State)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "data_product", &obj.DataProduct, UnmarshalDataProductIdentity)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "asset", &obj.Asset, UnmarshalAssetReference)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "tags", &obj.Tags)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "use_cases", &obj.UseCases, UnmarshalUseCase)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "domain", &obj.Domain, UnmarshalDomain)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "parts_out", &obj.PartsOut, UnmarshalDataProductPart)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "published_by", &obj.PublishedBy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "published_at", &obj.PublishedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "contract_terms", &obj.ContractTerms, UnmarshalDataProductContractTerms)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_by", &obj.CreatedBy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_at", &obj.CreatedAt)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

func (*DpxV1) NewDataProductVersionPatch(dataProductVersion *DataProductVersion) (_patch []JSONPatchOperation) {
	if (dataProductVersion.Version != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/version"),
			Value: dataProductVersion.Version,
		})
	}
	if (dataProductVersion.State != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/state"),
			Value: dataProductVersion.State,
		})
	}
	if (dataProductVersion.DataProduct != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/data_product"),
			Value: dataProductVersion.DataProduct,
		})
	}
	if (dataProductVersion.Name != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/name"),
			Value: dataProductVersion.Name,
		})
	}
	if (dataProductVersion.Description != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/description"),
			Value: dataProductVersion.Description,
		})
	}
	if (dataProductVersion.ID != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/id"),
			Value: dataProductVersion.ID,
		})
	}
	if (dataProductVersion.Asset != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/asset"),
			Value: dataProductVersion.Asset,
		})
	}
	if (dataProductVersion.Tags != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/tags"),
			Value: dataProductVersion.Tags,
		})
	}
	if (dataProductVersion.UseCases != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/use_cases"),
			Value: dataProductVersion.UseCases,
		})
	}
	if (dataProductVersion.Domain != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/domain"),
			Value: dataProductVersion.Domain,
		})
	}
	if (dataProductVersion.Type != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/type"),
			Value: dataProductVersion.Type,
		})
	}
	if (dataProductVersion.PartsOut != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/parts_out"),
			Value: dataProductVersion.PartsOut,
		})
	}
	if (dataProductVersion.PublishedBy != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/published_by"),
			Value: dataProductVersion.PublishedBy,
		})
	}
	if (dataProductVersion.PublishedAt != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/published_at"),
			Value: dataProductVersion.PublishedAt,
		})
	}
	if (dataProductVersion.ContractTerms != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/contract_terms"),
			Value: dataProductVersion.ContractTerms,
		})
	}
	if (dataProductVersion.CreatedBy != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/created_by"),
			Value: dataProductVersion.CreatedBy,
		})
	}
	if (dataProductVersion.CreatedAt != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/created_at"),
			Value: dataProductVersion.CreatedAt,
		})
	}
	return
}

// DataProductVersionCollection : A collection of data product version summaries.
type DataProductVersionCollection struct {
	// Set a limit on the number of results returned.
	Limit *int64 `json:"limit" validate:"required"`

	// First page in the collection.
	First *FirstPage `json:"first" validate:"required"`

	// Next page in the collection.
	Next *NextPage `json:"next,omitempty"`

	// Collection of data product versions.
	DataProductVersions []DataProductVersionSummary `json:"data_product_versions" validate:"required"`
}

// UnmarshalDataProductVersionCollection unmarshals an instance of DataProductVersionCollection from the specified map of raw messages.
func UnmarshalDataProductVersionCollection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DataProductVersionCollection)
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "first", &obj.First, UnmarshalFirstPage)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "next", &obj.Next, UnmarshalNextPage)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "data_product_versions", &obj.DataProductVersions, UnmarshalDataProductVersionSummary)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Retrieve the value to be passed to a request to access the next page of results
func (resp *DataProductVersionCollection) GetNextStart() (*string, error) {
	if core.IsNil(resp.Next) {
		return nil, nil
	}
	return resp.Next.Start, nil
}

// DataProductVersionSummary : DataProductVersionSummary struct
type DataProductVersionSummary struct {
	// The data product version number.
	Version *string `json:"version" validate:"required"`

	// The state of the data product version.
	State *string `json:"state" validate:"required"`

	// Data product identifier.
	DataProduct *DataProductIdentity `json:"data_product" validate:"required"`

	// The name of the data product version. A name can contain letters, numbers, understores, dashes, spaces or periods.
	// Names are mutable and reusable.
	Name *string `json:"name" validate:"required"`

	// The description of the data product version.
	Description *string `json:"description" validate:"required"`

	// The identifier of the data product version.
	ID *string `json:"id" validate:"required"`

	Asset *AssetReference `json:"asset" validate:"required"`
}

// Constants associated with the DataProductVersionSummary.State property.
// The state of the data product version.
const (
	DataProductVersionSummary_State_Available = "available"
	DataProductVersionSummary_State_Draft = "draft"
	DataProductVersionSummary_State_Retired = "retired"
)

// UnmarshalDataProductVersionSummary unmarshals an instance of DataProductVersionSummary from the specified map of raw messages.
func UnmarshalDataProductVersionSummary(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DataProductVersionSummary)
	err = core.UnmarshalPrimitive(m, "version", &obj.Version)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "state", &obj.State)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "data_product", &obj.DataProduct, UnmarshalDataProductIdentity)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "asset", &obj.Asset, UnmarshalAssetReference)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DeleteContractTermsDocumentOptions : The DeleteContractTermsDocument options.
type DeleteContractTermsDocumentOptions struct {
	// Data product version id.
	DataProductVersionID *string `json:"data_product_version_id" validate:"required,ne="`

	// Contract terms id.
	ContractTermsID *string `json:"contract_terms_id" validate:"required,ne="`

	// Document id.
	DocumentID *string `json:"document_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteContractTermsDocumentOptions : Instantiate DeleteContractTermsDocumentOptions
func (*DpxV1) NewDeleteContractTermsDocumentOptions(dataProductVersionID string, contractTermsID string, documentID string) *DeleteContractTermsDocumentOptions {
	return &DeleteContractTermsDocumentOptions{
		DataProductVersionID: core.StringPtr(dataProductVersionID),
		ContractTermsID: core.StringPtr(contractTermsID),
		DocumentID: core.StringPtr(documentID),
	}
}

// SetDataProductVersionID : Allow user to set DataProductVersionID
func (_options *DeleteContractTermsDocumentOptions) SetDataProductVersionID(dataProductVersionID string) *DeleteContractTermsDocumentOptions {
	_options.DataProductVersionID = core.StringPtr(dataProductVersionID)
	return _options
}

// SetContractTermsID : Allow user to set ContractTermsID
func (_options *DeleteContractTermsDocumentOptions) SetContractTermsID(contractTermsID string) *DeleteContractTermsDocumentOptions {
	_options.ContractTermsID = core.StringPtr(contractTermsID)
	return _options
}

// SetDocumentID : Allow user to set DocumentID
func (_options *DeleteContractTermsDocumentOptions) SetDocumentID(documentID string) *DeleteContractTermsDocumentOptions {
	_options.DocumentID = core.StringPtr(documentID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteContractTermsDocumentOptions) SetHeaders(param map[string]string) *DeleteContractTermsDocumentOptions {
	options.Headers = param
	return options
}

// DeleteDataProductVersionOptions : The DeleteDataProductVersion options.
type DeleteDataProductVersionOptions struct {
	// Data product version ID.
	ID *string `json:"id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteDataProductVersionOptions : Instantiate DeleteDataProductVersionOptions
func (*DpxV1) NewDeleteDataProductVersionOptions(id string) *DeleteDataProductVersionOptions {
	return &DeleteDataProductVersionOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (_options *DeleteDataProductVersionOptions) SetID(id string) *DeleteDataProductVersionOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteDataProductVersionOptions) SetHeaders(param map[string]string) *DeleteDataProductVersionOptions {
	options.Headers = param
	return options
}

// DeliveryMethod : DeliveryMethod struct
type DeliveryMethod struct {
	// The ID of the delivery method.
	ID *string `json:"id" validate:"required"`

	// Data product exchange container.
	Container *ContainerReference `json:"container" validate:"required"`
}

// NewDeliveryMethod : Instantiate DeliveryMethod (Generic Model Constructor)
func (*DpxV1) NewDeliveryMethod(id string, container *ContainerReference) (_model *DeliveryMethod, err error) {
	_model = &DeliveryMethod{
		ID: core.StringPtr(id),
		Container: container,
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalDeliveryMethod unmarshals an instance of DeliveryMethod from the specified map of raw messages.
func UnmarshalDeliveryMethod(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DeliveryMethod)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "container", &obj.Container, UnmarshalContainerReference)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Domain : The business domain associated with the data product version.
type Domain struct {
	// The ID of the domain.
	ID *string `json:"id" validate:"required"`

	// The display name of the domain.
	Name *string `json:"name,omitempty"`

	// Data product exchange container.
	Container *ContainerReference `json:"container,omitempty"`
}

// NewDomain : Instantiate Domain (Generic Model Constructor)
func (*DpxV1) NewDomain(id string) (_model *Domain, err error) {
	_model = &Domain{
		ID: core.StringPtr(id),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalDomain unmarshals an instance of Domain from the specified map of raw messages.
func UnmarshalDomain(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Domain)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "container", &obj.Container, UnmarshalContainerReference)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ErrorModelResource : Detailed error information.
type ErrorModelResource struct {
	// Error code.
	Code *string `json:"code,omitempty"`

	// Error message.
	Message *string `json:"message,omitempty"`

	// Extra information about the error.
	Extra map[string]interface{} `json:"extra,omitempty"`

	// More info message.
	MoreInfo *string `json:"more_info,omitempty"`
}

// Constants associated with the ErrorModelResource.Code property.
// Error code.
const (
	ErrorModelResource_Code_AlreadyExists = "already_exists"
	ErrorModelResource_Code_ConfigurationError = "configuration_error"
	ErrorModelResource_Code_Conflict = "conflict"
	ErrorModelResource_Code_ConstraintViolation = "constraint_violation"
	ErrorModelResource_Code_CreateError = "create_error"
	ErrorModelResource_Code_DataError = "data_error"
	ErrorModelResource_Code_DatabaseError = "database_error"
	ErrorModelResource_Code_DatabaseQueryError = "database_query_error"
	ErrorModelResource_Code_DatabaseUsageLimits = "database_usage_limits"
	ErrorModelResource_Code_DeleteError = "delete_error"
	ErrorModelResource_Code_Deleted = "deleted"
	ErrorModelResource_Code_DependentServiceError = "dependent_service_error"
	ErrorModelResource_Code_DoesNotExist = "does_not_exist"
	ErrorModelResource_Code_EntitlementEnforcement = "entitlement_enforcement"
	ErrorModelResource_Code_FetchError = "fetch_error"
	ErrorModelResource_Code_Forbidden = "forbidden"
	ErrorModelResource_Code_GovernancePolicyDenial = "governance_policy_denial"
	ErrorModelResource_Code_InactiveUser = "inactive_user"
	ErrorModelResource_Code_InvalidParameter = "invalid_parameter"
	ErrorModelResource_Code_MissingRequiredValue = "missing_required_value"
	ErrorModelResource_Code_NotAuthenticated = "not_authenticated"
	ErrorModelResource_Code_NotAuthorized = "not_authorized"
	ErrorModelResource_Code_NotImplemented = "not_implemented"
	ErrorModelResource_Code_RequestBodyError = "request_body_error"
	ErrorModelResource_Code_TooManyRequests = "too_many_requests"
	ErrorModelResource_Code_UnableToPerform = "unable_to_perform"
	ErrorModelResource_Code_UnexpectedException = "unexpected_exception"
	ErrorModelResource_Code_UpdateError = "update_error"
)

// UnmarshalErrorModelResource unmarshals an instance of ErrorModelResource from the specified map of raw messages.
func UnmarshalErrorModelResource(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ErrorModelResource)
	err = core.UnmarshalPrimitive(m, "code", &obj.Code)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "message", &obj.Message)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "extra", &obj.Extra)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "more_info", &obj.MoreInfo)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// FirstPage : First page in the collection.
type FirstPage struct {
	// Link to the first page in the collection.
	Href *string `json:"href" validate:"required"`
}

// UnmarshalFirstPage unmarshals an instance of FirstPage from the specified map of raw messages.
func UnmarshalFirstPage(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(FirstPage)
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetContractTermsDocumentOptions : The GetContractTermsDocument options.
type GetContractTermsDocumentOptions struct {
	// Data product version id.
	DataProductVersionID *string `json:"data_product_version_id" validate:"required,ne="`

	// Contract terms id.
	ContractTermsID *string `json:"contract_terms_id" validate:"required,ne="`

	// Document id.
	DocumentID *string `json:"document_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetContractTermsDocumentOptions : Instantiate GetContractTermsDocumentOptions
func (*DpxV1) NewGetContractTermsDocumentOptions(dataProductVersionID string, contractTermsID string, documentID string) *GetContractTermsDocumentOptions {
	return &GetContractTermsDocumentOptions{
		DataProductVersionID: core.StringPtr(dataProductVersionID),
		ContractTermsID: core.StringPtr(contractTermsID),
		DocumentID: core.StringPtr(documentID),
	}
}

// SetDataProductVersionID : Allow user to set DataProductVersionID
func (_options *GetContractTermsDocumentOptions) SetDataProductVersionID(dataProductVersionID string) *GetContractTermsDocumentOptions {
	_options.DataProductVersionID = core.StringPtr(dataProductVersionID)
	return _options
}

// SetContractTermsID : Allow user to set ContractTermsID
func (_options *GetContractTermsDocumentOptions) SetContractTermsID(contractTermsID string) *GetContractTermsDocumentOptions {
	_options.ContractTermsID = core.StringPtr(contractTermsID)
	return _options
}

// SetDocumentID : Allow user to set DocumentID
func (_options *GetContractTermsDocumentOptions) SetDocumentID(documentID string) *GetContractTermsDocumentOptions {
	_options.DocumentID = core.StringPtr(documentID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetContractTermsDocumentOptions) SetHeaders(param map[string]string) *GetContractTermsDocumentOptions {
	options.Headers = param
	return options
}

// GetDataProductOptions : The GetDataProduct options.
type GetDataProductOptions struct {
	// Data product id.
	ID *string `json:"id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetDataProductOptions : Instantiate GetDataProductOptions
func (*DpxV1) NewGetDataProductOptions(id string) *GetDataProductOptions {
	return &GetDataProductOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (_options *GetDataProductOptions) SetID(id string) *GetDataProductOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetDataProductOptions) SetHeaders(param map[string]string) *GetDataProductOptions {
	options.Headers = param
	return options
}

// GetDataProductVersionOptions : The GetDataProductVersion options.
type GetDataProductVersionOptions struct {
	// Data product version ID.
	ID *string `json:"id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetDataProductVersionOptions : Instantiate GetDataProductVersionOptions
func (*DpxV1) NewGetDataProductVersionOptions(id string) *GetDataProductVersionOptions {
	return &GetDataProductVersionOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (_options *GetDataProductVersionOptions) SetID(id string) *GetDataProductVersionOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetDataProductVersionOptions) SetHeaders(param map[string]string) *GetDataProductVersionOptions {
	options.Headers = param
	return options
}

// GetInitializeStatusOptions : The GetInitializeStatus options.
type GetInitializeStatusOptions struct {
	// Container ID of the data product catalog. If not supplied, the data product catalog will be looked up by using the
	// uid of the default data product catalog.
	ContainerID *string `json:"container.id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetInitializeStatusOptions : Instantiate GetInitializeStatusOptions
func (*DpxV1) NewGetInitializeStatusOptions() *GetInitializeStatusOptions {
	return &GetInitializeStatusOptions{}
}

// SetContainerID : Allow user to set ContainerID
func (_options *GetInitializeStatusOptions) SetContainerID(containerID string) *GetInitializeStatusOptions {
	_options.ContainerID = core.StringPtr(containerID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetInitializeStatusOptions) SetHeaders(param map[string]string) *GetInitializeStatusOptions {
	options.Headers = param
	return options
}

// InitializeOptions : The Initialize options.
type InitializeOptions struct {
	// Data product exchange container.
	Container *ContainerReference `json:"container,omitempty"`

	// If data product exchange has already been initialized in the account, re-initialize happens forcefully when this is
	// set to true.
	Force *bool `json:"force,omitempty"`

	// Set this to true when reinitalizing the configuration.
	Reinitialize *bool `json:"reinitialize,omitempty"`

	// List of configuration options to (re-)initialize.
	Include []string `json:"include,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the InitializeOptions.Include property.
const (
	InitializeOptions_Include_DataProductSamples = "data_product_samples"
	InitializeOptions_Include_DeliveryMethods = "delivery_methods"
	InitializeOptions_Include_DomainsMultiIndustry = "domains_multi_industry"
)

// NewInitializeOptions : Instantiate InitializeOptions
func (*DpxV1) NewInitializeOptions() *InitializeOptions {
	return &InitializeOptions{}
}

// SetContainer : Allow user to set Container
func (_options *InitializeOptions) SetContainer(container *ContainerReference) *InitializeOptions {
	_options.Container = container
	return _options
}

// SetForce : Allow user to set Force
func (_options *InitializeOptions) SetForce(force bool) *InitializeOptions {
	_options.Force = core.BoolPtr(force)
	return _options
}

// SetReinitialize : Allow user to set Reinitialize
func (_options *InitializeOptions) SetReinitialize(reinitialize bool) *InitializeOptions {
	_options.Reinitialize = core.BoolPtr(reinitialize)
	return _options
}

// SetInclude : Allow user to set Include
func (_options *InitializeOptions) SetInclude(include []string) *InitializeOptions {
	_options.Include = include
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *InitializeOptions) SetHeaders(param map[string]string) *InitializeOptions {
	options.Headers = param
	return options
}

// InitializeResource : Resource defining initialization parameters.
type InitializeResource struct {
	// Data product exchange container.
	Container *ContainerReference `json:"container,omitempty"`

	// Link to monitor the status of the initialize operation.
	Href *string `json:"href,omitempty"`

	// Status of the initialize operation.
	Status *string `json:"status,omitempty"`

	// The id to trace the failed initialization operation.
	Trace *string `json:"trace,omitempty"`

	// Set of errors on the latest initialization request.
	Errors []ErrorModelResource `json:"errors,omitempty"`

	// Start time of the last initialization.
	LastStartedAt *strfmt.DateTime `json:"last_started_at,omitempty"`

	// End time of the last initialization.
	LastFinishedAt *strfmt.DateTime `json:"last_finished_at,omitempty"`

	// Initialized options.
	InitializedOptions []InitializedOption `json:"initialized_options,omitempty"`
}

// Constants associated with the InitializeResource.Status property.
// Status of the initialize operation.
const (
	InitializeResource_Status_Failed = "failed"
	InitializeResource_Status_InProgress = "in_progress"
	InitializeResource_Status_NotStarted = "not_started"
	InitializeResource_Status_Succeeded = "succeeded"
)

// UnmarshalInitializeResource unmarshals an instance of InitializeResource from the specified map of raw messages.
func UnmarshalInitializeResource(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(InitializeResource)
	err = core.UnmarshalModel(m, "container", &obj.Container, UnmarshalContainerReference)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "trace", &obj.Trace)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "errors", &obj.Errors, UnmarshalErrorModelResource)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "last_started_at", &obj.LastStartedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "last_finished_at", &obj.LastFinishedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "initialized_options", &obj.InitializedOptions, UnmarshalInitializedOption)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// InitializedOption : List of options successfully initialized.
type InitializedOption struct {
	// The name of the option.
	Name *string `json:"name,omitempty"`

	// The version of the option.
	Version *int64 `json:"version,omitempty"`
}

// UnmarshalInitializedOption unmarshals an instance of InitializedOption from the specified map of raw messages.
func UnmarshalInitializedOption(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(InitializedOption)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "version", &obj.Version)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// JSONPatchOperation : This model represents an individual patch operation to be performed on a JSON document, as defined by RFC 6902.
type JSONPatchOperation struct {
	// The operation to be performed.
	Op *string `json:"op" validate:"required"`

	// The JSON Pointer that identifies the field that is the target of the operation.
	Path *string `json:"path" validate:"required"`

	// The JSON Pointer that identifies the field that is the source of the operation.
	From *string `json:"from,omitempty"`

	// The value to be used within the operation.
	Value interface{} `json:"value,omitempty"`
}

// Constants associated with the JSONPatchOperation.Op property.
// The operation to be performed.
const (
	JSONPatchOperation_Op_Add = "add"
	JSONPatchOperation_Op_Copy = "copy"
	JSONPatchOperation_Op_Move = "move"
	JSONPatchOperation_Op_Remove = "remove"
	JSONPatchOperation_Op_Replace = "replace"
	JSONPatchOperation_Op_Test = "test"
)

// NewJSONPatchOperation : Instantiate JSONPatchOperation (Generic Model Constructor)
func (*DpxV1) NewJSONPatchOperation(op string, path string) (_model *JSONPatchOperation, err error) {
	_model = &JSONPatchOperation{
		Op: core.StringPtr(op),
		Path: core.StringPtr(path),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalJSONPatchOperation unmarshals an instance of JSONPatchOperation from the specified map of raw messages.
func UnmarshalJSONPatchOperation(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(JSONPatchOperation)
	err = core.UnmarshalPrimitive(m, "op", &obj.Op)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "path", &obj.Path)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "from", &obj.From)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ListDataProductVersionsOptions : The ListDataProductVersions options.
type ListDataProductVersionsOptions struct {
	// Filter the list of data product versions by container id.
	AssetContainerID *string `json:"asset.container.id,omitempty"`

	// Filter the list of data product versions by data product id.
	DataProduct *string `json:"data_product,omitempty"`

	// Filter the list of data product versions by state. States are: draft, available and retired.
	State *string `json:"state,omitempty"`

	// Filter the list of data product versions by version number.
	Version *string `json:"version,omitempty"`

	// Limit the number of data products in the results. The maximum limit is 200.
	Limit *int64 `json:"limit,omitempty"`

	// Start token for pagination.
	Start *string `json:"start,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the ListDataProductVersionsOptions.State property.
// Filter the list of data product versions by state. States are: draft, available and retired.
const (
	ListDataProductVersionsOptions_State_Available = "available"
	ListDataProductVersionsOptions_State_Draft = "draft"
	ListDataProductVersionsOptions_State_Retired = "retired"
)

// NewListDataProductVersionsOptions : Instantiate ListDataProductVersionsOptions
func (*DpxV1) NewListDataProductVersionsOptions() *ListDataProductVersionsOptions {
	return &ListDataProductVersionsOptions{}
}

// SetAssetContainerID : Allow user to set AssetContainerID
func (_options *ListDataProductVersionsOptions) SetAssetContainerID(assetContainerID string) *ListDataProductVersionsOptions {
	_options.AssetContainerID = core.StringPtr(assetContainerID)
	return _options
}

// SetDataProduct : Allow user to set DataProduct
func (_options *ListDataProductVersionsOptions) SetDataProduct(dataProduct string) *ListDataProductVersionsOptions {
	_options.DataProduct = core.StringPtr(dataProduct)
	return _options
}

// SetState : Allow user to set State
func (_options *ListDataProductVersionsOptions) SetState(state string) *ListDataProductVersionsOptions {
	_options.State = core.StringPtr(state)
	return _options
}

// SetVersion : Allow user to set Version
func (_options *ListDataProductVersionsOptions) SetVersion(version string) *ListDataProductVersionsOptions {
	_options.Version = core.StringPtr(version)
	return _options
}

// SetLimit : Allow user to set Limit
func (_options *ListDataProductVersionsOptions) SetLimit(limit int64) *ListDataProductVersionsOptions {
	_options.Limit = core.Int64Ptr(limit)
	return _options
}

// SetStart : Allow user to set Start
func (_options *ListDataProductVersionsOptions) SetStart(start string) *ListDataProductVersionsOptions {
	_options.Start = core.StringPtr(start)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListDataProductVersionsOptions) SetHeaders(param map[string]string) *ListDataProductVersionsOptions {
	options.Headers = param
	return options
}

// ListDataProductsOptions : The ListDataProducts options.
type ListDataProductsOptions struct {
	// Limit the number of data products in the results. The maximum limit is 200.
	Limit *int64 `json:"limit,omitempty"`

	// Start token for pagination.
	Start *string `json:"start,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListDataProductsOptions : Instantiate ListDataProductsOptions
func (*DpxV1) NewListDataProductsOptions() *ListDataProductsOptions {
	return &ListDataProductsOptions{}
}

// SetLimit : Allow user to set Limit
func (_options *ListDataProductsOptions) SetLimit(limit int64) *ListDataProductsOptions {
	_options.Limit = core.Int64Ptr(limit)
	return _options
}

// SetStart : Allow user to set Start
func (_options *ListDataProductsOptions) SetStart(start string) *ListDataProductsOptions {
	_options.Start = core.StringPtr(start)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListDataProductsOptions) SetHeaders(param map[string]string) *ListDataProductsOptions {
	options.Headers = param
	return options
}

// NextPage : Next page in the collection.
type NextPage struct {
	// Link to the next page in the collection.
	Href *string `json:"href" validate:"required"`

	// Start token for pagination to the next page in the collection.
	Start *string `json:"start" validate:"required"`
}

// UnmarshalNextPage unmarshals an instance of NextPage from the specified map of raw messages.
func UnmarshalNextPage(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(NextPage)
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "start", &obj.Start)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// UpdateContractTermsDocumentOptions : The UpdateContractTermsDocument options.
type UpdateContractTermsDocumentOptions struct {
	// Data product version id.
	DataProductVersionID *string `json:"data_product_version_id" validate:"required,ne="`

	// Contract terms id.
	ContractTermsID *string `json:"contract_terms_id" validate:"required,ne="`

	// Document id.
	DocumentID *string `json:"document_id" validate:"required,ne="`

	// A set of patch operations as defined in RFC 6902. See http://jsonpatch.com/ for more information.
	JSONPatchInstructions []JSONPatchOperation `json:"jsonPatchInstructions" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateContractTermsDocumentOptions : Instantiate UpdateContractTermsDocumentOptions
func (*DpxV1) NewUpdateContractTermsDocumentOptions(dataProductVersionID string, contractTermsID string, documentID string, jsonPatchInstructions []JSONPatchOperation) *UpdateContractTermsDocumentOptions {
	return &UpdateContractTermsDocumentOptions{
		DataProductVersionID: core.StringPtr(dataProductVersionID),
		ContractTermsID: core.StringPtr(contractTermsID),
		DocumentID: core.StringPtr(documentID),
		JSONPatchInstructions: jsonPatchInstructions,
	}
}

// SetDataProductVersionID : Allow user to set DataProductVersionID
func (_options *UpdateContractTermsDocumentOptions) SetDataProductVersionID(dataProductVersionID string) *UpdateContractTermsDocumentOptions {
	_options.DataProductVersionID = core.StringPtr(dataProductVersionID)
	return _options
}

// SetContractTermsID : Allow user to set ContractTermsID
func (_options *UpdateContractTermsDocumentOptions) SetContractTermsID(contractTermsID string) *UpdateContractTermsDocumentOptions {
	_options.ContractTermsID = core.StringPtr(contractTermsID)
	return _options
}

// SetDocumentID : Allow user to set DocumentID
func (_options *UpdateContractTermsDocumentOptions) SetDocumentID(documentID string) *UpdateContractTermsDocumentOptions {
	_options.DocumentID = core.StringPtr(documentID)
	return _options
}

// SetJSONPatchInstructions : Allow user to set JSONPatchInstructions
func (_options *UpdateContractTermsDocumentOptions) SetJSONPatchInstructions(jsonPatchInstructions []JSONPatchOperation) *UpdateContractTermsDocumentOptions {
	_options.JSONPatchInstructions = jsonPatchInstructions
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateContractTermsDocumentOptions) SetHeaders(param map[string]string) *UpdateContractTermsDocumentOptions {
	options.Headers = param
	return options
}

// UpdateDataProductVersionOptions : The UpdateDataProductVersion options.
type UpdateDataProductVersionOptions struct {
	// Data product version ID.
	ID *string `json:"id" validate:"required,ne="`

	// A set of patch operations as defined in RFC 6902. See http://jsonpatch.com/ for more information.
	JSONPatchInstructions []JSONPatchOperation `json:"jsonPatchInstructions" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateDataProductVersionOptions : Instantiate UpdateDataProductVersionOptions
func (*DpxV1) NewUpdateDataProductVersionOptions(id string, jsonPatchInstructions []JSONPatchOperation) *UpdateDataProductVersionOptions {
	return &UpdateDataProductVersionOptions{
		ID: core.StringPtr(id),
		JSONPatchInstructions: jsonPatchInstructions,
	}
}

// SetID : Allow user to set ID
func (_options *UpdateDataProductVersionOptions) SetID(id string) *UpdateDataProductVersionOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetJSONPatchInstructions : Allow user to set JSONPatchInstructions
func (_options *UpdateDataProductVersionOptions) SetJSONPatchInstructions(jsonPatchInstructions []JSONPatchOperation) *UpdateDataProductVersionOptions {
	_options.JSONPatchInstructions = jsonPatchInstructions
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateDataProductVersionOptions) SetHeaders(param map[string]string) *UpdateDataProductVersionOptions {
	options.Headers = param
	return options
}

// UseCase : UseCase struct
type UseCase struct {
	// The id of the use case associated with the data product.
	ID *string `json:"id" validate:"required"`

	// The display name of the use case associated with the data product.
	Name *string `json:"name,omitempty"`

	// Data product exchange container.
	Container *ContainerReference `json:"container,omitempty"`
}

// NewUseCase : Instantiate UseCase (Generic Model Constructor)
func (*DpxV1) NewUseCase(id string) (_model *UseCase, err error) {
	_model = &UseCase{
		ID: core.StringPtr(id),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalUseCase unmarshals an instance of UseCase from the specified map of raw messages.
func UnmarshalUseCase(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(UseCase)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "container", &obj.Container, UnmarshalContainerReference)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

//
// DataProductsPager can be used to simplify the use of the "ListDataProducts" method.
//
type DataProductsPager struct {
	hasNext bool
	options *ListDataProductsOptions
	client  *DpxV1
	pageContext struct {
		next *string
	}
}

// NewDataProductsPager returns a new DataProductsPager instance.
func (dpx *DpxV1) NewDataProductsPager(options *ListDataProductsOptions) (pager *DataProductsPager, err error) {
	if options.Start != nil && *options.Start != "" {
		err = fmt.Errorf("the 'options.Start' field should not be set")
		return
	}

	var optionsCopy ListDataProductsOptions = *options
	pager = &DataProductsPager{
		hasNext: true,
		options: &optionsCopy,
		client:  dpx,
	}
	return
}

// HasNext returns true if there are potentially more results to be retrieved.
func (pager *DataProductsPager) HasNext() bool {
	return pager.hasNext
}

// GetNextWithContext returns the next page of results using the specified Context.
func (pager *DataProductsPager) GetNextWithContext(ctx context.Context) (page []DataProduct, err error) {
	if !pager.HasNext() {
		return nil, fmt.Errorf("no more results available")
	}

	pager.options.Start = pager.pageContext.next

	result, _, err := pager.client.ListDataProductsWithContext(ctx, pager.options)
	if err != nil {
		return
	}

	var next *string
	if result.Next != nil {
		next = result.Next.Start
	}
	pager.pageContext.next = next
	pager.hasNext = (pager.pageContext.next != nil)
	page = result.DataProducts

	return
}

// GetAllWithContext returns all results by invoking GetNextWithContext() repeatedly
// until all pages of results have been retrieved.
func (pager *DataProductsPager) GetAllWithContext(ctx context.Context) (allItems []DataProduct, err error) {
	for pager.HasNext() {
		var nextPage []DataProduct
		nextPage, err = pager.GetNextWithContext(ctx)
		if err != nil {
			return
		}
		allItems = append(allItems, nextPage...)
	}
	return
}

// GetNext invokes GetNextWithContext() using context.Background() as the Context parameter.
func (pager *DataProductsPager) GetNext() (page []DataProduct, err error) {
	return pager.GetNextWithContext(context.Background())
}

// GetAll invokes GetAllWithContext() using context.Background() as the Context parameter.
func (pager *DataProductsPager) GetAll() (allItems []DataProduct, err error) {
	return pager.GetAllWithContext(context.Background())
}

//
// DataProductVersionsPager can be used to simplify the use of the "ListDataProductVersions" method.
//
type DataProductVersionsPager struct {
	hasNext bool
	options *ListDataProductVersionsOptions
	client  *DpxV1
	pageContext struct {
		next *string
	}
}

// NewDataProductVersionsPager returns a new DataProductVersionsPager instance.
func (dpx *DpxV1) NewDataProductVersionsPager(options *ListDataProductVersionsOptions) (pager *DataProductVersionsPager, err error) {
	if options.Start != nil && *options.Start != "" {
		err = fmt.Errorf("the 'options.Start' field should not be set")
		return
	}

	var optionsCopy ListDataProductVersionsOptions = *options
	pager = &DataProductVersionsPager{
		hasNext: true,
		options: &optionsCopy,
		client:  dpx,
	}
	return
}

// HasNext returns true if there are potentially more results to be retrieved.
func (pager *DataProductVersionsPager) HasNext() bool {
	return pager.hasNext
}

// GetNextWithContext returns the next page of results using the specified Context.
func (pager *DataProductVersionsPager) GetNextWithContext(ctx context.Context) (page []DataProductVersionSummary, err error) {
	if !pager.HasNext() {
		return nil, fmt.Errorf("no more results available")
	}

	pager.options.Start = pager.pageContext.next

	result, _, err := pager.client.ListDataProductVersionsWithContext(ctx, pager.options)
	if err != nil {
		return
	}

	var next *string
	if result.Next != nil {
		next = result.Next.Start
	}
	pager.pageContext.next = next
	pager.hasNext = (pager.pageContext.next != nil)
	page = result.DataProductVersions

	return
}

// GetAllWithContext returns all results by invoking GetNextWithContext() repeatedly
// until all pages of results have been retrieved.
func (pager *DataProductVersionsPager) GetAllWithContext(ctx context.Context) (allItems []DataProductVersionSummary, err error) {
	for pager.HasNext() {
		var nextPage []DataProductVersionSummary
		nextPage, err = pager.GetNextWithContext(ctx)
		if err != nil {
			return
		}
		allItems = append(allItems, nextPage...)
	}
	return
}

// GetNext invokes GetNextWithContext() using context.Background() as the Context parameter.
func (pager *DataProductVersionsPager) GetNext() (page []DataProductVersionSummary, err error) {
	return pager.GetNextWithContext(context.Background())
}

// GetAll invokes GetAllWithContext() using context.Background() as the Context parameter.
func (pager *DataProductVersionsPager) GetAll() (allItems []DataProductVersionSummary, err error) {
	return pager.GetAllWithContext(context.Background())
}
