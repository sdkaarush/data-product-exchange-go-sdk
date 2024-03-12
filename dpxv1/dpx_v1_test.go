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
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	"github.com/IBM/data-product-exchange-go-sdk/dpxv1"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gstruct"
)

var _ = Describe(`DpxV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(dpxService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(dpxService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
				URL: "https://dpxv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(dpxService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"DPX_URL":       "https://dpxv1/api",
				"DPX_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				dpxService, serviceErr := dpxv1.NewDpxV1UsingExternalConfig(&dpxv1.DpxV1Options{})
				Expect(dpxService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := dpxService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != dpxService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(dpxService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(dpxService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				dpxService, serviceErr := dpxv1.NewDpxV1UsingExternalConfig(&dpxv1.DpxV1Options{
					URL: "https://testService/api",
				})
				Expect(dpxService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(dpxService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := dpxService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != dpxService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(dpxService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(dpxService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				dpxService, serviceErr := dpxv1.NewDpxV1UsingExternalConfig(&dpxv1.DpxV1Options{})
				err := dpxService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(dpxService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(dpxService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := dpxService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != dpxService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(dpxService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(dpxService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"DPX_URL":       "https://dpxv1/api",
				"DPX_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			dpxService, serviceErr := dpxv1.NewDpxV1UsingExternalConfig(&dpxv1.DpxV1Options{})

			It(`Instantiate service client with error`, func() {
				Expect(dpxService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"DPX_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			dpxService, serviceErr := dpxv1.NewDpxV1UsingExternalConfig(&dpxv1.DpxV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(dpxService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = dpxv1.GetServiceURLForRegion("INVALID_REGION")
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
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Construct an instance of the GetInitializeStatusOptions model
				getInitializeStatusOptionsModel := new(dpxv1.GetInitializeStatusOptions)
				getInitializeStatusOptionsModel.ContainerID = core.StringPtr("testString")
				getInitializeStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := dpxService.GetInitializeStatus(getInitializeStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				dpxService.EnableRetries(0, 0)
				result, response, operationErr = dpxService.GetInitializeStatus(getInitializeStatusOptionsModel)
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
					fmt.Fprintf(res, "%s", `{"container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}, "href": "https://api.example.com/configuration/initialize/status?catalog_id=d29c42eb-7100-4b7a-8257-c196dbcca1cd", "status": "not_started", "trace": "Trace", "errors": [{"code": "request_body_error", "message": "Message", "extra": {"anyKey": "anyValue"}, "more_info": "MoreInfo"}], "last_started_at": "2023-08-21T15:24:06.021Z", "last_finished_at": "2023-08-21T20:24:34.450Z", "initialized_options": [{"name": "Name", "version": 1}]}`)
				}))
			})
			It(`Invoke GetInitializeStatus successfully with retries`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())
				dpxService.EnableRetries(0, 0)

				// Construct an instance of the GetInitializeStatusOptions model
				getInitializeStatusOptionsModel := new(dpxv1.GetInitializeStatusOptions)
				getInitializeStatusOptionsModel.ContainerID = core.StringPtr("testString")
				getInitializeStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := dpxService.GetInitializeStatusWithContext(ctx, getInitializeStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				dpxService.DisableRetries()
				result, response, operationErr := dpxService.GetInitializeStatus(getInitializeStatusOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = dpxService.GetInitializeStatusWithContext(ctx, getInitializeStatusOptionsModel)
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
					fmt.Fprintf(res, "%s", `{"container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}, "href": "https://api.example.com/configuration/initialize/status?catalog_id=d29c42eb-7100-4b7a-8257-c196dbcca1cd", "status": "not_started", "trace": "Trace", "errors": [{"code": "request_body_error", "message": "Message", "extra": {"anyKey": "anyValue"}, "more_info": "MoreInfo"}], "last_started_at": "2023-08-21T15:24:06.021Z", "last_finished_at": "2023-08-21T20:24:34.450Z", "initialized_options": [{"name": "Name", "version": 1}]}`)
				}))
			})
			It(`Invoke GetInitializeStatus successfully`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := dpxService.GetInitializeStatus(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetInitializeStatusOptions model
				getInitializeStatusOptionsModel := new(dpxv1.GetInitializeStatusOptions)
				getInitializeStatusOptionsModel.ContainerID = core.StringPtr("testString")
				getInitializeStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = dpxService.GetInitializeStatus(getInitializeStatusOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetInitializeStatus with error: Operation request error`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Construct an instance of the GetInitializeStatusOptions model
				getInitializeStatusOptionsModel := new(dpxv1.GetInitializeStatusOptions)
				getInitializeStatusOptionsModel.ContainerID = core.StringPtr("testString")
				getInitializeStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dpxService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := dpxService.GetInitializeStatus(getInitializeStatusOptionsModel)
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
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Construct an instance of the GetInitializeStatusOptions model
				getInitializeStatusOptionsModel := new(dpxv1.GetInitializeStatusOptions)
				getInitializeStatusOptionsModel.ContainerID = core.StringPtr("testString")
				getInitializeStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := dpxService.GetInitializeStatus(getInitializeStatusOptionsModel)
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
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Construct an instance of the ContainerReference model
				containerReferenceModel := new(dpxv1.ContainerReference)
				containerReferenceModel.ID = core.StringPtr("d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				containerReferenceModel.Type = core.StringPtr("catalog")

				// Construct an instance of the InitializeOptions model
				initializeOptionsModel := new(dpxv1.InitializeOptions)
				initializeOptionsModel.Container = containerReferenceModel
				initializeOptionsModel.Include = []string{"delivery_methods", "data_product_samples", "domains_multi_industry"}
				initializeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := dpxService.Initialize(initializeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				dpxService.EnableRetries(0, 0)
				result, response, operationErr = dpxService.Initialize(initializeOptionsModel)
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
					fmt.Fprintf(res, "%s", `{"container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}, "href": "https://api.example.com/configuration/initialize/status?catalog_id=d29c42eb-7100-4b7a-8257-c196dbcca1cd", "status": "not_started", "trace": "Trace", "errors": [{"code": "request_body_error", "message": "Message", "extra": {"anyKey": "anyValue"}, "more_info": "MoreInfo"}], "last_started_at": "2023-08-21T15:24:06.021Z", "last_finished_at": "2023-08-21T20:24:34.450Z", "initialized_options": [{"name": "Name", "version": 1}]}`)
				}))
			})
			It(`Invoke Initialize successfully with retries`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())
				dpxService.EnableRetries(0, 0)

				// Construct an instance of the ContainerReference model
				containerReferenceModel := new(dpxv1.ContainerReference)
				containerReferenceModel.ID = core.StringPtr("d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				containerReferenceModel.Type = core.StringPtr("catalog")

				// Construct an instance of the InitializeOptions model
				initializeOptionsModel := new(dpxv1.InitializeOptions)
				initializeOptionsModel.Container = containerReferenceModel
				initializeOptionsModel.Include = []string{"delivery_methods", "data_product_samples", "domains_multi_industry"}
				initializeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := dpxService.InitializeWithContext(ctx, initializeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				dpxService.DisableRetries()
				result, response, operationErr := dpxService.Initialize(initializeOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = dpxService.InitializeWithContext(ctx, initializeOptionsModel)
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
					fmt.Fprintf(res, "%s", `{"container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}, "href": "https://api.example.com/configuration/initialize/status?catalog_id=d29c42eb-7100-4b7a-8257-c196dbcca1cd", "status": "not_started", "trace": "Trace", "errors": [{"code": "request_body_error", "message": "Message", "extra": {"anyKey": "anyValue"}, "more_info": "MoreInfo"}], "last_started_at": "2023-08-21T15:24:06.021Z", "last_finished_at": "2023-08-21T20:24:34.450Z", "initialized_options": [{"name": "Name", "version": 1}]}`)
				}))
			})
			It(`Invoke Initialize successfully`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := dpxService.Initialize(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ContainerReference model
				containerReferenceModel := new(dpxv1.ContainerReference)
				containerReferenceModel.ID = core.StringPtr("d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				containerReferenceModel.Type = core.StringPtr("catalog")

				// Construct an instance of the InitializeOptions model
				initializeOptionsModel := new(dpxv1.InitializeOptions)
				initializeOptionsModel.Container = containerReferenceModel
				initializeOptionsModel.Include = []string{"delivery_methods", "data_product_samples", "domains_multi_industry"}
				initializeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = dpxService.Initialize(initializeOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke Initialize with error: Operation request error`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Construct an instance of the ContainerReference model
				containerReferenceModel := new(dpxv1.ContainerReference)
				containerReferenceModel.ID = core.StringPtr("d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				containerReferenceModel.Type = core.StringPtr("catalog")

				// Construct an instance of the InitializeOptions model
				initializeOptionsModel := new(dpxv1.InitializeOptions)
				initializeOptionsModel.Container = containerReferenceModel
				initializeOptionsModel.Include = []string{"delivery_methods", "data_product_samples", "domains_multi_industry"}
				initializeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dpxService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := dpxService.Initialize(initializeOptionsModel)
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
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Construct an instance of the ContainerReference model
				containerReferenceModel := new(dpxv1.ContainerReference)
				containerReferenceModel.ID = core.StringPtr("d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				containerReferenceModel.Type = core.StringPtr("catalog")

				// Construct an instance of the InitializeOptions model
				initializeOptionsModel := new(dpxv1.InitializeOptions)
				initializeOptionsModel.Container = containerReferenceModel
				initializeOptionsModel.Include = []string{"delivery_methods", "data_product_samples", "domains_multi_industry"}
				initializeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := dpxService.Initialize(initializeOptionsModel)
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
	Describe(`ManageApiKeys(manageApiKeysOptions *ManageApiKeysOptions)`, func() {
		manageApiKeysPath := "/data_product_exchange/v1/configuration/rotate_credentials"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(manageApiKeysPath))
					Expect(req.Method).To(Equal("POST"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke ManageApiKeys successfully`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := dpxService.ManageApiKeys(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the ManageApiKeysOptions model
				manageApiKeysOptionsModel := new(dpxv1.ManageApiKeysOptions)
				manageApiKeysOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = dpxService.ManageApiKeys(manageApiKeysOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke ManageApiKeys with error: Operation request error`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Construct an instance of the ManageApiKeysOptions model
				manageApiKeysOptionsModel := new(dpxv1.ManageApiKeysOptions)
				manageApiKeysOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dpxService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := dpxService.ManageApiKeys(manageApiKeysOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
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
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Construct an instance of the ListDataProductsOptions model
				listDataProductsOptionsModel := new(dpxv1.ListDataProductsOptions)
				listDataProductsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listDataProductsOptionsModel.Start = core.StringPtr("testString")
				listDataProductsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := dpxService.ListDataProducts(listDataProductsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				dpxService.EnableRetries(0, 0)
				result, response, operationErr = dpxService.ListDataProducts(listDataProductsOptionsModel)
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
					fmt.Fprintf(res, "%s", `{"limit": 200, "first": {"href": "https://api.example.com/collection"}, "next": {"href": "https://api.example.com/collection?start=eyJvZmZzZXQiOjAsImRvbmUiOnRydWV9", "start": "eyJvZmZzZXQiOjAsImRvbmUiOnRydWV9"}, "data_products": [{"id": "b38df608-d34b-4d58-8136-ed25e6c6684e", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}]}`)
				}))
			})
			It(`Invoke ListDataProducts successfully with retries`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())
				dpxService.EnableRetries(0, 0)

				// Construct an instance of the ListDataProductsOptions model
				listDataProductsOptionsModel := new(dpxv1.ListDataProductsOptions)
				listDataProductsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listDataProductsOptionsModel.Start = core.StringPtr("testString")
				listDataProductsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := dpxService.ListDataProductsWithContext(ctx, listDataProductsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				dpxService.DisableRetries()
				result, response, operationErr := dpxService.ListDataProducts(listDataProductsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = dpxService.ListDataProductsWithContext(ctx, listDataProductsOptionsModel)
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
					fmt.Fprintf(res, "%s", `{"limit": 200, "first": {"href": "https://api.example.com/collection"}, "next": {"href": "https://api.example.com/collection?start=eyJvZmZzZXQiOjAsImRvbmUiOnRydWV9", "start": "eyJvZmZzZXQiOjAsImRvbmUiOnRydWV9"}, "data_products": [{"id": "b38df608-d34b-4d58-8136-ed25e6c6684e", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}]}`)
				}))
			})
			It(`Invoke ListDataProducts successfully`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := dpxService.ListDataProducts(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListDataProductsOptions model
				listDataProductsOptionsModel := new(dpxv1.ListDataProductsOptions)
				listDataProductsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listDataProductsOptionsModel.Start = core.StringPtr("testString")
				listDataProductsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = dpxService.ListDataProducts(listDataProductsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListDataProducts with error: Operation request error`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Construct an instance of the ListDataProductsOptions model
				listDataProductsOptionsModel := new(dpxv1.ListDataProductsOptions)
				listDataProductsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listDataProductsOptionsModel.Start = core.StringPtr("testString")
				listDataProductsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dpxService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := dpxService.ListDataProducts(listDataProductsOptionsModel)
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
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Construct an instance of the ListDataProductsOptions model
				listDataProductsOptionsModel := new(dpxv1.ListDataProductsOptions)
				listDataProductsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listDataProductsOptionsModel.Start = core.StringPtr("testString")
				listDataProductsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := dpxService.ListDataProducts(listDataProductsOptionsModel)
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
				responseObject := new(dpxv1.DataProductSummaryCollection)
				nextObject := new(dpxv1.NextPage)
				nextObject.Start = core.StringPtr("abc-123")
				responseObject.Next = nextObject

				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(Equal(core.StringPtr("abc-123")))
			})
			It(`Invoke GetNextStart without a "Next" property in the response`, func() {
				responseObject := new(dpxv1.DataProductSummaryCollection)

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
						fmt.Fprintf(res, "%s", `{"next":{"start":"1"},"total_count":2,"limit":1,"data_products":[{"id":"b38df608-d34b-4d58-8136-ed25e6c6684e","container":{"id":"d29c42eb-7100-4b7a-8257-c196dbcca1cd","type":"catalog"}}]}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"total_count":2,"limit":1,"data_products":[{"id":"b38df608-d34b-4d58-8136-ed25e6c6684e","container":{"id":"d29c42eb-7100-4b7a-8257-c196dbcca1cd","type":"catalog"}}]}`)
					} else {
						res.WriteHeader(400)
					}
				}))
			})
			It(`Use DataProductsPager.GetNext successfully`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				listDataProductsOptionsModel := &dpxv1.ListDataProductsOptions{
					Limit: core.Int64Ptr(int64(10)),
				}

				pager, err := dpxService.NewDataProductsPager(listDataProductsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				var allResults []dpxv1.DataProductSummary
				for pager.HasNext() {
					nextPage, err := pager.GetNext()
					Expect(err).To(BeNil())
					Expect(nextPage).ToNot(BeNil())
					allResults = append(allResults, nextPage...)
				}
				Expect(len(allResults)).To(Equal(2))
			})
			It(`Use DataProductsPager.GetAll successfully`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				listDataProductsOptionsModel := &dpxv1.ListDataProductsOptions{
					Limit: core.Int64Ptr(int64(10)),
				}

				pager, err := dpxService.NewDataProductsPager(listDataProductsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				allResults, err := pager.GetAll()
				Expect(err).To(BeNil())
				Expect(allResults).ToNot(BeNil())
				Expect(len(allResults)).To(Equal(2))
			})
		})
	})
	Describe(`CreateDataProduct(createDataProductOptions *CreateDataProductOptions) - Operation response error`, func() {
		createDataProductPath := "/data_product_exchange/v1/data_products"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createDataProductPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateDataProduct with error: Operation response processing error`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Construct an instance of the DataProductIdentity model
				dataProductIdentityModel := new(dpxv1.DataProductIdentity)
				dataProductIdentityModel.ID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")

				// Construct an instance of the ContainerReference model
				containerReferenceModel := new(dpxv1.ContainerReference)
				containerReferenceModel.ID = core.StringPtr("d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				containerReferenceModel.Type = core.StringPtr("catalog")

				// Construct an instance of the AssetReference model
				assetReferenceModel := new(dpxv1.AssetReference)
				assetReferenceModel.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				assetReferenceModel.Container = containerReferenceModel

				// Construct an instance of the UseCase model
				useCaseModel := new(dpxv1.UseCase)
				useCaseModel.ID = core.StringPtr("testString")
				useCaseModel.Name = core.StringPtr("testString")
				useCaseModel.Container = containerReferenceModel

				// Construct an instance of the Domain model
				domainModel := new(dpxv1.Domain)
				domainModel.ID = core.StringPtr("testString")
				domainModel.Name = core.StringPtr("testString")
				domainModel.Container = containerReferenceModel

				// Construct an instance of the AssetPartReference model
				assetPartReferenceModel := new(dpxv1.AssetPartReference)
				assetPartReferenceModel.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				assetPartReferenceModel.Container = containerReferenceModel
				assetPartReferenceModel.Type = core.StringPtr("data_asset")

				// Construct an instance of the DeliveryMethod model
				deliveryMethodModel := new(dpxv1.DeliveryMethod)
				deliveryMethodModel.ID = core.StringPtr("09cf5fcc-cb9d-4995-a8e4-16517b25229f")
				deliveryMethodModel.Container = containerReferenceModel

				// Construct an instance of the DataProductPart model
				dataProductPartModel := new(dpxv1.DataProductPart)
				dataProductPartModel.Asset = assetPartReferenceModel
				dataProductPartModel.Revision = core.Int64Ptr(int64(1))
				dataProductPartModel.UpdatedAt = CreateMockDateTime("2023-07-01T22:22:34.876Z")
				dataProductPartModel.DeliveryMethods = []dpxv1.DeliveryMethod{*deliveryMethodModel}

				// Construct an instance of the ContractTermsDocumentAttachment model
				contractTermsDocumentAttachmentModel := new(dpxv1.ContractTermsDocumentAttachment)
				contractTermsDocumentAttachmentModel.ID = core.StringPtr("testString")

				// Construct an instance of the ContractTermsDocument model
				contractTermsDocumentModel := new(dpxv1.ContractTermsDocument)
				contractTermsDocumentModel.URL = core.StringPtr("testString")
				contractTermsDocumentModel.Type = core.StringPtr("terms_and_conditions")
				contractTermsDocumentModel.Name = core.StringPtr("testString")
				contractTermsDocumentModel.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				contractTermsDocumentModel.Attachment = contractTermsDocumentAttachmentModel
				contractTermsDocumentModel.UploadURL = core.StringPtr("testString")

				// Construct an instance of the DataProductContractTerms model
				dataProductContractTermsModel := new(dpxv1.DataProductContractTerms)
				dataProductContractTermsModel.Asset = assetReferenceModel
				dataProductContractTermsModel.ID = core.StringPtr("testString")
				dataProductContractTermsModel.Documents = []dpxv1.ContractTermsDocument{*contractTermsDocumentModel}

				// Construct an instance of the DataProductVersionPrototype model
				dataProductVersionPrototypeModel := new(dpxv1.DataProductVersionPrototype)
				dataProductVersionPrototypeModel.Version = core.StringPtr("1.0.0")
				dataProductVersionPrototypeModel.State = core.StringPtr("draft")
				dataProductVersionPrototypeModel.DataProduct = dataProductIdentityModel
				dataProductVersionPrototypeModel.Name = core.StringPtr("My New Data Product")
				dataProductVersionPrototypeModel.Description = core.StringPtr("This is a description of My Data Product.")
				dataProductVersionPrototypeModel.Asset = assetReferenceModel
				dataProductVersionPrototypeModel.Tags = []string{"testString"}
				dataProductVersionPrototypeModel.UseCases = []dpxv1.UseCase{*useCaseModel}
				dataProductVersionPrototypeModel.Domain = domainModel
				dataProductVersionPrototypeModel.Types = []string{"data"}
				dataProductVersionPrototypeModel.PartsOut = []dpxv1.DataProductPart{*dataProductPartModel}
				dataProductVersionPrototypeModel.ContractTerms = []dpxv1.DataProductContractTerms{*dataProductContractTermsModel}
				dataProductVersionPrototypeModel.IsRestricted = core.BoolPtr(true)

				// Construct an instance of the CreateDataProductOptions model
				createDataProductOptionsModel := new(dpxv1.CreateDataProductOptions)
				createDataProductOptionsModel.Drafts = []dpxv1.DataProductVersionPrototype{*dataProductVersionPrototypeModel}
				createDataProductOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := dpxService.CreateDataProduct(createDataProductOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				dpxService.EnableRetries(0, 0)
				result, response, operationErr = dpxService.CreateDataProduct(createDataProductOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateDataProduct(createDataProductOptions *CreateDataProductOptions)`, func() {
		createDataProductPath := "/data_product_exchange/v1/data_products"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createDataProductPath))
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
					fmt.Fprintf(res, "%s", `{"id": "b38df608-d34b-4d58-8136-ed25e6c6684e", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}, "latest_release": {"version": "1.0.0", "state": "draft", "data_product": {"id": "b38df608-d34b-4d58-8136-ed25e6c6684e"}, "name": "My Data Product", "description": "This is a description of My Data Product.", "id": "2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd", "asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}}, "drafts": [{"version": "1.0.0", "state": "draft", "data_product": {"id": "b38df608-d34b-4d58-8136-ed25e6c6684e"}, "name": "My Data Product", "description": "This is a description of My Data Product.", "id": "2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd", "asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}}]}`)
				}))
			})
			It(`Invoke CreateDataProduct successfully with retries`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())
				dpxService.EnableRetries(0, 0)

				// Construct an instance of the DataProductIdentity model
				dataProductIdentityModel := new(dpxv1.DataProductIdentity)
				dataProductIdentityModel.ID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")

				// Construct an instance of the ContainerReference model
				containerReferenceModel := new(dpxv1.ContainerReference)
				containerReferenceModel.ID = core.StringPtr("d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				containerReferenceModel.Type = core.StringPtr("catalog")

				// Construct an instance of the AssetReference model
				assetReferenceModel := new(dpxv1.AssetReference)
				assetReferenceModel.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				assetReferenceModel.Container = containerReferenceModel

				// Construct an instance of the UseCase model
				useCaseModel := new(dpxv1.UseCase)
				useCaseModel.ID = core.StringPtr("testString")
				useCaseModel.Name = core.StringPtr("testString")
				useCaseModel.Container = containerReferenceModel

				// Construct an instance of the Domain model
				domainModel := new(dpxv1.Domain)
				domainModel.ID = core.StringPtr("testString")
				domainModel.Name = core.StringPtr("testString")
				domainModel.Container = containerReferenceModel

				// Construct an instance of the AssetPartReference model
				assetPartReferenceModel := new(dpxv1.AssetPartReference)
				assetPartReferenceModel.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				assetPartReferenceModel.Container = containerReferenceModel
				assetPartReferenceModel.Type = core.StringPtr("data_asset")

				// Construct an instance of the DeliveryMethod model
				deliveryMethodModel := new(dpxv1.DeliveryMethod)
				deliveryMethodModel.ID = core.StringPtr("09cf5fcc-cb9d-4995-a8e4-16517b25229f")
				deliveryMethodModel.Container = containerReferenceModel

				// Construct an instance of the DataProductPart model
				dataProductPartModel := new(dpxv1.DataProductPart)
				dataProductPartModel.Asset = assetPartReferenceModel
				dataProductPartModel.Revision = core.Int64Ptr(int64(1))
				dataProductPartModel.UpdatedAt = CreateMockDateTime("2023-07-01T22:22:34.876Z")
				dataProductPartModel.DeliveryMethods = []dpxv1.DeliveryMethod{*deliveryMethodModel}

				// Construct an instance of the ContractTermsDocumentAttachment model
				contractTermsDocumentAttachmentModel := new(dpxv1.ContractTermsDocumentAttachment)
				contractTermsDocumentAttachmentModel.ID = core.StringPtr("testString")

				// Construct an instance of the ContractTermsDocument model
				contractTermsDocumentModel := new(dpxv1.ContractTermsDocument)
				contractTermsDocumentModel.URL = core.StringPtr("testString")
				contractTermsDocumentModel.Type = core.StringPtr("terms_and_conditions")
				contractTermsDocumentModel.Name = core.StringPtr("testString")
				contractTermsDocumentModel.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				contractTermsDocumentModel.Attachment = contractTermsDocumentAttachmentModel
				contractTermsDocumentModel.UploadURL = core.StringPtr("testString")

				// Construct an instance of the DataProductContractTerms model
				dataProductContractTermsModel := new(dpxv1.DataProductContractTerms)
				dataProductContractTermsModel.Asset = assetReferenceModel
				dataProductContractTermsModel.ID = core.StringPtr("testString")
				dataProductContractTermsModel.Documents = []dpxv1.ContractTermsDocument{*contractTermsDocumentModel}

				// Construct an instance of the DataProductVersionPrototype model
				dataProductVersionPrototypeModel := new(dpxv1.DataProductVersionPrototype)
				dataProductVersionPrototypeModel.Version = core.StringPtr("1.0.0")
				dataProductVersionPrototypeModel.State = core.StringPtr("draft")
				dataProductVersionPrototypeModel.DataProduct = dataProductIdentityModel
				dataProductVersionPrototypeModel.Name = core.StringPtr("My New Data Product")
				dataProductVersionPrototypeModel.Description = core.StringPtr("This is a description of My Data Product.")
				dataProductVersionPrototypeModel.Asset = assetReferenceModel
				dataProductVersionPrototypeModel.Tags = []string{"testString"}
				dataProductVersionPrototypeModel.UseCases = []dpxv1.UseCase{*useCaseModel}
				dataProductVersionPrototypeModel.Domain = domainModel
				dataProductVersionPrototypeModel.Types = []string{"data"}
				dataProductVersionPrototypeModel.PartsOut = []dpxv1.DataProductPart{*dataProductPartModel}
				dataProductVersionPrototypeModel.ContractTerms = []dpxv1.DataProductContractTerms{*dataProductContractTermsModel}
				dataProductVersionPrototypeModel.IsRestricted = core.BoolPtr(true)

				// Construct an instance of the CreateDataProductOptions model
				createDataProductOptionsModel := new(dpxv1.CreateDataProductOptions)
				createDataProductOptionsModel.Drafts = []dpxv1.DataProductVersionPrototype{*dataProductVersionPrototypeModel}
				createDataProductOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := dpxService.CreateDataProductWithContext(ctx, createDataProductOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				dpxService.DisableRetries()
				result, response, operationErr := dpxService.CreateDataProduct(createDataProductOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = dpxService.CreateDataProductWithContext(ctx, createDataProductOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createDataProductPath))
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
					fmt.Fprintf(res, "%s", `{"id": "b38df608-d34b-4d58-8136-ed25e6c6684e", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}, "latest_release": {"version": "1.0.0", "state": "draft", "data_product": {"id": "b38df608-d34b-4d58-8136-ed25e6c6684e"}, "name": "My Data Product", "description": "This is a description of My Data Product.", "id": "2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd", "asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}}, "drafts": [{"version": "1.0.0", "state": "draft", "data_product": {"id": "b38df608-d34b-4d58-8136-ed25e6c6684e"}, "name": "My Data Product", "description": "This is a description of My Data Product.", "id": "2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd", "asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}}]}`)
				}))
			})
			It(`Invoke CreateDataProduct successfully`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := dpxService.CreateDataProduct(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DataProductIdentity model
				dataProductIdentityModel := new(dpxv1.DataProductIdentity)
				dataProductIdentityModel.ID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")

				// Construct an instance of the ContainerReference model
				containerReferenceModel := new(dpxv1.ContainerReference)
				containerReferenceModel.ID = core.StringPtr("d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				containerReferenceModel.Type = core.StringPtr("catalog")

				// Construct an instance of the AssetReference model
				assetReferenceModel := new(dpxv1.AssetReference)
				assetReferenceModel.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				assetReferenceModel.Container = containerReferenceModel

				// Construct an instance of the UseCase model
				useCaseModel := new(dpxv1.UseCase)
				useCaseModel.ID = core.StringPtr("testString")
				useCaseModel.Name = core.StringPtr("testString")
				useCaseModel.Container = containerReferenceModel

				// Construct an instance of the Domain model
				domainModel := new(dpxv1.Domain)
				domainModel.ID = core.StringPtr("testString")
				domainModel.Name = core.StringPtr("testString")
				domainModel.Container = containerReferenceModel

				// Construct an instance of the AssetPartReference model
				assetPartReferenceModel := new(dpxv1.AssetPartReference)
				assetPartReferenceModel.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				assetPartReferenceModel.Container = containerReferenceModel
				assetPartReferenceModel.Type = core.StringPtr("data_asset")

				// Construct an instance of the DeliveryMethod model
				deliveryMethodModel := new(dpxv1.DeliveryMethod)
				deliveryMethodModel.ID = core.StringPtr("09cf5fcc-cb9d-4995-a8e4-16517b25229f")
				deliveryMethodModel.Container = containerReferenceModel

				// Construct an instance of the DataProductPart model
				dataProductPartModel := new(dpxv1.DataProductPart)
				dataProductPartModel.Asset = assetPartReferenceModel
				dataProductPartModel.Revision = core.Int64Ptr(int64(1))
				dataProductPartModel.UpdatedAt = CreateMockDateTime("2023-07-01T22:22:34.876Z")
				dataProductPartModel.DeliveryMethods = []dpxv1.DeliveryMethod{*deliveryMethodModel}

				// Construct an instance of the ContractTermsDocumentAttachment model
				contractTermsDocumentAttachmentModel := new(dpxv1.ContractTermsDocumentAttachment)
				contractTermsDocumentAttachmentModel.ID = core.StringPtr("testString")

				// Construct an instance of the ContractTermsDocument model
				contractTermsDocumentModel := new(dpxv1.ContractTermsDocument)
				contractTermsDocumentModel.URL = core.StringPtr("testString")
				contractTermsDocumentModel.Type = core.StringPtr("terms_and_conditions")
				contractTermsDocumentModel.Name = core.StringPtr("testString")
				contractTermsDocumentModel.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				contractTermsDocumentModel.Attachment = contractTermsDocumentAttachmentModel
				contractTermsDocumentModel.UploadURL = core.StringPtr("testString")

				// Construct an instance of the DataProductContractTerms model
				dataProductContractTermsModel := new(dpxv1.DataProductContractTerms)
				dataProductContractTermsModel.Asset = assetReferenceModel
				dataProductContractTermsModel.ID = core.StringPtr("testString")
				dataProductContractTermsModel.Documents = []dpxv1.ContractTermsDocument{*contractTermsDocumentModel}

				// Construct an instance of the DataProductVersionPrototype model
				dataProductVersionPrototypeModel := new(dpxv1.DataProductVersionPrototype)
				dataProductVersionPrototypeModel.Version = core.StringPtr("1.0.0")
				dataProductVersionPrototypeModel.State = core.StringPtr("draft")
				dataProductVersionPrototypeModel.DataProduct = dataProductIdentityModel
				dataProductVersionPrototypeModel.Name = core.StringPtr("My New Data Product")
				dataProductVersionPrototypeModel.Description = core.StringPtr("This is a description of My Data Product.")
				dataProductVersionPrototypeModel.Asset = assetReferenceModel
				dataProductVersionPrototypeModel.Tags = []string{"testString"}
				dataProductVersionPrototypeModel.UseCases = []dpxv1.UseCase{*useCaseModel}
				dataProductVersionPrototypeModel.Domain = domainModel
				dataProductVersionPrototypeModel.Types = []string{"data"}
				dataProductVersionPrototypeModel.PartsOut = []dpxv1.DataProductPart{*dataProductPartModel}
				dataProductVersionPrototypeModel.ContractTerms = []dpxv1.DataProductContractTerms{*dataProductContractTermsModel}
				dataProductVersionPrototypeModel.IsRestricted = core.BoolPtr(true)

				// Construct an instance of the CreateDataProductOptions model
				createDataProductOptionsModel := new(dpxv1.CreateDataProductOptions)
				createDataProductOptionsModel.Drafts = []dpxv1.DataProductVersionPrototype{*dataProductVersionPrototypeModel}
				createDataProductOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = dpxService.CreateDataProduct(createDataProductOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateDataProduct with error: Operation validation and request error`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Construct an instance of the DataProductIdentity model
				dataProductIdentityModel := new(dpxv1.DataProductIdentity)
				dataProductIdentityModel.ID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")

				// Construct an instance of the ContainerReference model
				containerReferenceModel := new(dpxv1.ContainerReference)
				containerReferenceModel.ID = core.StringPtr("d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				containerReferenceModel.Type = core.StringPtr("catalog")

				// Construct an instance of the AssetReference model
				assetReferenceModel := new(dpxv1.AssetReference)
				assetReferenceModel.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				assetReferenceModel.Container = containerReferenceModel

				// Construct an instance of the UseCase model
				useCaseModel := new(dpxv1.UseCase)
				useCaseModel.ID = core.StringPtr("testString")
				useCaseModel.Name = core.StringPtr("testString")
				useCaseModel.Container = containerReferenceModel

				// Construct an instance of the Domain model
				domainModel := new(dpxv1.Domain)
				domainModel.ID = core.StringPtr("testString")
				domainModel.Name = core.StringPtr("testString")
				domainModel.Container = containerReferenceModel

				// Construct an instance of the AssetPartReference model
				assetPartReferenceModel := new(dpxv1.AssetPartReference)
				assetPartReferenceModel.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				assetPartReferenceModel.Container = containerReferenceModel
				assetPartReferenceModel.Type = core.StringPtr("data_asset")

				// Construct an instance of the DeliveryMethod model
				deliveryMethodModel := new(dpxv1.DeliveryMethod)
				deliveryMethodModel.ID = core.StringPtr("09cf5fcc-cb9d-4995-a8e4-16517b25229f")
				deliveryMethodModel.Container = containerReferenceModel

				// Construct an instance of the DataProductPart model
				dataProductPartModel := new(dpxv1.DataProductPart)
				dataProductPartModel.Asset = assetPartReferenceModel
				dataProductPartModel.Revision = core.Int64Ptr(int64(1))
				dataProductPartModel.UpdatedAt = CreateMockDateTime("2023-07-01T22:22:34.876Z")
				dataProductPartModel.DeliveryMethods = []dpxv1.DeliveryMethod{*deliveryMethodModel}

				// Construct an instance of the ContractTermsDocumentAttachment model
				contractTermsDocumentAttachmentModel := new(dpxv1.ContractTermsDocumentAttachment)
				contractTermsDocumentAttachmentModel.ID = core.StringPtr("testString")

				// Construct an instance of the ContractTermsDocument model
				contractTermsDocumentModel := new(dpxv1.ContractTermsDocument)
				contractTermsDocumentModel.URL = core.StringPtr("testString")
				contractTermsDocumentModel.Type = core.StringPtr("terms_and_conditions")
				contractTermsDocumentModel.Name = core.StringPtr("testString")
				contractTermsDocumentModel.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				contractTermsDocumentModel.Attachment = contractTermsDocumentAttachmentModel
				contractTermsDocumentModel.UploadURL = core.StringPtr("testString")

				// Construct an instance of the DataProductContractTerms model
				dataProductContractTermsModel := new(dpxv1.DataProductContractTerms)
				dataProductContractTermsModel.Asset = assetReferenceModel
				dataProductContractTermsModel.ID = core.StringPtr("testString")
				dataProductContractTermsModel.Documents = []dpxv1.ContractTermsDocument{*contractTermsDocumentModel}

				// Construct an instance of the DataProductVersionPrototype model
				dataProductVersionPrototypeModel := new(dpxv1.DataProductVersionPrototype)
				dataProductVersionPrototypeModel.Version = core.StringPtr("1.0.0")
				dataProductVersionPrototypeModel.State = core.StringPtr("draft")
				dataProductVersionPrototypeModel.DataProduct = dataProductIdentityModel
				dataProductVersionPrototypeModel.Name = core.StringPtr("My New Data Product")
				dataProductVersionPrototypeModel.Description = core.StringPtr("This is a description of My Data Product.")
				dataProductVersionPrototypeModel.Asset = assetReferenceModel
				dataProductVersionPrototypeModel.Tags = []string{"testString"}
				dataProductVersionPrototypeModel.UseCases = []dpxv1.UseCase{*useCaseModel}
				dataProductVersionPrototypeModel.Domain = domainModel
				dataProductVersionPrototypeModel.Types = []string{"data"}
				dataProductVersionPrototypeModel.PartsOut = []dpxv1.DataProductPart{*dataProductPartModel}
				dataProductVersionPrototypeModel.ContractTerms = []dpxv1.DataProductContractTerms{*dataProductContractTermsModel}
				dataProductVersionPrototypeModel.IsRestricted = core.BoolPtr(true)

				// Construct an instance of the CreateDataProductOptions model
				createDataProductOptionsModel := new(dpxv1.CreateDataProductOptions)
				createDataProductOptionsModel.Drafts = []dpxv1.DataProductVersionPrototype{*dataProductVersionPrototypeModel}
				createDataProductOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dpxService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := dpxService.CreateDataProduct(createDataProductOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateDataProductOptions model with no property values
				createDataProductOptionsModelNew := new(dpxv1.CreateDataProductOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = dpxService.CreateDataProduct(createDataProductOptionsModelNew)
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
			It(`Invoke CreateDataProduct successfully`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Construct an instance of the DataProductIdentity model
				dataProductIdentityModel := new(dpxv1.DataProductIdentity)
				dataProductIdentityModel.ID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")

				// Construct an instance of the ContainerReference model
				containerReferenceModel := new(dpxv1.ContainerReference)
				containerReferenceModel.ID = core.StringPtr("d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				containerReferenceModel.Type = core.StringPtr("catalog")

				// Construct an instance of the AssetReference model
				assetReferenceModel := new(dpxv1.AssetReference)
				assetReferenceModel.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				assetReferenceModel.Container = containerReferenceModel

				// Construct an instance of the UseCase model
				useCaseModel := new(dpxv1.UseCase)
				useCaseModel.ID = core.StringPtr("testString")
				useCaseModel.Name = core.StringPtr("testString")
				useCaseModel.Container = containerReferenceModel

				// Construct an instance of the Domain model
				domainModel := new(dpxv1.Domain)
				domainModel.ID = core.StringPtr("testString")
				domainModel.Name = core.StringPtr("testString")
				domainModel.Container = containerReferenceModel

				// Construct an instance of the AssetPartReference model
				assetPartReferenceModel := new(dpxv1.AssetPartReference)
				assetPartReferenceModel.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				assetPartReferenceModel.Container = containerReferenceModel
				assetPartReferenceModel.Type = core.StringPtr("data_asset")

				// Construct an instance of the DeliveryMethod model
				deliveryMethodModel := new(dpxv1.DeliveryMethod)
				deliveryMethodModel.ID = core.StringPtr("09cf5fcc-cb9d-4995-a8e4-16517b25229f")
				deliveryMethodModel.Container = containerReferenceModel

				// Construct an instance of the DataProductPart model
				dataProductPartModel := new(dpxv1.DataProductPart)
				dataProductPartModel.Asset = assetPartReferenceModel
				dataProductPartModel.Revision = core.Int64Ptr(int64(1))
				dataProductPartModel.UpdatedAt = CreateMockDateTime("2023-07-01T22:22:34.876Z")
				dataProductPartModel.DeliveryMethods = []dpxv1.DeliveryMethod{*deliveryMethodModel}

				// Construct an instance of the ContractTermsDocumentAttachment model
				contractTermsDocumentAttachmentModel := new(dpxv1.ContractTermsDocumentAttachment)
				contractTermsDocumentAttachmentModel.ID = core.StringPtr("testString")

				// Construct an instance of the ContractTermsDocument model
				contractTermsDocumentModel := new(dpxv1.ContractTermsDocument)
				contractTermsDocumentModel.URL = core.StringPtr("testString")
				contractTermsDocumentModel.Type = core.StringPtr("terms_and_conditions")
				contractTermsDocumentModel.Name = core.StringPtr("testString")
				contractTermsDocumentModel.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				contractTermsDocumentModel.Attachment = contractTermsDocumentAttachmentModel
				contractTermsDocumentModel.UploadURL = core.StringPtr("testString")

				// Construct an instance of the DataProductContractTerms model
				dataProductContractTermsModel := new(dpxv1.DataProductContractTerms)
				dataProductContractTermsModel.Asset = assetReferenceModel
				dataProductContractTermsModel.ID = core.StringPtr("testString")
				dataProductContractTermsModel.Documents = []dpxv1.ContractTermsDocument{*contractTermsDocumentModel}

				// Construct an instance of the DataProductVersionPrototype model
				dataProductVersionPrototypeModel := new(dpxv1.DataProductVersionPrototype)
				dataProductVersionPrototypeModel.Version = core.StringPtr("1.0.0")
				dataProductVersionPrototypeModel.State = core.StringPtr("draft")
				dataProductVersionPrototypeModel.DataProduct = dataProductIdentityModel
				dataProductVersionPrototypeModel.Name = core.StringPtr("My New Data Product")
				dataProductVersionPrototypeModel.Description = core.StringPtr("This is a description of My Data Product.")
				dataProductVersionPrototypeModel.Asset = assetReferenceModel
				dataProductVersionPrototypeModel.Tags = []string{"testString"}
				dataProductVersionPrototypeModel.UseCases = []dpxv1.UseCase{*useCaseModel}
				dataProductVersionPrototypeModel.Domain = domainModel
				dataProductVersionPrototypeModel.Types = []string{"data"}
				dataProductVersionPrototypeModel.PartsOut = []dpxv1.DataProductPart{*dataProductPartModel}
				dataProductVersionPrototypeModel.ContractTerms = []dpxv1.DataProductContractTerms{*dataProductContractTermsModel}
				dataProductVersionPrototypeModel.IsRestricted = core.BoolPtr(true)

				// Construct an instance of the CreateDataProductOptions model
				createDataProductOptionsModel := new(dpxv1.CreateDataProductOptions)
				createDataProductOptionsModel.Drafts = []dpxv1.DataProductVersionPrototype{*dataProductVersionPrototypeModel}
				createDataProductOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := dpxService.CreateDataProduct(createDataProductOptionsModel)
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
		getDataProductPath := "/data_product_exchange/v1/data_products/b38df608-d34b-4d58-8136-ed25e6c6684e"
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
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Construct an instance of the GetDataProductOptions model
				getDataProductOptionsModel := new(dpxv1.GetDataProductOptions)
				getDataProductOptionsModel.DataProductID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")
				getDataProductOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := dpxService.GetDataProduct(getDataProductOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				dpxService.EnableRetries(0, 0)
				result, response, operationErr = dpxService.GetDataProduct(getDataProductOptionsModel)
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
		getDataProductPath := "/data_product_exchange/v1/data_products/b38df608-d34b-4d58-8136-ed25e6c6684e"
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
					fmt.Fprintf(res, "%s", `{"id": "b38df608-d34b-4d58-8136-ed25e6c6684e", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}, "latest_release": {"version": "1.0.0", "state": "draft", "data_product": {"id": "b38df608-d34b-4d58-8136-ed25e6c6684e"}, "name": "My Data Product", "description": "This is a description of My Data Product.", "id": "2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd", "asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}}, "drafts": [{"version": "1.0.0", "state": "draft", "data_product": {"id": "b38df608-d34b-4d58-8136-ed25e6c6684e"}, "name": "My Data Product", "description": "This is a description of My Data Product.", "id": "2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd", "asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}}]}`)
				}))
			})
			It(`Invoke GetDataProduct successfully with retries`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())
				dpxService.EnableRetries(0, 0)

				// Construct an instance of the GetDataProductOptions model
				getDataProductOptionsModel := new(dpxv1.GetDataProductOptions)
				getDataProductOptionsModel.DataProductID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")
				getDataProductOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := dpxService.GetDataProductWithContext(ctx, getDataProductOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				dpxService.DisableRetries()
				result, response, operationErr := dpxService.GetDataProduct(getDataProductOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = dpxService.GetDataProductWithContext(ctx, getDataProductOptionsModel)
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
					fmt.Fprintf(res, "%s", `{"id": "b38df608-d34b-4d58-8136-ed25e6c6684e", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}, "latest_release": {"version": "1.0.0", "state": "draft", "data_product": {"id": "b38df608-d34b-4d58-8136-ed25e6c6684e"}, "name": "My Data Product", "description": "This is a description of My Data Product.", "id": "2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd", "asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}}, "drafts": [{"version": "1.0.0", "state": "draft", "data_product": {"id": "b38df608-d34b-4d58-8136-ed25e6c6684e"}, "name": "My Data Product", "description": "This is a description of My Data Product.", "id": "2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd", "asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}}]}`)
				}))
			})
			It(`Invoke GetDataProduct successfully`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := dpxService.GetDataProduct(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetDataProductOptions model
				getDataProductOptionsModel := new(dpxv1.GetDataProductOptions)
				getDataProductOptionsModel.DataProductID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")
				getDataProductOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = dpxService.GetDataProduct(getDataProductOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetDataProduct with error: Operation validation and request error`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Construct an instance of the GetDataProductOptions model
				getDataProductOptionsModel := new(dpxv1.GetDataProductOptions)
				getDataProductOptionsModel.DataProductID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")
				getDataProductOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dpxService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := dpxService.GetDataProduct(getDataProductOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetDataProductOptions model with no property values
				getDataProductOptionsModelNew := new(dpxv1.GetDataProductOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = dpxService.GetDataProduct(getDataProductOptionsModelNew)
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
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Construct an instance of the GetDataProductOptions model
				getDataProductOptionsModel := new(dpxv1.GetDataProductOptions)
				getDataProductOptionsModel.DataProductID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")
				getDataProductOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := dpxService.GetDataProduct(getDataProductOptionsModel)
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
	Describe(`CompleteDraftContractTermsDocument(completeDraftContractTermsDocumentOptions *CompleteDraftContractTermsDocumentOptions) - Operation response error`, func() {
		completeDraftContractTermsDocumentPath := "/data_product_exchange/v1/data_products/b38df608-d34b-4d58-8136-ed25e6c6684e/drafts/2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd/contract_terms/598183cd-b910-4e8d-9a97-97097afda3c1@e4fe2f87-0e56-46dd-b3b8-e9af32309e82/documents/testString/complete"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(completeDraftContractTermsDocumentPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CompleteDraftContractTermsDocument with error: Operation response processing error`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Construct an instance of the CompleteDraftContractTermsDocumentOptions model
				completeDraftContractTermsDocumentOptionsModel := new(dpxv1.CompleteDraftContractTermsDocumentOptions)
				completeDraftContractTermsDocumentOptionsModel.DataProductID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")
				completeDraftContractTermsDocumentOptionsModel.DraftID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				completeDraftContractTermsDocumentOptionsModel.ContractTermsID = core.StringPtr("598183cd-b910-4e8d-9a97-97097afda3c1@e4fe2f87-0e56-46dd-b3b8-e9af32309e82")
				completeDraftContractTermsDocumentOptionsModel.DocumentID = core.StringPtr("testString")
				completeDraftContractTermsDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := dpxService.CompleteDraftContractTermsDocument(completeDraftContractTermsDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				dpxService.EnableRetries(0, 0)
				result, response, operationErr = dpxService.CompleteDraftContractTermsDocument(completeDraftContractTermsDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CompleteDraftContractTermsDocument(completeDraftContractTermsDocumentOptions *CompleteDraftContractTermsDocumentOptions)`, func() {
		completeDraftContractTermsDocumentPath := "/data_product_exchange/v1/data_products/b38df608-d34b-4d58-8136-ed25e6c6684e/drafts/2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd/contract_terms/598183cd-b910-4e8d-9a97-97097afda3c1@e4fe2f87-0e56-46dd-b3b8-e9af32309e82/documents/testString/complete"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(completeDraftContractTermsDocumentPath))
					Expect(req.Method).To(Equal("POST"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"url": "URL", "type": "terms_and_conditions", "name": "Name", "id": "2b0bf220-079c-11ee-be56-0242ac120002", "attachment": {"id": "ID"}, "upload_url": "UploadURL"}`)
				}))
			})
			It(`Invoke CompleteDraftContractTermsDocument successfully with retries`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())
				dpxService.EnableRetries(0, 0)

				// Construct an instance of the CompleteDraftContractTermsDocumentOptions model
				completeDraftContractTermsDocumentOptionsModel := new(dpxv1.CompleteDraftContractTermsDocumentOptions)
				completeDraftContractTermsDocumentOptionsModel.DataProductID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")
				completeDraftContractTermsDocumentOptionsModel.DraftID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				completeDraftContractTermsDocumentOptionsModel.ContractTermsID = core.StringPtr("598183cd-b910-4e8d-9a97-97097afda3c1@e4fe2f87-0e56-46dd-b3b8-e9af32309e82")
				completeDraftContractTermsDocumentOptionsModel.DocumentID = core.StringPtr("testString")
				completeDraftContractTermsDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := dpxService.CompleteDraftContractTermsDocumentWithContext(ctx, completeDraftContractTermsDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				dpxService.DisableRetries()
				result, response, operationErr := dpxService.CompleteDraftContractTermsDocument(completeDraftContractTermsDocumentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = dpxService.CompleteDraftContractTermsDocumentWithContext(ctx, completeDraftContractTermsDocumentOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(completeDraftContractTermsDocumentPath))
					Expect(req.Method).To(Equal("POST"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"url": "URL", "type": "terms_and_conditions", "name": "Name", "id": "2b0bf220-079c-11ee-be56-0242ac120002", "attachment": {"id": "ID"}, "upload_url": "UploadURL"}`)
				}))
			})
			It(`Invoke CompleteDraftContractTermsDocument successfully`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := dpxService.CompleteDraftContractTermsDocument(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CompleteDraftContractTermsDocumentOptions model
				completeDraftContractTermsDocumentOptionsModel := new(dpxv1.CompleteDraftContractTermsDocumentOptions)
				completeDraftContractTermsDocumentOptionsModel.DataProductID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")
				completeDraftContractTermsDocumentOptionsModel.DraftID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				completeDraftContractTermsDocumentOptionsModel.ContractTermsID = core.StringPtr("598183cd-b910-4e8d-9a97-97097afda3c1@e4fe2f87-0e56-46dd-b3b8-e9af32309e82")
				completeDraftContractTermsDocumentOptionsModel.DocumentID = core.StringPtr("testString")
				completeDraftContractTermsDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = dpxService.CompleteDraftContractTermsDocument(completeDraftContractTermsDocumentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CompleteDraftContractTermsDocument with error: Operation validation and request error`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Construct an instance of the CompleteDraftContractTermsDocumentOptions model
				completeDraftContractTermsDocumentOptionsModel := new(dpxv1.CompleteDraftContractTermsDocumentOptions)
				completeDraftContractTermsDocumentOptionsModel.DataProductID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")
				completeDraftContractTermsDocumentOptionsModel.DraftID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				completeDraftContractTermsDocumentOptionsModel.ContractTermsID = core.StringPtr("598183cd-b910-4e8d-9a97-97097afda3c1@e4fe2f87-0e56-46dd-b3b8-e9af32309e82")
				completeDraftContractTermsDocumentOptionsModel.DocumentID = core.StringPtr("testString")
				completeDraftContractTermsDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dpxService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := dpxService.CompleteDraftContractTermsDocument(completeDraftContractTermsDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CompleteDraftContractTermsDocumentOptions model with no property values
				completeDraftContractTermsDocumentOptionsModelNew := new(dpxv1.CompleteDraftContractTermsDocumentOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = dpxService.CompleteDraftContractTermsDocument(completeDraftContractTermsDocumentOptionsModelNew)
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
			It(`Invoke CompleteDraftContractTermsDocument successfully`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Construct an instance of the CompleteDraftContractTermsDocumentOptions model
				completeDraftContractTermsDocumentOptionsModel := new(dpxv1.CompleteDraftContractTermsDocumentOptions)
				completeDraftContractTermsDocumentOptionsModel.DataProductID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")
				completeDraftContractTermsDocumentOptionsModel.DraftID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				completeDraftContractTermsDocumentOptionsModel.ContractTermsID = core.StringPtr("598183cd-b910-4e8d-9a97-97097afda3c1@e4fe2f87-0e56-46dd-b3b8-e9af32309e82")
				completeDraftContractTermsDocumentOptionsModel.DocumentID = core.StringPtr("testString")
				completeDraftContractTermsDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := dpxService.CompleteDraftContractTermsDocument(completeDraftContractTermsDocumentOptionsModel)
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
	Describe(`ListDataProductDrafts(listDataProductDraftsOptions *ListDataProductDraftsOptions) - Operation response error`, func() {
		listDataProductDraftsPath := "/data_product_exchange/v1/data_products/b38df608-d34b-4d58-8136-ed25e6c6684e/drafts"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listDataProductDraftsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["asset.container.id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListDataProductDrafts with error: Operation response processing error`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Construct an instance of the ListDataProductDraftsOptions model
				listDataProductDraftsOptionsModel := new(dpxv1.ListDataProductDraftsOptions)
				listDataProductDraftsOptionsModel.DataProductID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")
				listDataProductDraftsOptionsModel.AssetContainerID = core.StringPtr("testString")
				listDataProductDraftsOptionsModel.Version = core.StringPtr("testString")
				listDataProductDraftsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listDataProductDraftsOptionsModel.Start = core.StringPtr("testString")
				listDataProductDraftsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := dpxService.ListDataProductDrafts(listDataProductDraftsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				dpxService.EnableRetries(0, 0)
				result, response, operationErr = dpxService.ListDataProductDrafts(listDataProductDraftsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListDataProductDrafts(listDataProductDraftsOptions *ListDataProductDraftsOptions)`, func() {
		listDataProductDraftsPath := "/data_product_exchange/v1/data_products/b38df608-d34b-4d58-8136-ed25e6c6684e/drafts"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listDataProductDraftsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["asset.container.id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"limit": 200, "first": {"href": "https://api.example.com/collection"}, "next": {"href": "https://api.example.com/collection?start=eyJvZmZzZXQiOjAsImRvbmUiOnRydWV9", "start": "eyJvZmZzZXQiOjAsImRvbmUiOnRydWV9"}, "drafts": [{"version": "1.0.0", "state": "draft", "data_product": {"id": "b38df608-d34b-4d58-8136-ed25e6c6684e"}, "name": "My Data Product", "description": "This is a description of My Data Product.", "id": "2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd", "asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}}]}`)
				}))
			})
			It(`Invoke ListDataProductDrafts successfully with retries`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())
				dpxService.EnableRetries(0, 0)

				// Construct an instance of the ListDataProductDraftsOptions model
				listDataProductDraftsOptionsModel := new(dpxv1.ListDataProductDraftsOptions)
				listDataProductDraftsOptionsModel.DataProductID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")
				listDataProductDraftsOptionsModel.AssetContainerID = core.StringPtr("testString")
				listDataProductDraftsOptionsModel.Version = core.StringPtr("testString")
				listDataProductDraftsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listDataProductDraftsOptionsModel.Start = core.StringPtr("testString")
				listDataProductDraftsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := dpxService.ListDataProductDraftsWithContext(ctx, listDataProductDraftsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				dpxService.DisableRetries()
				result, response, operationErr := dpxService.ListDataProductDrafts(listDataProductDraftsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = dpxService.ListDataProductDraftsWithContext(ctx, listDataProductDraftsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listDataProductDraftsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["asset.container.id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"limit": 200, "first": {"href": "https://api.example.com/collection"}, "next": {"href": "https://api.example.com/collection?start=eyJvZmZzZXQiOjAsImRvbmUiOnRydWV9", "start": "eyJvZmZzZXQiOjAsImRvbmUiOnRydWV9"}, "drafts": [{"version": "1.0.0", "state": "draft", "data_product": {"id": "b38df608-d34b-4d58-8136-ed25e6c6684e"}, "name": "My Data Product", "description": "This is a description of My Data Product.", "id": "2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd", "asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}}]}`)
				}))
			})
			It(`Invoke ListDataProductDrafts successfully`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := dpxService.ListDataProductDrafts(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListDataProductDraftsOptions model
				listDataProductDraftsOptionsModel := new(dpxv1.ListDataProductDraftsOptions)
				listDataProductDraftsOptionsModel.DataProductID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")
				listDataProductDraftsOptionsModel.AssetContainerID = core.StringPtr("testString")
				listDataProductDraftsOptionsModel.Version = core.StringPtr("testString")
				listDataProductDraftsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listDataProductDraftsOptionsModel.Start = core.StringPtr("testString")
				listDataProductDraftsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = dpxService.ListDataProductDrafts(listDataProductDraftsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListDataProductDrafts with error: Operation validation and request error`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Construct an instance of the ListDataProductDraftsOptions model
				listDataProductDraftsOptionsModel := new(dpxv1.ListDataProductDraftsOptions)
				listDataProductDraftsOptionsModel.DataProductID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")
				listDataProductDraftsOptionsModel.AssetContainerID = core.StringPtr("testString")
				listDataProductDraftsOptionsModel.Version = core.StringPtr("testString")
				listDataProductDraftsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listDataProductDraftsOptionsModel.Start = core.StringPtr("testString")
				listDataProductDraftsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dpxService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := dpxService.ListDataProductDrafts(listDataProductDraftsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListDataProductDraftsOptions model with no property values
				listDataProductDraftsOptionsModelNew := new(dpxv1.ListDataProductDraftsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = dpxService.ListDataProductDrafts(listDataProductDraftsOptionsModelNew)
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
			It(`Invoke ListDataProductDrafts successfully`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Construct an instance of the ListDataProductDraftsOptions model
				listDataProductDraftsOptionsModel := new(dpxv1.ListDataProductDraftsOptions)
				listDataProductDraftsOptionsModel.DataProductID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")
				listDataProductDraftsOptionsModel.AssetContainerID = core.StringPtr("testString")
				listDataProductDraftsOptionsModel.Version = core.StringPtr("testString")
				listDataProductDraftsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listDataProductDraftsOptionsModel.Start = core.StringPtr("testString")
				listDataProductDraftsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := dpxService.ListDataProductDrafts(listDataProductDraftsOptionsModel)
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
				responseObject := new(dpxv1.DataProductDraftCollection)
				nextObject := new(dpxv1.NextPage)
				nextObject.Start = core.StringPtr("abc-123")
				responseObject.Next = nextObject

				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(Equal(core.StringPtr("abc-123")))
			})
			It(`Invoke GetNextStart without a "Next" property in the response`, func() {
				responseObject := new(dpxv1.DataProductDraftCollection)

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
					Expect(req.URL.EscapedPath()).To(Equal(listDataProductDraftsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					requestNumber++
					if requestNumber == 1 {
						fmt.Fprintf(res, "%s", `{"next":{"start":"1"},"total_count":2,"limit":1,"drafts":[{"version":"1.0.0","state":"draft","data_product":{"id":"b38df608-d34b-4d58-8136-ed25e6c6684e"},"name":"My Data Product","description":"This is a description of My Data Product.","id":"2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd","asset":{"id":"2b0bf220-079c-11ee-be56-0242ac120002","container":{"id":"d29c42eb-7100-4b7a-8257-c196dbcca1cd","type":"catalog"}}}]}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"total_count":2,"limit":1,"drafts":[{"version":"1.0.0","state":"draft","data_product":{"id":"b38df608-d34b-4d58-8136-ed25e6c6684e"},"name":"My Data Product","description":"This is a description of My Data Product.","id":"2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd","asset":{"id":"2b0bf220-079c-11ee-be56-0242ac120002","container":{"id":"d29c42eb-7100-4b7a-8257-c196dbcca1cd","type":"catalog"}}}]}`)
					} else {
						res.WriteHeader(400)
					}
				}))
			})
			It(`Use DataProductDraftsPager.GetNext successfully`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				listDataProductDraftsOptionsModel := &dpxv1.ListDataProductDraftsOptions{
					DataProductID:    core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e"),
					AssetContainerID: core.StringPtr("testString"),
					Version:          core.StringPtr("testString"),
					Limit:            core.Int64Ptr(int64(10)),
				}

				pager, err := dpxService.NewDataProductDraftsPager(listDataProductDraftsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				var allResults []dpxv1.DataProductVersionSummary
				for pager.HasNext() {
					nextPage, err := pager.GetNext()
					Expect(err).To(BeNil())
					Expect(nextPage).ToNot(BeNil())
					allResults = append(allResults, nextPage...)
				}
				Expect(len(allResults)).To(Equal(2))
			})
			It(`Use DataProductDraftsPager.GetAll successfully`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				listDataProductDraftsOptionsModel := &dpxv1.ListDataProductDraftsOptions{
					DataProductID:    core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e"),
					AssetContainerID: core.StringPtr("testString"),
					Version:          core.StringPtr("testString"),
					Limit:            core.Int64Ptr(int64(10)),
				}

				pager, err := dpxService.NewDataProductDraftsPager(listDataProductDraftsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				allResults, err := pager.GetAll()
				Expect(err).To(BeNil())
				Expect(allResults).ToNot(BeNil())
				Expect(len(allResults)).To(Equal(2))
			})
		})
	})
	Describe(`CreateDataProductDraft(createDataProductDraftOptions *CreateDataProductDraftOptions) - Operation response error`, func() {
		createDataProductDraftPath := "/data_product_exchange/v1/data_products/b38df608-d34b-4d58-8136-ed25e6c6684e/drafts"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createDataProductDraftPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateDataProductDraft with error: Operation response processing error`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Construct an instance of the ContainerReference model
				containerReferenceModel := new(dpxv1.ContainerReference)
				containerReferenceModel.ID = core.StringPtr("d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				containerReferenceModel.Type = core.StringPtr("catalog")

				// Construct an instance of the AssetReference model
				assetReferenceModel := new(dpxv1.AssetReference)
				assetReferenceModel.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				assetReferenceModel.Container = containerReferenceModel

				// Construct an instance of the DataProductIdentity model
				dataProductIdentityModel := new(dpxv1.DataProductIdentity)
				dataProductIdentityModel.ID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")

				// Construct an instance of the UseCase model
				useCaseModel := new(dpxv1.UseCase)
				useCaseModel.ID = core.StringPtr("testString")
				useCaseModel.Name = core.StringPtr("testString")
				useCaseModel.Container = containerReferenceModel

				// Construct an instance of the Domain model
				domainModel := new(dpxv1.Domain)
				domainModel.ID = core.StringPtr("testString")
				domainModel.Name = core.StringPtr("testString")
				domainModel.Container = containerReferenceModel

				// Construct an instance of the AssetPartReference model
				assetPartReferenceModel := new(dpxv1.AssetPartReference)
				assetPartReferenceModel.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				assetPartReferenceModel.Container = containerReferenceModel
				assetPartReferenceModel.Type = core.StringPtr("data_asset")

				// Construct an instance of the DeliveryMethod model
				deliveryMethodModel := new(dpxv1.DeliveryMethod)
				deliveryMethodModel.ID = core.StringPtr("09cf5fcc-cb9d-4995-a8e4-16517b25229f")
				deliveryMethodModel.Container = containerReferenceModel

				// Construct an instance of the DataProductPart model
				dataProductPartModel := new(dpxv1.DataProductPart)
				dataProductPartModel.Asset = assetPartReferenceModel
				dataProductPartModel.Revision = core.Int64Ptr(int64(1))
				dataProductPartModel.UpdatedAt = CreateMockDateTime("2023-07-01T22:22:34.876Z")
				dataProductPartModel.DeliveryMethods = []dpxv1.DeliveryMethod{*deliveryMethodModel}

				// Construct an instance of the ContractTermsDocumentAttachment model
				contractTermsDocumentAttachmentModel := new(dpxv1.ContractTermsDocumentAttachment)
				contractTermsDocumentAttachmentModel.ID = core.StringPtr("testString")

				// Construct an instance of the ContractTermsDocument model
				contractTermsDocumentModel := new(dpxv1.ContractTermsDocument)
				contractTermsDocumentModel.URL = core.StringPtr("testString")
				contractTermsDocumentModel.Type = core.StringPtr("terms_and_conditions")
				contractTermsDocumentModel.Name = core.StringPtr("testString")
				contractTermsDocumentModel.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				contractTermsDocumentModel.Attachment = contractTermsDocumentAttachmentModel
				contractTermsDocumentModel.UploadURL = core.StringPtr("testString")

				// Construct an instance of the DataProductContractTerms model
				dataProductContractTermsModel := new(dpxv1.DataProductContractTerms)
				dataProductContractTermsModel.Asset = assetReferenceModel
				dataProductContractTermsModel.ID = core.StringPtr("testString")
				dataProductContractTermsModel.Documents = []dpxv1.ContractTermsDocument{*contractTermsDocumentModel}

				// Construct an instance of the CreateDataProductDraftOptions model
				createDataProductDraftOptionsModel := new(dpxv1.CreateDataProductDraftOptions)
				createDataProductDraftOptionsModel.DataProductID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")
				createDataProductDraftOptionsModel.Asset = assetReferenceModel
				createDataProductDraftOptionsModel.Version = core.StringPtr("1.2.0")
				createDataProductDraftOptionsModel.State = core.StringPtr("draft")
				createDataProductDraftOptionsModel.DataProduct = dataProductIdentityModel
				createDataProductDraftOptionsModel.Name = core.StringPtr("testString")
				createDataProductDraftOptionsModel.Description = core.StringPtr("testString")
				createDataProductDraftOptionsModel.Tags = []string{"testString"}
				createDataProductDraftOptionsModel.UseCases = []dpxv1.UseCase{*useCaseModel}
				createDataProductDraftOptionsModel.Domain = domainModel
				createDataProductDraftOptionsModel.Types = []string{"data"}
				createDataProductDraftOptionsModel.PartsOut = []dpxv1.DataProductPart{*dataProductPartModel}
				createDataProductDraftOptionsModel.ContractTerms = []dpxv1.DataProductContractTerms{*dataProductContractTermsModel}
				createDataProductDraftOptionsModel.IsRestricted = core.BoolPtr(true)
				createDataProductDraftOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := dpxService.CreateDataProductDraft(createDataProductDraftOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				dpxService.EnableRetries(0, 0)
				result, response, operationErr = dpxService.CreateDataProductDraft(createDataProductDraftOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateDataProductDraft(createDataProductDraftOptions *CreateDataProductDraftOptions)`, func() {
		createDataProductDraftPath := "/data_product_exchange/v1/data_products/b38df608-d34b-4d58-8136-ed25e6c6684e/drafts"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createDataProductDraftPath))
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
					fmt.Fprintf(res, "%s", `{"version": "1.0.0", "state": "draft", "data_product": {"id": "b38df608-d34b-4d58-8136-ed25e6c6684e"}, "name": "My Data Product", "description": "This is a description of My Data Product.", "id": "2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd", "asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "tags": ["Tags"], "use_cases": [{"id": "ID", "name": "Name", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}], "domain": {"id": "ID", "name": "Name", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "types": ["data"], "parts_out": [{"asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}, "type": "data_asset"}, "revision": 1, "updated_at": "2023-07-01T22:22:34.876Z", "delivery_methods": [{"id": "09cf5fcc-cb9d-4995-a8e4-16517b25229f", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}]}], "published_by": "PublishedBy", "published_at": "2019-01-01T12:00:00.000Z", "contract_terms": [{"asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "id": "ID", "documents": [{"url": "URL", "type": "terms_and_conditions", "name": "Name", "id": "2b0bf220-079c-11ee-be56-0242ac120002", "attachment": {"id": "ID"}, "upload_url": "UploadURL"}]}], "created_by": "CreatedBy", "created_at": "2019-01-01T12:00:00.000Z", "is_restricted": true}`)
				}))
			})
			It(`Invoke CreateDataProductDraft successfully with retries`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())
				dpxService.EnableRetries(0, 0)

				// Construct an instance of the ContainerReference model
				containerReferenceModel := new(dpxv1.ContainerReference)
				containerReferenceModel.ID = core.StringPtr("d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				containerReferenceModel.Type = core.StringPtr("catalog")

				// Construct an instance of the AssetReference model
				assetReferenceModel := new(dpxv1.AssetReference)
				assetReferenceModel.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				assetReferenceModel.Container = containerReferenceModel

				// Construct an instance of the DataProductIdentity model
				dataProductIdentityModel := new(dpxv1.DataProductIdentity)
				dataProductIdentityModel.ID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")

				// Construct an instance of the UseCase model
				useCaseModel := new(dpxv1.UseCase)
				useCaseModel.ID = core.StringPtr("testString")
				useCaseModel.Name = core.StringPtr("testString")
				useCaseModel.Container = containerReferenceModel

				// Construct an instance of the Domain model
				domainModel := new(dpxv1.Domain)
				domainModel.ID = core.StringPtr("testString")
				domainModel.Name = core.StringPtr("testString")
				domainModel.Container = containerReferenceModel

				// Construct an instance of the AssetPartReference model
				assetPartReferenceModel := new(dpxv1.AssetPartReference)
				assetPartReferenceModel.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				assetPartReferenceModel.Container = containerReferenceModel
				assetPartReferenceModel.Type = core.StringPtr("data_asset")

				// Construct an instance of the DeliveryMethod model
				deliveryMethodModel := new(dpxv1.DeliveryMethod)
				deliveryMethodModel.ID = core.StringPtr("09cf5fcc-cb9d-4995-a8e4-16517b25229f")
				deliveryMethodModel.Container = containerReferenceModel

				// Construct an instance of the DataProductPart model
				dataProductPartModel := new(dpxv1.DataProductPart)
				dataProductPartModel.Asset = assetPartReferenceModel
				dataProductPartModel.Revision = core.Int64Ptr(int64(1))
				dataProductPartModel.UpdatedAt = CreateMockDateTime("2023-07-01T22:22:34.876Z")
				dataProductPartModel.DeliveryMethods = []dpxv1.DeliveryMethod{*deliveryMethodModel}

				// Construct an instance of the ContractTermsDocumentAttachment model
				contractTermsDocumentAttachmentModel := new(dpxv1.ContractTermsDocumentAttachment)
				contractTermsDocumentAttachmentModel.ID = core.StringPtr("testString")

				// Construct an instance of the ContractTermsDocument model
				contractTermsDocumentModel := new(dpxv1.ContractTermsDocument)
				contractTermsDocumentModel.URL = core.StringPtr("testString")
				contractTermsDocumentModel.Type = core.StringPtr("terms_and_conditions")
				contractTermsDocumentModel.Name = core.StringPtr("testString")
				contractTermsDocumentModel.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				contractTermsDocumentModel.Attachment = contractTermsDocumentAttachmentModel
				contractTermsDocumentModel.UploadURL = core.StringPtr("testString")

				// Construct an instance of the DataProductContractTerms model
				dataProductContractTermsModel := new(dpxv1.DataProductContractTerms)
				dataProductContractTermsModel.Asset = assetReferenceModel
				dataProductContractTermsModel.ID = core.StringPtr("testString")
				dataProductContractTermsModel.Documents = []dpxv1.ContractTermsDocument{*contractTermsDocumentModel}

				// Construct an instance of the CreateDataProductDraftOptions model
				createDataProductDraftOptionsModel := new(dpxv1.CreateDataProductDraftOptions)
				createDataProductDraftOptionsModel.DataProductID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")
				createDataProductDraftOptionsModel.Asset = assetReferenceModel
				createDataProductDraftOptionsModel.Version = core.StringPtr("1.2.0")
				createDataProductDraftOptionsModel.State = core.StringPtr("draft")
				createDataProductDraftOptionsModel.DataProduct = dataProductIdentityModel
				createDataProductDraftOptionsModel.Name = core.StringPtr("testString")
				createDataProductDraftOptionsModel.Description = core.StringPtr("testString")
				createDataProductDraftOptionsModel.Tags = []string{"testString"}
				createDataProductDraftOptionsModel.UseCases = []dpxv1.UseCase{*useCaseModel}
				createDataProductDraftOptionsModel.Domain = domainModel
				createDataProductDraftOptionsModel.Types = []string{"data"}
				createDataProductDraftOptionsModel.PartsOut = []dpxv1.DataProductPart{*dataProductPartModel}
				createDataProductDraftOptionsModel.ContractTerms = []dpxv1.DataProductContractTerms{*dataProductContractTermsModel}
				createDataProductDraftOptionsModel.IsRestricted = core.BoolPtr(true)
				createDataProductDraftOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := dpxService.CreateDataProductDraftWithContext(ctx, createDataProductDraftOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				dpxService.DisableRetries()
				result, response, operationErr := dpxService.CreateDataProductDraft(createDataProductDraftOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = dpxService.CreateDataProductDraftWithContext(ctx, createDataProductDraftOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createDataProductDraftPath))
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
					fmt.Fprintf(res, "%s", `{"version": "1.0.0", "state": "draft", "data_product": {"id": "b38df608-d34b-4d58-8136-ed25e6c6684e"}, "name": "My Data Product", "description": "This is a description of My Data Product.", "id": "2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd", "asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "tags": ["Tags"], "use_cases": [{"id": "ID", "name": "Name", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}], "domain": {"id": "ID", "name": "Name", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "types": ["data"], "parts_out": [{"asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}, "type": "data_asset"}, "revision": 1, "updated_at": "2023-07-01T22:22:34.876Z", "delivery_methods": [{"id": "09cf5fcc-cb9d-4995-a8e4-16517b25229f", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}]}], "published_by": "PublishedBy", "published_at": "2019-01-01T12:00:00.000Z", "contract_terms": [{"asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "id": "ID", "documents": [{"url": "URL", "type": "terms_and_conditions", "name": "Name", "id": "2b0bf220-079c-11ee-be56-0242ac120002", "attachment": {"id": "ID"}, "upload_url": "UploadURL"}]}], "created_by": "CreatedBy", "created_at": "2019-01-01T12:00:00.000Z", "is_restricted": true}`)
				}))
			})
			It(`Invoke CreateDataProductDraft successfully`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := dpxService.CreateDataProductDraft(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ContainerReference model
				containerReferenceModel := new(dpxv1.ContainerReference)
				containerReferenceModel.ID = core.StringPtr("d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				containerReferenceModel.Type = core.StringPtr("catalog")

				// Construct an instance of the AssetReference model
				assetReferenceModel := new(dpxv1.AssetReference)
				assetReferenceModel.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				assetReferenceModel.Container = containerReferenceModel

				// Construct an instance of the DataProductIdentity model
				dataProductIdentityModel := new(dpxv1.DataProductIdentity)
				dataProductIdentityModel.ID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")

				// Construct an instance of the UseCase model
				useCaseModel := new(dpxv1.UseCase)
				useCaseModel.ID = core.StringPtr("testString")
				useCaseModel.Name = core.StringPtr("testString")
				useCaseModel.Container = containerReferenceModel

				// Construct an instance of the Domain model
				domainModel := new(dpxv1.Domain)
				domainModel.ID = core.StringPtr("testString")
				domainModel.Name = core.StringPtr("testString")
				domainModel.Container = containerReferenceModel

				// Construct an instance of the AssetPartReference model
				assetPartReferenceModel := new(dpxv1.AssetPartReference)
				assetPartReferenceModel.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				assetPartReferenceModel.Container = containerReferenceModel
				assetPartReferenceModel.Type = core.StringPtr("data_asset")

				// Construct an instance of the DeliveryMethod model
				deliveryMethodModel := new(dpxv1.DeliveryMethod)
				deliveryMethodModel.ID = core.StringPtr("09cf5fcc-cb9d-4995-a8e4-16517b25229f")
				deliveryMethodModel.Container = containerReferenceModel

				// Construct an instance of the DataProductPart model
				dataProductPartModel := new(dpxv1.DataProductPart)
				dataProductPartModel.Asset = assetPartReferenceModel
				dataProductPartModel.Revision = core.Int64Ptr(int64(1))
				dataProductPartModel.UpdatedAt = CreateMockDateTime("2023-07-01T22:22:34.876Z")
				dataProductPartModel.DeliveryMethods = []dpxv1.DeliveryMethod{*deliveryMethodModel}

				// Construct an instance of the ContractTermsDocumentAttachment model
				contractTermsDocumentAttachmentModel := new(dpxv1.ContractTermsDocumentAttachment)
				contractTermsDocumentAttachmentModel.ID = core.StringPtr("testString")

				// Construct an instance of the ContractTermsDocument model
				contractTermsDocumentModel := new(dpxv1.ContractTermsDocument)
				contractTermsDocumentModel.URL = core.StringPtr("testString")
				contractTermsDocumentModel.Type = core.StringPtr("terms_and_conditions")
				contractTermsDocumentModel.Name = core.StringPtr("testString")
				contractTermsDocumentModel.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				contractTermsDocumentModel.Attachment = contractTermsDocumentAttachmentModel
				contractTermsDocumentModel.UploadURL = core.StringPtr("testString")

				// Construct an instance of the DataProductContractTerms model
				dataProductContractTermsModel := new(dpxv1.DataProductContractTerms)
				dataProductContractTermsModel.Asset = assetReferenceModel
				dataProductContractTermsModel.ID = core.StringPtr("testString")
				dataProductContractTermsModel.Documents = []dpxv1.ContractTermsDocument{*contractTermsDocumentModel}

				// Construct an instance of the CreateDataProductDraftOptions model
				createDataProductDraftOptionsModel := new(dpxv1.CreateDataProductDraftOptions)
				createDataProductDraftOptionsModel.DataProductID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")
				createDataProductDraftOptionsModel.Asset = assetReferenceModel
				createDataProductDraftOptionsModel.Version = core.StringPtr("1.2.0")
				createDataProductDraftOptionsModel.State = core.StringPtr("draft")
				createDataProductDraftOptionsModel.DataProduct = dataProductIdentityModel
				createDataProductDraftOptionsModel.Name = core.StringPtr("testString")
				createDataProductDraftOptionsModel.Description = core.StringPtr("testString")
				createDataProductDraftOptionsModel.Tags = []string{"testString"}
				createDataProductDraftOptionsModel.UseCases = []dpxv1.UseCase{*useCaseModel}
				createDataProductDraftOptionsModel.Domain = domainModel
				createDataProductDraftOptionsModel.Types = []string{"data"}
				createDataProductDraftOptionsModel.PartsOut = []dpxv1.DataProductPart{*dataProductPartModel}
				createDataProductDraftOptionsModel.ContractTerms = []dpxv1.DataProductContractTerms{*dataProductContractTermsModel}
				createDataProductDraftOptionsModel.IsRestricted = core.BoolPtr(true)
				createDataProductDraftOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = dpxService.CreateDataProductDraft(createDataProductDraftOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateDataProductDraft with error: Operation validation and request error`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Construct an instance of the ContainerReference model
				containerReferenceModel := new(dpxv1.ContainerReference)
				containerReferenceModel.ID = core.StringPtr("d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				containerReferenceModel.Type = core.StringPtr("catalog")

				// Construct an instance of the AssetReference model
				assetReferenceModel := new(dpxv1.AssetReference)
				assetReferenceModel.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				assetReferenceModel.Container = containerReferenceModel

				// Construct an instance of the DataProductIdentity model
				dataProductIdentityModel := new(dpxv1.DataProductIdentity)
				dataProductIdentityModel.ID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")

				// Construct an instance of the UseCase model
				useCaseModel := new(dpxv1.UseCase)
				useCaseModel.ID = core.StringPtr("testString")
				useCaseModel.Name = core.StringPtr("testString")
				useCaseModel.Container = containerReferenceModel

				// Construct an instance of the Domain model
				domainModel := new(dpxv1.Domain)
				domainModel.ID = core.StringPtr("testString")
				domainModel.Name = core.StringPtr("testString")
				domainModel.Container = containerReferenceModel

				// Construct an instance of the AssetPartReference model
				assetPartReferenceModel := new(dpxv1.AssetPartReference)
				assetPartReferenceModel.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				assetPartReferenceModel.Container = containerReferenceModel
				assetPartReferenceModel.Type = core.StringPtr("data_asset")

				// Construct an instance of the DeliveryMethod model
				deliveryMethodModel := new(dpxv1.DeliveryMethod)
				deliveryMethodModel.ID = core.StringPtr("09cf5fcc-cb9d-4995-a8e4-16517b25229f")
				deliveryMethodModel.Container = containerReferenceModel

				// Construct an instance of the DataProductPart model
				dataProductPartModel := new(dpxv1.DataProductPart)
				dataProductPartModel.Asset = assetPartReferenceModel
				dataProductPartModel.Revision = core.Int64Ptr(int64(1))
				dataProductPartModel.UpdatedAt = CreateMockDateTime("2023-07-01T22:22:34.876Z")
				dataProductPartModel.DeliveryMethods = []dpxv1.DeliveryMethod{*deliveryMethodModel}

				// Construct an instance of the ContractTermsDocumentAttachment model
				contractTermsDocumentAttachmentModel := new(dpxv1.ContractTermsDocumentAttachment)
				contractTermsDocumentAttachmentModel.ID = core.StringPtr("testString")

				// Construct an instance of the ContractTermsDocument model
				contractTermsDocumentModel := new(dpxv1.ContractTermsDocument)
				contractTermsDocumentModel.URL = core.StringPtr("testString")
				contractTermsDocumentModel.Type = core.StringPtr("terms_and_conditions")
				contractTermsDocumentModel.Name = core.StringPtr("testString")
				contractTermsDocumentModel.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				contractTermsDocumentModel.Attachment = contractTermsDocumentAttachmentModel
				contractTermsDocumentModel.UploadURL = core.StringPtr("testString")

				// Construct an instance of the DataProductContractTerms model
				dataProductContractTermsModel := new(dpxv1.DataProductContractTerms)
				dataProductContractTermsModel.Asset = assetReferenceModel
				dataProductContractTermsModel.ID = core.StringPtr("testString")
				dataProductContractTermsModel.Documents = []dpxv1.ContractTermsDocument{*contractTermsDocumentModel}

				// Construct an instance of the CreateDataProductDraftOptions model
				createDataProductDraftOptionsModel := new(dpxv1.CreateDataProductDraftOptions)
				createDataProductDraftOptionsModel.DataProductID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")
				createDataProductDraftOptionsModel.Asset = assetReferenceModel
				createDataProductDraftOptionsModel.Version = core.StringPtr("1.2.0")
				createDataProductDraftOptionsModel.State = core.StringPtr("draft")
				createDataProductDraftOptionsModel.DataProduct = dataProductIdentityModel
				createDataProductDraftOptionsModel.Name = core.StringPtr("testString")
				createDataProductDraftOptionsModel.Description = core.StringPtr("testString")
				createDataProductDraftOptionsModel.Tags = []string{"testString"}
				createDataProductDraftOptionsModel.UseCases = []dpxv1.UseCase{*useCaseModel}
				createDataProductDraftOptionsModel.Domain = domainModel
				createDataProductDraftOptionsModel.Types = []string{"data"}
				createDataProductDraftOptionsModel.PartsOut = []dpxv1.DataProductPart{*dataProductPartModel}
				createDataProductDraftOptionsModel.ContractTerms = []dpxv1.DataProductContractTerms{*dataProductContractTermsModel}
				createDataProductDraftOptionsModel.IsRestricted = core.BoolPtr(true)
				createDataProductDraftOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dpxService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := dpxService.CreateDataProductDraft(createDataProductDraftOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateDataProductDraftOptions model with no property values
				createDataProductDraftOptionsModelNew := new(dpxv1.CreateDataProductDraftOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = dpxService.CreateDataProductDraft(createDataProductDraftOptionsModelNew)
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
			It(`Invoke CreateDataProductDraft successfully`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Construct an instance of the ContainerReference model
				containerReferenceModel := new(dpxv1.ContainerReference)
				containerReferenceModel.ID = core.StringPtr("d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				containerReferenceModel.Type = core.StringPtr("catalog")

				// Construct an instance of the AssetReference model
				assetReferenceModel := new(dpxv1.AssetReference)
				assetReferenceModel.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				assetReferenceModel.Container = containerReferenceModel

				// Construct an instance of the DataProductIdentity model
				dataProductIdentityModel := new(dpxv1.DataProductIdentity)
				dataProductIdentityModel.ID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")

				// Construct an instance of the UseCase model
				useCaseModel := new(dpxv1.UseCase)
				useCaseModel.ID = core.StringPtr("testString")
				useCaseModel.Name = core.StringPtr("testString")
				useCaseModel.Container = containerReferenceModel

				// Construct an instance of the Domain model
				domainModel := new(dpxv1.Domain)
				domainModel.ID = core.StringPtr("testString")
				domainModel.Name = core.StringPtr("testString")
				domainModel.Container = containerReferenceModel

				// Construct an instance of the AssetPartReference model
				assetPartReferenceModel := new(dpxv1.AssetPartReference)
				assetPartReferenceModel.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				assetPartReferenceModel.Container = containerReferenceModel
				assetPartReferenceModel.Type = core.StringPtr("data_asset")

				// Construct an instance of the DeliveryMethod model
				deliveryMethodModel := new(dpxv1.DeliveryMethod)
				deliveryMethodModel.ID = core.StringPtr("09cf5fcc-cb9d-4995-a8e4-16517b25229f")
				deliveryMethodModel.Container = containerReferenceModel

				// Construct an instance of the DataProductPart model
				dataProductPartModel := new(dpxv1.DataProductPart)
				dataProductPartModel.Asset = assetPartReferenceModel
				dataProductPartModel.Revision = core.Int64Ptr(int64(1))
				dataProductPartModel.UpdatedAt = CreateMockDateTime("2023-07-01T22:22:34.876Z")
				dataProductPartModel.DeliveryMethods = []dpxv1.DeliveryMethod{*deliveryMethodModel}

				// Construct an instance of the ContractTermsDocumentAttachment model
				contractTermsDocumentAttachmentModel := new(dpxv1.ContractTermsDocumentAttachment)
				contractTermsDocumentAttachmentModel.ID = core.StringPtr("testString")

				// Construct an instance of the ContractTermsDocument model
				contractTermsDocumentModel := new(dpxv1.ContractTermsDocument)
				contractTermsDocumentModel.URL = core.StringPtr("testString")
				contractTermsDocumentModel.Type = core.StringPtr("terms_and_conditions")
				contractTermsDocumentModel.Name = core.StringPtr("testString")
				contractTermsDocumentModel.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				contractTermsDocumentModel.Attachment = contractTermsDocumentAttachmentModel
				contractTermsDocumentModel.UploadURL = core.StringPtr("testString")

				// Construct an instance of the DataProductContractTerms model
				dataProductContractTermsModel := new(dpxv1.DataProductContractTerms)
				dataProductContractTermsModel.Asset = assetReferenceModel
				dataProductContractTermsModel.ID = core.StringPtr("testString")
				dataProductContractTermsModel.Documents = []dpxv1.ContractTermsDocument{*contractTermsDocumentModel}

				// Construct an instance of the CreateDataProductDraftOptions model
				createDataProductDraftOptionsModel := new(dpxv1.CreateDataProductDraftOptions)
				createDataProductDraftOptionsModel.DataProductID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")
				createDataProductDraftOptionsModel.Asset = assetReferenceModel
				createDataProductDraftOptionsModel.Version = core.StringPtr("1.2.0")
				createDataProductDraftOptionsModel.State = core.StringPtr("draft")
				createDataProductDraftOptionsModel.DataProduct = dataProductIdentityModel
				createDataProductDraftOptionsModel.Name = core.StringPtr("testString")
				createDataProductDraftOptionsModel.Description = core.StringPtr("testString")
				createDataProductDraftOptionsModel.Tags = []string{"testString"}
				createDataProductDraftOptionsModel.UseCases = []dpxv1.UseCase{*useCaseModel}
				createDataProductDraftOptionsModel.Domain = domainModel
				createDataProductDraftOptionsModel.Types = []string{"data"}
				createDataProductDraftOptionsModel.PartsOut = []dpxv1.DataProductPart{*dataProductPartModel}
				createDataProductDraftOptionsModel.ContractTerms = []dpxv1.DataProductContractTerms{*dataProductContractTermsModel}
				createDataProductDraftOptionsModel.IsRestricted = core.BoolPtr(true)
				createDataProductDraftOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := dpxService.CreateDataProductDraft(createDataProductDraftOptionsModel)
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
	Describe(`CreateDraftContractTermsDocument(createDraftContractTermsDocumentOptions *CreateDraftContractTermsDocumentOptions) - Operation response error`, func() {
		createDraftContractTermsDocumentPath := "/data_product_exchange/v1/data_products/b38df608-d34b-4d58-8136-ed25e6c6684e/drafts/2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd/contract_terms/598183cd-b910-4e8d-9a97-97097afda3c1@e4fe2f87-0e56-46dd-b3b8-e9af32309e82/documents"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createDraftContractTermsDocumentPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateDraftContractTermsDocument with error: Operation response processing error`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Construct an instance of the ContractTermsDocumentAttachment model
				contractTermsDocumentAttachmentModel := new(dpxv1.ContractTermsDocumentAttachment)
				contractTermsDocumentAttachmentModel.ID = core.StringPtr("testString")

				// Construct an instance of the CreateDraftContractTermsDocumentOptions model
				createDraftContractTermsDocumentOptionsModel := new(dpxv1.CreateDraftContractTermsDocumentOptions)
				createDraftContractTermsDocumentOptionsModel.DataProductID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")
				createDraftContractTermsDocumentOptionsModel.DraftID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				createDraftContractTermsDocumentOptionsModel.ContractTermsID = core.StringPtr("598183cd-b910-4e8d-9a97-97097afda3c1@e4fe2f87-0e56-46dd-b3b8-e9af32309e82")
				createDraftContractTermsDocumentOptionsModel.Type = core.StringPtr("terms_and_conditions")
				createDraftContractTermsDocumentOptionsModel.Name = core.StringPtr("Terms and conditions document")
				createDraftContractTermsDocumentOptionsModel.ID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")
				createDraftContractTermsDocumentOptionsModel.URL = core.StringPtr("testString")
				createDraftContractTermsDocumentOptionsModel.Attachment = contractTermsDocumentAttachmentModel
				createDraftContractTermsDocumentOptionsModel.UploadURL = core.StringPtr("testString")
				createDraftContractTermsDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := dpxService.CreateDraftContractTermsDocument(createDraftContractTermsDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				dpxService.EnableRetries(0, 0)
				result, response, operationErr = dpxService.CreateDraftContractTermsDocument(createDraftContractTermsDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateDraftContractTermsDocument(createDraftContractTermsDocumentOptions *CreateDraftContractTermsDocumentOptions)`, func() {
		createDraftContractTermsDocumentPath := "/data_product_exchange/v1/data_products/b38df608-d34b-4d58-8136-ed25e6c6684e/drafts/2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd/contract_terms/598183cd-b910-4e8d-9a97-97097afda3c1@e4fe2f87-0e56-46dd-b3b8-e9af32309e82/documents"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createDraftContractTermsDocumentPath))
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
					fmt.Fprintf(res, "%s", `{"url": "URL", "type": "terms_and_conditions", "name": "Name", "id": "2b0bf220-079c-11ee-be56-0242ac120002", "attachment": {"id": "ID"}, "upload_url": "UploadURL"}`)
				}))
			})
			It(`Invoke CreateDraftContractTermsDocument successfully with retries`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())
				dpxService.EnableRetries(0, 0)

				// Construct an instance of the ContractTermsDocumentAttachment model
				contractTermsDocumentAttachmentModel := new(dpxv1.ContractTermsDocumentAttachment)
				contractTermsDocumentAttachmentModel.ID = core.StringPtr("testString")

				// Construct an instance of the CreateDraftContractTermsDocumentOptions model
				createDraftContractTermsDocumentOptionsModel := new(dpxv1.CreateDraftContractTermsDocumentOptions)
				createDraftContractTermsDocumentOptionsModel.DataProductID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")
				createDraftContractTermsDocumentOptionsModel.DraftID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				createDraftContractTermsDocumentOptionsModel.ContractTermsID = core.StringPtr("598183cd-b910-4e8d-9a97-97097afda3c1@e4fe2f87-0e56-46dd-b3b8-e9af32309e82")
				createDraftContractTermsDocumentOptionsModel.Type = core.StringPtr("terms_and_conditions")
				createDraftContractTermsDocumentOptionsModel.Name = core.StringPtr("Terms and conditions document")
				createDraftContractTermsDocumentOptionsModel.ID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")
				createDraftContractTermsDocumentOptionsModel.URL = core.StringPtr("testString")
				createDraftContractTermsDocumentOptionsModel.Attachment = contractTermsDocumentAttachmentModel
				createDraftContractTermsDocumentOptionsModel.UploadURL = core.StringPtr("testString")
				createDraftContractTermsDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := dpxService.CreateDraftContractTermsDocumentWithContext(ctx, createDraftContractTermsDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				dpxService.DisableRetries()
				result, response, operationErr := dpxService.CreateDraftContractTermsDocument(createDraftContractTermsDocumentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = dpxService.CreateDraftContractTermsDocumentWithContext(ctx, createDraftContractTermsDocumentOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createDraftContractTermsDocumentPath))
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
					fmt.Fprintf(res, "%s", `{"url": "URL", "type": "terms_and_conditions", "name": "Name", "id": "2b0bf220-079c-11ee-be56-0242ac120002", "attachment": {"id": "ID"}, "upload_url": "UploadURL"}`)
				}))
			})
			It(`Invoke CreateDraftContractTermsDocument successfully`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := dpxService.CreateDraftContractTermsDocument(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ContractTermsDocumentAttachment model
				contractTermsDocumentAttachmentModel := new(dpxv1.ContractTermsDocumentAttachment)
				contractTermsDocumentAttachmentModel.ID = core.StringPtr("testString")

				// Construct an instance of the CreateDraftContractTermsDocumentOptions model
				createDraftContractTermsDocumentOptionsModel := new(dpxv1.CreateDraftContractTermsDocumentOptions)
				createDraftContractTermsDocumentOptionsModel.DataProductID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")
				createDraftContractTermsDocumentOptionsModel.DraftID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				createDraftContractTermsDocumentOptionsModel.ContractTermsID = core.StringPtr("598183cd-b910-4e8d-9a97-97097afda3c1@e4fe2f87-0e56-46dd-b3b8-e9af32309e82")
				createDraftContractTermsDocumentOptionsModel.Type = core.StringPtr("terms_and_conditions")
				createDraftContractTermsDocumentOptionsModel.Name = core.StringPtr("Terms and conditions document")
				createDraftContractTermsDocumentOptionsModel.ID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")
				createDraftContractTermsDocumentOptionsModel.URL = core.StringPtr("testString")
				createDraftContractTermsDocumentOptionsModel.Attachment = contractTermsDocumentAttachmentModel
				createDraftContractTermsDocumentOptionsModel.UploadURL = core.StringPtr("testString")
				createDraftContractTermsDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = dpxService.CreateDraftContractTermsDocument(createDraftContractTermsDocumentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateDraftContractTermsDocument with error: Operation validation and request error`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Construct an instance of the ContractTermsDocumentAttachment model
				contractTermsDocumentAttachmentModel := new(dpxv1.ContractTermsDocumentAttachment)
				contractTermsDocumentAttachmentModel.ID = core.StringPtr("testString")

				// Construct an instance of the CreateDraftContractTermsDocumentOptions model
				createDraftContractTermsDocumentOptionsModel := new(dpxv1.CreateDraftContractTermsDocumentOptions)
				createDraftContractTermsDocumentOptionsModel.DataProductID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")
				createDraftContractTermsDocumentOptionsModel.DraftID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				createDraftContractTermsDocumentOptionsModel.ContractTermsID = core.StringPtr("598183cd-b910-4e8d-9a97-97097afda3c1@e4fe2f87-0e56-46dd-b3b8-e9af32309e82")
				createDraftContractTermsDocumentOptionsModel.Type = core.StringPtr("terms_and_conditions")
				createDraftContractTermsDocumentOptionsModel.Name = core.StringPtr("Terms and conditions document")
				createDraftContractTermsDocumentOptionsModel.ID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")
				createDraftContractTermsDocumentOptionsModel.URL = core.StringPtr("testString")
				createDraftContractTermsDocumentOptionsModel.Attachment = contractTermsDocumentAttachmentModel
				createDraftContractTermsDocumentOptionsModel.UploadURL = core.StringPtr("testString")
				createDraftContractTermsDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dpxService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := dpxService.CreateDraftContractTermsDocument(createDraftContractTermsDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateDraftContractTermsDocumentOptions model with no property values
				createDraftContractTermsDocumentOptionsModelNew := new(dpxv1.CreateDraftContractTermsDocumentOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = dpxService.CreateDraftContractTermsDocument(createDraftContractTermsDocumentOptionsModelNew)
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
			It(`Invoke CreateDraftContractTermsDocument successfully`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Construct an instance of the ContractTermsDocumentAttachment model
				contractTermsDocumentAttachmentModel := new(dpxv1.ContractTermsDocumentAttachment)
				contractTermsDocumentAttachmentModel.ID = core.StringPtr("testString")

				// Construct an instance of the CreateDraftContractTermsDocumentOptions model
				createDraftContractTermsDocumentOptionsModel := new(dpxv1.CreateDraftContractTermsDocumentOptions)
				createDraftContractTermsDocumentOptionsModel.DataProductID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")
				createDraftContractTermsDocumentOptionsModel.DraftID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				createDraftContractTermsDocumentOptionsModel.ContractTermsID = core.StringPtr("598183cd-b910-4e8d-9a97-97097afda3c1@e4fe2f87-0e56-46dd-b3b8-e9af32309e82")
				createDraftContractTermsDocumentOptionsModel.Type = core.StringPtr("terms_and_conditions")
				createDraftContractTermsDocumentOptionsModel.Name = core.StringPtr("Terms and conditions document")
				createDraftContractTermsDocumentOptionsModel.ID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")
				createDraftContractTermsDocumentOptionsModel.URL = core.StringPtr("testString")
				createDraftContractTermsDocumentOptionsModel.Attachment = contractTermsDocumentAttachmentModel
				createDraftContractTermsDocumentOptionsModel.UploadURL = core.StringPtr("testString")
				createDraftContractTermsDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := dpxService.CreateDraftContractTermsDocument(createDraftContractTermsDocumentOptionsModel)
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
	Describe(`GetDataProductDraft(getDataProductDraftOptions *GetDataProductDraftOptions) - Operation response error`, func() {
		getDataProductDraftPath := "/data_product_exchange/v1/data_products/b38df608-d34b-4d58-8136-ed25e6c6684e/drafts/2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDataProductDraftPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetDataProductDraft with error: Operation response processing error`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Construct an instance of the GetDataProductDraftOptions model
				getDataProductDraftOptionsModel := new(dpxv1.GetDataProductDraftOptions)
				getDataProductDraftOptionsModel.DataProductID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")
				getDataProductDraftOptionsModel.DraftID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				getDataProductDraftOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := dpxService.GetDataProductDraft(getDataProductDraftOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				dpxService.EnableRetries(0, 0)
				result, response, operationErr = dpxService.GetDataProductDraft(getDataProductDraftOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetDataProductDraft(getDataProductDraftOptions *GetDataProductDraftOptions)`, func() {
		getDataProductDraftPath := "/data_product_exchange/v1/data_products/b38df608-d34b-4d58-8136-ed25e6c6684e/drafts/2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDataProductDraftPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"version": "1.0.0", "state": "draft", "data_product": {"id": "b38df608-d34b-4d58-8136-ed25e6c6684e"}, "name": "My Data Product", "description": "This is a description of My Data Product.", "id": "2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd", "asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "tags": ["Tags"], "use_cases": [{"id": "ID", "name": "Name", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}], "domain": {"id": "ID", "name": "Name", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "types": ["data"], "parts_out": [{"asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}, "type": "data_asset"}, "revision": 1, "updated_at": "2023-07-01T22:22:34.876Z", "delivery_methods": [{"id": "09cf5fcc-cb9d-4995-a8e4-16517b25229f", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}]}], "published_by": "PublishedBy", "published_at": "2019-01-01T12:00:00.000Z", "contract_terms": [{"asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "id": "ID", "documents": [{"url": "URL", "type": "terms_and_conditions", "name": "Name", "id": "2b0bf220-079c-11ee-be56-0242ac120002", "attachment": {"id": "ID"}, "upload_url": "UploadURL"}]}], "created_by": "CreatedBy", "created_at": "2019-01-01T12:00:00.000Z", "is_restricted": true}`)
				}))
			})
			It(`Invoke GetDataProductDraft successfully with retries`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())
				dpxService.EnableRetries(0, 0)

				// Construct an instance of the GetDataProductDraftOptions model
				getDataProductDraftOptionsModel := new(dpxv1.GetDataProductDraftOptions)
				getDataProductDraftOptionsModel.DataProductID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")
				getDataProductDraftOptionsModel.DraftID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				getDataProductDraftOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := dpxService.GetDataProductDraftWithContext(ctx, getDataProductDraftOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				dpxService.DisableRetries()
				result, response, operationErr := dpxService.GetDataProductDraft(getDataProductDraftOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = dpxService.GetDataProductDraftWithContext(ctx, getDataProductDraftOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getDataProductDraftPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"version": "1.0.0", "state": "draft", "data_product": {"id": "b38df608-d34b-4d58-8136-ed25e6c6684e"}, "name": "My Data Product", "description": "This is a description of My Data Product.", "id": "2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd", "asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "tags": ["Tags"], "use_cases": [{"id": "ID", "name": "Name", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}], "domain": {"id": "ID", "name": "Name", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "types": ["data"], "parts_out": [{"asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}, "type": "data_asset"}, "revision": 1, "updated_at": "2023-07-01T22:22:34.876Z", "delivery_methods": [{"id": "09cf5fcc-cb9d-4995-a8e4-16517b25229f", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}]}], "published_by": "PublishedBy", "published_at": "2019-01-01T12:00:00.000Z", "contract_terms": [{"asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "id": "ID", "documents": [{"url": "URL", "type": "terms_and_conditions", "name": "Name", "id": "2b0bf220-079c-11ee-be56-0242ac120002", "attachment": {"id": "ID"}, "upload_url": "UploadURL"}]}], "created_by": "CreatedBy", "created_at": "2019-01-01T12:00:00.000Z", "is_restricted": true}`)
				}))
			})
			It(`Invoke GetDataProductDraft successfully`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := dpxService.GetDataProductDraft(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetDataProductDraftOptions model
				getDataProductDraftOptionsModel := new(dpxv1.GetDataProductDraftOptions)
				getDataProductDraftOptionsModel.DataProductID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")
				getDataProductDraftOptionsModel.DraftID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				getDataProductDraftOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = dpxService.GetDataProductDraft(getDataProductDraftOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetDataProductDraft with error: Operation validation and request error`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Construct an instance of the GetDataProductDraftOptions model
				getDataProductDraftOptionsModel := new(dpxv1.GetDataProductDraftOptions)
				getDataProductDraftOptionsModel.DataProductID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")
				getDataProductDraftOptionsModel.DraftID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				getDataProductDraftOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dpxService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := dpxService.GetDataProductDraft(getDataProductDraftOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetDataProductDraftOptions model with no property values
				getDataProductDraftOptionsModelNew := new(dpxv1.GetDataProductDraftOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = dpxService.GetDataProductDraft(getDataProductDraftOptionsModelNew)
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
			It(`Invoke GetDataProductDraft successfully`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Construct an instance of the GetDataProductDraftOptions model
				getDataProductDraftOptionsModel := new(dpxv1.GetDataProductDraftOptions)
				getDataProductDraftOptionsModel.DataProductID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")
				getDataProductDraftOptionsModel.DraftID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				getDataProductDraftOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := dpxService.GetDataProductDraft(getDataProductDraftOptionsModel)
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
	Describe(`DeleteDataProductDraft(deleteDataProductDraftOptions *DeleteDataProductDraftOptions)`, func() {
		deleteDataProductDraftPath := "/data_product_exchange/v1/data_products/b38df608-d34b-4d58-8136-ed25e6c6684e/drafts/2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteDataProductDraftPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteDataProductDraft successfully`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := dpxService.DeleteDataProductDraft(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteDataProductDraftOptions model
				deleteDataProductDraftOptionsModel := new(dpxv1.DeleteDataProductDraftOptions)
				deleteDataProductDraftOptionsModel.DataProductID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")
				deleteDataProductDraftOptionsModel.DraftID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				deleteDataProductDraftOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = dpxService.DeleteDataProductDraft(deleteDataProductDraftOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteDataProductDraft with error: Operation validation and request error`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Construct an instance of the DeleteDataProductDraftOptions model
				deleteDataProductDraftOptionsModel := new(dpxv1.DeleteDataProductDraftOptions)
				deleteDataProductDraftOptionsModel.DataProductID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")
				deleteDataProductDraftOptionsModel.DraftID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				deleteDataProductDraftOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dpxService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := dpxService.DeleteDataProductDraft(deleteDataProductDraftOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteDataProductDraftOptions model with no property values
				deleteDataProductDraftOptionsModelNew := new(dpxv1.DeleteDataProductDraftOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = dpxService.DeleteDataProductDraft(deleteDataProductDraftOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateDataProductDraft(updateDataProductDraftOptions *UpdateDataProductDraftOptions) - Operation response error`, func() {
		updateDataProductDraftPath := "/data_product_exchange/v1/data_products/b38df608-d34b-4d58-8136-ed25e6c6684e/drafts/2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateDataProductDraftPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateDataProductDraft with error: Operation response processing error`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(dpxv1.JSONPatchOperation)
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = "testString"

				// Construct an instance of the UpdateDataProductDraftOptions model
				updateDataProductDraftOptionsModel := new(dpxv1.UpdateDataProductDraftOptions)
				updateDataProductDraftOptionsModel.DataProductID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")
				updateDataProductDraftOptionsModel.DraftID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				updateDataProductDraftOptionsModel.JSONPatchInstructions = []dpxv1.JSONPatchOperation{*jsonPatchOperationModel}
				updateDataProductDraftOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := dpxService.UpdateDataProductDraft(updateDataProductDraftOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				dpxService.EnableRetries(0, 0)
				result, response, operationErr = dpxService.UpdateDataProductDraft(updateDataProductDraftOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateDataProductDraft(updateDataProductDraftOptions *UpdateDataProductDraftOptions)`, func() {
		updateDataProductDraftPath := "/data_product_exchange/v1/data_products/b38df608-d34b-4d58-8136-ed25e6c6684e/drafts/2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateDataProductDraftPath))
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
					fmt.Fprintf(res, "%s", `{"version": "1.0.0", "state": "draft", "data_product": {"id": "b38df608-d34b-4d58-8136-ed25e6c6684e"}, "name": "My Data Product", "description": "This is a description of My Data Product.", "id": "2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd", "asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "tags": ["Tags"], "use_cases": [{"id": "ID", "name": "Name", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}], "domain": {"id": "ID", "name": "Name", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "types": ["data"], "parts_out": [{"asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}, "type": "data_asset"}, "revision": 1, "updated_at": "2023-07-01T22:22:34.876Z", "delivery_methods": [{"id": "09cf5fcc-cb9d-4995-a8e4-16517b25229f", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}]}], "published_by": "PublishedBy", "published_at": "2019-01-01T12:00:00.000Z", "contract_terms": [{"asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "id": "ID", "documents": [{"url": "URL", "type": "terms_and_conditions", "name": "Name", "id": "2b0bf220-079c-11ee-be56-0242ac120002", "attachment": {"id": "ID"}, "upload_url": "UploadURL"}]}], "created_by": "CreatedBy", "created_at": "2019-01-01T12:00:00.000Z", "is_restricted": true}`)
				}))
			})
			It(`Invoke UpdateDataProductDraft successfully with retries`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())
				dpxService.EnableRetries(0, 0)

				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(dpxv1.JSONPatchOperation)
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = "testString"

				// Construct an instance of the UpdateDataProductDraftOptions model
				updateDataProductDraftOptionsModel := new(dpxv1.UpdateDataProductDraftOptions)
				updateDataProductDraftOptionsModel.DataProductID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")
				updateDataProductDraftOptionsModel.DraftID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				updateDataProductDraftOptionsModel.JSONPatchInstructions = []dpxv1.JSONPatchOperation{*jsonPatchOperationModel}
				updateDataProductDraftOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := dpxService.UpdateDataProductDraftWithContext(ctx, updateDataProductDraftOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				dpxService.DisableRetries()
				result, response, operationErr := dpxService.UpdateDataProductDraft(updateDataProductDraftOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = dpxService.UpdateDataProductDraftWithContext(ctx, updateDataProductDraftOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateDataProductDraftPath))
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
					fmt.Fprintf(res, "%s", `{"version": "1.0.0", "state": "draft", "data_product": {"id": "b38df608-d34b-4d58-8136-ed25e6c6684e"}, "name": "My Data Product", "description": "This is a description of My Data Product.", "id": "2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd", "asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "tags": ["Tags"], "use_cases": [{"id": "ID", "name": "Name", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}], "domain": {"id": "ID", "name": "Name", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "types": ["data"], "parts_out": [{"asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}, "type": "data_asset"}, "revision": 1, "updated_at": "2023-07-01T22:22:34.876Z", "delivery_methods": [{"id": "09cf5fcc-cb9d-4995-a8e4-16517b25229f", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}]}], "published_by": "PublishedBy", "published_at": "2019-01-01T12:00:00.000Z", "contract_terms": [{"asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "id": "ID", "documents": [{"url": "URL", "type": "terms_and_conditions", "name": "Name", "id": "2b0bf220-079c-11ee-be56-0242ac120002", "attachment": {"id": "ID"}, "upload_url": "UploadURL"}]}], "created_by": "CreatedBy", "created_at": "2019-01-01T12:00:00.000Z", "is_restricted": true}`)
				}))
			})
			It(`Invoke UpdateDataProductDraft successfully`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := dpxService.UpdateDataProductDraft(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(dpxv1.JSONPatchOperation)
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = "testString"

				// Construct an instance of the UpdateDataProductDraftOptions model
				updateDataProductDraftOptionsModel := new(dpxv1.UpdateDataProductDraftOptions)
				updateDataProductDraftOptionsModel.DataProductID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")
				updateDataProductDraftOptionsModel.DraftID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				updateDataProductDraftOptionsModel.JSONPatchInstructions = []dpxv1.JSONPatchOperation{*jsonPatchOperationModel}
				updateDataProductDraftOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = dpxService.UpdateDataProductDraft(updateDataProductDraftOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateDataProductDraft with error: Operation validation and request error`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(dpxv1.JSONPatchOperation)
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = "testString"

				// Construct an instance of the UpdateDataProductDraftOptions model
				updateDataProductDraftOptionsModel := new(dpxv1.UpdateDataProductDraftOptions)
				updateDataProductDraftOptionsModel.DataProductID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")
				updateDataProductDraftOptionsModel.DraftID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				updateDataProductDraftOptionsModel.JSONPatchInstructions = []dpxv1.JSONPatchOperation{*jsonPatchOperationModel}
				updateDataProductDraftOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dpxService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := dpxService.UpdateDataProductDraft(updateDataProductDraftOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateDataProductDraftOptions model with no property values
				updateDataProductDraftOptionsModelNew := new(dpxv1.UpdateDataProductDraftOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = dpxService.UpdateDataProductDraft(updateDataProductDraftOptionsModelNew)
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
			It(`Invoke UpdateDataProductDraft successfully`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(dpxv1.JSONPatchOperation)
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = "testString"

				// Construct an instance of the UpdateDataProductDraftOptions model
				updateDataProductDraftOptionsModel := new(dpxv1.UpdateDataProductDraftOptions)
				updateDataProductDraftOptionsModel.DataProductID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")
				updateDataProductDraftOptionsModel.DraftID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				updateDataProductDraftOptionsModel.JSONPatchInstructions = []dpxv1.JSONPatchOperation{*jsonPatchOperationModel}
				updateDataProductDraftOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := dpxService.UpdateDataProductDraft(updateDataProductDraftOptionsModel)
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
	Describe(`GetDraftContractTermsDocument(getDraftContractTermsDocumentOptions *GetDraftContractTermsDocumentOptions) - Operation response error`, func() {
		getDraftContractTermsDocumentPath := "/data_product_exchange/v1/data_products/b38df608-d34b-4d58-8136-ed25e6c6684e/drafts/2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd/contract_terms/598183cd-b910-4e8d-9a97-97097afda3c1@e4fe2f87-0e56-46dd-b3b8-e9af32309e82/documents/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDraftContractTermsDocumentPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetDraftContractTermsDocument with error: Operation response processing error`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Construct an instance of the GetDraftContractTermsDocumentOptions model
				getDraftContractTermsDocumentOptionsModel := new(dpxv1.GetDraftContractTermsDocumentOptions)
				getDraftContractTermsDocumentOptionsModel.DataProductID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")
				getDraftContractTermsDocumentOptionsModel.DraftID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				getDraftContractTermsDocumentOptionsModel.ContractTermsID = core.StringPtr("598183cd-b910-4e8d-9a97-97097afda3c1@e4fe2f87-0e56-46dd-b3b8-e9af32309e82")
				getDraftContractTermsDocumentOptionsModel.DocumentID = core.StringPtr("testString")
				getDraftContractTermsDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := dpxService.GetDraftContractTermsDocument(getDraftContractTermsDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				dpxService.EnableRetries(0, 0)
				result, response, operationErr = dpxService.GetDraftContractTermsDocument(getDraftContractTermsDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetDraftContractTermsDocument(getDraftContractTermsDocumentOptions *GetDraftContractTermsDocumentOptions)`, func() {
		getDraftContractTermsDocumentPath := "/data_product_exchange/v1/data_products/b38df608-d34b-4d58-8136-ed25e6c6684e/drafts/2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd/contract_terms/598183cd-b910-4e8d-9a97-97097afda3c1@e4fe2f87-0e56-46dd-b3b8-e9af32309e82/documents/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDraftContractTermsDocumentPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"url": "URL", "type": "terms_and_conditions", "name": "Name", "id": "2b0bf220-079c-11ee-be56-0242ac120002", "attachment": {"id": "ID"}, "upload_url": "UploadURL"}`)
				}))
			})
			It(`Invoke GetDraftContractTermsDocument successfully with retries`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())
				dpxService.EnableRetries(0, 0)

				// Construct an instance of the GetDraftContractTermsDocumentOptions model
				getDraftContractTermsDocumentOptionsModel := new(dpxv1.GetDraftContractTermsDocumentOptions)
				getDraftContractTermsDocumentOptionsModel.DataProductID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")
				getDraftContractTermsDocumentOptionsModel.DraftID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				getDraftContractTermsDocumentOptionsModel.ContractTermsID = core.StringPtr("598183cd-b910-4e8d-9a97-97097afda3c1@e4fe2f87-0e56-46dd-b3b8-e9af32309e82")
				getDraftContractTermsDocumentOptionsModel.DocumentID = core.StringPtr("testString")
				getDraftContractTermsDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := dpxService.GetDraftContractTermsDocumentWithContext(ctx, getDraftContractTermsDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				dpxService.DisableRetries()
				result, response, operationErr := dpxService.GetDraftContractTermsDocument(getDraftContractTermsDocumentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = dpxService.GetDraftContractTermsDocumentWithContext(ctx, getDraftContractTermsDocumentOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getDraftContractTermsDocumentPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"url": "URL", "type": "terms_and_conditions", "name": "Name", "id": "2b0bf220-079c-11ee-be56-0242ac120002", "attachment": {"id": "ID"}, "upload_url": "UploadURL"}`)
				}))
			})
			It(`Invoke GetDraftContractTermsDocument successfully`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := dpxService.GetDraftContractTermsDocument(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetDraftContractTermsDocumentOptions model
				getDraftContractTermsDocumentOptionsModel := new(dpxv1.GetDraftContractTermsDocumentOptions)
				getDraftContractTermsDocumentOptionsModel.DataProductID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")
				getDraftContractTermsDocumentOptionsModel.DraftID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				getDraftContractTermsDocumentOptionsModel.ContractTermsID = core.StringPtr("598183cd-b910-4e8d-9a97-97097afda3c1@e4fe2f87-0e56-46dd-b3b8-e9af32309e82")
				getDraftContractTermsDocumentOptionsModel.DocumentID = core.StringPtr("testString")
				getDraftContractTermsDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = dpxService.GetDraftContractTermsDocument(getDraftContractTermsDocumentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetDraftContractTermsDocument with error: Operation validation and request error`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Construct an instance of the GetDraftContractTermsDocumentOptions model
				getDraftContractTermsDocumentOptionsModel := new(dpxv1.GetDraftContractTermsDocumentOptions)
				getDraftContractTermsDocumentOptionsModel.DataProductID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")
				getDraftContractTermsDocumentOptionsModel.DraftID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				getDraftContractTermsDocumentOptionsModel.ContractTermsID = core.StringPtr("598183cd-b910-4e8d-9a97-97097afda3c1@e4fe2f87-0e56-46dd-b3b8-e9af32309e82")
				getDraftContractTermsDocumentOptionsModel.DocumentID = core.StringPtr("testString")
				getDraftContractTermsDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dpxService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := dpxService.GetDraftContractTermsDocument(getDraftContractTermsDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetDraftContractTermsDocumentOptions model with no property values
				getDraftContractTermsDocumentOptionsModelNew := new(dpxv1.GetDraftContractTermsDocumentOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = dpxService.GetDraftContractTermsDocument(getDraftContractTermsDocumentOptionsModelNew)
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
			It(`Invoke GetDraftContractTermsDocument successfully`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Construct an instance of the GetDraftContractTermsDocumentOptions model
				getDraftContractTermsDocumentOptionsModel := new(dpxv1.GetDraftContractTermsDocumentOptions)
				getDraftContractTermsDocumentOptionsModel.DataProductID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")
				getDraftContractTermsDocumentOptionsModel.DraftID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				getDraftContractTermsDocumentOptionsModel.ContractTermsID = core.StringPtr("598183cd-b910-4e8d-9a97-97097afda3c1@e4fe2f87-0e56-46dd-b3b8-e9af32309e82")
				getDraftContractTermsDocumentOptionsModel.DocumentID = core.StringPtr("testString")
				getDraftContractTermsDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := dpxService.GetDraftContractTermsDocument(getDraftContractTermsDocumentOptionsModel)
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
	Describe(`DeleteDraftContractTermsDocument(deleteDraftContractTermsDocumentOptions *DeleteDraftContractTermsDocumentOptions)`, func() {
		deleteDraftContractTermsDocumentPath := "/data_product_exchange/v1/data_products/b38df608-d34b-4d58-8136-ed25e6c6684e/drafts/2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd/contract_terms/598183cd-b910-4e8d-9a97-97097afda3c1@e4fe2f87-0e56-46dd-b3b8-e9af32309e82/documents/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteDraftContractTermsDocumentPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteDraftContractTermsDocument successfully`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := dpxService.DeleteDraftContractTermsDocument(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteDraftContractTermsDocumentOptions model
				deleteDraftContractTermsDocumentOptionsModel := new(dpxv1.DeleteDraftContractTermsDocumentOptions)
				deleteDraftContractTermsDocumentOptionsModel.DataProductID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")
				deleteDraftContractTermsDocumentOptionsModel.DraftID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				deleteDraftContractTermsDocumentOptionsModel.ContractTermsID = core.StringPtr("598183cd-b910-4e8d-9a97-97097afda3c1@e4fe2f87-0e56-46dd-b3b8-e9af32309e82")
				deleteDraftContractTermsDocumentOptionsModel.DocumentID = core.StringPtr("testString")
				deleteDraftContractTermsDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = dpxService.DeleteDraftContractTermsDocument(deleteDraftContractTermsDocumentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteDraftContractTermsDocument with error: Operation validation and request error`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Construct an instance of the DeleteDraftContractTermsDocumentOptions model
				deleteDraftContractTermsDocumentOptionsModel := new(dpxv1.DeleteDraftContractTermsDocumentOptions)
				deleteDraftContractTermsDocumentOptionsModel.DataProductID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")
				deleteDraftContractTermsDocumentOptionsModel.DraftID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				deleteDraftContractTermsDocumentOptionsModel.ContractTermsID = core.StringPtr("598183cd-b910-4e8d-9a97-97097afda3c1@e4fe2f87-0e56-46dd-b3b8-e9af32309e82")
				deleteDraftContractTermsDocumentOptionsModel.DocumentID = core.StringPtr("testString")
				deleteDraftContractTermsDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dpxService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := dpxService.DeleteDraftContractTermsDocument(deleteDraftContractTermsDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteDraftContractTermsDocumentOptions model with no property values
				deleteDraftContractTermsDocumentOptionsModelNew := new(dpxv1.DeleteDraftContractTermsDocumentOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = dpxService.DeleteDraftContractTermsDocument(deleteDraftContractTermsDocumentOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateDraftContractTermsDocument(updateDraftContractTermsDocumentOptions *UpdateDraftContractTermsDocumentOptions) - Operation response error`, func() {
		updateDraftContractTermsDocumentPath := "/data_product_exchange/v1/data_products/b38df608-d34b-4d58-8136-ed25e6c6684e/drafts/2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd/contract_terms/598183cd-b910-4e8d-9a97-97097afda3c1@e4fe2f87-0e56-46dd-b3b8-e9af32309e82/documents/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateDraftContractTermsDocumentPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateDraftContractTermsDocument with error: Operation response processing error`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(dpxv1.JSONPatchOperation)
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = "testString"

				// Construct an instance of the UpdateDraftContractTermsDocumentOptions model
				updateDraftContractTermsDocumentOptionsModel := new(dpxv1.UpdateDraftContractTermsDocumentOptions)
				updateDraftContractTermsDocumentOptionsModel.DataProductID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")
				updateDraftContractTermsDocumentOptionsModel.DraftID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				updateDraftContractTermsDocumentOptionsModel.ContractTermsID = core.StringPtr("598183cd-b910-4e8d-9a97-97097afda3c1@e4fe2f87-0e56-46dd-b3b8-e9af32309e82")
				updateDraftContractTermsDocumentOptionsModel.DocumentID = core.StringPtr("testString")
				updateDraftContractTermsDocumentOptionsModel.JSONPatchInstructions = []dpxv1.JSONPatchOperation{*jsonPatchOperationModel}
				updateDraftContractTermsDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := dpxService.UpdateDraftContractTermsDocument(updateDraftContractTermsDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				dpxService.EnableRetries(0, 0)
				result, response, operationErr = dpxService.UpdateDraftContractTermsDocument(updateDraftContractTermsDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateDraftContractTermsDocument(updateDraftContractTermsDocumentOptions *UpdateDraftContractTermsDocumentOptions)`, func() {
		updateDraftContractTermsDocumentPath := "/data_product_exchange/v1/data_products/b38df608-d34b-4d58-8136-ed25e6c6684e/drafts/2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd/contract_terms/598183cd-b910-4e8d-9a97-97097afda3c1@e4fe2f87-0e56-46dd-b3b8-e9af32309e82/documents/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateDraftContractTermsDocumentPath))
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
					fmt.Fprintf(res, "%s", `{"url": "URL", "type": "terms_and_conditions", "name": "Name", "id": "2b0bf220-079c-11ee-be56-0242ac120002", "attachment": {"id": "ID"}, "upload_url": "UploadURL"}`)
				}))
			})
			It(`Invoke UpdateDraftContractTermsDocument successfully with retries`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())
				dpxService.EnableRetries(0, 0)

				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(dpxv1.JSONPatchOperation)
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = "testString"

				// Construct an instance of the UpdateDraftContractTermsDocumentOptions model
				updateDraftContractTermsDocumentOptionsModel := new(dpxv1.UpdateDraftContractTermsDocumentOptions)
				updateDraftContractTermsDocumentOptionsModel.DataProductID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")
				updateDraftContractTermsDocumentOptionsModel.DraftID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				updateDraftContractTermsDocumentOptionsModel.ContractTermsID = core.StringPtr("598183cd-b910-4e8d-9a97-97097afda3c1@e4fe2f87-0e56-46dd-b3b8-e9af32309e82")
				updateDraftContractTermsDocumentOptionsModel.DocumentID = core.StringPtr("testString")
				updateDraftContractTermsDocumentOptionsModel.JSONPatchInstructions = []dpxv1.JSONPatchOperation{*jsonPatchOperationModel}
				updateDraftContractTermsDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := dpxService.UpdateDraftContractTermsDocumentWithContext(ctx, updateDraftContractTermsDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				dpxService.DisableRetries()
				result, response, operationErr := dpxService.UpdateDraftContractTermsDocument(updateDraftContractTermsDocumentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = dpxService.UpdateDraftContractTermsDocumentWithContext(ctx, updateDraftContractTermsDocumentOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateDraftContractTermsDocumentPath))
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
					fmt.Fprintf(res, "%s", `{"url": "URL", "type": "terms_and_conditions", "name": "Name", "id": "2b0bf220-079c-11ee-be56-0242ac120002", "attachment": {"id": "ID"}, "upload_url": "UploadURL"}`)
				}))
			})
			It(`Invoke UpdateDraftContractTermsDocument successfully`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := dpxService.UpdateDraftContractTermsDocument(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(dpxv1.JSONPatchOperation)
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = "testString"

				// Construct an instance of the UpdateDraftContractTermsDocumentOptions model
				updateDraftContractTermsDocumentOptionsModel := new(dpxv1.UpdateDraftContractTermsDocumentOptions)
				updateDraftContractTermsDocumentOptionsModel.DataProductID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")
				updateDraftContractTermsDocumentOptionsModel.DraftID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				updateDraftContractTermsDocumentOptionsModel.ContractTermsID = core.StringPtr("598183cd-b910-4e8d-9a97-97097afda3c1@e4fe2f87-0e56-46dd-b3b8-e9af32309e82")
				updateDraftContractTermsDocumentOptionsModel.DocumentID = core.StringPtr("testString")
				updateDraftContractTermsDocumentOptionsModel.JSONPatchInstructions = []dpxv1.JSONPatchOperation{*jsonPatchOperationModel}
				updateDraftContractTermsDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = dpxService.UpdateDraftContractTermsDocument(updateDraftContractTermsDocumentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateDraftContractTermsDocument with error: Operation validation and request error`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(dpxv1.JSONPatchOperation)
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = "testString"

				// Construct an instance of the UpdateDraftContractTermsDocumentOptions model
				updateDraftContractTermsDocumentOptionsModel := new(dpxv1.UpdateDraftContractTermsDocumentOptions)
				updateDraftContractTermsDocumentOptionsModel.DataProductID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")
				updateDraftContractTermsDocumentOptionsModel.DraftID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				updateDraftContractTermsDocumentOptionsModel.ContractTermsID = core.StringPtr("598183cd-b910-4e8d-9a97-97097afda3c1@e4fe2f87-0e56-46dd-b3b8-e9af32309e82")
				updateDraftContractTermsDocumentOptionsModel.DocumentID = core.StringPtr("testString")
				updateDraftContractTermsDocumentOptionsModel.JSONPatchInstructions = []dpxv1.JSONPatchOperation{*jsonPatchOperationModel}
				updateDraftContractTermsDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dpxService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := dpxService.UpdateDraftContractTermsDocument(updateDraftContractTermsDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateDraftContractTermsDocumentOptions model with no property values
				updateDraftContractTermsDocumentOptionsModelNew := new(dpxv1.UpdateDraftContractTermsDocumentOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = dpxService.UpdateDraftContractTermsDocument(updateDraftContractTermsDocumentOptionsModelNew)
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
			It(`Invoke UpdateDraftContractTermsDocument successfully`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(dpxv1.JSONPatchOperation)
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = "testString"

				// Construct an instance of the UpdateDraftContractTermsDocumentOptions model
				updateDraftContractTermsDocumentOptionsModel := new(dpxv1.UpdateDraftContractTermsDocumentOptions)
				updateDraftContractTermsDocumentOptionsModel.DataProductID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")
				updateDraftContractTermsDocumentOptionsModel.DraftID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				updateDraftContractTermsDocumentOptionsModel.ContractTermsID = core.StringPtr("598183cd-b910-4e8d-9a97-97097afda3c1@e4fe2f87-0e56-46dd-b3b8-e9af32309e82")
				updateDraftContractTermsDocumentOptionsModel.DocumentID = core.StringPtr("testString")
				updateDraftContractTermsDocumentOptionsModel.JSONPatchInstructions = []dpxv1.JSONPatchOperation{*jsonPatchOperationModel}
				updateDraftContractTermsDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := dpxService.UpdateDraftContractTermsDocument(updateDraftContractTermsDocumentOptionsModel)
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
	Describe(`PublishDataProductDraft(publishDataProductDraftOptions *PublishDataProductDraftOptions) - Operation response error`, func() {
		publishDataProductDraftPath := "/data_product_exchange/v1/data_products/b38df608-d34b-4d58-8136-ed25e6c6684e/drafts/2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd/publish"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(publishDataProductDraftPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke PublishDataProductDraft with error: Operation response processing error`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Construct an instance of the PublishDataProductDraftOptions model
				publishDataProductDraftOptionsModel := new(dpxv1.PublishDataProductDraftOptions)
				publishDataProductDraftOptionsModel.DataProductID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")
				publishDataProductDraftOptionsModel.DraftID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				publishDataProductDraftOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := dpxService.PublishDataProductDraft(publishDataProductDraftOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				dpxService.EnableRetries(0, 0)
				result, response, operationErr = dpxService.PublishDataProductDraft(publishDataProductDraftOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`PublishDataProductDraft(publishDataProductDraftOptions *PublishDataProductDraftOptions)`, func() {
		publishDataProductDraftPath := "/data_product_exchange/v1/data_products/b38df608-d34b-4d58-8136-ed25e6c6684e/drafts/2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd/publish"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(publishDataProductDraftPath))
					Expect(req.Method).To(Equal("POST"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"version": "1.0.0", "state": "draft", "data_product": {"id": "b38df608-d34b-4d58-8136-ed25e6c6684e"}, "name": "My Data Product", "description": "This is a description of My Data Product.", "id": "2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd", "asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "tags": ["Tags"], "use_cases": [{"id": "ID", "name": "Name", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}], "domain": {"id": "ID", "name": "Name", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "types": ["data"], "parts_out": [{"asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}, "type": "data_asset"}, "revision": 1, "updated_at": "2023-07-01T22:22:34.876Z", "delivery_methods": [{"id": "09cf5fcc-cb9d-4995-a8e4-16517b25229f", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}]}], "published_by": "PublishedBy", "published_at": "2019-01-01T12:00:00.000Z", "contract_terms": [{"asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "id": "ID", "documents": [{"url": "URL", "type": "terms_and_conditions", "name": "Name", "id": "2b0bf220-079c-11ee-be56-0242ac120002", "attachment": {"id": "ID"}, "upload_url": "UploadURL"}]}], "created_by": "CreatedBy", "created_at": "2019-01-01T12:00:00.000Z", "is_restricted": true}`)
				}))
			})
			It(`Invoke PublishDataProductDraft successfully with retries`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())
				dpxService.EnableRetries(0, 0)

				// Construct an instance of the PublishDataProductDraftOptions model
				publishDataProductDraftOptionsModel := new(dpxv1.PublishDataProductDraftOptions)
				publishDataProductDraftOptionsModel.DataProductID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")
				publishDataProductDraftOptionsModel.DraftID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				publishDataProductDraftOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := dpxService.PublishDataProductDraftWithContext(ctx, publishDataProductDraftOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				dpxService.DisableRetries()
				result, response, operationErr := dpxService.PublishDataProductDraft(publishDataProductDraftOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = dpxService.PublishDataProductDraftWithContext(ctx, publishDataProductDraftOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(publishDataProductDraftPath))
					Expect(req.Method).To(Equal("POST"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"version": "1.0.0", "state": "draft", "data_product": {"id": "b38df608-d34b-4d58-8136-ed25e6c6684e"}, "name": "My Data Product", "description": "This is a description of My Data Product.", "id": "2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd", "asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "tags": ["Tags"], "use_cases": [{"id": "ID", "name": "Name", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}], "domain": {"id": "ID", "name": "Name", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "types": ["data"], "parts_out": [{"asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}, "type": "data_asset"}, "revision": 1, "updated_at": "2023-07-01T22:22:34.876Z", "delivery_methods": [{"id": "09cf5fcc-cb9d-4995-a8e4-16517b25229f", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}]}], "published_by": "PublishedBy", "published_at": "2019-01-01T12:00:00.000Z", "contract_terms": [{"asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "id": "ID", "documents": [{"url": "URL", "type": "terms_and_conditions", "name": "Name", "id": "2b0bf220-079c-11ee-be56-0242ac120002", "attachment": {"id": "ID"}, "upload_url": "UploadURL"}]}], "created_by": "CreatedBy", "created_at": "2019-01-01T12:00:00.000Z", "is_restricted": true}`)
				}))
			})
			It(`Invoke PublishDataProductDraft successfully`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := dpxService.PublishDataProductDraft(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the PublishDataProductDraftOptions model
				publishDataProductDraftOptionsModel := new(dpxv1.PublishDataProductDraftOptions)
				publishDataProductDraftOptionsModel.DataProductID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")
				publishDataProductDraftOptionsModel.DraftID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				publishDataProductDraftOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = dpxService.PublishDataProductDraft(publishDataProductDraftOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke PublishDataProductDraft with error: Operation validation and request error`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Construct an instance of the PublishDataProductDraftOptions model
				publishDataProductDraftOptionsModel := new(dpxv1.PublishDataProductDraftOptions)
				publishDataProductDraftOptionsModel.DataProductID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")
				publishDataProductDraftOptionsModel.DraftID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				publishDataProductDraftOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dpxService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := dpxService.PublishDataProductDraft(publishDataProductDraftOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the PublishDataProductDraftOptions model with no property values
				publishDataProductDraftOptionsModelNew := new(dpxv1.PublishDataProductDraftOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = dpxService.PublishDataProductDraft(publishDataProductDraftOptionsModelNew)
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
			It(`Invoke PublishDataProductDraft successfully`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Construct an instance of the PublishDataProductDraftOptions model
				publishDataProductDraftOptionsModel := new(dpxv1.PublishDataProductDraftOptions)
				publishDataProductDraftOptionsModel.DataProductID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")
				publishDataProductDraftOptionsModel.DraftID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				publishDataProductDraftOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := dpxService.PublishDataProductDraft(publishDataProductDraftOptionsModel)
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
	Describe(`GetDataProductRelease(getDataProductReleaseOptions *GetDataProductReleaseOptions) - Operation response error`, func() {
		getDataProductReleasePath := "/data_product_exchange/v1/data_products/b38df608-d34b-4d58-8136-ed25e6c6684e/releases/2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDataProductReleasePath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetDataProductRelease with error: Operation response processing error`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Construct an instance of the GetDataProductReleaseOptions model
				getDataProductReleaseOptionsModel := new(dpxv1.GetDataProductReleaseOptions)
				getDataProductReleaseOptionsModel.DataProductID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")
				getDataProductReleaseOptionsModel.ReleaseID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				getDataProductReleaseOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := dpxService.GetDataProductRelease(getDataProductReleaseOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				dpxService.EnableRetries(0, 0)
				result, response, operationErr = dpxService.GetDataProductRelease(getDataProductReleaseOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetDataProductRelease(getDataProductReleaseOptions *GetDataProductReleaseOptions)`, func() {
		getDataProductReleasePath := "/data_product_exchange/v1/data_products/b38df608-d34b-4d58-8136-ed25e6c6684e/releases/2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDataProductReleasePath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"version": "1.0.0", "state": "draft", "data_product": {"id": "b38df608-d34b-4d58-8136-ed25e6c6684e"}, "name": "My Data Product", "description": "This is a description of My Data Product.", "id": "2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd", "asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "tags": ["Tags"], "use_cases": [{"id": "ID", "name": "Name", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}], "domain": {"id": "ID", "name": "Name", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "types": ["data"], "parts_out": [{"asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}, "type": "data_asset"}, "revision": 1, "updated_at": "2023-07-01T22:22:34.876Z", "delivery_methods": [{"id": "09cf5fcc-cb9d-4995-a8e4-16517b25229f", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}]}], "published_by": "PublishedBy", "published_at": "2019-01-01T12:00:00.000Z", "contract_terms": [{"asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "id": "ID", "documents": [{"url": "URL", "type": "terms_and_conditions", "name": "Name", "id": "2b0bf220-079c-11ee-be56-0242ac120002", "attachment": {"id": "ID"}, "upload_url": "UploadURL"}]}], "created_by": "CreatedBy", "created_at": "2019-01-01T12:00:00.000Z", "is_restricted": true}`)
				}))
			})
			It(`Invoke GetDataProductRelease successfully with retries`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())
				dpxService.EnableRetries(0, 0)

				// Construct an instance of the GetDataProductReleaseOptions model
				getDataProductReleaseOptionsModel := new(dpxv1.GetDataProductReleaseOptions)
				getDataProductReleaseOptionsModel.DataProductID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")
				getDataProductReleaseOptionsModel.ReleaseID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				getDataProductReleaseOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := dpxService.GetDataProductReleaseWithContext(ctx, getDataProductReleaseOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				dpxService.DisableRetries()
				result, response, operationErr := dpxService.GetDataProductRelease(getDataProductReleaseOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = dpxService.GetDataProductReleaseWithContext(ctx, getDataProductReleaseOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getDataProductReleasePath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"version": "1.0.0", "state": "draft", "data_product": {"id": "b38df608-d34b-4d58-8136-ed25e6c6684e"}, "name": "My Data Product", "description": "This is a description of My Data Product.", "id": "2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd", "asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "tags": ["Tags"], "use_cases": [{"id": "ID", "name": "Name", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}], "domain": {"id": "ID", "name": "Name", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "types": ["data"], "parts_out": [{"asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}, "type": "data_asset"}, "revision": 1, "updated_at": "2023-07-01T22:22:34.876Z", "delivery_methods": [{"id": "09cf5fcc-cb9d-4995-a8e4-16517b25229f", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}]}], "published_by": "PublishedBy", "published_at": "2019-01-01T12:00:00.000Z", "contract_terms": [{"asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "id": "ID", "documents": [{"url": "URL", "type": "terms_and_conditions", "name": "Name", "id": "2b0bf220-079c-11ee-be56-0242ac120002", "attachment": {"id": "ID"}, "upload_url": "UploadURL"}]}], "created_by": "CreatedBy", "created_at": "2019-01-01T12:00:00.000Z", "is_restricted": true}`)
				}))
			})
			It(`Invoke GetDataProductRelease successfully`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := dpxService.GetDataProductRelease(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetDataProductReleaseOptions model
				getDataProductReleaseOptionsModel := new(dpxv1.GetDataProductReleaseOptions)
				getDataProductReleaseOptionsModel.DataProductID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")
				getDataProductReleaseOptionsModel.ReleaseID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				getDataProductReleaseOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = dpxService.GetDataProductRelease(getDataProductReleaseOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetDataProductRelease with error: Operation validation and request error`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Construct an instance of the GetDataProductReleaseOptions model
				getDataProductReleaseOptionsModel := new(dpxv1.GetDataProductReleaseOptions)
				getDataProductReleaseOptionsModel.DataProductID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")
				getDataProductReleaseOptionsModel.ReleaseID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				getDataProductReleaseOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dpxService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := dpxService.GetDataProductRelease(getDataProductReleaseOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetDataProductReleaseOptions model with no property values
				getDataProductReleaseOptionsModelNew := new(dpxv1.GetDataProductReleaseOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = dpxService.GetDataProductRelease(getDataProductReleaseOptionsModelNew)
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
			It(`Invoke GetDataProductRelease successfully`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Construct an instance of the GetDataProductReleaseOptions model
				getDataProductReleaseOptionsModel := new(dpxv1.GetDataProductReleaseOptions)
				getDataProductReleaseOptionsModel.DataProductID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")
				getDataProductReleaseOptionsModel.ReleaseID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				getDataProductReleaseOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := dpxService.GetDataProductRelease(getDataProductReleaseOptionsModel)
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
	Describe(`UpdateDataProductRelease(updateDataProductReleaseOptions *UpdateDataProductReleaseOptions) - Operation response error`, func() {
		updateDataProductReleasePath := "/data_product_exchange/v1/data_products/b38df608-d34b-4d58-8136-ed25e6c6684e/releases/2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateDataProductReleasePath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateDataProductRelease with error: Operation response processing error`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(dpxv1.JSONPatchOperation)
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = "testString"

				// Construct an instance of the UpdateDataProductReleaseOptions model
				updateDataProductReleaseOptionsModel := new(dpxv1.UpdateDataProductReleaseOptions)
				updateDataProductReleaseOptionsModel.DataProductID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")
				updateDataProductReleaseOptionsModel.ReleaseID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				updateDataProductReleaseOptionsModel.JSONPatchInstructions = []dpxv1.JSONPatchOperation{*jsonPatchOperationModel}
				updateDataProductReleaseOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := dpxService.UpdateDataProductRelease(updateDataProductReleaseOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				dpxService.EnableRetries(0, 0)
				result, response, operationErr = dpxService.UpdateDataProductRelease(updateDataProductReleaseOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateDataProductRelease(updateDataProductReleaseOptions *UpdateDataProductReleaseOptions)`, func() {
		updateDataProductReleasePath := "/data_product_exchange/v1/data_products/b38df608-d34b-4d58-8136-ed25e6c6684e/releases/2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateDataProductReleasePath))
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
					fmt.Fprintf(res, "%s", `{"version": "1.0.0", "state": "draft", "data_product": {"id": "b38df608-d34b-4d58-8136-ed25e6c6684e"}, "name": "My Data Product", "description": "This is a description of My Data Product.", "id": "2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd", "asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "tags": ["Tags"], "use_cases": [{"id": "ID", "name": "Name", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}], "domain": {"id": "ID", "name": "Name", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "types": ["data"], "parts_out": [{"asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}, "type": "data_asset"}, "revision": 1, "updated_at": "2023-07-01T22:22:34.876Z", "delivery_methods": [{"id": "09cf5fcc-cb9d-4995-a8e4-16517b25229f", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}]}], "published_by": "PublishedBy", "published_at": "2019-01-01T12:00:00.000Z", "contract_terms": [{"asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "id": "ID", "documents": [{"url": "URL", "type": "terms_and_conditions", "name": "Name", "id": "2b0bf220-079c-11ee-be56-0242ac120002", "attachment": {"id": "ID"}, "upload_url": "UploadURL"}]}], "created_by": "CreatedBy", "created_at": "2019-01-01T12:00:00.000Z", "is_restricted": true}`)
				}))
			})
			It(`Invoke UpdateDataProductRelease successfully with retries`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())
				dpxService.EnableRetries(0, 0)

				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(dpxv1.JSONPatchOperation)
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = "testString"

				// Construct an instance of the UpdateDataProductReleaseOptions model
				updateDataProductReleaseOptionsModel := new(dpxv1.UpdateDataProductReleaseOptions)
				updateDataProductReleaseOptionsModel.DataProductID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")
				updateDataProductReleaseOptionsModel.ReleaseID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				updateDataProductReleaseOptionsModel.JSONPatchInstructions = []dpxv1.JSONPatchOperation{*jsonPatchOperationModel}
				updateDataProductReleaseOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := dpxService.UpdateDataProductReleaseWithContext(ctx, updateDataProductReleaseOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				dpxService.DisableRetries()
				result, response, operationErr := dpxService.UpdateDataProductRelease(updateDataProductReleaseOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = dpxService.UpdateDataProductReleaseWithContext(ctx, updateDataProductReleaseOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateDataProductReleasePath))
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
					fmt.Fprintf(res, "%s", `{"version": "1.0.0", "state": "draft", "data_product": {"id": "b38df608-d34b-4d58-8136-ed25e6c6684e"}, "name": "My Data Product", "description": "This is a description of My Data Product.", "id": "2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd", "asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "tags": ["Tags"], "use_cases": [{"id": "ID", "name": "Name", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}], "domain": {"id": "ID", "name": "Name", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "types": ["data"], "parts_out": [{"asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}, "type": "data_asset"}, "revision": 1, "updated_at": "2023-07-01T22:22:34.876Z", "delivery_methods": [{"id": "09cf5fcc-cb9d-4995-a8e4-16517b25229f", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}]}], "published_by": "PublishedBy", "published_at": "2019-01-01T12:00:00.000Z", "contract_terms": [{"asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "id": "ID", "documents": [{"url": "URL", "type": "terms_and_conditions", "name": "Name", "id": "2b0bf220-079c-11ee-be56-0242ac120002", "attachment": {"id": "ID"}, "upload_url": "UploadURL"}]}], "created_by": "CreatedBy", "created_at": "2019-01-01T12:00:00.000Z", "is_restricted": true}`)
				}))
			})
			It(`Invoke UpdateDataProductRelease successfully`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := dpxService.UpdateDataProductRelease(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(dpxv1.JSONPatchOperation)
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = "testString"

				// Construct an instance of the UpdateDataProductReleaseOptions model
				updateDataProductReleaseOptionsModel := new(dpxv1.UpdateDataProductReleaseOptions)
				updateDataProductReleaseOptionsModel.DataProductID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")
				updateDataProductReleaseOptionsModel.ReleaseID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				updateDataProductReleaseOptionsModel.JSONPatchInstructions = []dpxv1.JSONPatchOperation{*jsonPatchOperationModel}
				updateDataProductReleaseOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = dpxService.UpdateDataProductRelease(updateDataProductReleaseOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateDataProductRelease with error: Operation validation and request error`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(dpxv1.JSONPatchOperation)
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = "testString"

				// Construct an instance of the UpdateDataProductReleaseOptions model
				updateDataProductReleaseOptionsModel := new(dpxv1.UpdateDataProductReleaseOptions)
				updateDataProductReleaseOptionsModel.DataProductID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")
				updateDataProductReleaseOptionsModel.ReleaseID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				updateDataProductReleaseOptionsModel.JSONPatchInstructions = []dpxv1.JSONPatchOperation{*jsonPatchOperationModel}
				updateDataProductReleaseOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dpxService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := dpxService.UpdateDataProductRelease(updateDataProductReleaseOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateDataProductReleaseOptions model with no property values
				updateDataProductReleaseOptionsModelNew := new(dpxv1.UpdateDataProductReleaseOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = dpxService.UpdateDataProductRelease(updateDataProductReleaseOptionsModelNew)
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
			It(`Invoke UpdateDataProductRelease successfully`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(dpxv1.JSONPatchOperation)
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = "testString"

				// Construct an instance of the UpdateDataProductReleaseOptions model
				updateDataProductReleaseOptionsModel := new(dpxv1.UpdateDataProductReleaseOptions)
				updateDataProductReleaseOptionsModel.DataProductID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")
				updateDataProductReleaseOptionsModel.ReleaseID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				updateDataProductReleaseOptionsModel.JSONPatchInstructions = []dpxv1.JSONPatchOperation{*jsonPatchOperationModel}
				updateDataProductReleaseOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := dpxService.UpdateDataProductRelease(updateDataProductReleaseOptionsModel)
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
	Describe(`GetReleaseContractTermsDocument(getReleaseContractTermsDocumentOptions *GetReleaseContractTermsDocumentOptions) - Operation response error`, func() {
		getReleaseContractTermsDocumentPath := "/data_product_exchange/v1/data_products/b38df608-d34b-4d58-8136-ed25e6c6684e/releases/2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd/contract_terms/598183cd-b910-4e8d-9a97-97097afda3c1@e4fe2f87-0e56-46dd-b3b8-e9af32309e82/documents/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getReleaseContractTermsDocumentPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetReleaseContractTermsDocument with error: Operation response processing error`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Construct an instance of the GetReleaseContractTermsDocumentOptions model
				getReleaseContractTermsDocumentOptionsModel := new(dpxv1.GetReleaseContractTermsDocumentOptions)
				getReleaseContractTermsDocumentOptionsModel.DataProductID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")
				getReleaseContractTermsDocumentOptionsModel.ReleaseID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				getReleaseContractTermsDocumentOptionsModel.ContractTermsID = core.StringPtr("598183cd-b910-4e8d-9a97-97097afda3c1@e4fe2f87-0e56-46dd-b3b8-e9af32309e82")
				getReleaseContractTermsDocumentOptionsModel.DocumentID = core.StringPtr("testString")
				getReleaseContractTermsDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := dpxService.GetReleaseContractTermsDocument(getReleaseContractTermsDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				dpxService.EnableRetries(0, 0)
				result, response, operationErr = dpxService.GetReleaseContractTermsDocument(getReleaseContractTermsDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetReleaseContractTermsDocument(getReleaseContractTermsDocumentOptions *GetReleaseContractTermsDocumentOptions)`, func() {
		getReleaseContractTermsDocumentPath := "/data_product_exchange/v1/data_products/b38df608-d34b-4d58-8136-ed25e6c6684e/releases/2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd/contract_terms/598183cd-b910-4e8d-9a97-97097afda3c1@e4fe2f87-0e56-46dd-b3b8-e9af32309e82/documents/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getReleaseContractTermsDocumentPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"url": "URL", "type": "terms_and_conditions", "name": "Name", "id": "2b0bf220-079c-11ee-be56-0242ac120002", "attachment": {"id": "ID"}, "upload_url": "UploadURL"}`)
				}))
			})
			It(`Invoke GetReleaseContractTermsDocument successfully with retries`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())
				dpxService.EnableRetries(0, 0)

				// Construct an instance of the GetReleaseContractTermsDocumentOptions model
				getReleaseContractTermsDocumentOptionsModel := new(dpxv1.GetReleaseContractTermsDocumentOptions)
				getReleaseContractTermsDocumentOptionsModel.DataProductID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")
				getReleaseContractTermsDocumentOptionsModel.ReleaseID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				getReleaseContractTermsDocumentOptionsModel.ContractTermsID = core.StringPtr("598183cd-b910-4e8d-9a97-97097afda3c1@e4fe2f87-0e56-46dd-b3b8-e9af32309e82")
				getReleaseContractTermsDocumentOptionsModel.DocumentID = core.StringPtr("testString")
				getReleaseContractTermsDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := dpxService.GetReleaseContractTermsDocumentWithContext(ctx, getReleaseContractTermsDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				dpxService.DisableRetries()
				result, response, operationErr := dpxService.GetReleaseContractTermsDocument(getReleaseContractTermsDocumentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = dpxService.GetReleaseContractTermsDocumentWithContext(ctx, getReleaseContractTermsDocumentOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getReleaseContractTermsDocumentPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"url": "URL", "type": "terms_and_conditions", "name": "Name", "id": "2b0bf220-079c-11ee-be56-0242ac120002", "attachment": {"id": "ID"}, "upload_url": "UploadURL"}`)
				}))
			})
			It(`Invoke GetReleaseContractTermsDocument successfully`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := dpxService.GetReleaseContractTermsDocument(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetReleaseContractTermsDocumentOptions model
				getReleaseContractTermsDocumentOptionsModel := new(dpxv1.GetReleaseContractTermsDocumentOptions)
				getReleaseContractTermsDocumentOptionsModel.DataProductID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")
				getReleaseContractTermsDocumentOptionsModel.ReleaseID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				getReleaseContractTermsDocumentOptionsModel.ContractTermsID = core.StringPtr("598183cd-b910-4e8d-9a97-97097afda3c1@e4fe2f87-0e56-46dd-b3b8-e9af32309e82")
				getReleaseContractTermsDocumentOptionsModel.DocumentID = core.StringPtr("testString")
				getReleaseContractTermsDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = dpxService.GetReleaseContractTermsDocument(getReleaseContractTermsDocumentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetReleaseContractTermsDocument with error: Operation validation and request error`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Construct an instance of the GetReleaseContractTermsDocumentOptions model
				getReleaseContractTermsDocumentOptionsModel := new(dpxv1.GetReleaseContractTermsDocumentOptions)
				getReleaseContractTermsDocumentOptionsModel.DataProductID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")
				getReleaseContractTermsDocumentOptionsModel.ReleaseID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				getReleaseContractTermsDocumentOptionsModel.ContractTermsID = core.StringPtr("598183cd-b910-4e8d-9a97-97097afda3c1@e4fe2f87-0e56-46dd-b3b8-e9af32309e82")
				getReleaseContractTermsDocumentOptionsModel.DocumentID = core.StringPtr("testString")
				getReleaseContractTermsDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dpxService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := dpxService.GetReleaseContractTermsDocument(getReleaseContractTermsDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetReleaseContractTermsDocumentOptions model with no property values
				getReleaseContractTermsDocumentOptionsModelNew := new(dpxv1.GetReleaseContractTermsDocumentOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = dpxService.GetReleaseContractTermsDocument(getReleaseContractTermsDocumentOptionsModelNew)
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
			It(`Invoke GetReleaseContractTermsDocument successfully`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Construct an instance of the GetReleaseContractTermsDocumentOptions model
				getReleaseContractTermsDocumentOptionsModel := new(dpxv1.GetReleaseContractTermsDocumentOptions)
				getReleaseContractTermsDocumentOptionsModel.DataProductID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")
				getReleaseContractTermsDocumentOptionsModel.ReleaseID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				getReleaseContractTermsDocumentOptionsModel.ContractTermsID = core.StringPtr("598183cd-b910-4e8d-9a97-97097afda3c1@e4fe2f87-0e56-46dd-b3b8-e9af32309e82")
				getReleaseContractTermsDocumentOptionsModel.DocumentID = core.StringPtr("testString")
				getReleaseContractTermsDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := dpxService.GetReleaseContractTermsDocument(getReleaseContractTermsDocumentOptionsModel)
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
	Describe(`ListDataProductReleases(listDataProductReleasesOptions *ListDataProductReleasesOptions) - Operation response error`, func() {
		listDataProductReleasesPath := "/data_product_exchange/v1/data_products/b38df608-d34b-4d58-8136-ed25e6c6684e/releases"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listDataProductReleasesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["asset.container.id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListDataProductReleases with error: Operation response processing error`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Construct an instance of the ListDataProductReleasesOptions model
				listDataProductReleasesOptionsModel := new(dpxv1.ListDataProductReleasesOptions)
				listDataProductReleasesOptionsModel.DataProductID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")
				listDataProductReleasesOptionsModel.AssetContainerID = core.StringPtr("testString")
				listDataProductReleasesOptionsModel.State = []string{"available"}
				listDataProductReleasesOptionsModel.Version = core.StringPtr("testString")
				listDataProductReleasesOptionsModel.Limit = core.Int64Ptr(int64(10))
				listDataProductReleasesOptionsModel.Start = core.StringPtr("testString")
				listDataProductReleasesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := dpxService.ListDataProductReleases(listDataProductReleasesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				dpxService.EnableRetries(0, 0)
				result, response, operationErr = dpxService.ListDataProductReleases(listDataProductReleasesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListDataProductReleases(listDataProductReleasesOptions *ListDataProductReleasesOptions)`, func() {
		listDataProductReleasesPath := "/data_product_exchange/v1/data_products/b38df608-d34b-4d58-8136-ed25e6c6684e/releases"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listDataProductReleasesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["asset.container.id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"limit": 200, "first": {"href": "https://api.example.com/collection"}, "next": {"href": "https://api.example.com/collection?start=eyJvZmZzZXQiOjAsImRvbmUiOnRydWV9", "start": "eyJvZmZzZXQiOjAsImRvbmUiOnRydWV9"}, "releases": [{"version": "1.0.0", "state": "draft", "data_product": {"id": "b38df608-d34b-4d58-8136-ed25e6c6684e"}, "name": "My Data Product", "description": "This is a description of My Data Product.", "id": "2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd", "asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}}]}`)
				}))
			})
			It(`Invoke ListDataProductReleases successfully with retries`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())
				dpxService.EnableRetries(0, 0)

				// Construct an instance of the ListDataProductReleasesOptions model
				listDataProductReleasesOptionsModel := new(dpxv1.ListDataProductReleasesOptions)
				listDataProductReleasesOptionsModel.DataProductID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")
				listDataProductReleasesOptionsModel.AssetContainerID = core.StringPtr("testString")
				listDataProductReleasesOptionsModel.State = []string{"available"}
				listDataProductReleasesOptionsModel.Version = core.StringPtr("testString")
				listDataProductReleasesOptionsModel.Limit = core.Int64Ptr(int64(10))
				listDataProductReleasesOptionsModel.Start = core.StringPtr("testString")
				listDataProductReleasesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := dpxService.ListDataProductReleasesWithContext(ctx, listDataProductReleasesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				dpxService.DisableRetries()
				result, response, operationErr := dpxService.ListDataProductReleases(listDataProductReleasesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = dpxService.ListDataProductReleasesWithContext(ctx, listDataProductReleasesOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listDataProductReleasesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["asset.container.id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"limit": 200, "first": {"href": "https://api.example.com/collection"}, "next": {"href": "https://api.example.com/collection?start=eyJvZmZzZXQiOjAsImRvbmUiOnRydWV9", "start": "eyJvZmZzZXQiOjAsImRvbmUiOnRydWV9"}, "releases": [{"version": "1.0.0", "state": "draft", "data_product": {"id": "b38df608-d34b-4d58-8136-ed25e6c6684e"}, "name": "My Data Product", "description": "This is a description of My Data Product.", "id": "2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd", "asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}}]}`)
				}))
			})
			It(`Invoke ListDataProductReleases successfully`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := dpxService.ListDataProductReleases(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListDataProductReleasesOptions model
				listDataProductReleasesOptionsModel := new(dpxv1.ListDataProductReleasesOptions)
				listDataProductReleasesOptionsModel.DataProductID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")
				listDataProductReleasesOptionsModel.AssetContainerID = core.StringPtr("testString")
				listDataProductReleasesOptionsModel.State = []string{"available"}
				listDataProductReleasesOptionsModel.Version = core.StringPtr("testString")
				listDataProductReleasesOptionsModel.Limit = core.Int64Ptr(int64(10))
				listDataProductReleasesOptionsModel.Start = core.StringPtr("testString")
				listDataProductReleasesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = dpxService.ListDataProductReleases(listDataProductReleasesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListDataProductReleases with error: Operation validation and request error`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Construct an instance of the ListDataProductReleasesOptions model
				listDataProductReleasesOptionsModel := new(dpxv1.ListDataProductReleasesOptions)
				listDataProductReleasesOptionsModel.DataProductID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")
				listDataProductReleasesOptionsModel.AssetContainerID = core.StringPtr("testString")
				listDataProductReleasesOptionsModel.State = []string{"available"}
				listDataProductReleasesOptionsModel.Version = core.StringPtr("testString")
				listDataProductReleasesOptionsModel.Limit = core.Int64Ptr(int64(10))
				listDataProductReleasesOptionsModel.Start = core.StringPtr("testString")
				listDataProductReleasesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dpxService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := dpxService.ListDataProductReleases(listDataProductReleasesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListDataProductReleasesOptions model with no property values
				listDataProductReleasesOptionsModelNew := new(dpxv1.ListDataProductReleasesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = dpxService.ListDataProductReleases(listDataProductReleasesOptionsModelNew)
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
			It(`Invoke ListDataProductReleases successfully`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Construct an instance of the ListDataProductReleasesOptions model
				listDataProductReleasesOptionsModel := new(dpxv1.ListDataProductReleasesOptions)
				listDataProductReleasesOptionsModel.DataProductID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")
				listDataProductReleasesOptionsModel.AssetContainerID = core.StringPtr("testString")
				listDataProductReleasesOptionsModel.State = []string{"available"}
				listDataProductReleasesOptionsModel.Version = core.StringPtr("testString")
				listDataProductReleasesOptionsModel.Limit = core.Int64Ptr(int64(10))
				listDataProductReleasesOptionsModel.Start = core.StringPtr("testString")
				listDataProductReleasesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := dpxService.ListDataProductReleases(listDataProductReleasesOptionsModel)
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
				responseObject := new(dpxv1.DataProductReleaseCollection)
				nextObject := new(dpxv1.NextPage)
				nextObject.Start = core.StringPtr("abc-123")
				responseObject.Next = nextObject

				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(Equal(core.StringPtr("abc-123")))
			})
			It(`Invoke GetNextStart without a "Next" property in the response`, func() {
				responseObject := new(dpxv1.DataProductReleaseCollection)

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
					Expect(req.URL.EscapedPath()).To(Equal(listDataProductReleasesPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					requestNumber++
					if requestNumber == 1 {
						fmt.Fprintf(res, "%s", `{"next":{"start":"1"},"total_count":2,"limit":1,"releases":[{"version":"1.0.0","state":"draft","data_product":{"id":"b38df608-d34b-4d58-8136-ed25e6c6684e"},"name":"My Data Product","description":"This is a description of My Data Product.","id":"2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd","asset":{"id":"2b0bf220-079c-11ee-be56-0242ac120002","container":{"id":"d29c42eb-7100-4b7a-8257-c196dbcca1cd","type":"catalog"}}}]}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"total_count":2,"limit":1,"releases":[{"version":"1.0.0","state":"draft","data_product":{"id":"b38df608-d34b-4d58-8136-ed25e6c6684e"},"name":"My Data Product","description":"This is a description of My Data Product.","id":"2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd","asset":{"id":"2b0bf220-079c-11ee-be56-0242ac120002","container":{"id":"d29c42eb-7100-4b7a-8257-c196dbcca1cd","type":"catalog"}}}]}`)
					} else {
						res.WriteHeader(400)
					}
				}))
			})
			It(`Use DataProductReleasesPager.GetNext successfully`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				listDataProductReleasesOptionsModel := &dpxv1.ListDataProductReleasesOptions{
					DataProductID:    core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e"),
					AssetContainerID: core.StringPtr("testString"),
					State:            []string{"available"},
					Version:          core.StringPtr("testString"),
					Limit:            core.Int64Ptr(int64(10)),
				}

				pager, err := dpxService.NewDataProductReleasesPager(listDataProductReleasesOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				var allResults []dpxv1.DataProductVersionSummary
				for pager.HasNext() {
					nextPage, err := pager.GetNext()
					Expect(err).To(BeNil())
					Expect(nextPage).ToNot(BeNil())
					allResults = append(allResults, nextPage...)
				}
				Expect(len(allResults)).To(Equal(2))
			})
			It(`Use DataProductReleasesPager.GetAll successfully`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				listDataProductReleasesOptionsModel := &dpxv1.ListDataProductReleasesOptions{
					DataProductID:    core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e"),
					AssetContainerID: core.StringPtr("testString"),
					State:            []string{"available"},
					Version:          core.StringPtr("testString"),
					Limit:            core.Int64Ptr(int64(10)),
				}

				pager, err := dpxService.NewDataProductReleasesPager(listDataProductReleasesOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				allResults, err := pager.GetAll()
				Expect(err).To(BeNil())
				Expect(allResults).ToNot(BeNil())
				Expect(len(allResults)).To(Equal(2))
			})
		})
	})
	Describe(`RetireDataProductRelease(retireDataProductReleaseOptions *RetireDataProductReleaseOptions) - Operation response error`, func() {
		retireDataProductReleasePath := "/data_product_exchange/v1/data_products/b38df608-d34b-4d58-8136-ed25e6c6684e/releases/2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd/retire"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(retireDataProductReleasePath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke RetireDataProductRelease with error: Operation response processing error`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Construct an instance of the RetireDataProductReleaseOptions model
				retireDataProductReleaseOptionsModel := new(dpxv1.RetireDataProductReleaseOptions)
				retireDataProductReleaseOptionsModel.DataProductID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")
				retireDataProductReleaseOptionsModel.ReleaseID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				retireDataProductReleaseOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := dpxService.RetireDataProductRelease(retireDataProductReleaseOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				dpxService.EnableRetries(0, 0)
				result, response, operationErr = dpxService.RetireDataProductRelease(retireDataProductReleaseOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`RetireDataProductRelease(retireDataProductReleaseOptions *RetireDataProductReleaseOptions)`, func() {
		retireDataProductReleasePath := "/data_product_exchange/v1/data_products/b38df608-d34b-4d58-8136-ed25e6c6684e/releases/2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd/retire"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(retireDataProductReleasePath))
					Expect(req.Method).To(Equal("POST"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"version": "1.0.0", "state": "draft", "data_product": {"id": "b38df608-d34b-4d58-8136-ed25e6c6684e"}, "name": "My Data Product", "description": "This is a description of My Data Product.", "id": "2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd", "asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "tags": ["Tags"], "use_cases": [{"id": "ID", "name": "Name", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}], "domain": {"id": "ID", "name": "Name", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "types": ["data"], "parts_out": [{"asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}, "type": "data_asset"}, "revision": 1, "updated_at": "2023-07-01T22:22:34.876Z", "delivery_methods": [{"id": "09cf5fcc-cb9d-4995-a8e4-16517b25229f", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}]}], "published_by": "PublishedBy", "published_at": "2019-01-01T12:00:00.000Z", "contract_terms": [{"asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "id": "ID", "documents": [{"url": "URL", "type": "terms_and_conditions", "name": "Name", "id": "2b0bf220-079c-11ee-be56-0242ac120002", "attachment": {"id": "ID"}, "upload_url": "UploadURL"}]}], "created_by": "CreatedBy", "created_at": "2019-01-01T12:00:00.000Z", "is_restricted": true}`)
				}))
			})
			It(`Invoke RetireDataProductRelease successfully with retries`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())
				dpxService.EnableRetries(0, 0)

				// Construct an instance of the RetireDataProductReleaseOptions model
				retireDataProductReleaseOptionsModel := new(dpxv1.RetireDataProductReleaseOptions)
				retireDataProductReleaseOptionsModel.DataProductID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")
				retireDataProductReleaseOptionsModel.ReleaseID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				retireDataProductReleaseOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := dpxService.RetireDataProductReleaseWithContext(ctx, retireDataProductReleaseOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				dpxService.DisableRetries()
				result, response, operationErr := dpxService.RetireDataProductRelease(retireDataProductReleaseOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = dpxService.RetireDataProductReleaseWithContext(ctx, retireDataProductReleaseOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(retireDataProductReleasePath))
					Expect(req.Method).To(Equal("POST"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"version": "1.0.0", "state": "draft", "data_product": {"id": "b38df608-d34b-4d58-8136-ed25e6c6684e"}, "name": "My Data Product", "description": "This is a description of My Data Product.", "id": "2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd", "asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "tags": ["Tags"], "use_cases": [{"id": "ID", "name": "Name", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}], "domain": {"id": "ID", "name": "Name", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "types": ["data"], "parts_out": [{"asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}, "type": "data_asset"}, "revision": 1, "updated_at": "2023-07-01T22:22:34.876Z", "delivery_methods": [{"id": "09cf5fcc-cb9d-4995-a8e4-16517b25229f", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}]}], "published_by": "PublishedBy", "published_at": "2019-01-01T12:00:00.000Z", "contract_terms": [{"asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "id": "ID", "documents": [{"url": "URL", "type": "terms_and_conditions", "name": "Name", "id": "2b0bf220-079c-11ee-be56-0242ac120002", "attachment": {"id": "ID"}, "upload_url": "UploadURL"}]}], "created_by": "CreatedBy", "created_at": "2019-01-01T12:00:00.000Z", "is_restricted": true}`)
				}))
			})
			It(`Invoke RetireDataProductRelease successfully`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := dpxService.RetireDataProductRelease(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the RetireDataProductReleaseOptions model
				retireDataProductReleaseOptionsModel := new(dpxv1.RetireDataProductReleaseOptions)
				retireDataProductReleaseOptionsModel.DataProductID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")
				retireDataProductReleaseOptionsModel.ReleaseID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				retireDataProductReleaseOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = dpxService.RetireDataProductRelease(retireDataProductReleaseOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke RetireDataProductRelease with error: Operation validation and request error`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Construct an instance of the RetireDataProductReleaseOptions model
				retireDataProductReleaseOptionsModel := new(dpxv1.RetireDataProductReleaseOptions)
				retireDataProductReleaseOptionsModel.DataProductID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")
				retireDataProductReleaseOptionsModel.ReleaseID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				retireDataProductReleaseOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dpxService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := dpxService.RetireDataProductRelease(retireDataProductReleaseOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the RetireDataProductReleaseOptions model with no property values
				retireDataProductReleaseOptionsModelNew := new(dpxv1.RetireDataProductReleaseOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = dpxService.RetireDataProductRelease(retireDataProductReleaseOptionsModelNew)
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
			It(`Invoke RetireDataProductRelease successfully`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Construct an instance of the RetireDataProductReleaseOptions model
				retireDataProductReleaseOptionsModel := new(dpxv1.RetireDataProductReleaseOptions)
				retireDataProductReleaseOptionsModel.DataProductID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")
				retireDataProductReleaseOptionsModel.ReleaseID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				retireDataProductReleaseOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := dpxService.RetireDataProductRelease(retireDataProductReleaseOptionsModel)
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
			dpxService, _ := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
				URL:           "http://dpxv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewAssetPartReference successfully`, func() {
				var container *dpxv1.ContainerReference = nil
				_, err := dpxService.NewAssetPartReference(container)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewAssetReference successfully`, func() {
				var container *dpxv1.ContainerReference = nil
				_, err := dpxService.NewAssetReference(container)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewCompleteDraftContractTermsDocumentOptions successfully`, func() {
				// Construct an instance of the CompleteDraftContractTermsDocumentOptions model
				dataProductID := "b38df608-d34b-4d58-8136-ed25e6c6684e"
				draftID := "2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd"
				contractTermsID := "598183cd-b910-4e8d-9a97-97097afda3c1@e4fe2f87-0e56-46dd-b3b8-e9af32309e82"
				documentID := "testString"
				completeDraftContractTermsDocumentOptionsModel := dpxService.NewCompleteDraftContractTermsDocumentOptions(dataProductID, draftID, contractTermsID, documentID)
				completeDraftContractTermsDocumentOptionsModel.SetDataProductID("b38df608-d34b-4d58-8136-ed25e6c6684e")
				completeDraftContractTermsDocumentOptionsModel.SetDraftID("2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				completeDraftContractTermsDocumentOptionsModel.SetContractTermsID("598183cd-b910-4e8d-9a97-97097afda3c1@e4fe2f87-0e56-46dd-b3b8-e9af32309e82")
				completeDraftContractTermsDocumentOptionsModel.SetDocumentID("testString")
				completeDraftContractTermsDocumentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(completeDraftContractTermsDocumentOptionsModel).ToNot(BeNil())
				Expect(completeDraftContractTermsDocumentOptionsModel.DataProductID).To(Equal(core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")))
				Expect(completeDraftContractTermsDocumentOptionsModel.DraftID).To(Equal(core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd")))
				Expect(completeDraftContractTermsDocumentOptionsModel.ContractTermsID).To(Equal(core.StringPtr("598183cd-b910-4e8d-9a97-97097afda3c1@e4fe2f87-0e56-46dd-b3b8-e9af32309e82")))
				Expect(completeDraftContractTermsDocumentOptionsModel.DocumentID).To(Equal(core.StringPtr("testString")))
				Expect(completeDraftContractTermsDocumentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewContainerReference successfully`, func() {
				id := "d29c42eb-7100-4b7a-8257-c196dbcca1cd"
				_model, err := dpxService.NewContainerReference(id)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewContractTermsDocument successfully`, func() {
				typeVar := "terms_and_conditions"
				name := "testString"
				id := "2b0bf220-079c-11ee-be56-0242ac120002"
				_model, err := dpxService.NewContractTermsDocument(typeVar, name, id)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewContractTermsDocumentPatch successfully`, func() {
				// Construct an instance of the ContractTermsDocumentAttachment model
				contractTermsDocumentAttachmentModel := new(dpxv1.ContractTermsDocumentAttachment)
				contractTermsDocumentAttachmentModel.ID = core.StringPtr("testString")

				// Construct an instance of the ContractTermsDocument model
				contractTermsDocument := new(dpxv1.ContractTermsDocument)
				contractTermsDocument.URL = core.StringPtr("testString")
				contractTermsDocument.Type = core.StringPtr("terms_and_conditions")
				contractTermsDocument.Name = core.StringPtr("testString")
				contractTermsDocument.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				contractTermsDocument.Attachment = contractTermsDocumentAttachmentModel
				contractTermsDocument.UploadURL = core.StringPtr("testString")

				contractTermsDocumentPatch := dpxService.NewContractTermsDocumentPatch(contractTermsDocument)
				Expect(contractTermsDocumentPatch).ToNot(BeNil())

				_path := func(op interface{}) string {
					return *op.(dpxv1.JSONPatchOperation).Path
				}
				Expect(contractTermsDocumentPatch).To(MatchAllElements(_path, Elements{
					"/url": MatchAllFields(Fields{
						"Op":    PointTo(Equal(dpxv1.JSONPatchOperation_Op_Add)),
						"Path":  PointTo(Equal("/url")),
						"From":  BeNil(),
						"Value": Equal(contractTermsDocument.URL),
					}),
					"/type": MatchAllFields(Fields{
						"Op":    PointTo(Equal(dpxv1.JSONPatchOperation_Op_Add)),
						"Path":  PointTo(Equal("/type")),
						"From":  BeNil(),
						"Value": Equal(contractTermsDocument.Type),
					}),
					"/name": MatchAllFields(Fields{
						"Op":    PointTo(Equal(dpxv1.JSONPatchOperation_Op_Add)),
						"Path":  PointTo(Equal("/name")),
						"From":  BeNil(),
						"Value": Equal(contractTermsDocument.Name),
					}),
					"/id": MatchAllFields(Fields{
						"Op":    PointTo(Equal(dpxv1.JSONPatchOperation_Op_Add)),
						"Path":  PointTo(Equal("/id")),
						"From":  BeNil(),
						"Value": Equal(contractTermsDocument.ID),
					}),
					"/attachment": MatchAllFields(Fields{
						"Op":    PointTo(Equal(dpxv1.JSONPatchOperation_Op_Add)),
						"Path":  PointTo(Equal("/attachment")),
						"From":  BeNil(),
						"Value": Equal(contractTermsDocument.Attachment),
					}),
					"/upload_url": MatchAllFields(Fields{
						"Op":    PointTo(Equal(dpxv1.JSONPatchOperation_Op_Add)),
						"Path":  PointTo(Equal("/upload_url")),
						"From":  BeNil(),
						"Value": Equal(contractTermsDocument.UploadURL),
					}),
				}))
			})
			It(`Invoke NewCreateDataProductDraftOptions successfully`, func() {
				// Construct an instance of the ContainerReference model
				containerReferenceModel := new(dpxv1.ContainerReference)
				Expect(containerReferenceModel).ToNot(BeNil())
				containerReferenceModel.ID = core.StringPtr("d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				containerReferenceModel.Type = core.StringPtr("catalog")
				Expect(containerReferenceModel.ID).To(Equal(core.StringPtr("d29c42eb-7100-4b7a-8257-c196dbcca1cd")))
				Expect(containerReferenceModel.Type).To(Equal(core.StringPtr("catalog")))

				// Construct an instance of the AssetReference model
				assetReferenceModel := new(dpxv1.AssetReference)
				Expect(assetReferenceModel).ToNot(BeNil())
				assetReferenceModel.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				assetReferenceModel.Container = containerReferenceModel
				Expect(assetReferenceModel.ID).To(Equal(core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")))
				Expect(assetReferenceModel.Container).To(Equal(containerReferenceModel))

				// Construct an instance of the DataProductIdentity model
				dataProductIdentityModel := new(dpxv1.DataProductIdentity)
				Expect(dataProductIdentityModel).ToNot(BeNil())
				dataProductIdentityModel.ID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")
				Expect(dataProductIdentityModel.ID).To(Equal(core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")))

				// Construct an instance of the UseCase model
				useCaseModel := new(dpxv1.UseCase)
				Expect(useCaseModel).ToNot(BeNil())
				useCaseModel.ID = core.StringPtr("testString")
				useCaseModel.Name = core.StringPtr("testString")
				useCaseModel.Container = containerReferenceModel
				Expect(useCaseModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(useCaseModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(useCaseModel.Container).To(Equal(containerReferenceModel))

				// Construct an instance of the Domain model
				domainModel := new(dpxv1.Domain)
				Expect(domainModel).ToNot(BeNil())
				domainModel.ID = core.StringPtr("testString")
				domainModel.Name = core.StringPtr("testString")
				domainModel.Container = containerReferenceModel
				Expect(domainModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(domainModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(domainModel.Container).To(Equal(containerReferenceModel))

				// Construct an instance of the AssetPartReference model
				assetPartReferenceModel := new(dpxv1.AssetPartReference)
				Expect(assetPartReferenceModel).ToNot(BeNil())
				assetPartReferenceModel.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				assetPartReferenceModel.Container = containerReferenceModel
				assetPartReferenceModel.Type = core.StringPtr("data_asset")
				Expect(assetPartReferenceModel.ID).To(Equal(core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")))
				Expect(assetPartReferenceModel.Container).To(Equal(containerReferenceModel))
				Expect(assetPartReferenceModel.Type).To(Equal(core.StringPtr("data_asset")))

				// Construct an instance of the DeliveryMethod model
				deliveryMethodModel := new(dpxv1.DeliveryMethod)
				Expect(deliveryMethodModel).ToNot(BeNil())
				deliveryMethodModel.ID = core.StringPtr("09cf5fcc-cb9d-4995-a8e4-16517b25229f")
				deliveryMethodModel.Container = containerReferenceModel
				Expect(deliveryMethodModel.ID).To(Equal(core.StringPtr("09cf5fcc-cb9d-4995-a8e4-16517b25229f")))
				Expect(deliveryMethodModel.Container).To(Equal(containerReferenceModel))

				// Construct an instance of the DataProductPart model
				dataProductPartModel := new(dpxv1.DataProductPart)
				Expect(dataProductPartModel).ToNot(BeNil())
				dataProductPartModel.Asset = assetPartReferenceModel
				dataProductPartModel.Revision = core.Int64Ptr(int64(1))
				dataProductPartModel.UpdatedAt = CreateMockDateTime("2023-07-01T22:22:34.876Z")
				dataProductPartModel.DeliveryMethods = []dpxv1.DeliveryMethod{*deliveryMethodModel}
				Expect(dataProductPartModel.Asset).To(Equal(assetPartReferenceModel))
				Expect(dataProductPartModel.Revision).To(Equal(core.Int64Ptr(int64(1))))
				Expect(dataProductPartModel.UpdatedAt).To(Equal(CreateMockDateTime("2023-07-01T22:22:34.876Z")))
				Expect(dataProductPartModel.DeliveryMethods).To(Equal([]dpxv1.DeliveryMethod{*deliveryMethodModel}))

				// Construct an instance of the ContractTermsDocumentAttachment model
				contractTermsDocumentAttachmentModel := new(dpxv1.ContractTermsDocumentAttachment)
				Expect(contractTermsDocumentAttachmentModel).ToNot(BeNil())
				contractTermsDocumentAttachmentModel.ID = core.StringPtr("testString")
				Expect(contractTermsDocumentAttachmentModel.ID).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the ContractTermsDocument model
				contractTermsDocumentModel := new(dpxv1.ContractTermsDocument)
				Expect(contractTermsDocumentModel).ToNot(BeNil())
				contractTermsDocumentModel.URL = core.StringPtr("testString")
				contractTermsDocumentModel.Type = core.StringPtr("terms_and_conditions")
				contractTermsDocumentModel.Name = core.StringPtr("testString")
				contractTermsDocumentModel.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				contractTermsDocumentModel.Attachment = contractTermsDocumentAttachmentModel
				contractTermsDocumentModel.UploadURL = core.StringPtr("testString")
				Expect(contractTermsDocumentModel.URL).To(Equal(core.StringPtr("testString")))
				Expect(contractTermsDocumentModel.Type).To(Equal(core.StringPtr("terms_and_conditions")))
				Expect(contractTermsDocumentModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(contractTermsDocumentModel.ID).To(Equal(core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")))
				Expect(contractTermsDocumentModel.Attachment).To(Equal(contractTermsDocumentAttachmentModel))
				Expect(contractTermsDocumentModel.UploadURL).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the DataProductContractTerms model
				dataProductContractTermsModel := new(dpxv1.DataProductContractTerms)
				Expect(dataProductContractTermsModel).ToNot(BeNil())
				dataProductContractTermsModel.Asset = assetReferenceModel
				dataProductContractTermsModel.ID = core.StringPtr("testString")
				dataProductContractTermsModel.Documents = []dpxv1.ContractTermsDocument{*contractTermsDocumentModel}
				Expect(dataProductContractTermsModel.Asset).To(Equal(assetReferenceModel))
				Expect(dataProductContractTermsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(dataProductContractTermsModel.Documents).To(Equal([]dpxv1.ContractTermsDocument{*contractTermsDocumentModel}))

				// Construct an instance of the CreateDataProductDraftOptions model
				dataProductID := "b38df608-d34b-4d58-8136-ed25e6c6684e"
				var createDataProductDraftOptionsAsset *dpxv1.AssetReference = nil
				createDataProductDraftOptionsModel := dpxService.NewCreateDataProductDraftOptions(dataProductID, createDataProductDraftOptionsAsset)
				createDataProductDraftOptionsModel.SetDataProductID("b38df608-d34b-4d58-8136-ed25e6c6684e")
				createDataProductDraftOptionsModel.SetAsset(assetReferenceModel)
				createDataProductDraftOptionsModel.SetVersion("1.2.0")
				createDataProductDraftOptionsModel.SetState("draft")
				createDataProductDraftOptionsModel.SetDataProduct(dataProductIdentityModel)
				createDataProductDraftOptionsModel.SetName("testString")
				createDataProductDraftOptionsModel.SetDescription("testString")
				createDataProductDraftOptionsModel.SetTags([]string{"testString"})
				createDataProductDraftOptionsModel.SetUseCases([]dpxv1.UseCase{*useCaseModel})
				createDataProductDraftOptionsModel.SetDomain(domainModel)
				createDataProductDraftOptionsModel.SetTypes([]string{"data"})
				createDataProductDraftOptionsModel.SetPartsOut([]dpxv1.DataProductPart{*dataProductPartModel})
				createDataProductDraftOptionsModel.SetContractTerms([]dpxv1.DataProductContractTerms{*dataProductContractTermsModel})
				createDataProductDraftOptionsModel.SetIsRestricted(true)
				createDataProductDraftOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createDataProductDraftOptionsModel).ToNot(BeNil())
				Expect(createDataProductDraftOptionsModel.DataProductID).To(Equal(core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")))
				Expect(createDataProductDraftOptionsModel.Asset).To(Equal(assetReferenceModel))
				Expect(createDataProductDraftOptionsModel.Version).To(Equal(core.StringPtr("1.2.0")))
				Expect(createDataProductDraftOptionsModel.State).To(Equal(core.StringPtr("draft")))
				Expect(createDataProductDraftOptionsModel.DataProduct).To(Equal(dataProductIdentityModel))
				Expect(createDataProductDraftOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(createDataProductDraftOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(createDataProductDraftOptionsModel.Tags).To(Equal([]string{"testString"}))
				Expect(createDataProductDraftOptionsModel.UseCases).To(Equal([]dpxv1.UseCase{*useCaseModel}))
				Expect(createDataProductDraftOptionsModel.Domain).To(Equal(domainModel))
				Expect(createDataProductDraftOptionsModel.Types).To(Equal([]string{"data"}))
				Expect(createDataProductDraftOptionsModel.PartsOut).To(Equal([]dpxv1.DataProductPart{*dataProductPartModel}))
				Expect(createDataProductDraftOptionsModel.ContractTerms).To(Equal([]dpxv1.DataProductContractTerms{*dataProductContractTermsModel}))
				Expect(createDataProductDraftOptionsModel.IsRestricted).To(Equal(core.BoolPtr(true)))
				Expect(createDataProductDraftOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateDataProductOptions successfully`, func() {
				// Construct an instance of the DataProductIdentity model
				dataProductIdentityModel := new(dpxv1.DataProductIdentity)
				Expect(dataProductIdentityModel).ToNot(BeNil())
				dataProductIdentityModel.ID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")
				Expect(dataProductIdentityModel.ID).To(Equal(core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")))

				// Construct an instance of the ContainerReference model
				containerReferenceModel := new(dpxv1.ContainerReference)
				Expect(containerReferenceModel).ToNot(BeNil())
				containerReferenceModel.ID = core.StringPtr("d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				containerReferenceModel.Type = core.StringPtr("catalog")
				Expect(containerReferenceModel.ID).To(Equal(core.StringPtr("d29c42eb-7100-4b7a-8257-c196dbcca1cd")))
				Expect(containerReferenceModel.Type).To(Equal(core.StringPtr("catalog")))

				// Construct an instance of the AssetReference model
				assetReferenceModel := new(dpxv1.AssetReference)
				Expect(assetReferenceModel).ToNot(BeNil())
				assetReferenceModel.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				assetReferenceModel.Container = containerReferenceModel
				Expect(assetReferenceModel.ID).To(Equal(core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")))
				Expect(assetReferenceModel.Container).To(Equal(containerReferenceModel))

				// Construct an instance of the UseCase model
				useCaseModel := new(dpxv1.UseCase)
				Expect(useCaseModel).ToNot(BeNil())
				useCaseModel.ID = core.StringPtr("testString")
				useCaseModel.Name = core.StringPtr("testString")
				useCaseModel.Container = containerReferenceModel
				Expect(useCaseModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(useCaseModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(useCaseModel.Container).To(Equal(containerReferenceModel))

				// Construct an instance of the Domain model
				domainModel := new(dpxv1.Domain)
				Expect(domainModel).ToNot(BeNil())
				domainModel.ID = core.StringPtr("testString")
				domainModel.Name = core.StringPtr("testString")
				domainModel.Container = containerReferenceModel
				Expect(domainModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(domainModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(domainModel.Container).To(Equal(containerReferenceModel))

				// Construct an instance of the AssetPartReference model
				assetPartReferenceModel := new(dpxv1.AssetPartReference)
				Expect(assetPartReferenceModel).ToNot(BeNil())
				assetPartReferenceModel.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				assetPartReferenceModel.Container = containerReferenceModel
				assetPartReferenceModel.Type = core.StringPtr("data_asset")
				Expect(assetPartReferenceModel.ID).To(Equal(core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")))
				Expect(assetPartReferenceModel.Container).To(Equal(containerReferenceModel))
				Expect(assetPartReferenceModel.Type).To(Equal(core.StringPtr("data_asset")))

				// Construct an instance of the DeliveryMethod model
				deliveryMethodModel := new(dpxv1.DeliveryMethod)
				Expect(deliveryMethodModel).ToNot(BeNil())
				deliveryMethodModel.ID = core.StringPtr("09cf5fcc-cb9d-4995-a8e4-16517b25229f")
				deliveryMethodModel.Container = containerReferenceModel
				Expect(deliveryMethodModel.ID).To(Equal(core.StringPtr("09cf5fcc-cb9d-4995-a8e4-16517b25229f")))
				Expect(deliveryMethodModel.Container).To(Equal(containerReferenceModel))

				// Construct an instance of the DataProductPart model
				dataProductPartModel := new(dpxv1.DataProductPart)
				Expect(dataProductPartModel).ToNot(BeNil())
				dataProductPartModel.Asset = assetPartReferenceModel
				dataProductPartModel.Revision = core.Int64Ptr(int64(1))
				dataProductPartModel.UpdatedAt = CreateMockDateTime("2023-07-01T22:22:34.876Z")
				dataProductPartModel.DeliveryMethods = []dpxv1.DeliveryMethod{*deliveryMethodModel}
				Expect(dataProductPartModel.Asset).To(Equal(assetPartReferenceModel))
				Expect(dataProductPartModel.Revision).To(Equal(core.Int64Ptr(int64(1))))
				Expect(dataProductPartModel.UpdatedAt).To(Equal(CreateMockDateTime("2023-07-01T22:22:34.876Z")))
				Expect(dataProductPartModel.DeliveryMethods).To(Equal([]dpxv1.DeliveryMethod{*deliveryMethodModel}))

				// Construct an instance of the ContractTermsDocumentAttachment model
				contractTermsDocumentAttachmentModel := new(dpxv1.ContractTermsDocumentAttachment)
				Expect(contractTermsDocumentAttachmentModel).ToNot(BeNil())
				contractTermsDocumentAttachmentModel.ID = core.StringPtr("testString")
				Expect(contractTermsDocumentAttachmentModel.ID).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the ContractTermsDocument model
				contractTermsDocumentModel := new(dpxv1.ContractTermsDocument)
				Expect(contractTermsDocumentModel).ToNot(BeNil())
				contractTermsDocumentModel.URL = core.StringPtr("testString")
				contractTermsDocumentModel.Type = core.StringPtr("terms_and_conditions")
				contractTermsDocumentModel.Name = core.StringPtr("testString")
				contractTermsDocumentModel.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				contractTermsDocumentModel.Attachment = contractTermsDocumentAttachmentModel
				contractTermsDocumentModel.UploadURL = core.StringPtr("testString")
				Expect(contractTermsDocumentModel.URL).To(Equal(core.StringPtr("testString")))
				Expect(contractTermsDocumentModel.Type).To(Equal(core.StringPtr("terms_and_conditions")))
				Expect(contractTermsDocumentModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(contractTermsDocumentModel.ID).To(Equal(core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")))
				Expect(contractTermsDocumentModel.Attachment).To(Equal(contractTermsDocumentAttachmentModel))
				Expect(contractTermsDocumentModel.UploadURL).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the DataProductContractTerms model
				dataProductContractTermsModel := new(dpxv1.DataProductContractTerms)
				Expect(dataProductContractTermsModel).ToNot(BeNil())
				dataProductContractTermsModel.Asset = assetReferenceModel
				dataProductContractTermsModel.ID = core.StringPtr("testString")
				dataProductContractTermsModel.Documents = []dpxv1.ContractTermsDocument{*contractTermsDocumentModel}
				Expect(dataProductContractTermsModel.Asset).To(Equal(assetReferenceModel))
				Expect(dataProductContractTermsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(dataProductContractTermsModel.Documents).To(Equal([]dpxv1.ContractTermsDocument{*contractTermsDocumentModel}))

				// Construct an instance of the DataProductVersionPrototype model
				dataProductVersionPrototypeModel := new(dpxv1.DataProductVersionPrototype)
				Expect(dataProductVersionPrototypeModel).ToNot(BeNil())
				dataProductVersionPrototypeModel.Version = core.StringPtr("1.0.0")
				dataProductVersionPrototypeModel.State = core.StringPtr("draft")
				dataProductVersionPrototypeModel.DataProduct = dataProductIdentityModel
				dataProductVersionPrototypeModel.Name = core.StringPtr("My New Data Product")
				dataProductVersionPrototypeModel.Description = core.StringPtr("This is a description of My Data Product.")
				dataProductVersionPrototypeModel.Asset = assetReferenceModel
				dataProductVersionPrototypeModel.Tags = []string{"testString"}
				dataProductVersionPrototypeModel.UseCases = []dpxv1.UseCase{*useCaseModel}
				dataProductVersionPrototypeModel.Domain = domainModel
				dataProductVersionPrototypeModel.Types = []string{"data"}
				dataProductVersionPrototypeModel.PartsOut = []dpxv1.DataProductPart{*dataProductPartModel}
				dataProductVersionPrototypeModel.ContractTerms = []dpxv1.DataProductContractTerms{*dataProductContractTermsModel}
				dataProductVersionPrototypeModel.IsRestricted = core.BoolPtr(true)
				Expect(dataProductVersionPrototypeModel.Version).To(Equal(core.StringPtr("1.0.0")))
				Expect(dataProductVersionPrototypeModel.State).To(Equal(core.StringPtr("draft")))
				Expect(dataProductVersionPrototypeModel.DataProduct).To(Equal(dataProductIdentityModel))
				Expect(dataProductVersionPrototypeModel.Name).To(Equal(core.StringPtr("My New Data Product")))
				Expect(dataProductVersionPrototypeModel.Description).To(Equal(core.StringPtr("This is a description of My Data Product.")))
				Expect(dataProductVersionPrototypeModel.Asset).To(Equal(assetReferenceModel))
				Expect(dataProductVersionPrototypeModel.Tags).To(Equal([]string{"testString"}))
				Expect(dataProductVersionPrototypeModel.UseCases).To(Equal([]dpxv1.UseCase{*useCaseModel}))
				Expect(dataProductVersionPrototypeModel.Domain).To(Equal(domainModel))
				Expect(dataProductVersionPrototypeModel.Types).To(Equal([]string{"data"}))
				Expect(dataProductVersionPrototypeModel.PartsOut).To(Equal([]dpxv1.DataProductPart{*dataProductPartModel}))
				Expect(dataProductVersionPrototypeModel.ContractTerms).To(Equal([]dpxv1.DataProductContractTerms{*dataProductContractTermsModel}))
				Expect(dataProductVersionPrototypeModel.IsRestricted).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the CreateDataProductOptions model
				createDataProductOptionsDrafts := []dpxv1.DataProductVersionPrototype{}
				createDataProductOptionsModel := dpxService.NewCreateDataProductOptions(createDataProductOptionsDrafts)
				createDataProductOptionsModel.SetDrafts([]dpxv1.DataProductVersionPrototype{*dataProductVersionPrototypeModel})
				createDataProductOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createDataProductOptionsModel).ToNot(BeNil())
				Expect(createDataProductOptionsModel.Drafts).To(Equal([]dpxv1.DataProductVersionPrototype{*dataProductVersionPrototypeModel}))
				Expect(createDataProductOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateDraftContractTermsDocumentOptions successfully`, func() {
				// Construct an instance of the ContractTermsDocumentAttachment model
				contractTermsDocumentAttachmentModel := new(dpxv1.ContractTermsDocumentAttachment)
				Expect(contractTermsDocumentAttachmentModel).ToNot(BeNil())
				contractTermsDocumentAttachmentModel.ID = core.StringPtr("testString")
				Expect(contractTermsDocumentAttachmentModel.ID).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the CreateDraftContractTermsDocumentOptions model
				dataProductID := "b38df608-d34b-4d58-8136-ed25e6c6684e"
				draftID := "2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd"
				contractTermsID := "598183cd-b910-4e8d-9a97-97097afda3c1@e4fe2f87-0e56-46dd-b3b8-e9af32309e82"
				createDraftContractTermsDocumentOptionsType := "terms_and_conditions"
				createDraftContractTermsDocumentOptionsName := "Terms and conditions document"
				createDraftContractTermsDocumentOptionsID := "b38df608-d34b-4d58-8136-ed25e6c6684e"
				url := "teststring"
				createDraftContractTermsDocumentOptionsModel := dpxService.NewCreateDraftContractTermsDocumentOptions(dataProductID, draftID, contractTermsID, createDraftContractTermsDocumentOptionsType, createDraftContractTermsDocumentOptionsName, createDraftContractTermsDocumentOptionsID, url)
				createDraftContractTermsDocumentOptionsModel.SetDataProductID("b38df608-d34b-4d58-8136-ed25e6c6684e")
				createDraftContractTermsDocumentOptionsModel.SetDraftID("2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				createDraftContractTermsDocumentOptionsModel.SetContractTermsID("598183cd-b910-4e8d-9a97-97097afda3c1@e4fe2f87-0e56-46dd-b3b8-e9af32309e82")
				createDraftContractTermsDocumentOptionsModel.SetType("terms_and_conditions")
				createDraftContractTermsDocumentOptionsModel.SetName("Terms and conditions document")
				createDraftContractTermsDocumentOptionsModel.SetID("b38df608-d34b-4d58-8136-ed25e6c6684e")
				createDraftContractTermsDocumentOptionsModel.SetURL("testString")
				createDraftContractTermsDocumentOptionsModel.SetAttachment(contractTermsDocumentAttachmentModel)
				createDraftContractTermsDocumentOptionsModel.SetUploadURL("testString")
				createDraftContractTermsDocumentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createDraftContractTermsDocumentOptionsModel).ToNot(BeNil())
				Expect(createDraftContractTermsDocumentOptionsModel.DataProductID).To(Equal(core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")))
				Expect(createDraftContractTermsDocumentOptionsModel.DraftID).To(Equal(core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd")))
				Expect(createDraftContractTermsDocumentOptionsModel.ContractTermsID).To(Equal(core.StringPtr("598183cd-b910-4e8d-9a97-97097afda3c1@e4fe2f87-0e56-46dd-b3b8-e9af32309e82")))
				Expect(createDraftContractTermsDocumentOptionsModel.Type).To(Equal(core.StringPtr("terms_and_conditions")))
				Expect(createDraftContractTermsDocumentOptionsModel.Name).To(Equal(core.StringPtr("Terms and conditions document")))
				Expect(createDraftContractTermsDocumentOptionsModel.ID).To(Equal(core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")))
				Expect(createDraftContractTermsDocumentOptionsModel.URL).To(Equal(core.StringPtr("testString")))
				Expect(createDraftContractTermsDocumentOptionsModel.Attachment).To(Equal(contractTermsDocumentAttachmentModel))
				Expect(createDraftContractTermsDocumentOptionsModel.UploadURL).To(Equal(core.StringPtr("testString")))
				Expect(createDraftContractTermsDocumentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDataProductIdentity successfully`, func() {
				id := "b38df608-d34b-4d58-8136-ed25e6c6684e"
				_model, err := dpxService.NewDataProductIdentity(id)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewDataProductPart successfully`, func() {
				var asset *dpxv1.AssetPartReference = nil
				_, err := dpxService.NewDataProductPart(asset)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewDataProductVersionPatch successfully`, func() {
				// Construct an instance of the DataProductIdentity model
				dataProductIdentityModel := new(dpxv1.DataProductIdentity)
				dataProductIdentityModel.ID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")

				// Construct an instance of the ContainerReference model
				containerReferenceModel := new(dpxv1.ContainerReference)
				containerReferenceModel.ID = core.StringPtr("d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				containerReferenceModel.Type = core.StringPtr("catalog")

				// Construct an instance of the AssetReference model
				assetReferenceModel := new(dpxv1.AssetReference)
				assetReferenceModel.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				assetReferenceModel.Container = containerReferenceModel

				// Construct an instance of the UseCase model
				useCaseModel := new(dpxv1.UseCase)
				useCaseModel.ID = core.StringPtr("testString")
				useCaseModel.Name = core.StringPtr("testString")
				useCaseModel.Container = containerReferenceModel

				// Construct an instance of the Domain model
				domainModel := new(dpxv1.Domain)
				domainModel.ID = core.StringPtr("testString")
				domainModel.Name = core.StringPtr("testString")
				domainModel.Container = containerReferenceModel

				// Construct an instance of the AssetPartReference model
				assetPartReferenceModel := new(dpxv1.AssetPartReference)
				assetPartReferenceModel.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				assetPartReferenceModel.Container = containerReferenceModel
				assetPartReferenceModel.Type = core.StringPtr("data_asset")

				// Construct an instance of the DeliveryMethod model
				deliveryMethodModel := new(dpxv1.DeliveryMethod)
				deliveryMethodModel.ID = core.StringPtr("09cf5fcc-cb9d-4995-a8e4-16517b25229f")
				deliveryMethodModel.Container = containerReferenceModel

				// Construct an instance of the DataProductPart model
				dataProductPartModel := new(dpxv1.DataProductPart)
				dataProductPartModel.Asset = assetPartReferenceModel
				dataProductPartModel.Revision = core.Int64Ptr(int64(1))
				dataProductPartModel.UpdatedAt = CreateMockDateTime("2023-07-01T22:22:34.876Z")
				dataProductPartModel.DeliveryMethods = []dpxv1.DeliveryMethod{*deliveryMethodModel}

				// Construct an instance of the ContractTermsDocumentAttachment model
				contractTermsDocumentAttachmentModel := new(dpxv1.ContractTermsDocumentAttachment)
				contractTermsDocumentAttachmentModel.ID = core.StringPtr("testString")

				// Construct an instance of the ContractTermsDocument model
				contractTermsDocumentModel := new(dpxv1.ContractTermsDocument)
				contractTermsDocumentModel.URL = core.StringPtr("testString")
				contractTermsDocumentModel.Type = core.StringPtr("terms_and_conditions")
				contractTermsDocumentModel.Name = core.StringPtr("testString")
				contractTermsDocumentModel.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				contractTermsDocumentModel.Attachment = contractTermsDocumentAttachmentModel
				contractTermsDocumentModel.UploadURL = core.StringPtr("testString")

				// Construct an instance of the DataProductContractTerms model
				dataProductContractTermsModel := new(dpxv1.DataProductContractTerms)
				dataProductContractTermsModel.Asset = assetReferenceModel
				dataProductContractTermsModel.ID = core.StringPtr("testString")
				dataProductContractTermsModel.Documents = []dpxv1.ContractTermsDocument{*contractTermsDocumentModel}

				// Construct an instance of the DataProductVersion model
				dataProductVersion := new(dpxv1.DataProductVersion)
				dataProductVersion.Version = core.StringPtr("1.0.0")
				dataProductVersion.State = core.StringPtr("draft")
				dataProductVersion.DataProduct = dataProductIdentityModel
				dataProductVersion.Name = core.StringPtr("My Data Product")
				dataProductVersion.Description = core.StringPtr("This is a description of My Data Product.")
				dataProductVersion.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				dataProductVersion.Asset = assetReferenceModel
				dataProductVersion.Tags = []string{"testString"}
				dataProductVersion.UseCases = []dpxv1.UseCase{*useCaseModel}
				dataProductVersion.Domain = domainModel
				dataProductVersion.Types = []string{"data"}
				dataProductVersion.PartsOut = []dpxv1.DataProductPart{*dataProductPartModel}
				dataProductVersion.PublishedBy = core.StringPtr("testString")
				dataProductVersion.PublishedAt = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				dataProductVersion.ContractTerms = []dpxv1.DataProductContractTerms{*dataProductContractTermsModel}
				dataProductVersion.CreatedBy = core.StringPtr("testString")
				dataProductVersion.CreatedAt = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				dataProductVersion.IsRestricted = core.BoolPtr(true)

				dataProductVersionPatch := dpxService.NewDataProductVersionPatch(dataProductVersion)
				Expect(dataProductVersionPatch).ToNot(BeNil())

				_path := func(op interface{}) string {
					return *op.(dpxv1.JSONPatchOperation).Path
				}
				Expect(dataProductVersionPatch).To(MatchAllElements(_path, Elements{
					"/version": MatchAllFields(Fields{
						"Op":    PointTo(Equal(dpxv1.JSONPatchOperation_Op_Add)),
						"Path":  PointTo(Equal("/version")),
						"From":  BeNil(),
						"Value": Equal(dataProductVersion.Version),
					}),
					"/state": MatchAllFields(Fields{
						"Op":    PointTo(Equal(dpxv1.JSONPatchOperation_Op_Add)),
						"Path":  PointTo(Equal("/state")),
						"From":  BeNil(),
						"Value": Equal(dataProductVersion.State),
					}),
					"/data_product": MatchAllFields(Fields{
						"Op":    PointTo(Equal(dpxv1.JSONPatchOperation_Op_Add)),
						"Path":  PointTo(Equal("/data_product")),
						"From":  BeNil(),
						"Value": Equal(dataProductVersion.DataProduct),
					}),
					"/name": MatchAllFields(Fields{
						"Op":    PointTo(Equal(dpxv1.JSONPatchOperation_Op_Add)),
						"Path":  PointTo(Equal("/name")),
						"From":  BeNil(),
						"Value": Equal(dataProductVersion.Name),
					}),
					"/description": MatchAllFields(Fields{
						"Op":    PointTo(Equal(dpxv1.JSONPatchOperation_Op_Add)),
						"Path":  PointTo(Equal("/description")),
						"From":  BeNil(),
						"Value": Equal(dataProductVersion.Description),
					}),
					"/id": MatchAllFields(Fields{
						"Op":    PointTo(Equal(dpxv1.JSONPatchOperation_Op_Add)),
						"Path":  PointTo(Equal("/id")),
						"From":  BeNil(),
						"Value": Equal(dataProductVersion.ID),
					}),
					"/asset": MatchAllFields(Fields{
						"Op":    PointTo(Equal(dpxv1.JSONPatchOperation_Op_Add)),
						"Path":  PointTo(Equal("/asset")),
						"From":  BeNil(),
						"Value": Equal(dataProductVersion.Asset),
					}),
					"/tags": MatchAllFields(Fields{
						"Op":    PointTo(Equal(dpxv1.JSONPatchOperation_Op_Add)),
						"Path":  PointTo(Equal("/tags")),
						"From":  BeNil(),
						"Value": Equal(dataProductVersion.Tags),
					}),
					"/use_cases": MatchAllFields(Fields{
						"Op":    PointTo(Equal(dpxv1.JSONPatchOperation_Op_Add)),
						"Path":  PointTo(Equal("/use_cases")),
						"From":  BeNil(),
						"Value": Equal(dataProductVersion.UseCases),
					}),
					"/domain": MatchAllFields(Fields{
						"Op":    PointTo(Equal(dpxv1.JSONPatchOperation_Op_Add)),
						"Path":  PointTo(Equal("/domain")),
						"From":  BeNil(),
						"Value": Equal(dataProductVersion.Domain),
					}),
					"/types": MatchAllFields(Fields{
						"Op":    PointTo(Equal(dpxv1.JSONPatchOperation_Op_Add)),
						"Path":  PointTo(Equal("/types")),
						"From":  BeNil(),
						"Value": Equal(dataProductVersion.Types),
					}),
					"/parts_out": MatchAllFields(Fields{
						"Op":    PointTo(Equal(dpxv1.JSONPatchOperation_Op_Add)),
						"Path":  PointTo(Equal("/parts_out")),
						"From":  BeNil(),
						"Value": Equal(dataProductVersion.PartsOut),
					}),
					"/published_by": MatchAllFields(Fields{
						"Op":    PointTo(Equal(dpxv1.JSONPatchOperation_Op_Add)),
						"Path":  PointTo(Equal("/published_by")),
						"From":  BeNil(),
						"Value": Equal(dataProductVersion.PublishedBy),
					}),
					"/published_at": MatchAllFields(Fields{
						"Op":    PointTo(Equal(dpxv1.JSONPatchOperation_Op_Add)),
						"Path":  PointTo(Equal("/published_at")),
						"From":  BeNil(),
						"Value": Equal(dataProductVersion.PublishedAt),
					}),
					"/contract_terms": MatchAllFields(Fields{
						"Op":    PointTo(Equal(dpxv1.JSONPatchOperation_Op_Add)),
						"Path":  PointTo(Equal("/contract_terms")),
						"From":  BeNil(),
						"Value": Equal(dataProductVersion.ContractTerms),
					}),
					"/created_by": MatchAllFields(Fields{
						"Op":    PointTo(Equal(dpxv1.JSONPatchOperation_Op_Add)),
						"Path":  PointTo(Equal("/created_by")),
						"From":  BeNil(),
						"Value": Equal(dataProductVersion.CreatedBy),
					}),
					"/created_at": MatchAllFields(Fields{
						"Op":    PointTo(Equal(dpxv1.JSONPatchOperation_Op_Add)),
						"Path":  PointTo(Equal("/created_at")),
						"From":  BeNil(),
						"Value": Equal(dataProductVersion.CreatedAt),
					}),
					"/is_restricted": MatchAllFields(Fields{
						"Op":    PointTo(Equal(dpxv1.JSONPatchOperation_Op_Add)),
						"Path":  PointTo(Equal("/is_restricted")),
						"From":  BeNil(),
						"Value": Equal(dataProductVersion.IsRestricted),
					}),
				}))
			})
			It(`Invoke NewDataProductVersionPrototype successfully`, func() {
				var asset *dpxv1.AssetReference = nil
				_, err := dpxService.NewDataProductVersionPrototype(asset)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewDeleteDataProductDraftOptions successfully`, func() {
				// Construct an instance of the DeleteDataProductDraftOptions model
				dataProductID := "b38df608-d34b-4d58-8136-ed25e6c6684e"
				draftID := "2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd"
				deleteDataProductDraftOptionsModel := dpxService.NewDeleteDataProductDraftOptions(dataProductID, draftID)
				deleteDataProductDraftOptionsModel.SetDataProductID("b38df608-d34b-4d58-8136-ed25e6c6684e")
				deleteDataProductDraftOptionsModel.SetDraftID("2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				deleteDataProductDraftOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteDataProductDraftOptionsModel).ToNot(BeNil())
				Expect(deleteDataProductDraftOptionsModel.DataProductID).To(Equal(core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")))
				Expect(deleteDataProductDraftOptionsModel.DraftID).To(Equal(core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd")))
				Expect(deleteDataProductDraftOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteDraftContractTermsDocumentOptions successfully`, func() {
				// Construct an instance of the DeleteDraftContractTermsDocumentOptions model
				dataProductID := "b38df608-d34b-4d58-8136-ed25e6c6684e"
				draftID := "2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd"
				contractTermsID := "598183cd-b910-4e8d-9a97-97097afda3c1@e4fe2f87-0e56-46dd-b3b8-e9af32309e82"
				documentID := "testString"
				deleteDraftContractTermsDocumentOptionsModel := dpxService.NewDeleteDraftContractTermsDocumentOptions(dataProductID, draftID, contractTermsID, documentID)
				deleteDraftContractTermsDocumentOptionsModel.SetDataProductID("b38df608-d34b-4d58-8136-ed25e6c6684e")
				deleteDraftContractTermsDocumentOptionsModel.SetDraftID("2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				deleteDraftContractTermsDocumentOptionsModel.SetContractTermsID("598183cd-b910-4e8d-9a97-97097afda3c1@e4fe2f87-0e56-46dd-b3b8-e9af32309e82")
				deleteDraftContractTermsDocumentOptionsModel.SetDocumentID("testString")
				deleteDraftContractTermsDocumentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteDraftContractTermsDocumentOptionsModel).ToNot(BeNil())
				Expect(deleteDraftContractTermsDocumentOptionsModel.DataProductID).To(Equal(core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")))
				Expect(deleteDraftContractTermsDocumentOptionsModel.DraftID).To(Equal(core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd")))
				Expect(deleteDraftContractTermsDocumentOptionsModel.ContractTermsID).To(Equal(core.StringPtr("598183cd-b910-4e8d-9a97-97097afda3c1@e4fe2f87-0e56-46dd-b3b8-e9af32309e82")))
				Expect(deleteDraftContractTermsDocumentOptionsModel.DocumentID).To(Equal(core.StringPtr("testString")))
				Expect(deleteDraftContractTermsDocumentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeliveryMethod successfully`, func() {
				id := "09cf5fcc-cb9d-4995-a8e4-16517b25229f"
				var container *dpxv1.ContainerReference = nil
				_, err := dpxService.NewDeliveryMethod(id, container)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewDomain successfully`, func() {
				id := "testString"
				_model, err := dpxService.NewDomain(id)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewGetDataProductDraftOptions successfully`, func() {
				// Construct an instance of the GetDataProductDraftOptions model
				dataProductID := "b38df608-d34b-4d58-8136-ed25e6c6684e"
				draftID := "2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd"
				getDataProductDraftOptionsModel := dpxService.NewGetDataProductDraftOptions(dataProductID, draftID)
				getDataProductDraftOptionsModel.SetDataProductID("b38df608-d34b-4d58-8136-ed25e6c6684e")
				getDataProductDraftOptionsModel.SetDraftID("2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				getDataProductDraftOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getDataProductDraftOptionsModel).ToNot(BeNil())
				Expect(getDataProductDraftOptionsModel.DataProductID).To(Equal(core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")))
				Expect(getDataProductDraftOptionsModel.DraftID).To(Equal(core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd")))
				Expect(getDataProductDraftOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetDataProductOptions successfully`, func() {
				// Construct an instance of the GetDataProductOptions model
				dataProductID := "b38df608-d34b-4d58-8136-ed25e6c6684e"
				getDataProductOptionsModel := dpxService.NewGetDataProductOptions(dataProductID)
				getDataProductOptionsModel.SetDataProductID("b38df608-d34b-4d58-8136-ed25e6c6684e")
				getDataProductOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getDataProductOptionsModel).ToNot(BeNil())
				Expect(getDataProductOptionsModel.DataProductID).To(Equal(core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")))
				Expect(getDataProductOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetDataProductReleaseOptions successfully`, func() {
				// Construct an instance of the GetDataProductReleaseOptions model
				dataProductID := "b38df608-d34b-4d58-8136-ed25e6c6684e"
				releaseID := "2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd"
				getDataProductReleaseOptionsModel := dpxService.NewGetDataProductReleaseOptions(dataProductID, releaseID)
				getDataProductReleaseOptionsModel.SetDataProductID("b38df608-d34b-4d58-8136-ed25e6c6684e")
				getDataProductReleaseOptionsModel.SetReleaseID("2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				getDataProductReleaseOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getDataProductReleaseOptionsModel).ToNot(BeNil())
				Expect(getDataProductReleaseOptionsModel.DataProductID).To(Equal(core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")))
				Expect(getDataProductReleaseOptionsModel.ReleaseID).To(Equal(core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd")))
				Expect(getDataProductReleaseOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetDraftContractTermsDocumentOptions successfully`, func() {
				// Construct an instance of the GetDraftContractTermsDocumentOptions model
				dataProductID := "b38df608-d34b-4d58-8136-ed25e6c6684e"
				draftID := "2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd"
				contractTermsID := "598183cd-b910-4e8d-9a97-97097afda3c1@e4fe2f87-0e56-46dd-b3b8-e9af32309e82"
				documentID := "testString"
				getDraftContractTermsDocumentOptionsModel := dpxService.NewGetDraftContractTermsDocumentOptions(dataProductID, draftID, contractTermsID, documentID)
				getDraftContractTermsDocumentOptionsModel.SetDataProductID("b38df608-d34b-4d58-8136-ed25e6c6684e")
				getDraftContractTermsDocumentOptionsModel.SetDraftID("2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				getDraftContractTermsDocumentOptionsModel.SetContractTermsID("598183cd-b910-4e8d-9a97-97097afda3c1@e4fe2f87-0e56-46dd-b3b8-e9af32309e82")
				getDraftContractTermsDocumentOptionsModel.SetDocumentID("testString")
				getDraftContractTermsDocumentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getDraftContractTermsDocumentOptionsModel).ToNot(BeNil())
				Expect(getDraftContractTermsDocumentOptionsModel.DataProductID).To(Equal(core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")))
				Expect(getDraftContractTermsDocumentOptionsModel.DraftID).To(Equal(core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd")))
				Expect(getDraftContractTermsDocumentOptionsModel.ContractTermsID).To(Equal(core.StringPtr("598183cd-b910-4e8d-9a97-97097afda3c1@e4fe2f87-0e56-46dd-b3b8-e9af32309e82")))
				Expect(getDraftContractTermsDocumentOptionsModel.DocumentID).To(Equal(core.StringPtr("testString")))
				Expect(getDraftContractTermsDocumentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetInitializeStatusOptions successfully`, func() {
				// Construct an instance of the GetInitializeStatusOptions model
				getInitializeStatusOptionsModel := dpxService.NewGetInitializeStatusOptions()
				getInitializeStatusOptionsModel.SetContainerID("testString")
				getInitializeStatusOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getInitializeStatusOptionsModel).ToNot(BeNil())
				Expect(getInitializeStatusOptionsModel.ContainerID).To(Equal(core.StringPtr("testString")))
				Expect(getInitializeStatusOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetReleaseContractTermsDocumentOptions successfully`, func() {
				// Construct an instance of the GetReleaseContractTermsDocumentOptions model
				dataProductID := "b38df608-d34b-4d58-8136-ed25e6c6684e"
				releaseID := "2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd"
				contractTermsID := "598183cd-b910-4e8d-9a97-97097afda3c1@e4fe2f87-0e56-46dd-b3b8-e9af32309e82"
				documentID := "testString"
				getReleaseContractTermsDocumentOptionsModel := dpxService.NewGetReleaseContractTermsDocumentOptions(dataProductID, releaseID, contractTermsID, documentID)
				getReleaseContractTermsDocumentOptionsModel.SetDataProductID("b38df608-d34b-4d58-8136-ed25e6c6684e")
				getReleaseContractTermsDocumentOptionsModel.SetReleaseID("2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				getReleaseContractTermsDocumentOptionsModel.SetContractTermsID("598183cd-b910-4e8d-9a97-97097afda3c1@e4fe2f87-0e56-46dd-b3b8-e9af32309e82")
				getReleaseContractTermsDocumentOptionsModel.SetDocumentID("testString")
				getReleaseContractTermsDocumentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getReleaseContractTermsDocumentOptionsModel).ToNot(BeNil())
				Expect(getReleaseContractTermsDocumentOptionsModel.DataProductID).To(Equal(core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")))
				Expect(getReleaseContractTermsDocumentOptionsModel.ReleaseID).To(Equal(core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd")))
				Expect(getReleaseContractTermsDocumentOptionsModel.ContractTermsID).To(Equal(core.StringPtr("598183cd-b910-4e8d-9a97-97097afda3c1@e4fe2f87-0e56-46dd-b3b8-e9af32309e82")))
				Expect(getReleaseContractTermsDocumentOptionsModel.DocumentID).To(Equal(core.StringPtr("testString")))
				Expect(getReleaseContractTermsDocumentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewInitializeOptions successfully`, func() {
				// Construct an instance of the ContainerReference model
				containerReferenceModel := new(dpxv1.ContainerReference)
				Expect(containerReferenceModel).ToNot(BeNil())
				containerReferenceModel.ID = core.StringPtr("d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				containerReferenceModel.Type = core.StringPtr("catalog")
				Expect(containerReferenceModel.ID).To(Equal(core.StringPtr("d29c42eb-7100-4b7a-8257-c196dbcca1cd")))
				Expect(containerReferenceModel.Type).To(Equal(core.StringPtr("catalog")))

				// Construct an instance of the InitializeOptions model
				initializeOptionsModel := dpxService.NewInitializeOptions()
				initializeOptionsModel.SetContainer(containerReferenceModel)
				initializeOptionsModel.SetInclude([]string{"delivery_methods", "data_product_samples", "domains_multi_industry"})
				initializeOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(initializeOptionsModel).ToNot(BeNil())
				Expect(initializeOptionsModel.Container).To(Equal(containerReferenceModel))
				Expect(initializeOptionsModel.Include).To(Equal([]string{"delivery_methods", "data_product_samples", "domains_multi_industry"}))
				Expect(initializeOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewJSONPatchOperation successfully`, func() {
				op := "add"
				path := "testString"
				_model, err := dpxService.NewJSONPatchOperation(op, path)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewListDataProductDraftsOptions successfully`, func() {
				// Construct an instance of the ListDataProductDraftsOptions model
				dataProductID := "b38df608-d34b-4d58-8136-ed25e6c6684e"
				listDataProductDraftsOptionsModel := dpxService.NewListDataProductDraftsOptions(dataProductID)
				listDataProductDraftsOptionsModel.SetDataProductID("b38df608-d34b-4d58-8136-ed25e6c6684e")
				listDataProductDraftsOptionsModel.SetAssetContainerID("testString")
				listDataProductDraftsOptionsModel.SetVersion("testString")
				listDataProductDraftsOptionsModel.SetLimit(int64(10))
				listDataProductDraftsOptionsModel.SetStart("testString")
				listDataProductDraftsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listDataProductDraftsOptionsModel).ToNot(BeNil())
				Expect(listDataProductDraftsOptionsModel.DataProductID).To(Equal(core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")))
				Expect(listDataProductDraftsOptionsModel.AssetContainerID).To(Equal(core.StringPtr("testString")))
				Expect(listDataProductDraftsOptionsModel.Version).To(Equal(core.StringPtr("testString")))
				Expect(listDataProductDraftsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(10))))
				Expect(listDataProductDraftsOptionsModel.Start).To(Equal(core.StringPtr("testString")))
				Expect(listDataProductDraftsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListDataProductReleasesOptions successfully`, func() {
				// Construct an instance of the ListDataProductReleasesOptions model
				dataProductID := "b38df608-d34b-4d58-8136-ed25e6c6684e"
				listDataProductReleasesOptionsModel := dpxService.NewListDataProductReleasesOptions(dataProductID)
				listDataProductReleasesOptionsModel.SetDataProductID("b38df608-d34b-4d58-8136-ed25e6c6684e")
				listDataProductReleasesOptionsModel.SetAssetContainerID("testString")
				listDataProductReleasesOptionsModel.SetState([]string{"available"})
				listDataProductReleasesOptionsModel.SetVersion("testString")
				listDataProductReleasesOptionsModel.SetLimit(int64(10))
				listDataProductReleasesOptionsModel.SetStart("testString")
				listDataProductReleasesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listDataProductReleasesOptionsModel).ToNot(BeNil())
				Expect(listDataProductReleasesOptionsModel.DataProductID).To(Equal(core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")))
				Expect(listDataProductReleasesOptionsModel.AssetContainerID).To(Equal(core.StringPtr("testString")))
				Expect(listDataProductReleasesOptionsModel.State).To(Equal([]string{"available"}))
				Expect(listDataProductReleasesOptionsModel.Version).To(Equal(core.StringPtr("testString")))
				Expect(listDataProductReleasesOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(10))))
				Expect(listDataProductReleasesOptionsModel.Start).To(Equal(core.StringPtr("testString")))
				Expect(listDataProductReleasesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListDataProductsOptions successfully`, func() {
				// Construct an instance of the ListDataProductsOptions model
				listDataProductsOptionsModel := dpxService.NewListDataProductsOptions()
				listDataProductsOptionsModel.SetLimit(int64(10))
				listDataProductsOptionsModel.SetStart("testString")
				listDataProductsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listDataProductsOptionsModel).ToNot(BeNil())
				Expect(listDataProductsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(10))))
				Expect(listDataProductsOptionsModel.Start).To(Equal(core.StringPtr("testString")))
				Expect(listDataProductsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewManageApiKeysOptions successfully`, func() {
				// Construct an instance of the ManageApiKeysOptions model
				manageApiKeysOptionsModel := dpxService.NewManageApiKeysOptions()
				manageApiKeysOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(manageApiKeysOptionsModel).ToNot(BeNil())
				Expect(manageApiKeysOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPublishDataProductDraftOptions successfully`, func() {
				// Construct an instance of the PublishDataProductDraftOptions model
				dataProductID := "b38df608-d34b-4d58-8136-ed25e6c6684e"
				draftID := "2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd"
				publishDataProductDraftOptionsModel := dpxService.NewPublishDataProductDraftOptions(dataProductID, draftID)
				publishDataProductDraftOptionsModel.SetDataProductID("b38df608-d34b-4d58-8136-ed25e6c6684e")
				publishDataProductDraftOptionsModel.SetDraftID("2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				publishDataProductDraftOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(publishDataProductDraftOptionsModel).ToNot(BeNil())
				Expect(publishDataProductDraftOptionsModel.DataProductID).To(Equal(core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")))
				Expect(publishDataProductDraftOptionsModel.DraftID).To(Equal(core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd")))
				Expect(publishDataProductDraftOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewRetireDataProductReleaseOptions successfully`, func() {
				// Construct an instance of the RetireDataProductReleaseOptions model
				dataProductID := "b38df608-d34b-4d58-8136-ed25e6c6684e"
				releaseID := "2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd"
				retireDataProductReleaseOptionsModel := dpxService.NewRetireDataProductReleaseOptions(dataProductID, releaseID)
				retireDataProductReleaseOptionsModel.SetDataProductID("b38df608-d34b-4d58-8136-ed25e6c6684e")
				retireDataProductReleaseOptionsModel.SetReleaseID("2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				retireDataProductReleaseOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(retireDataProductReleaseOptionsModel).ToNot(BeNil())
				Expect(retireDataProductReleaseOptionsModel.DataProductID).To(Equal(core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")))
				Expect(retireDataProductReleaseOptionsModel.ReleaseID).To(Equal(core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd")))
				Expect(retireDataProductReleaseOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateDataProductDraftOptions successfully`, func() {
				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(dpxv1.JSONPatchOperation)
				Expect(jsonPatchOperationModel).ToNot(BeNil())
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = "testString"
				Expect(jsonPatchOperationModel.Op).To(Equal(core.StringPtr("add")))
				Expect(jsonPatchOperationModel.Path).To(Equal(core.StringPtr("testString")))
				Expect(jsonPatchOperationModel.From).To(Equal(core.StringPtr("testString")))
				Expect(jsonPatchOperationModel.Value).To(Equal("testString"))

				// Construct an instance of the UpdateDataProductDraftOptions model
				dataProductID := "b38df608-d34b-4d58-8136-ed25e6c6684e"
				draftID := "2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd"
				jsonPatchInstructions := []dpxv1.JSONPatchOperation{}
				updateDataProductDraftOptionsModel := dpxService.NewUpdateDataProductDraftOptions(dataProductID, draftID, jsonPatchInstructions)
				updateDataProductDraftOptionsModel.SetDataProductID("b38df608-d34b-4d58-8136-ed25e6c6684e")
				updateDataProductDraftOptionsModel.SetDraftID("2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				updateDataProductDraftOptionsModel.SetJSONPatchInstructions([]dpxv1.JSONPatchOperation{*jsonPatchOperationModel})
				updateDataProductDraftOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateDataProductDraftOptionsModel).ToNot(BeNil())
				Expect(updateDataProductDraftOptionsModel.DataProductID).To(Equal(core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")))
				Expect(updateDataProductDraftOptionsModel.DraftID).To(Equal(core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd")))
				Expect(updateDataProductDraftOptionsModel.JSONPatchInstructions).To(Equal([]dpxv1.JSONPatchOperation{*jsonPatchOperationModel}))
				Expect(updateDataProductDraftOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateDataProductReleaseOptions successfully`, func() {
				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(dpxv1.JSONPatchOperation)
				Expect(jsonPatchOperationModel).ToNot(BeNil())
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = "testString"
				Expect(jsonPatchOperationModel.Op).To(Equal(core.StringPtr("add")))
				Expect(jsonPatchOperationModel.Path).To(Equal(core.StringPtr("testString")))
				Expect(jsonPatchOperationModel.From).To(Equal(core.StringPtr("testString")))
				Expect(jsonPatchOperationModel.Value).To(Equal("testString"))

				// Construct an instance of the UpdateDataProductReleaseOptions model
				dataProductID := "b38df608-d34b-4d58-8136-ed25e6c6684e"
				releaseID := "2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd"
				jsonPatchInstructions := []dpxv1.JSONPatchOperation{}
				updateDataProductReleaseOptionsModel := dpxService.NewUpdateDataProductReleaseOptions(dataProductID, releaseID, jsonPatchInstructions)
				updateDataProductReleaseOptionsModel.SetDataProductID("b38df608-d34b-4d58-8136-ed25e6c6684e")
				updateDataProductReleaseOptionsModel.SetReleaseID("2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				updateDataProductReleaseOptionsModel.SetJSONPatchInstructions([]dpxv1.JSONPatchOperation{*jsonPatchOperationModel})
				updateDataProductReleaseOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateDataProductReleaseOptionsModel).ToNot(BeNil())
				Expect(updateDataProductReleaseOptionsModel.DataProductID).To(Equal(core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")))
				Expect(updateDataProductReleaseOptionsModel.ReleaseID).To(Equal(core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd")))
				Expect(updateDataProductReleaseOptionsModel.JSONPatchInstructions).To(Equal([]dpxv1.JSONPatchOperation{*jsonPatchOperationModel}))
				Expect(updateDataProductReleaseOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateDraftContractTermsDocumentOptions successfully`, func() {
				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(dpxv1.JSONPatchOperation)
				Expect(jsonPatchOperationModel).ToNot(BeNil())
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = "testString"
				Expect(jsonPatchOperationModel.Op).To(Equal(core.StringPtr("add")))
				Expect(jsonPatchOperationModel.Path).To(Equal(core.StringPtr("testString")))
				Expect(jsonPatchOperationModel.From).To(Equal(core.StringPtr("testString")))
				Expect(jsonPatchOperationModel.Value).To(Equal("testString"))

				// Construct an instance of the UpdateDraftContractTermsDocumentOptions model
				dataProductID := "b38df608-d34b-4d58-8136-ed25e6c6684e"
				draftID := "2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd"
				contractTermsID := "598183cd-b910-4e8d-9a97-97097afda3c1@e4fe2f87-0e56-46dd-b3b8-e9af32309e82"
				documentID := "testString"
				jsonPatchInstructions := []dpxv1.JSONPatchOperation{}
				updateDraftContractTermsDocumentOptionsModel := dpxService.NewUpdateDraftContractTermsDocumentOptions(dataProductID, draftID, contractTermsID, documentID, jsonPatchInstructions)
				updateDraftContractTermsDocumentOptionsModel.SetDataProductID("b38df608-d34b-4d58-8136-ed25e6c6684e")
				updateDraftContractTermsDocumentOptionsModel.SetDraftID("2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				updateDraftContractTermsDocumentOptionsModel.SetContractTermsID("598183cd-b910-4e8d-9a97-97097afda3c1@e4fe2f87-0e56-46dd-b3b8-e9af32309e82")
				updateDraftContractTermsDocumentOptionsModel.SetDocumentID("testString")
				updateDraftContractTermsDocumentOptionsModel.SetJSONPatchInstructions([]dpxv1.JSONPatchOperation{*jsonPatchOperationModel})
				updateDraftContractTermsDocumentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateDraftContractTermsDocumentOptionsModel).ToNot(BeNil())
				Expect(updateDraftContractTermsDocumentOptionsModel.DataProductID).To(Equal(core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")))
				Expect(updateDraftContractTermsDocumentOptionsModel.DraftID).To(Equal(core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd")))
				Expect(updateDraftContractTermsDocumentOptionsModel.ContractTermsID).To(Equal(core.StringPtr("598183cd-b910-4e8d-9a97-97097afda3c1@e4fe2f87-0e56-46dd-b3b8-e9af32309e82")))
				Expect(updateDraftContractTermsDocumentOptionsModel.DocumentID).To(Equal(core.StringPtr("testString")))
				Expect(updateDraftContractTermsDocumentOptionsModel.JSONPatchInstructions).To(Equal([]dpxv1.JSONPatchOperation{*jsonPatchOperationModel}))
				Expect(updateDraftContractTermsDocumentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUseCase successfully`, func() {
				id := "testString"
				_model, err := dpxService.NewUseCase(id)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
		})
	})
	Describe(`Model unmarshaling tests`, func() {
		It(`Invoke UnmarshalAssetPartReference successfully`, func() {
			// Construct an instance of the model.
			model := new(dpxv1.AssetPartReference)
			model.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
			model.Container = nil
			model.Type = core.StringPtr("data_asset")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *dpxv1.AssetPartReference
			err = dpxv1.UnmarshalAssetPartReference(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalAssetReference successfully`, func() {
			// Construct an instance of the model.
			model := new(dpxv1.AssetReference)
			model.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
			model.Container = nil

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *dpxv1.AssetReference
			err = dpxv1.UnmarshalAssetReference(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalContainerReference successfully`, func() {
			// Construct an instance of the model.
			model := new(dpxv1.ContainerReference)
			model.ID = core.StringPtr("d29c42eb-7100-4b7a-8257-c196dbcca1cd")
			model.Type = core.StringPtr("catalog")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *dpxv1.ContainerReference
			err = dpxv1.UnmarshalContainerReference(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalContractTermsDocument successfully`, func() {
			// Construct an instance of the model.
			model := new(dpxv1.ContractTermsDocument)
			model.URL = core.StringPtr("testString")
			model.Type = core.StringPtr("terms_and_conditions")
			model.Name = core.StringPtr("testString")
			model.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
			model.Attachment = nil
			model.UploadURL = core.StringPtr("testString")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *dpxv1.ContractTermsDocument
			err = dpxv1.UnmarshalContractTermsDocument(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalContractTermsDocumentAttachment successfully`, func() {
			// Construct an instance of the model.
			model := new(dpxv1.ContractTermsDocumentAttachment)
			model.ID = core.StringPtr("testString")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *dpxv1.ContractTermsDocumentAttachment
			err = dpxv1.UnmarshalContractTermsDocumentAttachment(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalDataProductContractTerms successfully`, func() {
			// Construct an instance of the model.
			model := new(dpxv1.DataProductContractTerms)
			model.Asset = nil
			model.ID = core.StringPtr("testString")
			model.Documents = nil

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *dpxv1.DataProductContractTerms
			err = dpxv1.UnmarshalDataProductContractTerms(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalDataProductIdentity successfully`, func() {
			// Construct an instance of the model.
			model := new(dpxv1.DataProductIdentity)
			model.ID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *dpxv1.DataProductIdentity
			err = dpxv1.UnmarshalDataProductIdentity(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalDataProductPart successfully`, func() {
			// Construct an instance of the model.
			model := new(dpxv1.DataProductPart)
			model.Asset = nil
			model.Revision = core.Int64Ptr(int64(1))
			model.UpdatedAt = CreateMockDateTime("2023-07-01T22:22:34.876Z")
			model.DeliveryMethods = nil

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *dpxv1.DataProductPart
			err = dpxv1.UnmarshalDataProductPart(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalDataProductVersionPrototype successfully`, func() {
			// Construct an instance of the model.
			model := new(dpxv1.DataProductVersionPrototype)
			model.Version = core.StringPtr("1.0.0")
			model.State = core.StringPtr("draft")
			model.DataProduct = nil
			model.Name = core.StringPtr("My Data Product")
			model.Description = core.StringPtr("This is a description of My Data Product.")
			model.Asset = nil
			model.Tags = []string{"testString"}
			model.UseCases = nil
			model.Domain = nil
			model.Types = []string{"data"}
			model.PartsOut = nil
			model.ContractTerms = nil
			model.IsRestricted = core.BoolPtr(true)

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *dpxv1.DataProductVersionPrototype
			err = dpxv1.UnmarshalDataProductVersionPrototype(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalDeliveryMethod successfully`, func() {
			// Construct an instance of the model.
			model := new(dpxv1.DeliveryMethod)
			model.ID = core.StringPtr("09cf5fcc-cb9d-4995-a8e4-16517b25229f")
			model.Container = nil

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *dpxv1.DeliveryMethod
			err = dpxv1.UnmarshalDeliveryMethod(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalDomain successfully`, func() {
			// Construct an instance of the model.
			model := new(dpxv1.Domain)
			model.ID = core.StringPtr("testString")
			model.Name = core.StringPtr("testString")
			model.Container = nil

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *dpxv1.Domain
			err = dpxv1.UnmarshalDomain(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalJSONPatchOperation successfully`, func() {
			// Construct an instance of the model.
			model := new(dpxv1.JSONPatchOperation)
			model.Op = core.StringPtr("add")
			model.Path = core.StringPtr("testString")
			model.From = core.StringPtr("testString")
			model.Value = "testString"

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *dpxv1.JSONPatchOperation
			err = dpxv1.UnmarshalJSONPatchOperation(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalUseCase successfully`, func() {
			// Construct an instance of the model.
			model := new(dpxv1.UseCase)
			model.ID = core.StringPtr("testString")
			model.Name = core.StringPtr("testString")
			model.Container = nil

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *dpxv1.UseCase
			err = dpxv1.UnmarshalUseCase(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
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
	ba := []byte(mockData)
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
