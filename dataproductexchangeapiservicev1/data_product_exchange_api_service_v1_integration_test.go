// +build integration

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
	"fmt"
	"log"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/IBM/data-product-exchange-go-sdk/dataproductexchangeapiservicev1"
)

/**
 * This file contains an integration test for the dataproductexchangeapiservicev1 package.
 *
 * Notes:
 *
 * The integration test will automatically skip tests if the required config file is not available.
 */

var _ = Describe(`DataProductExchangeApiServiceV1 Integration Tests`, func() {
	const externalConfigFile = "../data_product_exchange_api_service_v1.env"

	var (
		err          error
		dataProductExchangeApiServiceService *dataproductexchangeapiservicev1.DataProductExchangeApiServiceV1
		serviceURL   string
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
		Skip("External configuration is not available, skipping tests...")
	}

	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping tests: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(dataproductexchangeapiservicev1.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}
			serviceURL = config["URL"]
			if serviceURL == "" {
				Skip("Unable to load service URL configuration property, skipping tests")
			}

			fmt.Fprintf(GinkgoWriter, "Service URL: %v\n", serviceURL)
			shouldSkipTest = func() {}
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {
			dataProductExchangeApiServiceServiceOptions := &dataproductexchangeapiservicev1.DataProductExchangeApiServiceV1Options{}

			dataProductExchangeApiServiceService, err = dataproductexchangeapiservicev1.NewDataProductExchangeApiServiceV1UsingExternalConfig(dataProductExchangeApiServiceServiceOptions)
			Expect(err).To(BeNil())
			Expect(dataProductExchangeApiServiceService).ToNot(BeNil())
			Expect(dataProductExchangeApiServiceService.Service.Options.URL).To(Equal(serviceURL))

			core.SetLogger(core.NewLogger(core.LevelDebug, log.New(GinkgoWriter, "", log.LstdFlags), log.New(GinkgoWriter, "", log.LstdFlags)))
			dataProductExchangeApiServiceService.EnableRetries(4, 30*time.Second)
		})
	})

	Describe(`Initialize - Initialize resources in a data product exchange`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`Initialize(initializeOptions *InitializeOptions)`, func() {
			// containerReferenceModel := &dataproductexchangeapiservicev1.ContainerReference{
			// 	ID: core.StringPtr("d29c42eb-7100-4b7a-8257-c196dbcca1cd"),
			// 	Type: core.StringPtr("catalog"),
			// }

			initializeOptions := &dataproductexchangeapiservicev1.InitializeOptions{
				Container: nil,
				Include: []string{"delivery_methods", "data_product_samples", "domains_multi_industry"},
			}

			initializeResource, response, err := dataProductExchangeApiServiceService.Initialize(initializeOptions)
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
	})

	Describe(`CreateDataProductVersion - Create a new data product version`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateDataProductVersion(createDataProductVersionOptions *CreateDataProductVersionOptions)`, func() {
			containerReferenceModel := &dataproductexchangeapiservicev1.ContainerReference{
				ID: core.StringPtr(createDataProductVersionByCatalogIdLink),
				Type: core.StringPtr("catalog"),
			}

			// dataProductIdentityModel := &dataproductexchangeapiservicev1.DataProductIdentity{
			// 	ID: core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e"),
			// }

			// useCaseModel := &dataproductexchangeapiservicev1.UseCase{
			// 	ID: core.StringPtr("testString"),
			// 	Name: core.StringPtr("testString"),
			// 	Container: containerReferenceModel,
			// }

			// domainModel := &dataproductexchangeapiservicev1.Domain{
			// 	ID: core.StringPtr("testString"),
			// 	Name: core.StringPtr("testString"),
			// 	Container: containerReferenceModel,
			// }

			// assetPartReferenceModel := &dataproductexchangeapiservicev1.AssetPartReference{
			// 	ID: core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002"),
			// 	Container: containerReferenceModel,
			// 	Type: core.StringPtr("data_asset"),
			// }

			// deliveryMethodModel := &dataproductexchangeapiservicev1.DeliveryMethod{
			// 	ID: core.StringPtr("09cf5fcc-cb9d-4995-a8e4-16517b25229f"),
			// 	Container: containerReferenceModel,
			// }

			// dataProductPartModel := &dataproductexchangeapiservicev1.DataProductPart{
			// 	Asset: assetPartReferenceModel,
			// 	Revision: core.Int64Ptr(int64(1)),
			// 	UpdatedAt: CreateMockDateTime("2023-07-01T22:22:34.876Z"),
			// 	DeliveryMethods: []dataproductexchangeapiservicev1.DeliveryMethod{*deliveryMethodModel},
			// }

			createDataProductVersionOptions := &dataproductexchangeapiservicev1.CreateDataProductVersionOptions{
				Container: containerReferenceModel,
				// Version: core.StringPtr("testString"),
				// State: core.StringPtr("draft"),
				// DataProduct: dataProductIdentityModel,
				Name: core.StringPtr("My New Data Product"),
				Description: core.StringPtr("testString"),
				// Tags: []string{"testString"},
				// UseCases: []dataproductexchangeapiservicev1.UseCase{*useCaseModel},
				// Domain: domainModel,
				Type: []string{"data"},
				// PartsOut: []dataproductexchangeapiservicev1.DataProductPart{*dataProductPartModel},
			}

			dataProductVersion, response, err := dataProductExchangeApiServiceService.CreateDataProductVersion(createDataProductVersionOptions)
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
	})

	Describe(`GetInitializeStatus - Get the status of resources initialization in data product exchange`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetInitializeStatus(getInitializeStatusOptions *GetInitializeStatusOptions)`, func() {
			getInitializeStatusOptions := &dataproductexchangeapiservicev1.GetInitializeStatusOptions{
				ContainerID: &getStatusByCatalogIdLink,
			}

			initializeResource, response, err := dataProductExchangeApiServiceService.GetInitializeStatus(getInitializeStatusOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(initializeResource).ToNot(BeNil())
		})
	})

	Describe(`GetDataProduct - Retrieve a data product identified by id`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetDataProduct(getDataProductOptions *GetDataProductOptions)`, func() {
			getDataProductOptions := &dataproductexchangeapiservicev1.GetDataProductOptions{
				ID: &getDataProductByUserIdLink,
			}

			dataProduct, response, err := dataProductExchangeApiServiceService.GetDataProduct(getDataProductOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dataProduct).ToNot(BeNil())
		})
	})

	Describe(`ListDataProducts - Retrieve a list of data products`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListDataProducts(listDataProductsOptions *ListDataProductsOptions) with pagination`, func(){
			listDataProductsOptions := &dataproductexchangeapiservicev1.ListDataProductsOptions{
				Limit: core.Int64Ptr(int64(10)),
				// Start: core.StringPtr("testString"),
			}

			listDataProductsOptions.Start = nil
			listDataProductsOptions.Limit = core.Int64Ptr(1)

			var allResults []dataproductexchangeapiservicev1.DataProduct
			for {
				dataProductCollection, response, err := dataProductExchangeApiServiceService.ListDataProducts(listDataProductsOptions)
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(dataProductCollection).ToNot(BeNil())
				allResults = append(allResults, dataProductCollection.DataProducts...)

				listDataProductsOptions.Start, err = dataProductCollection.GetNextStart()
				Expect(err).To(BeNil())

				if listDataProductsOptions.Start == nil {
					break
				}
			}
			fmt.Fprintf(GinkgoWriter, "Retrieved a total of %d item(s) with pagination.\n", len(allResults))
		})
		It(`ListDataProducts(listDataProductsOptions *ListDataProductsOptions) using DataProductsPager`, func(){
			listDataProductsOptions := &dataproductexchangeapiservicev1.ListDataProductsOptions{
				Limit: core.Int64Ptr(int64(10)),
			}

			// Test GetNext().
			pager, err := dataProductExchangeApiServiceService.NewDataProductsPager(listDataProductsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []dataproductexchangeapiservicev1.DataProduct
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				Expect(err).To(BeNil())
				Expect(nextPage).ToNot(BeNil())
				allResults = append(allResults, nextPage...)
			}

			// Test GetAll().
			pager, err = dataProductExchangeApiServiceService.NewDataProductsPager(listDataProductsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			allItems, err := pager.GetAll()
			Expect(err).To(BeNil())
			Expect(allItems).ToNot(BeNil())

			Expect(len(allItems)).To(Equal(len(allResults)))
			fmt.Fprintf(GinkgoWriter, "ListDataProducts() returned a total of %d item(s) using DataProductsPager.\n", len(allResults))
		})
	})

	Describe(`ListDataProductVersions - Retrieve a list of data product versions`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListDataProductVersions(listDataProductVersionsOptions *ListDataProductVersionsOptions) with pagination`, func(){
			listDataProductVersionsOptions := &dataproductexchangeapiservicev1.ListDataProductVersionsOptions{
				AssetContainerID: &getListOfDataProductByCatalogIdLink,
				// DataProduct: core.StringPtr("testString"),
				// State: core.StringPtr("draft"),
				// Version: core.StringPtr("testString"),
				// Limit: core.Int64Ptr(int64(10)),
				// Start: core.StringPtr("testString"),
			}

			listDataProductVersionsOptions.Start = nil
			listDataProductVersionsOptions.Limit = core.Int64Ptr(1)

			var allResults []dataproductexchangeapiservicev1.DataProductVersionSummary
			for {
				dataProductVersionCollection, response, err := dataProductExchangeApiServiceService.ListDataProductVersions(listDataProductVersionsOptions)
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(dataProductVersionCollection).ToNot(BeNil())
				allResults = append(allResults, dataProductVersionCollection.DataProductVersions...)

				listDataProductVersionsOptions.Start, err = dataProductVersionCollection.GetNextStart()
				Expect(err).To(BeNil())

				if listDataProductVersionsOptions.Start == nil {
					break
				}
			}
			fmt.Fprintf(GinkgoWriter, "Retrieved a total of %d item(s) with pagination.\n", len(allResults))
		})
		It(`ListDataProductVersions(listDataProductVersionsOptions *ListDataProductVersionsOptions) using DataProductVersionsPager`, func(){
			listDataProductVersionsOptions := &dataproductexchangeapiservicev1.ListDataProductVersionsOptions{
				AssetContainerID: &getListOfDataProductByCatalogIdLink,
				// DataProduct: core.StringPtr("testString"),
				// State: core.StringPtr("draft"),
				// Version: core.StringPtr("testString"),
				Limit: core.Int64Ptr(int64(10)),
			}

			// Test GetNext().
			pager, err := dataProductExchangeApiServiceService.NewDataProductVersionsPager(listDataProductVersionsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []dataproductexchangeapiservicev1.DataProductVersionSummary
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				Expect(err).To(BeNil())
				Expect(nextPage).ToNot(BeNil())
				allResults = append(allResults, nextPage...)
			}

			// Test GetAll().
			pager, err = dataProductExchangeApiServiceService.NewDataProductVersionsPager(listDataProductVersionsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			allItems, err := pager.GetAll()
			Expect(err).To(BeNil())
			Expect(allItems).ToNot(BeNil())

			Expect(len(allItems)).To(Equal(len(allResults)))
			fmt.Fprintf(GinkgoWriter, "ListDataProductVersions() returned a total of %d item(s) using DataProductVersionsPager.\n", len(allResults))
		})
	})

	Describe(`GetDataProductVersion - Retrieve a data product version identified by ID`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetDataProductVersion(getDataProductVersionOptions *GetDataProductVersionOptions)`, func() {
			getDataProductVersionOptions := &dataproductexchangeapiservicev1.GetDataProductVersionOptions{
				ID: &getDataProductVersionByUserIdLink,
			}

			dataProductVersion, response, err := dataProductExchangeApiServiceService.GetDataProductVersion(getDataProductVersionOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dataProductVersion).ToNot(BeNil())
		})
	})

	Describe(`UpdateDataProductVersion - Update the data product version identified by ID`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateDataProductVersion(updateDataProductVersionOptions *UpdateDataProductVersionOptions)`, func() {
			jsonPatchOperationModel := &dataproductexchangeapiservicev1.JSONPatchOperation{
				Op: core.StringPtr("replace"),
				Path: core.StringPtr("/description"),
				// From: core.StringPtr("testString"),
				Value: core.StringPtr("This is the updated golang description"),
			}

			updateDataProductVersionOptions := &dataproductexchangeapiservicev1.UpdateDataProductVersionOptions{
				ID: &updateDataProductVersionByUserIdLink,
				JSONPatchInstructions: []dataproductexchangeapiservicev1.JSONPatchOperation{*jsonPatchOperationModel},
			}

			dataProductVersion, response, err := dataProductExchangeApiServiceService.UpdateDataProductVersion(updateDataProductVersionOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dataProductVersion).ToNot(BeNil())
		})
	})

	// Describe(`DeliverDataProductVersion - Deliver a data product identified by id`, func() {
	// 	BeforeEach(func() {
	// 		shouldSkipTest()
	// 	})
	// 	It(`DeliverDataProductVersion(deliverDataProductVersionOptions *DeliverDataProductVersionOptions)`, func() {
	// 		itemReferenceModel := &dataproductexchangeapiservicev1.ItemReference{
	// 			ID: core.StringPtr("d29c42eb-7100-4b7a-8257-c196dbcca1cd"),
	// 		}

	// 		orderReferenceModel := &dataproductexchangeapiservicev1.OrderReference{
	// 			ID: core.StringPtr("4705e047-1808-459a-805f-d5d13c947637"),
	// 			Items: []dataproductexchangeapiservicev1.ItemReference{*itemReferenceModel},
	// 		}

	// 		deliverDataProductVersionOptions := &dataproductexchangeapiservicev1.DeliverDataProductVersionOptions{
	// 			ID: &deliverDataProductVersionByUserIdLink,
	// 			Order: orderReferenceModel,
	// 		}

	// 		deliveryResource, response, err := dataProductExchangeApiServiceService.DeliverDataProductVersion(deliverDataProductVersionOptions)
	// 		Expect(err).To(BeNil())
	// 		Expect(response.StatusCode).To(Equal(202))
	// 		Expect(deliveryResource).ToNot(BeNil())
	// 	})
	// })

	Describe(`DeleteDataProductVersion - Delete a data product version identified by ID`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteDataProductVersion(deleteDataProductVersionOptions *DeleteDataProductVersionOptions)`, func() {
			deleteDataProductVersionOptions := &dataproductexchangeapiservicev1.DeleteDataProductVersionOptions{
				ID: &deleteDataProductVersionByUserIdLink,
			}

			response, err := dataProductExchangeApiServiceService.DeleteDataProductVersion(deleteDataProductVersionOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})
})

//
// Utility functions are declared in the unit test file
//
