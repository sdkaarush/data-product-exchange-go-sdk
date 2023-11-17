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
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gstruct"
	"github.ibm.com/wdp-gov/data-product-go-sdk/dataproductexchangeapiservicev1"
)

var _ = Describe(`DataProductExchangeApiServiceV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			dataProductExchangeApiServiceService, serviceErr := dataproductexchangeapiservicev1.NewDataProductExchangeApiServiceV1(&dataproductexchangeapiservicev1.DataProductExchangeApiServiceV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(dataProductExchangeApiServiceService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			dataProductExchangeApiServiceService, serviceErr := dataproductexchangeapiservicev1.NewDataProductExchangeApiServiceV1(&dataproductexchangeapiservicev1.DataProductExchangeApiServiceV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(dataProductExchangeApiServiceService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			dataProductExchangeApiServiceService, serviceErr := dataproductexchangeapiservicev1.NewDataProductExchangeApiServiceV1(&dataproductexchangeapiservicev1.DataProductExchangeApiServiceV1Options{
				URL: "https://dataproductexchangeapiservicev1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(dataProductExchangeApiServiceService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"DATA_PRODUCT_EXCHANGE_API_SERVICE_URL": "https://dataproductexchangeapiservicev1/api",
				"DATA_PRODUCT_EXCHANGE_API_SERVICE_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				dataProductExchangeApiServiceService, serviceErr := dataproductexchangeapiservicev1.NewDataProductExchangeApiServiceV1UsingExternalConfig(&dataproductexchangeapiservicev1.DataProductExchangeApiServiceV1Options{
				})
				Expect(dataProductExchangeApiServiceService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := dataProductExchangeApiServiceService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != dataProductExchangeApiServiceService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(dataProductExchangeApiServiceService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(dataProductExchangeApiServiceService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				dataProductExchangeApiServiceService, serviceErr := dataproductexchangeapiservicev1.NewDataProductExchangeApiServiceV1UsingExternalConfig(&dataproductexchangeapiservicev1.DataProductExchangeApiServiceV1Options{
					URL: "https://testService/api",
				})
				Expect(dataProductExchangeApiServiceService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(dataProductExchangeApiServiceService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := dataProductExchangeApiServiceService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != dataProductExchangeApiServiceService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(dataProductExchangeApiServiceService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(dataProductExchangeApiServiceService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				dataProductExchangeApiServiceService, serviceErr := dataproductexchangeapiservicev1.NewDataProductExchangeApiServiceV1UsingExternalConfig(&dataproductexchangeapiservicev1.DataProductExchangeApiServiceV1Options{
				})
				err := dataProductExchangeApiServiceService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(dataProductExchangeApiServiceService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(dataProductExchangeApiServiceService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := dataProductExchangeApiServiceService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != dataProductExchangeApiServiceService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(dataProductExchangeApiServiceService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(dataProductExchangeApiServiceService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"DATA_PRODUCT_EXCHANGE_API_SERVICE_URL": "https://dataproductexchangeapiservicev1/api",
				"DATA_PRODUCT_EXCHANGE_API_SERVICE_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			dataProductExchangeApiServiceService, serviceErr := dataproductexchangeapiservicev1.NewDataProductExchangeApiServiceV1UsingExternalConfig(&dataproductexchangeapiservicev1.DataProductExchangeApiServiceV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(dataProductExchangeApiServiceService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"DATA_PRODUCT_EXCHANGE_API_SERVICE_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			dataProductExchangeApiServiceService, serviceErr := dataproductexchangeapiservicev1.NewDataProductExchangeApiServiceV1UsingExternalConfig(&dataproductexchangeapiservicev1.DataProductExchangeApiServiceV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(dataProductExchangeApiServiceService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = dataproductexchangeapiservicev1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`GetInitializeStatus(getInitializeStatusOptions *GetInitializeStatusOptions) - Operation response error`, func() {
		getInitializeStatusPath := "/data_product_exchange/v1/configuration/initialize/status"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getInitializeStatusPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["container.id"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetInitializeStatus with error: Operation response processing error`, func() {
				dataProductExchangeApiServiceService, serviceErr := dataproductexchangeapiservicev1.NewDataProductExchangeApiServiceV1(&dataproductexchangeapiservicev1.DataProductExchangeApiServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductExchangeApiServiceService).ToNot(BeNil())

				// Construct an instance of the GetInitializeStatusOptions model
				getInitializeStatusOptionsModel := new(dataproductexchangeapiservicev1.GetInitializeStatusOptions)
				getInitializeStatusOptionsModel.ContainerID = core.StringPtr("testString")
				getInitializeStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := dataProductExchangeApiServiceService.GetInitializeStatus(getInitializeStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				dataProductExchangeApiServiceService.EnableRetries(0, 0)
				result, response, operationErr = dataProductExchangeApiServiceService.GetInitializeStatus(getInitializeStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetInitializeStatus(getInitializeStatusOptions *GetInitializeStatusOptions)`, func() {
		getInitializeStatusPath := "/data_product_exchange/v1/configuration/initialize/status"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getInitializeStatusPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["container.id"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}, "href": "http://api.example.com/configuration/initialize/status?catalog_id=d29c42eb-7100-4b7a-8257-c196dbcca1cd", "status": "not_started", "trace": "Trace", "errors": [{"code": "Code", "target": {"type": "field", "name": "Name"}, "message": "Message", "more_info": "MoreInfo"}], "last_started_at": "2023-08-21T15:24:06.021Z", "last_finished_at": "2023-08-21T20:24:34.450Z", "initialized_options": [{"name": "Name", "version": 1}]}`)
				}))
			})
			It(`Invoke GetInitializeStatus successfully with retries`, func() {
				dataProductExchangeApiServiceService, serviceErr := dataproductexchangeapiservicev1.NewDataProductExchangeApiServiceV1(&dataproductexchangeapiservicev1.DataProductExchangeApiServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductExchangeApiServiceService).ToNot(BeNil())
				dataProductExchangeApiServiceService.EnableRetries(0, 0)

				// Construct an instance of the GetInitializeStatusOptions model
				getInitializeStatusOptionsModel := new(dataproductexchangeapiservicev1.GetInitializeStatusOptions)
				getInitializeStatusOptionsModel.ContainerID = core.StringPtr("testString")
				getInitializeStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := dataProductExchangeApiServiceService.GetInitializeStatusWithContext(ctx, getInitializeStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				dataProductExchangeApiServiceService.DisableRetries()
				result, response, operationErr := dataProductExchangeApiServiceService.GetInitializeStatus(getInitializeStatusOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = dataProductExchangeApiServiceService.GetInitializeStatusWithContext(ctx, getInitializeStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getInitializeStatusPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["container.id"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}, "href": "http://api.example.com/configuration/initialize/status?catalog_id=d29c42eb-7100-4b7a-8257-c196dbcca1cd", "status": "not_started", "trace": "Trace", "errors": [{"code": "Code", "target": {"type": "field", "name": "Name"}, "message": "Message", "more_info": "MoreInfo"}], "last_started_at": "2023-08-21T15:24:06.021Z", "last_finished_at": "2023-08-21T20:24:34.450Z", "initialized_options": [{"name": "Name", "version": 1}]}`)
				}))
			})
			It(`Invoke GetInitializeStatus successfully`, func() {
				dataProductExchangeApiServiceService, serviceErr := dataproductexchangeapiservicev1.NewDataProductExchangeApiServiceV1(&dataproductexchangeapiservicev1.DataProductExchangeApiServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductExchangeApiServiceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := dataProductExchangeApiServiceService.GetInitializeStatus(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetInitializeStatusOptions model
				getInitializeStatusOptionsModel := new(dataproductexchangeapiservicev1.GetInitializeStatusOptions)
				getInitializeStatusOptionsModel.ContainerID = core.StringPtr("testString")
				getInitializeStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = dataProductExchangeApiServiceService.GetInitializeStatus(getInitializeStatusOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetInitializeStatus with error: Operation request error`, func() {
				dataProductExchangeApiServiceService, serviceErr := dataproductexchangeapiservicev1.NewDataProductExchangeApiServiceV1(&dataproductexchangeapiservicev1.DataProductExchangeApiServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductExchangeApiServiceService).ToNot(BeNil())

				// Construct an instance of the GetInitializeStatusOptions model
				getInitializeStatusOptionsModel := new(dataproductexchangeapiservicev1.GetInitializeStatusOptions)
				getInitializeStatusOptionsModel.ContainerID = core.StringPtr("testString")
				getInitializeStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dataProductExchangeApiServiceService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := dataProductExchangeApiServiceService.GetInitializeStatus(getInitializeStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetInitializeStatus successfully`, func() {
				dataProductExchangeApiServiceService, serviceErr := dataproductexchangeapiservicev1.NewDataProductExchangeApiServiceV1(&dataproductexchangeapiservicev1.DataProductExchangeApiServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductExchangeApiServiceService).ToNot(BeNil())

				// Construct an instance of the GetInitializeStatusOptions model
				getInitializeStatusOptionsModel := new(dataproductexchangeapiservicev1.GetInitializeStatusOptions)
				getInitializeStatusOptionsModel.ContainerID = core.StringPtr("testString")
				getInitializeStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := dataProductExchangeApiServiceService.GetInitializeStatus(getInitializeStatusOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Initialize(initializeOptions *InitializeOptions) - Operation response error`, func() {
		initializePath := "/data_product_exchange/v1/configuration/initialize"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(initializePath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke Initialize with error: Operation response processing error`, func() {
				dataProductExchangeApiServiceService, serviceErr := dataproductexchangeapiservicev1.NewDataProductExchangeApiServiceV1(&dataproductexchangeapiservicev1.DataProductExchangeApiServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductExchangeApiServiceService).ToNot(BeNil())

				// Construct an instance of the ContainerReference model
				containerReferenceModel := new(dataproductexchangeapiservicev1.ContainerReference)
				containerReferenceModel.ID = core.StringPtr("d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				containerReferenceModel.Type = core.StringPtr("catalog")

				// Construct an instance of the InitializeOptions model
				initializeOptionsModel := new(dataproductexchangeapiservicev1.InitializeOptions)
				initializeOptionsModel.Container = containerReferenceModel
				initializeOptionsModel.Include = []string{"delivery_methods", "data_product_samples", "domains_multi_industry"}
				initializeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := dataProductExchangeApiServiceService.Initialize(initializeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				dataProductExchangeApiServiceService.EnableRetries(0, 0)
				result, response, operationErr = dataProductExchangeApiServiceService.Initialize(initializeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Initialize(initializeOptions *InitializeOptions)`, func() {
		initializePath := "/data_product_exchange/v1/configuration/initialize"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(initializePath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}, "href": "http://api.example.com/configuration/initialize/status?catalog_id=d29c42eb-7100-4b7a-8257-c196dbcca1cd", "status": "not_started", "trace": "Trace", "errors": [{"code": "Code", "target": {"type": "field", "name": "Name"}, "message": "Message", "more_info": "MoreInfo"}], "last_started_at": "2023-08-21T15:24:06.021Z", "last_finished_at": "2023-08-21T20:24:34.450Z", "initialized_options": [{"name": "Name", "version": 1}]}`)
				}))
			})
			It(`Invoke Initialize successfully with retries`, func() {
				dataProductExchangeApiServiceService, serviceErr := dataproductexchangeapiservicev1.NewDataProductExchangeApiServiceV1(&dataproductexchangeapiservicev1.DataProductExchangeApiServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductExchangeApiServiceService).ToNot(BeNil())
				dataProductExchangeApiServiceService.EnableRetries(0, 0)

				// Construct an instance of the ContainerReference model
				containerReferenceModel := new(dataproductexchangeapiservicev1.ContainerReference)
				containerReferenceModel.ID = core.StringPtr("d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				containerReferenceModel.Type = core.StringPtr("catalog")

				// Construct an instance of the InitializeOptions model
				initializeOptionsModel := new(dataproductexchangeapiservicev1.InitializeOptions)
				initializeOptionsModel.Container = containerReferenceModel
				initializeOptionsModel.Include = []string{"delivery_methods", "data_product_samples", "domains_multi_industry"}
				initializeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := dataProductExchangeApiServiceService.InitializeWithContext(ctx, initializeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				dataProductExchangeApiServiceService.DisableRetries()
				result, response, operationErr := dataProductExchangeApiServiceService.Initialize(initializeOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = dataProductExchangeApiServiceService.InitializeWithContext(ctx, initializeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(initializePath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}, "href": "http://api.example.com/configuration/initialize/status?catalog_id=d29c42eb-7100-4b7a-8257-c196dbcca1cd", "status": "not_started", "trace": "Trace", "errors": [{"code": "Code", "target": {"type": "field", "name": "Name"}, "message": "Message", "more_info": "MoreInfo"}], "last_started_at": "2023-08-21T15:24:06.021Z", "last_finished_at": "2023-08-21T20:24:34.450Z", "initialized_options": [{"name": "Name", "version": 1}]}`)
				}))
			})
			It(`Invoke Initialize successfully`, func() {
				dataProductExchangeApiServiceService, serviceErr := dataproductexchangeapiservicev1.NewDataProductExchangeApiServiceV1(&dataproductexchangeapiservicev1.DataProductExchangeApiServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductExchangeApiServiceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := dataProductExchangeApiServiceService.Initialize(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ContainerReference model
				containerReferenceModel := new(dataproductexchangeapiservicev1.ContainerReference)
				containerReferenceModel.ID = core.StringPtr("d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				containerReferenceModel.Type = core.StringPtr("catalog")

				// Construct an instance of the InitializeOptions model
				initializeOptionsModel := new(dataproductexchangeapiservicev1.InitializeOptions)
				initializeOptionsModel.Container = containerReferenceModel
				initializeOptionsModel.Include = []string{"delivery_methods", "data_product_samples", "domains_multi_industry"}
				initializeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = dataProductExchangeApiServiceService.Initialize(initializeOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke Initialize with error: Operation request error`, func() {
				dataProductExchangeApiServiceService, serviceErr := dataproductexchangeapiservicev1.NewDataProductExchangeApiServiceV1(&dataproductexchangeapiservicev1.DataProductExchangeApiServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductExchangeApiServiceService).ToNot(BeNil())

				// Construct an instance of the ContainerReference model
				containerReferenceModel := new(dataproductexchangeapiservicev1.ContainerReference)
				containerReferenceModel.ID = core.StringPtr("d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				containerReferenceModel.Type = core.StringPtr("catalog")

				// Construct an instance of the InitializeOptions model
				initializeOptionsModel := new(dataproductexchangeapiservicev1.InitializeOptions)
				initializeOptionsModel.Container = containerReferenceModel
				initializeOptionsModel.Include = []string{"delivery_methods", "data_product_samples", "domains_multi_industry"}
				initializeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dataProductExchangeApiServiceService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := dataProductExchangeApiServiceService.Initialize(initializeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(202)
				}))
			})
			It(`Invoke Initialize successfully`, func() {
				dataProductExchangeApiServiceService, serviceErr := dataproductexchangeapiservicev1.NewDataProductExchangeApiServiceV1(&dataproductexchangeapiservicev1.DataProductExchangeApiServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductExchangeApiServiceService).ToNot(BeNil())

				// Construct an instance of the ContainerReference model
				containerReferenceModel := new(dataproductexchangeapiservicev1.ContainerReference)
				containerReferenceModel.ID = core.StringPtr("d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				containerReferenceModel.Type = core.StringPtr("catalog")

				// Construct an instance of the InitializeOptions model
				initializeOptionsModel := new(dataproductexchangeapiservicev1.InitializeOptions)
				initializeOptionsModel.Container = containerReferenceModel
				initializeOptionsModel.Include = []string{"delivery_methods", "data_product_samples", "domains_multi_industry"}
				initializeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := dataProductExchangeApiServiceService.Initialize(initializeOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetDataProduct(getDataProductOptions *GetDataProductOptions) - Operation response error`, func() {
		getDataProductPath := "/data_product_exchange/v1/data_products/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDataProductPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetDataProduct with error: Operation response processing error`, func() {
				dataProductExchangeApiServiceService, serviceErr := dataproductexchangeapiservicev1.NewDataProductExchangeApiServiceV1(&dataproductexchangeapiservicev1.DataProductExchangeApiServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductExchangeApiServiceService).ToNot(BeNil())

				// Construct an instance of the GetDataProductOptions model
				getDataProductOptionsModel := new(dataproductexchangeapiservicev1.GetDataProductOptions)
				getDataProductOptionsModel.ID = core.StringPtr("testString")
				getDataProductOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := dataProductExchangeApiServiceService.GetDataProduct(getDataProductOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				dataProductExchangeApiServiceService.EnableRetries(0, 0)
				result, response, operationErr = dataProductExchangeApiServiceService.GetDataProduct(getDataProductOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetDataProduct(getDataProductOptions *GetDataProductOptions)`, func() {
		getDataProductPath := "/data_product_exchange/v1/data_products/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDataProductPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "b38df608-d34b-4d58-8136-ed25e6c6684e", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}, "name": "Sample Data Product"}`)
				}))
			})
			It(`Invoke GetDataProduct successfully with retries`, func() {
				dataProductExchangeApiServiceService, serviceErr := dataproductexchangeapiservicev1.NewDataProductExchangeApiServiceV1(&dataproductexchangeapiservicev1.DataProductExchangeApiServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductExchangeApiServiceService).ToNot(BeNil())
				dataProductExchangeApiServiceService.EnableRetries(0, 0)

				// Construct an instance of the GetDataProductOptions model
				getDataProductOptionsModel := new(dataproductexchangeapiservicev1.GetDataProductOptions)
				getDataProductOptionsModel.ID = core.StringPtr("testString")
				getDataProductOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := dataProductExchangeApiServiceService.GetDataProductWithContext(ctx, getDataProductOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				dataProductExchangeApiServiceService.DisableRetries()
				result, response, operationErr := dataProductExchangeApiServiceService.GetDataProduct(getDataProductOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = dataProductExchangeApiServiceService.GetDataProductWithContext(ctx, getDataProductOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDataProductPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "b38df608-d34b-4d58-8136-ed25e6c6684e", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}, "name": "Sample Data Product"}`)
				}))
			})
			It(`Invoke GetDataProduct successfully`, func() {
				dataProductExchangeApiServiceService, serviceErr := dataproductexchangeapiservicev1.NewDataProductExchangeApiServiceV1(&dataproductexchangeapiservicev1.DataProductExchangeApiServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductExchangeApiServiceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := dataProductExchangeApiServiceService.GetDataProduct(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetDataProductOptions model
				getDataProductOptionsModel := new(dataproductexchangeapiservicev1.GetDataProductOptions)
				getDataProductOptionsModel.ID = core.StringPtr("testString")
				getDataProductOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = dataProductExchangeApiServiceService.GetDataProduct(getDataProductOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetDataProduct with error: Operation validation and request error`, func() {
				dataProductExchangeApiServiceService, serviceErr := dataproductexchangeapiservicev1.NewDataProductExchangeApiServiceV1(&dataproductexchangeapiservicev1.DataProductExchangeApiServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductExchangeApiServiceService).ToNot(BeNil())

				// Construct an instance of the GetDataProductOptions model
				getDataProductOptionsModel := new(dataproductexchangeapiservicev1.GetDataProductOptions)
				getDataProductOptionsModel.ID = core.StringPtr("testString")
				getDataProductOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dataProductExchangeApiServiceService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := dataProductExchangeApiServiceService.GetDataProduct(getDataProductOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetDataProductOptions model with no property values
				getDataProductOptionsModelNew := new(dataproductexchangeapiservicev1.GetDataProductOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = dataProductExchangeApiServiceService.GetDataProduct(getDataProductOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetDataProduct successfully`, func() {
				dataProductExchangeApiServiceService, serviceErr := dataproductexchangeapiservicev1.NewDataProductExchangeApiServiceV1(&dataproductexchangeapiservicev1.DataProductExchangeApiServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductExchangeApiServiceService).ToNot(BeNil())

				// Construct an instance of the GetDataProductOptions model
				getDataProductOptionsModel := new(dataproductexchangeapiservicev1.GetDataProductOptions)
				getDataProductOptionsModel.ID = core.StringPtr("testString")
				getDataProductOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := dataProductExchangeApiServiceService.GetDataProduct(getDataProductOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListDataProducts(listDataProductsOptions *ListDataProductsOptions) - Operation response error`, func() {
		listDataProductsPath := "/data_product_exchange/v1/data_products"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listDataProductsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListDataProducts with error: Operation response processing error`, func() {
				dataProductExchangeApiServiceService, serviceErr := dataproductexchangeapiservicev1.NewDataProductExchangeApiServiceV1(&dataproductexchangeapiservicev1.DataProductExchangeApiServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductExchangeApiServiceService).ToNot(BeNil())

				// Construct an instance of the ListDataProductsOptions model
				listDataProductsOptionsModel := new(dataproductexchangeapiservicev1.ListDataProductsOptions)
				listDataProductsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listDataProductsOptionsModel.Start = core.StringPtr("testString")
				listDataProductsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := dataProductExchangeApiServiceService.ListDataProducts(listDataProductsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				dataProductExchangeApiServiceService.EnableRetries(0, 0)
				result, response, operationErr = dataProductExchangeApiServiceService.ListDataProducts(listDataProductsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListDataProducts(listDataProductsOptions *ListDataProductsOptions)`, func() {
		listDataProductsPath := "/data_product_exchange/v1/data_products"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listDataProductsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"limit": 200, "first": {"href": "http://api.example.com/collection"}, "next": {"href": "http://api.example.com/collection?start=eyJvZmZzZXQiOjAsImRvbmUiOnRydWV9", "start": "eyJvZmZzZXQiOjAsImRvbmUiOnRydWV9"}, "data_products": [{"id": "b38df608-d34b-4d58-8136-ed25e6c6684e", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}, "name": "Sample Data Product"}]}`)
				}))
			})
			It(`Invoke ListDataProducts successfully with retries`, func() {
				dataProductExchangeApiServiceService, serviceErr := dataproductexchangeapiservicev1.NewDataProductExchangeApiServiceV1(&dataproductexchangeapiservicev1.DataProductExchangeApiServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductExchangeApiServiceService).ToNot(BeNil())
				dataProductExchangeApiServiceService.EnableRetries(0, 0)

				// Construct an instance of the ListDataProductsOptions model
				listDataProductsOptionsModel := new(dataproductexchangeapiservicev1.ListDataProductsOptions)
				listDataProductsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listDataProductsOptionsModel.Start = core.StringPtr("testString")
				listDataProductsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := dataProductExchangeApiServiceService.ListDataProductsWithContext(ctx, listDataProductsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				dataProductExchangeApiServiceService.DisableRetries()
				result, response, operationErr := dataProductExchangeApiServiceService.ListDataProducts(listDataProductsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = dataProductExchangeApiServiceService.ListDataProductsWithContext(ctx, listDataProductsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listDataProductsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"limit": 200, "first": {"href": "http://api.example.com/collection"}, "next": {"href": "http://api.example.com/collection?start=eyJvZmZzZXQiOjAsImRvbmUiOnRydWV9", "start": "eyJvZmZzZXQiOjAsImRvbmUiOnRydWV9"}, "data_products": [{"id": "b38df608-d34b-4d58-8136-ed25e6c6684e", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}, "name": "Sample Data Product"}]}`)
				}))
			})
			It(`Invoke ListDataProducts successfully`, func() {
				dataProductExchangeApiServiceService, serviceErr := dataproductexchangeapiservicev1.NewDataProductExchangeApiServiceV1(&dataproductexchangeapiservicev1.DataProductExchangeApiServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductExchangeApiServiceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := dataProductExchangeApiServiceService.ListDataProducts(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListDataProductsOptions model
				listDataProductsOptionsModel := new(dataproductexchangeapiservicev1.ListDataProductsOptions)
				listDataProductsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listDataProductsOptionsModel.Start = core.StringPtr("testString")
				listDataProductsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = dataProductExchangeApiServiceService.ListDataProducts(listDataProductsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListDataProducts with error: Operation request error`, func() {
				dataProductExchangeApiServiceService, serviceErr := dataproductexchangeapiservicev1.NewDataProductExchangeApiServiceV1(&dataproductexchangeapiservicev1.DataProductExchangeApiServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductExchangeApiServiceService).ToNot(BeNil())

				// Construct an instance of the ListDataProductsOptions model
				listDataProductsOptionsModel := new(dataproductexchangeapiservicev1.ListDataProductsOptions)
				listDataProductsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listDataProductsOptionsModel.Start = core.StringPtr("testString")
				listDataProductsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dataProductExchangeApiServiceService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := dataProductExchangeApiServiceService.ListDataProducts(listDataProductsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListDataProducts successfully`, func() {
				dataProductExchangeApiServiceService, serviceErr := dataproductexchangeapiservicev1.NewDataProductExchangeApiServiceV1(&dataproductexchangeapiservicev1.DataProductExchangeApiServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductExchangeApiServiceService).ToNot(BeNil())

				// Construct an instance of the ListDataProductsOptions model
				listDataProductsOptionsModel := new(dataproductexchangeapiservicev1.ListDataProductsOptions)
				listDataProductsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listDataProductsOptionsModel.Start = core.StringPtr("testString")
				listDataProductsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := dataProductExchangeApiServiceService.ListDataProducts(listDataProductsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Test pagination helper method on response`, func() {
			It(`Invoke GetNextStart successfully`, func() {
				responseObject := new(dataproductexchangeapiservicev1.DataProductCollection)
				nextObject := new(dataproductexchangeapiservicev1.NextPage)
				nextObject.Start = core.StringPtr("abc-123")
				responseObject.Next = nextObject

				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(Equal(core.StringPtr("abc-123")))
			})
			It(`Invoke GetNextStart without a "Next" property in the response`, func() {
				responseObject := new(dataproductexchangeapiservicev1.DataProductCollection)

				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
		})
		Context(`Using mock server endpoint - paginated response`, func() {
			BeforeEach(func() {
				var requestNumber int = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listDataProductsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					requestNumber++
					if requestNumber == 1 {
						fmt.Fprintf(res, "%s", `{"next":{"start":"1"},"total_count":2,"limit":1,"data_products":[{"id":"b38df608-d34b-4d58-8136-ed25e6c6684e","container":{"id":"d29c42eb-7100-4b7a-8257-c196dbcca1cd","type":"catalog"},"name":"Sample Data Product"}]}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"total_count":2,"limit":1,"data_products":[{"id":"b38df608-d34b-4d58-8136-ed25e6c6684e","container":{"id":"d29c42eb-7100-4b7a-8257-c196dbcca1cd","type":"catalog"},"name":"Sample Data Product"}]}`)
					} else {
						res.WriteHeader(400)
					}
				}))
			})
			It(`Use DataProductsPager.GetNext successfully`, func() {
				dataProductExchangeApiServiceService, serviceErr := dataproductexchangeapiservicev1.NewDataProductExchangeApiServiceV1(&dataproductexchangeapiservicev1.DataProductExchangeApiServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductExchangeApiServiceService).ToNot(BeNil())

				listDataProductsOptionsModel := &dataproductexchangeapiservicev1.ListDataProductsOptions{
					Limit: core.Int64Ptr(int64(10)),
				}

				pager, err := dataProductExchangeApiServiceService.NewDataProductsPager(listDataProductsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				var allResults []dataproductexchangeapiservicev1.DataProduct
				for pager.HasNext() {
					nextPage, err := pager.GetNext()
					Expect(err).To(BeNil())
					Expect(nextPage).ToNot(BeNil())
					allResults = append(allResults, nextPage...)
				}
				Expect(len(allResults)).To(Equal(2))
			})
			It(`Use DataProductsPager.GetAll successfully`, func() {
				dataProductExchangeApiServiceService, serviceErr := dataproductexchangeapiservicev1.NewDataProductExchangeApiServiceV1(&dataproductexchangeapiservicev1.DataProductExchangeApiServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductExchangeApiServiceService).ToNot(BeNil())

				listDataProductsOptionsModel := &dataproductexchangeapiservicev1.ListDataProductsOptions{
					Limit: core.Int64Ptr(int64(10)),
				}

				pager, err := dataProductExchangeApiServiceService.NewDataProductsPager(listDataProductsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				allResults, err := pager.GetAll()
				Expect(err).To(BeNil())
				Expect(allResults).ToNot(BeNil())
				Expect(len(allResults)).To(Equal(2))
			})
		})
	})
	Describe(`ListDataProductVersions(listDataProductVersionsOptions *ListDataProductVersionsOptions) - Operation response error`, func() {
		listDataProductVersionsPath := "/data_product_exchange/v1/data_product_versions"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listDataProductVersionsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["asset.container.id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["data_product"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["state"]).To(Equal([]string{"draft"}))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListDataProductVersions with error: Operation response processing error`, func() {
				dataProductExchangeApiServiceService, serviceErr := dataproductexchangeapiservicev1.NewDataProductExchangeApiServiceV1(&dataproductexchangeapiservicev1.DataProductExchangeApiServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductExchangeApiServiceService).ToNot(BeNil())

				// Construct an instance of the ListDataProductVersionsOptions model
				listDataProductVersionsOptionsModel := new(dataproductexchangeapiservicev1.ListDataProductVersionsOptions)
				listDataProductVersionsOptionsModel.AssetContainerID = core.StringPtr("testString")
				listDataProductVersionsOptionsModel.DataProduct = core.StringPtr("testString")
				listDataProductVersionsOptionsModel.State = core.StringPtr("draft")
				listDataProductVersionsOptionsModel.Version = core.StringPtr("testString")
				listDataProductVersionsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listDataProductVersionsOptionsModel.Start = core.StringPtr("testString")
				listDataProductVersionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := dataProductExchangeApiServiceService.ListDataProductVersions(listDataProductVersionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				dataProductExchangeApiServiceService.EnableRetries(0, 0)
				result, response, operationErr = dataProductExchangeApiServiceService.ListDataProductVersions(listDataProductVersionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListDataProductVersions(listDataProductVersionsOptions *ListDataProductVersionsOptions)`, func() {
		listDataProductVersionsPath := "/data_product_exchange/v1/data_product_versions"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listDataProductVersionsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["asset.container.id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["data_product"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["state"]).To(Equal([]string{"draft"}))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"limit": 200, "first": {"href": "http://api.example.com/collection"}, "next": {"href": "http://api.example.com/collection?start=eyJvZmZzZXQiOjAsImRvbmUiOnRydWV9", "start": "eyJvZmZzZXQiOjAsImRvbmUiOnRydWV9"}, "data_product_versions": [{"version": "1.0.0", "state": "draft", "data_product": {"id": "b38df608-d34b-4d58-8136-ed25e6c6684e"}, "name": "My Data Product", "description": "This is a description of My Data Product.", "id": "2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd", "asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}}]}`)
				}))
			})
			It(`Invoke ListDataProductVersions successfully with retries`, func() {
				dataProductExchangeApiServiceService, serviceErr := dataproductexchangeapiservicev1.NewDataProductExchangeApiServiceV1(&dataproductexchangeapiservicev1.DataProductExchangeApiServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductExchangeApiServiceService).ToNot(BeNil())
				dataProductExchangeApiServiceService.EnableRetries(0, 0)

				// Construct an instance of the ListDataProductVersionsOptions model
				listDataProductVersionsOptionsModel := new(dataproductexchangeapiservicev1.ListDataProductVersionsOptions)
				listDataProductVersionsOptionsModel.AssetContainerID = core.StringPtr("testString")
				listDataProductVersionsOptionsModel.DataProduct = core.StringPtr("testString")
				listDataProductVersionsOptionsModel.State = core.StringPtr("draft")
				listDataProductVersionsOptionsModel.Version = core.StringPtr("testString")
				listDataProductVersionsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listDataProductVersionsOptionsModel.Start = core.StringPtr("testString")
				listDataProductVersionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := dataProductExchangeApiServiceService.ListDataProductVersionsWithContext(ctx, listDataProductVersionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				dataProductExchangeApiServiceService.DisableRetries()
				result, response, operationErr := dataProductExchangeApiServiceService.ListDataProductVersions(listDataProductVersionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = dataProductExchangeApiServiceService.ListDataProductVersionsWithContext(ctx, listDataProductVersionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listDataProductVersionsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["asset.container.id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["data_product"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["state"]).To(Equal([]string{"draft"}))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"limit": 200, "first": {"href": "http://api.example.com/collection"}, "next": {"href": "http://api.example.com/collection?start=eyJvZmZzZXQiOjAsImRvbmUiOnRydWV9", "start": "eyJvZmZzZXQiOjAsImRvbmUiOnRydWV9"}, "data_product_versions": [{"version": "1.0.0", "state": "draft", "data_product": {"id": "b38df608-d34b-4d58-8136-ed25e6c6684e"}, "name": "My Data Product", "description": "This is a description of My Data Product.", "id": "2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd", "asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}}]}`)
				}))
			})
			It(`Invoke ListDataProductVersions successfully`, func() {
				dataProductExchangeApiServiceService, serviceErr := dataproductexchangeapiservicev1.NewDataProductExchangeApiServiceV1(&dataproductexchangeapiservicev1.DataProductExchangeApiServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductExchangeApiServiceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := dataProductExchangeApiServiceService.ListDataProductVersions(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListDataProductVersionsOptions model
				listDataProductVersionsOptionsModel := new(dataproductexchangeapiservicev1.ListDataProductVersionsOptions)
				listDataProductVersionsOptionsModel.AssetContainerID = core.StringPtr("testString")
				listDataProductVersionsOptionsModel.DataProduct = core.StringPtr("testString")
				listDataProductVersionsOptionsModel.State = core.StringPtr("draft")
				listDataProductVersionsOptionsModel.Version = core.StringPtr("testString")
				listDataProductVersionsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listDataProductVersionsOptionsModel.Start = core.StringPtr("testString")
				listDataProductVersionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = dataProductExchangeApiServiceService.ListDataProductVersions(listDataProductVersionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListDataProductVersions with error: Operation request error`, func() {
				dataProductExchangeApiServiceService, serviceErr := dataproductexchangeapiservicev1.NewDataProductExchangeApiServiceV1(&dataproductexchangeapiservicev1.DataProductExchangeApiServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductExchangeApiServiceService).ToNot(BeNil())

				// Construct an instance of the ListDataProductVersionsOptions model
				listDataProductVersionsOptionsModel := new(dataproductexchangeapiservicev1.ListDataProductVersionsOptions)
				listDataProductVersionsOptionsModel.AssetContainerID = core.StringPtr("testString")
				listDataProductVersionsOptionsModel.DataProduct = core.StringPtr("testString")
				listDataProductVersionsOptionsModel.State = core.StringPtr("draft")
				listDataProductVersionsOptionsModel.Version = core.StringPtr("testString")
				listDataProductVersionsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listDataProductVersionsOptionsModel.Start = core.StringPtr("testString")
				listDataProductVersionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dataProductExchangeApiServiceService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := dataProductExchangeApiServiceService.ListDataProductVersions(listDataProductVersionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListDataProductVersions successfully`, func() {
				dataProductExchangeApiServiceService, serviceErr := dataproductexchangeapiservicev1.NewDataProductExchangeApiServiceV1(&dataproductexchangeapiservicev1.DataProductExchangeApiServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductExchangeApiServiceService).ToNot(BeNil())

				// Construct an instance of the ListDataProductVersionsOptions model
				listDataProductVersionsOptionsModel := new(dataproductexchangeapiservicev1.ListDataProductVersionsOptions)
				listDataProductVersionsOptionsModel.AssetContainerID = core.StringPtr("testString")
				listDataProductVersionsOptionsModel.DataProduct = core.StringPtr("testString")
				listDataProductVersionsOptionsModel.State = core.StringPtr("draft")
				listDataProductVersionsOptionsModel.Version = core.StringPtr("testString")
				listDataProductVersionsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listDataProductVersionsOptionsModel.Start = core.StringPtr("testString")
				listDataProductVersionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := dataProductExchangeApiServiceService.ListDataProductVersions(listDataProductVersionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Test pagination helper method on response`, func() {
			It(`Invoke GetNextStart successfully`, func() {
				responseObject := new(dataproductexchangeapiservicev1.DataProductVersionCollection)
				nextObject := new(dataproductexchangeapiservicev1.NextPage)
				nextObject.Start = core.StringPtr("abc-123")
				responseObject.Next = nextObject

				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(Equal(core.StringPtr("abc-123")))
			})
			It(`Invoke GetNextStart without a "Next" property in the response`, func() {
				responseObject := new(dataproductexchangeapiservicev1.DataProductVersionCollection)

				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
		})
		Context(`Using mock server endpoint - paginated response`, func() {
			BeforeEach(func() {
				var requestNumber int = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listDataProductVersionsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					requestNumber++
					if requestNumber == 1 {
						fmt.Fprintf(res, "%s", `{"next":{"start":"1"},"data_product_versions":[{"version":"1.0.0","state":"draft","data_product":{"id":"b38df608-d34b-4d58-8136-ed25e6c6684e"},"name":"My Data Product","description":"This is a description of My Data Product.","id":"2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd","asset":{"id":"2b0bf220-079c-11ee-be56-0242ac120002","container":{"id":"d29c42eb-7100-4b7a-8257-c196dbcca1cd","type":"catalog"}}}],"total_count":2,"limit":1}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"data_product_versions":[{"version":"1.0.0","state":"draft","data_product":{"id":"b38df608-d34b-4d58-8136-ed25e6c6684e"},"name":"My Data Product","description":"This is a description of My Data Product.","id":"2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd","asset":{"id":"2b0bf220-079c-11ee-be56-0242ac120002","container":{"id":"d29c42eb-7100-4b7a-8257-c196dbcca1cd","type":"catalog"}}}],"total_count":2,"limit":1}`)
					} else {
						res.WriteHeader(400)
					}
				}))
			})
			It(`Use DataProductVersionsPager.GetNext successfully`, func() {
				dataProductExchangeApiServiceService, serviceErr := dataproductexchangeapiservicev1.NewDataProductExchangeApiServiceV1(&dataproductexchangeapiservicev1.DataProductExchangeApiServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductExchangeApiServiceService).ToNot(BeNil())

				listDataProductVersionsOptionsModel := &dataproductexchangeapiservicev1.ListDataProductVersionsOptions{
					AssetContainerID: core.StringPtr("testString"),
					DataProduct: core.StringPtr("testString"),
					State: core.StringPtr("draft"),
					Version: core.StringPtr("testString"),
					Limit: core.Int64Ptr(int64(10)),
				}

				pager, err := dataProductExchangeApiServiceService.NewDataProductVersionsPager(listDataProductVersionsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				var allResults []dataproductexchangeapiservicev1.DataProductVersionSummary
				for pager.HasNext() {
					nextPage, err := pager.GetNext()
					Expect(err).To(BeNil())
					Expect(nextPage).ToNot(BeNil())
					allResults = append(allResults, nextPage...)
				}
				Expect(len(allResults)).To(Equal(2))
			})
			It(`Use DataProductVersionsPager.GetAll successfully`, func() {
				dataProductExchangeApiServiceService, serviceErr := dataproductexchangeapiservicev1.NewDataProductExchangeApiServiceV1(&dataproductexchangeapiservicev1.DataProductExchangeApiServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductExchangeApiServiceService).ToNot(BeNil())

				listDataProductVersionsOptionsModel := &dataproductexchangeapiservicev1.ListDataProductVersionsOptions{
					AssetContainerID: core.StringPtr("testString"),
					DataProduct: core.StringPtr("testString"),
					State: core.StringPtr("draft"),
					Version: core.StringPtr("testString"),
					Limit: core.Int64Ptr(int64(10)),
				}

				pager, err := dataProductExchangeApiServiceService.NewDataProductVersionsPager(listDataProductVersionsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				allResults, err := pager.GetAll()
				Expect(err).To(BeNil())
				Expect(allResults).ToNot(BeNil())
				Expect(len(allResults)).To(Equal(2))
			})
		})
	})
	Describe(`CreateDataProductVersion(createDataProductVersionOptions *CreateDataProductVersionOptions) - Operation response error`, func() {
		createDataProductVersionPath := "/data_product_exchange/v1/data_product_versions"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createDataProductVersionPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateDataProductVersion with error: Operation response processing error`, func() {
				dataProductExchangeApiServiceService, serviceErr := dataproductexchangeapiservicev1.NewDataProductExchangeApiServiceV1(&dataproductexchangeapiservicev1.DataProductExchangeApiServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductExchangeApiServiceService).ToNot(BeNil())

				// Construct an instance of the ContainerReference model
				containerReferenceModel := new(dataproductexchangeapiservicev1.ContainerReference)
				containerReferenceModel.ID = core.StringPtr("d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				containerReferenceModel.Type = core.StringPtr("catalog")

				// Construct an instance of the DataProductIdentity model
				dataProductIdentityModel := new(dataproductexchangeapiservicev1.DataProductIdentity)
				dataProductIdentityModel.ID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")

				// Construct an instance of the UseCase model
				useCaseModel := new(dataproductexchangeapiservicev1.UseCase)
				useCaseModel.ID = core.StringPtr("testString")
				useCaseModel.Name = core.StringPtr("testString")
				useCaseModel.Container = containerReferenceModel

				// Construct an instance of the Domain model
				domainModel := new(dataproductexchangeapiservicev1.Domain)
				domainModel.ID = core.StringPtr("testString")
				domainModel.Name = core.StringPtr("testString")
				domainModel.Container = containerReferenceModel

				// Construct an instance of the AssetPartReference model
				assetPartReferenceModel := new(dataproductexchangeapiservicev1.AssetPartReference)
				assetPartReferenceModel.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				assetPartReferenceModel.Container = containerReferenceModel
				assetPartReferenceModel.Type = core.StringPtr("data_asset")

				// Construct an instance of the DeliveryMethod model
				deliveryMethodModel := new(dataproductexchangeapiservicev1.DeliveryMethod)
				deliveryMethodModel.ID = core.StringPtr("09cf5fcc-cb9d-4995-a8e4-16517b25229f")
				deliveryMethodModel.Container = containerReferenceModel

				// Construct an instance of the DataProductPart model
				dataProductPartModel := new(dataproductexchangeapiservicev1.DataProductPart)
				dataProductPartModel.Asset = assetPartReferenceModel
				dataProductPartModel.Revision = core.Int64Ptr(int64(1))
				dataProductPartModel.UpdatedAt = CreateMockDateTime("2023-07-01T22:22:34.876Z")
				dataProductPartModel.DeliveryMethods = []dataproductexchangeapiservicev1.DeliveryMethod{*deliveryMethodModel}

				// Construct an instance of the CreateDataProductVersionOptions model
				createDataProductVersionOptionsModel := new(dataproductexchangeapiservicev1.CreateDataProductVersionOptions)
				createDataProductVersionOptionsModel.Container = containerReferenceModel
				createDataProductVersionOptionsModel.Version = core.StringPtr("testString")
				createDataProductVersionOptionsModel.State = core.StringPtr("draft")
				createDataProductVersionOptionsModel.DataProduct = dataProductIdentityModel
				createDataProductVersionOptionsModel.Name = core.StringPtr("My New Data Product")
				createDataProductVersionOptionsModel.Description = core.StringPtr("testString")
				createDataProductVersionOptionsModel.Tags = []string{"testString"}
				createDataProductVersionOptionsModel.UseCases = []dataproductexchangeapiservicev1.UseCase{*useCaseModel}
				createDataProductVersionOptionsModel.Domain = domainModel
				createDataProductVersionOptionsModel.Type = []string{"data"}
				createDataProductVersionOptionsModel.PartsOut = []dataproductexchangeapiservicev1.DataProductPart{*dataProductPartModel}
				createDataProductVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := dataProductExchangeApiServiceService.CreateDataProductVersion(createDataProductVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				dataProductExchangeApiServiceService.EnableRetries(0, 0)
				result, response, operationErr = dataProductExchangeApiServiceService.CreateDataProductVersion(createDataProductVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateDataProductVersion(createDataProductVersionOptions *CreateDataProductVersionOptions)`, func() {
		createDataProductVersionPath := "/data_product_exchange/v1/data_product_versions"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createDataProductVersionPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"version": "1.0.0", "state": "draft", "data_product": {"id": "b38df608-d34b-4d58-8136-ed25e6c6684e"}, "name": "My Data Product", "description": "This is a description of My Data Product.", "id": "2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd", "asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "tags": ["Tags"], "use_cases": [{"id": "ID", "name": "Name", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}], "domain": {"id": "ID", "name": "Name", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "type": ["data"], "parts_out": [{"asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}, "type": "data_asset"}, "revision": 1, "updated_at": "2023-07-01T22:22:34.876Z", "delivery_methods": [{"id": "09cf5fcc-cb9d-4995-a8e4-16517b25229f", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}]}], "published_by": "PublishedBy", "published_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "created_at": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke CreateDataProductVersion successfully with retries`, func() {
				dataProductExchangeApiServiceService, serviceErr := dataproductexchangeapiservicev1.NewDataProductExchangeApiServiceV1(&dataproductexchangeapiservicev1.DataProductExchangeApiServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductExchangeApiServiceService).ToNot(BeNil())
				dataProductExchangeApiServiceService.EnableRetries(0, 0)

				// Construct an instance of the ContainerReference model
				containerReferenceModel := new(dataproductexchangeapiservicev1.ContainerReference)
				containerReferenceModel.ID = core.StringPtr("d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				containerReferenceModel.Type = core.StringPtr("catalog")

				// Construct an instance of the DataProductIdentity model
				dataProductIdentityModel := new(dataproductexchangeapiservicev1.DataProductIdentity)
				dataProductIdentityModel.ID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")

				// Construct an instance of the UseCase model
				useCaseModel := new(dataproductexchangeapiservicev1.UseCase)
				useCaseModel.ID = core.StringPtr("testString")
				useCaseModel.Name = core.StringPtr("testString")
				useCaseModel.Container = containerReferenceModel

				// Construct an instance of the Domain model
				domainModel := new(dataproductexchangeapiservicev1.Domain)
				domainModel.ID = core.StringPtr("testString")
				domainModel.Name = core.StringPtr("testString")
				domainModel.Container = containerReferenceModel

				// Construct an instance of the AssetPartReference model
				assetPartReferenceModel := new(dataproductexchangeapiservicev1.AssetPartReference)
				assetPartReferenceModel.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				assetPartReferenceModel.Container = containerReferenceModel
				assetPartReferenceModel.Type = core.StringPtr("data_asset")

				// Construct an instance of the DeliveryMethod model
				deliveryMethodModel := new(dataproductexchangeapiservicev1.DeliveryMethod)
				deliveryMethodModel.ID = core.StringPtr("09cf5fcc-cb9d-4995-a8e4-16517b25229f")
				deliveryMethodModel.Container = containerReferenceModel

				// Construct an instance of the DataProductPart model
				dataProductPartModel := new(dataproductexchangeapiservicev1.DataProductPart)
				dataProductPartModel.Asset = assetPartReferenceModel
				dataProductPartModel.Revision = core.Int64Ptr(int64(1))
				dataProductPartModel.UpdatedAt = CreateMockDateTime("2023-07-01T22:22:34.876Z")
				dataProductPartModel.DeliveryMethods = []dataproductexchangeapiservicev1.DeliveryMethod{*deliveryMethodModel}

				// Construct an instance of the CreateDataProductVersionOptions model
				createDataProductVersionOptionsModel := new(dataproductexchangeapiservicev1.CreateDataProductVersionOptions)
				createDataProductVersionOptionsModel.Container = containerReferenceModel
				createDataProductVersionOptionsModel.Version = core.StringPtr("testString")
				createDataProductVersionOptionsModel.State = core.StringPtr("draft")
				createDataProductVersionOptionsModel.DataProduct = dataProductIdentityModel
				createDataProductVersionOptionsModel.Name = core.StringPtr("My New Data Product")
				createDataProductVersionOptionsModel.Description = core.StringPtr("testString")
				createDataProductVersionOptionsModel.Tags = []string{"testString"}
				createDataProductVersionOptionsModel.UseCases = []dataproductexchangeapiservicev1.UseCase{*useCaseModel}
				createDataProductVersionOptionsModel.Domain = domainModel
				createDataProductVersionOptionsModel.Type = []string{"data"}
				createDataProductVersionOptionsModel.PartsOut = []dataproductexchangeapiservicev1.DataProductPart{*dataProductPartModel}
				createDataProductVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := dataProductExchangeApiServiceService.CreateDataProductVersionWithContext(ctx, createDataProductVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				dataProductExchangeApiServiceService.DisableRetries()
				result, response, operationErr := dataProductExchangeApiServiceService.CreateDataProductVersion(createDataProductVersionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = dataProductExchangeApiServiceService.CreateDataProductVersionWithContext(ctx, createDataProductVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createDataProductVersionPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"version": "1.0.0", "state": "draft", "data_product": {"id": "b38df608-d34b-4d58-8136-ed25e6c6684e"}, "name": "My Data Product", "description": "This is a description of My Data Product.", "id": "2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd", "asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "tags": ["Tags"], "use_cases": [{"id": "ID", "name": "Name", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}], "domain": {"id": "ID", "name": "Name", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "type": ["data"], "parts_out": [{"asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}, "type": "data_asset"}, "revision": 1, "updated_at": "2023-07-01T22:22:34.876Z", "delivery_methods": [{"id": "09cf5fcc-cb9d-4995-a8e4-16517b25229f", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}]}], "published_by": "PublishedBy", "published_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "created_at": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke CreateDataProductVersion successfully`, func() {
				dataProductExchangeApiServiceService, serviceErr := dataproductexchangeapiservicev1.NewDataProductExchangeApiServiceV1(&dataproductexchangeapiservicev1.DataProductExchangeApiServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductExchangeApiServiceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := dataProductExchangeApiServiceService.CreateDataProductVersion(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ContainerReference model
				containerReferenceModel := new(dataproductexchangeapiservicev1.ContainerReference)
				containerReferenceModel.ID = core.StringPtr("d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				containerReferenceModel.Type = core.StringPtr("catalog")

				// Construct an instance of the DataProductIdentity model
				dataProductIdentityModel := new(dataproductexchangeapiservicev1.DataProductIdentity)
				dataProductIdentityModel.ID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")

				// Construct an instance of the UseCase model
				useCaseModel := new(dataproductexchangeapiservicev1.UseCase)
				useCaseModel.ID = core.StringPtr("testString")
				useCaseModel.Name = core.StringPtr("testString")
				useCaseModel.Container = containerReferenceModel

				// Construct an instance of the Domain model
				domainModel := new(dataproductexchangeapiservicev1.Domain)
				domainModel.ID = core.StringPtr("testString")
				domainModel.Name = core.StringPtr("testString")
				domainModel.Container = containerReferenceModel

				// Construct an instance of the AssetPartReference model
				assetPartReferenceModel := new(dataproductexchangeapiservicev1.AssetPartReference)
				assetPartReferenceModel.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				assetPartReferenceModel.Container = containerReferenceModel
				assetPartReferenceModel.Type = core.StringPtr("data_asset")

				// Construct an instance of the DeliveryMethod model
				deliveryMethodModel := new(dataproductexchangeapiservicev1.DeliveryMethod)
				deliveryMethodModel.ID = core.StringPtr("09cf5fcc-cb9d-4995-a8e4-16517b25229f")
				deliveryMethodModel.Container = containerReferenceModel

				// Construct an instance of the DataProductPart model
				dataProductPartModel := new(dataproductexchangeapiservicev1.DataProductPart)
				dataProductPartModel.Asset = assetPartReferenceModel
				dataProductPartModel.Revision = core.Int64Ptr(int64(1))
				dataProductPartModel.UpdatedAt = CreateMockDateTime("2023-07-01T22:22:34.876Z")
				dataProductPartModel.DeliveryMethods = []dataproductexchangeapiservicev1.DeliveryMethod{*deliveryMethodModel}

				// Construct an instance of the CreateDataProductVersionOptions model
				createDataProductVersionOptionsModel := new(dataproductexchangeapiservicev1.CreateDataProductVersionOptions)
				createDataProductVersionOptionsModel.Container = containerReferenceModel
				createDataProductVersionOptionsModel.Version = core.StringPtr("testString")
				createDataProductVersionOptionsModel.State = core.StringPtr("draft")
				createDataProductVersionOptionsModel.DataProduct = dataProductIdentityModel
				createDataProductVersionOptionsModel.Name = core.StringPtr("My New Data Product")
				createDataProductVersionOptionsModel.Description = core.StringPtr("testString")
				createDataProductVersionOptionsModel.Tags = []string{"testString"}
				createDataProductVersionOptionsModel.UseCases = []dataproductexchangeapiservicev1.UseCase{*useCaseModel}
				createDataProductVersionOptionsModel.Domain = domainModel
				createDataProductVersionOptionsModel.Type = []string{"data"}
				createDataProductVersionOptionsModel.PartsOut = []dataproductexchangeapiservicev1.DataProductPart{*dataProductPartModel}
				createDataProductVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = dataProductExchangeApiServiceService.CreateDataProductVersion(createDataProductVersionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateDataProductVersion with error: Operation validation and request error`, func() {
				dataProductExchangeApiServiceService, serviceErr := dataproductexchangeapiservicev1.NewDataProductExchangeApiServiceV1(&dataproductexchangeapiservicev1.DataProductExchangeApiServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductExchangeApiServiceService).ToNot(BeNil())

				// Construct an instance of the ContainerReference model
				containerReferenceModel := new(dataproductexchangeapiservicev1.ContainerReference)
				containerReferenceModel.ID = core.StringPtr("d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				containerReferenceModel.Type = core.StringPtr("catalog")

				// Construct an instance of the DataProductIdentity model
				dataProductIdentityModel := new(dataproductexchangeapiservicev1.DataProductIdentity)
				dataProductIdentityModel.ID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")

				// Construct an instance of the UseCase model
				useCaseModel := new(dataproductexchangeapiservicev1.UseCase)
				useCaseModel.ID = core.StringPtr("testString")
				useCaseModel.Name = core.StringPtr("testString")
				useCaseModel.Container = containerReferenceModel

				// Construct an instance of the Domain model
				domainModel := new(dataproductexchangeapiservicev1.Domain)
				domainModel.ID = core.StringPtr("testString")
				domainModel.Name = core.StringPtr("testString")
				domainModel.Container = containerReferenceModel

				// Construct an instance of the AssetPartReference model
				assetPartReferenceModel := new(dataproductexchangeapiservicev1.AssetPartReference)
				assetPartReferenceModel.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				assetPartReferenceModel.Container = containerReferenceModel
				assetPartReferenceModel.Type = core.StringPtr("data_asset")

				// Construct an instance of the DeliveryMethod model
				deliveryMethodModel := new(dataproductexchangeapiservicev1.DeliveryMethod)
				deliveryMethodModel.ID = core.StringPtr("09cf5fcc-cb9d-4995-a8e4-16517b25229f")
				deliveryMethodModel.Container = containerReferenceModel

				// Construct an instance of the DataProductPart model
				dataProductPartModel := new(dataproductexchangeapiservicev1.DataProductPart)
				dataProductPartModel.Asset = assetPartReferenceModel
				dataProductPartModel.Revision = core.Int64Ptr(int64(1))
				dataProductPartModel.UpdatedAt = CreateMockDateTime("2023-07-01T22:22:34.876Z")
				dataProductPartModel.DeliveryMethods = []dataproductexchangeapiservicev1.DeliveryMethod{*deliveryMethodModel}

				// Construct an instance of the CreateDataProductVersionOptions model
				createDataProductVersionOptionsModel := new(dataproductexchangeapiservicev1.CreateDataProductVersionOptions)
				createDataProductVersionOptionsModel.Container = containerReferenceModel
				createDataProductVersionOptionsModel.Version = core.StringPtr("testString")
				createDataProductVersionOptionsModel.State = core.StringPtr("draft")
				createDataProductVersionOptionsModel.DataProduct = dataProductIdentityModel
				createDataProductVersionOptionsModel.Name = core.StringPtr("My New Data Product")
				createDataProductVersionOptionsModel.Description = core.StringPtr("testString")
				createDataProductVersionOptionsModel.Tags = []string{"testString"}
				createDataProductVersionOptionsModel.UseCases = []dataproductexchangeapiservicev1.UseCase{*useCaseModel}
				createDataProductVersionOptionsModel.Domain = domainModel
				createDataProductVersionOptionsModel.Type = []string{"data"}
				createDataProductVersionOptionsModel.PartsOut = []dataproductexchangeapiservicev1.DataProductPart{*dataProductPartModel}
				createDataProductVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dataProductExchangeApiServiceService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := dataProductExchangeApiServiceService.CreateDataProductVersion(createDataProductVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateDataProductVersionOptions model with no property values
				createDataProductVersionOptionsModelNew := new(dataproductexchangeapiservicev1.CreateDataProductVersionOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = dataProductExchangeApiServiceService.CreateDataProductVersion(createDataProductVersionOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(201)
				}))
			})
			It(`Invoke CreateDataProductVersion successfully`, func() {
				dataProductExchangeApiServiceService, serviceErr := dataproductexchangeapiservicev1.NewDataProductExchangeApiServiceV1(&dataproductexchangeapiservicev1.DataProductExchangeApiServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductExchangeApiServiceService).ToNot(BeNil())

				// Construct an instance of the ContainerReference model
				containerReferenceModel := new(dataproductexchangeapiservicev1.ContainerReference)
				containerReferenceModel.ID = core.StringPtr("d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				containerReferenceModel.Type = core.StringPtr("catalog")

				// Construct an instance of the DataProductIdentity model
				dataProductIdentityModel := new(dataproductexchangeapiservicev1.DataProductIdentity)
				dataProductIdentityModel.ID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")

				// Construct an instance of the UseCase model
				useCaseModel := new(dataproductexchangeapiservicev1.UseCase)
				useCaseModel.ID = core.StringPtr("testString")
				useCaseModel.Name = core.StringPtr("testString")
				useCaseModel.Container = containerReferenceModel

				// Construct an instance of the Domain model
				domainModel := new(dataproductexchangeapiservicev1.Domain)
				domainModel.ID = core.StringPtr("testString")
				domainModel.Name = core.StringPtr("testString")
				domainModel.Container = containerReferenceModel

				// Construct an instance of the AssetPartReference model
				assetPartReferenceModel := new(dataproductexchangeapiservicev1.AssetPartReference)
				assetPartReferenceModel.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				assetPartReferenceModel.Container = containerReferenceModel
				assetPartReferenceModel.Type = core.StringPtr("data_asset")

				// Construct an instance of the DeliveryMethod model
				deliveryMethodModel := new(dataproductexchangeapiservicev1.DeliveryMethod)
				deliveryMethodModel.ID = core.StringPtr("09cf5fcc-cb9d-4995-a8e4-16517b25229f")
				deliveryMethodModel.Container = containerReferenceModel

				// Construct an instance of the DataProductPart model
				dataProductPartModel := new(dataproductexchangeapiservicev1.DataProductPart)
				dataProductPartModel.Asset = assetPartReferenceModel
				dataProductPartModel.Revision = core.Int64Ptr(int64(1))
				dataProductPartModel.UpdatedAt = CreateMockDateTime("2023-07-01T22:22:34.876Z")
				dataProductPartModel.DeliveryMethods = []dataproductexchangeapiservicev1.DeliveryMethod{*deliveryMethodModel}

				// Construct an instance of the CreateDataProductVersionOptions model
				createDataProductVersionOptionsModel := new(dataproductexchangeapiservicev1.CreateDataProductVersionOptions)
				createDataProductVersionOptionsModel.Container = containerReferenceModel
				createDataProductVersionOptionsModel.Version = core.StringPtr("testString")
				createDataProductVersionOptionsModel.State = core.StringPtr("draft")
				createDataProductVersionOptionsModel.DataProduct = dataProductIdentityModel
				createDataProductVersionOptionsModel.Name = core.StringPtr("My New Data Product")
				createDataProductVersionOptionsModel.Description = core.StringPtr("testString")
				createDataProductVersionOptionsModel.Tags = []string{"testString"}
				createDataProductVersionOptionsModel.UseCases = []dataproductexchangeapiservicev1.UseCase{*useCaseModel}
				createDataProductVersionOptionsModel.Domain = domainModel
				createDataProductVersionOptionsModel.Type = []string{"data"}
				createDataProductVersionOptionsModel.PartsOut = []dataproductexchangeapiservicev1.DataProductPart{*dataProductPartModel}
				createDataProductVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := dataProductExchangeApiServiceService.CreateDataProductVersion(createDataProductVersionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetDataProductVersion(getDataProductVersionOptions *GetDataProductVersionOptions) - Operation response error`, func() {
		getDataProductVersionPath := "/data_product_exchange/v1/data_product_versions/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDataProductVersionPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetDataProductVersion with error: Operation response processing error`, func() {
				dataProductExchangeApiServiceService, serviceErr := dataproductexchangeapiservicev1.NewDataProductExchangeApiServiceV1(&dataproductexchangeapiservicev1.DataProductExchangeApiServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductExchangeApiServiceService).ToNot(BeNil())

				// Construct an instance of the GetDataProductVersionOptions model
				getDataProductVersionOptionsModel := new(dataproductexchangeapiservicev1.GetDataProductVersionOptions)
				getDataProductVersionOptionsModel.ID = core.StringPtr("testString")
				getDataProductVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := dataProductExchangeApiServiceService.GetDataProductVersion(getDataProductVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				dataProductExchangeApiServiceService.EnableRetries(0, 0)
				result, response, operationErr = dataProductExchangeApiServiceService.GetDataProductVersion(getDataProductVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetDataProductVersion(getDataProductVersionOptions *GetDataProductVersionOptions)`, func() {
		getDataProductVersionPath := "/data_product_exchange/v1/data_product_versions/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDataProductVersionPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"version": "1.0.0", "state": "draft", "data_product": {"id": "b38df608-d34b-4d58-8136-ed25e6c6684e"}, "name": "My Data Product", "description": "This is a description of My Data Product.", "id": "2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd", "asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "tags": ["Tags"], "use_cases": [{"id": "ID", "name": "Name", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}], "domain": {"id": "ID", "name": "Name", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "type": ["data"], "parts_out": [{"asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}, "type": "data_asset"}, "revision": 1, "updated_at": "2023-07-01T22:22:34.876Z", "delivery_methods": [{"id": "09cf5fcc-cb9d-4995-a8e4-16517b25229f", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}]}], "published_by": "PublishedBy", "published_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "created_at": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke GetDataProductVersion successfully with retries`, func() {
				dataProductExchangeApiServiceService, serviceErr := dataproductexchangeapiservicev1.NewDataProductExchangeApiServiceV1(&dataproductexchangeapiservicev1.DataProductExchangeApiServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductExchangeApiServiceService).ToNot(BeNil())
				dataProductExchangeApiServiceService.EnableRetries(0, 0)

				// Construct an instance of the GetDataProductVersionOptions model
				getDataProductVersionOptionsModel := new(dataproductexchangeapiservicev1.GetDataProductVersionOptions)
				getDataProductVersionOptionsModel.ID = core.StringPtr("testString")
				getDataProductVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := dataProductExchangeApiServiceService.GetDataProductVersionWithContext(ctx, getDataProductVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				dataProductExchangeApiServiceService.DisableRetries()
				result, response, operationErr := dataProductExchangeApiServiceService.GetDataProductVersion(getDataProductVersionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = dataProductExchangeApiServiceService.GetDataProductVersionWithContext(ctx, getDataProductVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDataProductVersionPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"version": "1.0.0", "state": "draft", "data_product": {"id": "b38df608-d34b-4d58-8136-ed25e6c6684e"}, "name": "My Data Product", "description": "This is a description of My Data Product.", "id": "2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd", "asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "tags": ["Tags"], "use_cases": [{"id": "ID", "name": "Name", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}], "domain": {"id": "ID", "name": "Name", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "type": ["data"], "parts_out": [{"asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}, "type": "data_asset"}, "revision": 1, "updated_at": "2023-07-01T22:22:34.876Z", "delivery_methods": [{"id": "09cf5fcc-cb9d-4995-a8e4-16517b25229f", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}]}], "published_by": "PublishedBy", "published_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "created_at": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke GetDataProductVersion successfully`, func() {
				dataProductExchangeApiServiceService, serviceErr := dataproductexchangeapiservicev1.NewDataProductExchangeApiServiceV1(&dataproductexchangeapiservicev1.DataProductExchangeApiServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductExchangeApiServiceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := dataProductExchangeApiServiceService.GetDataProductVersion(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetDataProductVersionOptions model
				getDataProductVersionOptionsModel := new(dataproductexchangeapiservicev1.GetDataProductVersionOptions)
				getDataProductVersionOptionsModel.ID = core.StringPtr("testString")
				getDataProductVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = dataProductExchangeApiServiceService.GetDataProductVersion(getDataProductVersionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetDataProductVersion with error: Operation validation and request error`, func() {
				dataProductExchangeApiServiceService, serviceErr := dataproductexchangeapiservicev1.NewDataProductExchangeApiServiceV1(&dataproductexchangeapiservicev1.DataProductExchangeApiServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductExchangeApiServiceService).ToNot(BeNil())

				// Construct an instance of the GetDataProductVersionOptions model
				getDataProductVersionOptionsModel := new(dataproductexchangeapiservicev1.GetDataProductVersionOptions)
				getDataProductVersionOptionsModel.ID = core.StringPtr("testString")
				getDataProductVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dataProductExchangeApiServiceService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := dataProductExchangeApiServiceService.GetDataProductVersion(getDataProductVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetDataProductVersionOptions model with no property values
				getDataProductVersionOptionsModelNew := new(dataproductexchangeapiservicev1.GetDataProductVersionOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = dataProductExchangeApiServiceService.GetDataProductVersion(getDataProductVersionOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetDataProductVersion successfully`, func() {
				dataProductExchangeApiServiceService, serviceErr := dataproductexchangeapiservicev1.NewDataProductExchangeApiServiceV1(&dataproductexchangeapiservicev1.DataProductExchangeApiServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductExchangeApiServiceService).ToNot(BeNil())

				// Construct an instance of the GetDataProductVersionOptions model
				getDataProductVersionOptionsModel := new(dataproductexchangeapiservicev1.GetDataProductVersionOptions)
				getDataProductVersionOptionsModel.ID = core.StringPtr("testString")
				getDataProductVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := dataProductExchangeApiServiceService.GetDataProductVersion(getDataProductVersionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteDataProductVersion(deleteDataProductVersionOptions *DeleteDataProductVersionOptions)`, func() {
		deleteDataProductVersionPath := "/data_product_exchange/v1/data_product_versions/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteDataProductVersionPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteDataProductVersion successfully`, func() {
				dataProductExchangeApiServiceService, serviceErr := dataproductexchangeapiservicev1.NewDataProductExchangeApiServiceV1(&dataproductexchangeapiservicev1.DataProductExchangeApiServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductExchangeApiServiceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := dataProductExchangeApiServiceService.DeleteDataProductVersion(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteDataProductVersionOptions model
				deleteDataProductVersionOptionsModel := new(dataproductexchangeapiservicev1.DeleteDataProductVersionOptions)
				deleteDataProductVersionOptionsModel.ID = core.StringPtr("testString")
				deleteDataProductVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = dataProductExchangeApiServiceService.DeleteDataProductVersion(deleteDataProductVersionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteDataProductVersion with error: Operation validation and request error`, func() {
				dataProductExchangeApiServiceService, serviceErr := dataproductexchangeapiservicev1.NewDataProductExchangeApiServiceV1(&dataproductexchangeapiservicev1.DataProductExchangeApiServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductExchangeApiServiceService).ToNot(BeNil())

				// Construct an instance of the DeleteDataProductVersionOptions model
				deleteDataProductVersionOptionsModel := new(dataproductexchangeapiservicev1.DeleteDataProductVersionOptions)
				deleteDataProductVersionOptionsModel.ID = core.StringPtr("testString")
				deleteDataProductVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dataProductExchangeApiServiceService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := dataProductExchangeApiServiceService.DeleteDataProductVersion(deleteDataProductVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteDataProductVersionOptions model with no property values
				deleteDataProductVersionOptionsModelNew := new(dataproductexchangeapiservicev1.DeleteDataProductVersionOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = dataProductExchangeApiServiceService.DeleteDataProductVersion(deleteDataProductVersionOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateDataProductVersion(updateDataProductVersionOptions *UpdateDataProductVersionOptions) - Operation response error`, func() {
		updateDataProductVersionPath := "/data_product_exchange/v1/data_product_versions/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateDataProductVersionPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateDataProductVersion with error: Operation response processing error`, func() {
				dataProductExchangeApiServiceService, serviceErr := dataproductexchangeapiservicev1.NewDataProductExchangeApiServiceV1(&dataproductexchangeapiservicev1.DataProductExchangeApiServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductExchangeApiServiceService).ToNot(BeNil())

				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(dataproductexchangeapiservicev1.JSONPatchOperation)
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = core.StringPtr("testString")

				// Construct an instance of the UpdateDataProductVersionOptions model
				updateDataProductVersionOptionsModel := new(dataproductexchangeapiservicev1.UpdateDataProductVersionOptions)
				updateDataProductVersionOptionsModel.ID = core.StringPtr("testString")
				updateDataProductVersionOptionsModel.JSONPatchInstructions = []dataproductexchangeapiservicev1.JSONPatchOperation{*jsonPatchOperationModel}
				updateDataProductVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := dataProductExchangeApiServiceService.UpdateDataProductVersion(updateDataProductVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				dataProductExchangeApiServiceService.EnableRetries(0, 0)
				result, response, operationErr = dataProductExchangeApiServiceService.UpdateDataProductVersion(updateDataProductVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateDataProductVersion(updateDataProductVersionOptions *UpdateDataProductVersionOptions)`, func() {
		updateDataProductVersionPath := "/data_product_exchange/v1/data_product_versions/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateDataProductVersionPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"version": "1.0.0", "state": "draft", "data_product": {"id": "b38df608-d34b-4d58-8136-ed25e6c6684e"}, "name": "My Data Product", "description": "This is a description of My Data Product.", "id": "2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd", "asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "tags": ["Tags"], "use_cases": [{"id": "ID", "name": "Name", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}], "domain": {"id": "ID", "name": "Name", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "type": ["data"], "parts_out": [{"asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}, "type": "data_asset"}, "revision": 1, "updated_at": "2023-07-01T22:22:34.876Z", "delivery_methods": [{"id": "09cf5fcc-cb9d-4995-a8e4-16517b25229f", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}]}], "published_by": "PublishedBy", "published_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "created_at": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke UpdateDataProductVersion successfully with retries`, func() {
				dataProductExchangeApiServiceService, serviceErr := dataproductexchangeapiservicev1.NewDataProductExchangeApiServiceV1(&dataproductexchangeapiservicev1.DataProductExchangeApiServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductExchangeApiServiceService).ToNot(BeNil())
				dataProductExchangeApiServiceService.EnableRetries(0, 0)

				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(dataproductexchangeapiservicev1.JSONPatchOperation)
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = core.StringPtr("testString")

				// Construct an instance of the UpdateDataProductVersionOptions model
				updateDataProductVersionOptionsModel := new(dataproductexchangeapiservicev1.UpdateDataProductVersionOptions)
				updateDataProductVersionOptionsModel.ID = core.StringPtr("testString")
				updateDataProductVersionOptionsModel.JSONPatchInstructions = []dataproductexchangeapiservicev1.JSONPatchOperation{*jsonPatchOperationModel}
				updateDataProductVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := dataProductExchangeApiServiceService.UpdateDataProductVersionWithContext(ctx, updateDataProductVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				dataProductExchangeApiServiceService.DisableRetries()
				result, response, operationErr := dataProductExchangeApiServiceService.UpdateDataProductVersion(updateDataProductVersionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = dataProductExchangeApiServiceService.UpdateDataProductVersionWithContext(ctx, updateDataProductVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateDataProductVersionPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"version": "1.0.0", "state": "draft", "data_product": {"id": "b38df608-d34b-4d58-8136-ed25e6c6684e"}, "name": "My Data Product", "description": "This is a description of My Data Product.", "id": "2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd", "asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "tags": ["Tags"], "use_cases": [{"id": "ID", "name": "Name", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}], "domain": {"id": "ID", "name": "Name", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "type": ["data"], "parts_out": [{"asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}, "type": "data_asset"}, "revision": 1, "updated_at": "2023-07-01T22:22:34.876Z", "delivery_methods": [{"id": "09cf5fcc-cb9d-4995-a8e4-16517b25229f", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}]}], "published_by": "PublishedBy", "published_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "created_at": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke UpdateDataProductVersion successfully`, func() {
				dataProductExchangeApiServiceService, serviceErr := dataproductexchangeapiservicev1.NewDataProductExchangeApiServiceV1(&dataproductexchangeapiservicev1.DataProductExchangeApiServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductExchangeApiServiceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := dataProductExchangeApiServiceService.UpdateDataProductVersion(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(dataproductexchangeapiservicev1.JSONPatchOperation)
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = core.StringPtr("testString")

				// Construct an instance of the UpdateDataProductVersionOptions model
				updateDataProductVersionOptionsModel := new(dataproductexchangeapiservicev1.UpdateDataProductVersionOptions)
				updateDataProductVersionOptionsModel.ID = core.StringPtr("testString")
				updateDataProductVersionOptionsModel.JSONPatchInstructions = []dataproductexchangeapiservicev1.JSONPatchOperation{*jsonPatchOperationModel}
				updateDataProductVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = dataProductExchangeApiServiceService.UpdateDataProductVersion(updateDataProductVersionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateDataProductVersion with error: Operation validation and request error`, func() {
				dataProductExchangeApiServiceService, serviceErr := dataproductexchangeapiservicev1.NewDataProductExchangeApiServiceV1(&dataproductexchangeapiservicev1.DataProductExchangeApiServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductExchangeApiServiceService).ToNot(BeNil())

				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(dataproductexchangeapiservicev1.JSONPatchOperation)
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = core.StringPtr("testString")

				// Construct an instance of the UpdateDataProductVersionOptions model
				updateDataProductVersionOptionsModel := new(dataproductexchangeapiservicev1.UpdateDataProductVersionOptions)
				updateDataProductVersionOptionsModel.ID = core.StringPtr("testString")
				updateDataProductVersionOptionsModel.JSONPatchInstructions = []dataproductexchangeapiservicev1.JSONPatchOperation{*jsonPatchOperationModel}
				updateDataProductVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dataProductExchangeApiServiceService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := dataProductExchangeApiServiceService.UpdateDataProductVersion(updateDataProductVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateDataProductVersionOptions model with no property values
				updateDataProductVersionOptionsModelNew := new(dataproductexchangeapiservicev1.UpdateDataProductVersionOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = dataProductExchangeApiServiceService.UpdateDataProductVersion(updateDataProductVersionOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke UpdateDataProductVersion successfully`, func() {
				dataProductExchangeApiServiceService, serviceErr := dataproductexchangeapiservicev1.NewDataProductExchangeApiServiceV1(&dataproductexchangeapiservicev1.DataProductExchangeApiServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductExchangeApiServiceService).ToNot(BeNil())

				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(dataproductexchangeapiservicev1.JSONPatchOperation)
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = core.StringPtr("testString")

				// Construct an instance of the UpdateDataProductVersionOptions model
				updateDataProductVersionOptionsModel := new(dataproductexchangeapiservicev1.UpdateDataProductVersionOptions)
				updateDataProductVersionOptionsModel.ID = core.StringPtr("testString")
				updateDataProductVersionOptionsModel.JSONPatchInstructions = []dataproductexchangeapiservicev1.JSONPatchOperation{*jsonPatchOperationModel}
				updateDataProductVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := dataProductExchangeApiServiceService.UpdateDataProductVersion(updateDataProductVersionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeliverDataProductVersion(deliverDataProductVersionOptions *DeliverDataProductVersionOptions) - Operation response error`, func() {
		deliverDataProductVersionPath := "/data_product_exchange/v1/data_product_versions/testString/deliver"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deliverDataProductVersionPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeliverDataProductVersion with error: Operation response processing error`, func() {
				dataProductExchangeApiServiceService, serviceErr := dataproductexchangeapiservicev1.NewDataProductExchangeApiServiceV1(&dataproductexchangeapiservicev1.DataProductExchangeApiServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductExchangeApiServiceService).ToNot(BeNil())

				// Construct an instance of the ItemReference model
				itemReferenceModel := new(dataproductexchangeapiservicev1.ItemReference)
				itemReferenceModel.ID = core.StringPtr("d29c42eb-7100-4b7a-8257-c196dbcca1cd")

				// Construct an instance of the OrderReference model
				orderReferenceModel := new(dataproductexchangeapiservicev1.OrderReference)
				orderReferenceModel.ID = core.StringPtr("4705e047-1808-459a-805f-d5d13c947637")
				orderReferenceModel.Items = []dataproductexchangeapiservicev1.ItemReference{*itemReferenceModel}

				// Construct an instance of the DeliverDataProductVersionOptions model
				deliverDataProductVersionOptionsModel := new(dataproductexchangeapiservicev1.DeliverDataProductVersionOptions)
				deliverDataProductVersionOptionsModel.ID = core.StringPtr("testString")
				deliverDataProductVersionOptionsModel.Order = orderReferenceModel
				deliverDataProductVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := dataProductExchangeApiServiceService.DeliverDataProductVersion(deliverDataProductVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				dataProductExchangeApiServiceService.EnableRetries(0, 0)
				result, response, operationErr = dataProductExchangeApiServiceService.DeliverDataProductVersion(deliverDataProductVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeliverDataProductVersion(deliverDataProductVersionOptions *DeliverDataProductVersionOptions)`, func() {
		deliverDataProductVersionPath := "/data_product_exchange/v1/data_product_versions/testString/deliver"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deliverDataProductVersionPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"status": "not_started", "href": "Href"}`)
				}))
			})
			It(`Invoke DeliverDataProductVersion successfully with retries`, func() {
				dataProductExchangeApiServiceService, serviceErr := dataproductexchangeapiservicev1.NewDataProductExchangeApiServiceV1(&dataproductexchangeapiservicev1.DataProductExchangeApiServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductExchangeApiServiceService).ToNot(BeNil())
				dataProductExchangeApiServiceService.EnableRetries(0, 0)

				// Construct an instance of the ItemReference model
				itemReferenceModel := new(dataproductexchangeapiservicev1.ItemReference)
				itemReferenceModel.ID = core.StringPtr("d29c42eb-7100-4b7a-8257-c196dbcca1cd")

				// Construct an instance of the OrderReference model
				orderReferenceModel := new(dataproductexchangeapiservicev1.OrderReference)
				orderReferenceModel.ID = core.StringPtr("4705e047-1808-459a-805f-d5d13c947637")
				orderReferenceModel.Items = []dataproductexchangeapiservicev1.ItemReference{*itemReferenceModel}

				// Construct an instance of the DeliverDataProductVersionOptions model
				deliverDataProductVersionOptionsModel := new(dataproductexchangeapiservicev1.DeliverDataProductVersionOptions)
				deliverDataProductVersionOptionsModel.ID = core.StringPtr("testString")
				deliverDataProductVersionOptionsModel.Order = orderReferenceModel
				deliverDataProductVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := dataProductExchangeApiServiceService.DeliverDataProductVersionWithContext(ctx, deliverDataProductVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				dataProductExchangeApiServiceService.DisableRetries()
				result, response, operationErr := dataProductExchangeApiServiceService.DeliverDataProductVersion(deliverDataProductVersionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = dataProductExchangeApiServiceService.DeliverDataProductVersionWithContext(ctx, deliverDataProductVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deliverDataProductVersionPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"status": "not_started", "href": "Href"}`)
				}))
			})
			It(`Invoke DeliverDataProductVersion successfully`, func() {
				dataProductExchangeApiServiceService, serviceErr := dataproductexchangeapiservicev1.NewDataProductExchangeApiServiceV1(&dataproductexchangeapiservicev1.DataProductExchangeApiServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductExchangeApiServiceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := dataProductExchangeApiServiceService.DeliverDataProductVersion(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ItemReference model
				itemReferenceModel := new(dataproductexchangeapiservicev1.ItemReference)
				itemReferenceModel.ID = core.StringPtr("d29c42eb-7100-4b7a-8257-c196dbcca1cd")

				// Construct an instance of the OrderReference model
				orderReferenceModel := new(dataproductexchangeapiservicev1.OrderReference)
				orderReferenceModel.ID = core.StringPtr("4705e047-1808-459a-805f-d5d13c947637")
				orderReferenceModel.Items = []dataproductexchangeapiservicev1.ItemReference{*itemReferenceModel}

				// Construct an instance of the DeliverDataProductVersionOptions model
				deliverDataProductVersionOptionsModel := new(dataproductexchangeapiservicev1.DeliverDataProductVersionOptions)
				deliverDataProductVersionOptionsModel.ID = core.StringPtr("testString")
				deliverDataProductVersionOptionsModel.Order = orderReferenceModel
				deliverDataProductVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = dataProductExchangeApiServiceService.DeliverDataProductVersion(deliverDataProductVersionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke DeliverDataProductVersion with error: Operation validation and request error`, func() {
				dataProductExchangeApiServiceService, serviceErr := dataproductexchangeapiservicev1.NewDataProductExchangeApiServiceV1(&dataproductexchangeapiservicev1.DataProductExchangeApiServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductExchangeApiServiceService).ToNot(BeNil())

				// Construct an instance of the ItemReference model
				itemReferenceModel := new(dataproductexchangeapiservicev1.ItemReference)
				itemReferenceModel.ID = core.StringPtr("d29c42eb-7100-4b7a-8257-c196dbcca1cd")

				// Construct an instance of the OrderReference model
				orderReferenceModel := new(dataproductexchangeapiservicev1.OrderReference)
				orderReferenceModel.ID = core.StringPtr("4705e047-1808-459a-805f-d5d13c947637")
				orderReferenceModel.Items = []dataproductexchangeapiservicev1.ItemReference{*itemReferenceModel}

				// Construct an instance of the DeliverDataProductVersionOptions model
				deliverDataProductVersionOptionsModel := new(dataproductexchangeapiservicev1.DeliverDataProductVersionOptions)
				deliverDataProductVersionOptionsModel.ID = core.StringPtr("testString")
				deliverDataProductVersionOptionsModel.Order = orderReferenceModel
				deliverDataProductVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dataProductExchangeApiServiceService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := dataProductExchangeApiServiceService.DeliverDataProductVersion(deliverDataProductVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeliverDataProductVersionOptions model with no property values
				deliverDataProductVersionOptionsModelNew := new(dataproductexchangeapiservicev1.DeliverDataProductVersionOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = dataProductExchangeApiServiceService.DeliverDataProductVersion(deliverDataProductVersionOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(202)
				}))
			})
			It(`Invoke DeliverDataProductVersion successfully`, func() {
				dataProductExchangeApiServiceService, serviceErr := dataproductexchangeapiservicev1.NewDataProductExchangeApiServiceV1(&dataproductexchangeapiservicev1.DataProductExchangeApiServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductExchangeApiServiceService).ToNot(BeNil())

				// Construct an instance of the ItemReference model
				itemReferenceModel := new(dataproductexchangeapiservicev1.ItemReference)
				itemReferenceModel.ID = core.StringPtr("d29c42eb-7100-4b7a-8257-c196dbcca1cd")

				// Construct an instance of the OrderReference model
				orderReferenceModel := new(dataproductexchangeapiservicev1.OrderReference)
				orderReferenceModel.ID = core.StringPtr("4705e047-1808-459a-805f-d5d13c947637")
				orderReferenceModel.Items = []dataproductexchangeapiservicev1.ItemReference{*itemReferenceModel}

				// Construct an instance of the DeliverDataProductVersionOptions model
				deliverDataProductVersionOptionsModel := new(dataproductexchangeapiservicev1.DeliverDataProductVersionOptions)
				deliverDataProductVersionOptionsModel.ID = core.StringPtr("testString")
				deliverDataProductVersionOptionsModel.Order = orderReferenceModel
				deliverDataProductVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := dataProductExchangeApiServiceService.DeliverDataProductVersion(deliverDataProductVersionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Model constructor tests`, func() {
		Context(`Using a service client instance`, func() {
			dataProductExchangeApiServiceService, _ := dataproductexchangeapiservicev1.NewDataProductExchangeApiServiceV1(&dataproductexchangeapiservicev1.DataProductExchangeApiServiceV1Options{
				URL:           "http://dataproductexchangeapiservicev1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewAssetPartReference successfully`, func() {
				id := "2b0bf220-079c-11ee-be56-0242ac120002"
				var container *dataproductexchangeapiservicev1.ContainerReference = nil
				_, err := dataProductExchangeApiServiceService.NewAssetPartReference(id, container)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewContainerReference successfully`, func() {
				id := "d29c42eb-7100-4b7a-8257-c196dbcca1cd"
				_model, err := dataProductExchangeApiServiceService.NewContainerReference(id)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewCreateDataProductVersionOptions successfully`, func() {
				// Construct an instance of the ContainerReference model
				containerReferenceModel := new(dataproductexchangeapiservicev1.ContainerReference)
				Expect(containerReferenceModel).ToNot(BeNil())
				containerReferenceModel.ID = core.StringPtr("d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				containerReferenceModel.Type = core.StringPtr("catalog")
				Expect(containerReferenceModel.ID).To(Equal(core.StringPtr("d29c42eb-7100-4b7a-8257-c196dbcca1cd")))
				Expect(containerReferenceModel.Type).To(Equal(core.StringPtr("catalog")))

				// Construct an instance of the DataProductIdentity model
				dataProductIdentityModel := new(dataproductexchangeapiservicev1.DataProductIdentity)
				Expect(dataProductIdentityModel).ToNot(BeNil())
				dataProductIdentityModel.ID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")
				Expect(dataProductIdentityModel.ID).To(Equal(core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")))

				// Construct an instance of the UseCase model
				useCaseModel := new(dataproductexchangeapiservicev1.UseCase)
				Expect(useCaseModel).ToNot(BeNil())
				useCaseModel.ID = core.StringPtr("testString")
				useCaseModel.Name = core.StringPtr("testString")
				useCaseModel.Container = containerReferenceModel
				Expect(useCaseModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(useCaseModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(useCaseModel.Container).To(Equal(containerReferenceModel))

				// Construct an instance of the Domain model
				domainModel := new(dataproductexchangeapiservicev1.Domain)
				Expect(domainModel).ToNot(BeNil())
				domainModel.ID = core.StringPtr("testString")
				domainModel.Name = core.StringPtr("testString")
				domainModel.Container = containerReferenceModel
				Expect(domainModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(domainModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(domainModel.Container).To(Equal(containerReferenceModel))

				// Construct an instance of the AssetPartReference model
				assetPartReferenceModel := new(dataproductexchangeapiservicev1.AssetPartReference)
				Expect(assetPartReferenceModel).ToNot(BeNil())
				assetPartReferenceModel.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				assetPartReferenceModel.Container = containerReferenceModel
				assetPartReferenceModel.Type = core.StringPtr("data_asset")
				Expect(assetPartReferenceModel.ID).To(Equal(core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")))
				Expect(assetPartReferenceModel.Container).To(Equal(containerReferenceModel))
				Expect(assetPartReferenceModel.Type).To(Equal(core.StringPtr("data_asset")))

				// Construct an instance of the DeliveryMethod model
				deliveryMethodModel := new(dataproductexchangeapiservicev1.DeliveryMethod)
				Expect(deliveryMethodModel).ToNot(BeNil())
				deliveryMethodModel.ID = core.StringPtr("09cf5fcc-cb9d-4995-a8e4-16517b25229f")
				deliveryMethodModel.Container = containerReferenceModel
				Expect(deliveryMethodModel.ID).To(Equal(core.StringPtr("09cf5fcc-cb9d-4995-a8e4-16517b25229f")))
				Expect(deliveryMethodModel.Container).To(Equal(containerReferenceModel))

				// Construct an instance of the DataProductPart model
				dataProductPartModel := new(dataproductexchangeapiservicev1.DataProductPart)
				Expect(dataProductPartModel).ToNot(BeNil())
				dataProductPartModel.Asset = assetPartReferenceModel
				dataProductPartModel.Revision = core.Int64Ptr(int64(1))
				dataProductPartModel.UpdatedAt = CreateMockDateTime("2023-07-01T22:22:34.876Z")
				dataProductPartModel.DeliveryMethods = []dataproductexchangeapiservicev1.DeliveryMethod{*deliveryMethodModel}
				Expect(dataProductPartModel.Asset).To(Equal(assetPartReferenceModel))
				Expect(dataProductPartModel.Revision).To(Equal(core.Int64Ptr(int64(1))))
				Expect(dataProductPartModel.UpdatedAt).To(Equal(CreateMockDateTime("2023-07-01T22:22:34.876Z")))
				Expect(dataProductPartModel.DeliveryMethods).To(Equal([]dataproductexchangeapiservicev1.DeliveryMethod{*deliveryMethodModel}))

				// Construct an instance of the CreateDataProductVersionOptions model
				var createDataProductVersionOptionsContainer *dataproductexchangeapiservicev1.ContainerReference = nil
				createDataProductVersionOptionsModel := dataProductExchangeApiServiceService.NewCreateDataProductVersionOptions(createDataProductVersionOptionsContainer)
				createDataProductVersionOptionsModel.SetContainer(containerReferenceModel)
				createDataProductVersionOptionsModel.SetVersion("testString")
				createDataProductVersionOptionsModel.SetState("draft")
				createDataProductVersionOptionsModel.SetDataProduct(dataProductIdentityModel)
				createDataProductVersionOptionsModel.SetName("My New Data Product")
				createDataProductVersionOptionsModel.SetDescription("testString")
				createDataProductVersionOptionsModel.SetTags([]string{"testString"})
				createDataProductVersionOptionsModel.SetUseCases([]dataproductexchangeapiservicev1.UseCase{*useCaseModel})
				createDataProductVersionOptionsModel.SetDomain(domainModel)
				createDataProductVersionOptionsModel.SetType([]string{"data"})
				createDataProductVersionOptionsModel.SetPartsOut([]dataproductexchangeapiservicev1.DataProductPart{*dataProductPartModel})
				createDataProductVersionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createDataProductVersionOptionsModel).ToNot(BeNil())
				Expect(createDataProductVersionOptionsModel.Container).To(Equal(containerReferenceModel))
				Expect(createDataProductVersionOptionsModel.Version).To(Equal(core.StringPtr("testString")))
				Expect(createDataProductVersionOptionsModel.State).To(Equal(core.StringPtr("draft")))
				Expect(createDataProductVersionOptionsModel.DataProduct).To(Equal(dataProductIdentityModel))
				Expect(createDataProductVersionOptionsModel.Name).To(Equal(core.StringPtr("My New Data Product")))
				Expect(createDataProductVersionOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(createDataProductVersionOptionsModel.Tags).To(Equal([]string{"testString"}))
				Expect(createDataProductVersionOptionsModel.UseCases).To(Equal([]dataproductexchangeapiservicev1.UseCase{*useCaseModel}))
				Expect(createDataProductVersionOptionsModel.Domain).To(Equal(domainModel))
				Expect(createDataProductVersionOptionsModel.Type).To(Equal([]string{"data"}))
				Expect(createDataProductVersionOptionsModel.PartsOut).To(Equal([]dataproductexchangeapiservicev1.DataProductPart{*dataProductPartModel}))
				Expect(createDataProductVersionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDataProductIdentity successfully`, func() {
				id := "b38df608-d34b-4d58-8136-ed25e6c6684e"
				_model, err := dataProductExchangeApiServiceService.NewDataProductIdentity(id)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewDataProductPart successfully`, func() {
				var asset *dataproductexchangeapiservicev1.AssetPartReference = nil
				_, err := dataProductExchangeApiServiceService.NewDataProductPart(asset)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewDataProductVersionPatch successfully`, func() {
				// Construct an instance of the DataProductIdentity model
				dataProductIdentityModel := new(dataproductexchangeapiservicev1.DataProductIdentity)
				dataProductIdentityModel.ID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")

				// Construct an instance of the ContainerReference model
				containerReferenceModel := new(dataproductexchangeapiservicev1.ContainerReference)
				containerReferenceModel.ID = core.StringPtr("d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				containerReferenceModel.Type = core.StringPtr("catalog")

				// Construct an instance of the AssetReference model
				assetReferenceModel := new(dataproductexchangeapiservicev1.AssetReference)
				assetReferenceModel.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				assetReferenceModel.Container = containerReferenceModel

				// Construct an instance of the UseCase model
				useCaseModel := new(dataproductexchangeapiservicev1.UseCase)
				useCaseModel.ID = core.StringPtr("testString")
				useCaseModel.Name = core.StringPtr("testString")
				useCaseModel.Container = containerReferenceModel

				// Construct an instance of the Domain model
				domainModel := new(dataproductexchangeapiservicev1.Domain)
				domainModel.ID = core.StringPtr("testString")
				domainModel.Name = core.StringPtr("testString")
				domainModel.Container = containerReferenceModel

				// Construct an instance of the AssetPartReference model
				assetPartReferenceModel := new(dataproductexchangeapiservicev1.AssetPartReference)
				assetPartReferenceModel.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				assetPartReferenceModel.Container = containerReferenceModel
				assetPartReferenceModel.Type = core.StringPtr("data_asset")

				// Construct an instance of the DeliveryMethod model
				deliveryMethodModel := new(dataproductexchangeapiservicev1.DeliveryMethod)
				deliveryMethodModel.ID = core.StringPtr("09cf5fcc-cb9d-4995-a8e4-16517b25229f")
				deliveryMethodModel.Container = containerReferenceModel

				// Construct an instance of the DataProductPart model
				dataProductPartModel := new(dataproductexchangeapiservicev1.DataProductPart)
				dataProductPartModel.Asset = assetPartReferenceModel
				dataProductPartModel.Revision = core.Int64Ptr(int64(1))
				dataProductPartModel.UpdatedAt = CreateMockDateTime("2023-07-01T22:22:34.876Z")
				dataProductPartModel.DeliveryMethods = []dataproductexchangeapiservicev1.DeliveryMethod{*deliveryMethodModel}

				// Construct an instance of the DataProductVersion model
				dataProductVersion := new(dataproductexchangeapiservicev1.DataProductVersion)
				dataProductVersion.Version = core.StringPtr("1.0.0")
				dataProductVersion.State = core.StringPtr("draft")
				dataProductVersion.DataProduct = dataProductIdentityModel
				dataProductVersion.Name = core.StringPtr("My Data Product")
				dataProductVersion.Description = core.StringPtr("This is a description of My Data Product.")
				dataProductVersion.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				dataProductVersion.Asset = assetReferenceModel
				dataProductVersion.Tags = []string{"testString"}
				dataProductVersion.UseCases = []dataproductexchangeapiservicev1.UseCase{*useCaseModel}
				dataProductVersion.Domain = domainModel
				dataProductVersion.Type = []string{"data"}
				dataProductVersion.PartsOut = []dataproductexchangeapiservicev1.DataProductPart{*dataProductPartModel}
				dataProductVersion.PublishedBy = core.StringPtr("testString")
				dataProductVersion.PublishedAt = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				dataProductVersion.CreatedBy = core.StringPtr("testString")
				dataProductVersion.CreatedAt = CreateMockDateTime("2019-01-01T12:00:00.000Z")

				dataProductVersionPatch := dataProductExchangeApiServiceService.NewDataProductVersionPatch(dataProductVersion)
				Expect(dataProductVersionPatch).ToNot(BeNil())

				_path := func(op interface{}) string {
					return *op.(dataproductexchangeapiservicev1.JSONPatchOperation).Path
				}
				Expect(dataProductVersionPatch).To(MatchAllElements(_path, Elements{
				"/version": MatchAllFields(Fields{
					"Op": PointTo(Equal(dataproductexchangeapiservicev1.JSONPatchOperation_Op_Add)),
					"Path": PointTo(Equal("/version")),
					"From": BeNil(),
					"Value": Equal(dataProductVersion.Version),
					}),
				"/state": MatchAllFields(Fields{
					"Op": PointTo(Equal(dataproductexchangeapiservicev1.JSONPatchOperation_Op_Add)),
					"Path": PointTo(Equal("/state")),
					"From": BeNil(),
					"Value": Equal(dataProductVersion.State),
					}),
				"/data_product": MatchAllFields(Fields{
					"Op": PointTo(Equal(dataproductexchangeapiservicev1.JSONPatchOperation_Op_Add)),
					"Path": PointTo(Equal("/data_product")),
					"From": BeNil(),
					"Value": Equal(dataProductVersion.DataProduct),
					}),
				"/name": MatchAllFields(Fields{
					"Op": PointTo(Equal(dataproductexchangeapiservicev1.JSONPatchOperation_Op_Add)),
					"Path": PointTo(Equal("/name")),
					"From": BeNil(),
					"Value": Equal(dataProductVersion.Name),
					}),
				"/description": MatchAllFields(Fields{
					"Op": PointTo(Equal(dataproductexchangeapiservicev1.JSONPatchOperation_Op_Add)),
					"Path": PointTo(Equal("/description")),
					"From": BeNil(),
					"Value": Equal(dataProductVersion.Description),
					}),
				"/id": MatchAllFields(Fields{
					"Op": PointTo(Equal(dataproductexchangeapiservicev1.JSONPatchOperation_Op_Add)),
					"Path": PointTo(Equal("/id")),
					"From": BeNil(),
					"Value": Equal(dataProductVersion.ID),
					}),
				"/asset": MatchAllFields(Fields{
					"Op": PointTo(Equal(dataproductexchangeapiservicev1.JSONPatchOperation_Op_Add)),
					"Path": PointTo(Equal("/asset")),
					"From": BeNil(),
					"Value": Equal(dataProductVersion.Asset),
					}),
				"/tags": MatchAllFields(Fields{
					"Op": PointTo(Equal(dataproductexchangeapiservicev1.JSONPatchOperation_Op_Add)),
					"Path": PointTo(Equal("/tags")),
					"From": BeNil(),
					"Value": Equal(dataProductVersion.Tags),
					}),
				"/use_cases": MatchAllFields(Fields{
					"Op": PointTo(Equal(dataproductexchangeapiservicev1.JSONPatchOperation_Op_Add)),
					"Path": PointTo(Equal("/use_cases")),
					"From": BeNil(),
					"Value": Equal(dataProductVersion.UseCases),
					}),
				"/domain": MatchAllFields(Fields{
					"Op": PointTo(Equal(dataproductexchangeapiservicev1.JSONPatchOperation_Op_Add)),
					"Path": PointTo(Equal("/domain")),
					"From": BeNil(),
					"Value": Equal(dataProductVersion.Domain),
					}),
				"/type": MatchAllFields(Fields{
					"Op": PointTo(Equal(dataproductexchangeapiservicev1.JSONPatchOperation_Op_Add)),
					"Path": PointTo(Equal("/type")),
					"From": BeNil(),
					"Value": Equal(dataProductVersion.Type),
					}),
				"/parts_out": MatchAllFields(Fields{
					"Op": PointTo(Equal(dataproductexchangeapiservicev1.JSONPatchOperation_Op_Add)),
					"Path": PointTo(Equal("/parts_out")),
					"From": BeNil(),
					"Value": Equal(dataProductVersion.PartsOut),
					}),
				"/published_by": MatchAllFields(Fields{
					"Op": PointTo(Equal(dataproductexchangeapiservicev1.JSONPatchOperation_Op_Add)),
					"Path": PointTo(Equal("/published_by")),
					"From": BeNil(),
					"Value": Equal(dataProductVersion.PublishedBy),
					}),
				"/published_at": MatchAllFields(Fields{
					"Op": PointTo(Equal(dataproductexchangeapiservicev1.JSONPatchOperation_Op_Add)),
					"Path": PointTo(Equal("/published_at")),
					"From": BeNil(),
					"Value": Equal(dataProductVersion.PublishedAt),
					}),
				"/created_by": MatchAllFields(Fields{
					"Op": PointTo(Equal(dataproductexchangeapiservicev1.JSONPatchOperation_Op_Add)),
					"Path": PointTo(Equal("/created_by")),
					"From": BeNil(),
					"Value": Equal(dataProductVersion.CreatedBy),
					}),
				"/created_at": MatchAllFields(Fields{
					"Op": PointTo(Equal(dataproductexchangeapiservicev1.JSONPatchOperation_Op_Add)),
					"Path": PointTo(Equal("/created_at")),
					"From": BeNil(),
					"Value": Equal(dataProductVersion.CreatedAt),
					}),
				}))
			})
			It(`Invoke NewDeleteDataProductVersionOptions successfully`, func() {
				// Construct an instance of the DeleteDataProductVersionOptions model
				id := "testString"
				deleteDataProductVersionOptionsModel := dataProductExchangeApiServiceService.NewDeleteDataProductVersionOptions(id)
				deleteDataProductVersionOptionsModel.SetID("testString")
				deleteDataProductVersionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteDataProductVersionOptionsModel).ToNot(BeNil())
				Expect(deleteDataProductVersionOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(deleteDataProductVersionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeliverDataProductVersionOptions successfully`, func() {
				// Construct an instance of the ItemReference model
				itemReferenceModel := new(dataproductexchangeapiservicev1.ItemReference)
				Expect(itemReferenceModel).ToNot(BeNil())
				itemReferenceModel.ID = core.StringPtr("d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				Expect(itemReferenceModel.ID).To(Equal(core.StringPtr("d29c42eb-7100-4b7a-8257-c196dbcca1cd")))

				// Construct an instance of the OrderReference model
				orderReferenceModel := new(dataproductexchangeapiservicev1.OrderReference)
				Expect(orderReferenceModel).ToNot(BeNil())
				orderReferenceModel.ID = core.StringPtr("4705e047-1808-459a-805f-d5d13c947637")
				orderReferenceModel.Items = []dataproductexchangeapiservicev1.ItemReference{*itemReferenceModel}
				Expect(orderReferenceModel.ID).To(Equal(core.StringPtr("4705e047-1808-459a-805f-d5d13c947637")))
				Expect(orderReferenceModel.Items).To(Equal([]dataproductexchangeapiservicev1.ItemReference{*itemReferenceModel}))

				// Construct an instance of the DeliverDataProductVersionOptions model
				id := "testString"
				deliverDataProductVersionOptionsModel := dataProductExchangeApiServiceService.NewDeliverDataProductVersionOptions(id)
				deliverDataProductVersionOptionsModel.SetID("testString")
				deliverDataProductVersionOptionsModel.SetOrder(orderReferenceModel)
				deliverDataProductVersionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deliverDataProductVersionOptionsModel).ToNot(BeNil())
				Expect(deliverDataProductVersionOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(deliverDataProductVersionOptionsModel.Order).To(Equal(orderReferenceModel))
				Expect(deliverDataProductVersionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeliveryMethod successfully`, func() {
				id := "09cf5fcc-cb9d-4995-a8e4-16517b25229f"
				var container *dataproductexchangeapiservicev1.ContainerReference = nil
				_, err := dataProductExchangeApiServiceService.NewDeliveryMethod(id, container)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewDomain successfully`, func() {
				id := "testString"
				name := "testString"
				_model, err := dataProductExchangeApiServiceService.NewDomain(id, name)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewGetDataProductOptions successfully`, func() {
				// Construct an instance of the GetDataProductOptions model
				id := "testString"
				getDataProductOptionsModel := dataProductExchangeApiServiceService.NewGetDataProductOptions(id)
				getDataProductOptionsModel.SetID("testString")
				getDataProductOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getDataProductOptionsModel).ToNot(BeNil())
				Expect(getDataProductOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getDataProductOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetDataProductVersionOptions successfully`, func() {
				// Construct an instance of the GetDataProductVersionOptions model
				id := "testString"
				getDataProductVersionOptionsModel := dataProductExchangeApiServiceService.NewGetDataProductVersionOptions(id)
				getDataProductVersionOptionsModel.SetID("testString")
				getDataProductVersionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getDataProductVersionOptionsModel).ToNot(BeNil())
				Expect(getDataProductVersionOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getDataProductVersionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetInitializeStatusOptions successfully`, func() {
				// Construct an instance of the GetInitializeStatusOptions model
				getInitializeStatusOptionsModel := dataProductExchangeApiServiceService.NewGetInitializeStatusOptions()
				getInitializeStatusOptionsModel.SetContainerID("testString")
				getInitializeStatusOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getInitializeStatusOptionsModel).ToNot(BeNil())
				Expect(getInitializeStatusOptionsModel.ContainerID).To(Equal(core.StringPtr("testString")))
				Expect(getInitializeStatusOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewInitializeOptions successfully`, func() {
				// Construct an instance of the ContainerReference model
				containerReferenceModel := new(dataproductexchangeapiservicev1.ContainerReference)
				Expect(containerReferenceModel).ToNot(BeNil())
				containerReferenceModel.ID = core.StringPtr("d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				containerReferenceModel.Type = core.StringPtr("catalog")
				Expect(containerReferenceModel.ID).To(Equal(core.StringPtr("d29c42eb-7100-4b7a-8257-c196dbcca1cd")))
				Expect(containerReferenceModel.Type).To(Equal(core.StringPtr("catalog")))

				// Construct an instance of the InitializeOptions model
				initializeOptionsModel := dataProductExchangeApiServiceService.NewInitializeOptions()
				initializeOptionsModel.SetContainer(containerReferenceModel)
				initializeOptionsModel.SetInclude([]string{"delivery_methods", "data_product_samples", "domains_multi_industry"})
				initializeOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(initializeOptionsModel).ToNot(BeNil())
				Expect(initializeOptionsModel.Container).To(Equal(containerReferenceModel))
				Expect(initializeOptionsModel.Include).To(Equal([]string{"delivery_methods", "data_product_samples", "domains_multi_industry"}))
				Expect(initializeOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewItemReference successfully`, func() {
				id := "2b0bf220-079c-11ee-be56-0242ac120002"
				_model, err := dataProductExchangeApiServiceService.NewItemReference(id)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewJSONPatchOperation successfully`, func() {
				op := "add"
				path := "testString"
				_model, err := dataProductExchangeApiServiceService.NewJSONPatchOperation(op, path)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewListDataProductVersionsOptions successfully`, func() {
				// Construct an instance of the ListDataProductVersionsOptions model
				listDataProductVersionsOptionsModel := dataProductExchangeApiServiceService.NewListDataProductVersionsOptions()
				listDataProductVersionsOptionsModel.SetAssetContainerID("testString")
				listDataProductVersionsOptionsModel.SetDataProduct("testString")
				listDataProductVersionsOptionsModel.SetState("draft")
				listDataProductVersionsOptionsModel.SetVersion("testString")
				listDataProductVersionsOptionsModel.SetLimit(int64(10))
				listDataProductVersionsOptionsModel.SetStart("testString")
				listDataProductVersionsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listDataProductVersionsOptionsModel).ToNot(BeNil())
				Expect(listDataProductVersionsOptionsModel.AssetContainerID).To(Equal(core.StringPtr("testString")))
				Expect(listDataProductVersionsOptionsModel.DataProduct).To(Equal(core.StringPtr("testString")))
				Expect(listDataProductVersionsOptionsModel.State).To(Equal(core.StringPtr("draft")))
				Expect(listDataProductVersionsOptionsModel.Version).To(Equal(core.StringPtr("testString")))
				Expect(listDataProductVersionsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(10))))
				Expect(listDataProductVersionsOptionsModel.Start).To(Equal(core.StringPtr("testString")))
				Expect(listDataProductVersionsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListDataProductsOptions successfully`, func() {
				// Construct an instance of the ListDataProductsOptions model
				listDataProductsOptionsModel := dataProductExchangeApiServiceService.NewListDataProductsOptions()
				listDataProductsOptionsModel.SetLimit(int64(10))
				listDataProductsOptionsModel.SetStart("testString")
				listDataProductsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listDataProductsOptionsModel).ToNot(BeNil())
				Expect(listDataProductsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(10))))
				Expect(listDataProductsOptionsModel.Start).To(Equal(core.StringPtr("testString")))
				Expect(listDataProductsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewOrderReference successfully`, func() {
				id := "2b0bf220-079c-11ee-be56-0242ac120002"
				_model, err := dataProductExchangeApiServiceService.NewOrderReference(id)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewUpdateDataProductVersionOptions successfully`, func() {
				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(dataproductexchangeapiservicev1.JSONPatchOperation)
				Expect(jsonPatchOperationModel).ToNot(BeNil())
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = core.StringPtr("testString")
				Expect(jsonPatchOperationModel.Op).To(Equal(core.StringPtr("add")))
				Expect(jsonPatchOperationModel.Path).To(Equal(core.StringPtr("testString")))
				Expect(jsonPatchOperationModel.From).To(Equal(core.StringPtr("testString")))
				Expect(jsonPatchOperationModel.Value).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the UpdateDataProductVersionOptions model
				id := "testString"
				jsonPatchInstructions := []dataproductexchangeapiservicev1.JSONPatchOperation{}
				updateDataProductVersionOptionsModel := dataProductExchangeApiServiceService.NewUpdateDataProductVersionOptions(id, jsonPatchInstructions)
				updateDataProductVersionOptionsModel.SetID("testString")
				updateDataProductVersionOptionsModel.SetJSONPatchInstructions([]dataproductexchangeapiservicev1.JSONPatchOperation{*jsonPatchOperationModel})
				updateDataProductVersionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateDataProductVersionOptionsModel).ToNot(BeNil())
				Expect(updateDataProductVersionOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(updateDataProductVersionOptionsModel.JSONPatchInstructions).To(Equal([]dataproductexchangeapiservicev1.JSONPatchOperation{*jsonPatchOperationModel}))
				Expect(updateDataProductVersionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUseCase successfully`, func() {
				id := "testString"
				name := "testString"
				_model, err := dataProductExchangeApiServiceService.NewUseCase(id, name)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
		})
	})
	Describe(`Utility function tests`, func() {
		It(`Invoke CreateMockByteArray() successfully`, func() {
			mockByteArray := CreateMockByteArray("This is a test")
			Expect(mockByteArray).ToNot(BeNil())
		})
		It(`Invoke CreateMockUUID() successfully`, func() {
			mockUUID := CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
			Expect(mockUUID).ToNot(BeNil())
		})
		It(`Invoke CreateMockReader() successfully`, func() {
			mockReader := CreateMockReader("This is a test.")
			Expect(mockReader).ToNot(BeNil())
		})
		It(`Invoke CreateMockDate() successfully`, func() {
			mockDate := CreateMockDate("2019-01-01")
			Expect(mockDate).ToNot(BeNil())
		})
		It(`Invoke CreateMockDateTime() successfully`, func() {
			mockDateTime := CreateMockDateTime("2019-01-01T12:00:00.000Z")
			Expect(mockDateTime).ToNot(BeNil())
		})
	})
})

//
// Utility functions used by the generated test code
//

func CreateMockByteArray(mockData string) *[]byte {
	ba := make([]byte, 0)
	ba = append(ba, mockData...)
	return &ba
}

func CreateMockUUID(mockData string) *strfmt.UUID {
	uuid := strfmt.UUID(mockData)
	return &uuid
}

func CreateMockReader(mockData string) io.ReadCloser {
	return io.NopCloser(bytes.NewReader([]byte(mockData)))
}

func CreateMockDate(mockData string) *strfmt.Date {
	d, err := core.ParseDate(mockData)
	if err != nil {
		return nil
	}
	return &d
}

func CreateMockDateTime(mockData string) *strfmt.DateTime {
	d, err := core.ParseDateTime(mockData)
	if err != nil {
		return nil
	}
	return &d
}

func SetTestEnvironment(testEnvironment map[string]string) {
	for key, value := range testEnvironment {
		os.Setenv(key, value)
	}
}

func ClearTestEnvironment(testEnvironment map[string]string) {
	for key := range testEnvironment {
		os.Unsetenv(key)
	}
}
