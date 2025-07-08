# \NetworkLoggingV1NetworkLoggingConfigurationApiAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListNetworkLoggingConfigurations**](NetworkLoggingV1NetworkLoggingConfigurationApiAPI.md#ListNetworkLoggingConfigurations) | **Get** /v1/network-logging/configurations | List Network Logging Configurations



## ListNetworkLoggingConfigurations

> NetworkLoggingConfigurationListResponse ListNetworkLoggingConfigurations(ctx).ResourceType(resourceType).WithCount(withCount).Limit(limit).Marker(marker).Sort(sort).ResourceId(resourceId).ResourceName(resourceName).Execute()

List Network Logging Configurations



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	resourceType := openapiclient.NetworkLoggingResourceType("FIREWALL") // NetworkLoggingResourceType | Resource Type
	withCount := "true" // string | with count (optional)
	limit := int32(20) // int32 | limit (optional)
	marker := "607e0938521643b5b4b266f343fae693" // string | marker (optional)
	sort := "created_at:desc" // string | sort (optional)
	resourceId := "resourceId_example" // string | Resource ID (optional)
	resourceName := "resourceName_example" // string | Resource Name (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkLoggingV1NetworkLoggingConfigurationApiAPI.ListNetworkLoggingConfigurations(context.Background()).ResourceType(resourceType).WithCount(withCount).Limit(limit).Marker(marker).Sort(sort).ResourceId(resourceId).ResourceName(resourceName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkLoggingV1NetworkLoggingConfigurationApiAPI.ListNetworkLoggingConfigurations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListNetworkLoggingConfigurations`: NetworkLoggingConfigurationListResponse
	fmt.Fprintf(os.Stdout, "Response from `NetworkLoggingV1NetworkLoggingConfigurationApiAPI.ListNetworkLoggingConfigurations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNetworkLoggingConfigurationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resourceType** | [**NetworkLoggingResourceType**](NetworkLoggingResourceType.md) | Resource Type | 
 **withCount** | **string** | with count | 
 **limit** | **int32** | limit | 
 **marker** | **string** | marker | 
 **sort** | **string** | sort | 
 **resourceId** | **string** | Resource ID | 
 **resourceName** | **string** | Resource Name | 

### Return type

[**NetworkLoggingConfigurationListResponse**](NetworkLoggingConfigurationListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

