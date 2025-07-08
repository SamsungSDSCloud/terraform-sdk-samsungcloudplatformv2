# \ResourcemanagerV1ResourceGroupsAPIsAPI

All URIs are relative to *https://resourcemanager.kr-west1.dev2.samsungsdscloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateResourceGroup**](ResourcemanagerV1ResourceGroupsAPIsAPI.md#CreateResourceGroup) | **Post** /v1/resource-groups | CreateResourceGroup
[**DeleteResourceGroup**](ResourcemanagerV1ResourceGroupsAPIsAPI.md#DeleteResourceGroup) | **Delete** /v1/resource-groups/{resource_group_id} | DeleteResourceGroup
[**DeleteResourceGroups**](ResourcemanagerV1ResourceGroupsAPIsAPI.md#DeleteResourceGroups) | **Delete** /v1/resource-groups | DeleteResourceGroups
[**ListResourceGroups**](ResourcemanagerV1ResourceGroupsAPIsAPI.md#ListResourceGroups) | **Get** /v1/resource-groups | ListResourceGroups
[**SetResourceGroup**](ResourcemanagerV1ResourceGroupsAPIsAPI.md#SetResourceGroup) | **Put** /v1/resource-groups/{resource_group_id} | SetResourceGroup
[**ShowResourceGroup**](ResourcemanagerV1ResourceGroupsAPIsAPI.md#ShowResourceGroup) | **Get** /v1/resource-groups/{resource_group_id} | ShowResourceGroup
[**ShowResourceGroupResources**](ResourcemanagerV1ResourceGroupsAPIsAPI.md#ShowResourceGroupResources) | **Post** /v1/resource-groups/resources | ShowResourceGroupResources
[**ShowResourceGroupResourcesByGroupId**](ResourcemanagerV1ResourceGroupsAPIsAPI.md#ShowResourceGroupResourcesByGroupId) | **Get** /v1/resource-groups/{resource_group_id}/resources | ShowResourceGroupResources



## CreateResourceGroup

> ResourceGroupCreateResponse CreateResourceGroup(ctx).ResourceGroupCreateRequest(resourceGroupCreateRequest).Execute()

CreateResourceGroup



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
	resourceGroupCreateRequest := *openapiclient.NewResourceGroupCreateRequest("Description_example", "Name_example") // ResourceGroupCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourcemanagerV1ResourceGroupsAPIsAPI.CreateResourceGroup(context.Background()).ResourceGroupCreateRequest(resourceGroupCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcemanagerV1ResourceGroupsAPIsAPI.CreateResourceGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateResourceGroup`: ResourceGroupCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `ResourcemanagerV1ResourceGroupsAPIsAPI.CreateResourceGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateResourceGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resourceGroupCreateRequest** | [**ResourceGroupCreateRequest**](ResourceGroupCreateRequest.md) |  | 

### Return type

[**ResourceGroupCreateResponse**](ResourceGroupCreateResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteResourceGroup

> DeleteResourceGroup(ctx, resourceGroupId).Execute()

DeleteResourceGroup



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
	resourceGroupId := "resourceGroupId_example" // string | 리소스 그룹 ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ResourcemanagerV1ResourceGroupsAPIsAPI.DeleteResourceGroup(context.Background(), resourceGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcemanagerV1ResourceGroupsAPIsAPI.DeleteResourceGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceGroupId** | **string** | 리소스 그룹 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteResourceGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteResourceGroups

> DeleteResourceGroups(ctx).ResourceGroupsDeleteRequest(resourceGroupsDeleteRequest).Execute()

DeleteResourceGroups



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
	resourceGroupsDeleteRequest := *openapiclient.NewResourceGroupsDeleteRequest() // ResourceGroupsDeleteRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ResourcemanagerV1ResourceGroupsAPIsAPI.DeleteResourceGroups(context.Background()).ResourceGroupsDeleteRequest(resourceGroupsDeleteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcemanagerV1ResourceGroupsAPIsAPI.DeleteResourceGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteResourceGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resourceGroupsDeleteRequest** | [**ResourceGroupsDeleteRequest**](ResourceGroupsDeleteRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListResourceGroups

> ResourceGroupPageResponse ListResourceGroups(ctx).Size(size).Page(page).Sort(sort).Id(id).Name(name).Execute()

ListResourceGroups



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
	size := int32(20) // int32 | size (optional)
	page := int32(0) // int32 | page (optional)
	sort := "created_at:desc" // string | sort (optional)
	id := "id_example" // string | 리소스 그룹 ID (optional)
	name := "name_example" // string | 리소스 그룹명 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourcemanagerV1ResourceGroupsAPIsAPI.ListResourceGroups(context.Background()).Size(size).Page(page).Sort(sort).Id(id).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcemanagerV1ResourceGroupsAPIsAPI.ListResourceGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListResourceGroups`: ResourceGroupPageResponse
	fmt.Fprintf(os.Stdout, "Response from `ResourcemanagerV1ResourceGroupsAPIsAPI.ListResourceGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListResourceGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **int32** | size | 
 **page** | **int32** | page | 
 **sort** | **string** | sort | 
 **id** | **string** | 리소스 그룹 ID | 
 **name** | **string** | 리소스 그룹명 | 

### Return type

[**ResourceGroupPageResponse**](ResourceGroupPageResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetResourceGroup

> ResourceGroupCreateResponse SetResourceGroup(ctx, resourceGroupId).ResourceGroupCreateRequest(resourceGroupCreateRequest).Execute()

SetResourceGroup



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
	resourceGroupId := "resourceGroupId_example" // string | 리소스 그룹 ID
	resourceGroupCreateRequest := *openapiclient.NewResourceGroupCreateRequest("Description_example", "Name_example") // ResourceGroupCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourcemanagerV1ResourceGroupsAPIsAPI.SetResourceGroup(context.Background(), resourceGroupId).ResourceGroupCreateRequest(resourceGroupCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcemanagerV1ResourceGroupsAPIsAPI.SetResourceGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetResourceGroup`: ResourceGroupCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `ResourcemanagerV1ResourceGroupsAPIsAPI.SetResourceGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceGroupId** | **string** | 리소스 그룹 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetResourceGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **resourceGroupCreateRequest** | [**ResourceGroupCreateRequest**](ResourceGroupCreateRequest.md) |  | 

### Return type

[**ResourceGroupCreateResponse**](ResourceGroupCreateResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowResourceGroup

> ResourceGroupShowResponse ShowResourceGroup(ctx, resourceGroupId).Execute()

ShowResourceGroup



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
	resourceGroupId := "resourceGroupId_example" // string | 리소스 그룹 ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourcemanagerV1ResourceGroupsAPIsAPI.ShowResourceGroup(context.Background(), resourceGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcemanagerV1ResourceGroupsAPIsAPI.ShowResourceGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowResourceGroup`: ResourceGroupShowResponse
	fmt.Fprintf(os.Stdout, "Response from `ResourcemanagerV1ResourceGroupsAPIsAPI.ShowResourceGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceGroupId** | **string** | 리소스 그룹 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowResourceGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResourceGroupShowResponse**](ResourceGroupShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowResourceGroupResources

> ResourcePageResponse ShowResourceGroupResources(ctx).ResourceSearchRequest(resourceSearchRequest).Size(size).Page(page).Sort(sort).Execute()

ShowResourceGroupResources



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
	resourceSearchRequest := *openapiclient.NewResourceSearchRequest("Tags_example") // ResourceSearchRequest | 
	size := int32(20) // int32 | size (optional)
	page := int32(0) // int32 | page (optional)
	sort := "created_at:desc" // string | sort (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourcemanagerV1ResourceGroupsAPIsAPI.ShowResourceGroupResources(context.Background()).ResourceSearchRequest(resourceSearchRequest).Size(size).Page(page).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcemanagerV1ResourceGroupsAPIsAPI.ShowResourceGroupResources``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowResourceGroupResources`: ResourcePageResponse
	fmt.Fprintf(os.Stdout, "Response from `ResourcemanagerV1ResourceGroupsAPIsAPI.ShowResourceGroupResources`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiShowResourceGroupResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resourceSearchRequest** | [**ResourceSearchRequest**](ResourceSearchRequest.md) |  | 
 **size** | **int32** | size | 
 **page** | **int32** | page | 
 **sort** | **string** | sort | 

### Return type

[**ResourcePageResponse**](ResourcePageResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowResourceGroupResourcesByGroupId

> ResourcePageResponse ShowResourceGroupResourcesByGroupId(ctx, resourceGroupId).Size(size).Page(page).Sort(sort).Execute()

ShowResourceGroupResources



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
	resourceGroupId := "resourceGroupId_example" // string | 리소스 그룹 ID
	size := int32(20) // int32 | size (optional)
	page := int32(0) // int32 | page (optional)
	sort := "created_at:desc" // string | sort (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourcemanagerV1ResourceGroupsAPIsAPI.ShowResourceGroupResourcesByGroupId(context.Background(), resourceGroupId).Size(size).Page(page).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcemanagerV1ResourceGroupsAPIsAPI.ShowResourceGroupResourcesByGroupId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowResourceGroupResourcesByGroupId`: ResourcePageResponse
	fmt.Fprintf(os.Stdout, "Response from `ResourcemanagerV1ResourceGroupsAPIsAPI.ShowResourceGroupResourcesByGroupId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceGroupId** | **string** | 리소스 그룹 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowResourceGroupResourcesByGroupIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **size** | **int32** | size | 
 **page** | **int32** | page | 
 **sort** | **string** | sort | 

### Return type

[**ResourcePageResponse**](ResourcePageResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

