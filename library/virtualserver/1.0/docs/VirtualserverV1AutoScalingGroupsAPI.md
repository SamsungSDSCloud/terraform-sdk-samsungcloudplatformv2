# \VirtualserverV1AutoScalingGroupsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAutoScalingGroup**](VirtualserverV1AutoScalingGroupsAPI.md#CreateAutoScalingGroup) | **Post** /v1/auto-scaling-groups | Create Auto-Scaling Group
[**DeleteAutoScalingGroup**](VirtualserverV1AutoScalingGroupsAPI.md#DeleteAutoScalingGroup) | **Delete** /v1/auto-scaling-groups/{auto_scaling_group_id} | Delete Auto-Scaling Group
[**ListAutoScalingGroups**](VirtualserverV1AutoScalingGroupsAPI.md#ListAutoScalingGroups) | **Get** /v1/auto-scaling-groups | List Auto-Scaling Groups
[**ShowAutoScalingGroup**](VirtualserverV1AutoScalingGroupsAPI.md#ShowAutoScalingGroup) | **Get** /v1/auto-scaling-groups/{auto_scaling_group_id} | Show Auto-Scaling Group
[**UpdateAutoScalingGroup**](VirtualserverV1AutoScalingGroupsAPI.md#UpdateAutoScalingGroup) | **Put** /v1/auto-scaling-groups/{auto_scaling_group_id} | Update Auto-Scaling Group
[**UpdateAutoScalingGroupServerCount**](VirtualserverV1AutoScalingGroupsAPI.md#UpdateAutoScalingGroupServerCount) | **Put** /v1/auto-scaling-groups/{auto_scaling_group_id}/server-count | Update Auto-Scaling Group server count



## CreateAutoScalingGroup

> AutoScalingGroupShowResponse CreateAutoScalingGroup(ctx).AutoScalingGroupCreateRequest(autoScalingGroupCreateRequest).Execute()

Create Auto-Scaling Group



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
	autoScalingGroupCreateRequest := *openapiclient.NewAutoScalingGroupCreateRequest(int32(1), true, "b5aea5a675fc4f6b9e0fcd1288354c5f", int32(2), int32(1), "auto-scaling-group-name", "server-name-prefix", []string{"SubnetIds_example"}) // AutoScalingGroupCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualserverV1AutoScalingGroupsAPI.CreateAutoScalingGroup(context.Background()).AutoScalingGroupCreateRequest(autoScalingGroupCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1AutoScalingGroupsAPI.CreateAutoScalingGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAutoScalingGroup`: AutoScalingGroupShowResponse
	fmt.Fprintf(os.Stdout, "Response from `VirtualserverV1AutoScalingGroupsAPI.CreateAutoScalingGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAutoScalingGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **autoScalingGroupCreateRequest** | [**AutoScalingGroupCreateRequest**](AutoScalingGroupCreateRequest.md) |  | 

### Return type

[**AutoScalingGroupShowResponse**](AutoScalingGroupShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAutoScalingGroup

> DeleteAutoScalingGroup(ctx, autoScalingGroupId).Execute()

Delete Auto-Scaling Group



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
	autoScalingGroupId := "52613bd852b04b39adcb15a8364d856d" // string | Auto-Scaling Group ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VirtualserverV1AutoScalingGroupsAPI.DeleteAutoScalingGroup(context.Background(), autoScalingGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1AutoScalingGroupsAPI.DeleteAutoScalingGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**autoScalingGroupId** | **string** | Auto-Scaling Group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAutoScalingGroupRequest struct via the builder pattern


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


## ListAutoScalingGroups

> AutoScalingGroupListResponse ListAutoScalingGroups(ctx).Name(name).State(state).LaunchConfigurationId(launchConfigurationId).LaunchConfigurationName(launchConfigurationName).LbServerGroupId(lbServerGroupId).SecurityGroupId(securityGroupId).SubnetId(subnetId).VpcId(vpcId).CreatedBy(createdBy).Offset(offset).Limit(limit).Sort(sort).Execute()

List Auto-Scaling Groups



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
	name := "auto-scaling-group-name" // string | Auto-Scaling Group name (optional)
	state := "IN_SERVICE" // string | Auto-Scaling Group state (optional)
	launchConfigurationId := "b5aea5a675fc4f6b9e0fcd1288354c5f" // string | Launch Configuration ID (optional)
	launchConfigurationName := "launch-configuration-name" // string | Launch Configuration name (optional)
	lbServerGroupId := "d06e87d3-ca9a-461b-8e88-077a542a7335" // string | LB Server Group ID (optional)
	securityGroupId := "c09c3f05-03d9-443f-b27a-40e0f973c75f" // string | Security Group ID (optional)
	subnetId := "78b2ee3f074145c580ab6abac4821f97" // string | Subnet ID (optional)
	vpcId := "cc976b621087484ea5fd527f4b78708b" // string | VPC ID (optional)
	createdBy := "97e6b22c9a4143789ca522df7457a32f" // string | Created by (optional)
	offset := int32(0) // int32 | Offset (optional)
	limit := int32(20) // int32 | Limit (optional)
	sort := "created_at:desc" // string | Sort (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualserverV1AutoScalingGroupsAPI.ListAutoScalingGroups(context.Background()).Name(name).State(state).LaunchConfigurationId(launchConfigurationId).LaunchConfigurationName(launchConfigurationName).LbServerGroupId(lbServerGroupId).SecurityGroupId(securityGroupId).SubnetId(subnetId).VpcId(vpcId).CreatedBy(createdBy).Offset(offset).Limit(limit).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1AutoScalingGroupsAPI.ListAutoScalingGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAutoScalingGroups`: AutoScalingGroupListResponse
	fmt.Fprintf(os.Stdout, "Response from `VirtualserverV1AutoScalingGroupsAPI.ListAutoScalingGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAutoScalingGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | Auto-Scaling Group name | 
 **state** | **string** | Auto-Scaling Group state | 
 **launchConfigurationId** | **string** | Launch Configuration ID | 
 **launchConfigurationName** | **string** | Launch Configuration name | 
 **lbServerGroupId** | **string** | LB Server Group ID | 
 **securityGroupId** | **string** | Security Group ID | 
 **subnetId** | **string** | Subnet ID | 
 **vpcId** | **string** | VPC ID | 
 **createdBy** | **string** | Created by | 
 **offset** | **int32** | Offset | 
 **limit** | **int32** | Limit | 
 **sort** | **string** | Sort | 

### Return type

[**AutoScalingGroupListResponse**](AutoScalingGroupListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowAutoScalingGroup

> AutoScalingGroupShowResponse ShowAutoScalingGroup(ctx, autoScalingGroupId).Execute()

Show Auto-Scaling Group



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
	autoScalingGroupId := "52613bd852b04b39adcb15a8364d856d" // string | Auto-Scaling Group ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualserverV1AutoScalingGroupsAPI.ShowAutoScalingGroup(context.Background(), autoScalingGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1AutoScalingGroupsAPI.ShowAutoScalingGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowAutoScalingGroup`: AutoScalingGroupShowResponse
	fmt.Fprintf(os.Stdout, "Response from `VirtualserverV1AutoScalingGroupsAPI.ShowAutoScalingGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**autoScalingGroupId** | **string** | Auto-Scaling Group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowAutoScalingGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AutoScalingGroupShowResponse**](AutoScalingGroupShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAutoScalingGroup

> AutoScalingGroupShowResponse UpdateAutoScalingGroup(ctx, autoScalingGroupId).AutoScalingGroupSetRequest(autoScalingGroupSetRequest).Execute()

Update Auto-Scaling Group



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
	autoScalingGroupId := "52613bd852b04b39adcb15a8364d856d" // string | Auto-Scaling Group ID
	autoScalingGroupSetRequest := *openapiclient.NewAutoScalingGroupSetRequest() // AutoScalingGroupSetRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualserverV1AutoScalingGroupsAPI.UpdateAutoScalingGroup(context.Background(), autoScalingGroupId).AutoScalingGroupSetRequest(autoScalingGroupSetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1AutoScalingGroupsAPI.UpdateAutoScalingGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAutoScalingGroup`: AutoScalingGroupShowResponse
	fmt.Fprintf(os.Stdout, "Response from `VirtualserverV1AutoScalingGroupsAPI.UpdateAutoScalingGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**autoScalingGroupId** | **string** | Auto-Scaling Group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAutoScalingGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **autoScalingGroupSetRequest** | [**AutoScalingGroupSetRequest**](AutoScalingGroupSetRequest.md) |  | 

### Return type

[**AutoScalingGroupShowResponse**](AutoScalingGroupShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAutoScalingGroupServerCount

> AutoScalingGroupShowResponse UpdateAutoScalingGroupServerCount(ctx, autoScalingGroupId).AutoScalingGroupSetServerCountRequest(autoScalingGroupSetServerCountRequest).Execute()

Update Auto-Scaling Group server count



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
	autoScalingGroupId := "52613bd852b04b39adcb15a8364d856d" // string | Auto-Scaling Group ID
	autoScalingGroupSetServerCountRequest := *openapiclient.NewAutoScalingGroupSetServerCountRequest() // AutoScalingGroupSetServerCountRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualserverV1AutoScalingGroupsAPI.UpdateAutoScalingGroupServerCount(context.Background(), autoScalingGroupId).AutoScalingGroupSetServerCountRequest(autoScalingGroupSetServerCountRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1AutoScalingGroupsAPI.UpdateAutoScalingGroupServerCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAutoScalingGroupServerCount`: AutoScalingGroupShowResponse
	fmt.Fprintf(os.Stdout, "Response from `VirtualserverV1AutoScalingGroupsAPI.UpdateAutoScalingGroupServerCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**autoScalingGroupId** | **string** | Auto-Scaling Group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAutoScalingGroupServerCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **autoScalingGroupSetServerCountRequest** | [**AutoScalingGroupSetServerCountRequest**](AutoScalingGroupSetServerCountRequest.md) |  | 

### Return type

[**AutoScalingGroupShowResponse**](AutoScalingGroupShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

