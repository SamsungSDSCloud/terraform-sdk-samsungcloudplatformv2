# \VirtualserverV1ServerGroupsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateServerGroup**](VirtualserverV1ServerGroupsAPI.md#CreateServerGroup) | **Post** /v1/server-groups | Create Server Group
[**DeleteServerGroup**](VirtualserverV1ServerGroupsAPI.md#DeleteServerGroup) | **Delete** /v1/server-groups/{server_group_id} | Delete Server Group
[**ListServerGroups**](VirtualserverV1ServerGroupsAPI.md#ListServerGroups) | **Get** /v1/server-groups | List Server Groups
[**ShowServerGroup**](VirtualserverV1ServerGroupsAPI.md#ShowServerGroup) | **Get** /v1/server-groups/{server_group_id} | Show Server Group



## CreateServerGroup

> ServerGroup CreateServerGroup(ctx).ServerGroupCreateRequest(serverGroupCreateRequest).Execute()

Create Server Group



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
	serverGroupCreateRequest := *openapiclient.NewServerGroupCreateRequest("test", "anti-affinity") // ServerGroupCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualserverV1ServerGroupsAPI.CreateServerGroup(context.Background()).ServerGroupCreateRequest(serverGroupCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1ServerGroupsAPI.CreateServerGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateServerGroup`: ServerGroup
	fmt.Fprintf(os.Stdout, "Response from `VirtualserverV1ServerGroupsAPI.CreateServerGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateServerGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serverGroupCreateRequest** | [**ServerGroupCreateRequest**](ServerGroupCreateRequest.md) |  | 

### Return type

[**ServerGroup**](ServerGroup.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteServerGroup

> DeleteServerGroup(ctx, serverGroupId).Execute()

Delete Server Group



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
	serverGroupId := "616fb98f-46ca-475e-917e-2563e5a8cd19" // string | Server Group ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VirtualserverV1ServerGroupsAPI.DeleteServerGroup(context.Background(), serverGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1ServerGroupsAPI.DeleteServerGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverGroupId** | **string** | Server Group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteServerGroupRequest struct via the builder pattern


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


## ListServerGroups

> ServerGroupListResponse ListServerGroups(ctx).Limit(limit).Offset(offset).Execute()

List Server Groups



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
	limit := int32(20) // int32 | Limit (optional)
	offset := int32(10) // int32 | Offset (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualserverV1ServerGroupsAPI.ListServerGroups(context.Background()).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1ServerGroupsAPI.ListServerGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListServerGroups`: ServerGroupListResponse
	fmt.Fprintf(os.Stdout, "Response from `VirtualserverV1ServerGroupsAPI.ListServerGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListServerGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit | 
 **offset** | **int32** | Offset | 

### Return type

[**ServerGroupListResponse**](ServerGroupListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowServerGroup

> ServerGroup ShowServerGroup(ctx, serverGroupId).Execute()

Show Server Group



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
	serverGroupId := "616fb98f-46ca-475e-917e-2563e5a8cd19" // string | Server Group ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualserverV1ServerGroupsAPI.ShowServerGroup(context.Background(), serverGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1ServerGroupsAPI.ShowServerGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowServerGroup`: ServerGroup
	fmt.Fprintf(os.Stdout, "Response from `VirtualserverV1ServerGroupsAPI.ShowServerGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverGroupId** | **string** | Server Group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowServerGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServerGroup**](ServerGroup.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

