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
 * IBM OpenAPI SDK Code Generator Version: 3.79.0-2eb6af3d-20230905-174838
 */

// Package dataproductexchangeapiservicev1 : Operations and models for the DataProductExchangeApiServiceV1 service
package dataproductexchangeapiservicev1

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/go-openapi/strfmt"
	common "github.com/IBM/data-product-exchange-go-sdk/common"
)

// DataProductExchangeApiServiceV1 : Data Product Exchange API Service
//
// API Version: 1.0.0
type DataProductExchangeApiServiceV1 struct {
	Service *core.BaseService
}

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "data_product_exchange_api_service"

// DataProductExchangeApiServiceV1Options : Service options
type DataProductExchangeApiServiceV1Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator
}

// NewDataProductExchangeApiServiceV1UsingExternalConfig : constructs an instance of DataProductExchangeApiServiceV1 with passed in options and external configuration.
func NewDataProductExchangeApiServiceV1UsingExternalConfig(options *DataProductExchangeApiServiceV1Options) (dataProductExchangeApiService *DataProductExchangeApiServiceV1, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			return
		}
	}

	dataProductExchangeApiService, err = NewDataProductExchangeApiServiceV1(options)
	if err != nil {
		return
	}

	err = dataProductExchangeApiService.Service.ConfigureService(options.ServiceName)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = dataProductExchangeApiService.Service.SetServiceURL(options.URL)
	}
	return
}

// NewDataProductExchangeApiServiceV1 : constructs an instance of DataProductExchangeApiServiceV1 with passed in options.
func NewDataProductExchangeApiServiceV1(options *DataProductExchangeApiServiceV1Options) (service *DataProductExchangeApiServiceV1, err error) {
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

	service = &DataProductExchangeApiServiceV1{
		Service: baseService,
	}

	return
}

// GetServiceURLForRegion returns the service URL to be used for the specified region
func GetServiceURLForRegion(region string) (string, error) {
	return "", fmt.Errorf("service does not support regional URLs")
}

// Clone makes a copy of "dataProductExchangeApiService" suitable for processing requests.
func (dataProductExchangeApiService *DataProductExchangeApiServiceV1) Clone() *DataProductExchangeApiServiceV1 {
	if core.IsNil(dataProductExchangeApiService) {
		return nil
	}
	clone := *dataProductExchangeApiService
	clone.Service = dataProductExchangeApiService.Service.Clone()
	return &clone
}

// SetServiceURL sets the service URL
func (dataProductExchangeApiService *DataProductExchangeApiServiceV1) SetServiceURL(url string) error {
	return dataProductExchangeApiService.Service.SetServiceURL(url)
}

// GetServiceURL returns the service URL
func (dataProductExchangeApiService *DataProductExchangeApiServiceV1) GetServiceURL() string {
	return dataProductExchangeApiService.Service.GetServiceURL()
}

// SetDefaultHeaders sets HTTP headers to be sent in every request
func (dataProductExchangeApiService *DataProductExchangeApiServiceV1) SetDefaultHeaders(headers http.Header) {
	dataProductExchangeApiService.Service.SetDefaultHeaders(headers)
}

// SetEnableGzipCompression sets the service's EnableGzipCompression field
func (dataProductExchangeApiService *DataProductExchangeApiServiceV1) SetEnableGzipCompression(enableGzip bool) {
	dataProductExchangeApiService.Service.SetEnableGzipCompression(enableGzip)
}

// GetEnableGzipCompression returns the service's EnableGzipCompression field
func (dataProductExchangeApiService *DataProductExchangeApiServiceV1) GetEnableGzipCompression() bool {
	return dataProductExchangeApiService.Service.GetEnableGzipCompression()
}

// EnableRetries enables automatic retries for requests invoked for this service instance.
// If either parameter is specified as 0, then a default value is used instead.
func (dataProductExchangeApiService *DataProductExchangeApiServiceV1) EnableRetries(maxRetries int, maxRetryInterval time.Duration) {
	dataProductExchangeApiService.Service.EnableRetries(maxRetries, maxRetryInterval)
}

