//go:build examples

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
	"encoding/json"
	"fmt"
	"os"

	"github.com/IBM/data-product-exchange-go-sdk/dpxv1"
	"github.com/IBM/go-sdk-core/v5/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//
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
//
var _ = Describe(`DpxV1 Examples Tests`, func() {

	const externalConfigFile = "../dpx_v1.env"

	var (
		dpxService *dpxv1.DpxV1
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

			containerReferenceModel := &dpxv1.ContainerReference{
				ID: core.StringPtr(createDataProductVersionByCatalogIdLink),
			}

			createDataProductVersionOptions := dpxService.NewCreateDataProductVersionOptions(
				containerReferenceModel,
			)
			createDataProductVersionOptions.SetName("My New Data Product")

			dataProductVersion, response, err := dpxService.CreateDataProductVersion(createDataProductVersionOptions)
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
		It(`GetDataProductVersion request example`, func() {
			fmt.Println("\nGetDataProductVersion() result:")
			// begin-get_data_product_version

			getDataProductVersionOptions := dpxService.NewGetDataProductVersionOptions(
				getDataProductVersionByUserIdLink,
			)

			dataProductVersion, response, err := dpxService.GetDataProductVersion(getDataProductVersionOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(dataProductVersion, "", "  ")
			fmt.Println(string(b))

			// end-get_data_product_version

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
		It(`CreateContractTermsDocument request example`, func() {
			fmt.Println("\nCreateContractTermsDocument() result:")
			// begin-create_contract_terms_document

			createContractTermsDocumentOptions := dpxService.NewCreateContractTermsDocumentOptions(
				uploadContractTermsDocumentsByVersionIdLink,
				uploadContractDocumentsByContractIdLink,
				"terms_and_conditions",
				"Terms and conditions document",
				"testString",
			)

			contractTermsDocument, response, err := dpxService.CreateContractTermsDocument(createContractTermsDocumentOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(contractTermsDocument, "", "  ")
			fmt.Println(string(b))

			// end-create_contract_terms_document

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
		It(`GetDataProduct request example`, func() {
			fmt.Println("\nGetDataProduct() result:")
			// begin-get_data_product

			getDataProductOptions := dpxService.NewGetDataProductOptions(
				getDataProductByUserIdLink,
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

			var allResults []dpxv1.DataProduct
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
			listDataProductVersionsOptions := &dpxv1.ListDataProductVersionsOptions{
				AssetContainerID: &getListOfDataProductByCatalogIdLink,
				DataProduct: core.StringPtr("testString"),
				State: core.StringPtr("draft"),
				Version: core.StringPtr("testString"),
				Limit: core.Int64Ptr(int64(10)),
			}

			pager, err := dpxService.NewDataProductVersionsPager(listDataProductVersionsOptions)
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
			// end-list_data_product_versions
		})
		It(`UpdateDataProductVersion request example`, func() {
			fmt.Println("\nUpdateDataProductVersion() result:")
			// begin-update_data_product_version

			jsonPatchOperationModel := &dpxv1.JSONPatchOperation{
				Op: core.StringPtr("add"),
				Path: core.StringPtr("testString"),
			}

			updateDataProductVersionOptions := dpxService.NewUpdateDataProductVersionOptions(
				updateDataProductVersionByUserIdLink,
				[]dpxv1.JSONPatchOperation{*jsonPatchOperationModel},
			)

			dataProductVersion, response, err := dpxService.UpdateDataProductVersion(updateDataProductVersionOptions)
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
		It(`CompleteContractTermsDocument request example`, func() {
			fmt.Println("\nCompleteContractTermsDocument() result:")
			// begin-complete_contract_terms_document

			completeContractTermsDocumentOptions := dpxService.NewCompleteContractTermsDocumentOptions(
				completeContractTermsDocumentByVersionIdLink,
				completeContractTermsDocumentByContractIdLink,
				completeContractTermsDocumentLink,
			)

			contractTermsDocument, response, err := dpxService.CompleteContractTermsDocument(completeContractTermsDocumentOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(contractTermsDocument, "", "  ")
			fmt.Println(string(b))

			// end-complete_contract_terms_document

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(contractTermsDocument).ToNot(BeNil())
		})
		It(`GetContractTermsDocument request example`, func() {
			fmt.Println("\nGetContractTermsDocument() result:")
			// begin-get_contract_terms_document

			getContractTermsDocumentOptions := dpxService.NewGetContractTermsDocumentOptions(
				getContractTermsDocumentByVersionIdLink,
				getContractTermsDocumentsByContractIdLink,
				getContractTermsDocumentByIdLink,
			)

			contractTermsDocument, response, err := dpxService.GetContractTermsDocument(getContractTermsDocumentOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(contractTermsDocument, "", "  ")
			fmt.Println(string(b))

			// end-get_contract_terms_document

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(contractTermsDocument).ToNot(BeNil())
		})
		It(`UpdateContractTermsDocument request example`, func() {
			fmt.Println("\nUpdateContractTermsDocument() result:")
			// begin-update_contract_terms_document

			jsonPatchOperationModel := &dpxv1.JSONPatchOperation{
				Op: core.StringPtr("add"),
				Path: core.StringPtr("testString"),
			}

			updateContractTermsDocumentOptions := dpxService.NewUpdateContractTermsDocumentOptions(
				updateContractTermsDocumentByVersionIdLink,
				updateContractTermsDocumentByContractIdLink,
				updateContractTermsDocumentLink,
				[]dpxv1.JSONPatchOperation{*jsonPatchOperationModel},
			)

			contractTermsDocument, response, err := dpxService.UpdateContractTermsDocument(updateContractTermsDocumentOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(contractTermsDocument, "", "  ")
			fmt.Println(string(b))

			// end-update_contract_terms_document

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(contractTermsDocument).ToNot(BeNil())
		})
		It(`DeleteContractTermsDocument request example`, func() {
			// begin-delete_contract_terms_document

			deleteContractTermsDocumentOptions := dpxService.NewDeleteContractTermsDocumentOptions(
				deleteContractTermsDocumentByVersionIdLink,
				deleteContractTermsDocumentByContractIdLink,
				deleteContractTermsDocumentLink,
			)

			response, err := dpxService.DeleteContractTermsDocument(deleteContractTermsDocumentOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteContractTermsDocument(): %d\n", response.StatusCode)
			}

			// end-delete_contract_terms_document

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`DeleteDataProductVersion request example`, func() {
			// begin-delete_data_product_version

			deleteDataProductVersionOptions := dpxService.NewDeleteDataProductVersionOptions(
				deleteDataProductVersionByUserIdLink,
			)

			response, err := dpxService.DeleteDataProductVersion(deleteDataProductVersionOptions)
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
