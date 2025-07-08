# \LoadbalancerV1LBServerGroupsApiAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLbServerGroup**](LoadbalancerV1LBServerGroupsApiAPI.md#CreateLbServerGroup) | **Post** /v1/lb-server-groups | CreateLBServerGroup
[**DeleteLbServerGroup**](LoadbalancerV1LBServerGroupsApiAPI.md#DeleteLbServerGroup) | **Delete** /v1/lb-server-groups/{lb_server_group_id} | DeleteLBServerGroup
[**ListLbServerGroups**](LoadbalancerV1LBServerGroupsApiAPI.md#ListLbServerGroups) | **Get** /v1/lb-server-groups | ListLBServerGroups
[**SetLbServerGroup**](LoadbalancerV1LBServerGroupsApiAPI.md#SetLbServerGroup) | **Put** /v1/lb-server-groups/{lb_server_group_id} | SetLBServerGroup
[**ShowLbServerGroup**](LoadbalancerV1LBServerGroupsApiAPI.md#ShowLbServerGroup) | **Get** /v1/lb-server-groups/{lb_server_group_id} | ShowLBServerGroup



## CreateLbServerGroup

> LbServerGroupShowResponse CreateLbServerGroup(ctx).LbServerGroupCreateRequest(lbServerGroupCreateRequest).Execute()

CreateLBServerGroup



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
	lbServerGroupCreateRequest := *openapiclient.NewLbServerGroupCreateRequest(*openapiclient.NewLbServerGroupCreate(openapiclient.LbServerGroupLbMethod("ROUND_ROBIN"), "ServerGroup01", openapiclient.LbServerGroupProtocol("TCP"), "60fba45cb6c811efba41ba92e4fe7200", "8acceeb6920c4fc494490d864f67f0b5")) // LbServerGroupCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoadbalancerV1LBServerGroupsApiAPI.CreateLbServerGroup(context.Background()).LbServerGroupCreateRequest(lbServerGroupCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadbalancerV1LBServerGroupsApiAPI.CreateLbServerGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateLbServerGroup`: LbServerGroupShowResponse
	fmt.Fprintf(os.Stdout, "Response from `LoadbalancerV1LBServerGroupsApiAPI.CreateLbServerGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateLbServerGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lbServerGroupCreateRequest** | [**LbServerGroupCreateRequest**](LbServerGroupCreateRequest.md) |  | 

### Return type

[**LbServerGroupShowResponse**](LbServerGroupShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLbServerGroup

> DeleteLbServerGroup(ctx, lbServerGroupId).Execute()

DeleteLBServerGroup



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
	lbServerGroupId := "46c68101-8e33-4530-85ca-7c8db54e0076" // string | LB Server Group ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LoadbalancerV1LBServerGroupsApiAPI.DeleteLbServerGroup(context.Background(), lbServerGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadbalancerV1LBServerGroupsApiAPI.DeleteLbServerGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lbServerGroupId** | **string** | LB Server Group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLbServerGroupRequest struct via the builder pattern


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


## ListLbServerGroups

> LbServerGroupListResponse ListLbServerGroups(ctx).Size(size).Page(page).Sort(sort).Name(name).Protocol(protocol).SubnetId(subnetId).VpcId(vpcId).LbHealthCheckId(lbHealthCheckId).LbMethod(lbMethod).Execute()

ListLBServerGroups



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
	name := "ServerGroup01" // string | LB Server Group name (optional)
	protocol := *openapiclient.NewProtocol1() // Protocol1 | Protocol (optional)
	subnetId := "60fba45cb6c811efba41ba92e4fe7200" // string | Service Subnet ID (optional)
	vpcId := "8acceeb6920c4fc494490d864f67f0b5" // string | VPC ID (optional)
	lbHealthCheckId := "46c681018e33453085ca7c8db54e0076" // string | LB Health Check ID (optional)
	lbMethod := *openapiclient.NewLbMethod() // LbMethod | LB Method (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoadbalancerV1LBServerGroupsApiAPI.ListLbServerGroups(context.Background()).Size(size).Page(page).Sort(sort).Name(name).Protocol(protocol).SubnetId(subnetId).VpcId(vpcId).LbHealthCheckId(lbHealthCheckId).LbMethod(lbMethod).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadbalancerV1LBServerGroupsApiAPI.ListLbServerGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListLbServerGroups`: LbServerGroupListResponse
	fmt.Fprintf(os.Stdout, "Response from `LoadbalancerV1LBServerGroupsApiAPI.ListLbServerGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListLbServerGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **int32** | size | 
 **page** | **int32** | page | 
 **sort** | **string** | sort | 
 **name** | **string** | LB Server Group name | 
 **protocol** | [**Protocol1**](Protocol1.md) | Protocol | 
 **subnetId** | **string** | Service Subnet ID | 
 **vpcId** | **string** | VPC ID | 
 **lbHealthCheckId** | **string** | LB Health Check ID | 
 **lbMethod** | [**LbMethod**](LbMethod.md) | LB Method | 

### Return type

[**LbServerGroupListResponse**](LbServerGroupListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetLbServerGroup

> LbServerGroupShowResponse SetLbServerGroup(ctx, lbServerGroupId).LbServerGroupSetRequest(lbServerGroupSetRequest).Execute()

SetLBServerGroup



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
	lbServerGroupId := "46c68101-8e33-4530-85ca-7c8db54e0076" // string | LB Server Group ID
	lbServerGroupSetRequest := *openapiclient.NewLbServerGroupSetRequest(*openapiclient.NewLbServerGroupSet()) // LbServerGroupSetRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoadbalancerV1LBServerGroupsApiAPI.SetLbServerGroup(context.Background(), lbServerGroupId).LbServerGroupSetRequest(lbServerGroupSetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadbalancerV1LBServerGroupsApiAPI.SetLbServerGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetLbServerGroup`: LbServerGroupShowResponse
	fmt.Fprintf(os.Stdout, "Response from `LoadbalancerV1LBServerGroupsApiAPI.SetLbServerGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lbServerGroupId** | **string** | LB Server Group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetLbServerGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **lbServerGroupSetRequest** | [**LbServerGroupSetRequest**](LbServerGroupSetRequest.md) |  | 

### Return type

[**LbServerGroupShowResponse**](LbServerGroupShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowLbServerGroup

> LbServerGroupShowResponse ShowLbServerGroup(ctx, lbServerGroupId).Execute()

ShowLBServerGroup



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
	lbServerGroupId := "46c68101-8e33-4530-85ca-7c8db54e0076" // string | LB Server Group ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoadbalancerV1LBServerGroupsApiAPI.ShowLbServerGroup(context.Background(), lbServerGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadbalancerV1LBServerGroupsApiAPI.ShowLbServerGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowLbServerGroup`: LbServerGroupShowResponse
	fmt.Fprintf(os.Stdout, "Response from `LoadbalancerV1LBServerGroupsApiAPI.ShowLbServerGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lbServerGroupId** | **string** | LB Server Group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowLbServerGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LbServerGroupShowResponse**](LbServerGroupShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

