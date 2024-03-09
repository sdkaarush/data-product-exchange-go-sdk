//go:build integration

/**
 * (C) Copyright IBM Corp. 2024.
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
		err        error
		dpxService *dpxv1.DpxV1
		serviceURL string
		config     map[string]string

		// Variables to hold link values
		containerIdLink           string
		contractTermsIdLink       string
		dataProductIdLink         string
		documentIdLink            string
		draftIdLink               string
		optionalDataProductIdLink string
		releaseIdLink             string
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

	Describe(`Initialize - Initialize resources`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`Initialize(initializeOptions *InitializeOptions)`, func() {
			// containerReferenceModel := &dpxv1.ContainerReference{
			// 	ID:   &containerIdLink,
			// 	Type: core.StringPtr("catalog"),
			// }

			initializeOptions := &dpxv1.InitializeOptions{
				Container: nil,
				Include:   []string{"delivery_methods", "data_product_samples", "domains_multi_industry"},
			}

			initializeResource, response, err := dpxService.Initialize(initializeOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(initializeResource).ToNot(BeNil())

			containerIdLink = *initializeResource.Container.ID
			fmt.Fprintf(GinkgoWriter, "Saved containerIdLink value: %v\n", containerIdLink)
		})
	})

	Describe(`CreateDataProduct - Create a new data product`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateDataProduct(createDataProductOptions *CreateDataProductOptions)`, func() {
			// dataProductIdentityModel := &dpxv1.DataProductIdentity{
			// 	ID: core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e"),
			// }

			containerReferenceModel := &dpxv1.ContainerReference{
				ID:   &containerIdLink,
				Type: core.StringPtr("catalog"),
			}

			assetReferenceModel := &dpxv1.AssetReference{
				ID:        core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002"),
				Container: containerReferenceModel,
			}

			// useCaseModel := &dpxv1.UseCase{
			// 	ID:        core.StringPtr("testString"),
			// 	Name:      core.StringPtr("testString"),
			// 	Container: containerReferenceModel,
			// }

			// domainModel := &dpxv1.Domain{
			// 	ID:        core.StringPtr("testString"),
			// 	Name:      core.StringPtr("testString"),
			// 	Container: containerReferenceModel,
			// }

			// assetPartReferenceModel := &dpxv1.AssetPartReference{
			// 	ID:        core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002"),
			// 	Container: containerReferenceModel,
			// 	Type:      core.StringPtr("data_asset"),
			// }

			// deliveryMethodModel := &dpxv1.DeliveryMethod{
			// 	ID:        core.StringPtr("09cf5fcc-cb9d-4995-a8e4-16517b25229f"),
			// 	Container: containerReferenceModel,
			// }

			// dataProductPartModel := &dpxv1.DataProductPart{
			// 	Asset:           assetPartReferenceModel,
			// 	Revision:        core.Int64Ptr(int64(1)),
			// 	UpdatedAt:       CreateMockDateTime("2023-07-01T22:22:34.876Z"),
			// 	DeliveryMethods: []dpxv1.DeliveryMethod{*deliveryMethodModel},
			// }

			// contractTermsDocumentAttachmentModel := &dpxv1.ContractTermsDocumentAttachment{
			// 	ID: core.StringPtr("testString"),
			// }

			// contractTermsDocumentModel := &dpxv1.ContractTermsDocument{
			// 	URL:        core.StringPtr("testString"),
			// 	Type:       core.StringPtr("terms_and_conditions"),
			// 	Name:       core.StringPtr("testString"),
			// 	ID:         core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002"),
			// 	Attachment: contractTermsDocumentAttachmentModel,
			// 	UploadURL:  core.StringPtr("testString"),
			// }

			// dataProductContractTermsModel := &dpxv1.DataProductContractTerms{
			// 	Asset:     assetReferenceModel,
			// 	ID:        &contractTermsIdLink,
			// 	Documents: []dpxv1.ContractTermsDocument{*contractTermsDocumentModel},
			// }

			dataProductVersionPrototypeModel := &dpxv1.DataProductVersionPrototype{
				// Version:       core.StringPtr("1.0.0"),
				// State:         core.StringPtr("draft"),
				// DataProduct:   dataProductIdentityModel,
				Name:        core.StringPtr("My New Data Product"),
				Description: core.StringPtr("This is a description of My Data Product."),
				Asset:       assetReferenceModel,
				// Tags:          []string{"testString"},
				// UseCases:      []dpxv1.UseCase{*useCaseModel},
				// Domain:        domainModel,
				Types: []string{"data"},
				// PartsOut:      []dpxv1.DataProductPart{*dataProductPartModel},
				// ContractTerms: []dpxv1.DataProductContractTerms{*dataProductContractTermsModel},
				// IsRestricted:  core.BoolPtr(true),
			}

			createDataProductOptions := &dpxv1.CreateDataProductOptions{
				Drafts: []dpxv1.DataProductVersionPrototype{*dataProductVersionPrototypeModel},
			}

			dataProduct, response, err := dpxService.CreateDataProduct(createDataProductOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dataProduct).ToNot(BeNil())

			optionalDataProductIdLink = *dataProduct.ID
			fmt.Fprintf(GinkgoWriter, "Saved optionalDataProductIdLink value: %v\n", optionalDataProductIdLink)
			dataProductIdLink = *dataProduct.ID
			fmt.Fprintf(GinkgoWriter, "Saved dataProductIdLink value: %v\n", dataProductIdLink)
		})
	})

	Describe(`CreateDataProductDraft - Create a new draft of an existing data product`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateDataProductDraft(createDataProductDraftOptions *CreateDataProductDraftOptions)`, func() {
			containerReferenceModel := &dpxv1.ContainerReference{
				ID:   &containerIdLink,
				Type: core.StringPtr("catalog"),
			}

			assetReferenceModel := &dpxv1.AssetReference{
				ID:        core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002"),
				Container: containerReferenceModel,
			}

			dataProductIdentityModel := &dpxv1.DataProductIdentity{
				ID: &dataProductIdLink,
			}

			// useCaseModel := &dpxv1.UseCase{
			// 	ID:        core.StringPtr("testString"),
			// 	Name:      core.StringPtr("testString"),
			// 	Container: containerReferenceModel,
			// }

			domainModel := &dpxv1.Domain{
				ID:        core.StringPtr("918c0bfd-6943-4468-b74f-bc111018e0d1"),
				Name:      core.StringPtr("Customer Service"),
				Container: containerReferenceModel,
			}

			// assetPartReferenceModel := &dpxv1.AssetPartReference{
			// 	ID:        core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002"),
			// 	Container: containerReferenceModel,
			// 	Type:      core.StringPtr("data_asset"),
			// }

			// deliveryMethodModel := &dpxv1.DeliveryMethod{
			// 	ID:        core.StringPtr("09cf5fcc-cb9d-4995-a8e4-16517b25229f"),
			// 	Container: containerReferenceModel,
			// }

			// dataProductPartModel := &dpxv1.DataProductPart{
			// 	Asset:           assetPartReferenceModel,
			// 	Revision:        core.Int64Ptr(int64(1)),
			// 	UpdatedAt:       CreateMockDateTime("2023-07-01T22:22:34.876Z"),
			// 	DeliveryMethods: []dpxv1.DeliveryMethod{*deliveryMethodModel},
			// }

			// contractTermsDocumentAttachmentModel := &dpxv1.ContractTermsDocumentAttachment{
			// 	ID: core.StringPtr("testString"),
			// }

			// contractTermsDocumentModel := &dpxv1.ContractTermsDocument{
			// 	URL:        core.StringPtr("testString"),
			// 	Type:       core.StringPtr("terms_and_conditions"),
			// 	Name:       core.StringPtr("testString"),
			// 	ID:         core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002"),
			// 	Attachment: contractTermsDocumentAttachmentModel,
			// 	UploadURL:  core.StringPtr("testString"),
			// }

			// dataProductContractTermsModel := &dpxv1.DataProductContractTerms{
			// 	Asset:     assetReferenceModel,
			// 	ID:        &contractTermsIdLink,
			// 	Documents: []dpxv1.ContractTermsDocument{*contractTermsDocumentModel},
			// }

			createDataProductDraftOptions := &dpxv1.CreateDataProductDraftOptions{
				DataProductID: &dataProductIdLink,
				Asset:         assetReferenceModel,
				Version:       core.StringPtr("1.2.0"),
				// State:         core.StringPtr("draft"),
				DataProduct: dataProductIdentityModel,
				Name:        core.StringPtr("data_product_test"),
				Description: core.StringPtr("testString"),
				// Tags:          []string{"testString"},
				// UseCases:      []dpxv1.UseCase{*useCaseModel},
				Domain: domainModel,
				Types:  []string{"data"},
				// PartsOut:      []dpxv1.DataProductPart{*dataProductPartModel},
				// ContractTerms: []dpxv1.DataProductContractTerms{*dataProductContractTermsModel},
				IsRestricted: core.BoolPtr(true),
			}

			dataProductVersion, response, err := dpxService.CreateDataProductDraft(createDataProductDraftOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(dataProductVersion).ToNot(BeNil())

			draftIdLink = *dataProductVersion.ID
			fmt.Fprintf(GinkgoWriter, "Saved draftIdLink value: %v\n", draftIdLink)
			contractTermsIdLink = *dataProductVersion.ContractTerms[0].ID
			fmt.Fprintf(GinkgoWriter, "Saved contractTermsIdLink value: %v\n", contractTermsIdLink)
		})
	})

	Describe(`DeleteDataProductDraft - Delete a data product draft identified by ID`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteDataProductDraft(deleteDataProductDraftOptions *DeleteDataProductDraftOptions)`, func() {
			deleteDataProductDraftOptions := &dpxv1.DeleteDataProductDraftOptions{
				DataProductID: &optionalDataProductIdLink,
				DraftID:       &draftIdLink,
			}

			response, err := dpxService.DeleteDataProductDraft(deleteDataProductDraftOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`CreateDataProductDraftAgain	 - Create a new draft of an existing data product`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateDataProductDraft(createDataProductDraftOptions *CreateDataProductDraftOptions)`, func() {
			containerReferenceModel := &dpxv1.ContainerReference{
				ID:   &containerIdLink,
				Type: core.StringPtr("catalog"),
			}

			assetReferenceModel := &dpxv1.AssetReference{
				ID:        core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002"),
				Container: containerReferenceModel,
			}

			dataProductIdentityModel := &dpxv1.DataProductIdentity{
				ID: &dataProductIdLink,
			}

			// useCaseModel := &dpxv1.UseCase{
			// 	ID:        core.StringPtr("testString"),
			// 	Name:      core.StringPtr("testString"),
			// 	Container: containerReferenceModel,
			// }

			domainModel := &dpxv1.Domain{
				ID:        core.StringPtr("918c0bfd-6943-4468-b74f-bc111018e0d1"),
				Name:      core.StringPtr("Customer Service"),
				Container: containerReferenceModel,
			}

			// assetPartReferenceModel := &dpxv1.AssetPartReference{
			// 	ID:        core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002"),
			// 	Container: containerReferenceModel,
			// 	Type:      core.StringPtr("data_asset"),
			// }

			// deliveryMethodModel := &dpxv1.DeliveryMethod{
			// 	ID:        core.StringPtr("09cf5fcc-cb9d-4995-a8e4-16517b25229f"),
			// 	Container: containerReferenceModel,
			// }

			// dataProductPartModel := &dpxv1.DataProductPart{
			// 	Asset:           assetPartReferenceModel,
			// 	Revision:        core.Int64Ptr(int64(1)),
			// 	UpdatedAt:       CreateMockDateTime("2023-07-01T22:22:34.876Z"),
			// 	DeliveryMethods: []dpxv1.DeliveryMethod{*deliveryMethodModel},
			// }

			// contractTermsDocumentAttachmentModel := &dpxv1.ContractTermsDocumentAttachment{
			// 	ID: core.StringPtr("testString"),
			// }

			// contractTermsDocumentModel := &dpxv1.ContractTermsDocument{
			// 	URL:        core.StringPtr("testString"),
			// 	Type:       core.StringPtr("terms_and_conditions"),
			// 	Name:       core.StringPtr("testString"),
			// 	ID:         core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002"),
			// 	Attachment: contractTermsDocumentAttachmentModel,
			// 	UploadURL:  core.StringPtr("testString"),
			// }

			// dataProductContractTermsModel := &dpxv1.DataProductContractTerms{
			// 	Asset:     assetReferenceModel,
			// 	ID:        &contractTermsIdLink,
			// 	Documents: []dpxv1.ContractTermsDocument{*contractTermsDocumentModel},
			// }

			createDataProductDraftOptions := &dpxv1.CreateDataProductDraftOptions{
				DataProductID: &dataProductIdLink,
				Asset:         assetReferenceModel,
				Version:       core.StringPtr("1.2.0"),
				// State:         core.StringPtr("draft"),
				DataProduct: dataProductIdentityModel,
				Name:        core.StringPtr("data_product_test"),
				Description: core.StringPtr("testString"),
				// Tags:          []string{"testString"},
				// UseCases:      []dpxv1.UseCase{*useCaseModel},
				Domain: domainModel,
				Types:  []string{"data"},
				// PartsOut:      []dpxv1.DataProductPart{*dataProductPartModel},
				// ContractTerms: []dpxv1.DataProductContractTerms{*dataProductContractTermsModel},
				IsRestricted: core.BoolPtr(true),
			}

			dataProductVersion, response, err := dpxService.CreateDataProductDraft(createDataProductDraftOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(dataProductVersion).ToNot(BeNil())

			draftIdLink = *dataProductVersion.ID
			fmt.Fprintf(GinkgoWriter, "Saved draftIdLink value: %v\n", draftIdLink)
			contractTermsIdLink = *dataProductVersion.ContractTerms[0].ID
			fmt.Fprintf(GinkgoWriter, "Saved contractTermsIdLink value: %v\n", contractTermsIdLink)
		})
	})

	Describe(`CreateDraftContractTermsDocument - Upload a contract document to the data product draft contract terms`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateDraftContractTermsDocument(createDraftContractTermsDocumentOptions *CreateDraftContractTermsDocumentOptions)`, func() {
			// contractTermsDocumentAttachmentModel := &dpxv1.ContractTermsDocumentAttachment{
			// 	ID: core.StringPtr("testString"),
			// }

			createDraftContractTermsDocumentOptions := &dpxv1.CreateDraftContractTermsDocumentOptions{
				DataProductID:   &optionalDataProductIdLink,
				DraftID:         &draftIdLink,
				ContractTermsID: &contractTermsIdLink,
				Type:            core.StringPtr("terms_and_conditions"),
				Name:            core.StringPtr("Terms and conditions document"),
				ID:              core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e"),
				URL:             core.StringPtr("https://www.google.com"),
				// Attachment:      contractTermsDocumentAttachmentModel,
				// UploadURL:       core.StringPtr("testString"),
			}

			contractTermsDocument, response, err := dpxService.CreateDraftContractTermsDocument(createDraftContractTermsDocumentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(contractTermsDocument).ToNot(BeNil())

			documentIdLink = *contractTermsDocument.ID
			fmt.Fprintf(GinkgoWriter, "Saved documentIdLink value: %v\n", documentIdLink)
		})
	})

	Describe(`DeleteDraftContractTermsDocument - Delete a contract document`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteDraftContractTermsDocument(deleteDraftContractTermsDocumentOptions *DeleteDraftContractTermsDocumentOptions)`, func() {
			deleteDraftContractTermsDocumentOptions := &dpxv1.DeleteDraftContractTermsDocumentOptions{
				DataProductID:   &optionalDataProductIdLink,
				DraftID:         &draftIdLink,
				ContractTermsID: &contractTermsIdLink,
				DocumentID:      &documentIdLink,
			}

			response, err := dpxService.DeleteDraftContractTermsDocument(deleteDraftContractTermsDocumentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`CreateDraftContractTermsDocumentAgain - Upload a contract document to the data product draft contract terms`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateDraftContractTermsDocument(createDraftContractTermsDocumentOptions *CreateDraftContractTermsDocumentOptions)`, func() {
			// contractTermsDocumentAttachmentModel := &dpxv1.ContractTermsDocumentAttachment{
			// 	ID: core.StringPtr("testString"),
			// }

			createDraftContractTermsDocumentOptions := &dpxv1.CreateDraftContractTermsDocumentOptions{
				DataProductID:   &optionalDataProductIdLink,
				DraftID:         &draftIdLink,
				ContractTermsID: &contractTermsIdLink,
				Type:            core.StringPtr("terms_and_conditions"),
				Name:            core.StringPtr("Terms and conditions document"),
				ID:              core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e"),
				URL:             core.StringPtr("https://www.google.com"),
				// Attachment:      contractTermsDocumentAttachmentModel,
				// UploadURL:       core.StringPtr("testString"),
			}

			contractTermsDocument, response, err := dpxService.CreateDraftContractTermsDocument(createDraftContractTermsDocumentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(contractTermsDocument).ToNot(BeNil())

			documentIdLink = *contractTermsDocument.ID
			fmt.Fprintf(GinkgoWriter, "Saved documentIdLink value: %v\n", documentIdLink)
		})
	})

	Describe(`UpdateDataProductDraft - Update the data product draft identified by ID`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateDataProductDraft(updateDataProductDraftOptions *UpdateDataProductDraftOptions)`, func() {
			// Define the asset structure
			asset := map[string]interface{}{
				"id": "669a570b-31f7-4c84-bfd1-851282ab5b86",
				"container": map[string]string{
					"id":   "b6eb50b4-ace4-4dab-b2c4-318bb4c032a6",
					"type": "catalog",
				},
			}

			// Create a list to hold the asset object
			partsOutList := []map[string]interface{}{{"asset": asset}}

			jsonPatchOperationModel := &dpxv1.JSONPatchOperation{
				Op:   core.StringPtr("add"),
				Path: core.StringPtr("/parts_out"),
				// From:  core.StringPtr("testString"),
				Value: partsOutList,
			}

			updateDataProductDraftOptions := &dpxv1.UpdateDataProductDraftOptions{
				DataProductID:         &optionalDataProductIdLink,
				DraftID:               &draftIdLink,
				JSONPatchInstructions: []dpxv1.JSONPatchOperation{*jsonPatchOperationModel},
			}

			dataProductVersion, response, err := dpxService.UpdateDataProductDraft(updateDataProductDraftOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dataProductVersion).ToNot(BeNil())
		})
	})

	Describe(`GetDataProductDraft - Get a draft of an existing data product`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetDataProductDraft(getDataProductDraftOptions *GetDataProductDraftOptions)`, func() {
			getDataProductDraftOptions := &dpxv1.GetDataProductDraftOptions{
				DataProductID: &optionalDataProductIdLink,
				DraftID:       &draftIdLink,
			}

			dataProductVersion, response, err := dpxService.GetDataProductDraft(getDataProductDraftOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dataProductVersion).ToNot(BeNil())
		})
	})

	Describe(`UpdateDraftContractTermsDocument - Update a contract document`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateDraftContractTermsDocument(updateDraftContractTermsDocumentOptions *UpdateDraftContractTermsDocumentOptions)`, func() {
			jsonPatchOperationModel := &dpxv1.JSONPatchOperation{
				Op:    core.StringPtr("replace"),
				Path:  core.StringPtr("/url"),
				Value: "https://google.com",
			}

			updateDraftContractTermsDocumentOptions := &dpxv1.UpdateDraftContractTermsDocumentOptions{
				DataProductID:         &optionalDataProductIdLink,
				DraftID:               &draftIdLink,
				ContractTermsID:       &contractTermsIdLink,
				DocumentID:            &documentIdLink,
				JSONPatchInstructions: []dpxv1.JSONPatchOperation{*jsonPatchOperationModel},
			}

			contractTermsDocument, response, err := dpxService.UpdateDraftContractTermsDocument(updateDraftContractTermsDocumentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(contractTermsDocument).ToNot(BeNil())
		})
	})

	Describe(`GetInitializeStatus - Get resource initialization status`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetInitializeStatus(getInitializeStatusOptions *GetInitializeStatusOptions)`, func() {
			getInitializeStatusOptions := &dpxv1.GetInitializeStatusOptions{
				ContainerID: &containerIdLink,
			}

			initializeResource, response, err := dpxService.GetInitializeStatus(getInitializeStatusOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(initializeResource).ToNot(BeNil())
		})
	})

	Describe(`GetDraftContractTermsDocument - Get a contract document`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetDraftContractTermsDocument(getDraftContractTermsDocumentOptions *GetDraftContractTermsDocumentOptions)`, func() {
			getDraftContractTermsDocumentOptions := &dpxv1.GetDraftContractTermsDocumentOptions{
				DataProductID:   &optionalDataProductIdLink,
				DraftID:         &draftIdLink,
				ContractTermsID: &contractTermsIdLink,
				DocumentID:      &documentIdLink,
			}

			contractTermsDocument, response, err := dpxService.GetDraftContractTermsDocument(getDraftContractTermsDocumentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(contractTermsDocument).ToNot(BeNil())
		})
	})

	Describe(`PublishDataProductDraft - Publish a draft of an existing data product`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PublishDataProductDraft(publishDataProductDraftOptions *PublishDataProductDraftOptions)`, func() {
			publishDataProductDraftOptions := &dpxv1.PublishDataProductDraftOptions{
				DataProductID: &optionalDataProductIdLink,
				DraftID:       &draftIdLink,
			}

			dataProductVersion, response, err := dpxService.PublishDataProductDraft(publishDataProductDraftOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dataProductVersion).ToNot(BeNil())

			releaseIdLink = *dataProductVersion.ID
			fmt.Fprintf(GinkgoWriter, "Saved releaseIdLink value: %v\n", releaseIdLink)
		})
	})

	Describe(`ManageApiKeys - Rotate credentials for a Data Product Exchange instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ManageApiKeys(manageApiKeysOptions *ManageApiKeysOptions)`, func() {
			manageApiKeysOptions := &dpxv1.ManageApiKeysOptions{}

			response, err := dpxService.ManageApiKeys(manageApiKeysOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`ListDataProducts - Retrieve a list of data products`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListDataProducts(listDataProductsOptions *ListDataProductsOptions) with pagination`, func() {
			listDataProductsOptions := &dpxv1.ListDataProductsOptions{
				Limit: core.Int64Ptr(int64(10)),
			}

			listDataProductsOptions.Start = nil
			listDataProductsOptions.Limit = core.Int64Ptr(1)

			var allResults []dpxv1.DataProductSummary
			for {
				dataProductSummaryCollection, response, err := dpxService.ListDataProducts(listDataProductsOptions)
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(dataProductSummaryCollection).ToNot(BeNil())
				allResults = append(allResults, dataProductSummaryCollection.DataProducts...)

				listDataProductsOptions.Start, err = dataProductSummaryCollection.GetNextStart()
				Expect(err).To(BeNil())

				if listDataProductsOptions.Start == nil {
					break
				}
			}
			fmt.Fprintf(GinkgoWriter, "Retrieved a total of %d item(s) with pagination.\n", len(allResults))
		})
		It(`ListDataProducts(listDataProductsOptions *ListDataProductsOptions) using DataProductsPager`, func() {
			listDataProductsOptions := &dpxv1.ListDataProductsOptions{
				Limit: core.Int64Ptr(int64(10)),
			}

			// Test GetNext().
			pager, err := dpxService.NewDataProductsPager(listDataProductsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []dpxv1.DataProductSummary
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

	Describe(`GetDataProduct - Retrieve a data product identified by id`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetDataProduct(getDataProductOptions *GetDataProductOptions)`, func() {
			getDataProductOptions := &dpxv1.GetDataProductOptions{
				DataProductID: &dataProductIdLink,
			}

			dataProduct, response, err := dpxService.GetDataProduct(getDataProductOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dataProduct).ToNot(BeNil())
		})
	})

	Describe(`CompleteDraftContractTermsDocument - Complete a contract document upload operation`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CompleteDraftContractTermsDocument(completeDraftContractTermsDocumentOptions *CompleteDraftContractTermsDocumentOptions)`, func() {
			completeDraftContractTermsDocumentOptions := &dpxv1.CompleteDraftContractTermsDocumentOptions{
				DataProductID:   &optionalDataProductIdLink,
				DraftID:         &draftIdLink,
				ContractTermsID: &contractTermsIdLink,
				DocumentID:      &documentIdLink,
			}

			contractTermsDocument, response, err := dpxService.CompleteDraftContractTermsDocument(completeDraftContractTermsDocumentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(contractTermsDocument).ToNot(BeNil())
		})
	})

	Describe(`ListDataProductDrafts - Retrieve a list of data product drafts`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListDataProductDrafts(listDataProductDraftsOptions *ListDataProductDraftsOptions) with pagination`, func() {
			listDataProductDraftsOptions := &dpxv1.ListDataProductDraftsOptions{
				DataProductID: &optionalDataProductIdLink,
				Limit:         core.Int64Ptr(int64(10)),
			}

			listDataProductDraftsOptions.Start = nil
			listDataProductDraftsOptions.Limit = core.Int64Ptr(1)

			var allResults []dpxv1.DataProductVersionSummary
			for {
				dataProductDraftCollection, response, err := dpxService.ListDataProductDrafts(listDataProductDraftsOptions)
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(dataProductDraftCollection).ToNot(BeNil())
				allResults = append(allResults, dataProductDraftCollection.Drafts...)

				listDataProductDraftsOptions.Start, err = dataProductDraftCollection.GetNextStart()
				Expect(err).To(BeNil())

				if listDataProductDraftsOptions.Start == nil {
					break
				}
			}
			fmt.Fprintf(GinkgoWriter, "Retrieved a total of %d item(s) with pagination.\n", len(allResults))
		})
		It(`ListDataProductDrafts(listDataProductDraftsOptions *ListDataProductDraftsOptions) using DataProductDraftsPager`, func() {
			listDataProductDraftsOptions := &dpxv1.ListDataProductDraftsOptions{
				DataProductID: &optionalDataProductIdLink,
				Limit:         core.Int64Ptr(int64(10)),
			}

			// Test GetNext().
			pager, err := dpxService.NewDataProductDraftsPager(listDataProductDraftsOptions)
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
			pager, err = dpxService.NewDataProductDraftsPager(listDataProductDraftsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			allItems, err := pager.GetAll()
			Expect(err).To(BeNil())
			Expect(allItems).ToNot(BeNil())

			Expect(len(allItems)).To(Equal(len(allResults)))
			fmt.Fprintf(GinkgoWriter, "ListDataProductDrafts() returned a total of %d item(s) using DataProductDraftsPager.\n", len(allResults))
		})
	})

	Describe(`GetDataProductRelease - Get a release of an existing data product`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetDataProductRelease(getDataProductReleaseOptions *GetDataProductReleaseOptions)`, func() {
			getDataProductReleaseOptions := &dpxv1.GetDataProductReleaseOptions{
				DataProductID: &optionalDataProductIdLink,
				ReleaseID:     &releaseIdLink,
			}

			dataProductVersion, response, err := dpxService.GetDataProductRelease(getDataProductReleaseOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dataProductVersion).ToNot(BeNil())
		})
	})

	Describe(`UpdateDataProductRelease - Update the data product release identified by ID`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateDataProductRelease(updateDataProductReleaseOptions *UpdateDataProductReleaseOptions)`, func() {
			jsonPatchOperationModel := &dpxv1.JSONPatchOperation{
				Op:    core.StringPtr("replace"),
				Path:  core.StringPtr("/description"),
				Value: "New Description",
			}

			updateDataProductReleaseOptions := &dpxv1.UpdateDataProductReleaseOptions{
				DataProductID:         &optionalDataProductIdLink,
				ReleaseID:             &releaseIdLink,
				JSONPatchInstructions: []dpxv1.JSONPatchOperation{*jsonPatchOperationModel},
			}

			dataProductVersion, response, err := dpxService.UpdateDataProductRelease(updateDataProductReleaseOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dataProductVersion).ToNot(BeNil())
		})
	})

	Describe(`GetReleaseContractTermsDocument - Get a contract document`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetReleaseContractTermsDocument(getReleaseContractTermsDocumentOptions *GetReleaseContractTermsDocumentOptions)`, func() {
			getReleaseContractTermsDocumentOptions := &dpxv1.GetReleaseContractTermsDocumentOptions{
				DataProductID:   &optionalDataProductIdLink,
				ReleaseID:       &releaseIdLink,
				ContractTermsID: &contractTermsIdLink,
				DocumentID:      &documentIdLink,
			}

			contractTermsDocument, response, err := dpxService.GetReleaseContractTermsDocument(getReleaseContractTermsDocumentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(contractTermsDocument).ToNot(BeNil())
		})
	})

	Describe(`ListDataProductReleases - Retrieve a list of data product releases`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListDataProductReleases(listDataProductReleasesOptions *ListDataProductReleasesOptions) with pagination`, func() {
			listDataProductReleasesOptions := &dpxv1.ListDataProductReleasesOptions{
				DataProductID: &optionalDataProductIdLink,
				// AssetContainerID: core.StringPtr("testString"),
				State: []string{"available"},
				// Version:          core.StringPtr("testString"),
				Limit: core.Int64Ptr(int64(10)),
				// Start:            core.StringPtr("testString"),
			}

			listDataProductReleasesOptions.Start = nil
			listDataProductReleasesOptions.Limit = core.Int64Ptr(1)

			var allResults []dpxv1.DataProductVersionSummary
			for {
				dataProductReleaseCollection, response, err := dpxService.ListDataProductReleases(listDataProductReleasesOptions)
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(dataProductReleaseCollection).ToNot(BeNil())
				allResults = append(allResults, dataProductReleaseCollection.Releases...)

				listDataProductReleasesOptions.Start, err = dataProductReleaseCollection.GetNextStart()
				Expect(err).To(BeNil())

				if listDataProductReleasesOptions.Start == nil {
					break
				}
			}
			fmt.Fprintf(GinkgoWriter, "Retrieved a total of %d item(s) with pagination.\n", len(allResults))
		})
		It(`ListDataProductReleases(listDataProductReleasesOptions *ListDataProductReleasesOptions) using DataProductReleasesPager`, func() {
			listDataProductReleasesOptions := &dpxv1.ListDataProductReleasesOptions{
				DataProductID: &optionalDataProductIdLink,
				// AssetContainerID: core.StringPtr("testString"),
				State: []string{"available"},
				// Version:          core.StringPtr("testString"),
				Limit: core.Int64Ptr(int64(10)),
			}

			// Test GetNext().
			pager, err := dpxService.NewDataProductReleasesPager(listDataProductReleasesOptions)
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
			pager, err = dpxService.NewDataProductReleasesPager(listDataProductReleasesOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			allItems, err := pager.GetAll()
			Expect(err).To(BeNil())
			Expect(allItems).ToNot(BeNil())

			Expect(len(allItems)).To(Equal(len(allResults)))
			fmt.Fprintf(GinkgoWriter, "ListDataProductReleases() returned a total of %d item(s) using DataProductReleasesPager.\n", len(allResults))
		})
	})

	Describe(`RetireDataProductRelease - Retire a release of an existing data product`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`RetireDataProductRelease(retireDataProductReleaseOptions *RetireDataProductReleaseOptions)`, func() {
			retireDataProductReleaseOptions := &dpxv1.RetireDataProductReleaseOptions{
				DataProductID: &optionalDataProductIdLink,
				ReleaseID:     &releaseIdLink,
			}

			dataProductVersion, response, err := dpxService.RetireDataProductRelease(retireDataProductReleaseOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dataProductVersion).ToNot(BeNil())
		})
	})

})

//
// Utility functions are declared in the unit test file
//
