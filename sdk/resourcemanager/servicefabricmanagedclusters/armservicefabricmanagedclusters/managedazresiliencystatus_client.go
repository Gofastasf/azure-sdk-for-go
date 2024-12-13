//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armservicefabricmanagedclusters

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

// ManagedAzResiliencyStatusClient contains the methods for the ManagedAzResiliencyStatus group.
// Don't use this type directly, use NewManagedAzResiliencyStatusClient() instead.
type ManagedAzResiliencyStatusClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewManagedAzResiliencyStatusClient creates a new instance of ManagedAzResiliencyStatusClient with the specified values.
//   - subscriptionID - The customer subscription identifier.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewManagedAzResiliencyStatusClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ManagedAzResiliencyStatusClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ManagedAzResiliencyStatusClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Action to get Az Resiliency Status of all the Base resources constituting Service Fabric Managed Clusters.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-09-01-preview
//   - resourceGroupName - The name of the resource group.
//   - clusterName - The name of the cluster resource.
//   - options - ManagedAzResiliencyStatusClientGetOptions contains the optional parameters for the ManagedAzResiliencyStatusClient.Get
//     method.
func (client *ManagedAzResiliencyStatusClient) Get(ctx context.Context, resourceGroupName string, clusterName string, options *ManagedAzResiliencyStatusClientGetOptions) (ManagedAzResiliencyStatusClientGetResponse, error) {
	var err error
	const operationName = "ManagedAzResiliencyStatusClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, clusterName, options)
	if err != nil {
		return ManagedAzResiliencyStatusClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ManagedAzResiliencyStatusClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ManagedAzResiliencyStatusClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ManagedAzResiliencyStatusClient) getCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, options *ManagedAzResiliencyStatusClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/managedClusters/{clusterName}/getazresiliencystatus"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ManagedAzResiliencyStatusClient) getHandleResponse(resp *http.Response) (ManagedAzResiliencyStatusClientGetResponse, error) {
	result := ManagedAzResiliencyStatusClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagedAzResiliencyStatus); err != nil {
		return ManagedAzResiliencyStatusClientGetResponse{}, err
	}
	return result, nil
}
