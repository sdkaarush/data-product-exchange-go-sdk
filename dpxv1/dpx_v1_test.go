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
	"bytes"
	"context"
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
				"DPX_URL": "https://dpxv1/api",
				"DPX_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				dpxService, serviceErr := dpxv1.NewDpxV1UsingExternalConfig(&dpxv1.DpxV1Options{
				})
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
				dpxService, serviceErr := dpxv1.NewDpxV1UsingExternalConfig(&dpxv1.DpxV1Options{
				})
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
				"DPX_URL": "https://dpxv1/api",
				"DPX_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			dpxService, serviceErr := dpxv1.NewDpxV1UsingExternalConfig(&dpxv1.DpxV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(dpxService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"DPX_AUTH_TYPE":   "NOAuth",
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
					fmt.Fprintf(res, "%s", `{"container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}, "href": "http://api.example.com/configuration/initialize/status?catalog_id=d29c42eb-7100-4b7a-8257-c196dbcca1cd", "status": "not_started", "trace": "Trace", "errors": [{"code": "request_body_error", "message": "Message", "extra": {"anyKey": "anyValue"}, "more_info": "MoreInfo"}], "last_started_at": "2023-08-21T15:24:06.021Z", "last_finished_at": "2023-08-21T20:24:34.450Z", "initialized_options": [{"name": "Name", "version": 1}]}`)
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
					fmt.Fprintf(res, "%s", `{"container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}, "href": "http://api.example.com/configuration/initialize/status?catalog_id=d29c42eb-7100-4b7a-8257-c196dbcca1cd", "status": "not_started", "trace": "Trace", "errors": [{"code": "request_body_error", "message": "Message", "extra": {"anyKey": "anyValue"}, "more_info": "MoreInfo"}], "last_started_at": "2023-08-21T15:24:06.021Z", "last_finished_at": "2023-08-21T20:24:34.450Z", "initialized_options": [{"name": "Name", "version": 1}]}`)
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
				initializeOptionsModel.Force = core.BoolPtr(true)
				initializeOptionsModel.Reinitialize = core.BoolPtr(true)
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
					fmt.Fprintf(res, "%s", `{"container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}, "href": "http://api.example.com/configuration/initialize/status?catalog_id=d29c42eb-7100-4b7a-8257-c196dbcca1cd", "status": "not_started", "trace": "Trace", "errors": [{"code": "request_body_error", "message": "Message", "extra": {"anyKey": "anyValue"}, "more_info": "MoreInfo"}], "last_started_at": "2023-08-21T15:24:06.021Z", "last_finished_at": "2023-08-21T20:24:34.450Z", "initialized_options": [{"name": "Name", "version": 1}]}`)
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
				initializeOptionsModel.Force = core.BoolPtr(true)
				initializeOptionsModel.Reinitialize = core.BoolPtr(true)
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
					fmt.Fprintf(res, "%s", `{"container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}, "href": "http://api.example.com/configuration/initialize/status?catalog_id=d29c42eb-7100-4b7a-8257-c196dbcca1cd", "status": "not_started", "trace": "Trace", "errors": [{"code": "request_body_error", "message": "Message", "extra": {"anyKey": "anyValue"}, "more_info": "MoreInfo"}], "last_started_at": "2023-08-21T15:24:06.021Z", "last_finished_at": "2023-08-21T20:24:34.450Z", "initialized_options": [{"name": "Name", "version": 1}]}`)
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
				initializeOptionsModel.Force = core.BoolPtr(true)
				initializeOptionsModel.Reinitialize = core.BoolPtr(true)
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
				initializeOptionsModel.Force = core.BoolPtr(true)
				initializeOptionsModel.Reinitialize = core.BoolPtr(true)
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
				initializeOptionsModel.Force = core.BoolPtr(true)
				initializeOptionsModel.Reinitialize = core.BoolPtr(true)
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
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Construct an instance of the GetDataProductOptions model
				getDataProductOptionsModel := new(dpxv1.GetDataProductOptions)
				getDataProductOptionsModel.ID = core.StringPtr("testString")
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
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())
				dpxService.EnableRetries(0, 0)

				// Construct an instance of the GetDataProductOptions model
				getDataProductOptionsModel := new(dpxv1.GetDataProductOptions)
				getDataProductOptionsModel.ID = core.StringPtr("testString")
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
					fmt.Fprintf(res, "%s", `{"id": "b38df608-d34b-4d58-8136-ed25e6c6684e", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}, "name": "Sample Data Product"}`)
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
				getDataProductOptionsModel.ID = core.StringPtr("testString")
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
				getDataProductOptionsModel.ID = core.StringPtr("testString")
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
				getDataProductOptionsModel.ID = core.StringPtr("testString")
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
					fmt.Fprintf(res, "%s", `{"limit": 200, "first": {"href": "http://api.example.com/collection"}, "next": {"href": "http://api.example.com/collection?start=eyJvZmZzZXQiOjAsImRvbmUiOnRydWV9", "start": "eyJvZmZzZXQiOjAsImRvbmUiOnRydWV9"}, "data_products": [{"id": "b38df608-d34b-4d58-8136-ed25e6c6684e", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}, "name": "Sample Data Product"}]}`)
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
					fmt.Fprintf(res, "%s", `{"limit": 200, "first": {"href": "http://api.example.com/collection"}, "next": {"href": "http://api.example.com/collection?start=eyJvZmZzZXQiOjAsImRvbmUiOnRydWV9", "start": "eyJvZmZzZXQiOjAsImRvbmUiOnRydWV9"}, "data_products": [{"id": "b38df608-d34b-4d58-8136-ed25e6c6684e", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}, "name": "Sample Data Product"}]}`)
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
				responseObject := new(dpxv1.DataProductCollection)
				nextObject := new(dpxv1.NextPage)
				nextObject.Start = core.StringPtr("abc-123")
				responseObject.Next = nextObject

				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(Equal(core.StringPtr("abc-123")))
			})
			It(`Invoke GetNextStart without a "Next" property in the response`, func() {
				responseObject := new(dpxv1.DataProductCollection)

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

				var allResults []dpxv1.DataProduct
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
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Construct an instance of the ListDataProductVersionsOptions model
				listDataProductVersionsOptionsModel := new(dpxv1.ListDataProductVersionsOptions)
				listDataProductVersionsOptionsModel.AssetContainerID = core.StringPtr("testString")
				listDataProductVersionsOptionsModel.DataProduct = core.StringPtr("testString")
				listDataProductVersionsOptionsModel.State = core.StringPtr("draft")
				listDataProductVersionsOptionsModel.Version = core.StringPtr("testString")
				listDataProductVersionsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listDataProductVersionsOptionsModel.Start = core.StringPtr("testString")
				listDataProductVersionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := dpxService.ListDataProductVersions(listDataProductVersionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				dpxService.EnableRetries(0, 0)
				result, response, operationErr = dpxService.ListDataProductVersions(listDataProductVersionsOptionsModel)
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
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())
				dpxService.EnableRetries(0, 0)

				// Construct an instance of the ListDataProductVersionsOptions model
				listDataProductVersionsOptionsModel := new(dpxv1.ListDataProductVersionsOptions)
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
				_, _, operationErr := dpxService.ListDataProductVersionsWithContext(ctx, listDataProductVersionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				dpxService.DisableRetries()
				result, response, operationErr := dpxService.ListDataProductVersions(listDataProductVersionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = dpxService.ListDataProductVersionsWithContext(ctx, listDataProductVersionsOptionsModel)
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
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := dpxService.ListDataProductVersions(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListDataProductVersionsOptions model
				listDataProductVersionsOptionsModel := new(dpxv1.ListDataProductVersionsOptions)
				listDataProductVersionsOptionsModel.AssetContainerID = core.StringPtr("testString")
				listDataProductVersionsOptionsModel.DataProduct = core.StringPtr("testString")
				listDataProductVersionsOptionsModel.State = core.StringPtr("draft")
				listDataProductVersionsOptionsModel.Version = core.StringPtr("testString")
				listDataProductVersionsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listDataProductVersionsOptionsModel.Start = core.StringPtr("testString")
				listDataProductVersionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = dpxService.ListDataProductVersions(listDataProductVersionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListDataProductVersions with error: Operation request error`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Construct an instance of the ListDataProductVersionsOptions model
				listDataProductVersionsOptionsModel := new(dpxv1.ListDataProductVersionsOptions)
				listDataProductVersionsOptionsModel.AssetContainerID = core.StringPtr("testString")
				listDataProductVersionsOptionsModel.DataProduct = core.StringPtr("testString")
				listDataProductVersionsOptionsModel.State = core.StringPtr("draft")
				listDataProductVersionsOptionsModel.Version = core.StringPtr("testString")
				listDataProductVersionsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listDataProductVersionsOptionsModel.Start = core.StringPtr("testString")
				listDataProductVersionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dpxService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := dpxService.ListDataProductVersions(listDataProductVersionsOptionsModel)
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
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Construct an instance of the ListDataProductVersionsOptions model
				listDataProductVersionsOptionsModel := new(dpxv1.ListDataProductVersionsOptions)
				listDataProductVersionsOptionsModel.AssetContainerID = core.StringPtr("testString")
				listDataProductVersionsOptionsModel.DataProduct = core.StringPtr("testString")
				listDataProductVersionsOptionsModel.State = core.StringPtr("draft")
				listDataProductVersionsOptionsModel.Version = core.StringPtr("testString")
				listDataProductVersionsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listDataProductVersionsOptionsModel.Start = core.StringPtr("testString")
				listDataProductVersionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := dpxService.ListDataProductVersions(listDataProductVersionsOptionsModel)
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
				responseObject := new(dpxv1.DataProductVersionCollection)
				nextObject := new(dpxv1.NextPage)
				nextObject.Start = core.StringPtr("abc-123")
				responseObject.Next = nextObject

				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(Equal(core.StringPtr("abc-123")))
			})
			It(`Invoke GetNextStart without a "Next" property in the response`, func() {
				responseObject := new(dpxv1.DataProductVersionCollection)

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
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				listDataProductVersionsOptionsModel := &dpxv1.ListDataProductVersionsOptions{
					AssetContainerID: core.StringPtr("testString"),
					DataProduct: core.StringPtr("testString"),
					State: core.StringPtr("draft"),
					Version: core.StringPtr("testString"),
					Limit: core.Int64Ptr(int64(10)),
				}

				pager, err := dpxService.NewDataProductVersionsPager(listDataProductVersionsOptionsModel)
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
			It(`Use DataProductVersionsPager.GetAll successfully`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				listDataProductVersionsOptionsModel := &dpxv1.ListDataProductVersionsOptions{
					AssetContainerID: core.StringPtr("testString"),
					DataProduct: core.StringPtr("testString"),
					State: core.StringPtr("draft"),
					Version: core.StringPtr("testString"),
					Limit: core.Int64Ptr(int64(10)),
				}

				pager, err := dpxService.NewDataProductVersionsPager(listDataProductVersionsOptionsModel)
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

				// Construct an instance of the AssetReference model
				assetReferenceModel := new(dpxv1.AssetReference)
				assetReferenceModel.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				assetReferenceModel.Container = containerReferenceModel

				// Construct an instance of the DataProductContractTerms model
				dataProductContractTermsModel := new(dpxv1.DataProductContractTerms)
				dataProductContractTermsModel.ID = core.StringPtr("testString")
				dataProductContractTermsModel.Documents = []dpxv1.ContractTermsDocument{*contractTermsDocumentModel}
				dataProductContractTermsModel.Asset = assetReferenceModel

				// Construct an instance of the CreateDataProductVersionOptions model
				createDataProductVersionOptionsModel := new(dpxv1.CreateDataProductVersionOptions)
				createDataProductVersionOptionsModel.Container = containerReferenceModel
				createDataProductVersionOptionsModel.Version = core.StringPtr("testString")
				createDataProductVersionOptionsModel.State = core.StringPtr("draft")
				createDataProductVersionOptionsModel.DataProduct = dataProductIdentityModel
				createDataProductVersionOptionsModel.Name = core.StringPtr("My New Data Product")
				createDataProductVersionOptionsModel.Description = core.StringPtr("testString")
				createDataProductVersionOptionsModel.Tags = []string{"testString"}
				createDataProductVersionOptionsModel.UseCases = []dpxv1.UseCase{*useCaseModel}
				createDataProductVersionOptionsModel.Domain = domainModel
				createDataProductVersionOptionsModel.Type = []string{"data"}
				createDataProductVersionOptionsModel.PartsOut = []dpxv1.DataProductPart{*dataProductPartModel}
				createDataProductVersionOptionsModel.ContractTerms = []dpxv1.DataProductContractTerms{*dataProductContractTermsModel}
				createDataProductVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := dpxService.CreateDataProductVersion(createDataProductVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				dpxService.EnableRetries(0, 0)
				result, response, operationErr = dpxService.CreateDataProductVersion(createDataProductVersionOptionsModel)
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
					fmt.Fprintf(res, "%s", `{"version": "1.0.0", "state": "draft", "data_product": {"id": "b38df608-d34b-4d58-8136-ed25e6c6684e"}, "name": "My Data Product", "description": "This is a description of My Data Product.", "id": "2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd", "asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "tags": ["Tags"], "use_cases": [{"id": "ID", "name": "Name", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}], "domain": {"id": "ID", "name": "Name", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "type": ["data"], "parts_out": [{"asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}, "type": "data_asset"}, "revision": 1, "updated_at": "2023-07-01T22:22:34.876Z", "delivery_methods": [{"id": "09cf5fcc-cb9d-4995-a8e4-16517b25229f", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}]}], "published_by": "PublishedBy", "published_at": "2019-01-01T12:00:00.000Z", "contract_terms": [{"id": "ID", "documents": [{"url": "URL", "type": "terms_and_conditions", "name": "Name", "id": "2b0bf220-079c-11ee-be56-0242ac120002", "attachment": {"id": "ID"}}], "asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}}], "created_by": "CreatedBy", "created_at": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke CreateDataProductVersion successfully with retries`, func() {
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

				// Construct an instance of the AssetReference model
				assetReferenceModel := new(dpxv1.AssetReference)
				assetReferenceModel.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				assetReferenceModel.Container = containerReferenceModel

				// Construct an instance of the DataProductContractTerms model
				dataProductContractTermsModel := new(dpxv1.DataProductContractTerms)
				dataProductContractTermsModel.ID = core.StringPtr("testString")
				dataProductContractTermsModel.Documents = []dpxv1.ContractTermsDocument{*contractTermsDocumentModel}
				dataProductContractTermsModel.Asset = assetReferenceModel

				// Construct an instance of the CreateDataProductVersionOptions model
				createDataProductVersionOptionsModel := new(dpxv1.CreateDataProductVersionOptions)
				createDataProductVersionOptionsModel.Container = containerReferenceModel
				createDataProductVersionOptionsModel.Version = core.StringPtr("testString")
				createDataProductVersionOptionsModel.State = core.StringPtr("draft")
				createDataProductVersionOptionsModel.DataProduct = dataProductIdentityModel
				createDataProductVersionOptionsModel.Name = core.StringPtr("My New Data Product")
				createDataProductVersionOptionsModel.Description = core.StringPtr("testString")
				createDataProductVersionOptionsModel.Tags = []string{"testString"}
				createDataProductVersionOptionsModel.UseCases = []dpxv1.UseCase{*useCaseModel}
				createDataProductVersionOptionsModel.Domain = domainModel
				createDataProductVersionOptionsModel.Type = []string{"data"}
				createDataProductVersionOptionsModel.PartsOut = []dpxv1.DataProductPart{*dataProductPartModel}
				createDataProductVersionOptionsModel.ContractTerms = []dpxv1.DataProductContractTerms{*dataProductContractTermsModel}
				createDataProductVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := dpxService.CreateDataProductVersionWithContext(ctx, createDataProductVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				dpxService.DisableRetries()
				result, response, operationErr := dpxService.CreateDataProductVersion(createDataProductVersionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = dpxService.CreateDataProductVersionWithContext(ctx, createDataProductVersionOptionsModel)
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
					fmt.Fprintf(res, "%s", `{"version": "1.0.0", "state": "draft", "data_product": {"id": "b38df608-d34b-4d58-8136-ed25e6c6684e"}, "name": "My Data Product", "description": "This is a description of My Data Product.", "id": "2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd", "asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "tags": ["Tags"], "use_cases": [{"id": "ID", "name": "Name", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}], "domain": {"id": "ID", "name": "Name", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "type": ["data"], "parts_out": [{"asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}, "type": "data_asset"}, "revision": 1, "updated_at": "2023-07-01T22:22:34.876Z", "delivery_methods": [{"id": "09cf5fcc-cb9d-4995-a8e4-16517b25229f", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}]}], "published_by": "PublishedBy", "published_at": "2019-01-01T12:00:00.000Z", "contract_terms": [{"id": "ID", "documents": [{"url": "URL", "type": "terms_and_conditions", "name": "Name", "id": "2b0bf220-079c-11ee-be56-0242ac120002", "attachment": {"id": "ID"}}], "asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}}], "created_by": "CreatedBy", "created_at": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke CreateDataProductVersion successfully`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := dpxService.CreateDataProductVersion(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ContainerReference model
				containerReferenceModel := new(dpxv1.ContainerReference)
				containerReferenceModel.ID = core.StringPtr("d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				containerReferenceModel.Type = core.StringPtr("catalog")

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

				// Construct an instance of the AssetReference model
				assetReferenceModel := new(dpxv1.AssetReference)
				assetReferenceModel.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				assetReferenceModel.Container = containerReferenceModel

				// Construct an instance of the DataProductContractTerms model
				dataProductContractTermsModel := new(dpxv1.DataProductContractTerms)
				dataProductContractTermsModel.ID = core.StringPtr("testString")
				dataProductContractTermsModel.Documents = []dpxv1.ContractTermsDocument{*contractTermsDocumentModel}
				dataProductContractTermsModel.Asset = assetReferenceModel

				// Construct an instance of the CreateDataProductVersionOptions model
				createDataProductVersionOptionsModel := new(dpxv1.CreateDataProductVersionOptions)
				createDataProductVersionOptionsModel.Container = containerReferenceModel
				createDataProductVersionOptionsModel.Version = core.StringPtr("testString")
				createDataProductVersionOptionsModel.State = core.StringPtr("draft")
				createDataProductVersionOptionsModel.DataProduct = dataProductIdentityModel
				createDataProductVersionOptionsModel.Name = core.StringPtr("My New Data Product")
				createDataProductVersionOptionsModel.Description = core.StringPtr("testString")
				createDataProductVersionOptionsModel.Tags = []string{"testString"}
				createDataProductVersionOptionsModel.UseCases = []dpxv1.UseCase{*useCaseModel}
				createDataProductVersionOptionsModel.Domain = domainModel
				createDataProductVersionOptionsModel.Type = []string{"data"}
				createDataProductVersionOptionsModel.PartsOut = []dpxv1.DataProductPart{*dataProductPartModel}
				createDataProductVersionOptionsModel.ContractTerms = []dpxv1.DataProductContractTerms{*dataProductContractTermsModel}
				createDataProductVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = dpxService.CreateDataProductVersion(createDataProductVersionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateDataProductVersion with error: Operation validation and request error`, func() {
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

				// Construct an instance of the AssetReference model
				assetReferenceModel := new(dpxv1.AssetReference)
				assetReferenceModel.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				assetReferenceModel.Container = containerReferenceModel

				// Construct an instance of the DataProductContractTerms model
				dataProductContractTermsModel := new(dpxv1.DataProductContractTerms)
				dataProductContractTermsModel.ID = core.StringPtr("testString")
				dataProductContractTermsModel.Documents = []dpxv1.ContractTermsDocument{*contractTermsDocumentModel}
				dataProductContractTermsModel.Asset = assetReferenceModel

				// Construct an instance of the CreateDataProductVersionOptions model
				createDataProductVersionOptionsModel := new(dpxv1.CreateDataProductVersionOptions)
				createDataProductVersionOptionsModel.Container = containerReferenceModel
				createDataProductVersionOptionsModel.Version = core.StringPtr("testString")
				createDataProductVersionOptionsModel.State = core.StringPtr("draft")
				createDataProductVersionOptionsModel.DataProduct = dataProductIdentityModel
				createDataProductVersionOptionsModel.Name = core.StringPtr("My New Data Product")
				createDataProductVersionOptionsModel.Description = core.StringPtr("testString")
				createDataProductVersionOptionsModel.Tags = []string{"testString"}
				createDataProductVersionOptionsModel.UseCases = []dpxv1.UseCase{*useCaseModel}
				createDataProductVersionOptionsModel.Domain = domainModel
				createDataProductVersionOptionsModel.Type = []string{"data"}
				createDataProductVersionOptionsModel.PartsOut = []dpxv1.DataProductPart{*dataProductPartModel}
				createDataProductVersionOptionsModel.ContractTerms = []dpxv1.DataProductContractTerms{*dataProductContractTermsModel}
				createDataProductVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dpxService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := dpxService.CreateDataProductVersion(createDataProductVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateDataProductVersionOptions model with no property values
				createDataProductVersionOptionsModelNew := new(dpxv1.CreateDataProductVersionOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = dpxService.CreateDataProductVersion(createDataProductVersionOptionsModelNew)
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

				// Construct an instance of the AssetReference model
				assetReferenceModel := new(dpxv1.AssetReference)
				assetReferenceModel.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				assetReferenceModel.Container = containerReferenceModel

				// Construct an instance of the DataProductContractTerms model
				dataProductContractTermsModel := new(dpxv1.DataProductContractTerms)
				dataProductContractTermsModel.ID = core.StringPtr("testString")
				dataProductContractTermsModel.Documents = []dpxv1.ContractTermsDocument{*contractTermsDocumentModel}
				dataProductContractTermsModel.Asset = assetReferenceModel

				// Construct an instance of the CreateDataProductVersionOptions model
				createDataProductVersionOptionsModel := new(dpxv1.CreateDataProductVersionOptions)
				createDataProductVersionOptionsModel.Container = containerReferenceModel
				createDataProductVersionOptionsModel.Version = core.StringPtr("testString")
				createDataProductVersionOptionsModel.State = core.StringPtr("draft")
				createDataProductVersionOptionsModel.DataProduct = dataProductIdentityModel
				createDataProductVersionOptionsModel.Name = core.StringPtr("My New Data Product")
				createDataProductVersionOptionsModel.Description = core.StringPtr("testString")
				createDataProductVersionOptionsModel.Tags = []string{"testString"}
				createDataProductVersionOptionsModel.UseCases = []dpxv1.UseCase{*useCaseModel}
				createDataProductVersionOptionsModel.Domain = domainModel
				createDataProductVersionOptionsModel.Type = []string{"data"}
				createDataProductVersionOptionsModel.PartsOut = []dpxv1.DataProductPart{*dataProductPartModel}
				createDataProductVersionOptionsModel.ContractTerms = []dpxv1.DataProductContractTerms{*dataProductContractTermsModel}
				createDataProductVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := dpxService.CreateDataProductVersion(createDataProductVersionOptionsModel)
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
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Construct an instance of the GetDataProductVersionOptions model
				getDataProductVersionOptionsModel := new(dpxv1.GetDataProductVersionOptions)
				getDataProductVersionOptionsModel.ID = core.StringPtr("testString")
				getDataProductVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := dpxService.GetDataProductVersion(getDataProductVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				dpxService.EnableRetries(0, 0)
				result, response, operationErr = dpxService.GetDataProductVersion(getDataProductVersionOptionsModel)
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
					fmt.Fprintf(res, "%s", `{"version": "1.0.0", "state": "draft", "data_product": {"id": "b38df608-d34b-4d58-8136-ed25e6c6684e"}, "name": "My Data Product", "description": "This is a description of My Data Product.", "id": "2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd", "asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "tags": ["Tags"], "use_cases": [{"id": "ID", "name": "Name", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}], "domain": {"id": "ID", "name": "Name", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "type": ["data"], "parts_out": [{"asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}, "type": "data_asset"}, "revision": 1, "updated_at": "2023-07-01T22:22:34.876Z", "delivery_methods": [{"id": "09cf5fcc-cb9d-4995-a8e4-16517b25229f", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}]}], "published_by": "PublishedBy", "published_at": "2019-01-01T12:00:00.000Z", "contract_terms": [{"id": "ID", "documents": [{"url": "URL", "type": "terms_and_conditions", "name": "Name", "id": "2b0bf220-079c-11ee-be56-0242ac120002", "attachment": {"id": "ID"}}], "asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}}], "created_by": "CreatedBy", "created_at": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke GetDataProductVersion successfully with retries`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())
				dpxService.EnableRetries(0, 0)

				// Construct an instance of the GetDataProductVersionOptions model
				getDataProductVersionOptionsModel := new(dpxv1.GetDataProductVersionOptions)
				getDataProductVersionOptionsModel.ID = core.StringPtr("testString")
				getDataProductVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := dpxService.GetDataProductVersionWithContext(ctx, getDataProductVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				dpxService.DisableRetries()
				result, response, operationErr := dpxService.GetDataProductVersion(getDataProductVersionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = dpxService.GetDataProductVersionWithContext(ctx, getDataProductVersionOptionsModel)
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
					fmt.Fprintf(res, "%s", `{"version": "1.0.0", "state": "draft", "data_product": {"id": "b38df608-d34b-4d58-8136-ed25e6c6684e"}, "name": "My Data Product", "description": "This is a description of My Data Product.", "id": "2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd", "asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "tags": ["Tags"], "use_cases": [{"id": "ID", "name": "Name", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}], "domain": {"id": "ID", "name": "Name", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "type": ["data"], "parts_out": [{"asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}, "type": "data_asset"}, "revision": 1, "updated_at": "2023-07-01T22:22:34.876Z", "delivery_methods": [{"id": "09cf5fcc-cb9d-4995-a8e4-16517b25229f", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}]}], "published_by": "PublishedBy", "published_at": "2019-01-01T12:00:00.000Z", "contract_terms": [{"id": "ID", "documents": [{"url": "URL", "type": "terms_and_conditions", "name": "Name", "id": "2b0bf220-079c-11ee-be56-0242ac120002", "attachment": {"id": "ID"}}], "asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}}], "created_by": "CreatedBy", "created_at": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke GetDataProductVersion successfully`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := dpxService.GetDataProductVersion(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetDataProductVersionOptions model
				getDataProductVersionOptionsModel := new(dpxv1.GetDataProductVersionOptions)
				getDataProductVersionOptionsModel.ID = core.StringPtr("testString")
				getDataProductVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = dpxService.GetDataProductVersion(getDataProductVersionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetDataProductVersion with error: Operation validation and request error`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Construct an instance of the GetDataProductVersionOptions model
				getDataProductVersionOptionsModel := new(dpxv1.GetDataProductVersionOptions)
				getDataProductVersionOptionsModel.ID = core.StringPtr("testString")
				getDataProductVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dpxService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := dpxService.GetDataProductVersion(getDataProductVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetDataProductVersionOptions model with no property values
				getDataProductVersionOptionsModelNew := new(dpxv1.GetDataProductVersionOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = dpxService.GetDataProductVersion(getDataProductVersionOptionsModelNew)
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
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Construct an instance of the GetDataProductVersionOptions model
				getDataProductVersionOptionsModel := new(dpxv1.GetDataProductVersionOptions)
				getDataProductVersionOptionsModel.ID = core.StringPtr("testString")
				getDataProductVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := dpxService.GetDataProductVersion(getDataProductVersionOptionsModel)
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
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := dpxService.DeleteDataProductVersion(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteDataProductVersionOptions model
				deleteDataProductVersionOptionsModel := new(dpxv1.DeleteDataProductVersionOptions)
				deleteDataProductVersionOptionsModel.ID = core.StringPtr("testString")
				deleteDataProductVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = dpxService.DeleteDataProductVersion(deleteDataProductVersionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteDataProductVersion with error: Operation validation and request error`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Construct an instance of the DeleteDataProductVersionOptions model
				deleteDataProductVersionOptionsModel := new(dpxv1.DeleteDataProductVersionOptions)
				deleteDataProductVersionOptionsModel.ID = core.StringPtr("testString")
				deleteDataProductVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dpxService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := dpxService.DeleteDataProductVersion(deleteDataProductVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteDataProductVersionOptions model with no property values
				deleteDataProductVersionOptionsModelNew := new(dpxv1.DeleteDataProductVersionOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = dpxService.DeleteDataProductVersion(deleteDataProductVersionOptionsModelNew)
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
				jsonPatchOperationModel.Value = core.StringPtr("testString")

				// Construct an instance of the UpdateDataProductVersionOptions model
				updateDataProductVersionOptionsModel := new(dpxv1.UpdateDataProductVersionOptions)
				updateDataProductVersionOptionsModel.ID = core.StringPtr("testString")
				updateDataProductVersionOptionsModel.JSONPatchInstructions = []dpxv1.JSONPatchOperation{*jsonPatchOperationModel}
				updateDataProductVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := dpxService.UpdateDataProductVersion(updateDataProductVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				dpxService.EnableRetries(0, 0)
				result, response, operationErr = dpxService.UpdateDataProductVersion(updateDataProductVersionOptionsModel)
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
					fmt.Fprintf(res, "%s", `{"version": "1.0.0", "state": "draft", "data_product": {"id": "b38df608-d34b-4d58-8136-ed25e6c6684e"}, "name": "My Data Product", "description": "This is a description of My Data Product.", "id": "2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd", "asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "tags": ["Tags"], "use_cases": [{"id": "ID", "name": "Name", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}], "domain": {"id": "ID", "name": "Name", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "type": ["data"], "parts_out": [{"asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}, "type": "data_asset"}, "revision": 1, "updated_at": "2023-07-01T22:22:34.876Z", "delivery_methods": [{"id": "09cf5fcc-cb9d-4995-a8e4-16517b25229f", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}]}], "published_by": "PublishedBy", "published_at": "2019-01-01T12:00:00.000Z", "contract_terms": [{"id": "ID", "documents": [{"url": "URL", "type": "terms_and_conditions", "name": "Name", "id": "2b0bf220-079c-11ee-be56-0242ac120002", "attachment": {"id": "ID"}}], "asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}}], "created_by": "CreatedBy", "created_at": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke UpdateDataProductVersion successfully with retries`, func() {
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
				jsonPatchOperationModel.Value = core.StringPtr("testString")

				// Construct an instance of the UpdateDataProductVersionOptions model
				updateDataProductVersionOptionsModel := new(dpxv1.UpdateDataProductVersionOptions)
				updateDataProductVersionOptionsModel.ID = core.StringPtr("testString")
				updateDataProductVersionOptionsModel.JSONPatchInstructions = []dpxv1.JSONPatchOperation{*jsonPatchOperationModel}
				updateDataProductVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := dpxService.UpdateDataProductVersionWithContext(ctx, updateDataProductVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				dpxService.DisableRetries()
				result, response, operationErr := dpxService.UpdateDataProductVersion(updateDataProductVersionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = dpxService.UpdateDataProductVersionWithContext(ctx, updateDataProductVersionOptionsModel)
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
					fmt.Fprintf(res, "%s", `{"version": "1.0.0", "state": "draft", "data_product": {"id": "b38df608-d34b-4d58-8136-ed25e6c6684e"}, "name": "My Data Product", "description": "This is a description of My Data Product.", "id": "2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd", "asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "tags": ["Tags"], "use_cases": [{"id": "ID", "name": "Name", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}], "domain": {"id": "ID", "name": "Name", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "type": ["data"], "parts_out": [{"asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}, "type": "data_asset"}, "revision": 1, "updated_at": "2023-07-01T22:22:34.876Z", "delivery_methods": [{"id": "09cf5fcc-cb9d-4995-a8e4-16517b25229f", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}]}], "published_by": "PublishedBy", "published_at": "2019-01-01T12:00:00.000Z", "contract_terms": [{"id": "ID", "documents": [{"url": "URL", "type": "terms_and_conditions", "name": "Name", "id": "2b0bf220-079c-11ee-be56-0242ac120002", "attachment": {"id": "ID"}}], "asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}}], "created_by": "CreatedBy", "created_at": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke UpdateDataProductVersion successfully`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := dpxService.UpdateDataProductVersion(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(dpxv1.JSONPatchOperation)
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = core.StringPtr("testString")

				// Construct an instance of the UpdateDataProductVersionOptions model
				updateDataProductVersionOptionsModel := new(dpxv1.UpdateDataProductVersionOptions)
				updateDataProductVersionOptionsModel.ID = core.StringPtr("testString")
				updateDataProductVersionOptionsModel.JSONPatchInstructions = []dpxv1.JSONPatchOperation{*jsonPatchOperationModel}
				updateDataProductVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = dpxService.UpdateDataProductVersion(updateDataProductVersionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateDataProductVersion with error: Operation validation and request error`, func() {
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
				jsonPatchOperationModel.Value = core.StringPtr("testString")

				// Construct an instance of the UpdateDataProductVersionOptions model
				updateDataProductVersionOptionsModel := new(dpxv1.UpdateDataProductVersionOptions)
				updateDataProductVersionOptionsModel.ID = core.StringPtr("testString")
				updateDataProductVersionOptionsModel.JSONPatchInstructions = []dpxv1.JSONPatchOperation{*jsonPatchOperationModel}
				updateDataProductVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dpxService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := dpxService.UpdateDataProductVersion(updateDataProductVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateDataProductVersionOptions model with no property values
				updateDataProductVersionOptionsModelNew := new(dpxv1.UpdateDataProductVersionOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = dpxService.UpdateDataProductVersion(updateDataProductVersionOptionsModelNew)
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
				jsonPatchOperationModel.Value = core.StringPtr("testString")

				// Construct an instance of the UpdateDataProductVersionOptions model
				updateDataProductVersionOptionsModel := new(dpxv1.UpdateDataProductVersionOptions)
				updateDataProductVersionOptionsModel.ID = core.StringPtr("testString")
				updateDataProductVersionOptionsModel.JSONPatchInstructions = []dpxv1.JSONPatchOperation{*jsonPatchOperationModel}
				updateDataProductVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := dpxService.UpdateDataProductVersion(updateDataProductVersionOptionsModel)
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
	Describe(`CompleteContractTermsDocument(completeContractTermsDocumentOptions *CompleteContractTermsDocumentOptions) - Operation response error`, func() {
		completeContractTermsDocumentPath := "/data_product_exchange/v1/data_product_versions/testString/contract_terms/testString/documents/testString/complete"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(completeContractTermsDocumentPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CompleteContractTermsDocument with error: Operation response processing error`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Construct an instance of the CompleteContractTermsDocumentOptions model
				completeContractTermsDocumentOptionsModel := new(dpxv1.CompleteContractTermsDocumentOptions)
				completeContractTermsDocumentOptionsModel.DataProductVersionID = core.StringPtr("testString")
				completeContractTermsDocumentOptionsModel.ContractTermsID = core.StringPtr("testString")
				completeContractTermsDocumentOptionsModel.DocumentID = core.StringPtr("testString")
				completeContractTermsDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := dpxService.CompleteContractTermsDocument(completeContractTermsDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				dpxService.EnableRetries(0, 0)
				result, response, operationErr = dpxService.CompleteContractTermsDocument(completeContractTermsDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CompleteContractTermsDocument(completeContractTermsDocumentOptions *CompleteContractTermsDocumentOptions)`, func() {
		completeContractTermsDocumentPath := "/data_product_exchange/v1/data_product_versions/testString/contract_terms/testString/documents/testString/complete"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(completeContractTermsDocumentPath))
					Expect(req.Method).To(Equal("POST"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"url": "URL", "type": "terms_and_conditions", "name": "Name", "id": "2b0bf220-079c-11ee-be56-0242ac120002", "attachment": {"id": "ID"}}`)
				}))
			})
			It(`Invoke CompleteContractTermsDocument successfully with retries`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())
				dpxService.EnableRetries(0, 0)

				// Construct an instance of the CompleteContractTermsDocumentOptions model
				completeContractTermsDocumentOptionsModel := new(dpxv1.CompleteContractTermsDocumentOptions)
				completeContractTermsDocumentOptionsModel.DataProductVersionID = core.StringPtr("testString")
				completeContractTermsDocumentOptionsModel.ContractTermsID = core.StringPtr("testString")
				completeContractTermsDocumentOptionsModel.DocumentID = core.StringPtr("testString")
				completeContractTermsDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := dpxService.CompleteContractTermsDocumentWithContext(ctx, completeContractTermsDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				dpxService.DisableRetries()
				result, response, operationErr := dpxService.CompleteContractTermsDocument(completeContractTermsDocumentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = dpxService.CompleteContractTermsDocumentWithContext(ctx, completeContractTermsDocumentOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(completeContractTermsDocumentPath))
					Expect(req.Method).To(Equal("POST"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"url": "URL", "type": "terms_and_conditions", "name": "Name", "id": "2b0bf220-079c-11ee-be56-0242ac120002", "attachment": {"id": "ID"}}`)
				}))
			})
			It(`Invoke CompleteContractTermsDocument successfully`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := dpxService.CompleteContractTermsDocument(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CompleteContractTermsDocumentOptions model
				completeContractTermsDocumentOptionsModel := new(dpxv1.CompleteContractTermsDocumentOptions)
				completeContractTermsDocumentOptionsModel.DataProductVersionID = core.StringPtr("testString")
				completeContractTermsDocumentOptionsModel.ContractTermsID = core.StringPtr("testString")
				completeContractTermsDocumentOptionsModel.DocumentID = core.StringPtr("testString")
				completeContractTermsDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = dpxService.CompleteContractTermsDocument(completeContractTermsDocumentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CompleteContractTermsDocument with error: Operation validation and request error`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Construct an instance of the CompleteContractTermsDocumentOptions model
				completeContractTermsDocumentOptionsModel := new(dpxv1.CompleteContractTermsDocumentOptions)
				completeContractTermsDocumentOptionsModel.DataProductVersionID = core.StringPtr("testString")
				completeContractTermsDocumentOptionsModel.ContractTermsID = core.StringPtr("testString")
				completeContractTermsDocumentOptionsModel.DocumentID = core.StringPtr("testString")
				completeContractTermsDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dpxService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := dpxService.CompleteContractTermsDocument(completeContractTermsDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CompleteContractTermsDocumentOptions model with no property values
				completeContractTermsDocumentOptionsModelNew := new(dpxv1.CompleteContractTermsDocumentOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = dpxService.CompleteContractTermsDocument(completeContractTermsDocumentOptionsModelNew)
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
			It(`Invoke CompleteContractTermsDocument successfully`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Construct an instance of the CompleteContractTermsDocumentOptions model
				completeContractTermsDocumentOptionsModel := new(dpxv1.CompleteContractTermsDocumentOptions)
				completeContractTermsDocumentOptionsModel.DataProductVersionID = core.StringPtr("testString")
				completeContractTermsDocumentOptionsModel.ContractTermsID = core.StringPtr("testString")
				completeContractTermsDocumentOptionsModel.DocumentID = core.StringPtr("testString")
				completeContractTermsDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := dpxService.CompleteContractTermsDocument(completeContractTermsDocumentOptionsModel)
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
	Describe(`CreateContractTermsDocument(createContractTermsDocumentOptions *CreateContractTermsDocumentOptions) - Operation response error`, func() {
		createContractTermsDocumentPath := "/data_product_exchange/v1/data_product_versions/testString/contract_terms/testString/documents"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createContractTermsDocumentPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateContractTermsDocument with error: Operation response processing error`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Construct an instance of the ContractTermsDocumentAttachment model
				contractTermsDocumentAttachmentModel := new(dpxv1.ContractTermsDocumentAttachment)
				contractTermsDocumentAttachmentModel.ID = core.StringPtr("testString")

				// Construct an instance of the CreateContractTermsDocumentOptions model
				createContractTermsDocumentOptionsModel := new(dpxv1.CreateContractTermsDocumentOptions)
				createContractTermsDocumentOptionsModel.DataProductVersionID = core.StringPtr("testString")
				createContractTermsDocumentOptionsModel.ContractTermsID = core.StringPtr("testString")
				createContractTermsDocumentOptionsModel.Type = core.StringPtr("terms_and_conditions")
				createContractTermsDocumentOptionsModel.Name = core.StringPtr("Terms and conditions document")
				createContractTermsDocumentOptionsModel.ID = core.StringPtr("testString")
				createContractTermsDocumentOptionsModel.URL = core.StringPtr("testString")
				createContractTermsDocumentOptionsModel.Attachment = contractTermsDocumentAttachmentModel
				createContractTermsDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := dpxService.CreateContractTermsDocument(createContractTermsDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				dpxService.EnableRetries(0, 0)
				result, response, operationErr = dpxService.CreateContractTermsDocument(createContractTermsDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateContractTermsDocument(createContractTermsDocumentOptions *CreateContractTermsDocumentOptions)`, func() {
		createContractTermsDocumentPath := "/data_product_exchange/v1/data_product_versions/testString/contract_terms/testString/documents"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createContractTermsDocumentPath))
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
					fmt.Fprintf(res, "%s", `{"url": "URL", "type": "terms_and_conditions", "name": "Name", "id": "2b0bf220-079c-11ee-be56-0242ac120002", "attachment": {"id": "ID"}}`)
				}))
			})
			It(`Invoke CreateContractTermsDocument successfully with retries`, func() {
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

				// Construct an instance of the CreateContractTermsDocumentOptions model
				createContractTermsDocumentOptionsModel := new(dpxv1.CreateContractTermsDocumentOptions)
				createContractTermsDocumentOptionsModel.DataProductVersionID = core.StringPtr("testString")
				createContractTermsDocumentOptionsModel.ContractTermsID = core.StringPtr("testString")
				createContractTermsDocumentOptionsModel.Type = core.StringPtr("terms_and_conditions")
				createContractTermsDocumentOptionsModel.Name = core.StringPtr("Terms and conditions document")
				createContractTermsDocumentOptionsModel.ID = core.StringPtr("testString")
				createContractTermsDocumentOptionsModel.URL = core.StringPtr("testString")
				createContractTermsDocumentOptionsModel.Attachment = contractTermsDocumentAttachmentModel
				createContractTermsDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := dpxService.CreateContractTermsDocumentWithContext(ctx, createContractTermsDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				dpxService.DisableRetries()
				result, response, operationErr := dpxService.CreateContractTermsDocument(createContractTermsDocumentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = dpxService.CreateContractTermsDocumentWithContext(ctx, createContractTermsDocumentOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createContractTermsDocumentPath))
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
					fmt.Fprintf(res, "%s", `{"url": "URL", "type": "terms_and_conditions", "name": "Name", "id": "2b0bf220-079c-11ee-be56-0242ac120002", "attachment": {"id": "ID"}}`)
				}))
			})
			It(`Invoke CreateContractTermsDocument successfully`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := dpxService.CreateContractTermsDocument(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ContractTermsDocumentAttachment model
				contractTermsDocumentAttachmentModel := new(dpxv1.ContractTermsDocumentAttachment)
				contractTermsDocumentAttachmentModel.ID = core.StringPtr("testString")

				// Construct an instance of the CreateContractTermsDocumentOptions model
				createContractTermsDocumentOptionsModel := new(dpxv1.CreateContractTermsDocumentOptions)
				createContractTermsDocumentOptionsModel.DataProductVersionID = core.StringPtr("testString")
				createContractTermsDocumentOptionsModel.ContractTermsID = core.StringPtr("testString")
				createContractTermsDocumentOptionsModel.Type = core.StringPtr("terms_and_conditions")
				createContractTermsDocumentOptionsModel.Name = core.StringPtr("Terms and conditions document")
				createContractTermsDocumentOptionsModel.ID = core.StringPtr("testString")
				createContractTermsDocumentOptionsModel.URL = core.StringPtr("testString")
				createContractTermsDocumentOptionsModel.Attachment = contractTermsDocumentAttachmentModel
				createContractTermsDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = dpxService.CreateContractTermsDocument(createContractTermsDocumentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateContractTermsDocument with error: Operation validation and request error`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Construct an instance of the ContractTermsDocumentAttachment model
				contractTermsDocumentAttachmentModel := new(dpxv1.ContractTermsDocumentAttachment)
				contractTermsDocumentAttachmentModel.ID = core.StringPtr("testString")

				// Construct an instance of the CreateContractTermsDocumentOptions model
				createContractTermsDocumentOptionsModel := new(dpxv1.CreateContractTermsDocumentOptions)
				createContractTermsDocumentOptionsModel.DataProductVersionID = core.StringPtr("testString")
				createContractTermsDocumentOptionsModel.ContractTermsID = core.StringPtr("testString")
				createContractTermsDocumentOptionsModel.Type = core.StringPtr("terms_and_conditions")
				createContractTermsDocumentOptionsModel.Name = core.StringPtr("Terms and conditions document")
				createContractTermsDocumentOptionsModel.ID = core.StringPtr("testString")
				createContractTermsDocumentOptionsModel.URL = core.StringPtr("testString")
				createContractTermsDocumentOptionsModel.Attachment = contractTermsDocumentAttachmentModel
				createContractTermsDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dpxService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := dpxService.CreateContractTermsDocument(createContractTermsDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateContractTermsDocumentOptions model with no property values
				createContractTermsDocumentOptionsModelNew := new(dpxv1.CreateContractTermsDocumentOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = dpxService.CreateContractTermsDocument(createContractTermsDocumentOptionsModelNew)
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
			It(`Invoke CreateContractTermsDocument successfully`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Construct an instance of the ContractTermsDocumentAttachment model
				contractTermsDocumentAttachmentModel := new(dpxv1.ContractTermsDocumentAttachment)
				contractTermsDocumentAttachmentModel.ID = core.StringPtr("testString")

				// Construct an instance of the CreateContractTermsDocumentOptions model
				createContractTermsDocumentOptionsModel := new(dpxv1.CreateContractTermsDocumentOptions)
				createContractTermsDocumentOptionsModel.DataProductVersionID = core.StringPtr("testString")
				createContractTermsDocumentOptionsModel.ContractTermsID = core.StringPtr("testString")
				createContractTermsDocumentOptionsModel.Type = core.StringPtr("terms_and_conditions")
				createContractTermsDocumentOptionsModel.Name = core.StringPtr("Terms and conditions document")
				createContractTermsDocumentOptionsModel.ID = core.StringPtr("testString")
				createContractTermsDocumentOptionsModel.URL = core.StringPtr("testString")
				createContractTermsDocumentOptionsModel.Attachment = contractTermsDocumentAttachmentModel
				createContractTermsDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := dpxService.CreateContractTermsDocument(createContractTermsDocumentOptionsModel)
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
	Describe(`GetContractTermsDocument(getContractTermsDocumentOptions *GetContractTermsDocumentOptions) - Operation response error`, func() {
		getContractTermsDocumentPath := "/data_product_exchange/v1/data_product_versions/testString/contract_terms/testString/documents/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getContractTermsDocumentPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetContractTermsDocument with error: Operation response processing error`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Construct an instance of the GetContractTermsDocumentOptions model
				getContractTermsDocumentOptionsModel := new(dpxv1.GetContractTermsDocumentOptions)
				getContractTermsDocumentOptionsModel.DataProductVersionID = core.StringPtr("testString")
				getContractTermsDocumentOptionsModel.ContractTermsID = core.StringPtr("testString")
				getContractTermsDocumentOptionsModel.DocumentID = core.StringPtr("testString")
				getContractTermsDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := dpxService.GetContractTermsDocument(getContractTermsDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				dpxService.EnableRetries(0, 0)
				result, response, operationErr = dpxService.GetContractTermsDocument(getContractTermsDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetContractTermsDocument(getContractTermsDocumentOptions *GetContractTermsDocumentOptions)`, func() {
		getContractTermsDocumentPath := "/data_product_exchange/v1/data_product_versions/testString/contract_terms/testString/documents/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getContractTermsDocumentPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"url": "URL", "type": "terms_and_conditions", "name": "Name", "id": "2b0bf220-079c-11ee-be56-0242ac120002", "attachment": {"id": "ID"}}`)
				}))
			})
			It(`Invoke GetContractTermsDocument successfully with retries`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())
				dpxService.EnableRetries(0, 0)

				// Construct an instance of the GetContractTermsDocumentOptions model
				getContractTermsDocumentOptionsModel := new(dpxv1.GetContractTermsDocumentOptions)
				getContractTermsDocumentOptionsModel.DataProductVersionID = core.StringPtr("testString")
				getContractTermsDocumentOptionsModel.ContractTermsID = core.StringPtr("testString")
				getContractTermsDocumentOptionsModel.DocumentID = core.StringPtr("testString")
				getContractTermsDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := dpxService.GetContractTermsDocumentWithContext(ctx, getContractTermsDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				dpxService.DisableRetries()
				result, response, operationErr := dpxService.GetContractTermsDocument(getContractTermsDocumentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = dpxService.GetContractTermsDocumentWithContext(ctx, getContractTermsDocumentOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getContractTermsDocumentPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"url": "URL", "type": "terms_and_conditions", "name": "Name", "id": "2b0bf220-079c-11ee-be56-0242ac120002", "attachment": {"id": "ID"}}`)
				}))
			})
			It(`Invoke GetContractTermsDocument successfully`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := dpxService.GetContractTermsDocument(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetContractTermsDocumentOptions model
				getContractTermsDocumentOptionsModel := new(dpxv1.GetContractTermsDocumentOptions)
				getContractTermsDocumentOptionsModel.DataProductVersionID = core.StringPtr("testString")
				getContractTermsDocumentOptionsModel.ContractTermsID = core.StringPtr("testString")
				getContractTermsDocumentOptionsModel.DocumentID = core.StringPtr("testString")
				getContractTermsDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = dpxService.GetContractTermsDocument(getContractTermsDocumentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetContractTermsDocument with error: Operation validation and request error`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Construct an instance of the GetContractTermsDocumentOptions model
				getContractTermsDocumentOptionsModel := new(dpxv1.GetContractTermsDocumentOptions)
				getContractTermsDocumentOptionsModel.DataProductVersionID = core.StringPtr("testString")
				getContractTermsDocumentOptionsModel.ContractTermsID = core.StringPtr("testString")
				getContractTermsDocumentOptionsModel.DocumentID = core.StringPtr("testString")
				getContractTermsDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dpxService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := dpxService.GetContractTermsDocument(getContractTermsDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetContractTermsDocumentOptions model with no property values
				getContractTermsDocumentOptionsModelNew := new(dpxv1.GetContractTermsDocumentOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = dpxService.GetContractTermsDocument(getContractTermsDocumentOptionsModelNew)
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
			It(`Invoke GetContractTermsDocument successfully`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Construct an instance of the GetContractTermsDocumentOptions model
				getContractTermsDocumentOptionsModel := new(dpxv1.GetContractTermsDocumentOptions)
				getContractTermsDocumentOptionsModel.DataProductVersionID = core.StringPtr("testString")
				getContractTermsDocumentOptionsModel.ContractTermsID = core.StringPtr("testString")
				getContractTermsDocumentOptionsModel.DocumentID = core.StringPtr("testString")
				getContractTermsDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := dpxService.GetContractTermsDocument(getContractTermsDocumentOptionsModel)
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
	Describe(`DeleteContractTermsDocument(deleteContractTermsDocumentOptions *DeleteContractTermsDocumentOptions)`, func() {
		deleteContractTermsDocumentPath := "/data_product_exchange/v1/data_product_versions/testString/contract_terms/testString/documents/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteContractTermsDocumentPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteContractTermsDocument successfully`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := dpxService.DeleteContractTermsDocument(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteContractTermsDocumentOptions model
				deleteContractTermsDocumentOptionsModel := new(dpxv1.DeleteContractTermsDocumentOptions)
				deleteContractTermsDocumentOptionsModel.DataProductVersionID = core.StringPtr("testString")
				deleteContractTermsDocumentOptionsModel.ContractTermsID = core.StringPtr("testString")
				deleteContractTermsDocumentOptionsModel.DocumentID = core.StringPtr("testString")
				deleteContractTermsDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = dpxService.DeleteContractTermsDocument(deleteContractTermsDocumentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteContractTermsDocument with error: Operation validation and request error`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Construct an instance of the DeleteContractTermsDocumentOptions model
				deleteContractTermsDocumentOptionsModel := new(dpxv1.DeleteContractTermsDocumentOptions)
				deleteContractTermsDocumentOptionsModel.DataProductVersionID = core.StringPtr("testString")
				deleteContractTermsDocumentOptionsModel.ContractTermsID = core.StringPtr("testString")
				deleteContractTermsDocumentOptionsModel.DocumentID = core.StringPtr("testString")
				deleteContractTermsDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dpxService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := dpxService.DeleteContractTermsDocument(deleteContractTermsDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteContractTermsDocumentOptions model with no property values
				deleteContractTermsDocumentOptionsModelNew := new(dpxv1.DeleteContractTermsDocumentOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = dpxService.DeleteContractTermsDocument(deleteContractTermsDocumentOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateContractTermsDocument(updateContractTermsDocumentOptions *UpdateContractTermsDocumentOptions) - Operation response error`, func() {
		updateContractTermsDocumentPath := "/data_product_exchange/v1/data_product_versions/testString/contract_terms/testString/documents/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateContractTermsDocumentPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateContractTermsDocument with error: Operation response processing error`, func() {
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
				jsonPatchOperationModel.Value = core.StringPtr("testString")

				// Construct an instance of the UpdateContractTermsDocumentOptions model
				updateContractTermsDocumentOptionsModel := new(dpxv1.UpdateContractTermsDocumentOptions)
				updateContractTermsDocumentOptionsModel.DataProductVersionID = core.StringPtr("testString")
				updateContractTermsDocumentOptionsModel.ContractTermsID = core.StringPtr("testString")
				updateContractTermsDocumentOptionsModel.DocumentID = core.StringPtr("testString")
				updateContractTermsDocumentOptionsModel.JSONPatchInstructions = []dpxv1.JSONPatchOperation{*jsonPatchOperationModel}
				updateContractTermsDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := dpxService.UpdateContractTermsDocument(updateContractTermsDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				dpxService.EnableRetries(0, 0)
				result, response, operationErr = dpxService.UpdateContractTermsDocument(updateContractTermsDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateContractTermsDocument(updateContractTermsDocumentOptions *UpdateContractTermsDocumentOptions)`, func() {
		updateContractTermsDocumentPath := "/data_product_exchange/v1/data_product_versions/testString/contract_terms/testString/documents/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateContractTermsDocumentPath))
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
					fmt.Fprintf(res, "%s", `{"url": "URL", "type": "terms_and_conditions", "name": "Name", "id": "2b0bf220-079c-11ee-be56-0242ac120002", "attachment": {"id": "ID"}}`)
				}))
			})
			It(`Invoke UpdateContractTermsDocument successfully with retries`, func() {
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
				jsonPatchOperationModel.Value = core.StringPtr("testString")

				// Construct an instance of the UpdateContractTermsDocumentOptions model
				updateContractTermsDocumentOptionsModel := new(dpxv1.UpdateContractTermsDocumentOptions)
				updateContractTermsDocumentOptionsModel.DataProductVersionID = core.StringPtr("testString")
				updateContractTermsDocumentOptionsModel.ContractTermsID = core.StringPtr("testString")
				updateContractTermsDocumentOptionsModel.DocumentID = core.StringPtr("testString")
				updateContractTermsDocumentOptionsModel.JSONPatchInstructions = []dpxv1.JSONPatchOperation{*jsonPatchOperationModel}
				updateContractTermsDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := dpxService.UpdateContractTermsDocumentWithContext(ctx, updateContractTermsDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				dpxService.DisableRetries()
				result, response, operationErr := dpxService.UpdateContractTermsDocument(updateContractTermsDocumentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = dpxService.UpdateContractTermsDocumentWithContext(ctx, updateContractTermsDocumentOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateContractTermsDocumentPath))
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
					fmt.Fprintf(res, "%s", `{"url": "URL", "type": "terms_and_conditions", "name": "Name", "id": "2b0bf220-079c-11ee-be56-0242ac120002", "attachment": {"id": "ID"}}`)
				}))
			})
			It(`Invoke UpdateContractTermsDocument successfully`, func() {
				dpxService, serviceErr := dpxv1.NewDpxV1(&dpxv1.DpxV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dpxService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := dpxService.UpdateContractTermsDocument(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(dpxv1.JSONPatchOperation)
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = core.StringPtr("testString")

				// Construct an instance of the UpdateContractTermsDocumentOptions model
				updateContractTermsDocumentOptionsModel := new(dpxv1.UpdateContractTermsDocumentOptions)
				updateContractTermsDocumentOptionsModel.DataProductVersionID = core.StringPtr("testString")
				updateContractTermsDocumentOptionsModel.ContractTermsID = core.StringPtr("testString")
				updateContractTermsDocumentOptionsModel.DocumentID = core.StringPtr("testString")
				updateContractTermsDocumentOptionsModel.JSONPatchInstructions = []dpxv1.JSONPatchOperation{*jsonPatchOperationModel}
				updateContractTermsDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = dpxService.UpdateContractTermsDocument(updateContractTermsDocumentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateContractTermsDocument with error: Operation validation and request error`, func() {
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
				jsonPatchOperationModel.Value = core.StringPtr("testString")

				// Construct an instance of the UpdateContractTermsDocumentOptions model
				updateContractTermsDocumentOptionsModel := new(dpxv1.UpdateContractTermsDocumentOptions)
				updateContractTermsDocumentOptionsModel.DataProductVersionID = core.StringPtr("testString")
				updateContractTermsDocumentOptionsModel.ContractTermsID = core.StringPtr("testString")
				updateContractTermsDocumentOptionsModel.DocumentID = core.StringPtr("testString")
				updateContractTermsDocumentOptionsModel.JSONPatchInstructions = []dpxv1.JSONPatchOperation{*jsonPatchOperationModel}
				updateContractTermsDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dpxService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := dpxService.UpdateContractTermsDocument(updateContractTermsDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateContractTermsDocumentOptions model with no property values
				updateContractTermsDocumentOptionsModelNew := new(dpxv1.UpdateContractTermsDocumentOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = dpxService.UpdateContractTermsDocument(updateContractTermsDocumentOptionsModelNew)
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
			It(`Invoke UpdateContractTermsDocument successfully`, func() {
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
				jsonPatchOperationModel.Value = core.StringPtr("testString")

				// Construct an instance of the UpdateContractTermsDocumentOptions model
				updateContractTermsDocumentOptionsModel := new(dpxv1.UpdateContractTermsDocumentOptions)
				updateContractTermsDocumentOptionsModel.DataProductVersionID = core.StringPtr("testString")
				updateContractTermsDocumentOptionsModel.ContractTermsID = core.StringPtr("testString")
				updateContractTermsDocumentOptionsModel.DocumentID = core.StringPtr("testString")
				updateContractTermsDocumentOptionsModel.JSONPatchInstructions = []dpxv1.JSONPatchOperation{*jsonPatchOperationModel}
				updateContractTermsDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := dpxService.UpdateContractTermsDocument(updateContractTermsDocumentOptionsModel)
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
				id := "2b0bf220-079c-11ee-be56-0242ac120002"
				var container *dpxv1.ContainerReference = nil
				_, err := dpxService.NewAssetPartReference(id, container)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewAssetReference successfully`, func() {
				id := "2b0bf220-079c-11ee-be56-0242ac120002"
				var container *dpxv1.ContainerReference = nil
				_, err := dpxService.NewAssetReference(id, container)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewCompleteContractTermsDocumentOptions successfully`, func() {
				// Construct an instance of the CompleteContractTermsDocumentOptions model
				dataProductVersionID := "testString"
				contractTermsID := "testString"
				documentID := "testString"
				completeContractTermsDocumentOptionsModel := dpxService.NewCompleteContractTermsDocumentOptions(dataProductVersionID, contractTermsID, documentID)
				completeContractTermsDocumentOptionsModel.SetDataProductVersionID("testString")
				completeContractTermsDocumentOptionsModel.SetContractTermsID("testString")
				completeContractTermsDocumentOptionsModel.SetDocumentID("testString")
				completeContractTermsDocumentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(completeContractTermsDocumentOptionsModel).ToNot(BeNil())
				Expect(completeContractTermsDocumentOptionsModel.DataProductVersionID).To(Equal(core.StringPtr("testString")))
				Expect(completeContractTermsDocumentOptionsModel.ContractTermsID).To(Equal(core.StringPtr("testString")))
				Expect(completeContractTermsDocumentOptionsModel.DocumentID).To(Equal(core.StringPtr("testString")))
				Expect(completeContractTermsDocumentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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

				contractTermsDocumentPatch := dpxService.NewContractTermsDocumentPatch(contractTermsDocument)
				Expect(contractTermsDocumentPatch).ToNot(BeNil())

				_path := func(op interface{}) string {
					return *op.(dpxv1.JSONPatchOperation).Path
				}
				Expect(contractTermsDocumentPatch).To(MatchAllElements(_path, Elements{
				"/url": MatchAllFields(Fields{
					"Op": PointTo(Equal(dpxv1.JSONPatchOperation_Op_Add)),
					"Path": PointTo(Equal("/url")),
					"From": BeNil(),
					"Value": Equal(contractTermsDocument.URL),
					}),
				"/type": MatchAllFields(Fields{
					"Op": PointTo(Equal(dpxv1.JSONPatchOperation_Op_Add)),
					"Path": PointTo(Equal("/type")),
					"From": BeNil(),
					"Value": Equal(contractTermsDocument.Type),
					}),
				"/name": MatchAllFields(Fields{
					"Op": PointTo(Equal(dpxv1.JSONPatchOperation_Op_Add)),
					"Path": PointTo(Equal("/name")),
					"From": BeNil(),
					"Value": Equal(contractTermsDocument.Name),
					}),
				"/id": MatchAllFields(Fields{
					"Op": PointTo(Equal(dpxv1.JSONPatchOperation_Op_Add)),
					"Path": PointTo(Equal("/id")),
					"From": BeNil(),
					"Value": Equal(contractTermsDocument.ID),
					}),
				"/attachment": MatchAllFields(Fields{
					"Op": PointTo(Equal(dpxv1.JSONPatchOperation_Op_Add)),
					"Path": PointTo(Equal("/attachment")),
					"From": BeNil(),
					"Value": Equal(contractTermsDocument.Attachment),
					}),
				}))
			})
			It(`Invoke NewCreateContractTermsDocumentOptions successfully`, func() {
				// Construct an instance of the ContractTermsDocumentAttachment model
				contractTermsDocumentAttachmentModel := new(dpxv1.ContractTermsDocumentAttachment)
				Expect(contractTermsDocumentAttachmentModel).ToNot(BeNil())
				contractTermsDocumentAttachmentModel.ID = core.StringPtr("testString")
				Expect(contractTermsDocumentAttachmentModel.ID).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the CreateContractTermsDocumentOptions model
				dataProductVersionID := "testString"
				contractTermsID := "testString"
				createContractTermsDocumentOptionsType := "terms_and_conditions"
				createContractTermsDocumentOptionsName := "Terms and conditions document"
				createContractTermsDocumentOptionsID := "testString"
				createContractTermsDocumentOptionsModel := dpxService.NewCreateContractTermsDocumentOptions(dataProductVersionID, contractTermsID, createContractTermsDocumentOptionsType, createContractTermsDocumentOptionsName, createContractTermsDocumentOptionsID)
				createContractTermsDocumentOptionsModel.SetDataProductVersionID("testString")
				createContractTermsDocumentOptionsModel.SetContractTermsID("testString")
				createContractTermsDocumentOptionsModel.SetType("terms_and_conditions")
				createContractTermsDocumentOptionsModel.SetName("Terms and conditions document")
				createContractTermsDocumentOptionsModel.SetID("testString")
				createContractTermsDocumentOptionsModel.SetURL("testString")
				createContractTermsDocumentOptionsModel.SetAttachment(contractTermsDocumentAttachmentModel)
				createContractTermsDocumentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createContractTermsDocumentOptionsModel).ToNot(BeNil())
				Expect(createContractTermsDocumentOptionsModel.DataProductVersionID).To(Equal(core.StringPtr("testString")))
				Expect(createContractTermsDocumentOptionsModel.ContractTermsID).To(Equal(core.StringPtr("testString")))
				Expect(createContractTermsDocumentOptionsModel.Type).To(Equal(core.StringPtr("terms_and_conditions")))
				Expect(createContractTermsDocumentOptionsModel.Name).To(Equal(core.StringPtr("Terms and conditions document")))
				Expect(createContractTermsDocumentOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(createContractTermsDocumentOptionsModel.URL).To(Equal(core.StringPtr("testString")))
				Expect(createContractTermsDocumentOptionsModel.Attachment).To(Equal(contractTermsDocumentAttachmentModel))
				Expect(createContractTermsDocumentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateDataProductVersionOptions successfully`, func() {
				// Construct an instance of the ContainerReference model
				containerReferenceModel := new(dpxv1.ContainerReference)
				Expect(containerReferenceModel).ToNot(BeNil())
				containerReferenceModel.ID = core.StringPtr("d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				containerReferenceModel.Type = core.StringPtr("catalog")
				Expect(containerReferenceModel.ID).To(Equal(core.StringPtr("d29c42eb-7100-4b7a-8257-c196dbcca1cd")))
				Expect(containerReferenceModel.Type).To(Equal(core.StringPtr("catalog")))

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
				Expect(contractTermsDocumentModel.URL).To(Equal(core.StringPtr("testString")))
				Expect(contractTermsDocumentModel.Type).To(Equal(core.StringPtr("terms_and_conditions")))
				Expect(contractTermsDocumentModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(contractTermsDocumentModel.ID).To(Equal(core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")))
				Expect(contractTermsDocumentModel.Attachment).To(Equal(contractTermsDocumentAttachmentModel))

				// Construct an instance of the AssetReference model
				assetReferenceModel := new(dpxv1.AssetReference)
				Expect(assetReferenceModel).ToNot(BeNil())
				assetReferenceModel.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				assetReferenceModel.Container = containerReferenceModel
				Expect(assetReferenceModel.ID).To(Equal(core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")))
				Expect(assetReferenceModel.Container).To(Equal(containerReferenceModel))

				// Construct an instance of the DataProductContractTerms model
				dataProductContractTermsModel := new(dpxv1.DataProductContractTerms)
				Expect(dataProductContractTermsModel).ToNot(BeNil())
				dataProductContractTermsModel.ID = core.StringPtr("testString")
				dataProductContractTermsModel.Documents = []dpxv1.ContractTermsDocument{*contractTermsDocumentModel}
				dataProductContractTermsModel.Asset = assetReferenceModel
				Expect(dataProductContractTermsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(dataProductContractTermsModel.Documents).To(Equal([]dpxv1.ContractTermsDocument{*contractTermsDocumentModel}))
				Expect(dataProductContractTermsModel.Asset).To(Equal(assetReferenceModel))

				// Construct an instance of the CreateDataProductVersionOptions model
				var createDataProductVersionOptionsContainer *dpxv1.ContainerReference = nil
				createDataProductVersionOptionsModel := dpxService.NewCreateDataProductVersionOptions(createDataProductVersionOptionsContainer)
				createDataProductVersionOptionsModel.SetContainer(containerReferenceModel)
				createDataProductVersionOptionsModel.SetVersion("testString")
				createDataProductVersionOptionsModel.SetState("draft")
				createDataProductVersionOptionsModel.SetDataProduct(dataProductIdentityModel)
				createDataProductVersionOptionsModel.SetName("My New Data Product")
				createDataProductVersionOptionsModel.SetDescription("testString")
				createDataProductVersionOptionsModel.SetTags([]string{"testString"})
				createDataProductVersionOptionsModel.SetUseCases([]dpxv1.UseCase{*useCaseModel})
				createDataProductVersionOptionsModel.SetDomain(domainModel)
				createDataProductVersionOptionsModel.SetType([]string{"data"})
				createDataProductVersionOptionsModel.SetPartsOut([]dpxv1.DataProductPart{*dataProductPartModel})
				createDataProductVersionOptionsModel.SetContractTerms([]dpxv1.DataProductContractTerms{*dataProductContractTermsModel})
				createDataProductVersionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createDataProductVersionOptionsModel).ToNot(BeNil())
				Expect(createDataProductVersionOptionsModel.Container).To(Equal(containerReferenceModel))
				Expect(createDataProductVersionOptionsModel.Version).To(Equal(core.StringPtr("testString")))
				Expect(createDataProductVersionOptionsModel.State).To(Equal(core.StringPtr("draft")))
				Expect(createDataProductVersionOptionsModel.DataProduct).To(Equal(dataProductIdentityModel))
				Expect(createDataProductVersionOptionsModel.Name).To(Equal(core.StringPtr("My New Data Product")))
				Expect(createDataProductVersionOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(createDataProductVersionOptionsModel.Tags).To(Equal([]string{"testString"}))
				Expect(createDataProductVersionOptionsModel.UseCases).To(Equal([]dpxv1.UseCase{*useCaseModel}))
				Expect(createDataProductVersionOptionsModel.Domain).To(Equal(domainModel))
				Expect(createDataProductVersionOptionsModel.Type).To(Equal([]string{"data"}))
				Expect(createDataProductVersionOptionsModel.PartsOut).To(Equal([]dpxv1.DataProductPart{*dataProductPartModel}))
				Expect(createDataProductVersionOptionsModel.ContractTerms).To(Equal([]dpxv1.DataProductContractTerms{*dataProductContractTermsModel}))
				Expect(createDataProductVersionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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

				// Construct an instance of the DataProductContractTerms model
				dataProductContractTermsModel := new(dpxv1.DataProductContractTerms)
				dataProductContractTermsModel.ID = core.StringPtr("testString")
				dataProductContractTermsModel.Documents = []dpxv1.ContractTermsDocument{*contractTermsDocumentModel}
				dataProductContractTermsModel.Asset = assetReferenceModel

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
				dataProductVersion.Type = []string{"data"}
				dataProductVersion.PartsOut = []dpxv1.DataProductPart{*dataProductPartModel}
				dataProductVersion.PublishedBy = core.StringPtr("testString")
				dataProductVersion.PublishedAt = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				dataProductVersion.ContractTerms = []dpxv1.DataProductContractTerms{*dataProductContractTermsModel}
				dataProductVersion.CreatedBy = core.StringPtr("testString")
				dataProductVersion.CreatedAt = CreateMockDateTime("2019-01-01T12:00:00.000Z")

				dataProductVersionPatch := dpxService.NewDataProductVersionPatch(dataProductVersion)
				Expect(dataProductVersionPatch).ToNot(BeNil())

				_path := func(op interface{}) string {
					return *op.(dpxv1.JSONPatchOperation).Path
				}
				Expect(dataProductVersionPatch).To(MatchAllElements(_path, Elements{
				"/version": MatchAllFields(Fields{
					"Op": PointTo(Equal(dpxv1.JSONPatchOperation_Op_Add)),
					"Path": PointTo(Equal("/version")),
					"From": BeNil(),
					"Value": Equal(dataProductVersion.Version),
					}),
				"/state": MatchAllFields(Fields{
					"Op": PointTo(Equal(dpxv1.JSONPatchOperation_Op_Add)),
					"Path": PointTo(Equal("/state")),
					"From": BeNil(),
					"Value": Equal(dataProductVersion.State),
					}),
				"/data_product": MatchAllFields(Fields{
					"Op": PointTo(Equal(dpxv1.JSONPatchOperation_Op_Add)),
					"Path": PointTo(Equal("/data_product")),
					"From": BeNil(),
					"Value": Equal(dataProductVersion.DataProduct),
					}),
				"/name": MatchAllFields(Fields{
					"Op": PointTo(Equal(dpxv1.JSONPatchOperation_Op_Add)),
					"Path": PointTo(Equal("/name")),
					"From": BeNil(),
					"Value": Equal(dataProductVersion.Name),
					}),
				"/description": MatchAllFields(Fields{
					"Op": PointTo(Equal(dpxv1.JSONPatchOperation_Op_Add)),
					"Path": PointTo(Equal("/description")),
					"From": BeNil(),
					"Value": Equal(dataProductVersion.Description),
					}),
				"/id": MatchAllFields(Fields{
					"Op": PointTo(Equal(dpxv1.JSONPatchOperation_Op_Add)),
					"Path": PointTo(Equal("/id")),
					"From": BeNil(),
					"Value": Equal(dataProductVersion.ID),
					}),
				"/asset": MatchAllFields(Fields{
					"Op": PointTo(Equal(dpxv1.JSONPatchOperation_Op_Add)),
					"Path": PointTo(Equal("/asset")),
					"From": BeNil(),
					"Value": Equal(dataProductVersion.Asset),
					}),
				"/tags": MatchAllFields(Fields{
					"Op": PointTo(Equal(dpxv1.JSONPatchOperation_Op_Add)),
					"Path": PointTo(Equal("/tags")),
					"From": BeNil(),
					"Value": Equal(dataProductVersion.Tags),
					}),
				"/use_cases": MatchAllFields(Fields{
					"Op": PointTo(Equal(dpxv1.JSONPatchOperation_Op_Add)),
					"Path": PointTo(Equal("/use_cases")),
					"From": BeNil(),
					"Value": Equal(dataProductVersion.UseCases),
					}),
				"/domain": MatchAllFields(Fields{
					"Op": PointTo(Equal(dpxv1.JSONPatchOperation_Op_Add)),
					"Path": PointTo(Equal("/domain")),
					"From": BeNil(),
					"Value": Equal(dataProductVersion.Domain),
					}),
				"/type": MatchAllFields(Fields{
					"Op": PointTo(Equal(dpxv1.JSONPatchOperation_Op_Add)),
					"Path": PointTo(Equal("/type")),
					"From": BeNil(),
					"Value": Equal(dataProductVersion.Type),
					}),
				"/parts_out": MatchAllFields(Fields{
					"Op": PointTo(Equal(dpxv1.JSONPatchOperation_Op_Add)),
					"Path": PointTo(Equal("/parts_out")),
					"From": BeNil(),
					"Value": Equal(dataProductVersion.PartsOut),
					}),
				"/published_by": MatchAllFields(Fields{
					"Op": PointTo(Equal(dpxv1.JSONPatchOperation_Op_Add)),
					"Path": PointTo(Equal("/published_by")),
					"From": BeNil(),
					"Value": Equal(dataProductVersion.PublishedBy),
					}),
				"/published_at": MatchAllFields(Fields{
					"Op": PointTo(Equal(dpxv1.JSONPatchOperation_Op_Add)),
					"Path": PointTo(Equal("/published_at")),
					"From": BeNil(),
					"Value": Equal(dataProductVersion.PublishedAt),
					}),
				"/contract_terms": MatchAllFields(Fields{
					"Op": PointTo(Equal(dpxv1.JSONPatchOperation_Op_Add)),
					"Path": PointTo(Equal("/contract_terms")),
					"From": BeNil(),
					"Value": Equal(dataProductVersion.ContractTerms),
					}),
				"/created_by": MatchAllFields(Fields{
					"Op": PointTo(Equal(dpxv1.JSONPatchOperation_Op_Add)),
					"Path": PointTo(Equal("/created_by")),
					"From": BeNil(),
					"Value": Equal(dataProductVersion.CreatedBy),
					}),
				"/created_at": MatchAllFields(Fields{
					"Op": PointTo(Equal(dpxv1.JSONPatchOperation_Op_Add)),
					"Path": PointTo(Equal("/created_at")),
					"From": BeNil(),
					"Value": Equal(dataProductVersion.CreatedAt),
					}),
				}))
			})
			It(`Invoke NewDeleteContractTermsDocumentOptions successfully`, func() {
				// Construct an instance of the DeleteContractTermsDocumentOptions model
				dataProductVersionID := "testString"
				contractTermsID := "testString"
				documentID := "testString"
				deleteContractTermsDocumentOptionsModel := dpxService.NewDeleteContractTermsDocumentOptions(dataProductVersionID, contractTermsID, documentID)
				deleteContractTermsDocumentOptionsModel.SetDataProductVersionID("testString")
				deleteContractTermsDocumentOptionsModel.SetContractTermsID("testString")
				deleteContractTermsDocumentOptionsModel.SetDocumentID("testString")
				deleteContractTermsDocumentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteContractTermsDocumentOptionsModel).ToNot(BeNil())
				Expect(deleteContractTermsDocumentOptionsModel.DataProductVersionID).To(Equal(core.StringPtr("testString")))
				Expect(deleteContractTermsDocumentOptionsModel.ContractTermsID).To(Equal(core.StringPtr("testString")))
				Expect(deleteContractTermsDocumentOptionsModel.DocumentID).To(Equal(core.StringPtr("testString")))
				Expect(deleteContractTermsDocumentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteDataProductVersionOptions successfully`, func() {
				// Construct an instance of the DeleteDataProductVersionOptions model
				id := "testString"
				deleteDataProductVersionOptionsModel := dpxService.NewDeleteDataProductVersionOptions(id)
				deleteDataProductVersionOptionsModel.SetID("testString")
				deleteDataProductVersionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteDataProductVersionOptionsModel).ToNot(BeNil())
				Expect(deleteDataProductVersionOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(deleteDataProductVersionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
			It(`Invoke NewGetContractTermsDocumentOptions successfully`, func() {
				// Construct an instance of the GetContractTermsDocumentOptions model
				dataProductVersionID := "testString"
				contractTermsID := "testString"
				documentID := "testString"
				getContractTermsDocumentOptionsModel := dpxService.NewGetContractTermsDocumentOptions(dataProductVersionID, contractTermsID, documentID)
				getContractTermsDocumentOptionsModel.SetDataProductVersionID("testString")
				getContractTermsDocumentOptionsModel.SetContractTermsID("testString")
				getContractTermsDocumentOptionsModel.SetDocumentID("testString")
				getContractTermsDocumentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getContractTermsDocumentOptionsModel).ToNot(BeNil())
				Expect(getContractTermsDocumentOptionsModel.DataProductVersionID).To(Equal(core.StringPtr("testString")))
				Expect(getContractTermsDocumentOptionsModel.ContractTermsID).To(Equal(core.StringPtr("testString")))
				Expect(getContractTermsDocumentOptionsModel.DocumentID).To(Equal(core.StringPtr("testString")))
				Expect(getContractTermsDocumentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetDataProductOptions successfully`, func() {
				// Construct an instance of the GetDataProductOptions model
				id := "testString"
				getDataProductOptionsModel := dpxService.NewGetDataProductOptions(id)
				getDataProductOptionsModel.SetID("testString")
				getDataProductOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getDataProductOptionsModel).ToNot(BeNil())
				Expect(getDataProductOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getDataProductOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetDataProductVersionOptions successfully`, func() {
				// Construct an instance of the GetDataProductVersionOptions model
				id := "testString"
				getDataProductVersionOptionsModel := dpxService.NewGetDataProductVersionOptions(id)
				getDataProductVersionOptionsModel.SetID("testString")
				getDataProductVersionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getDataProductVersionOptionsModel).ToNot(BeNil())
				Expect(getDataProductVersionOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getDataProductVersionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
				initializeOptionsModel.SetForce(true)
				initializeOptionsModel.SetReinitialize(true)
				initializeOptionsModel.SetInclude([]string{"delivery_methods", "data_product_samples", "domains_multi_industry"})
				initializeOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(initializeOptionsModel).ToNot(BeNil())
				Expect(initializeOptionsModel.Container).To(Equal(containerReferenceModel))
				Expect(initializeOptionsModel.Force).To(Equal(core.BoolPtr(true)))
				Expect(initializeOptionsModel.Reinitialize).To(Equal(core.BoolPtr(true)))
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
			It(`Invoke NewListDataProductVersionsOptions successfully`, func() {
				// Construct an instance of the ListDataProductVersionsOptions model
				listDataProductVersionsOptionsModel := dpxService.NewListDataProductVersionsOptions()
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
				listDataProductsOptionsModel := dpxService.NewListDataProductsOptions()
				listDataProductsOptionsModel.SetLimit(int64(10))
				listDataProductsOptionsModel.SetStart("testString")
				listDataProductsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listDataProductsOptionsModel).ToNot(BeNil())
				Expect(listDataProductsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(10))))
				Expect(listDataProductsOptionsModel.Start).To(Equal(core.StringPtr("testString")))
				Expect(listDataProductsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateContractTermsDocumentOptions successfully`, func() {
				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(dpxv1.JSONPatchOperation)
				Expect(jsonPatchOperationModel).ToNot(BeNil())
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = core.StringPtr("testString")
				Expect(jsonPatchOperationModel.Op).To(Equal(core.StringPtr("add")))
				Expect(jsonPatchOperationModel.Path).To(Equal(core.StringPtr("testString")))
				Expect(jsonPatchOperationModel.From).To(Equal(core.StringPtr("testString")))
				Expect(jsonPatchOperationModel.Value).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the UpdateContractTermsDocumentOptions model
				dataProductVersionID := "testString"
				contractTermsID := "testString"
				documentID := "testString"
				jsonPatchInstructions := []dpxv1.JSONPatchOperation{}
				updateContractTermsDocumentOptionsModel := dpxService.NewUpdateContractTermsDocumentOptions(dataProductVersionID, contractTermsID, documentID, jsonPatchInstructions)
				updateContractTermsDocumentOptionsModel.SetDataProductVersionID("testString")
				updateContractTermsDocumentOptionsModel.SetContractTermsID("testString")
				updateContractTermsDocumentOptionsModel.SetDocumentID("testString")
				updateContractTermsDocumentOptionsModel.SetJSONPatchInstructions([]dpxv1.JSONPatchOperation{*jsonPatchOperationModel})
				updateContractTermsDocumentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateContractTermsDocumentOptionsModel).ToNot(BeNil())
				Expect(updateContractTermsDocumentOptionsModel.DataProductVersionID).To(Equal(core.StringPtr("testString")))
				Expect(updateContractTermsDocumentOptionsModel.ContractTermsID).To(Equal(core.StringPtr("testString")))
				Expect(updateContractTermsDocumentOptionsModel.DocumentID).To(Equal(core.StringPtr("testString")))
				Expect(updateContractTermsDocumentOptionsModel.JSONPatchInstructions).To(Equal([]dpxv1.JSONPatchOperation{*jsonPatchOperationModel}))
				Expect(updateContractTermsDocumentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateDataProductVersionOptions successfully`, func() {
				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(dpxv1.JSONPatchOperation)
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
				jsonPatchInstructions := []dpxv1.JSONPatchOperation{}
				updateDataProductVersionOptionsModel := dpxService.NewUpdateDataProductVersionOptions(id, jsonPatchInstructions)
				updateDataProductVersionOptionsModel.SetID("testString")
				updateDataProductVersionOptionsModel.SetJSONPatchInstructions([]dpxv1.JSONPatchOperation{*jsonPatchOperationModel})
				updateDataProductVersionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateDataProductVersionOptionsModel).ToNot(BeNil())
				Expect(updateDataProductVersionOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(updateDataProductVersionOptionsModel.JSONPatchInstructions).To(Equal([]dpxv1.JSONPatchOperation{*jsonPatchOperationModel}))
				Expect(updateDataProductVersionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUseCase successfully`, func() {
				id := "testString"
				_model, err := dpxService.NewUseCase(id)
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
