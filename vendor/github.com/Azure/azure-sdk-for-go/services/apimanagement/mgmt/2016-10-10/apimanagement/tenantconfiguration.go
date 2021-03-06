package apimanagement

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"net/http"
)

// TenantConfigurationClient is the apiManagement Client
type TenantConfigurationClient struct {
	ManagementClient
}

// NewTenantConfigurationClient creates an instance of the TenantConfigurationClient client.
func NewTenantConfigurationClient(subscriptionID string) TenantConfigurationClient {
	return NewTenantConfigurationClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewTenantConfigurationClientWithBaseURI creates an instance of the TenantConfigurationClient client.
func NewTenantConfigurationClientWithBaseURI(baseURI string, subscriptionID string) TenantConfigurationClient {
	return TenantConfigurationClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Deploy this operation applies changes from the specified Git branch to the configuration database. This is a long
// running operation and could take several minutes to complete. This method may poll for completion. Polling can be
// canceled by passing the cancel channel argument. The channel will be used to cancel polling and any outstanding HTTP
// requests.
//
// resourceGroupName is the name of the resource group. serviceName is the name of the API Management service.
// parameters is deploy Configuration parameters.
func (client TenantConfigurationClient) Deploy(resourceGroupName string, serviceName string, parameters DeployConfigurationParameters, cancel <-chan struct{}) (<-chan OperationResultContract, <-chan error) {
	resultChan := make(chan OperationResultContract, 1)
	errChan := make(chan error, 1)
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}},
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.Branch", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		errChan <- validation.NewErrorWithValidationError(err, "apimanagement.TenantConfigurationClient", "Deploy")
		close(errChan)
		close(resultChan)
		return resultChan, errChan
	}

	go func() {
		var err error
		var result OperationResultContract
		defer func() {
			if err != nil {
				errChan <- err
			}
			resultChan <- result
			close(resultChan)
			close(errChan)
		}()
		req, err := client.DeployPreparer(resourceGroupName, serviceName, parameters, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "apimanagement.TenantConfigurationClient", "Deploy", nil, "Failure preparing request")
			return
		}

		resp, err := client.DeploySender(req)
		if err != nil {
			result.Response = autorest.Response{Response: resp}
			err = autorest.NewErrorWithError(err, "apimanagement.TenantConfigurationClient", "Deploy", resp, "Failure sending request")
			return
		}

		result, err = client.DeployResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "apimanagement.TenantConfigurationClient", "Deploy", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// DeployPreparer prepares the Deploy request.
func (client TenantConfigurationClient) DeployPreparer(resourceGroupName string, serviceName string, parameters DeployConfigurationParameters, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-10-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/tenant/configuration/deploy", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// DeploySender sends the Deploy request. The method will close the
// http.Response Body if it receives an error.
func (client TenantConfigurationClient) DeploySender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client),
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// DeployResponder handles the response to the Deploy request. The method always
// closes the http.Response Body.
func (client TenantConfigurationClient) DeployResponder(resp *http.Response) (result OperationResultContract, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Save this operation creates a commit with the current configuration snapshot to the specified branch in the
// repository. This is a long running operation and could take several minutes to complete. This method may poll for
// completion. Polling can be canceled by passing the cancel channel argument. The channel will be used to cancel
// polling and any outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group. serviceName is the name of the API Management service.
// parameters is save Configuration parameters.
func (client TenantConfigurationClient) Save(resourceGroupName string, serviceName string, parameters SaveConfigurationParameter, cancel <-chan struct{}) (<-chan OperationResultContract, <-chan error) {
	resultChan := make(chan OperationResultContract, 1)
	errChan := make(chan error, 1)
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}},
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.Branch", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		errChan <- validation.NewErrorWithValidationError(err, "apimanagement.TenantConfigurationClient", "Save")
		close(errChan)
		close(resultChan)
		return resultChan, errChan
	}

	go func() {
		var err error
		var result OperationResultContract
		defer func() {
			if err != nil {
				errChan <- err
			}
			resultChan <- result
			close(resultChan)
			close(errChan)
		}()
		req, err := client.SavePreparer(resourceGroupName, serviceName, parameters, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "apimanagement.TenantConfigurationClient", "Save", nil, "Failure preparing request")
			return
		}

		resp, err := client.SaveSender(req)
		if err != nil {
			result.Response = autorest.Response{Response: resp}
			err = autorest.NewErrorWithError(err, "apimanagement.TenantConfigurationClient", "Save", resp, "Failure sending request")
			return
		}

		result, err = client.SaveResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "apimanagement.TenantConfigurationClient", "Save", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// SavePreparer prepares the Save request.
func (client TenantConfigurationClient) SavePreparer(resourceGroupName string, serviceName string, parameters SaveConfigurationParameter, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-10-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/tenant/configuration/save", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// SaveSender sends the Save request. The method will close the
// http.Response Body if it receives an error.
func (client TenantConfigurationClient) SaveSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client),
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// SaveResponder handles the response to the Save request. The method always
// closes the http.Response Body.
func (client TenantConfigurationClient) SaveResponder(resp *http.Response) (result OperationResultContract, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Validate this operation validates the changes in the specified Git branch. This is a long running operation and
// could take several minutes to complete. This method may poll for completion. Polling can be canceled by passing the
// cancel channel argument. The channel will be used to cancel polling and any outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group. serviceName is the name of the API Management service.
// parameters is validate Configuration parameters.
func (client TenantConfigurationClient) Validate(resourceGroupName string, serviceName string, parameters DeployConfigurationParameters, cancel <-chan struct{}) (<-chan OperationResultContract, <-chan error) {
	resultChan := make(chan OperationResultContract, 1)
	errChan := make(chan error, 1)
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}},
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.Branch", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		errChan <- validation.NewErrorWithValidationError(err, "apimanagement.TenantConfigurationClient", "Validate")
		close(errChan)
		close(resultChan)
		return resultChan, errChan
	}

	go func() {
		var err error
		var result OperationResultContract
		defer func() {
			if err != nil {
				errChan <- err
			}
			resultChan <- result
			close(resultChan)
			close(errChan)
		}()
		req, err := client.ValidatePreparer(resourceGroupName, serviceName, parameters, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "apimanagement.TenantConfigurationClient", "Validate", nil, "Failure preparing request")
			return
		}

		resp, err := client.ValidateSender(req)
		if err != nil {
			result.Response = autorest.Response{Response: resp}
			err = autorest.NewErrorWithError(err, "apimanagement.TenantConfigurationClient", "Validate", resp, "Failure sending request")
			return
		}

		result, err = client.ValidateResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "apimanagement.TenantConfigurationClient", "Validate", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// ValidatePreparer prepares the Validate request.
func (client TenantConfigurationClient) ValidatePreparer(resourceGroupName string, serviceName string, parameters DeployConfigurationParameters, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-10-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/tenant/configuration/validate", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// ValidateSender sends the Validate request. The method will close the
// http.Response Body if it receives an error.
func (client TenantConfigurationClient) ValidateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client),
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// ValidateResponder handles the response to the Validate request. The method always
// closes the http.Response Body.
func (client TenantConfigurationClient) ValidateResponder(resp *http.Response) (result OperationResultContract, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
