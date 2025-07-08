# \BaremetalBlockstorageV1VolumeGroupV1APIsAPI

All URIs are relative to *http://70.50.166.119:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddVolumeGroupMembers**](BaremetalBlockstorageV1VolumeGroupV1APIsAPI.md#AddVolumeGroupMembers) | **Put** /v1/volume-groups/{volume_group_id}/members/add | Add Volume Group Member on Volume Group
[**CreateVolumeGroup**](BaremetalBlockstorageV1VolumeGroupV1APIsAPI.md#CreateVolumeGroup) | **Post** /v1/volume-groups | Create Volume Group
[**ListVolumeGroups**](BaremetalBlockstorageV1VolumeGroupV1APIsAPI.md#ListVolumeGroups) | **Get** /v1/volume-groups | List Volume Groups
[**RemoveVolumeGroupMembers**](BaremetalBlockstorageV1VolumeGroupV1APIsAPI.md#RemoveVolumeGroupMembers) | **Put** /v1/volume-groups/{volume_group_id}/members/remove | Remove Volume Group Member on Volume Group
[**ShowVolumeGroup**](BaremetalBlockstorageV1VolumeGroupV1APIsAPI.md#ShowVolumeGroup) | **Get** /v1/volume-groups/{volume_group_id} | Show Volume Group



## AddVolumeGroupMembers

> VolumeGroupMemberResponse AddVolumeGroupMembers(ctx, volumeGroupId).VolumeGroupMemberRequest(volumeGroupMemberRequest).Execute()

Add Volume Group Member on Volume Group



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
	volumeGroupId := "c9ecae8b973b425c81c3817893cd8063" // string | Volume group id
	volumeGroupMemberRequest := *openapiclient.NewVolumeGroupMemberRequest([]string{"VolumeIds_example"}) // VolumeGroupMemberRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BaremetalBlockstorageV1VolumeGroupV1APIsAPI.AddVolumeGroupMembers(context.Background(), volumeGroupId).VolumeGroupMemberRequest(volumeGroupMemberRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BaremetalBlockstorageV1VolumeGroupV1APIsAPI.AddVolumeGroupMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddVolumeGroupMembers`: VolumeGroupMemberResponse
	fmt.Fprintf(os.Stdout, "Response from `BaremetalBlockstorageV1VolumeGroupV1APIsAPI.AddVolumeGroupMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeGroupId** | **string** | Volume group id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddVolumeGroupMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **volumeGroupMemberRequest** | [**VolumeGroupMemberRequest**](VolumeGroupMemberRequest.md) |  | 

### Return type

[**VolumeGroupMemberResponse**](VolumeGroupMemberResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateVolumeGroup

> AsyncResponse CreateVolumeGroup(ctx).VolumeGroupCreationRequest(volumeGroupCreationRequest).Execute()

Create Volume Group



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
	volumeGroupCreationRequest := *openapiclient.NewVolumeGroupCreationRequest("cg-name01", []string{"VolumeIds_example"}) // VolumeGroupCreationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BaremetalBlockstorageV1VolumeGroupV1APIsAPI.CreateVolumeGroup(context.Background()).VolumeGroupCreationRequest(volumeGroupCreationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BaremetalBlockstorageV1VolumeGroupV1APIsAPI.CreateVolumeGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateVolumeGroup`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `BaremetalBlockstorageV1VolumeGroupV1APIsAPI.CreateVolumeGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateVolumeGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **volumeGroupCreationRequest** | [**VolumeGroupCreationRequest**](VolumeGroupCreationRequest.md) |  | 

### Return type

[**AsyncResponse**](AsyncResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVolumeGroups

> VolumeGroupListResponse ListVolumeGroups(ctx).Limit(limit).Offset(offset).Sort(sort).Name(name).Execute()

List Volume Groups



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
	limit := int32(20) // int32 | Number to be displayed on the page (optional) (default to 20)
	offset := int32(0) // int32 | Offset (optional) (default to 0)
	sort := "created_at:desc,name:asc" // string | Sort (optional)
	name := "cg-01" // string | Volume group name (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BaremetalBlockstorageV1VolumeGroupV1APIsAPI.ListVolumeGroups(context.Background()).Limit(limit).Offset(offset).Sort(sort).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BaremetalBlockstorageV1VolumeGroupV1APIsAPI.ListVolumeGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListVolumeGroups`: VolumeGroupListResponse
	fmt.Fprintf(os.Stdout, "Response from `BaremetalBlockstorageV1VolumeGroupV1APIsAPI.ListVolumeGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListVolumeGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Number to be displayed on the page | [default to 20]
 **offset** | **int32** | Offset | [default to 0]
 **sort** | **string** | Sort | 
 **name** | **string** | Volume group name | 

### Return type

[**VolumeGroupListResponse**](VolumeGroupListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveVolumeGroupMembers

> VolumeGroupMemberResponse RemoveVolumeGroupMembers(ctx, volumeGroupId).VolumeGroupMemberRequest(volumeGroupMemberRequest).Execute()

Remove Volume Group Member on Volume Group



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
	volumeGroupId := "c9ecae8b973b425c81c3817893cd8063" // string | Volume group id
	volumeGroupMemberRequest := *openapiclient.NewVolumeGroupMemberRequest([]string{"VolumeIds_example"}) // VolumeGroupMemberRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BaremetalBlockstorageV1VolumeGroupV1APIsAPI.RemoveVolumeGroupMembers(context.Background(), volumeGroupId).VolumeGroupMemberRequest(volumeGroupMemberRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BaremetalBlockstorageV1VolumeGroupV1APIsAPI.RemoveVolumeGroupMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveVolumeGroupMembers`: VolumeGroupMemberResponse
	fmt.Fprintf(os.Stdout, "Response from `BaremetalBlockstorageV1VolumeGroupV1APIsAPI.RemoveVolumeGroupMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeGroupId** | **string** | Volume group id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveVolumeGroupMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **volumeGroupMemberRequest** | [**VolumeGroupMemberRequest**](VolumeGroupMemberRequest.md) |  | 

### Return type

[**VolumeGroupMemberResponse**](VolumeGroupMemberResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowVolumeGroup

> VolumeGroupResponse ShowVolumeGroup(ctx, volumeGroupId).Execute()

Show Volume Group



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
	volumeGroupId := "c9ecae8b973b425c81c3817893cd8063" // string | Volume group id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BaremetalBlockstorageV1VolumeGroupV1APIsAPI.ShowVolumeGroup(context.Background(), volumeGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BaremetalBlockstorageV1VolumeGroupV1APIsAPI.ShowVolumeGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowVolumeGroup`: VolumeGroupResponse
	fmt.Fprintf(os.Stdout, "Response from `BaremetalBlockstorageV1VolumeGroupV1APIsAPI.ShowVolumeGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeGroupId** | **string** | Volume group id | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowVolumeGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VolumeGroupResponse**](VolumeGroupResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

