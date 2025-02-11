// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armhealthdataaiservices

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// DeidServicesClient contains the methods for the DeidServices group.
// Don't use this type directly, use NewDeidServicesClient() instead.
type DeidServicesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewDeidServicesClient creates a new instance of DeidServicesClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewDeidServicesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DeidServicesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &DeidServicesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreate - Create a DeidService
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-09-20
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - deidServiceName - The name of the deid service
//   - resource - Resource create parameters.
//   - options - DeidServicesClientBeginCreateOptions contains the optional parameters for the DeidServicesClient.BeginCreate
//     method.
func (client *DeidServicesClient) BeginCreate(ctx context.Context, resourceGroupName string, deidServiceName string, resource DeidService, options *DeidServicesClientBeginCreateOptions) (*runtime.Poller[DeidServicesClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, deidServiceName, resource, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[DeidServicesClientCreateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[DeidServicesClientCreateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Create - Create a DeidService
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-09-20
func (client *DeidServicesClient) create(ctx context.Context, resourceGroupName string, deidServiceName string, resource DeidService, options *DeidServicesClientBeginCreateOptions) (*http.Response, error) {
	var err error
	const operationName = "DeidServicesClient.BeginCreate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createCreateRequest(ctx, resourceGroupName, deidServiceName, resource, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createCreateRequest creates the Create request.
func (client *DeidServicesClient) createCreateRequest(ctx context.Context, resourceGroupName string, deidServiceName string, resource DeidService, _ *DeidServicesClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HealthDataAIServices/deidServices/{deidServiceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if deidServiceName == "" {
		return nil, errors.New("parameter deidServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deidServiceName}", url.PathEscape(deidServiceName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-09-20")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, resource); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Delete a DeidService
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-09-20
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - deidServiceName - The name of the deid service
//   - options - DeidServicesClientBeginDeleteOptions contains the optional parameters for the DeidServicesClient.BeginDelete
//     method.
func (client *DeidServicesClient) BeginDelete(ctx context.Context, resourceGroupName string, deidServiceName string, options *DeidServicesClientBeginDeleteOptions) (*runtime.Poller[DeidServicesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, deidServiceName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[DeidServicesClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[DeidServicesClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Delete a DeidService
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-09-20
func (client *DeidServicesClient) deleteOperation(ctx context.Context, resourceGroupName string, deidServiceName string, options *DeidServicesClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "DeidServicesClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, deidServiceName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *DeidServicesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, deidServiceName string, _ *DeidServicesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HealthDataAIServices/deidServices/{deidServiceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if deidServiceName == "" {
		return nil, errors.New("parameter deidServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deidServiceName}", url.PathEscape(deidServiceName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-09-20")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get a DeidService
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-09-20
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - deidServiceName - The name of the deid service
//   - options - DeidServicesClientGetOptions contains the optional parameters for the DeidServicesClient.Get method.
func (client *DeidServicesClient) Get(ctx context.Context, resourceGroupName string, deidServiceName string, options *DeidServicesClientGetOptions) (DeidServicesClientGetResponse, error) {
	var err error
	const operationName = "DeidServicesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, deidServiceName, options)
	if err != nil {
		return DeidServicesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DeidServicesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DeidServicesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *DeidServicesClient) getCreateRequest(ctx context.Context, resourceGroupName string, deidServiceName string, _ *DeidServicesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HealthDataAIServices/deidServices/{deidServiceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if deidServiceName == "" {
		return nil, errors.New("parameter deidServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deidServiceName}", url.PathEscape(deidServiceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-09-20")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DeidServicesClient) getHandleResponse(resp *http.Response) (DeidServicesClientGetResponse, error) {
	result := DeidServicesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DeidService); err != nil {
		return DeidServicesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - List DeidService resources by resource group
//
// Generated from API version 2024-09-20
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - DeidServicesClientListByResourceGroupOptions contains the optional parameters for the DeidServicesClient.NewListByResourceGroupPager
//     method.
func (client *DeidServicesClient) NewListByResourceGroupPager(resourceGroupName string, options *DeidServicesClientListByResourceGroupOptions) *runtime.Pager[DeidServicesClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[DeidServicesClientListByResourceGroupResponse]{
		More: func(page DeidServicesClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DeidServicesClientListByResourceGroupResponse) (DeidServicesClientListByResourceGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "DeidServicesClient.NewListByResourceGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			}, nil)
			if err != nil {
				return DeidServicesClientListByResourceGroupResponse{}, err
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *DeidServicesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, _ *DeidServicesClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HealthDataAIServices/deidServices"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-09-20")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *DeidServicesClient) listByResourceGroupHandleResponse(resp *http.Response) (DeidServicesClientListByResourceGroupResponse, error) {
	result := DeidServicesClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DeidServiceListResult); err != nil {
		return DeidServicesClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - List DeidService resources by subscription ID
//
// Generated from API version 2024-09-20
//   - options - DeidServicesClientListBySubscriptionOptions contains the optional parameters for the DeidServicesClient.NewListBySubscriptionPager
//     method.
func (client *DeidServicesClient) NewListBySubscriptionPager(options *DeidServicesClientListBySubscriptionOptions) *runtime.Pager[DeidServicesClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[DeidServicesClientListBySubscriptionResponse]{
		More: func(page DeidServicesClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DeidServicesClientListBySubscriptionResponse) (DeidServicesClientListBySubscriptionResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "DeidServicesClient.NewListBySubscriptionPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listBySubscriptionCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return DeidServicesClientListBySubscriptionResponse{}, err
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *DeidServicesClient) listBySubscriptionCreateRequest(ctx context.Context, _ *DeidServicesClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.HealthDataAIServices/deidServices"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-09-20")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *DeidServicesClient) listBySubscriptionHandleResponse(resp *http.Response) (DeidServicesClientListBySubscriptionResponse, error) {
	result := DeidServicesClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DeidServiceListResult); err != nil {
		return DeidServicesClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Update a DeidService
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-09-20
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - deidServiceName - The name of the deid service
//   - properties - The resource properties to be updated.
//   - options - DeidServicesClientBeginUpdateOptions contains the optional parameters for the DeidServicesClient.BeginUpdate
//     method.
func (client *DeidServicesClient) BeginUpdate(ctx context.Context, resourceGroupName string, deidServiceName string, properties DeidUpdate, options *DeidServicesClientBeginUpdateOptions) (*runtime.Poller[DeidServicesClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, deidServiceName, properties, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[DeidServicesClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[DeidServicesClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Update - Update a DeidService
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-09-20
func (client *DeidServicesClient) update(ctx context.Context, resourceGroupName string, deidServiceName string, properties DeidUpdate, options *DeidServicesClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "DeidServicesClient.BeginUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, deidServiceName, properties, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// updateCreateRequest creates the Update request.
func (client *DeidServicesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, deidServiceName string, properties DeidUpdate, _ *DeidServicesClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HealthDataAIServices/deidServices/{deidServiceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if deidServiceName == "" {
		return nil, errors.New("parameter deidServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deidServiceName}", url.PathEscape(deidServiceName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-09-20")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, properties); err != nil {
		return nil, err
	}
	return req, nil
}
