//go:build examples

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
	"encoding/json"
	"fmt"
	"os"

	"github.com/IBM/data-product-exchange-go-sdk/dpxv1"
	"github.com/IBM/go-sdk-core/v5/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// This file provides an example of how to use the DPX service.
//
// The following configuration properties are assumed to be defined:
// DPX_URL=<service base url>
// DPX_AUTH_TYPE=iam
// DPX_APIKEY=<IAM apikey>
// DPX_AUTH_URL=<IAM token service base URL - omit this if using the production environment>
//
// These configuration properties can be exported as environment variables, or stored
// in a configuration file and then:
// export IBM_CREDENTIALS_FILE=<name of configuration file>
var _ = Describe(`DpxV1 Examples Tests`, func() {

	const externalConfigFile = "../dpx_v1.env"

	var (
		dpxService *dpxv1.DpxV1
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
			config, err = core.GetServiceProperties(dpxv1.DefaultServiceName)
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

			dpxServiceOptions := &dpxv1.DpxV1Options{}

			dpxService, err = dpxv1.NewDpxV1UsingExternalConfig(dpxServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(dpxService).ToNot(BeNil())
		})
	})

	Describe(`DpxV1 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`Initialize request example`, func() {
			fmt.Println("\nInitialize() result:")
			// begin-initialize

			initializeOptions := dpxService.NewInitializeOptions()
			initializeOptions.SetInclude([]string{"delivery_methods", "data_product_samples", "domains_multi_industry"})

			initializeResource, response, err := dpxService.Initialize(initializeOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(initializeResource, "", "  ")
			fmt.Println(string(b))

			// end-initialize

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(initializeResource).ToNot(BeNil())

			containerIdLink = *initializeResource.Container.ID
			fmt.Fprintf(GinkgoWriter, "Saved containerIdLink value: %v\n", containerIdLink)
		})
		It(`CreateDataProduct request example`, func() {
			fmt.Println("\nCreateDataProduct() result:")
			// begin-create_data_product

			containerReferenceModel := &dpxv1.ContainerReference{
				ID: &containerIdLink,
			}

			assetReferenceModel := &dpxv1.AssetReference{
				Container: containerReferenceModel,
			}

			dataProductVersionPrototypeModel := &dpxv1.DataProductVersionPrototype{
				Name:  core.StringPtr("My New Data Product"),
				Asset: assetReferenceModel,
			}

			createDataProductOptions := dpxService.NewCreateDataProductOptions(
				[]dpxv1.DataProductVersionPrototype{*dataProductVersionPrototypeModel},
			)

			dataProduct, response, err := dpxService.CreateDataProduct(createDataProductOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(dataProduct, "", "  ")
			fmt.Println(string(b))

			// end-create_data_product

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dataProduct).ToNot(BeNil())

			optionalDataProductIdLink = *dataProduct.ID
			fmt.Fprintf(GinkgoWriter, "Saved optionalDataProductIdLink value: %v\n", optionalDataProductIdLink)
			dataProductIdLink = *dataProduct.ID
			fmt.Fprintf(GinkgoWriter, "Saved dataProductIdLink value: %v\n", dataProductIdLink)
		})
		It(`CreateDataProductDraft request example`, func() {
			fmt.Println("\nCreateDataProductDraft() result:")
			// begin-create_data_product_draft

			containerReferenceModel := &dpxv1.ContainerReference{
				ID: &containerIdLink,
			}

			assetReferenceModel := &dpxv1.AssetReference{
				Container: containerReferenceModel,
			}

			dataProductIdentityModel := &dpxv1.DataProductIdentity{
				ID: &dataProductIdLink,
			}

			domainModel := &dpxv1.Domain{
				ID:        core.StringPtr("918c0bfd-6943-4468-b74f-bc111018e0d1"),
				Name:      core.StringPtr("Customer Service"),
				Container: containerReferenceModel,
			}

			createDataProductDraftOptions := dpxService.NewCreateDataProductDraftOptions(
				dataProductIdLink,
				assetReferenceModel,
			)
			createDataProductDraftOptions.SetVersion("1.2.0")
			createDataProductDraftOptions.SetDataProduct(dataProductIdentityModel)
			createDataProductDraftOptions.SetDomain(domainModel)

			dataProductVersion, response, err := dpxService.CreateDataProductDraft(createDataProductDraftOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(dataProductVersion, "", "  ")
			fmt.Println(string(b))

			// end-create_data_product_draft

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(dataProductVersion).ToNot(BeNil())

			draftIdLink = *dataProductVersion.ID
			fmt.Fprintf(GinkgoWriter, "Saved draftIdLink value: %v\n", draftIdLink)
			contractTermsIdLink = *dataProductVersion.ContractTerms[0].ID
			fmt.Fprintf(GinkgoWriter, "Saved contractTermsIdLink value: %v\n", contractTermsIdLink)
		})
		It(`DeleteDataProductDraft request example`, func() {
			// begin-delete_data_product_draft

			deleteDataProductDraftOptions := dpxService.NewDeleteDataProductDraftOptions(
				optionalDataProductIdLink,
				draftIdLink,
			)

			response, err := dpxService.DeleteDataProductDraft(deleteDataProductDraftOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteDataProductDraft(): %d\n", response.StatusCode)
			}

			// end-delete_data_product_draft

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`CreateDataProductDraftAgain request example`, func() {
			fmt.Println("\nCreateDataProductDraft() result:")
			// begin-create_data_product_draft

			containerReferenceModel := &dpxv1.ContainerReference{
				ID: &containerIdLink,
			}

			assetReferenceModel := &dpxv1.AssetReference{
				Container: containerReferenceModel,
			}

			dataProductIdentityModel := &dpxv1.DataProductIdentity{
				ID: &dataProductIdLink,
			}

			domainModel := &dpxv1.Domain{
				ID:        core.StringPtr("918c0bfd-6943-4468-b74f-bc111018e0d1"),
				Name:      core.StringPtr("Customer Service"),
				Container: containerReferenceModel,
			}

			createDataProductDraftOptions := dpxService.NewCreateDataProductDraftOptions(
				dataProductIdLink,
				assetReferenceModel,
			)
			createDataProductDraftOptions.SetVersion("1.2.0")
			createDataProductDraftOptions.SetDataProduct(dataProductIdentityModel)
			createDataProductDraftOptions.SetDomain(domainModel)

			dataProductVersion, response, err := dpxService.CreateDataProductDraft(createDataProductDraftOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(dataProductVersion, "", "  ")
			fmt.Println(string(b))

			// end-create_data_product_draft

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(dataProductVersion).ToNot(BeNil())

			draftIdLink = *dataProductVersion.ID
			fmt.Fprintf(GinkgoWriter, "Saved draftIdLink value: %v\n", draftIdLink)
			contractTermsIdLink = *dataProductVersion.ContractTerms[0].ID
			fmt.Fprintf(GinkgoWriter, "Saved contractTermsIdLink value: %v\n", contractTermsIdLink)
		})
		It(`CreateDraftContractTermsDocument request example`, func() {
			fmt.Println("\nCreateDraftContractTermsDocument() result:")
			// begin-create_draft_contract_terms_document

			createDraftContractTermsDocumentOptions := dpxService.NewCreateDraftContractTermsDocumentOptions(
				optionalDataProductIdLink,
				draftIdLink,
				contractTermsIdLink,
				"terms_and_conditions",
				"Terms and conditions document",
				"b38df608-d34b-4d58-8136-ed25e6c6684e",
				"https://www.google.com",
			)

			contractTermsDocument, response, err := dpxService.CreateDraftContractTermsDocument(createDraftContractTermsDocumentOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(contractTermsDocument, "", "  ")
			fmt.Println(string(b))

			// end-create_draft_contract_terms_document

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(contractTermsDocument).ToNot(BeNil())

			documentIdLink = *contractTermsDocument.ID
			fmt.Fprintf(GinkgoWriter, "Saved documentIdLink value: %v\n", documentIdLink)
		})
		It(`DeleteDraftContractTermsDocument request example`, func() {
			// begin-delete_draft_contract_terms_document

			deleteDraftContractTermsDocumentOptions := dpxService.NewDeleteDraftContractTermsDocumentOptions(
				optionalDataProductIdLink,
				draftIdLink,
				contractTermsIdLink,
				documentIdLink,
			)

			response, err := dpxService.DeleteDraftContractTermsDocument(deleteDraftContractTermsDocumentOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteDraftContractTermsDocument(): %d\n", response.StatusCode)
			}

			// end-delete_draft_contract_terms_document

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`CreateDraftContractTermsDocumentAgain request example`, func() {
			fmt.Println("\nCreateDraftContractTermsDocument() result:")
			// begin-create_draft_contract_terms_document

			createDraftContractTermsDocumentOptions := dpxService.NewCreateDraftContractTermsDocumentOptions(
				optionalDataProductIdLink,
				draftIdLink,
				contractTermsIdLink,
				"terms_and_conditions",
				"Terms and conditions document",
				"b38df608-d34b-4d58-8136-ed25e6c6684e",
				"https://www.google.com",
			)

			contractTermsDocument, response, err := dpxService.CreateDraftContractTermsDocument(createDraftContractTermsDocumentOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(contractTermsDocument, "", "  ")
			fmt.Println(string(b))

			// end-create_draft_contract_terms_document

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(contractTermsDocument).ToNot(BeNil())

			documentIdLink = *contractTermsDocument.ID
			fmt.Fprintf(GinkgoWriter, "Saved documentIdLink value: %v\n", documentIdLink)
		})
		It(`UpdateDataProductDraft request example`, func() {
			fmt.Println("\nUpdateDataProductDraft() result:")
			// begin-update_data_product_draft

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
				Op:    core.StringPtr("add"),
				Path:  core.StringPtr("/parts_out"),
				Value: partsOutList,
			}

			updateDataProductDraftOptions := dpxService.NewUpdateDataProductDraftOptions(
				optionalDataProductIdLink,
				draftIdLink,
				[]dpxv1.JSONPatchOperation{*jsonPatchOperationModel},
			)

			dataProductVersion, response, err := dpxService.UpdateDataProductDraft(updateDataProductDraftOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(dataProductVersion, "", "  ")
			fmt.Println(string(b))

			// end-update_data_product_draft

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dataProductVersion).ToNot(BeNil())
		})
		It(`GetDataProductDraft request example`, func() {
			fmt.Println("\nGetDataProductDraft() result:")
			// begin-get_data_product_draft

			getDataProductDraftOptions := dpxService.NewGetDataProductDraftOptions(
				optionalDataProductIdLink,
				draftIdLink,
			)

			dataProductVersion, response, err := dpxService.GetDataProductDraft(getDataProductDraftOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(dataProductVersion, "", "  ")
			fmt.Println(string(b))

			// end-get_data_product_draft

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dataProductVersion).ToNot(BeNil())
		})
		It(`UpdateDraftContractTermsDocument request example`, func() {
			fmt.Println("\nUpdateDraftContractTermsDocument() result:")
			// begin-update_draft_contract_terms_document

			jsonPatchOperationModel := &dpxv1.JSONPatchOperation{
				Op:    core.StringPtr("replace"),
				Path:  core.StringPtr("/url"),
				Value: "https://google.com",
			}

			updateDraftContractTermsDocumentOptions := dpxService.NewUpdateDraftContractTermsDocumentOptions(
				optionalDataProductIdLink,
				draftIdLink,
				contractTermsIdLink,
				documentIdLink,
				[]dpxv1.JSONPatchOperation{*jsonPatchOperationModel},
			)

			contractTermsDocument, response, err := dpxService.UpdateDraftContractTermsDocument(updateDraftContractTermsDocumentOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(contractTermsDocument, "", "  ")
			fmt.Println(string(b))

			// end-update_draft_contract_terms_document

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(contractTermsDocument).ToNot(BeNil())
		})
		It(`GetInitializeStatus request example`, func() {
			fmt.Println("\nGetInitializeStatus() result:")
			// begin-get_initialize_status

			getInitializeStatusOptions := dpxService.NewGetInitializeStatusOptions()

			initializeResource, response, err := dpxService.GetInitializeStatus(getInitializeStatusOptions)
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
		It(`GetDraftContractTermsDocument request example`, func() {
			fmt.Println("\nGetDraftContractTermsDocument() result:")
			// begin-get_draft_contract_terms_document

			getDraftContractTermsDocumentOptions := dpxService.NewGetDraftContractTermsDocumentOptions(
				optionalDataProductIdLink,
				draftIdLink,
				contractTermsIdLink,
				documentIdLink,
			)

			contractTermsDocument, response, err := dpxService.GetDraftContractTermsDocument(getDraftContractTermsDocumentOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(contractTermsDocument, "", "  ")
			fmt.Println(string(b))

			// end-get_draft_contract_terms_document

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(contractTermsDocument).ToNot(BeNil())
		})
		It(`PublishDataProductDraft request example`, func() {
			fmt.Println("\nPublishDataProductDraft() result:")
			// begin-publish_data_product_draft

			publishDataProductDraftOptions := dpxService.NewPublishDataProductDraftOptions(
				optionalDataProductIdLink,
				draftIdLink,
			)

			dataProductVersion, response, err := dpxService.PublishDataProductDraft(publishDataProductDraftOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(dataProductVersion, "", "  ")
			fmt.Println(string(b))

			// end-publish_data_product_draft

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dataProductVersion).ToNot(BeNil())

			releaseIdLink = *dataProductVersion.ID
			fmt.Fprintf(GinkgoWriter, "Saved releaseIdLink value: %v\n", releaseIdLink)
		})
		It(`ManageApiKeys request example`, func() {
			// begin-manage_api_keys

			manageApiKeysOptions := dpxService.NewManageApiKeysOptions()

			response, err := dpxService.ManageApiKeys(manageApiKeysOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from ManageApiKeys(): %d\n", response.StatusCode)
			}

			// end-manage_api_keys

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`ListDataProducts request example`, func() {
			fmt.Println("\nListDataProducts() result:")
			// begin-list_data_products
			listDataProductsOptions := &dpxv1.ListDataProductsOptions{
				Limit: core.Int64Ptr(int64(10)),
			}

			pager, err := dpxService.NewDataProductsPager(listDataProductsOptions)
			if err != nil {
				panic(err)
			}

			var allResults []dpxv1.DataProductSummary
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
		It(`GetDataProduct request example`, func() {
			fmt.Println("\nGetDataProduct() result:")
			// begin-get_data_product

			getDataProductOptions := dpxService.NewGetDataProductOptions(
				dataProductIdLink,
			)

			dataProduct, response, err := dpxService.GetDataProduct(getDataProductOptions)
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
		It(`CompleteDraftContractTermsDocument request example`, func() {
			fmt.Println("\nCompleteDraftContractTermsDocument() result:")
			// begin-complete_draft_contract_terms_document

			completeDraftContractTermsDocumentOptions := dpxService.NewCompleteDraftContractTermsDocumentOptions(
				optionalDataProductIdLink,
				draftIdLink,
				contractTermsIdLink,
				documentIdLink,
			)

			contractTermsDocument, response, err := dpxService.CompleteDraftContractTermsDocument(completeDraftContractTermsDocumentOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(contractTermsDocument, "", "  ")
			fmt.Println(string(b))

			// end-complete_draft_contract_terms_document

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(contractTermsDocument).ToNot(BeNil())
		})
		It(`ListDataProductDrafts request example`, func() {
			fmt.Println("\nListDataProductDrafts() result:")
			// begin-list_data_product_drafts
			listDataProductDraftsOptions := &dpxv1.ListDataProductDraftsOptions{
				DataProductID: &optionalDataProductIdLink,
				Limit:         core.Int64Ptr(int64(10)),
			}

			pager, err := dpxService.NewDataProductDraftsPager(listDataProductDraftsOptions)
			if err != nil {
				panic(err)
			}

			var allResults []dpxv1.DataProductVersionSummary
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				if err != nil {
					panic(err)
				}
				allResults = append(allResults, nextPage...)
			}
			b, _ := json.MarshalIndent(allResults, "", "  ")
			fmt.Println(string(b))
			// end-list_data_product_drafts
		})
		It(`GetDataProductRelease request example`, func() {
			fmt.Println("\nGetDataProductRelease() result:")
			// begin-get_data_product_release

			getDataProductReleaseOptions := dpxService.NewGetDataProductReleaseOptions(
				optionalDataProductIdLink,
				releaseIdLink,
			)

			dataProductVersion, response, err := dpxService.GetDataProductRelease(getDataProductReleaseOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(dataProductVersion, "", "  ")
			fmt.Println(string(b))

			// end-get_data_product_release

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dataProductVersion).ToNot(BeNil())
		})
		It(`UpdateDataProductRelease request example`, func() {
			fmt.Println("\nUpdateDataProductRelease() result:")
			// begin-update_data_product_release

			jsonPatchOperationModel := &dpxv1.JSONPatchOperation{
				Op:    core.StringPtr("replace"),
				Path:  core.StringPtr("/description"),
				Value: "New Description",
			}

			updateDataProductReleaseOptions := dpxService.NewUpdateDataProductReleaseOptions(
				optionalDataProductIdLink,
				releaseIdLink,
				[]dpxv1.JSONPatchOperation{*jsonPatchOperationModel},
			)

			dataProductVersion, response, err := dpxService.UpdateDataProductRelease(updateDataProductReleaseOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(dataProductVersion, "", "  ")
			fmt.Println(string(b))

			// end-update_data_product_release

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dataProductVersion).ToNot(BeNil())
		})
		It(`GetReleaseContractTermsDocument request example`, func() {
			fmt.Println("\nGetReleaseContractTermsDocument() result:")
			// begin-get_release_contract_terms_document

			getReleaseContractTermsDocumentOptions := dpxService.NewGetReleaseContractTermsDocumentOptions(
				optionalDataProductIdLink,
				releaseIdLink,
				contractTermsIdLink,
				documentIdLink,
			)

			contractTermsDocument, response, err := dpxService.GetReleaseContractTermsDocument(getReleaseContractTermsDocumentOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(contractTermsDocument, "", "  ")
			fmt.Println(string(b))

			// end-get_release_contract_terms_document

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(contractTermsDocument).ToNot(BeNil())
		})
		It(`ListDataProductReleases request example`, func() {
			fmt.Println("\nListDataProductReleases() result:")
			// begin-list_data_product_releases
			listDataProductReleasesOptions := &dpxv1.ListDataProductReleasesOptions{
				DataProductID: &optionalDataProductIdLink,
				State:         []string{"available"},
				Limit:         core.Int64Ptr(int64(10)),
			}

			pager, err := dpxService.NewDataProductReleasesPager(listDataProductReleasesOptions)
			if err != nil {
				panic(err)
			}

			var allResults []dpxv1.DataProductVersionSummary
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				if err != nil {
					panic(err)
				}
				allResults = append(allResults, nextPage...)
			}
			b, _ := json.MarshalIndent(allResults, "", "  ")
			fmt.Println(string(b))
			// end-list_data_product_releases
		})
		It(`RetireDataProductRelease request example`, func() {
			fmt.Println("\nRetireDataProductRelease() result:")
			// begin-retire_data_product_release

			retireDataProductReleaseOptions := dpxService.NewRetireDataProductReleaseOptions(
				optionalDataProductIdLink,
				releaseIdLink,
			)

			dataProductVersion, response, err := dpxService.RetireDataProductRelease(retireDataProductReleaseOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(dataProductVersion, "", "  ")
			fmt.Println(string(b))

			// end-retire_data_product_release

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dataProductVersion).ToNot(BeNil())
		})
	})
})
