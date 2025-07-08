# \SkeV1RolesApiAPI

All URIs are relative to *https://scp-kubernetes-a.kr-west1.dev2.samsungsdscloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteRole**](SkeV1RolesApiAPI.md#DeleteRole) | **Delete** /v1/clusters/{cluster_id}/namespaces/{namespace_name}/roles | Delete Role
[**ListRoles**](SkeV1RolesApiAPI.md#ListRoles) | **Get** /v1/clusters/{cluster_id}/namespaces/{namespace_name}/roles | List Roles
[**ShowRole**](SkeV1RolesApiAPI.md#ShowRole) | **Get** /v1/clusters/{cluster_id}/namespaces/{namespace_name}/roles/{role_name} | Show Role



## DeleteRole

> DeleteRole(ctx, clusterId, namespaceName).Execute()

Delete Role



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SkeV1RolesApiAPI.DeleteRole(context.Background(), clusterId, namespaceName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SkeV1RolesApiAPI.DeleteRole``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## ListRoles

> RoleListResponse ListRoles(ctx, clusterId, namespaceName).Execute()

List Roles



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SkeV1RolesApiAPI.ListRoles(context.Background(), clusterId, namespaceName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SkeV1RolesApiAPI.ListRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListRoles`: RoleListResponse
	fmt.Fprintf(os.Stdout, "Response from `SkeV1RolesApiAPI.ListRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 
**namespaceName** | **string** | Namespace Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RoleListResponse**](RoleListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowRole

> RoleShowResponse ShowRole(ctx, clusterId, namespaceName, roleName).Execute()

Show Role



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
	roleName := "roleName_example" // string | Role Name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SkeV1RolesApiAPI.ShowRole(context.Background(), clusterId, namespaceName, roleName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SkeV1RolesApiAPI.ShowRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowRole`: RoleShowResponse
	fmt.Fprintf(os.Stdout, "Response from `SkeV1RolesApiAPI.ShowRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 
**namespaceName** | **string** | Namespace Name | 
**roleName** | **string** | Role Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**RoleShowResponse**](RoleShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

