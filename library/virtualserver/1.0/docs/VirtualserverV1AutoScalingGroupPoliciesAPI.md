# \VirtualserverV1AutoScalingGroupPoliciesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAutoScalingGroupPolicy**](VirtualserverV1AutoScalingGroupPoliciesAPI.md#CreateAutoScalingGroupPolicy) | **Post** /v1/auto-scaling-groups/{auto_scaling_group_id}/policies | Create Auto-Scaling Group Policy
[**DeleteAutoScalingGroupPolicy**](VirtualserverV1AutoScalingGroupPoliciesAPI.md#DeleteAutoScalingGroupPolicy) | **Delete** /v1/auto-scaling-groups/{auto_scaling_group_id}/policies/{policy_id} | Delete Auto-Scaling Group Policy
[**ListAutoScalingGroupPolicies**](VirtualserverV1AutoScalingGroupPoliciesAPI.md#ListAutoScalingGroupPolicies) | **Get** /v1/auto-scaling-groups/{auto_scaling_group_id}/policies | List Auto-Scaling Group Policies
[**ShowAutoScalingGroupPolicy**](VirtualserverV1AutoScalingGroupPoliciesAPI.md#ShowAutoScalingGroupPolicy) | **Get** /v1/auto-scaling-groups/{auto_scaling_group_id}/policies/{policy_id} | Show Auto-Scaling Group Policy
[**UpdateAutoScalingGroupPolicy**](VirtualserverV1AutoScalingGroupPoliciesAPI.md#UpdateAutoScalingGroupPolicy) | **Put** /v1/auto-scaling-groups/{auto_scaling_group_id}/policies/{policy_id} | Update Auto-Scaling Group Policy



## CreateAutoScalingGroupPolicy

> AutoScalingGroupPolicyShowResponse CreateAutoScalingGroupPolicy(ctx, autoScalingGroupId).AutoScalingGroupPolicyCreateRequest(autoScalingGroupPolicyCreateRequest).Execute()

Create Auto-Scaling Group Policy



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
	autoScalingGroupPolicyCreateRequest := *openapiclient.NewAutoScalingGroupPolicyCreateRequest("ge", int32(300), int32(1), "AVG", "CPU", "policy-name", "AMOUNT", "SCALE_OUT", int32(10), int32(60)) // AutoScalingGroupPolicyCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualserverV1AutoScalingGroupPoliciesAPI.CreateAutoScalingGroupPolicy(context.Background(), autoScalingGroupId).AutoScalingGroupPolicyCreateRequest(autoScalingGroupPolicyCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1AutoScalingGroupPoliciesAPI.CreateAutoScalingGroupPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAutoScalingGroupPolicy`: AutoScalingGroupPolicyShowResponse
	fmt.Fprintf(os.Stdout, "Response from `VirtualserverV1AutoScalingGroupPoliciesAPI.CreateAutoScalingGroupPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**autoScalingGroupId** | **string** | Auto-Scaling Group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAutoScalingGroupPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **autoScalingGroupPolicyCreateRequest** | [**AutoScalingGroupPolicyCreateRequest**](AutoScalingGroupPolicyCreateRequest.md) |  | 

### Return type

[**AutoScalingGroupPolicyShowResponse**](AutoScalingGroupPolicyShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAutoScalingGroupPolicy

> DeleteAutoScalingGroupPolicy(ctx, autoScalingGroupId, policyId).Execute()

Delete Auto-Scaling Group Policy



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
	policyId := "0761d76085f54363bab07909baf69841" // string | Policy ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VirtualserverV1AutoScalingGroupPoliciesAPI.DeleteAutoScalingGroupPolicy(context.Background(), autoScalingGroupId, policyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1AutoScalingGroupPoliciesAPI.DeleteAutoScalingGroupPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**autoScalingGroupId** | **string** | Auto-Scaling Group ID | 
**policyId** | **string** | Policy ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAutoScalingGroupPolicyRequest struct via the builder pattern


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


## ListAutoScalingGroupPolicies

> AutoScalingGroupPolicyListResponse ListAutoScalingGroupPolicies(ctx, autoScalingGroupId).Name(name).MetricMethod(metricMethod).MetricType(metricType).ScaleType(scaleType).Offset(offset).Limit(limit).Sort(sort).Execute()

List Auto-Scaling Group Policies



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
	name := "policy-name" // string | Policy name (optional)
	metricMethod := openapiclient.AutoScalingGroupPolicyMetricMethod("AVG") // AutoScalingGroupPolicyMetricMethod | Metric method (optional)
	metricType := "CPU" // string | Metric type (optional)
	scaleType := openapiclient.AutoScalingGroupPolicyScaleType("SCALE_OUT") // AutoScalingGroupPolicyScaleType | Scale type (optional)
	offset := int32(0) // int32 | Offset (optional)
	limit := int32(20) // int32 | Limit (optional)
	sort := "created_at:desc" // string | Sort (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualserverV1AutoScalingGroupPoliciesAPI.ListAutoScalingGroupPolicies(context.Background(), autoScalingGroupId).Name(name).MetricMethod(metricMethod).MetricType(metricType).ScaleType(scaleType).Offset(offset).Limit(limit).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1AutoScalingGroupPoliciesAPI.ListAutoScalingGroupPolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAutoScalingGroupPolicies`: AutoScalingGroupPolicyListResponse
	fmt.Fprintf(os.Stdout, "Response from `VirtualserverV1AutoScalingGroupPoliciesAPI.ListAutoScalingGroupPolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**autoScalingGroupId** | **string** | Auto-Scaling Group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAutoScalingGroupPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** | Policy name | 
 **metricMethod** | [**AutoScalingGroupPolicyMetricMethod**](AutoScalingGroupPolicyMetricMethod.md) | Metric method | 
 **metricType** | **string** | Metric type | 
 **scaleType** | [**AutoScalingGroupPolicyScaleType**](AutoScalingGroupPolicyScaleType.md) | Scale type | 
 **offset** | **int32** | Offset | 
 **limit** | **int32** | Limit | 
 **sort** | **string** | Sort | 

### Return type

[**AutoScalingGroupPolicyListResponse**](AutoScalingGroupPolicyListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowAutoScalingGroupPolicy

> AutoScalingGroupPolicyShowResponse ShowAutoScalingGroupPolicy(ctx, autoScalingGroupId, policyId).Execute()

Show Auto-Scaling Group Policy



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
	policyId := "0761d76085f54363bab07909baf69841" // string | Policy ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualserverV1AutoScalingGroupPoliciesAPI.ShowAutoScalingGroupPolicy(context.Background(), autoScalingGroupId, policyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1AutoScalingGroupPoliciesAPI.ShowAutoScalingGroupPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowAutoScalingGroupPolicy`: AutoScalingGroupPolicyShowResponse
	fmt.Fprintf(os.Stdout, "Response from `VirtualserverV1AutoScalingGroupPoliciesAPI.ShowAutoScalingGroupPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**autoScalingGroupId** | **string** | Auto-Scaling Group ID | 
**policyId** | **string** | Policy ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowAutoScalingGroupPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AutoScalingGroupPolicyShowResponse**](AutoScalingGroupPolicyShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAutoScalingGroupPolicy

> AutoScalingGroupPolicyShowResponse UpdateAutoScalingGroupPolicy(ctx, autoScalingGroupId, policyId).AutoScalingGroupPolicyUpdateRequest(autoScalingGroupPolicyUpdateRequest).Execute()

Update Auto-Scaling Group Policy



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
	policyId := "0761d76085f54363bab07909baf69841" // string | Policy ID
	autoScalingGroupPolicyUpdateRequest := *openapiclient.NewAutoScalingGroupPolicyUpdateRequest() // AutoScalingGroupPolicyUpdateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualserverV1AutoScalingGroupPoliciesAPI.UpdateAutoScalingGroupPolicy(context.Background(), autoScalingGroupId, policyId).AutoScalingGroupPolicyUpdateRequest(autoScalingGroupPolicyUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1AutoScalingGroupPoliciesAPI.UpdateAutoScalingGroupPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAutoScalingGroupPolicy`: AutoScalingGroupPolicyShowResponse
	fmt.Fprintf(os.Stdout, "Response from `VirtualserverV1AutoScalingGroupPoliciesAPI.UpdateAutoScalingGroupPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**autoScalingGroupId** | **string** | Auto-Scaling Group ID | 
**policyId** | **string** | Policy ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAutoScalingGroupPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **autoScalingGroupPolicyUpdateRequest** | [**AutoScalingGroupPolicyUpdateRequest**](AutoScalingGroupPolicyUpdateRequest.md) |  | 

### Return type

[**AutoScalingGroupPolicyShowResponse**](AutoScalingGroupPolicyShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