// DisableRetries disables automatic retries for requests invoked for this service instance.
func (dataProductExchangeApiService *DataProductExchangeApiServiceV1) DisableRetries() {
	dataProductExchangeApiService.Service.DisableRetries()
}

// GetInitializeStatus : Get the status of resources initialization in data product exchange
// Use this API to get the status of the resource initialization in data product exchange. <br/><br/>If the data product
// catalog exists but has never been initialized, the status will be "not_started".<br/>If the data product catalog
// exists and has been or is being initialized, the response will contain the status of the last or current
// initialization.If the initialization failed, the "errors" and the "trace" fields will contain the error(s)
// encountered during the initialization and the id to trace the error(s).<br/>If the data product catalog doesn't
// exist, a HTTP 404 response will be returned.
func (dataProductExchangeApiService *DataProductExchangeApiServiceV1) GetInitializeStatus(getInitializeStatusOptions *GetInitializeStatusOptions) (result *InitializeResource, response *core.DetailedResponse, err error) {
	return dataProductExchangeApiService.GetInitializeStatusWithContext(context.Background(), getInitializeStatusOptions)
}

// GetInitializeStatusWithContext is an alternate form of the GetInitializeStatus method which supports a Context parameter
func (dataProductExchangeApiService *DataProductExchangeApiServiceV1) GetInitializeStatusWithContext(ctx context.Context, getInitializeStatusOptions *GetInitializeStatusOptions) (result *InitializeResource, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getInitializeStatusOptions, "getInitializeStatusOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dataProductExchangeApiService.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dataProductExchangeApiService.Service.Options.URL, `/data_product_exchange/v1/configuration/initialize/status`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range getInitializeStatusOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("data_product_exchange_api_service", "V1", "GetInitializeStatus")
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
	response, err = dataProductExchangeApiService.Service.Request(request, &rawResponse)
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
// data product catalog will be created.<br/>For first time initialization, it is recommended that `delivery_methods`
// and at least one domain taxonomy is included in the initialize operation.<br/><br/>If the data product exchange has
// already been initialized, you may call this API again to initialize new resources, such as new delivery methods.In
// this case, specify the default data product catalog container information.
func (dataProductExchangeApiService *DataProductExchangeApiServiceV1) Initialize(initializeOptions *InitializeOptions) (result *InitializeResource, response *core.DetailedResponse, err error) {
	return dataProductExchangeApiService.InitializeWithContext(context.Background(), initializeOptions)
}

// InitializeWithContext is an alternate form of the Initialize method which supports a Context parameter
func (dataProductExchangeApiService *DataProductExchangeApiServiceV1) InitializeWithContext(ctx context.Context, initializeOptions *InitializeOptions) (result *InitializeResource, response *core.DetailedResponse, err error) {
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
	builder.EnableGzipCompression = dataProductExchangeApiService.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dataProductExchangeApiService.Service.Options.URL, `/data_product_exchange/v1/configuration/initialize`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range initializeOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("data_product_exchange_api_service", "V1", "Initialize")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if initializeOptions.Container != nil {
		body["container"] = initializeOptions.Container
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
	response, err = dataProductExchangeApiService.Service.Request(request, &rawResponse)
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
func (dataProductExchangeApiService *DataProductExchangeApiServiceV1) GetDataProduct(getDataProductOptions *GetDataProductOptions) (result *DataProduct, response *core.DetailedResponse, err error) {
	return dataProductExchangeApiService.GetDataProductWithContext(context.Background(), getDataProductOptions)
}

// GetDataProductWithContext is an alternate form of the GetDataProduct method which supports a Context parameter
func (dataProductExchangeApiService *DataProductExchangeApiServiceV1) GetDataProductWithContext(ctx context.Context, getDataProductOptions *GetDataProductOptions) (result *DataProduct, response *core.DetailedResponse, err error) {
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
	builder.EnableGzipCompression = dataProductExchangeApiService.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dataProductExchangeApiService.Service.Options.URL, `/data_product_exchange/v1/data_products/{id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getDataProductOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("data_product_exchange_api_service", "V1", "GetDataProduct")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dataProductExchangeApiService.Service.Request(request, &rawResponse)
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
func (dataProductExchangeApiService *DataProductExchangeApiServiceV1) ListDataProducts(listDataProductsOptions *ListDataProductsOptions) (result *DataProductCollection, response *core.DetailedResponse, err error) {
	return dataProductExchangeApiService.ListDataProductsWithContext(context.Background(), listDataProductsOptions)
}

// ListDataProductsWithContext is an alternate form of the ListDataProducts method which supports a Context parameter
func (dataProductExchangeApiService *DataProductExchangeApiServiceV1) ListDataProductsWithContext(ctx context.Context, listDataProductsOptions *ListDataProductsOptions) (result *DataProductCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listDataProductsOptions, "listDataProductsOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dataProductExchangeApiService.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dataProductExchangeApiService.Service.Options.URL, `/data_product_exchange/v1/data_products`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range listDataProductsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("data_product_exchange_api_service", "V1", "ListDataProducts")
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
	response, err = dataProductExchangeApiService.Service.Request(request, &rawResponse)
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
func (dataProductExchangeApiService *DataProductExchangeApiServiceV1) ListDataProductVersions(listDataProductVersionsOptions *ListDataProductVersionsOptions) (result *DataProductVersionCollection, response *core.DetailedResponse, err error) {
	return dataProductExchangeApiService.ListDataProductVersionsWithContext(context.Background(), listDataProductVersionsOptions)
}

// ListDataProductVersionsWithContext is an alternate form of the ListDataProductVersions method which supports a Context parameter
func (dataProductExchangeApiService *DataProductExchangeApiServiceV1) ListDataProductVersionsWithContext(ctx context.Context, listDataProductVersionsOptions *ListDataProductVersionsOptions) (result *DataProductVersionCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listDataProductVersionsOptions, "listDataProductVersionsOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dataProductExchangeApiService.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dataProductExchangeApiService.Service.Options.URL, `/data_product_exchange/v1/data_product_versions`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range listDataProductVersionsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("data_product_exchange_api_service", "V1", "ListDataProductVersions")
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
	response, err = dataProductExchangeApiService.Service.Request(request, &rawResponse)
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
func (dataProductExchangeApiService *DataProductExchangeApiServiceV1) CreateDataProductVersion(createDataProductVersionOptions *CreateDataProductVersionOptions) (result *DataProductVersion, response *core.DetailedResponse, err error) {
	return dataProductExchangeApiService.CreateDataProductVersionWithContext(context.Background(), createDataProductVersionOptions)
}

// CreateDataProductVersionWithContext is an alternate form of the CreateDataProductVersion method which supports a Context parameter
func (dataProductExchangeApiService *DataProductExchangeApiServiceV1) CreateDataProductVersionWithContext(ctx context.Context, createDataProductVersionOptions *CreateDataProductVersionOptions) (result *DataProductVersion, response *core.DetailedResponse, err error) {
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
	builder.EnableGzipCompression = dataProductExchangeApiService.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dataProductExchangeApiService.Service.Options.URL, `/data_product_exchange/v1/data_product_versions`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range createDataProductVersionOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("data_product_exchange_api_service", "V1", "CreateDataProductVersion")
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
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dataProductExchangeApiService.Service.Request(request, &rawResponse)
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
func (dataProductExchangeApiService *DataProductExchangeApiServiceV1) GetDataProductVersion(getDataProductVersionOptions *GetDataProductVersionOptions) (result *DataProductVersion, response *core.DetailedResponse, err error) {
	return dataProductExchangeApiService.GetDataProductVersionWithContext(context.Background(), getDataProductVersionOptions)
}

// GetDataProductVersionWithContext is an alternate form of the GetDataProductVersion method which supports a Context parameter
func (dataProductExchangeApiService *DataProductExchangeApiServiceV1) GetDataProductVersionWithContext(ctx context.Context, getDataProductVersionOptions *GetDataProductVersionOptions) (result *DataProductVersion, response *core.DetailedResponse, err error) {
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
	builder.EnableGzipCompression = dataProductExchangeApiService.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dataProductExchangeApiService.Service.Options.URL, `/data_product_exchange/v1/data_product_versions/{id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getDataProductVersionOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("data_product_exchange_api_service", "V1", "GetDataProductVersion")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dataProductExchangeApiService.Service.Request(request, &rawResponse)
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
func (dataProductExchangeApiService *DataProductExchangeApiServiceV1) DeleteDataProductVersion(deleteDataProductVersionOptions *DeleteDataProductVersionOptions) (response *core.DetailedResponse, err error) {
	return dataProductExchangeApiService.DeleteDataProductVersionWithContext(context.Background(), deleteDataProductVersionOptions)
}

// DeleteDataProductVersionWithContext is an alternate form of the DeleteDataProductVersion method which supports a Context parameter
func (dataProductExchangeApiService *DataProductExchangeApiServiceV1) DeleteDataProductVersionWithContext(ctx context.Context, deleteDataProductVersionOptions *DeleteDataProductVersionOptions) (response *core.DetailedResponse, err error) {
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
	builder.EnableGzipCompression = dataProductExchangeApiService.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dataProductExchangeApiService.Service.Options.URL, `/data_product_exchange/v1/data_product_versions/{id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteDataProductVersionOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("data_product_exchange_api_service", "V1", "DeleteDataProductVersion")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = dataProductExchangeApiService.Service.Request(request, nil)

	return
}

// UpdateDataProductVersion : Update the data product version identified by ID
// Use this API to update the properties of a data product version identified by a valid ID.<br/><br/>Specify patch
// operations using http://jsonpatch.com/ syntax.<br/><br/>Supported patch operations include:<br/><br/>- Update the
// properties of a data product<br/><br/>- Add/Remove parts from a data product<br/><br/>- Add/Remove use cases from a
// data product<br/><br/>- Update the data product state<br/><br/>.
func (dataProductExchangeApiService *DataProductExchangeApiServiceV1) UpdateDataProductVersion(updateDataProductVersionOptions *UpdateDataProductVersionOptions) (result *DataProductVersion, response *core.DetailedResponse, err error) {
	return dataProductExchangeApiService.UpdateDataProductVersionWithContext(context.Background(), updateDataProductVersionOptions)
}

// UpdateDataProductVersionWithContext is an alternate form of the UpdateDataProductVersion method which supports a Context parameter
func (dataProductExchangeApiService *DataProductExchangeApiServiceV1) UpdateDataProductVersionWithContext(ctx context.Context, updateDataProductVersionOptions *UpdateDataProductVersionOptions) (result *DataProductVersion, response *core.DetailedResponse, err error) {
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
	builder.EnableGzipCompression = dataProductExchangeApiService.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dataProductExchangeApiService.Service.Options.URL, `/data_product_exchange/v1/data_product_versions/{id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateDataProductVersionOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("data_product_exchange_api_service", "V1", "UpdateDataProductVersion")
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
	response, err = dataProductExchangeApiService.Service.Request(request, &rawResponse)
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

// DeliverDataProductVersion : Deliver a data product identified by id
// Deliver a data product version identified by id.
func (dataProductExchangeApiService *DataProductExchangeApiServiceV1) DeliverDataProductVersion(deliverDataProductVersionOptions *DeliverDataProductVersionOptions) (result *DeliveryResource, response *core.DetailedResponse, err error) {
	return dataProductExchangeApiService.DeliverDataProductVersionWithContext(context.Background(), deliverDataProductVersionOptions)
}

// DeliverDataProductVersionWithContext is an alternate form of the DeliverDataProductVersion method which supports a Context parameter
func (dataProductExchangeApiService *DataProductExchangeApiServiceV1) DeliverDataProductVersionWithContext(ctx context.Context, deliverDataProductVersionOptions *DeliverDataProductVersionOptions) (result *DeliveryResource, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deliverDataProductVersionOptions, "deliverDataProductVersionOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deliverDataProductVersionOptions, "deliverDataProductVersionOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *deliverDataProductVersionOptions.ID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dataProductExchangeApiService.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dataProductExchangeApiService.Service.Options.URL, `/data_product_exchange/v1/data_product_versions/{id}/deliver`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deliverDataProductVersionOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("data_product_exchange_api_service", "V1", "DeliverDataProductVersion")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if deliverDataProductVersionOptions.Order != nil {
		body["order"] = deliverDataProductVersionOptions.Order
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
	response, err = dataProductExchangeApiService.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDeliveryResource)
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
func (*DataProductExchangeApiServiceV1) NewAssetPartReference(id string, container *ContainerReference) (_model *AssetPartReference, err error) {
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

// AssetReference : The asset referenced by the data product version.
type AssetReference struct {
	// The unique identifier of the asset.
	ID *string `json:"id" validate:"required"`

	// Data product exchange container.
	Container *ContainerReference `json:"container" validate:"required"`
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
func (*DataProductExchangeApiServiceV1) NewContainerReference(id string) (_model *ContainerReference, err error) {
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
func (*DataProductExchangeApiServiceV1) NewCreateDataProductVersionOptions(container *ContainerReference) *CreateDataProductVersionOptions {
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

// DataProductIdentity : Data product identifier.
type DataProductIdentity struct {
	// Data product identifier.
	ID *string `json:"id" validate:"required"`
}

// NewDataProductIdentity : Instantiate DataProductIdentity (Generic Model Constructor)
func (*DataProductExchangeApiServiceV1) NewDataProductIdentity(id string) (_model *DataProductIdentity, err error) {
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
func (*DataProductExchangeApiServiceV1) NewDataProductPart(asset *AssetPartReference) (_model *DataProductPart, err error) {
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

	// The asset referenced by the data product version.
	Asset *AssetReference `json:"asset" validate:"required"`

	// Tags on the data product.
	Tags []string `json:"tags,omitempty"`

	// A list of use cases associated with the data product version.
	UseCases []UseCase `json:"use_cases" validate:"required"`

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

func (*DataProductExchangeApiServiceV1) NewDataProductVersionPatch(dataProductVersion *DataProductVersion) (_patch []JSONPatchOperation) {
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

	// The asset referenced by the data product version.
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

// DeleteDataProductVersionOptions : The DeleteDataProductVersion options.
type DeleteDataProductVersionOptions struct {
	// Data product version ID.
	ID *string `json:"id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteDataProductVersionOptions : Instantiate DeleteDataProductVersionOptions
func (*DataProductExchangeApiServiceV1) NewDeleteDataProductVersionOptions(id string) *DeleteDataProductVersionOptions {
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

// DeliverDataProductVersionOptions : The DeliverDataProductVersion options.
type DeliverDataProductVersionOptions struct {
	// Data product version id.
	ID *string `json:"id" validate:"required,ne="`

	// The order for the data product that should be delivered as part of this delivery operation.
	Order *OrderReference `json:"order,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeliverDataProductVersionOptions : Instantiate DeliverDataProductVersionOptions
func (*DataProductExchangeApiServiceV1) NewDeliverDataProductVersionOptions(id string) *DeliverDataProductVersionOptions {
	return &DeliverDataProductVersionOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (_options *DeliverDataProductVersionOptions) SetID(id string) *DeliverDataProductVersionOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetOrder : Allow user to set Order
func (_options *DeliverDataProductVersionOptions) SetOrder(order *OrderReference) *DeliverDataProductVersionOptions {
	_options.Order = order
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeliverDataProductVersionOptions) SetHeaders(param map[string]string) *DeliverDataProductVersionOptions {
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
func (*DataProductExchangeApiServiceV1) NewDeliveryMethod(id string, container *ContainerReference) (_model *DeliveryMethod, err error) {
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

// DeliveryResource : DeliveryResource struct
type DeliveryResource struct {
	// Status of the deliver operation.
	Status *string `json:"status" validate:"required"`

	// Link to monitor the status of the deliver operation.
	Href *string `json:"href,omitempty"`
}

// Constants associated with the DeliveryResource.Status property.
// Status of the deliver operation.
const (
	DeliveryResource_Status_Delivered = "delivered"
	DeliveryResource_Status_Failed = "failed"
	DeliveryResource_Status_NotStarted = "not_started"
	DeliveryResource_Status_Received = "received"
	DeliveryResource_Status_Succeeded = "succeeded"
)

// UnmarshalDeliveryResource unmarshals an instance of DeliveryResource from the specified map of raw messages.
func UnmarshalDeliveryResource(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DeliveryResource)
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
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
	Name *string `json:"name" validate:"required"`

	// Data product exchange container.
	Container *ContainerReference `json:"container,omitempty"`
}

// NewDomain : Instantiate Domain (Generic Model Constructor)
func (*DataProductExchangeApiServiceV1) NewDomain(id string, name string) (_model *Domain, err error) {
	_model = &Domain{
		ID: core.StringPtr(id),
		Name: core.StringPtr(name),
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

// ErrorModel : ErrorModel struct
type ErrorModel struct {
	Code *string `json:"code,omitempty"`

	Target *ErrorTargetModel `json:"target,omitempty"`

	Message *string `json:"message,omitempty"`

	MoreInfo *string `json:"more_info,omitempty"`
}

// UnmarshalErrorModel unmarshals an instance of ErrorModel from the specified map of raw messages.
func UnmarshalErrorModel(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ErrorModel)
	err = core.UnmarshalPrimitive(m, "code", &obj.Code)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "target", &obj.Target, UnmarshalErrorTargetModel)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "message", &obj.Message)
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

// ErrorTargetModel : ErrorTargetModel struct
type ErrorTargetModel struct {
	Type *string `json:"type,omitempty"`

	Name *string `json:"name,omitempty"`
}

// Constants associated with the ErrorTargetModel.Type property.
const (
	ErrorTargetModel_Type_Field = "field"
	ErrorTargetModel_Type_Header = "header"
	ErrorTargetModel_Type_Parameter = "parameter"
)

// UnmarshalErrorTargetModel unmarshals an instance of ErrorTargetModel from the specified map of raw messages.
func UnmarshalErrorTargetModel(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ErrorTargetModel)
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
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

// GetDataProductOptions : The GetDataProduct options.
type GetDataProductOptions struct {
	// Data product id.
	ID *string `json:"id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetDataProductOptions : Instantiate GetDataProductOptions
func (*DataProductExchangeApiServiceV1) NewGetDataProductOptions(id string) *GetDataProductOptions {
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
func (*DataProductExchangeApiServiceV1) NewGetDataProductVersionOptions(id string) *GetDataProductVersionOptions {
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
func (*DataProductExchangeApiServiceV1) NewGetInitializeStatusOptions() *GetInitializeStatusOptions {
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

	// List of configuration options to initialize.
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
func (*DataProductExchangeApiServiceV1) NewInitializeOptions() *InitializeOptions {
	return &InitializeOptions{}
}

// SetContainer : Allow user to set Container
func (_options *InitializeOptions) SetContainer(container *ContainerReference) *InitializeOptions {
	_options.Container = container
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

// InitializeResource : InitializeResource struct
type InitializeResource struct {
	// Data product exchange container.
	Container *ContainerReference `json:"container,omitempty"`

	// Link to monitor the status of the initialize operation.
	Href *string `json:"href,omitempty"`

	// Status of the initialize operation.
	Status *string `json:"status,omitempty"`

	// The id to trace the failed initialization operation.
	Trace *string `json:"trace,omitempty"`

	// The error(s) encountered in the initialization operation.
	Errors []ErrorModel `json:"errors,omitempty"`

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
	err = core.UnmarshalModel(m, "errors", &obj.Errors, UnmarshalErrorModel)
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

// InitializedOption : Initialized options.
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

// ItemReference : ItemReference struct
type ItemReference struct {
	// The unique identifier of an item on an asset list representing a data product order.
	ID *string `json:"id" validate:"required"`
}

// NewItemReference : Instantiate ItemReference (Generic Model Constructor)
func (*DataProductExchangeApiServiceV1) NewItemReference(id string) (_model *ItemReference, err error) {
	_model = &ItemReference{
		ID: core.StringPtr(id),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalItemReference unmarshals an instance of ItemReference from the specified map of raw messages.
func UnmarshalItemReference(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ItemReference)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
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
func (*DataProductExchangeApiServiceV1) NewJSONPatchOperation(op string, path string) (_model *JSONPatchOperation, err error) {
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
func (*DataProductExchangeApiServiceV1) NewListDataProductVersionsOptions() *ListDataProductVersionsOptions {
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
func (*DataProductExchangeApiServiceV1) NewListDataProductsOptions() *ListDataProductsOptions {
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

// OrderReference : The order for the data product that should be delivered as part of this delivery operation.
type OrderReference struct {
	// The unique identifier of the asset list representing a data product order.
	ID *string `json:"id" validate:"required"`

	// The list of items to be delivered as part of this operation. This list can be a subset of items belonging to this
	// order. All items specified must belong to this order.
	Items []ItemReference `json:"items,omitempty"`
}

// NewOrderReference : Instantiate OrderReference (Generic Model Constructor)
func (*DataProductExchangeApiServiceV1) NewOrderReference(id string) (_model *OrderReference, err error) {
	_model = &OrderReference{
		ID: core.StringPtr(id),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalOrderReference unmarshals an instance of OrderReference from the specified map of raw messages.
func UnmarshalOrderReference(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(OrderReference)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "items", &obj.Items, UnmarshalItemReference)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
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
func (*DataProductExchangeApiServiceV1) NewUpdateDataProductVersionOptions(id string, jsonPatchInstructions []JSONPatchOperation) *UpdateDataProductVersionOptions {
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
	Name *string `json:"name" validate:"required"`

	// Data product exchange container.
	Container *ContainerReference `json:"container,omitempty"`
}

// NewUseCase : Instantiate UseCase (Generic Model Constructor)
func (*DataProductExchangeApiServiceV1) NewUseCase(id string, name string) (_model *UseCase, err error) {
	_model = &UseCase{
		ID: core.StringPtr(id),
		Name: core.StringPtr(name),
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
	client  *DataProductExchangeApiServiceV1
	pageContext struct {
		next *string
	}
}

// NewDataProductsPager returns a new DataProductsPager instance.
func (dataProductExchangeApiService *DataProductExchangeApiServiceV1) NewDataProductsPager(options *ListDataProductsOptions) (pager *DataProductsPager, err error) {
	if options.Start != nil && *options.Start != "" {
		err = fmt.Errorf("the 'options.Start' field should not be set")
		return
	}

	var optionsCopy ListDataProductsOptions = *options
	pager = &DataProductsPager{
		hasNext: true,
		options: &optionsCopy,
		client:  dataProductExchangeApiService,
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
	client  *DataProductExchangeApiServiceV1
	pageContext struct {
		next *string
	}
}

// NewDataProductVersionsPager returns a new DataProductVersionsPager instance.
func (dataProductExchangeApiService *DataProductExchangeApiServiceV1) NewDataProductVersionsPager(options *ListDataProductVersionsOptions) (pager *DataProductVersionsPager, err error) {
	if options.Start != nil && *options.Start != "" {
		err = fmt.Errorf("the 'options.Start' field should not be set")
		return
	}

	var optionsCopy ListDataProductVersionsOptions = *options
	pager = &DataProductVersionsPager{
		hasNext: true,
		options: &optionsCopy,
		client:  dataProductExchangeApiService,
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
