//go:build integration

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

package dpxv1_test

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/IBM/data-product-exchange-go-sdk/dpxv1"
	"github.com/IBM/go-sdk-core/v5/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

/**
 * This file contains an integration test for the dpxv1 package.
 *
 * Notes:
 *
 * The integration test will automatically skip tests if the required config file is not available.
 */

var _ = Describe(`DpxV1 Integration Tests`, func() {
	const externalConfigFile = "../dpx_v1.env"

	var (
		err          error
		dpxService *dpxv1.DpxV1
		serviceURL   string
		config       map[string]string

		// Variables to hold link values
		completeContractTermsDocumentByContractIdLink string
		completeContractTermsDocumentByVersionIdLink string
		completeContractTermsDocumentLink string
		createDataProductVersionByCatalogIdLink string
		deleteContractTermsDocumentByContractIdLink string
		deleteContractTermsDocumentByVersionIdLink string
		deleteContractTermsDocumentLink string
		deleteDataProductVersionByUserIdLink string
		deliverDataProductVersionByUserIdLink string
		getContractTermsDocumentByIdLink string
		getContractTermsDocumentByVersionIdLink string
		getContractTermsDocumentsByContractIdLink string
		getDataProductByUserIdLink string
		getDataProductVersionByUserIdLink string
		getListOfDataProductByCatalogIdLink string
		getStatusByCatalogIdLink string
		updateContractTermsDocumentByContractIdLink string
		updateContractTermsDocumentByVersionIdLink string
		updateContractTermsDocumentLink string
		updateDataProductVersionByUserIdLink string
		uploadContractDocumentsByContractIdLink string
		uploadContractTermsDocumentsByVersionIdLink string
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
			config, err = core.GetServiceProperties(dpxv1.DefaultServiceName)
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
			dpxServiceOptions := &dpxv1.DpxV1Options{}

			dpxService, err = dpxv1.NewDpxV1UsingExternalConfig(dpxServiceOptions)
			Expect(err).To(BeNil())
			Expect(dpxService).ToNot(BeNil())
			Expect(dpxService.Service.Options.URL).To(Equal(serviceURL))

			core.SetLogger(core.NewLogger(core.LevelDebug, log.New(GinkgoWriter, "", log.LstdFlags), log.New(GinkgoWriter, "", log.LstdFlags)))
			dpxService.EnableRetries(4, 30*time.Second)
		})
	})

	Describe(`Initialize - Initialize resources in a data product exchange`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`Initialize(initializeOptions *InitializeOptions)`, func() {
			// containerReferenceModel := &dpxv1.ContainerReference{
			// 	ID: core.StringPtr("d29c42eb-7100-4b7a-8257-c196dbcca1cd"),
			// 	Type: core.StringPtr("catalog"),
			// }

			initializeOptions := &dpxv1.InitializeOptions{
				Container: nil,
				Force: core.BoolPtr(true),
				Reinitialize: core.BoolPtr(true),
				Include: []string{"delivery_methods", "data_product_samples", "domains_multi_industry"},
			}

			initializeResource, response, err := dpxService.Initialize(initializeOptions)
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
			containerReferenceModel := &dpxv1.ContainerReference{
				ID: core.StringPtr(createDataProductVersionByCatalogIdLink),
				Type: core.StringPtr("catalog"),
			}

			dataProductIdentityModel := &dpxv1.DataProductIdentity{
				ID: core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e"),
			}

			useCaseModel := &dpxv1.UseCase{
				ID: core.StringPtr("testString"),
				Name: core.StringPtr("testString"),
				Container: containerReferenceModel,
			}

			domainModel := &dpxv1.Domain{
				ID: core.StringPtr("testString"),
				Name: core.StringPtr("testString"),
				Container: containerReferenceModel,
			}

			assetPartReferenceModel := &dpxv1.AssetPartReference{
				ID: core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002"),
				Container: containerReferenceModel,
				Type: core.StringPtr("data_asset"),
			}

			deliveryMethodModel := &dpxv1.DeliveryMethod{
				ID: core.StringPtr("09cf5fcc-cb9d-4995-a8e4-16517b25229f"),
				Container: containerReferenceModel,
			}

			dataProductPartModel := &dpxv1.DataProductPart{
				Asset: assetPartReferenceModel,
				Revision: core.Int64Ptr(int64(1)),
				UpdatedAt: CreateMockDateTime("2023-07-01T22:22:34.876Z"),
				DeliveryMethods: []dpxv1.DeliveryMethod{*deliveryMethodModel},
			}

			contractTermsDocumentAttachmentModel := &dpxv1.ContractTermsDocumentAttachment{
				ID: core.StringPtr("testString"),
			}

			contractTermsDocumentModel := &dpxv1.ContractTermsDocument{
				URL: core.StringPtr("testString"),
				Type: core.StringPtr("terms_and_conditions"),
				Name: core.StringPtr("testString"),
				ID: core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002"),
				Attachment: contractTermsDocumentAttachmentModel,
			}

			assetReferenceModel := &dpxv1.AssetReference{
				ID: core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002"),
				Container: containerReferenceModel,
			}

			dataProductContractTermsModel := &dpxv1.DataProductContractTerms{
				ID: core.StringPtr("testString"),
				Documents: []dpxv1.ContractTermsDocument{*contractTermsDocumentModel},
				Asset: assetReferenceModel,
			}

			createDataProductVersionOptions := &dpxv1.CreateDataProductVersionOptions{
				Container: containerReferenceModel,
				Version: core.StringPtr("testString"),
				State: core.StringPtr("draft"),
				DataProduct: dataProductIdentityModel,
				Name: core.StringPtr("My New Data Product"),
				Description: core.StringPtr("testString"),
				Tags: []string{"testString"},
				UseCases: []dpxv1.UseCase{*useCaseModel},
				Domain: domainModel,
				Type: []string{"data"},
				PartsOut: []dpxv1.DataProductPart{*dataProductPartModel},
				ContractTerms: []dpxv1.DataProductContractTerms{*dataProductContractTermsModel},
			}

			dataProductVersion, response, err := dpxService.CreateDataProductVersion(createDataProductVersionOptions)
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
			completeContractTermsDocumentByVersionIdLink = *dataProductVersion.ID
			fmt.Fprintf(GinkgoWriter, "Saved completeContractTermsDocumentByVersionIdLink value: %v\n", completeContractTermsDocumentByVersionIdLink)
			uploadContractTermsDocumentsByVersionIdLink = *dataProductVersion.ID
			fmt.Fprintf(GinkgoWriter, "Saved uploadContractTermsDocumentsByVersionIdLink value: %v\n", uploadContractTermsDocumentsByVersionIdLink)
			getContractTermsDocumentByVersionIdLink = *dataProductVersion.ID
			fmt.Fprintf(GinkgoWriter, "Saved getContractTermsDocumentByVersionIdLink value: %v\n", getContractTermsDocumentByVersionIdLink)
			deleteContractTermsDocumentByVersionIdLink = *dataProductVersion.ID
			fmt.Fprintf(GinkgoWriter, "Saved deleteContractTermsDocumentByVersionIdLink value: %v\n", deleteContractTermsDocumentByVersionIdLink)
			updateContractTermsDocumentByVersionIdLink = *dataProductVersion.ID
			fmt.Fprintf(GinkgoWriter, "Saved updateContractTermsDocumentByVersionIdLink value: %v\n", updateContractTermsDocumentByVersionIdLink)
		})
	})

	Describe(`GetDataProductVersion - Retrieve a data product version identified by ID`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetDataProductVersion(getDataProductVersionOptions *GetDataProductVersionOptions)`, func() {
			getDataProductVersionOptions := &dpxv1.GetDataProductVersionOptions{
				ID: &getDataProductVersionByUserIdLink,
			}

			dataProductVersion, response, err := dpxService.GetDataProductVersion(getDataProductVersionOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dataProductVersion).ToNot(BeNil())

			completeContractTermsDocumentByContractIdLink = *dataProductVersion.ContractTerms[4].ID
			fmt.Fprintf(GinkgoWriter, "Saved completeContractTermsDocumentByContractIdLink value: %v\n", completeContractTermsDocumentByContractIdLink)
			uploadContractDocumentsByContractIdLink = *dataProductVersion.ContractTerms[4].ID
			fmt.Fprintf(GinkgoWriter, "Saved uploadContractDocumentsByContractIdLink value: %v\n", uploadContractDocumentsByContractIdLink)
			getContractTermsDocumentsByContractIdLink = *dataProductVersion.ContractTerms[4].ID
			fmt.Fprintf(GinkgoWriter, "Saved getContractTermsDocumentsByContractIdLink value: %v\n", getContractTermsDocumentsByContractIdLink)
			deleteContractTermsDocumentByContractIdLink = *dataProductVersion.ContractTerms[4].ID
			fmt.Fprintf(GinkgoWriter, "Saved deleteContractTermsDocumentByContractIdLink value: %v\n", deleteContractTermsDocumentByContractIdLink)
			updateContractTermsDocumentByContractIdLink = *dataProductVersion.ContractTerms[4].ID
			fmt.Fprintf(GinkgoWriter, "Saved updateContractTermsDocumentByContractIdLink value: %v\n", updateContractTermsDocumentByContractIdLink)
		})
	})

	Describe(`CreateContractTermsDocument - Upload a contract document to the Data Product Version contract terms`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateContractTermsDocument(createContractTermsDocumentOptions *CreateContractTermsDocumentOptions)`, func() {
			contractTermsDocumentAttachmentModel := &dpxv1.ContractTermsDocumentAttachment{
				ID: core.StringPtr("testString"),
			}

			createContractTermsDocumentOptions := &dpxv1.CreateContractTermsDocumentOptions{
				DataProductVersionID: &uploadContractTermsDocumentsByVersionIdLink,
				ContractTermsID: &uploadContractDocumentsByContractIdLink,
				Type: core.StringPtr("terms_and_conditions"),
				Name: core.StringPtr("Terms and conditions document"),
				ID: core.StringPtr("testString"),
				URL: core.StringPtr("testString"),
				Attachment: contractTermsDocumentAttachmentModel,
			}

			contractTermsDocument, response, err := dpxService.CreateContractTermsDocument(createContractTermsDocumentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(contractTermsDocument).ToNot(BeNil())

			updateContractTermsDocumentLink = *contractTermsDocument.ID
			fmt.Fprintf(GinkgoWriter, "Saved updateContractTermsDocumentLink value: %v\n", updateContractTermsDocumentLink)
			deleteContractTermsDocumentLink = *contractTermsDocument.ID
			fmt.Fprintf(GinkgoWriter, "Saved deleteContractTermsDocumentLink value: %v\n", deleteContractTermsDocumentLink)
			completeContractTermsDocumentLink = *contractTermsDocument.ID
			fmt.Fprintf(GinkgoWriter, "Saved completeContractTermsDocumentLink value: %v\n", completeContractTermsDocumentLink)
			getContractTermsDocumentByIdLink = *contractTermsDocument.ID
			fmt.Fprintf(GinkgoWriter, "Saved getContractTermsDocumentByIdLink value: %v\n", getContractTermsDocumentByIdLink)
		})
	})

	Describe(`GetInitializeStatus - Get the status of resources initialization in data product exchange`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetInitializeStatus(getInitializeStatusOptions *GetInitializeStatusOptions)`, func() {
			getInitializeStatusOptions := &dpxv1.GetInitializeStatusOptions{
				ContainerID: &getStatusByCatalogIdLink,
			}

			initializeResource, response, err := dpxService.GetInitializeStatus(getInitializeStatusOptions)
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
			getDataProductOptions := &dpxv1.GetDataProductOptions{
				ID: &getDataProductByUserIdLink,
			}

			dataProduct, response, err := dpxService.GetDataProduct(getDataProductOptions)
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
			listDataProductsOptions := &dpxv1.ListDataProductsOptions{
				Limit: core.Int64Ptr(int64(10)),
				Start: core.StringPtr("testString"),
			}

			listDataProductsOptions.Start = nil
			listDataProductsOptions.Limit = core.Int64Ptr(1)

			var allResults []dpxv1.DataProduct
			for {
				dataProductCollection, response, err := dpxService.ListDataProducts(listDataProductsOptions)
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
			listDataProductsOptions := &dpxv1.ListDataProductsOptions{
				Limit: core.Int64Ptr(int64(10)),
			}

			// Test GetNext().
			pager, err := dpxService.NewDataProductsPager(listDataProductsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []dpxv1.DataProduct
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				Expect(err).To(BeNil())
				Expect(nextPage).ToNot(BeNil())
				allResults = append(allResults, nextPage...)
			}

			// Test GetAll().
			pager, err = dpxService.NewDataProductsPager(listDataProductsOptions)
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
			listDataProductVersionsOptions := &dpxv1.ListDataProductVersionsOptions{
				AssetContainerID: &getListOfDataProductByCatalogIdLink,
				DataProduct: core.StringPtr("testString"),
				State: core.StringPtr("draft"),
				Version: core.StringPtr("testString"),
				Limit: core.Int64Ptr(int64(10)),
				Start: core.StringPtr("testString"),
			}

			listDataProductVersionsOptions.Start = nil
			listDataProductVersionsOptions.Limit = core.Int64Ptr(1)

			var allResults []dpxv1.DataProductVersionSummary
			for {
				dataProductVersionCollection, response, err := dpxService.ListDataProductVersions(listDataProductVersionsOptions)
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
			listDataProductVersionsOptions := &dpxv1.ListDataProductVersionsOptions{
				AssetContainerID: &getListOfDataProductByCatalogIdLink,
				DataProduct: core.StringPtr("testString"),
				State: core.StringPtr("draft"),
				Version: core.StringPtr("testString"),
				Limit: core.Int64Ptr(int64(10)),
			}

			// Test GetNext().
			pager, err := dpxService.NewDataProductVersionsPager(listDataProductVersionsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []dpxv1.DataProductVersionSummary
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				Expect(err).To(BeNil())
				Expect(nextPage).ToNot(BeNil())
				allResults = append(allResults, nextPage...)
			}

			// Test GetAll().
			pager, err = dpxService.NewDataProductVersionsPager(listDataProductVersionsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			allItems, err := pager.GetAll()
			Expect(err).To(BeNil())
			Expect(allItems).ToNot(BeNil())

			Expect(len(allItems)).To(Equal(len(allResults)))
			fmt.Fprintf(GinkgoWriter, "ListDataProductVersions() returned a total of %d item(s) using DataProductVersionsPager.\n", len(allResults))
		})
	})

	Describe(`UpdateDataProductVersion - Update the data product version identified by ID`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateDataProductVersion(updateDataProductVersionOptions *UpdateDataProductVersionOptions)`, func() {
			jsonPatchOperationModel := &dpxv1.JSONPatchOperation{
				Op: core.StringPtr("add"),
				Path: core.StringPtr("testString"),
				From: core.StringPtr("testString"),
				Value: core.StringPtr("testString"),
			}

			updateDataProductVersionOptions := &dpxv1.UpdateDataProductVersionOptions{
				ID: &updateDataProductVersionByUserIdLink,
				JSONPatchInstructions: []dpxv1.JSONPatchOperation{*jsonPatchOperationModel},
			}

			dataProductVersion, response, err := dpxService.UpdateDataProductVersion(updateDataProductVersionOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dataProductVersion).ToNot(BeNil())
		})
	})

	Describe(`CompleteContractTermsDocument - Complete a contract document upload`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CompleteContractTermsDocument(completeContractTermsDocumentOptions *CompleteContractTermsDocumentOptions)`, func() {
			completeContractTermsDocumentOptions := &dpxv1.CompleteContractTermsDocumentOptions{
				DataProductVersionID: &completeContractTermsDocumentByVersionIdLink,
				ContractTermsID: &completeContractTermsDocumentByContractIdLink,
				DocumentID: &completeContractTermsDocumentLink,
			}

			contractTermsDocument, response, err := dpxService.CompleteContractTermsDocument(completeContractTermsDocumentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(contractTermsDocument).ToNot(BeNil())
		})
	})

	Describe(`GetContractTermsDocument - Get a contract document`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetContractTermsDocument(getContractTermsDocumentOptions *GetContractTermsDocumentOptions)`, func() {
			getContractTermsDocumentOptions := &dpxv1.GetContractTermsDocumentOptions{
				DataProductVersionID: &getContractTermsDocumentByVersionIdLink,
				ContractTermsID: &getContractTermsDocumentsByContractIdLink,
				DocumentID: &getContractTermsDocumentByIdLink,
			}

			contractTermsDocument, response, err := dpxService.GetContractTermsDocument(getContractTermsDocumentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(contractTermsDocument).ToNot(BeNil())
		})
	})

	Describe(`UpdateContractTermsDocument - Update a contract document`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateContractTermsDocument(updateContractTermsDocumentOptions *UpdateContractTermsDocumentOptions)`, func() {
			jsonPatchOperationModel := &dpxv1.JSONPatchOperation{
				Op: core.StringPtr("add"),
				Path: core.StringPtr("testString"),
				From: core.StringPtr("testString"),
				Value: core.StringPtr("testString"),
			}

			updateContractTermsDocumentOptions := &dpxv1.UpdateContractTermsDocumentOptions{
				DataProductVersionID: &updateContractTermsDocumentByVersionIdLink,
				ContractTermsID: &updateContractTermsDocumentByContractIdLink,
				DocumentID: &updateContractTermsDocumentLink,
				JSONPatchInstructions: []dpxv1.JSONPatchOperation{*jsonPatchOperationModel},
			}

			contractTermsDocument, response, err := dpxService.UpdateContractTermsDocument(updateContractTermsDocumentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(contractTermsDocument).ToNot(BeNil())
		})
	})

	Describe(`DeleteContractTermsDocument - Delete a contract document`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteContractTermsDocument(deleteContractTermsDocumentOptions *DeleteContractTermsDocumentOptions)`, func() {
			deleteContractTermsDocumentOptions := &dpxv1.DeleteContractTermsDocumentOptions{
				DataProductVersionID: &deleteContractTermsDocumentByVersionIdLink,
				ContractTermsID: &deleteContractTermsDocumentByContractIdLink,
				DocumentID: &deleteContractTermsDocumentLink,
			}

			response, err := dpxService.DeleteContractTermsDocument(deleteContractTermsDocumentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteDataProductVersion - Delete a data product version identified by ID`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteDataProductVersion(deleteDataProductVersionOptions *DeleteDataProductVersionOptions)`, func() {
			deleteDataProductVersionOptions := &dpxv1.DeleteDataProductVersionOptions{
				ID: &deleteDataProductVersionByUserIdLink,
			}

			response, err := dpxService.DeleteDataProductVersion(deleteDataProductVersionOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})
})

//
// Utility functions are declared in the unit test file
//
