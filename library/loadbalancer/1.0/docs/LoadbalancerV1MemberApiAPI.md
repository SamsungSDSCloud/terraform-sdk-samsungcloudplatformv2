# \LoadbalancerV1MemberApiAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddLbServerGroupMembers**](LoadbalancerV1MemberApiAPI.md#AddLbServerGroupMembers) | **Post** /v1/lb-server-groups/{lb_server_group_id}/members | AddLbServerGroupMembers
[**CheckMemberExistsInServerGroup**](LoadbalancerV1MemberApiAPI.md#CheckMemberExistsInServerGroup) | **Get** /v1/lb-server-groups/{lb_server_group_id}/members/check-member-exists | CheckMemberExistsInServerGroup
[**ListLbServerGroupMembers**](LoadbalancerV1MemberApiAPI.md#ListLbServerGroupMembers) | **Get** /v1/lb-server-groups/{lb_server_group_id}/members | ListLbServerGroupMembers
[**RemoveLbServerGroupMember**](LoadbalancerV1MemberApiAPI.md#RemoveLbServerGroupMember) | **Delete** /v1/lb-server-groups/{lb_server_group_id}/members/{member_id} | RemoveLbServerGroupMember
[**SetLbServerGroupMember**](LoadbalancerV1MemberApiAPI.md#SetLbServerGroupMember) | **Put** /v1/lb-server-groups/{lb_server_group_id}/members/{member_id} | SetLbServerGroupMember
[**ShowLbServerGroupMember**](LoadbalancerV1MemberApiAPI.md#ShowLbServerGroupMember) | **Get** /v1/lb-server-groups/{lb_server_group_id}/members/{member_id} | ShowLbServerGroupMember



## AddLbServerGroupMembers

> MemberListResponse AddLbServerGroupMembers(ctx, lbServerGroupId).MemberListCreateRequest(memberListCreateRequest).Execute()

AddLbServerGroupMembers



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
	lbServerGroupId := "lbServerGroupId_example" // string | ID of Lb Server Group
	memberListCreateRequest := *openapiclient.NewMemberListCreateRequest([]openapiclient.MemberCreateRequest{*openapiclient.NewMemberCreateRequest("192.168.0.1", int32(80), "virtualserver-1", "VM")}) // MemberListCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoadbalancerV1MemberApiAPI.AddLbServerGroupMembers(context.Background(), lbServerGroupId).MemberListCreateRequest(memberListCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadbalancerV1MemberApiAPI.AddLbServerGroupMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddLbServerGroupMembers`: MemberListResponse
	fmt.Fprintf(os.Stdout, "Response from `LoadbalancerV1MemberApiAPI.AddLbServerGroupMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lbServerGroupId** | **string** | ID of Lb Server Group | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddLbServerGroupMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **memberListCreateRequest** | [**MemberListCreateRequest**](MemberListCreateRequest.md) |  | 

### Return type

[**MemberListResponse**](MemberListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckMemberExistsInServerGroup

> CheckMemberExistsResponse CheckMemberExistsInServerGroup(ctx, lbServerGroupId).MemberIp(memberIp).MemberPort(memberPort).Execute()

CheckMemberExistsInServerGroup



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
	lbServerGroupId := "lbServerGroupId_example" // string | ID of Lb Server Group
	memberIp := "memberIp_example" // string | IP of Member
	memberPort := int32(56) // int32 | Protocol port of Member

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoadbalancerV1MemberApiAPI.CheckMemberExistsInServerGroup(context.Background(), lbServerGroupId).MemberIp(memberIp).MemberPort(memberPort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadbalancerV1MemberApiAPI.CheckMemberExistsInServerGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckMemberExistsInServerGroup`: CheckMemberExistsResponse
	fmt.Fprintf(os.Stdout, "Response from `LoadbalancerV1MemberApiAPI.CheckMemberExistsInServerGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lbServerGroupId** | **string** | ID of Lb Server Group | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckMemberExistsInServerGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **memberIp** | **string** | IP of Member | 
 **memberPort** | **int32** | Protocol port of Member | 

### Return type

[**CheckMemberExistsResponse**](CheckMemberExistsResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLbServerGroupMembers

> MemberWithHealthStateListResponse ListLbServerGroupMembers(ctx, lbServerGroupId).Size(size).Page(page).Sort(sort).Name(name).MemberIp(memberIp).MemberPort(memberPort).Execute()

ListLbServerGroupMembers



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
	lbServerGroupId := "lbServerGroupId_example" // string | ID of Lb Server Group
	size := int32(20) // int32 | size (optional)
	page := int32(0) // int32 | page (optional)
	sort := "created_at:desc" // string | sort (optional)
	name := "name_example" // string | Name of Member (optional)
	memberIp := "memberIp_example" // string | IP of Member (optional)
	memberPort := int32(56) // int32 | Protocol port of Member (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoadbalancerV1MemberApiAPI.ListLbServerGroupMembers(context.Background(), lbServerGroupId).Size(size).Page(page).Sort(sort).Name(name).MemberIp(memberIp).MemberPort(memberPort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadbalancerV1MemberApiAPI.ListLbServerGroupMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListLbServerGroupMembers`: MemberWithHealthStateListResponse
	fmt.Fprintf(os.Stdout, "Response from `LoadbalancerV1MemberApiAPI.ListLbServerGroupMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lbServerGroupId** | **string** | ID of Lb Server Group | 

### Other Parameters

Other parameters are passed through a pointer to a apiListLbServerGroupMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **size** | **int32** | size | 
 **page** | **int32** | page | 
 **sort** | **string** | sort | 
 **name** | **string** | Name of Member | 
 **memberIp** | **string** | IP of Member | 
 **memberPort** | **int32** | Protocol port of Member | 

### Return type

[**MemberWithHealthStateListResponse**](MemberWithHealthStateListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveLbServerGroupMember

> RemoveLbServerGroupMember(ctx, lbServerGroupId, memberId).Execute()

RemoveLbServerGroupMember



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
	lbServerGroupId := "lbServerGroupId_example" // string | ID of Lb Server Group
	memberId := "memberId_example" // string | ID of Member

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LoadbalancerV1MemberApiAPI.RemoveLbServerGroupMember(context.Background(), lbServerGroupId, memberId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadbalancerV1MemberApiAPI.RemoveLbServerGroupMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lbServerGroupId** | **string** | ID of Lb Server Group | 
**memberId** | **string** | ID of Member | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveLbServerGroupMemberRequest struct via the builder pattern


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


## SetLbServerGroupMember

> MemberShowResponse SetLbServerGroupMember(ctx, lbServerGroupId, memberId).MemberSetRequest(memberSetRequest).Execute()

SetLbServerGroupMember



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
	lbServerGroupId := "lbServerGroupId_example" // string | ID of Lb Server Group
	memberId := "memberId_example" // string | ID of Member
	memberSetRequest := *openapiclient.NewMemberSetRequest(*openapiclient.NewMemberSet()) // MemberSetRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoadbalancerV1MemberApiAPI.SetLbServerGroupMember(context.Background(), lbServerGroupId, memberId).MemberSetRequest(memberSetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadbalancerV1MemberApiAPI.SetLbServerGroupMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetLbServerGroupMember`: MemberShowResponse
	fmt.Fprintf(os.Stdout, "Response from `LoadbalancerV1MemberApiAPI.SetLbServerGroupMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lbServerGroupId** | **string** | ID of Lb Server Group | 
**memberId** | **string** | ID of Member | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetLbServerGroupMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **memberSetRequest** | [**MemberSetRequest**](MemberSetRequest.md) |  | 

### Return type

[**MemberShowResponse**](MemberShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowLbServerGroupMember

> MemberShowResponse ShowLbServerGroupMember(ctx, lbServerGroupId, memberId).Execute()

ShowLbServerGroupMember



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
	lbServerGroupId := "lbServerGroupId_example" // string | ID of Lb Server Group
	memberId := "memberId_example" // string | ID of Member

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoadbalancerV1MemberApiAPI.ShowLbServerGroupMember(context.Background(), lbServerGroupId, memberId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadbalancerV1MemberApiAPI.ShowLbServerGroupMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowLbServerGroupMember`: MemberShowResponse
	fmt.Fprintf(os.Stdout, "Response from `LoadbalancerV1MemberApiAPI.ShowLbServerGroupMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lbServerGroupId** | **string** | ID of Lb Server Group | 
**memberId** | **string** | ID of Member | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowLbServerGroupMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**MemberShowResponse**](MemberShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

