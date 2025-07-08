# \SkeV1RoleBindingsApiAPI

All URIs are relative to *https://scp-kubernetes-a.kr-west1.dev2.samsungsdscloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteRoleBinding**](SkeV1RoleBindingsApiAPI.md#DeleteRoleBinding) | **Delete** /v1/clusters/{cluster_id}/namespaces/{namespace_name}/rolebindings | Delete Role Binding
[**ListRoleBindings**](SkeV1RoleBindingsApiAPI.md#ListRoleBindings) | **Get** /v1/clusters/{cluster_id}/namespaces/{namespace_name}/rolebindings | List Role Bindings
[**ShowRoleBinding**](SkeV1RoleBindingsApiAPI.md#ShowRoleBinding) | **Get** /v1/clusters/{cluster_id}/namespaces/{namespace_name}/rolebindings/{role_binding_name} | Show Role Binding



## DeleteRoleBinding

> DeleteRoleBinding(ctx, clusterId, namespaceName).Name(name).Execute()

Delete Role Binding



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
	clusterId := "clusterId_example" // string | Cluster ID
	namespaceName := "namespaceName_example" // string | Namespace Name
	name := *openapiclient.NewName1() // Name1 | Names

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SkeV1RoleBindingsApiAPI.DeleteRoleBinding(context.Background(), clusterId, namespaceName).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SkeV1RoleBindingsApiAPI.DeleteRoleBinding``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 
**namespaceName** | **string** | Namespace Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRoleBindingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **name** | [**Name1**](Name1.md) | Names | 

### Return type

 (empty response body)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRoleBindings

> RoleBindingListResponse ListRoleBindings(ctx, clusterId, namespaceName).Size(size).Page(page).Sort(sort).Name(name).SystemObject(systemObject).Execute()

List Role Bindings



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
	clusterId := "clusterId_example" // string | Cluster ID
	namespaceName := "namespaceName_example" // string | Namespace Name
	size := int32(20) // int32 | size (optional)
	page := int32(0) // int32 | page (optional)
	sort := "created_at:desc" // string | sort (optional)
	name := "name_example" // string | Role Binding Name (optional)
	systemObject := true // bool | System Object (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SkeV1RoleBindingsApiAPI.ListRoleBindings(context.Background(), clusterId, namespaceName).Size(size).Page(page).Sort(sort).Name(name).SystemObject(systemObject).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SkeV1RoleBindingsApiAPI.ListRoleBindings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListRoleBindings`: RoleBindingListResponse
	fmt.Fprintf(os.Stdout, "Response from `SkeV1RoleBindingsApiAPI.ListRoleBindings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 
**namespaceName** | **string** | Namespace Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRoleBindingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **size** | **int32** | size | 
 **page** | **int32** | page | 
 **sort** | **string** | sort | 
 **name** | **string** | Role Binding Name | 
 **systemObject** | **bool** | System Object | 

### Return type

[**RoleBindingListResponse**](RoleBindingListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowRoleBinding

> RoleBindingShowResponse ShowRoleBinding(ctx, clusterId, namespaceName, roleBindingName).Execute()

Show Role Binding



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
	clusterId := "clusterId_example" // string | Cluster ID
	namespaceName := "namespaceName_example" // string | Namespace Name
	roleBindingName := "roleBindingName_example" // string | Role Binding Name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SkeV1RoleBindingsApiAPI.ShowRoleBinding(context.Background(), clusterId, namespaceName, roleBindingName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SkeV1RoleBindingsApiAPI.ShowRoleBinding``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowRoleBinding`: RoleBindingShowResponse
	fmt.Fprintf(os.Stdout, "Response from `SkeV1RoleBindingsApiAPI.ShowRoleBinding`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 
**namespaceName** | **string** | Namespace Name | 
**roleBindingName** | **string** | Role Binding Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowRoleBindingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**RoleBindingShowResponse**](RoleBindingShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

