// +build examples

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

package dataproductexchangeapiservicev1_test

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/IBM/data-product-exchange-go-sdk/dataproductexchangeapiservicev1"
)

//
// This file provides an example of how to use the Data Product Exchange API Service service.
//
// The following configuration properties are assumed to be defined:
// DATA_PRODUCT_EXCHANGE_API_SERVICE_URL=<service base url>
// DATA_PRODUCT_EXCHANGE_API_SERVICE_AUTH_TYPE=iam
// DATA_PRODUCT_EXCHANGE_API_SERVICE_APIKEY=<IAM apikey>
// DATA_PRODUCT_EXCHANGE_API_SERVICE_AUTH_URL=<IAM token service base URL - omit this if using the production environment>
//
// These configuration properties can be exported as environment variables, or stored
// in a configuration file and then:
// export IBM_CREDENTIALS_FILE=<name of configuration file>
//
var _ = Describe(`DataProductExchangeApiServiceV1 Examples Tests`, func() {

	const externalConfigFile = "../data_product_exchange_api_service_v1.env"

	var (
		dataProductExchangeApiServiceService *dataproductexchangeapiservicev1.DataProductExchangeApiServiceV1
		config       map[string]string

		// Variables to hold link values
		createDataProductVersionByCatalogIdLink string
		deleteDataProductVersionByUserIdLink string
		deliverDataProductVersionByUserIdLink string
		getDataProductByUserIdLink string
		getDataProductVersionByUserIdLink string
		getListOfDataProductByCatalogIdLink string
		getStatusByCatalogIdLink string
		updateDataProductVersionByUserIdLink string
	)

	var shouldSkipTest = func() {
		Skip("External configuration is not available, skipping examples...")
	}

	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			var err error
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping examples: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(dataproductexchangeapiservicev1.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping examples: " + err.Error())
			} else if len(config) == 0 {
				Skip("Unable to load service properties, skipping examples")
			}

			shouldSkipTest = func() {}
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {
			var err error

			// begin-common

			dataProductExchangeApiServiceServiceOptions := &dataproductexchangeapiservicev1.DataProductExchangeApiServiceV1Options{}

			dataProductExchangeApiServiceService, err = dataproductexchangeapiservicev1.NewDataProductExchangeApiServiceV1UsingExternalConfig(dataProductExchangeApiServiceServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(dataProductExchangeApiServiceService).ToNot(BeNil())
		})
	})

	Describe(`DataProductExchangeApiServiceV1 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`Initialize request example`, func() {
			fmt.Println("\nInitialize() result:")
			// begin-initialize

			initializeOptions := dataProductExchangeApiServiceService.NewInitializeOptions()
			initializeOptions.SetInclude([]string{"delivery_methods", "data_product_samples", "domains_multi_industry"})

			initializeResource, response, err := dataProductExchangeApiServiceService.Initialize(initializeOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(initializeResource, "", "  ")
			fmt.Println(string(b))

			// end-initialize

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(initializeResource).ToNot(BeNil())

			createDataProductVersionByCatalogIdLink = *initializeResource.Container.ID
			fmt.Fprintf(GinkgoWriter, "Saved createDataProductVersionByCatalogIdLink value: %v\n", createDataProductVersionByCatalogIdLink)
			getStatusByCatalogIdLink = *initializeResource.Container.ID
			fmt.Fprintf(GinkgoWriter, "Saved getStatusByCatalogIdLink value: %v\n", getStatusByCatalogIdLink)
			getListOfDataProductByCatalogIdLink = *initializeResource.Container.ID
			fmt.Fprintf(GinkgoWriter, "Saved getListOfDataProductByCatalogIdLink value: %v\n", getListOfDataProductByCatalogIdLink)
		})
		It(`CreateDataProductVersion request example`, func() {
			fmt.Println("\nCreateDataProductVersion() result:")
			// begin-create_data_product_version

			containerReferenceModel := &dataproductexchangeapiservicev1.ContainerReference{
				ID: core.StringPtr("d29c42eb-7100-4b7a-8257-c196dbcca1cd"),
			}

			createDataProductVersionOptions := dataProductExchangeApiServiceService.NewCreateDataProductVersionOptions(
				containerReferenceModel,
			)
			createDataProductVersionOptions.SetName("My New Data Product")

			dataProductVersion, response, err := dataProductExchangeApiServiceService.CreateDataProductVersion(createDataProductVersionOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(dataProductVersion, "", "  ")
			fmt.Println(string(b))

			// end-create_data_product_version

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(dataProductVersion).ToNot(BeNil())

			getDataProductVersionByUserIdLink = *dataProductVersion.ID
			fmt.Fprintf(GinkgoWriter, "Saved getDataProductVersionByUserIdLink value: %v\n", getDataProductVersionByUserIdLink)
			updateDataProductVersionByUserIdLink = *dataProductVersion.ID
			fmt.Fprintf(GinkgoWriter, "Saved updateDataProductVersionByUserIdLink value: %v\n", updateDataProductVersionByUserIdLink)
			deleteDataProductVersionByUserIdLink = *dataProductVersion.ID
			fmt.Fprintf(GinkgoWriter, "Saved deleteDataProductVersionByUserIdLink value: %v\n", deleteDataProductVersionByUserIdLink)
			getDataProductByUserIdLink = *dataProductVersion.DataProduct.ID
			fmt.Fprintf(GinkgoWriter, "Saved getDataProductByUserIdLink value: %v\n", getDataProductByUserIdLink)
			deliverDataProductVersionByUserIdLink = *dataProductVersion.ID
			fmt.Fprintf(GinkgoWriter, "Saved deliverDataProductVersionByUserIdLink value: %v\n", deliverDataProductVersionByUserIdLink)
		})
		It(`GetInitializeStatus request example`, func() {
			fmt.Println("\nGetInitializeStatus() result:")
			// begin-get_initialize_status

			getInitializeStatusOptions := dataProductExchangeApiServiceService.NewGetInitializeStatusOptions()

			initializeResource, response, err := dataProductExchangeApiServiceService.GetInitializeStatus(getInitializeStatusOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(initializeResource, "", "  ")
			fmt.Println(string(b))

			// end-get_initialize_status

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(initializeResource).ToNot(BeNil())
		})
		It(`GetDataProduct request example`, func() {
			fmt.Println("\nGetDataProduct() result:")
			// begin-get_data_product

			getDataProductOptions := dataProductExchangeApiServiceService.NewGetDataProductOptions(
				getDataProductByUserIdLink,
			)

			dataProduct, response, err := dataProductExchangeApiServiceService.GetDataProduct(getDataProductOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(dataProduct, "", "  ")
			fmt.Println(string(b))

			// end-get_data_product

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dataProduct).ToNot(BeNil())
		})
		It(`ListDataProducts request example`, func() {
			fmt.Println("\nListDataProducts() result:")
			// begin-list_data_products
			listDataProductsOptions := &dataproductexchangeapiservicev1.ListDataProductsOptions{
				Limit: core.Int64Ptr(int64(10)),
			}

			pager, err := dataProductExchangeApiServiceService.NewDataProductsPager(listDataProductsOptions)
			if err != nil {
				panic(err)
			}

			var allResults []dataproductexchangeapiservicev1.DataProduct
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				if err != nil {
					panic(err)
				}
				allResults = append(allResults, nextPage...)
			}
			b, _ := json.MarshalIndent(allResults, "", "  ")
			fmt.Println(string(b))
			// end-list_data_products
		})
		It(`ListDataProductVersions request example`, func() {
			fmt.Println("\nListDataProductVersions() result:")
			// begin-list_data_product_versions
			listDataProductVersionsOptions := &dataproductexchangeapiservicev1.ListDataProductVersionsOptions{
				AssetContainerID: &getListOfDataProductByCatalogIdLink,
				DataProduct: core.StringPtr("testString"),
				State: core.StringPtr("draft"),
				Version: core.StringPtr("testString"),
				Limit: core.Int64Ptr(int64(10)),
			}

			pager, err := dataProductExchangeApiServiceService.NewDataProductVersionsPager(listDataProductVersionsOptions)
			if err != nil {
				panic(err)
			}

			var allResults []dataproductexchangeapiservicev1.DataProductVersionSummary
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				if err != nil {
					panic(err)
				}
				allResults = append(allResults, nextPage...)
			}
			b, _ := json.MarshalIndent(allResults, "", "  ")
			fmt.Println(string(b))
			// end-list_data_product_versions
		})
		It(`GetDataProductVersion request example`, func() {
			fmt.Println("\nGetDataProductVersion() result:")
			// begin-get_data_product_version

			getDataProductVersionOptions := dataProductExchangeApiServiceService.NewGetDataProductVersionOptions(
				getDataProductVersionByUserIdLink,
			)

			dataProductVersion, response, err := dataProductExchangeApiServiceService.GetDataProductVersion(getDataProductVersionOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(dataProductVersion, "", "  ")
			fmt.Println(string(b))

			// end-get_data_product_version

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dataProductVersion).ToNot(BeNil())
		})
		It(`UpdateDataProductVersion request example`, func() {
			fmt.Println("\nUpdateDataProductVersion() result:")
			// begin-update_data_product_version

			jsonPatchOperationModel := &dataproductexchangeapiservicev1.JSONPatchOperation{
				Op: core.StringPtr("add"),
				Path: core.StringPtr("testString"),
			}

			updateDataProductVersionOptions := dataProductExchangeApiServiceService.NewUpdateDataProductVersionOptions(
				updateDataProductVersionByUserIdLink,
				[]dataproductexchangeapiservicev1.JSONPatchOperation{*jsonPatchOperationModel},
			)

			dataProductVersion, response, err := dataProductExchangeApiServiceService.UpdateDataProductVersion(updateDataProductVersionOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(dataProductVersion, "", "  ")
			fmt.Println(string(b))

			// end-update_data_product_version

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dataProductVersion).ToNot(BeNil())
		})
		It(`DeliverDataProductVersion request example`, func() {
			fmt.Println("\nDeliverDataProductVersion() result:")
			// begin-deliver_data_product_version

			itemReferenceModel := &dataproductexchangeapiservicev1.ItemReference{
				ID: core.StringPtr("d29c42eb-7100-4b7a-8257-c196dbcca1cd"),
			}

			orderReferenceModel := &dataproductexchangeapiservicev1.OrderReference{
				ID: core.StringPtr("4705e047-1808-459a-805f-d5d13c947637"),
				Items: []dataproductexchangeapiservicev1.ItemReference{*itemReferenceModel},
			}

			deliverDataProductVersionOptions := dataProductExchangeApiServiceService.NewDeliverDataProductVersionOptions(
				deliverDataProductVersionByUserIdLink,
			)
			deliverDataProductVersionOptions.SetOrder(orderReferenceModel)

			deliveryResource, response, err := dataProductExchangeApiServiceService.DeliverDataProductVersion(deliverDataProductVersionOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(deliveryResource, "", "  ")
			fmt.Println(string(b))

			// end-deliver_data_product_version

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(deliveryResource).ToNot(BeNil())
		})
		It(`DeleteDataProductVersion request example`, func() {
			// begin-delete_data_product_version

			deleteDataProductVersionOptions := dataProductExchangeApiServiceService.NewDeleteDataProductVersionOptions(
				deleteDataProductVersionByUserIdLink,
			)

			response, err := dataProductExchangeApiServiceService.DeleteDataProductVersion(deleteDataProductVersionOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteDataProductVersion(): %d\n", response.StatusCode)
			}

			// end-delete_data_product_version

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})
})
